package main

import (
	"fmt"
	"github.com/civet148/cosmos-cli/chain"
	"github.com/civet148/cosmos-cli/types"
	"github.com/civet148/cosmos-cli/utils"
	"github.com/civet148/log"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
)

const (
	Version     = "v0.1.0"
	ProgramName = "cosmos-cli"
)

var (
	BuildTime = "2023-10-11"
	GitCommit = ""
)

const (
	CMD_NAME_INIT = "init"
)

const (
	CMD_FLAG_NAME_CONFIG        = "config"
	CMD_FLAG_NAME_IGNITE_CMD    = "ignite-cmd"
	CMD_FLAG_NAME_NODE_CMD      = "node-cmd"
	CMD_FLAG_NAME_DEFAULT_DENOM = "default-denom"
	CMD_FLAG_NAME_CHAIN_ID      = "chain-id"
	CMD_FLAG_NAME_KEY_PHRASE    = "key-phrase"
)

func init() {
	log.SetLevel("debug")
}

func grace() {
	//capture signal of Ctrl+C and gracefully exit
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	go func() {
		for {
			select {
			case s := <-sigChannel:
				{
					if s != nil && s == os.Interrupt {
						fmt.Printf("Ctrl+C signal captured, program exiting...\n")
						close(sigChannel)
						os.Exit(0)
					}
				}
			}
		}
	}()
}

func main() {

	grace()

	local := []*cli.Command{
		initCmd,
	}
	app := &cli.App{
		Name:     ProgramName,
		Version:  fmt.Sprintf("%s %s commit %s", Version, BuildTime, GitCommit),
		Flags:    []cli.Flag{},
		Commands: local,
		Action:   nil,
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}

var initCmd = &cli.Command{
	Name:      CMD_NAME_INIT,
	Usage:     "set up cosmos nodes initializer",
	ArgsUsage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_CONFIG,
			Usage:   "config file path",
			Value:   types.DEFAULT_CONFIG_FILE,
			Aliases: []string{"c"},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_NODE_CMD,
			Usage:   "node command",
			Value:   types.DEFAULT_NODE_CMD,
			Aliases: []string{"n"},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_DEFAULT_DENOM,
			Usage:   "default denom for staking",
			Value:   types.DEFAULT_DENON,
			Aliases: []string{"d"},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_CHAIN_ID,
			Usage:   "chain id",
			Value:   types.DEFAULT_CHAIN_ID,
			Aliases: []string{""},
		},
		&cli.StringFlag{
			Name:    CMD_FLAG_NAME_KEY_PHRASE,
			Usage:   "pass phrase to protect keys",
			Value:   types.DEFAULT_KEY_PHRASE,
			Aliases: []string{"p"},
		},
	},
	Before: func(context *cli.Context) error {
		//check shells command installed or not before init chain
		cmd := utils.NewCmdExecutor(false)
		ok := cmd.Which(types.COMMAND_NAME_EXPECT)
		if !ok {
			if !cmd.Which(types.COMMAND_NAME_APT_GET) && !cmd.Which(types.COMMAND_NAME_APT) {
				return fmt.Errorf("%s command not found, please install expect first", types.COMMAND_NAME_EXPECT)
			}
			_, err := cmd.Shell("sudo apt-get install -y expect")
			if err != nil {
				return err
			}
		}
		return nil
	},
	Action: func(cctx *cli.Context) error {
		opt := &chain.InitOption{
			KeyringBackend: types.KEYRING_BACKEND,
			ConfigPath:     cctx.String(CMD_FLAG_NAME_CONFIG),
			IgniteCmd:      cctx.String(CMD_FLAG_NAME_IGNITE_CMD),
			NodeCmd:        cctx.String(CMD_FLAG_NAME_NODE_CMD),
			DefaultDenom:   cctx.String(CMD_FLAG_NAME_DEFAULT_DENOM),
			ChainID:        cctx.String(CMD_FLAG_NAME_CHAIN_ID),
			KeyPhrase:      cctx.String(CMD_FLAG_NAME_KEY_PHRASE),
		}
		service := chain.NewInitChain(opt)
		return service.Run()
	},
}
