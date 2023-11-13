package server

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

type Config struct {
	Address    string `default:":5000"`
	BufferSize int    `default:"1024" split_words:"true"`
}

type Server struct {
	config  *Config
	logger  *zap.SugaredLogger
	handler Handler
}

type Handler interface {
	Handle(conn net.PacketConn, addr net.Addr, buf []byte)
}

// New returns a new Server instance
func New(config *Config, logger *zap.SugaredLogger, handler Handler) *Server {
	return &Server{
		config:  config,
		logger:  logger,
		handler: handler,
	}
}

// Run starts the listener and handles incoming packets
func (s *Server) Run() error {
	conn, err := net.ListenPacket("udp", s.config.Address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, s.config.BufferSize)

		numBytesRead, addr, err := conn.ReadFrom(buf)
		if err != nil {
			s.logger.Errorf("failed to read from connection: %w", err)
			continue
		}

		go s.handler.Handle(conn, addr, buf[:numBytesRead])
	}
}
