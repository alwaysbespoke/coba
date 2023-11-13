package handlers

import "go.uber.org/zap"

type Handlers struct {
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
