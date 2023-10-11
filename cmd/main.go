package main

import (
	"fmt"
	"github.com/civet148/cosmos-cli/chain"
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
	CMD_FLAG_NAME_CONFIG = "config"
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
			Usage:   "ignite config file path",
			Value:   "config.yml",
			Aliases: []string{"c"},
		},
	},
	Action: func(cctx *cli.Context) error {
		opt := &chain.InitOption{
			IgniteConfigPath: cctx.String(CMD_FLAG_NAME_CONFIG),
		}
		service := chain.NewInitChain(opt)
		return service.Run()
	},
}
