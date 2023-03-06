package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/he-man-david/prysm-geth-proxy/internals/configs"
	"github.com/he-man-david/prysm-geth-proxy/internals/flags"
	"github.com/he-man-david/prysm-geth-proxy/proxy"
	"github.com/urfave/cli/v2"
)

var app = flags.NewCliApp()

func main() {
	log.Println("Starting PRYSM-GETH proxy server at....")
	app := flags.NewCliApp()

    app.Action = gethProxy

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

func gethProxy(c *cli.Context) error {
	cfg, err := configs.MakeProxyConfig(c)

	fmt.Printf(
	"===========================================================================================================\n" +
	"[PROXY] configuration: \n HTTPPort: %d, \n WSPort: %d, \n AuthPort: %d, \n AuthAddr: %s, "+
    "\n AuthVirtualHosts: %v, \n HTTPModules: %v, \n HTTPVirtualHosts: %v, \n HTTPTimeouts: %v, \n WSModules: %v, "+
    "\n JWTSecret: %s\n" +
	"===========================================================================================================\n",
    cfg.HTTPPort, cfg.WSPort, cfg.AuthPort, cfg.AuthAddr, cfg.AuthVirtualHosts, cfg.HTTPModules,
    cfg.HTTPVirtualHosts, cfg.HTTPTimeouts, cfg.WSModules, cfg.JWTSecret)

	if err != nil {
		log.Printf("[main] failed to create proxy configs :: ERR, %v", err)
		return err
	}
	p, pErr := proxy.New(cfg)
	if pErr != nil {
		log.Printf("[main] failed to create proxy instance :: ERR, %v", err)
		return err
	}

	if err = p.StartRPC(); err != nil {
		log.Printf("[main] failed to start proxy RPC servers :: ERR, %v", err)
		return err
	}
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Shutting down PRYSM-GETH proxy server...")
		p.StopRPC()
		os.Exit(0)
	}()

	select {}
}






