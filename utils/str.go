package utils

import (
	"github.com/civet148/cosmos-cli/types"
	"github.com/civet148/log"
	"net/url"
)

// ParseP2PPort parse comsos p2p port from listen address. eg. "tcp://0.0.0.0:26656"
func ParseP2PPort(strAddr string) string {
	u, err := url.Parse(strAddr)
	if err != nil {
		log.Warnf("p2p listen address %s is invalid", strAddr)
		return types.COSMOS_P2P_PORT
	}
	return u.Port()
}
