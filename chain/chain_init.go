package chain

import (
	"github.com/civet148/cosmos-cli/api"
	"github.com/civet148/cosmos-cli/shells"
	"github.com/civet148/cosmos-cli/types"
	"github.com/civet148/cosmos-cli/utils"
	"github.com/civet148/log"
	"gopkg.in/yaml.v2"
	"os"
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

type InitChain struct {
	option            *InitOption //init option
	nNodeCount        int         //node count
	strNode0Home      string      //node0 home
	strNode0Validator string      //node0 validator
	strKeyFile        string      //key file to save
}

func NewInitChain(opt *InitOption) api.ManagerApi {
	if opt == nil {
		panic("init option is nil")
	}
	return &InitChain{
		option:     opt,
		strKeyFile: types.EXPORT_KEY_FILE,
	}
}

func (c *InitChain) Run() error {
	ic, err := c.parseConfig()
	if err != nil {
		return err
	}
	err = c.checkConfig(ic)
	if err != nil {
		return err
	}
	return c.initNodes(ic)
}

func (c *InitChain) parseConfig() (*types.IgniteConfig, error) {
	strConfig := c.option.ConfigPath
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
	return &ic, nil
}

func (c *InitChain) checkConfig(ic *types.IgniteConfig) error {

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
	c.strNode0Home = ic.Validators[0].Home
	c.strNode0Validator = ic.Validators[0].Name
	return nil
}

func (c *InitChain) initNodes(ic *types.IgniteConfig) (err error) {
	opt := c.option
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
		cmdline = maker.MakeCmdLineKeysAdd(v.Name, c.strNode0Home, reenter)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		reenter = false
		if v.Name != c.strNode0Validator {

			//make keyring file directory
			cmdline = maker.MakeCmdLineMkdirKeyringFile(v.Home)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
			//copy keys file to current validator keyring dir
			cmdline = maker.MakeCmdLineCopyKeysFile(c.strNode0Home, v.Home)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}

			//show validator account key
			var strPrivateKey string
			cmdline = maker.MakeCmdLineKeysShow(v.Name, v.Home)
			strPrivateKey, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
			log.Printf("------------------------------- validator %s private key -------------------------------", v.Name)
			log.Printf(strPrivateKey)
		}

		//add all validator genesis account to first validator
		balances := ic.GetAccountBalances(v.Name)
		cmdline = maker.MakeCmdLineAddGenesisAccount(v.Name, c.strNode0Home, balances)
		_, err = cmd.Shell(cmdline)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		if v.Name != c.strNode0Validator {
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

		if v.Name != c.strNode0Validator {
			//copy other gentx to first validator
			cmdline = maker.MakeCmdLineCopyGenTxJSON(v.Home, c.strNode0Home)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
		}
	}

	//collect gentxs for first validator
	cmdline := maker.MakeCmdLineCollectGenTxs(c.strNode0Home)
	_, err = cmd.Shell(cmdline)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	cmdline = maker.MakeCmdLineValidateGenesis(c.strNode0Home)
	_, err = cmd.Shell(cmdline)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	//sync genesis.json to every validator except first validator
	for _, v := range ic.Validators {
		if v.Name != c.strNode0Validator {
			cmdline = maker.MakeCmdLineCopyGenesisFile(c.strNode0Home, v.Home)
			_, err = cmd.Shell(cmdline)
			if err != nil {
				log.Errorf(err.Error())
				return
			}
		}
	}
	return nil
}
