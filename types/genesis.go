package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	dist "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"time"
)

type GenesisConfig struct {
	GenesisTime     time.Time `json:"genesis_time"`
	ChainId         string    `json:"chain_id"`
	InitialHeight   string    `json:"initial_height"`
	ConsensusParams struct {
		Block struct {
			MaxBytes string `json:"max_bytes"`
			MaxGas   string `json:"max_gas"`
		} `json:"block"`
		Evidence struct {
			MaxAgeNumBlocks string `json:"max_age_num_blocks"`
			MaxAgeDuration  string `json:"max_age_duration"`
			MaxBytes        string `json:"max_bytes"`
		} `json:"evidence"`
		Validator struct {
			PubKeyTypes []string `json:"pub_key_types"`
		} `json:"validator"`
		Version struct {
			App string `json:"app"`
		} `json:"version"`
	} `json:"consensus_params"`
	AppHash  string `json:"app_hash"`
	AppState struct {
		Solomachine string `json:"06-solomachine"`
		Tendermint  string `json:"07-tendermint"`
		Auth        struct {
			Params struct {
				MaxMemoCharacters      string `json:"max_memo_characters"`
				TxSigLimit             string `json:"tx_sig_limit"`
				TxSizeCostPerByte      string `json:"tx_size_cost_per_byte"`
				SigVerifyCostEd25519   string `json:"sig_verify_cost_ed25519"`
				SigVerifyCostSecp256K1 string `json:"sig_verify_cost_secp256k1"`
			} `json:"params"`
			Accounts []struct {
				Type          string `json:"@type"`
				Address       string `json:"address"`
				PubKey        string `json:"pub_key"`
				AccountNumber string `json:"account_number"`
				Sequence      string `json:"sequence"`
			} `json:"accounts"`
		} `json:"auth"`
		Authz struct {
			Authorization []interface{} `json:"authorization"`
		} `json:"authz"`
		Bank struct {
			Params struct {
				SendEnabled        []bank.SendEnabled `json:"send_enabled"`
				DefaultSendEnabled bool               `json:"default_send_enabled"`
			} `json:"params"`
			Balances []struct {
				Address string `json:"address"`
				Coins   []struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"coins"`
			} `json:"balances"`
			Supply        []sdk.Coin         `json:"supply"`
			DenomMetadata []bank.Metadata    `json:"denom_metadata"`
			SendEnabled   []bank.SendEnabled `json:"send_enabled"`
		} `json:"bank"`
		Capability struct {
			Index  string        `json:"index"`
			Owners []interface{} `json:"owners"`
		} `json:"capability"`
		Consensus interface{} `json:"consensus"`
		Crisis    struct {
			ConstantFee struct {
				Amount string `json:"amount"`
				Denom  string `json:"denom"`
			} `json:"constant_fee"`
		} `json:"crisis"`
		Distribution struct {
			DelegatorStartingInfos []dist.DelegatorStartingInfoRecord `json:"delegator_starting_infos"`
			DelegatorWithdrawInfos []dist.DelegatorWithdrawInfo       `json:"delegator_withdraw_infos"`
			FeePool                struct {
				CommunityPool []dist.FeePool `json:"community_pool"`
			} `json:"fee_pool"`
			OutstandingRewards []sdk.DecCoins `json:"outstanding_rewards"`
			Params             struct {
				BaseProposerReward  string `json:"base_proposer_reward"`
				BonusProposerReward string `json:"bonus_proposer_reward"`
				CommunityTax        string `json:"community_tax"`
				WithdrawAddrEnabled bool   `json:"withdraw_addr_enabled"`
			} `json:"params"`
			PreviousProposer                string                                  `json:"previous_proposer"`
			ValidatorAccumulatedCommissions []dist.ValidatorAccumulatedCommission   `json:"validator_accumulated_commissions"`
			ValidatorCurrentRewards         []dist.ValidatorCurrentRewardsRecord    `json:"validator_current_rewards"`
			ValidatorHistoricalRewards      []dist.ValidatorHistoricalRewardsRecord `json:"validator_historical_rewards"`
			ValidatorSlashEvents            []dist.ValidatorSlashEvent              `json:"validator_slash_events"`
		} `json:"distribution"`
		Evidence struct {
			Evidence []interface{} `json:"evidence"`
		} `json:"evidence"`
		Feegrant struct {
			Allowances []interface{} `json:"allowances"`
		} `json:"feegrant"`
		Genutil struct {
			GenTxs []struct {
				Body struct {
					Messages []struct {
						Type        string `json:"@type"`
						Description struct {
							Moniker         string `json:"moniker"`
							Identity        string `json:"identity"`
							Website         string `json:"website"`
							SecurityContact string `json:"security_contact"`
							Details         string `json:"details"`
						} `json:"description"`
						Commission struct {
							Rate          string `json:"rate"`
							MaxRate       string `json:"max_rate"`
							MaxChangeRate string `json:"max_change_rate"`
						} `json:"commission"`
						MinSelfDelegation string `json:"min_self_delegation"`
						DelegatorAddress  string `json:"delegator_address"`
						ValidatorAddress  string `json:"validator_address"`
						Pubkey            struct {
							Type string `json:"@type"`
							Key  string `json:"key"`
						} `json:"pubkey"`
						Value struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"value"`
					} `json:"messages"`
					Memo                        string        `json:"memo"`
					TimeoutHeight               string        `json:"timeout_height"`
					ExtensionOptions            []interface{} `json:"extension_options"`
					NonCriticalExtensionOptions []interface{} `json:"non_critical_extension_options"`
				} `json:"body"`
				AuthInfo struct {
					SignerInfos []struct {
						PublicKey struct {
							Type string `json:"@type"`
							Key  string `json:"key"`
						} `json:"public_key"`
						ModeInfo struct {
							Single struct {
								Mode string `json:"mode"`
							} `json:"single"`
						} `json:"mode_info"`
						Sequence string `json:"sequence"`
					} `json:"signer_infos"`
					Fee struct {
						Amount   []interface{} `json:"amount"`
						GasLimit string        `json:"gas_limit"`
						Payer    string        `json:"payer"`
						Granter  string        `json:"granter"`
					} `json:"fee"`
					Tip interface{} `json:"tip"`
				} `json:"auth_info"`
				Signatures []string `json:"signatures"`
			} `json:"gen_txs"`
		} `json:"genutil"`
		Gov struct {
			DepositParams interface{}   `json:"deposit_params"`
			Deposits      []interface{} `json:"deposits"`
			Params        struct {
				BurnProposalDepositPrevote bool   `json:"burn_proposal_deposit_prevote"`
				BurnVoteQuorum             bool   `json:"burn_vote_quorum"`
				BurnVoteVeto               bool   `json:"burn_vote_veto"`
				MaxDepositPeriod           string `json:"max_deposit_period"`
				MinDeposit                 []struct {
					Amount string `json:"amount"`
					Denom  string `json:"denom"`
				} `json:"min_deposit"`
				MinInitialDepositRatio string `json:"min_initial_deposit_ratio"`
				Quorum                 string `json:"quorum"`
				Threshold              string `json:"threshold"`
				VetoThreshold          string `json:"veto_threshold"`
				VotingPeriod           string `json:"voting_period"`
			} `json:"params"`
			Proposals          []interface{} `json:"proposals"`
			StartingProposalId string        `json:"starting_proposal_id"`
			TallyParams        interface{}   `json:"tally_params"`
			Votes              []interface{} `json:"votes"`
			VotingParams       interface{}   `json:"voting_params"`
		} `json:"gov"`
		Group struct {
			GroupMembers   []interface{} `json:"group_members"`
			GroupPolicies  []interface{} `json:"group_policies"`
			GroupPolicySeq string        `json:"group_policy_seq"`
			GroupSeq       string        `json:"group_seq"`
			Groups         []interface{} `json:"groups"`
			ProposalSeq    string        `json:"proposal_seq"`
			Proposals      []interface{} `json:"proposals"`
			Votes          []interface{} `json:"votes"`
		} `json:"group"`
		Ibc struct {
			ChannelGenesis struct {
				AckSequences        []interface{} `json:"ack_sequences"`
				Acknowledgements    []interface{} `json:"acknowledgements"`
				Channels            []interface{} `json:"channels"`
				Commitments         []interface{} `json:"commitments"`
				NextChannelSequence string        `json:"next_channel_sequence"`
				Receipts            []interface{} `json:"receipts"`
				RecvSequences       []interface{} `json:"recv_sequences"`
				SendSequences       []interface{} `json:"send_sequences"`
			} `json:"channel_genesis"`
			ClientGenesis struct {
				Clients            []interface{} `json:"clients"`
				ClientsConsensus   []interface{} `json:"clients_consensus"`
				ClientsMetadata    []interface{} `json:"clients_metadata"`
				CreateLocalhost    bool          `json:"create_localhost"`
				NextClientSequence string        `json:"next_client_sequence"`
				Params             struct {
					AllowedClients []string `json:"allowed_clients"`
				} `json:"params"`
			} `json:"client_genesis"`
			ConnectionGenesis struct {
				ClientConnectionPaths  []interface{} `json:"client_connection_paths"`
				Connections            []interface{} `json:"connections"`
				NextConnectionSequence string        `json:"next_connection_sequence"`
				Params                 struct {
					MaxExpectedTimePerBlock string `json:"max_expected_time_per_block"`
				} `json:"params"`
			} `json:"connection_genesis"`
		} `json:"ibc"`
		Interchainaccounts struct {
			ControllerGenesisState struct {
				ActiveChannels     []interface{} `json:"active_channels"`
				InterchainAccounts []interface{} `json:"interchain_accounts"`
				Params             struct {
					ControllerEnabled bool `json:"controller_enabled"`
				} `json:"params"`
				Ports []interface{} `json:"ports"`
			} `json:"controller_genesis_state"`
			HostGenesisState struct {
				ActiveChannels     []interface{} `json:"active_channels"`
				InterchainAccounts []interface{} `json:"interchain_accounts"`
				Params             struct {
					AllowMessages []string `json:"allow_messages"`
					HostEnabled   bool     `json:"host_enabled"`
				} `json:"params"`
				Port string `json:"port"`
			} `json:"host_genesis_state"`
		} `json:"interchainaccounts"`
		Mint struct {
			Minter struct {
				AnnualProvisions string `json:"annual_provisions"`
				Inflation        string `json:"inflation"`
			} `json:"minter"`
			Params struct {
				BlocksPerYear       string `json:"blocks_per_year"`
				GoalBonded          string `json:"goal_bonded"`
				InflationMax        string `json:"inflation_max"`
				InflationMin        string `json:"inflation_min"`
				InflationRateChange string `json:"inflation_rate_change"`
				MintDenom           string `json:"mint_denom"`
				Reduction           struct {
					Enable          bool     `json:"enable"`
					TotalProvisions string   `json:"total_provisions"`
					Heights         []uint64 `json:"heights"`
				} `json:"reduction"`
			} `json:"params"`
		} `json:"mint"`
		Params   interface{} `json:"params"`
		Slashing struct {
			MissedBlocks []interface{} `json:"missed_blocks"`
			Params       struct {
				DowntimeJailDuration    string `json:"downtime_jail_duration"`
				MinSignedPerWindow      string `json:"min_signed_per_window"`
				SignedBlocksWindow      string `json:"signed_blocks_window"`
				SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
				SlashFractionDowntime   string `json:"slash_fraction_downtime"`
			} `json:"params"`
			SigningInfos []interface{} `json:"signing_infos"`
		} `json:"slashing"`
		Staking struct {
			Delegations         []interface{} `json:"delegations"`
			Exported            bool          `json:"exported"`
			LastTotalPower      string        `json:"last_total_power"`
			LastValidatorPowers []interface{} `json:"last_validator_powers"`
			Params              struct {
				BondDenom         string `json:"bond_denom"`
				HistoricalEntries int64  `json:"historical_entries"`
				MaxEntries        int64  `json:"max_entries"`
				MaxValidators     string `json:"max_validators"`
				MinCommissionRate string `json:"min_commission_rate"`
				UnbondingTime     string `json:"unbonding_time"`
			} `json:"params"`
			Redelegations        []interface{} `json:"redelegations"`
			UnbondingDelegations []interface{} `json:"unbonding_delegations"`
			Validators           []interface{} `json:"validators"`
		} `json:"staking"`
		Transfer struct {
			DenomTraces []interface{} `json:"denom_traces"`
			Params      struct {
				ReceiveEnabled bool `json:"receive_enabled"`
				SendEnabled    bool `json:"send_enabled"`
			} `json:"params"`
			PortId        string        `json:"port_id"`
			TotalEscrowed []interface{} `json:"total_escrowed"`
		} `json:"transfer"`
		Upgrade struct {
		} `json:"upgrade"`
		Vesting struct {
		} `json:"vesting"`
	} `json:"app_state"`
}
