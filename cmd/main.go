package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/alwaysbespoke/coba/internal/apis/udp"
	"github.com/alwaysbespoke/coba/internal/clients/k8"
	"github.com/alwaysbespoke/coba/internal/clients/sbcs"
	"github.com/alwaysbespoke/coba/internal/config"

	exitmanager "github.com/alwaysbespoke/coba/internal/exit-manager"
)

func main() {
	// create a logger instance
	prodLogger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}
	logger := prodLogger.Sugar()
	defer logger.Sync()

	logger.Info("starting service")

	//
	// configure the application
	//

	udpAPICfg := &udp.Config{}
	k8Cfg := &k8.Config{}
	sbcsCfg := &sbcs.Config{}

	cfg := config.New()
	cfg.Set("UDP_API", udpAPICfg)
	cfg.Set("K8", k8Cfg)
	cfg.Set("SBCS", sbcsCfg)
	cfg.Run()

	//
	//
	//

	// create an ExitManager instance
	exitMgr := exitmanager.New()
	go exitMgr.Run()

	//
	// create clients
	//

	// create K8 clients
	k8Clients := k8.New(k8Cfg, logger)

	// create SBCs client
	sbcsClient := sbcs.New(&sbcs.Input{
		Ctx:           exitMgr.Context(),
		Config:        &sbcs.Config{},
		Logger:        logger,
		SbcV1Client:   k8Clients.SbcV1Client,
		SbcV1Informer: k8Clients.SbcInformer,
	})

	// run the SBC controller
	go sbcsClient.RunController()

	//
	//
	//

	// create UDP API instance
	udpAPI := udp.New(&udp.Input{
		K8Clients:  k8Clients,
		SbcsClient: sbcsClient,
	})

	// run the APIs
	go func() {
		err := udpAPI.Run()
		logger.Errorf("server failure: %w", err)
	}()

	// wait until graceful shutdown has resolved before exiting the application
	exitMgr.Wait()

	logger.Info("closing service")
}
