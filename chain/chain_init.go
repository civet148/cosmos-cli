package chain

import (
	"fmt"
	"github.com/civet148/cosmos-cli/api"
	"github.com/civet148/cosmos-cli/shells"
	"github.com/civet148/cosmos-cli/types"
	"github.com/civet148/cosmos-cli/utils"
	"github.com/civet148/log"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
	"strings"
)

type InitOption struct {
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

type InitChain struct {
	peers             map[string]*NodePeer //node peer information
	option            *InitOption          //init option
	nNodeCount        int                  //node count
	strNode0Home      string               //node0 home
	strNode0Validator string               //node0 validator
	strKeyFile        string               //key file to save
}

func NewInitChain(opt *InitOption) api.ManagerApi {
	if opt == nil {
		panic("init option is nil")
	}
	return &InitChain{
		option:     opt,
		peers:      make(map[string]*NodePeer),
		strKeyFile: types.EXPORT_KEY_FILE,
	}
}

func (m *InitChain) Run() error {
	ic, err := m.parseConfig()
	if err != nil {
		return err
	}
	err = m.checkConfig(ic)
	if err != nil {
		return err
	}
	err = m.initNodes(ic)
	if err != nil {
		return err
	}
	err = m.updateAppConfig(ic)
	if err != nil {
		return err
	}
	err = m.updateCosmosConfig(ic)
	if err != nil {
		return err
	}
	err = m.updateGenesisConfig(ic)
	if err != nil {
		return err
	}
	err = m.syncGenesisFile(ic)
	if err != nil {
		return err
	}
	return nil
}

func (m *InitChain) parseConfig() (*types.IgniteConfig, error) {
	strConfig := m.option.ConfigPath
	data, err := os.ReadFile(strConfig)
	if err != nil {
		return nil, log.Errorf("open config file %s error [%v]", strConfig, err)
	}
	ic := types.IgniteConfig{}
	err = yaml.Unmarshal(data, &ic)
	if err != nil {
		return nil, log.Errorf("unmarshal config file %s error [%v]", strConfig, err)
	}
	//log.Json("ignite config json", ic)
	if ic.Genesis.ChainID != m.option.ChainID {
		m.option.ChainID = ic.Genesis.ChainID
	}
	return &ic, nil
}

func (m *InitChain) checkConfig(ic *types.IgniteConfig) error {

	if len(ic.Validators) == 0 {
		return log.Errorf("validators must not be empty")
	}
	if len(ic.Accounts) == 0 {
		return log.Errorf("accounts must not be empty")
	}
	if len(ic.Validators) > len(ic.Accounts) {
		return log.Errorf("validator count is more than accounts")
	}

	for _, v := range ic.Validators {
		var ok bool
		for _, a := range ic.Accounts {
			if v.Name == a.Name {
				ok = true
			}
		}
		if ok == false {
			return log.Errorf("validator [%v] account not exist", v.Name)
		}
		if v.IP == "" {
			return log.Errorf("validator [%v] ip is empty", v.Name)
		}
		if v.Home == "" {
			return log.Errorf("validator [%v] home is empty", v.Name)
		}
		if v.Config.P2P.Laddr == "" {
			return log.Errorf("validator [%v] config p2p listen address is empty", v.Name)
		}
		if v.Config.RPC.Laddr == "" {
			return log.Errorf("validator [%v] config rpc listen address is empty", v.Name)
		}
		if v.Config.Moniker == "" {
			return log.Errorf("validator [%v] config moniker is empty", v.Name)
		}
		if v.Config.Consensus.TimeoutCommit == "" {
			return log.Errorf("validator [%v] config timeout commit is empty", v.Name)
		}
		if v.Bonded == "" {
			return log.Errorf("validator [%v] bonded staking is empty", v.Name)
		}
	}
	for _, v := range ic.Accounts {
		if len(v.Coins) == 0 {
			return log.Errorf("account [%v] coins is empty", v.Name)
		}
	}
	strHome := ic.Validators[0].Home
	if strings.HasPrefix(strHome, "~") {
		strHome = strings.Replace(strHome, "~", "$HOME", -1)
		strHome = os.ExpandEnv(strHome)
		ic.Validators[0].Home = strHome
	}
	m.strNode0Home = ic.Validators[0].Home
	m.strNode0Validator = ic.Validators[0].Name
	return nil
}

func (m *InitChain) initNodes(ic *types.IgniteConfig) (err error) {
	opt := m.option
	maker := shells.NewChainMaker(opt.NodeCmd, opt.ChainID, opt.DefaultDenom, opt.KeyPhrase, opt.KeyringBackend)

	cmd := utils.NewCmdExecutor(opt.Debug)
	var reenter = true
	for _, v := range ic.Validators {
		err = os.RemoveAll(v.Home)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		//chain config and data init
		cmdline := maker.MakeCmdLineInit(v.Config.Moniker, v.Home)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		//add all validator account key to first validator keyring
		cmdline = maker.MakeCmdLineKeysAdd(v.Name, m.strNode0Home, reenter)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		reenter = false
		if v.Name != m.strNode0Validator {

			//make keyring file directory
			cmdline = maker.MakeCmdLineMkdirKeyringFile(v.Home)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
			//copy keys file to current validator keyring dir
			cmdline = maker.MakeCmdLineCopyKeysFile(m.strNode0Home, v.Home)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
		}

		//add all validator genesis account to first validator
		balances := ic.GetAccountBalances(v.Name)
		cmdline = maker.MakeCmdLineAddGenesisAccount(v.Name, m.strNode0Home, balances)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		if v.Name != m.strNode0Validator {
			//add self validator genesis account
			cmdline = maker.MakeCmdLineAddGenesisAccount(v.Name, v.Home, balances)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
		}
		//gen genesis tx for every validator
		strPort := utils.ParseP2PPort(v.Config.P2P.Laddr)
		cmdline = maker.MakeCmdLineGenTx(v.Name, v.Home, v.Bonded, v.IP, strPort)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		//collect gentxs for every validator
		cmdline = maker.MakeCmdLineCollectGenTxs(v.Home)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}

		if v.Name != m.strNode0Validator {
			//copy other gentx to first validator
			cmdline = maker.MakeCmdLineCopyGenTxJSON(v.Home, m.strNode0Home)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
		}
		//get node id and make peer info
		var strNodeId string
		cmdline = maker.MakeCmdLineShowNodeId(v.Home)
		strNodeId, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		np := &NodePeer{
			Name: v.Name,
			Peer: fmt.Sprintf("%s@%s", strNodeId, ic.GetValidatorHost(v.Name)),
		}
		m.peers[v.Name] = np
	}

	//collect gentxs for first validator
	cmdline := maker.MakeCmdLineCollectGenTxs(m.strNode0Home)
	_, err = cmd.Shell(cmdline)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	cmdline = maker.MakeCmdLineValidateGenesis(m.strNode0Home)
	_, err = cmd.Shell(cmdline)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	for _, v := range ic.Validators {
		//update persistent peers
		np := m.peers[v.Name]
		if np == nil {
			return log.Errorf("validator %s peer info not found", v.Name)
		}
		for _, p := range m.peers {
			if p.Name != v.Name {
				np.PersistentPeers = append(np.PersistentPeers, p.Peer)
			}
		}
	}
	return nil
}

func (m *InitChain) updateAppConfig(ic *types.IgniteConfig) (err error) {
	for _, v := range ic.Validators {
		vip := viper.New()
		strPath := utils.MakeCosmosConfigPath(v.Home, types.FILE_NAME_APP)
		vip.SetConfigFile(strPath)
		vip.SetConfigType("toml")
		if err = vip.ReadInConfig(); err != nil {
			return log.Errorf("load config [%s] error [%s]", strPath, err.Error())
		}

		//if v.App.MinimumGasPrices != "" {
		vip.Set("minimum-gas-prices", v.App.MinimumGasPrices)
		//}
		if v.App.API.Enable {
			vip.Set("api.enable", v.App.API.Enable)
		}
		if v.App.API.Address != "" {
			vip.Set("api.address", v.App.API.Address)
		}
		if v.App.API.EnabledUnsafeCors {
			vip.Set("api.enabled-unsafe-cors", v.App.API.EnabledUnsafeCors)
		}
		if v.App.Grpc.Enable {
			vip.Set("grpc.enable", v.App.Grpc.Enable)
		}
		if v.App.Grpc.Address != "" {
			vip.Set("grpc.address", v.App.Grpc.Address)
		}
		if v.App.GrpcWeb.Enable {
			vip.Set("grpc-web.enable", v.App.GrpcWeb.Enable)
		}
		if v.App.GrpcWeb.Address != "" {
			vip.Set("grpc-web.address", v.App.GrpcWeb.Address)
		}
		if v.App.GrpcWeb.EnableUnsafeCors {
			vip.Set("grpc-web.enable-unsafe-cors", v.App.GrpcWeb.EnableUnsafeCors)
		}

		//save to app.toml file
		if err = vip.WriteConfig(); err != nil {
			return log.Errorf(err.Error())
		}
	}
	return
}

func (m *InitChain) updateCosmosConfig(ic *types.IgniteConfig) (err error) {
	for _, v := range ic.Validators {

		vip := viper.New()
		strPath := utils.MakeCosmosConfigPath(v.Home, types.FILE_NAME_CONFIG)
		vip.SetConfigFile(strPath)
		vip.SetConfigType("toml")
		if err = vip.ReadInConfig(); err != nil {
			return log.Errorf("load config [%s] error [%s]", strPath, err.Error())
		}

		if v.Config.Moniker != "" {
			vip.Set("moniker", v.Config.Moniker)
		}
		if v.Config.ProxyApp != "" {
			vip.Set("proxy_app", v.Config.ProxyApp)
		}
		if v.Config.Consensus.TimeoutCommit != "" {
			vip.Set("consensus.timeout_commit", v.Config.Consensus.TimeoutCommit)
		}
		if v.Config.RPC.Laddr != "" {
			vip.Set("rpc.laddr", v.Config.RPC.Laddr)
		}
		if v.Config.RPC.MaxBodyBytes != "" {
			nMaxBodySize, err := strconv.Atoi(v.Config.RPC.MaxBodyBytes)
			if err != nil {
				return log.Errorf("rpc.max_body_bytes %s invalid number", v.Config.RPC.MaxBodyBytes)
			}
			vip.Set("rpc.max_body_bytes", nMaxBodySize)
		}
		if v.Config.P2P.Laddr != "" {
			vip.Set("p2p.laddr", v.Config.P2P.Laddr)
		}
		if v.Config.P2P.AllowDuplicateIP {
			vip.Set("p2p.allow_duplicate_ip", v.Config.P2P.AllowDuplicateIP)
		}
		if v.Config.Instrumentation.Prometheus {
			vip.Set("instrumentation.prometheus", v.Config.Instrumentation.Prometheus)
		}
		if v.Config.Instrumentation.PrometheusListenAddr != "" {
			vip.Set("instrumentation.prometheus_listen_addr", v.Config.Instrumentation.PrometheusListenAddr)
		}
		//update validator p2p persistent peers
		if np, ok := m.peers[v.Name]; ok {
			strPeers := strings.Join(np.PersistentPeers, ",")
			vip.Set("p2p.persistent_peers", strPeers)
		}
		//save to config.toml file
		if err = vip.WriteConfig(); err != nil {
			return log.Errorf(err.Error())
		}
	}
	return
}

func (m *InitChain) updateGenesisConfig(ic *types.IgniteConfig) (err error) {
	for _, v := range ic.Validators {
		vip := viper.New()
		strPath := utils.MakeCosmosConfigPath(v.Home, types.FILE_NAME_GENESIS)
		vip.SetConfigFile(strPath)
		vip.SetConfigType("json")
		if err = vip.ReadInConfig(); err != nil {
			return log.Errorf("load config [%s] error [%s]", strPath, err.Error())
		}

		genesis := ic.Genesis
		gov := ic.Genesis.AppState.Gov
		mint := ic.Genesis.AppState.Mint
		crisis := ic.Genesis.AppState.Crisis
		staking := ic.Genesis.AppState.Staking
		distr := ic.Genesis.AppState.Distribution
		bank := ic.Genesis.AppState.Bank
		consens := ic.Genesis.ConsensusParams

		//handle global parameters
		if genesis.ChainID != "" {
			vip.Set("chain_id", genesis.ChainID)
		}
		if genesis.InitialHeight != "" {
			vip.Set("initial_height", genesis.InitialHeight)
		}
		if genesis.GenesisTime != "" {
			vip.Set("genesis_time", genesis.GenesisTime)
		}

		//handle consensus parameters
		if consens.Version.App != "" {
			vip.Set("consensus_params.version.app", consens.Version.App)
		}
		if consens.Block.MaxBytes != "" {
			vip.Set("consensus_params.block.max_bytes", consens.Block.MaxBytes)
		}
		if consens.Block.MaxGas != "" {
			vip.Set("consensus_params.block.max_gas", consens.Block.MaxGas)
		}
		if len(consens.Validator.PubKeyTypes) != 0 {
			vip.Set("consensus_params.validator.pub_key_types", consens.Validator.PubKeyTypes)
		}
		if consens.Evidence.MaxAgeNumBlocks != "" {
			vip.Set("consensus_params.evidence.max_age_num_blocks", consens.Evidence.MaxAgeNumBlocks)
		}
		if consens.Evidence.MaxAgeDuration != "" {
			vip.Set("consensus_params.evidence.max_age_duration", consens.Evidence.MaxAgeDuration)
		}
		if consens.Evidence.MaxBytes != "" {
			vip.Set("consensus_params.evidence.max_bytes", consens.Evidence.MaxBytes)
		}

		//handle genesis app state of bank
		if len(bank.DenomMetadata) != 0 {
			vip.Set("app_state.bank.denom_metadata", bank.DenomMetadata)
		}

		//handle genesis app state of staking
		if staking.Params.BondDenom != "" {
			vip.Set("app_state.staking.params.bond_denom", staking.Params.BondDenom)
		}
		if staking.Params.MaxValidators != "" {
			count, _ := strconv.Atoi(staking.Params.MaxValidators)
			vip.Set("app_state.staking.params.max_validators", count)
		}

		//handle genesis app state of mint
		if mint.Minter.AnnualProvisions != "" {
			vip.Set("app_state.mint.minter.annual_provisions", mint.Minter.AnnualProvisions)
		}
		if mint.Minter.Inflation != "" {
			vip.Set("app_state.mint.minter.inflation", mint.Minter.Inflation)
		}
		if mint.Params.MintDenom != "" {
			vip.Set("app_state.mint.params.mint_denom", mint.Params.MintDenom)
		}
		if mint.Params.BlocksPerYear != "" {
			vip.Set("app_state.mint.params.blocks_per_year", mint.Params.BlocksPerYear)
		}
		if mint.Params.GoalBonded != "" {
			vip.Set("app_state.mint.params.goal_bonded", mint.Params.GoalBonded)
		}
		if mint.Params.InflationMax != "" {
			vip.Set("app_state.mint.params.inflation_max", mint.Params.InflationMax)
		}
		if mint.Params.InflationMin != "" {
			vip.Set("app_state.mint.params.inflation_min", mint.Params.InflationMin)
		}
		if mint.Params.Reduction.Enable {
			vip.Set("app_state.mint.params.reduction.enable", mint.Params.Reduction.Enable)
			vip.Set("app_state.mint.params.reduction.total_provisions", mint.Params.Reduction.TotalProvisions)
			vip.Set("app_state.mint.params.reduction.heights", mint.Params.Reduction.Heights)
		}

		//handle genesis app state of gov
		if len(gov.Params.MinDeposit) != 0 {
			vip.Set("app_state.gov.params.min_deposit", gov.Params.MinDeposit)
		}

		//handle genesis app state of distribution
		if distr.Params.BaseProposerReward != "" {
			vip.Set("app_state.distribution.params.base_proposer_reward", distr.Params.BaseProposerReward)
		}
		if distr.Params.BonusProposerReward != "" {
			vip.Set("app_state.distribution.params.bonus_proposer_reward", distr.Params.BonusProposerReward)
		}
		if distr.Params.CommunityTax != "" {
			vip.Set("app_state.distribution.params.community_tax", distr.Params.CommunityTax)
		}
		if distr.Params.WithdrawAddrEnabled {
			vip.Set("app_state.distribution.params.withdraw_addr_enabled", distr.Params.WithdrawAddrEnabled)
		}

		//handle genesis app state of crisis
		if crisis.ConstantFee.Amount != "" {
			vip.Set("app_state.crisis.constant_fee.amount", crisis.ConstantFee.Amount)
		}
		if crisis.ConstantFee.Denom != "" {
			vip.Set("app_state.crisis.constant_fee.denom", crisis.ConstantFee.Denom)
		}

		//save to genesis.json file
		if err = vip.WriteConfig(); err != nil {
			return log.Errorf(err.Error())
		}
	}
	return nil
}

func (m *InitChain) syncGenesisFile(ic *types.IgniteConfig) (err error) {
	opt := m.option
	cmd := utils.NewCmdExecutor(opt.Debug)
	maker := shells.NewChainMaker(opt.NodeCmd, opt.ChainID, opt.DefaultDenom, opt.KeyPhrase, opt.KeyringBackend)
	for _, v := range ic.Validators {
		//sync genesis.json to every validator except first validator
		if v.Name != m.strNode0Validator {
			cmdline := maker.MakeCmdLineCopyGenesisFile(m.strNode0Home, v.Home)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				return log.Errorf(err.Error())
			}
		}
	}
	return nil
}
