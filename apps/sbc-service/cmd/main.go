package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/alwaysbespoke/coba/pkg/config"

	api "github.com/alwaysbespoke/coba/apps/sbc-service/internal/apis/rest"
	exitmanager "github.com/alwaysbespoke/coba/pkg/exit-manager"
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
	apiCfg := &api.Config{}

	// process the configs
	cfg := config.New()
	cfg.Set("API", apiCfg)
	cfg.Run()

	/*
	*
	* Create the application clients
	*
	 */

	// create an ExitManager instance
	exitMgr, _ := exitmanager.New()
	go exitMgr.Run()

	/*
	*
	* Create and run the APIs and controllers
	*
	 */
	// create the REST API
	api := api.New(apiCfg, logger)

	// run the REST API
	go func() {
		err := api.Run()
		logger.Errorf("server failure: %w", err)
	}()

	// wait until graceful shutdown has resolved before exiting the application
	exitMgr.Wait()

	logger.Info("closing service")
}
