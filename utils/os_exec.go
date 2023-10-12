package utils

import (
	"fmt"
	"github.com/civet148/cosmos-cli/types"
	"github.com/civet148/log"
	"os/exec"
)

type CmdExecutor struct {
	Debug bool
}

func NewCmdExecutor(debug bool) *CmdExecutor {
	return &CmdExecutor{
		Debug: debug,
	}
}

func (m *CmdExecutor) Run(name string, args ...string) (output string, err error) {
	cmd := exec.Command(name, args...)
	log.Infof("executing [%s %v]...", name, FmtStringArgs(args...))
	var data []byte
	data, err = cmd.CombinedOutput()
	output = string(data)
	if err != nil {
		log.Errorf("execute command line [%s %v] error [%s]", name, FmtStringArgs(args...), err.Error())
		log.Printf(output)
		return
	}
	if m.Debug {
		log.Printf(output)
	}
	return
}

func (m *CmdExecutor) Shell(cmdline string) (output string, err error) {
	var args []string
	args = append(args, types.EXEC_SHELL_ARG)
	args = append(args, cmdline)
	return m.Run(types.EXEC_CMD_SHELL, args...)
}

func FmtStringArgs(args ...string) string {
	var as []interface{}
	for _, v := range args {
		as = append(as, " "+v)
	}
	return fmt.Sprint(as...)
}
