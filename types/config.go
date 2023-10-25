package types

import (
	"fmt"
	"net/url"
	"strings"
)

type IgniteConfig struct {
	Version  int `yaml:"version" json:"version"`
	Accounts []struct {
		Name  string   `yaml:"name" json:"name,omitempty"`
		Coins []string `yaml:"coins" json:"coins,omitempty"`
	} `yaml:"accounts" json:"accounts"`
	Client struct {
		Openapi struct {
			Path string `yaml:"path" json:"path,omitempty"`
		} `yaml:"openapi" json:"openapi"`
	} `yaml:"client" json:"client"`
	Validators []struct {
		Name   string `yaml:"name" json:"name"`
		Bonded string `yaml:"bonded" json:"bonded"`
		Home   string `yaml:"home" json:"home"`
		IP     string `yaml:"ip" json:"ip"`
		App    struct {
			MinimumGasPrices string `yaml:"minimum-gas-prices" json:"minimum-gas-prices"`
			API              struct {
				Enable            bool   `yaml:"enable" json:"enable,omitempty"`
				EnabledUnsafeCors bool   `yaml:"enabled-unsafe-cors" json:"enabled-unsafe-cors,omitempty"`
				Address           string `yaml:"address" json:"address,omitempty"`
			} `yaml:"api"`
			Grpc struct {
				Enable  bool   `yaml:"enable" json:"enable,omitempty"`
				Address string `yaml:"address" json:"address,omitempty"`
			} `yaml:"grpc"`
			GrpcWeb struct {
				Address          string `yaml:"address" json:"address,omitempty"`
				Enable           bool   `yaml:"enable" json:"enable,omitempty"`
				EnableUnsafeCors bool   `yaml:"enable-unsafe-cors" json:"enable-unsafe-cors,omitempty"`
			} `yaml:"grpc-web" json:"grpc-web"`
		} `yaml:"app" json:"app"`
		Config struct {
			Consensus struct {
				TimeoutCommit string `yaml:"timeout_commit" json:"timeout_commit,omitempty"`
			} `yaml:"consensus" json:"consensus"`
			ProxyApp string `yaml:"proxy_app" json:"proxy_app"`
			Moniker  string `yaml:"moniker" json:"moniker"`
			RPC      struct {
				MaxBodyBytes string `yaml:"max_body_bytes" json:"max_body_bytes"`
				Laddr        string `yaml:"laddr" json:"laddr,omitempty"`
			} `yaml:"rpc" json:"rpc"`
			P2P struct {
				Laddr            string `yaml:"laddr" json:"laddr,omitempty"`
				PersistentPeers  string `yaml:"persistent_peers" json:"persistent_peers,omitempty"`
				AllowDuplicateIP bool   `yaml:"allow_duplicate_ip" json:"allow_duplicate_ip,omitempty"`
			} `yaml:"p2p" json:"p2p"`
			Instrumentation struct {
				Prometheus           bool   `yaml:"prometheus" json:"prometheus"`
				PrometheusListenAddr string `yaml:"prometheus_listen_addr" json:"prometheus_listen_addr"`
			} `yaml:"instrumentation" json:"instrumentation"`
		} `yaml:"config" json:"config"`
	} `yaml:"validators" json:"validators"`
	Genesis struct {
		ConsensusParams struct {
			Block struct {
				MaxBytes string `yaml:"max_bytes" json:"max_bytes"`
				MaxGas   string `yaml:"max_gas" json:"max_gas"`
			} `yaml:"block" json:"block"`
			Evidence struct {
				MaxAgeNumBlocks string `yaml:"max_age_num_blocks" json:"max_age_num_blocks"`
				MaxAgeDuration  string `yaml:"max_age_duration" json:"max_age_duration"`
				MaxBytes        string `yaml:"max_bytes" json:"max_bytes"`
			} `yaml:"evidence" json:"evidence"`
			Validator struct {
				PubKeyTypes []string `yaml:"pub_key_types" json:"pub_key_types"`
			} `yaml:"validator" json:"validator"`
			Version struct {
				App string `yaml:"app" json:"app"`
			} `yaml:"version" json:"version"`
		} `yaml:"consensus_params" json:"consensus_params"`
		AppState struct {
			Bank struct {
				DenomMetadata []struct {
					Description string `yaml:"description" json:"description,omitempty"`
					Base        string `yaml:"base" json:"base,omitempty"`
					Display     string `yaml:"display" json:"display,omitempty"`
					Name        string `yaml:"name" json:"name,omitempty"`
					Symbol      string `yaml:"symbol" json:"symbol,omitempty"`
					URI         string `yaml:"uri" json:"uri,omitempty"`
					DenomUnits  []struct {
						Aliases  []interface{} `yaml:"aliases" json:"aliases,omitempty"`
						Denom    string        `yaml:"denom" json:"denom,omitempty"`
						Exponent int           `yaml:"exponent" json:"exponent,omitempty"`
					} `yaml:"denom_units" json:"denom_units,omitempty"`
				} `yaml:"denom_metadata" json:"denom_metadata"`
			} `yaml:"bank" json:"bank"`
			Staking struct {
				Params struct {
					BondDenom     string `yaml:"bond_denom" json:"bond_denom,omitempty"`
					MaxValidators string `yaml:"max_validators" json:"max_validators,omitempty"`
				} `yaml:"params" json:"params"`
			} `yaml:"staking" json:"staking"`
			Mint struct {
				Minter struct {
					AnnualProvisions string ` yaml:"annual_provisions" json:"annual_provisions,omitempty"`
					Inflation        string `yaml:"inflation" json:"inflation,omitempty"`
				} `yaml:"minter" json:"minter"`
				Params struct {
					MintDenom           string `yaml:"mint_denom" json:"mint_denom,omitempty"`
					BlocksPerYear       string `yaml:"blocks_per_year" json:"blocks_per_year"`
					GoalBonded          string `yaml:"goal_bonded" json:"goal_bonded"`
					InflationMax        string `yaml:"inflation_max" json:"inflation_max"`
					InflationMin        string `yaml:"inflation_min" json:"inflation_min"`
					InflationRateChange string `yaml:"inflation_rate_change" json:"inflation_rate_change"`
					Reduction           struct {
						Enable          bool     `yaml:"enable" json:"enable"`
						TotalProvisions string   `yaml:"total_provisions" json:"total_provisions"`
						Heights         []uint64 `yaml:"heights" json:"heights"`
					} `yaml:"reduction" json:"reduction"`
				} `yaml:"params" json:"params"`
			} `yaml:"mint" json:"mint"`
			Gov struct {
				Params struct {
					MinDeposit []struct {
						Amount string `yaml:"amount" json:"amount" `
						Denom  string `yaml:"denom" json:"denom" `
					} `yaml:"min_deposit" json:"min_deposit"`
				} `yaml:"params" json:"params"`
			} `yaml:"gov" json:"gov"`
			Distribution struct {
				Params struct {
					BaseProposerReward  string ` yaml:"base_proposer_reward" json:"base_proposer_reward,omitempty"`
					BonusProposerReward string `yaml:"bonus_proposer_reward" json:"bonus_proposer_reward,omitempty"`
					CommunityTax        string `yaml:"community_tax" json:"community_tax,omitempty"`
					WithdrawAddrEnabled bool   `yaml:"withdraw_addr_enabled" json:"withdraw_addr_enabled,omitempty"`
				} `yaml:"params" json:"params"`
			} `yaml:"distribution" json:"distribution"`
			Crisis struct {
				ConstantFee struct {
					Amount string `yaml:"amount" json:"amount,omitempty"`
					Denom  string `yaml:"denom" json:"denom,omitempty"`
				} `yaml:"constant_fee" json:"constant_fee"`
			} `yaml:"crisis" json:"crisis"`
		} `yaml:"app_state" json:"app_state"`
	} `yaml:"genesis" json:"genesis"`
}

func (m IgniteConfig) GetAccountBalances(name string) string {
	for _, a := range m.Accounts {
		if a.Name == name {
			return strings.Join(a.Coins, ",")
		}
	}
	return ""
}

func (m IgniteConfig) GetValidatorHost(strValidatorName string) string {
	for _, v := range m.Validators {
		if v.Name == strValidatorName {
			return fmt.Sprintf("%s:%s", v.IP, parseP2PPort(v.Config.P2P.Laddr))
		}
	}
	return "<N/A>"
}

func parseP2PPort(strAddr string) string {
	u, err := url.Parse(strAddr)
	if err != nil {
		return COSMOS_P2P_PORT
	}
	return u.Port()
}

type CosmosConfig struct {
	ProxyApp               string `toml:"proxy_app"`
	Moniker                string `toml:"moniker"`
	BlockSync              bool   `toml:"block_sync"`
	DbBackend              string `toml:"db_backend"`
	DbDir                  string `toml:"db_dir"`
	LogLevel               string `toml:"log_level"`
	LogFormat              string `toml:"log_format"`
	GenesisFile            string `toml:"genesis_file"`
	PrivValidatorKeyFile   string `toml:"priv_validator_key_file"`
	PrivValidatorStateFile string `toml:"priv_validator_state_file"`
	PrivValidatorLaddr     string `toml:"priv_validator_laddr"`
	NodeKeyFile            string `toml:"node_key_file"`
	Abci                   string `toml:"abci"`
	FilterPeers            bool   `toml:"filter_peers"`
	RPC                    struct {
		Laddr                                string   `toml:"laddr"`
		CorsAllowedOrigins                   []string `toml:"cors_allowed_origins"`
		CorsAllowedMethods                   []string `toml:"cors_allowed_methods"`
		CorsAllowedHeaders                   []string `toml:"cors_allowed_headers"`
		GrpcLaddr                            string   `toml:"grpc_laddr"`
		GrpcMaxOpenConnections               int      `toml:"grpc_max_open_connections"`
		Unsafe                               bool     `toml:"unsafe"`
		MaxOpenConnections                   int      `toml:"max_open_connections"`
		MaxSubscriptionClients               int      `toml:"max_subscription_clients"`
		MaxSubscriptionsPerClient            int      `toml:"max_subscriptions_per_client"`
		ExperimentalSubscriptionBufferSize   int      `toml:"experimental_subscription_buffer_size"`
		ExperimentalWebsocketWriteBufferSize int      `toml:"experimental_websocket_write_buffer_size"`
		ExperimentalCloseOnSlowClient        bool     `toml:"experimental_close_on_slow_client"`
		TimeoutBroadcastTxCommit             string   `toml:"timeout_broadcast_tx_commit"`
		MaxBodyBytes                         int      `toml:"max_body_bytes"`
		MaxHeaderBytes                       int      `toml:"max_header_bytes"`
		TLSCertFile                          string   `toml:"tls_cert_file"`
		TLSKeyFile                           string   `toml:"tls_key_file"`
		PprofLaddr                           string   `toml:"pprof_laddr"`
	} `toml:"rpc"`
	P2P struct {
		Laddr                        string `toml:"laddr"`
		ExternalAddress              string `toml:"external_address"`
		Seeds                        string `toml:"seeds"`
		PersistentPeers              string `toml:"persistent_peers"`
		Upnp                         bool   `toml:"upnp"`
		AddrBookFile                 string `toml:"addr_book_file"`
		AddrBookStrict               bool   `toml:"addr_book_strict"`
		MaxNumInboundPeers           int    `toml:"max_num_inbound_peers"`
		MaxNumOutboundPeers          int    `toml:"max_num_outbound_peers"`
		UnconditionalPeerIds         string `toml:"unconditional_peer_ids"`
		PersistentPeersMaxDialPeriod string `toml:"persistent_peers_max_dial_period"`
		FlushThrottleTimeout         string `toml:"flush_throttle_timeout"`
		MaxPacketMsgPayloadSize      int    `toml:"max_packet_msg_payload_size"`
		SendRate                     int    `toml:"send_rate"`
		RecvRate                     int    `toml:"recv_rate"`
		Pex                          bool   `toml:"pex"`
		SeedMode                     bool   `toml:"seed_mode"`
		PrivatePeerIds               string `toml:"private_peer_ids"`
		AllowDuplicateIP             bool   `toml:"allow_duplicate_ip"`
		HandshakeTimeout             string `toml:"handshake_timeout"`
		DialTimeout                  string `toml:"dial_timeout"`
	} `toml:"p2p"`
	Mempool struct {
		Version               string `toml:"version"`
		Recheck               bool   `toml:"recheck"`
		Broadcast             bool   `toml:"broadcast"`
		WalDir                string `toml:"wal_dir"`
		Size                  int    `toml:"size"`
		MaxTxsBytes           int    `toml:"max_txs_bytes"`
		CacheSize             int    `toml:"cache_size"`
		KeepInvalidTxsInCache bool   `toml:"keep-invalid-txs-in-cache"`
		MaxTxBytes            int    `toml:"max_tx_bytes"`
		MaxBatchBytes         int    `toml:"max_batch_bytes"`
		TTLDuration           string `toml:"ttl-duration"`
		TTLNumBlocks          int    `toml:"ttl-num-blocks"`
	} `toml:"mempool"`
	Statesync struct {
		Enable              bool   `toml:"enable"`
		RPCServers          string `toml:"rpc_servers"`
		TrustHeight         int    `toml:"trust_height"`
		TrustHash           string `toml:"trust_hash"`
		TrustPeriod         string `toml:"trust_period"`
		DiscoveryTime       string `toml:"discovery_time"`
		TempDir             string `toml:"temp_dir"`
		ChunkRequestTimeout string `toml:"chunk_request_timeout"`
		ChunkFetchers       string `toml:"chunk_fetchers"`
	} `toml:"statesync"`
	Blocksync struct {
		Version string `toml:"version"`
	} `toml:"blocksync"`
	Consensus struct {
		WalFile                     string `toml:"wal_file"`
		TimeoutPropose              string `toml:"timeout_propose"`
		TimeoutProposeDelta         string `toml:"timeout_propose_delta"`
		TimeoutPrevote              string `toml:"timeout_prevote"`
		TimeoutPrevoteDelta         string `toml:"timeout_prevote_delta"`
		TimeoutPrecommit            string `toml:"timeout_precommit"`
		TimeoutPrecommitDelta       string `toml:"timeout_precommit_delta"`
		TimeoutCommit               string `toml:"timeout_commit"`
		DoubleSignCheckHeight       int    `toml:"double_sign_check_height"`
		SkipTimeoutCommit           bool   `toml:"skip_timeout_commit"`
		CreateEmptyBlocks           bool   `toml:"create_empty_blocks"`
		CreateEmptyBlocksInterval   string `toml:"create_empty_blocks_interval"`
		PeerGossipSleepDuration     string `toml:"peer_gossip_sleep_duration"`
		PeerQueryMaj23SleepDuration string `toml:"peer_query_maj23_sleep_duration"`
	} `toml:"consensus"`
	Storage struct {
		DiscardAbciResponses bool `toml:"discard_abci_responses"`
	} `toml:"storage"`
	TxIndex struct {
		Indexer  string `toml:"indexer"`
		PsqlConn string `toml:"psql-conn"`
	} `toml:"tx_index"`
	Instrumentation struct {
		Prometheus           bool   `toml:"prometheus"`
		PrometheusListenAddr string `toml:"prometheus_listen_addr"`
		MaxOpenConnections   int    `toml:"max_open_connections"`
		Namespace            string `toml:"namespace"`
	} `toml:"instrumentation"`
}
