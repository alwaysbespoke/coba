package handler

import (
	"github.com/alwaysbespoke/coba/apps/sip-server/internal/clients/k8"
	"github.com/alwaysbespoke/coba/apps/sip-server/internal/clients/sbcs"
	"go.uber.org/zap"
)

type Config struct {
	Namespace string
}

type Handler struct {
	config     *Config
	logger     *zap.SugaredLogger
	K8Clients  *k8.Clients
	SbcsClient sbcs.Client
}

func New(config *Config, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		config: config,
		logger: logger,
	}
}
