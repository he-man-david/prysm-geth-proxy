package proxy

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/he-man-david/prysm-geth-proxy/http"
	rpcapis "github.com/he-man-david/prysm-geth-proxy/rpc_apis"
	"github.com/he-man-david/prysm-geth-proxy/utils"
)

type Proxy struct {
	config        *node.Config
	log           log.Logger
	rpcAPIs       []rpc.API   // List of APIs currently provided by the node
	http          *http.HttpServer //
	ws            *http.HttpServer //
	httpAuth      *http.HttpServer //
	wsAuth        *http.HttpServer //
	inprocHandler *rpc.Server // In-process RPC request handler to process the API requests
}

func New() (*Proxy, error) {
	conf := 
	p := &Proxy{
		inprocHandler: rpc.NewServer(),
	}

	// register required rpc APIs
	p.rpcAPIs = append(p.rpcAPIs, rpcapis.GetAPIs()...)

	// Check HTTP/WS prefixes are valid.
	if err := validatePrefix("HTTP", conf.HTTPPathPrefix); err != nil {
		return nil, err
	}
	if err := validatePrefix("WebSocket", conf.WSPathPrefix); err != nil {
		return nil, err
	}

	// Configure RPC servers.
	p.http = http.NewHttpServer(p.log, conf.HTTPTimeouts)
	p.httpAuth = http.NewHttpServer(p.log, conf.HTTPTimeouts)
	p.ws = http.NewHttpServer(p.log, rpc.DefaultHTTPTimeouts)
	p.wsAuth = http.NewHttpServer(p.log, rpc.DefaultHTTPTimeouts)

	return p, nil
}

func (p *Proxy) startRPC() error {
	var (
		servers           []*http.HttpServer
		openAPIs, allAPIs = p.getAPIs()
	)

	initHttp := func(server *http.HttpServer, port int) error {
		if err := server.SetListenAddr(p.config.HTTPHost, port); err != nil {
			return err
		}
		if err := server.EnableRPC(openAPIs, http.HttpConfig{
			CorsAllowedOrigins: p.config.HTTPCors,
			Vhosts:             p.config.HTTPVirtualHosts,
			Modules:            p.config.HTTPModules,
			Prefix: 			p.config.HTTPPathPrefix,
		}); err != nil {
			return err
		}
		servers = append(servers, server)
		return nil
	}

	initWS := func(port int) error {
		server := p.wsServerForPort(port, false)
		if err := server.SetListenAddr(p.config.WSHost, port); err != nil {
			return err
		}
		if err := server.EnableWS(openAPIs, http.WsConfig{
			Modules: p.config.WSModules,
			Origins: p.config.WSOrigins,
			Prefix:  p.config.WSPathPrefix,
		}); err != nil {
			return err
		}
		servers = append(servers, server)
		return nil
	}

	initAuth := func(port int, secret []byte) error {
		// Enable auth via HTTP
		server := p.httpAuth
		if err := server.SetListenAddr(p.config.AuthAddr, port); err != nil {
			return err
		}
		if err := server.EnableRPC(allAPIs, http.HttpConfig{
			CorsAllowedOrigins: DefaultAuthCors,
			Vhosts:             p.config.AuthVirtualHosts,
			Modules:            DefaultAuthModules,
			Prefix:             DefaultAuthPrefix,
			JwtSecret:          secret,
		}); err != nil {
			return err
		}
		servers = append(servers, server)
		// Enable auth via WS
		server = p.wsServerForPort(port, true)
		if err := server.SetListenAddr(p.config.AuthAddr, port); err != nil {
			return err
		}
		if err := server.EnableWS(allAPIs, http.WsConfig{
			Modules:   DefaultAuthModules,
			Origins:   DefaultAuthOrigins,
			Prefix:    DefaultAuthPrefix,
			JwtSecret: secret,
		}); err != nil {
			return err
		}
		servers = append(servers, server)
		return nil
	}

	// Set up HTTP.
	if p.config.HTTPHost != "" {
		// Configure legacy unauthenticated HTTP.
		if err := initHttp(p.http, p.config.HTTPPort); err != nil {
			return err
		}
	}
	// Configure WebSocket.
	if p.config.WSHost != "" {
		// legacy unauthenticated
		if err := initWS(p.config.WSPort); err != nil {
			return err
		}
	}
	// Configure authenticated API
	if len(openAPIs) != len(allAPIs) {
		jwtSecret, err := utils.ObtainJWTSecret(p.config.JWTSecret)
		if err != nil {
			return err
		}
		if err := initAuth(p.config.AuthPort, jwtSecret); err != nil {
			return err
		}
	}
	// Start the servers
	for _, server := range servers {
		if err := server.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Proxy) wsServerForPort(port int, authenticated bool) *http.HttpServer {
	httpServer, wsServer := p.http, p.ws
	if authenticated {
		httpServer, wsServer = p.httpAuth, p.wsAuth
	}
	if p.config.HTTPHost == "" || httpServer.Port == port {
		return httpServer
	}
	return wsServer
}

func (p *Proxy) stopRPC() {
	p.http.Stop()
	p.ws.Stop()
	p.httpAuth.Stop()
	p.wsAuth.Stop()
	p.stopInProc()
}

// startInProc registers all RPC APIs on the inproc server.
func (p *Proxy) startInProc(apis []rpc.API) error {
	for _, api := range apis {
		if err := p.inprocHandler.RegisterName(api.Namespace, api.Service); err != nil {
			return err
		}
	}
	return nil
}

// stopInProc terminates the in-process RPC endpoint.
func (p *Proxy) stopInProc() {
	p.inprocHandler.Stop()
}

// getAPIs return two sets of APIs, both the ones that do not require
// authentication, and the complete set
func (p *Proxy) getAPIs() (unauthenticated, all []rpc.API) {
	for _, api := range p.rpcAPIs {
		if !api.Authenticated {
			unauthenticated = append(unauthenticated, api)
		}
	}
	return unauthenticated, p.rpcAPIs
}


// validatePrefix checks if 'path' is a valid configuration value for the RPC prefix option.
func validatePrefix(what, path string) error {
	if path == "" {
		return nil
	}
	if path[0] != '/' {
		return fmt.Errorf(`%s RPC path prefix %q does not contain leading "/"`, what, path)
	}
	if strings.ContainsAny(path, "?#") {
		// This is just to avoid confusion. While these would match correctly (i.e. they'd
		// match if URL-escaped into path), it's not easy to understand for users when
		// setting that on the command line.
		return fmt.Errorf("%s RPC path prefix %q contains URL meta-characters", what, path)
	}
	return nil
}