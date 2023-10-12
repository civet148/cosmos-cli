package utils

import (
	"fmt"
	"testing"
)

func TestParseP2PPort(t *testing.T) {
	strPort := ParseP2PPort("tcp://0.0.0.0:26656")
	fmt.Printf("p2p port %s\n", strPort)
}
