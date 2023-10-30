package v1

import "sync"

type Method string

const (
	Ack      Method = "ACK"
	Bye      Method = "BYE"
	Cancel   Method = "CANCEL"
	Invite   Method = "INVITE"
	Options  Method = "OPTIONS"
	Register Method = "REGISTER"
)

var methodToStringMap = map[Method]string{
	Ack:      "ACK",
	Bye:      "BYE",
	Cancel:   "CANCEL",
	Invite:   "INVITE",
	Options:  "OPTIONS",
	Register: "REGISTER",
}

var methodToStringLock sync.RWMutex = sync.RWMutex{}

func (m Method) String() string {
	methodToStringLock.RLock()
	defer methodToStringLock.RUnlock()

	return methodToStringMap[m]
}

var methodFromStringMap = map[string]Method{
	"ACK":      Ack,
	"BYE":      Bye,
	"CANCEL":   Cancel,
	"INVITE":   Invite,
	"OPTIONS":  Options,
	"REGISTER": Register,
}

var methodFromStringLock sync.RWMutex = sync.RWMutex{}

func MethodFromString(s string) (Method, bool) {
	var method Method
	var isMethod bool

	methodFromStringLock.RLock()
	defer methodFromStringLock.RUnlock()

	method, isMethod = methodFromStringMap[s]

	return method, isMethod
}
