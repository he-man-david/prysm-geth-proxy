package configs

import (
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"
)

const (
	DefaultHTTPHost    = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort    = 8545        // Default TCP port for the HTTP RPC server
	DefaultWSHost      = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort      = 8546        // Default TCP port for the websocket RPC server
	DefaultAuthHost    = "localhost" // Default host interface for the authenticated apis
	DefaultAuthPort    = 8551        // Default port for the authenticated apis
)

var (
	DefaultAuthCors    = []string{"localhost"} // Default cors domain for the authenticated apis
	DefaultAuthVhosts  = []string{"localhost"} // Default virtual hosts for the authenticated apis
	DefaultAuthOrigins = []string{"localhost"} // Default origins for the authenticated apis
	DefaultAuthPrefix  = ""                    // Default prefix for the authenticated apis
	DefaultHTTPPAthPrefix = ""
	DefaultAuthModules = []string{"eth", "engine"}
	DefaultJWTPath     = "/Users/david/workspace/learnings/prysm-geth-proxy/jwt.hex"
)

var DefaultConfig = node.Config{
	HTTPPort:            DefaultHTTPPort,
	WSPort:              DefaultWSPort,
	AuthPort:            DefaultAuthPort,
	AuthAddr:            DefaultAuthHost,
	AuthVirtualHosts:    DefaultAuthVhosts,
	HTTPModules:         []string{"eth"},
	HTTPVirtualHosts:    []string{"localhost"},
	HTTPTimeouts:        rpc.DefaultHTTPTimeouts,
	WSModules:           []string{"eth"},
	JWTSecret:   		 DefaultJWTPath,
	HTTPPathPrefix:		 DefaultHTTPPAthPrefix,
	WSPathPrefix:		 DefaultHTTPPAthPrefix,

}

// makeProxyConfig loads geth configuration and creates a blank node instance.
func MakeProxyConfig(ctx *cli.Context) (*node.Config, error) {
	cfg := setFromFlags(ctx, &DefaultConfig)
	
	return cfg, nil
}

func setFromFlags(ctx *cli.Context, cfg *node.Config) (*node.Config) {
	if ctx.IsSet("http-host") {
		cfg.HTTPHost = ctx.String("http-host")
	}
	if ctx.IsSet("http-port") {
		cfg.HTTPPort = ctx.Int("http-port")
	}
	if ctx.IsSet("ws-host") {
		cfg.WSHost = ctx.String("ws-host")
	}
	if ctx.IsSet("ws-port") {
		cfg.WSPort = ctx.Int("ws-port")
	}
	if ctx.IsSet("auth-host") {
		cfg.AuthAddr = ctx.String("auth-host")
	}
	if ctx.IsSet("auth-port") {
		cfg.AuthPort = ctx.Int("auth-port")
	}
	if ctx.IsSet("auth-vhosts") {
		cfg.AuthVirtualHosts = ctx.StringSlice("auth-vhosts")
	}
	if ctx.IsSet("jwt-path") {
		cfg.JWTSecret = ctx.String("jwt-path")
	}

	return cfg
}