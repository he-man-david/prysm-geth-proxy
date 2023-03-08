package proxy

import (
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	DefaultHTTPHost    = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort    = 8545        // Default TCP port for the HTTP RPC server
	DefaultWSHost      = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort      = 8546        // Default TCP port for the websocket RPC server
	DefaultGraphQLHost = "localhost" // Default host interface for the GraphQL server
	DefaultGraphQLPort = 8547        // Default TCP port for the GraphQL server
	DefaultAuthHost    = "localhost" // Default host interface for the authenticated apis
	DefaultAuthPort    = 8551        // Default port for the authenticated apis
)

var (
	DefaultAuthCors    = []string{"localhost"} // Default cors domain for the authenticated apis
	DefaultAuthVhosts  = []string{"localhost"} // Default virtual hosts for the authenticated apis
	DefaultAuthOrigins = []string{"localhost"} // Default origins for the authenticated apis
	DefaultAuthPrefix  = ""                    // Default prefix for the authenticated apis
	DefaultAuthModules = []string{"eth", "engine"}
)

var DefaultConfig = node.Config{
	HTTPPort:            DefaultHTTPPort,
	AuthAddr:            DefaultAuthHost,
	AuthPort:            DefaultAuthPort,
	AuthVirtualHosts:    DefaultAuthVhosts,
	HTTPModules:         []string{"net", "web3", "eth"},
	HTTPVirtualHosts:    []string{"localhost"},
	HTTPTimeouts:        rpc.DefaultHTTPTimeouts,
	WSPort:              DefaultWSPort,
	WSModules:           []string{"net", "web3", "eth"},
	GraphQLVirtualHosts: []string{"localhost"},
	DBEngine: "",
}

// makeConfigNode loads geth configuration and creates a blank node instance.
func makeConfigNode() (*node.Config) {
	// Load defaults.
	// HTTPHost
	// HTTPCors
	// HTTPPathPrefix
	// WSPathPrefix
	// HTTPTimeouts
	// HTTPVirtualHosts
	// HTTPModules
	// WSHost
	// WSModules
	// WSOrigins
	// WSPathPrefix
	// AuthAddr
	// AuthVirtualHosts
	// HTTPPort
	// WSPort
	// JWTSecret
	cfg := DefaultConfig

	// utils.SetNodeConfig(ctx, &cfg.Node)
	// stack, err := node.New(&cfg.Node)
	// if err != nil {
	// 	utils.Fatalf("Failed to create the protocol stack: %v", err)
	// }

	// utils.SetEthConfig(ctx, stack, &cfg.Eth)
	// if ctx.IsSet(utils.EthStatsURLFlag.Name) {
	// 	cfg.Ethstats.URL = ctx.String(utils.EthStatsURLFlag.Name)
	// }

	return &cfg
}