package chain

import (
	"fmt"
	"github.com/civet148/cosmos-cli/api"
	"github.com/civet148/cosmos-cli/confile"
	"github.com/civet148/cosmos-cli/shells"
	"github.com/civet148/cosmos-cli/types"
	"github.com/civet148/cosmos-cli/utils"
	"github.com/civet148/log"
	"github.com/imdario/mergo"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type ChainBuilder struct {
	peers             map[string]*types.NodePeer //node peer information
	option            *types.Option              //init option
	nNodeCount        int                        //node count
	strNode0Home      string                     //node0 home
	strNode0Validator string                     //node0 validator
	strKeyFile        string                     //key file to save
	igniteConfigs     map[string]interface{}     //ignite config map
}

func NewChainBuilder(opt *types.Option) api.ManagerApi {
	if opt == nil {
		panic("init option is nil")
	}
	return &ChainBuilder{
		option:        opt,
		peers:         make(map[string]*types.NodePeer),
		strKeyFile:    types.EXPORT_KEY_FILE,
		igniteConfigs: make(map[string]interface{}),
	}
}

func (m *ChainBuilder) Run() (err error) {
	var ic *types.IgniteConfig
	if ic, err = m.parseConfig(); err != nil {
		return log.Errorf(err.Error())
	}
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
	err = m.mergeGenesisConfig(ic)
	if err != nil {
		return err
	}
	err = m.syncGenesisFile(ic)
	if err != nil {
		return err
	}
	return nil
}

func (m *ChainBuilder) parseConfig() (ic *types.IgniteConfig, err error) {
	vip := viper.New()
	strPath := m.option.ConfigPath
	vip.SetConfigFile(strPath)
	vip.SetConfigType("yaml")
	if err = vip.ReadInConfig(); err != nil {
		return nil, log.Errorf("load config [%s] error [%s]", strPath, err.Error())
	}
	m.igniteConfigs = vip.AllSettings()
	strConfig := m.option.ConfigPath
	data, err := os.ReadFile(strConfig)
	if err != nil {
		return nil, log.Errorf("open config file %s error [%v]", strConfig, err)
	}
	ic = &types.IgniteConfig{}
	err = yaml.Unmarshal(data, ic)
	if err != nil {
		return nil, log.Errorf("unmarshal config file %s error [%v]", strConfig, err)
	}
	if ic.Genesis.ChainID != m.option.ChainID {
		m.option.ChainID = ic.Genesis.ChainID
	}
	return ic, nil
}

func (m *ChainBuilder) checkConfig(ic *types.IgniteConfig) error {

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

func (m *ChainBuilder) initNodes(ic *types.IgniteConfig) (err error) {
	opt := m.option
	maker := shells.NewChainMaker(opt.NodeCmd, opt.ChainID, opt.DefaultDenom, opt.KeyPhrase, opt.KeyringBackend)

	cmd := utils.NewCmdExecutor(opt.Debug)
	var reenter = true
	var cmdline string
	var passwd = true
	if m.option.KeyringBackend == types.KEYRING_BACKEND_TEST {
		passwd = false
	}
	for _, v := range ic.Validators {
		err = os.RemoveAll(v.Home)
		if err != nil {
			log.Errorf(err.Error())
			return
		}

		//chain config and data init
		cmdline = maker.MakeCmdLineConfigKeyringBackend(v.Home)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		cmdline = maker.MakeCmdLineConfigChainID(v.Home)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		cmdline = maker.MakeCmdLineInit(v.Config.Moniker, v.Home)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		//add all validator account key to first validator keyring
		cmdline = maker.MakeCmdLineKeysAdd(v.Name, m.strNode0Home, reenter, passwd)
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
		cmdline = maker.MakeCmdLineAddGenesisAccount(v.Name, m.strNode0Home, balances, passwd)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		if v.Name != m.strNode0Validator {
			//add self validator genesis account
			cmdline = maker.MakeCmdLineAddGenesisAccount(v.Name, v.Home, balances, passwd)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
		}
		//gen genesis tx for every validator
		strPort := utils.ParseP2PPort(v.Config.P2P.Laddr)
		cmdline = maker.MakeCmdLineGenTx(v.Name, v.Home, v.Bonded, v.IP, strPort, passwd)
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
		np := &types.NodePeer{
			Name: v.Name,
			Peer: fmt.Sprintf("%s@%s", strNodeId, ic.GetValidatorHost(v.Name)),
		}
		m.peers[v.Name] = np
	}

	//collect gentxs for first validator
	cmdline = maker.MakeCmdLineCollectGenTxs(m.strNode0Home)
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

func (m *ChainBuilder) updateAppConfig(ic *types.IgniteConfig) (err error) {
	vals := m.igniteConfigs["validators"].([]interface{})
	for i, v := range ic.Validators {
		vip := viper.New()
		strPath := utils.MakeCosmosConfigPath(v.Home, types.FILE_NAME_APP)
		vip.SetConfigFile(strPath)
		vip.SetConfigType("toml")
		if err = vip.ReadInConfig(); err != nil {
			return log.Errorf("load config [%s] error [%s]", strPath, err.Error())
		}
		var genesis = make(map[string]interface{})
		genesis = vip.AllSettings()
		cf := confile.New(confile.DefaultTOMLEncodingCreator, strPath)
		if err = cf.Load(&genesis); err != nil {
			return err
		}
		igniteSettings := vals[i].(map[string]interface{})
		conf := igniteSettings["app"]
		log.Json("app config to update", conf)
		if err = mergo.Merge(&genesis, conf, mergo.WithOverride); err != nil {
			return err
		}
		err = cf.Save(genesis)
	}
	return
}

func (m *ChainBuilder) updateCosmosConfig(ic *types.IgniteConfig) (err error) {
	vals := m.igniteConfigs["validators"].([]interface{})
	for i, v := range ic.Validators {
		vip := viper.New()
		strPath := utils.MakeCosmosConfigPath(v.Home, types.FILE_NAME_CONFIG)
		vip.SetConfigFile(strPath)
		vip.SetConfigType("toml")
		if err = vip.ReadInConfig(); err != nil {
			return log.Errorf("load config [%s] error [%s]", strPath, err.Error())
		}
		//update validator p2p persistent peers
		var strPeers string
		if np, ok := m.peers[v.Name]; ok {
			strPeers = strings.Join(np.PersistentPeers, ",")
			vip.Set("p2p.persistent_peers", strPeers)
			log.Infof("[%s] p2p.persistent_peers=%s", v.Name, strPeers)
		}
		var genesis = make(map[string]interface{})
		genesis = vip.AllSettings()
		cf := confile.New(confile.DefaultTOMLEncodingCreator, strPath)
		if err = cf.Load(&genesis); err != nil {
			return err
		}
		igniteSettings := vals[i].(map[string]interface{})
		conf := igniteSettings["config"].(map[string]interface{})
		p2p := conf["p2p"].(map[string]interface{})
		p2p["persistent_peers"] = strPeers
		log.Json("cosmos config to update", conf)
		if err = mergo.Merge(&genesis, conf, mergo.WithOverride); err != nil {
			return err
		}
		err = cf.Save(genesis)
	}
	return
}

func (m *ChainBuilder) mergeGenesisConfig(ic *types.IgniteConfig) (err error) {
	for _, v := range ic.Validators {
		vip := viper.New()
		strPath := utils.MakeCosmosConfigPath(v.Home, types.FILE_NAME_GENESIS)
		vip.SetConfigFile(strPath)
		vip.SetConfigType("json")
		if err = vip.ReadInConfig(); err != nil {
			return log.Errorf("load config [%s] error [%s]", strPath, err.Error())
		}
		var genesis = make(map[string]interface{})
		genesis = vip.AllSettings()

		cf := confile.New(confile.DefaultJSONEncodingCreator, strPath)
		if err = cf.Load(&genesis); err != nil {
			return err
		}
		igniteSettings := m.igniteConfigs["genesis"]
		if err = mergo.Merge(&genesis, igniteSettings, mergo.WithOverride); err != nil {
			return err
		}
		err = cf.SaveJSON(genesis)
	}
	return nil
}

func (m *ChainBuilder) syncGenesisFile(ic *types.IgniteConfig) (err error) {
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
