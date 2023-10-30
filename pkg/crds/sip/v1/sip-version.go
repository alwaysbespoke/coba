package v1

import (
	"sync"
)

type SIPVersion string

const (
	SIP20 SIPVersion = "SIP/2.0"
)

var versionFromStringMap = map[string]SIPVersion{
	"SIP/2.0": SIP20,
}

var versionFromStringLock sync.RWMutex = sync.RWMutex{}

func VersionFromString(s string) (SIPVersion, bool) {
	var version SIPVersion
	var isVersion bool

	versionFromStringLock.RLock()
	defer versionFromStringLock.RUnlock()

	version, isVersion = versionFromStringMap[s]

	return version, isVersion
}

var versionToStringMap = map[SIPVersion]string{
	SIP20: "SIP/2.0",
}

var versionToStringLock sync.RWMutex = sync.RWMutex{}

func (v SIPVersion) String() string {
	versionToStringLock.RLock()
	defer versionToStringLock.RUnlock()

	return versionToStringMap[v]
}
