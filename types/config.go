package types

import (
	"fmt"
	"net/url"
	"strings"
	"time"
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
		ChainID         string `yaml:"chain_id" json:"chain_id""`
		InitialHeight   string `yaml:"initial_height" json:"initial_height"`
		GenesisTime     string `yaml:"genesis_time" json:"genesis_time"`
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
			Claims struct {
				ClaimsRecords []interface{} `json:"claims_records" yaml:"claims_records"`
				Params        struct {
					AirdropStartTime   time.Time `json:"airdrop_start_time" yaml:"airdrop_start_time"`
					AuthorizedChannels []string  `json:"authorized_channels" yaml:"authorized_channels"`
					ClaimsDenom        string    `json:"claims_denom" yaml:"claims_denom"`
					DurationOfDecay    string    `json:"duration_of_decay" yaml:"duration_of_decay"`
					DurationUntilDecay string    `json:"duration_until_decay" yaml:"duration_until_decay"`
					EnableClaims       bool      `json:"enable_claims" yaml:"enable_claims"`
					EvmChannels        []string  `json:"evm_channels" yaml:"evm_channels"`
				} `json:"params" yaml:"params"`
			} `json:"claims" yaml:"claims"`
			Evm struct {
				Accounts []interface{} `json:"accounts"`
				Params   struct {
					ActivePrecompiles   []string `json:"active_precompiles" yaml:"active_precompiles"`
					AllowUnprotectedTxs bool     `json:"allow_unprotected_txs" yaml:"allow_unprotected"`
					ChainConfig         struct {
						ArrowGlacierBlock   string `json:"arrow_glacier_block" yaml:"arrow_glacier_block"`
						BerlinBlock         string `json:"berlin_block" yaml:"berlin_block"`
						ByzantiumBlock      string `json:"byzantium_block" yaml:"byzantium_block"`
						CancunBlock         string `json:"cancun_block" yaml:"cancun_block"`
						ConstantinopleBlock string `json:"constantinople_block" yaml:"constantinople_block"`
						DaoForkBlock        string `json:"dao_fork_block" yaml:"dao_fork_block"`
						DaoForkSupport      bool   `json:"dao_fork_support" yaml:"dao_fork_support"`
						Eip150Block         string `json:"eip150_block" yaml:"eip_150_block"`
						Eip150Hash          string `json:"eip150_hash" yaml:"eip_150_hash"`
						Eip155Block         string `json:"eip155_block" yaml:"eip_155_block"`
						Eip158Block         string `json:"eip158_block" yaml:"eip_158_block"`
						GrayGlacierBlock    string `json:"gray_glacier_block" yaml:"gray_glacier_block"`
						HomesteadBlock      string `json:"homestead_block" yaml:"homestead_block"`
						IstanbulBlock       string `json:"istanbul_block" yaml:"istanbul_block"`
						LondonBlock         string `json:"london_block" yaml:"london_block"`
						MergeNetsplitBlock  string `json:"merge_netsplit_block" yaml:"merge_netsplit_block"`
						MuirGlacierBlock    string `json:"muir_glacier_block" yaml:"muir_glacier_block"`
						PetersburgBlock     string `json:"petersburg_block" yaml:"petersburg_block"`
						ShanghaiBlock       string `json:"shanghai_block" yaml:"shanghai_block"`
					} `json:"chain_config"`
					EnableCall   bool     `json:"enable_call" yaml:"enable_call"`
					EnableCreate bool     `json:"enable_create" yaml:"enable_create"`
					EvmDenom     string   `json:"evm_denom" yaml:"evm_denom"`
					ExtraEips    []string `json:"extra_eips" yaml:"extra_eips"`
				} `json:"params"`
			} `json:"evm"`

			Feemarket struct {
				BlockGas string `json:"block_gas" yaml:"block_gas"`
				Params   struct {
					BaseFee                  string `json:"base_fee" yaml:"base_fee"`
					BaseFeeChangeDenominator int    `json:"base_fee_change_denominator" yaml:"base_fee_change_denominator"`
					ElasticityMultiplier     int    `json:"elasticity_multiplier" yaml:"elasticity_multiplier"`
					EnableHeight             string `json:"enable_height" yaml:"enable_height"`
					MinGasMultiplier         string `json:"min_gas_multiplier" yaml:"min_gas_multiplier"`
					MinGasPrice              string `json:"min_gas_price" yaml:"min_gas_price"`
					NoBaseFee                bool   `json:"no_base_fee" yaml:"no_base_fee"`
				} `json:"params" yaml:"params"`
			} `json:"feemarket" yaml:"feemarket"`

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
