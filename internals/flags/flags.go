package flags

import (
	cfg "github.com/he-man-david/prysm-geth-proxy/internals/configs"
	"github.com/urfave/cli/v2"
)

func NewCliApp() *cli.App {
	cliapp := cli.NewApp()
	cliapp.Name = "prysm-geth-proxy"
	cliapp.Usage = "Proxy router for Prysm Consensus & Geth Execution clients"

	cliapp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "http-host",
			Value: cfg.DefaultHTTPHost,
			Usage: "host interface for the HTTP RPC server",
		},
		&cli.IntFlag{
			Name:  "http-port",
			Value: cfg.DefaultHTTPPort,
			Usage: "TCP port for the HTTP RPC server",
		},
		&cli.StringFlag{
			Name:  "ws-host",
			Value: cfg.DefaultWSHost,
			Usage: "host interface for the websocket RPC server",
		},
		&cli.IntFlag{
			Name:  "ws-port",
			Value: cfg.DefaultWSPort,
			Usage: "TCP port for the websocket RPC server",
		},
		&cli.StringFlag{
			Name:  "auth-host",
			Value: cfg.DefaultAuthHost,
			Usage: "host interface for the authenticated apis",
		},
		&cli.IntFlag{
			Name:  "auth-port",
			Value: cfg.DefaultAuthPort,
			Usage: "port for the authenticated apis",
		},
		&cli.StringSliceFlag{
			Name:  "auth-cors",
			Value: cli.NewStringSlice(cfg.DefaultAuthCors...),
			Usage: "cors domain for the authenticated apis",
		},
		&cli.StringSliceFlag{
			Name:  "auth-vhosts",
			Value: cli.NewStringSlice(cfg.DefaultAuthVhosts...),
			Usage: "virtual hosts for the authenticated apis",
		},
		&cli.StringSliceFlag{
			Name:  "auth-origins",
			Value: cli.NewStringSlice(cfg.DefaultAuthOrigins...),
			Usage: "origins for the authenticated apis",
		},
		&cli.StringFlag{
			Name:  "auth-prefix",
			Value: cfg.DefaultAuthPrefix,
			Usage: "prefix for the authenticated apis",
		},
		&cli.StringSliceFlag{
			Name:  "auth-modules",
			Value: cli.NewStringSlice(cfg.DefaultAuthModules...),
			Usage: "modules for the authenticated apis",
		},
		&cli.StringFlag{
			Name:  "jwt-path",
			Value: cfg.DefaultJWTPath,
			Usage: "path for JWT secret",
		},
	}
	
	return cliapp
}