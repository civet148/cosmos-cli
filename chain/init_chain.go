package chain

import (
	"github.com/civet148/cosmos-cli/api"
	"github.com/civet148/cosmos-cli/types"
	"github.com/civet148/log"
	"gopkg.in/yaml.v2"
	"os"
)

type InitOption struct {
	IgniteConfigPath string
}

type InitChain struct {
	option *InitOption
}

func NewInitChain(opt *InitOption) api.ManagerApi {
	if opt == nil {
		panic("init option is nil")
	}
	return &InitChain{option: opt}
}

func (c *InitChain) Run() error {
	data, err := os.ReadFile(c.option.IgniteConfigPath)
	if err != nil {
		return log.Errorf("open ignite config file error [%v]", err)
	}
	ic := types.IgniteConfig{}
	err = yaml.Unmarshal(data, &ic)
	if err != nil {
		return log.Errorf("unmarshal ignite config file error [%v]", err)
	}
	log.Json("ignite config json", ic)
	return nil
}
