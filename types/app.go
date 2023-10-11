package types

type AppConfig struct {
	AppDbBackend        string        `toml:"app-db-backend"`
	HaltHeight          int           `toml:"halt-height"`
	HaltTime            int           `toml:"halt-time"`
	IavlCacheSize       int           `toml:"iavl-cache-size"`
	IavlDisableFastnode bool          `toml:"iavl-disable-fastnode"`
	IavlLazyLoading     bool          `toml:"iavl-lazy-loading"`
	IndexEvents         []interface{} `toml:"index-events"`
	InterBlockCache     bool          `toml:"inter-block-cache"`
	MinRetainBlocks     int           `toml:"min-retain-blocks"`
	MinimumGasPrices    string        `toml:"minimum-gas-prices"`
	Pruning             string        `toml:"pruning"`
	PruningInterval     string        `toml:"pruning-interval"`
	PruningKeepRecent   string        `toml:"pruning-keep-recent"`
	API                 struct {
		Address            string `toml:"address"`
		Enable             bool   `toml:"enable"`
		EnabledUnsafeCors  bool   `toml:"enabled-unsafe-cors"`
		MaxOpenConnections int    `toml:"max-open-connections"`
		RPCMaxBodyBytes    int    `toml:"rpc-max-body-bytes"`
		RPCReadTimeout     int    `toml:"rpc-read-timeout"`
		RPCWriteTimeout    int    `toml:"rpc-write-timeout"`
		Swagger            bool   `toml:"swagger"`
	} `toml:"api"`
	Grpc struct {
		Address        string `toml:"address"`
		Enable         bool   `toml:"enable"`
		MaxRecvMsgSize string `toml:"max-recv-msg-size"`
		MaxSendMsgSize string `toml:"max-send-msg-size"`
	} `toml:"grpc"`
	GrpcWeb struct {
		Address          string `toml:"address"`
		Enable           bool   `toml:"enable"`
		EnableUnsafeCors bool   `toml:"enable-unsafe-cors"`
	} `toml:"grpc-web"`
	Mempool struct {
		MaxTxs string `toml:"max-txs"`
	} `toml:"mempool"`
	Rosetta struct {
		Address             string `toml:"address"`
		Blockchain          string `toml:"blockchain"`
		DenomToSuggest      string `toml:"denom-to-suggest"`
		Enable              bool   `toml:"enable"`
		EnableFeeSuggestion bool   `toml:"enable-fee-suggestion"`
		GasToSuggest        int    `toml:"gas-to-suggest"`
		Network             string `toml:"network"`
		Offline             bool   `toml:"offline"`
		Retries             int    `toml:"retries"`
	} `toml:"rosetta"`
	RPC struct {
		CorsAllowedOrigins []string `toml:"cors_allowed_origins"`
	} `toml:"rpc"`
	StateSync struct {
		SnapshotInterval   int `toml:"snapshot-interval"`
		SnapshotKeepRecent int `toml:"snapshot-keep-recent"`
	} `toml:"state-sync"`
	Store struct {
		Streamers []interface{} `toml:"streamers"`
	} `toml:"store"`
	Streamers struct {
		File struct {
			Fsync           string   `toml:"fsync"`
			Keys            []string `toml:"keys"`
			OutputMetadata  string   `toml:"output-metadata"`
			Prefix          string   `toml:"prefix"`
			StopNodeOnError string   `toml:"stop-node-on-error"`
			WriteDir        string   `toml:"write_dir"`
		} `toml:"file"`
	} `toml:"streamers"`
	Telemetry struct {
		EnableHostname          bool     `toml:"enable-hostname"`
		EnableHostnameLabel     bool     `toml:"enable-hostname-label"`
		EnableServiceLabel      bool     `toml:"enable-service-label"`
		Enabled                 bool     `toml:"enabled"`
		GlobalLabels            []string `toml:"global-labels"`
		PrometheusRetentionTime int      `toml:"prometheus-retention-time"`
		ServiceName             string   `toml:"service-name"`
	} `toml:"telemetry"`
}
