package types

type ClientConfig struct {
	BroadcastMode  string `toml:"broadcast-mode"`
	ChainID        string `toml:"chain-id"`
	KeyringBackend string `toml:"keyring-backend"`
	Node           string `toml:"node"`
	Output         string `toml:"output"`
}
