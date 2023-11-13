package handler

import "go.uber.org/zap"

type Config struct {
}

type Handler struct {
	config *Config
	logger *zap.SugaredLogger
}

func New(config *Config, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		config: config,
		logger: logger,
	}
}
