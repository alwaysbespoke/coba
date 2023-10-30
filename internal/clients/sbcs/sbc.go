package sbcs

import (
	"net"

	v1 "github.com/alwaysbespoke/coba/pkg/crds/sbc/v1"
)

type SBC struct {
	Obj  *v1.SBC
	Conn *net.UDPConn
}
