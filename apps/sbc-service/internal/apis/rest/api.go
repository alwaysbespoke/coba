package api

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/alwaysbespoke/coba/apps/sbc-service/internal/apis/rest/routes"

	v1handlers "github.com/alwaysbespoke/coba/apps/sbc-service/internal/apis/rest/handlers/v1"
)

type Config struct {
	Address string `default:":8000"`
}

type API struct {
	logger *zap.SugaredLogger
	server *http.Server
}

func New(config *Config, logger *zap.SugaredLogger) *API {
	v1Handlers := v1handlers.New(logger)

	router := routes.New(v1Handlers)

	server := &http.Server{
		Addr:    config.Address,
		Handler: router,
	}

	return &API{
		logger: logger,
		server: server,
	}
}

func (a *API) Run() error {
	err := a.server.ListenAndServe()
	err = fmt.Errorf("server failure: %w", err)
	a.logger.Error(err)
	return err
}
