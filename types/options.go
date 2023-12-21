package types

type Option struct {
	Debug          bool   // debug mode on
	ConfigPath     string // config file path
	NodeCmd        string // chain node command
	DefaultDenom   string // default denom
	ChainID        string // chain id
	KeyPhrase      string // pass phrase to protect keys
	KeyringBackend string // keyring backend
}

type NodePeer struct {
	Name            string
	Peer            string
	PersistentPeers []string
}
