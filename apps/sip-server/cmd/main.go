package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/alwaysbespoke/coba/apps/sip-server/internal/clients/k8"
	"github.com/alwaysbespoke/coba/apps/sip-server/internal/clients/sbcs"
	"github.com/alwaysbespoke/coba/pkg/config"

	tcphandler "github.com/alwaysbespoke/coba/apps/sip-server/internal/handlers/tcp"
	udphandler "github.com/alwaysbespoke/coba/apps/sip-server/internal/handlers/udp"
	exitmanager "github.com/alwaysbespoke/coba/pkg/exit-manager"
	tcpserver "github.com/alwaysbespoke/coba/pkg/servers/tcp"
	udpserver "github.com/alwaysbespoke/coba/pkg/servers/udp"
)

func main() {
	/*
	*
	* Setup the logger
	*
	 */
	// create a prod logger instance
	prodLogger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}

	// create a sugared logger instance
	logger := prodLogger.Sugar()
	defer logger.Sync()

	logger.Info("starting service")

	/*
	*
	* Configure the application
	*
	 */
	// create empty instances for the application configs
	udpHandlerCfg := &udphandler.Config{}
	udpServerCfg := &udpserver.Config{}
	tcpHandlerCfg := &tcphandler.Config{}
	tcpServerCfg := &tcpserver.Config{}
	k8Cfg := &k8.Config{}
	sbcsCfg := &sbcs.Config{}

	// process the configs
	cfg := config.New()
	cfg.Set("UDP_HANDLER", udpHandlerCfg)
	cfg.Set("UDP_SERVER", udpServerCfg)
	cfg.Set("TCP_HANDLER", tcpHandlerCfg)
	cfg.Set("TCP_SERVER", tcpServerCfg)
	cfg.Set("K8", k8Cfg)
	cfg.Set("SBCS", sbcsCfg)
	cfg.Run()

	/*
	*
	* Create the application clients
	*
	 */

	// create the K8 clients
	k8Clients := k8.New(k8Cfg, logger)

	// create an ExitManager instance
	exitMgr, ctx := exitmanager.New()
	go exitMgr.Run()

	// create the SBCs client
	sbcsClient := sbcs.New(&sbcs.Input{
		Ctx:           ctx,
		Config:        &sbcs.Config{},
		Logger:        logger,
		SbcV1Client:   k8Clients.SbcV1Client,
		SbcV1Informer: k8Clients.SbcInformer,
	})

	/*
	*
	* Create and run the APIs and controllers
	*
	 */
	// run the SBC controller
	go sbcsClient.RunController()

	// create the UDP handler
	udpHandler := udphandler.New(udpHandlerCfg, logger)

	// create the UDP server
	udpServer := udpserver.New(udpServerCfg, logger, udpHandler)

	// run the UDP server
	go func() {
		err := udpServer.Run()
		logger.Errorf("server failure: %w", err)
	}()

	// create the TCP handler
	tcpHandler := tcphandler.New(tcpHandlerCfg, logger)

	// create the TCP server
	tcpServer := tcpserver.New(tcpServerCfg, logger, tcpHandler)

	// run the TCP server
	go func() {
		err := tcpServer.Run()
		logger.Errorf("server failure: %w", err)
	}()

	// wait until graceful shutdown has resolved before exiting the application
	exitMgr.Wait()

	logger.Info("closing service")
}
