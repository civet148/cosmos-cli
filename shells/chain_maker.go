package shells

import (
	"fmt"
	"github.com/civet148/cosmos-cli/types"
	"github.com/civet148/log"
)

type ChainMaker struct {
	strChainID        string
	strNodeCmd        string
	strDefaultDenom   string
	strKeyPhrase      string
	strKeyringBackend string
}

func NewChainMaker(strNodeCmd, strChainID, strDefaultDenom, strKeyPhrase, strKeyringBackend string) *ChainMaker {
	if strNodeCmd == "" || strChainID == "" || strKeyPhrase == "" {
		log.Panic("node command, chain id and key phrase must not be empty")
	}
	return &ChainMaker{
		strChainID:        strChainID,
		strNodeCmd:        strNodeCmd,
		strDefaultDenom:   strDefaultDenom,
		strKeyPhrase:      strKeyPhrase,
		strKeyringBackend: strKeyringBackend,
	}
}

func (s *ChainMaker) NodeCmd() string {
	return s.strNodeCmd
}

func (s *ChainMaker) CopyCmd() string {
	return types.EXEC_CMD_COPY
}

func (s *ChainMaker) MkdirCmd() string {
	return types.EXEC_CMD_MKDIR
}

func (s *ChainMaker) MakeCmdLineConfigKeyringBackend(strHome string) string {
	return fmt.Sprintf("%s config keyring-backend %s --home %s", s.NodeCmd(), s.strKeyringBackend, strHome)
}

func (s *ChainMaker) MakeCmdLineConfigChainID(strHome string) string {
	return fmt.Sprintf("%s config chain-id %s --home %s", s.NodeCmd(), s.strChainID, strHome)
}

func (s *ChainMaker) MakeCmdLineInit(strMoniker, strHome string) string {
	if s.strDefaultDenom != "" {
		return fmt.Sprintf("%s init %s --chain-id %s --home %s --default-denom %s", s.NodeCmd(), strMoniker, s.strChainID, strHome, s.strDefaultDenom)
	}
	return fmt.Sprintf("%s init %s --chain-id %s --home %s", s.NodeCmd(), strMoniker, s.strChainID, strHome)
}

func (s *ChainMaker) MakeCmdLineKeysAdd(strAccName, strHome string, reenter, passwd bool) string {

	var strCmdLine string
	strSpawn := fmt.Sprintf("%s keys add %s --home %s --keyring-backend %s", s.NodeCmd(), strAccName, strHome, s.strKeyringBackend)
	if passwd {
		if reenter {
			strCmdLine = fmt.Sprintf(`
		expect <<-EOF
		spawn %s
		expect "Enter keyring passphrase"
		send "%s\r"
		expect "Re-enter keyring passphrase"
		send "%s\r"
		expect eof
		EOF
`, strSpawn, s.strKeyPhrase, s.strKeyPhrase)
		} else {
			strCmdLine = fmt.Sprintf(`
		expect <<-EOF
		spawn %s
		expect "Enter keyring passphrase"
		send "%s\r"
		expect eof
		EOF
`, strSpawn, s.strKeyPhrase)
		}
	} else {
		strCmdLine = strSpawn
	}
	return strCmdLine
}

func (s *ChainMaker) MakeCmdLineKeysShow(strAccName, strHome string) string {
	strSpawn := fmt.Sprintf("%s keys show %s --home %s --keyring-backend %s", s.NodeCmd(), strAccName, strHome, s.strKeyringBackend)
	strCmdLine := fmt.Sprintf(`
		expect <<-EOF
		spawn %s
		expect "Enter keyring passphrase"
		send "%s\r"
		expect eof
		EOF
`, strSpawn, s.strKeyPhrase)
	return strCmdLine
}

func (s *ChainMaker) MakeCmdLineKeysExport(strAccName, strHome string) string {
	strSpawn := fmt.Sprintf("%s keys export %s --home %s --keyring-backend %s", s.NodeCmd(), strAccName, strHome, s.strKeyringBackend)
	strCmdLine := fmt.Sprintf(`
		expect <<-EOF
		spawn %s
		expect "Enter passphrase to encrypt the exported key"
		send "%s\r"
		expect "Enter keyring passphrase"
		send "%s\r"
		expect eof
		EOF
`, strSpawn, s.strKeyPhrase, s.strKeyPhrase)
	return strCmdLine
}

func (s *ChainMaker) MakeCmdLineKeysImport(strAccName, strHome, strKeyFile string) string {
	strSpawn := fmt.Sprintf("%s keys import %s %s --home %s --keyring-backend %s", s.NodeCmd(), strAccName, strKeyFile, strHome, s.strKeyringBackend)
	strCmdLine := fmt.Sprintf(`
		expect <<-EOF
		spawn %s
		expect "Enter passphrase to decrypt your key"
		send "%s\r"
		expect "Enter keyring passphrase"
		send "%s\r"
		expect "Re-enter keyring passphrase"
  		send "%s\r"
		expect eof
		EOF
`, strSpawn, s.strKeyPhrase, s.strKeyPhrase, s.strKeyPhrase)
	return strCmdLine
}

func (s *ChainMaker) MakeCmdLineAddGenesisAccount(strAccName, strHome, strBalances string, passwd bool) string {
	var strCmdLine string
	strSpawn := fmt.Sprintf("%s add-genesis-account %s %s --home %s --keyring-backend %s", s.NodeCmd(), strAccName, strBalances, strHome, s.strKeyringBackend)
	if passwd {
		strCmdLine = fmt.Sprintf(`
		expect <<-EOF
		spawn %s
		expect "Enter keyring passphrase"
		send "%s\r"
		expect eof
		EOF
`, strSpawn, s.strKeyPhrase)
	} else {
		strCmdLine = strSpawn
	}
	return strCmdLine
}

func (s *ChainMaker) MakeCmdLineGenTx(strAccName, strHome, strStaking, strIP, strPort string, passwd bool) string {
	var strCmdLine string
	strSpawn := fmt.Sprintf("%s gentx %s %s --chain-id %s --ip %s --p2p-port %s --home %s --keyring-backend %s",
		s.NodeCmd(), strAccName, strStaking, s.strChainID, strIP, strPort, strHome, s.strKeyringBackend)
	if passwd {
		strCmdLine = fmt.Sprintf(`
		expect <<-EOF
		spawn %s
		expect "Enter keyring passphrase"
		send "%s\r"
		expect eof
		EOF
`, strSpawn, s.strKeyPhrase)
	} else {
		strCmdLine = strSpawn
	}
	return strCmdLine
}

func (s *ChainMaker) MakeCmdLineCopyGenTxJSON(strHomeSrc, strHomeDst string) string {
	return fmt.Sprintf("%s -f %s/config/gentx/*.json %s/config/gentx", s.CopyCmd(), strHomeSrc, strHomeDst)
}

func (s *ChainMaker) MakeCmdLineCollectGenTxs(strHome string) string {
	return fmt.Sprintf("%s collect-gentxs --home %s", s.NodeCmd(), strHome)
}

func (s *ChainMaker) MakeCmdLineValidateGenesis(strHome string) string {
	return fmt.Sprintf("%s validate-genesis --home %s", s.NodeCmd(), strHome)
}

func (s *ChainMaker) MakeCmdLineCopyGenesisFile(strHomeSrc, strHomeDst string) string {
	return fmt.Sprintf("%s -f %s/config/genesis.json %s/config", s.CopyCmd(), strHomeSrc, strHomeDst)
}

func (s *ChainMaker) MakeCmdLineCopyKeysFile(src, dest string) string {
	return fmt.Sprintf("%s -rf %s/keyring-* %s", s.CopyCmd(), src, dest)
}

func (s *ChainMaker) MakeCmdLineMkdirKeyringFile(strHome string) string {
	return fmt.Sprintf("%s -p %s/keyring-file", s.MkdirCmd(), strHome)
}

func (s *ChainMaker) MakeCmdLineShowNodeId(strHome string) string {
	return fmt.Sprintf("%s tendermint show-node-id --home %s", s.NodeCmd(), strHome)
}

func (s *ChainMaker) MakeCmdLineKeysShowAddrOnly(strHome string, name string, addrType string) string {
	var strCmdLine string
	strSpawn := fmt.Sprintf("%s keys show %s --bech %s -a --home %s", s.NodeCmd(), name, addrType, strHome)
	if s.strKeyringBackend != "test" {
		strCmdLine = fmt.Sprintf(`
		expect <<-EOF
		spawn %s
		expect "Enter keyring passphrase"
		send "%s\r"
		expect eof
		EOF
`, strSpawn, s.strKeyPhrase)
	} else {
		strCmdLine = strSpawn
	}
	return strCmdLine
}
