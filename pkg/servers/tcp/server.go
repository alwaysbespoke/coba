package server

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

type Config struct {
	Address    string `default:":6000"`
	BufferSize int    `default:"1024" split_words:"true"`
}

type Server struct {
	config  *Config
	logger  *zap.SugaredLogger
	handler Handler
}

type Handler interface {
	Handle(conn net.Conn, buf []byte)
}

// New returns a new Server instance
func New(config *Config, logger *zap.SugaredLogger, handler Handler) *Server {
	return &Server{
		config:  config,
		logger:  logger,
		handler: handler,
	}
}

// Run starts the listener and handles incoming connections
func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Errorf("failed to accept connection: %w", err)
			continue
		}

		buf := make([]byte, s.config.BufferSize)

		numBytesRead, err := conn.Read(buf)
		if err != nil {
			s.logger.Errorf("failed to read from connection: %w", err)
			continue
		}

		go s.handler.Handle(conn, buf[:numBytesRead])
	}
}
