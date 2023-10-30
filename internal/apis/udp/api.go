package udp

import (
	"fmt"
	"net"

	"go.uber.org/zap"

	"github.com/alwaysbespoke/coba/internal/clients/k8"
	"github.com/alwaysbespoke/coba/internal/clients/sbcs"
)

type Config struct {
	Address       string `default:":5000"`
	BufferSize    int    `default:"1024"`
	Namespace     string
	Configuration string
}

type API struct {
	*Input
}

type Input struct {
	Config     *Config
	Logger     *zap.SugaredLogger
	K8Clients  *k8.Clients
	SbcsClient sbcs.Client
}

// New returns a new API instance
func New(input *Input) *API {
	return &API{
		Input: input,
	}
}

// Run starts the listen and serve loop for the API
func (a *API) Run() error {
	conn, err := net.ListenPacket("udp", a.Config.Address)
	if err != nil {
		return fmt.Errorf("server failure: %w", err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, a.Config.BufferSize)

		numBytesRead, addr, err := conn.ReadFrom(buf)
		if err != nil {
			a.Logger.Errorf("server read error: %w", err)
			continue
		}

		go a.handlePacket(conn, addr, buf[:numBytesRead])
	}
}
