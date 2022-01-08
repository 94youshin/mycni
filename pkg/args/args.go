package args

import (
	"fmt"
	"os"
)

const (
	CommandEnvKey string = ""
)

const (
	AddCmd     string = "ADD"
	DelCmd     string = "DEL"
	CheckCmd   string = "CHECK"
	VersionCmd string = "VERSION"
)

type CmdEnv struct {
	CmdArgKey   string
	CmdArgValue *string
	ReqForCmd   map[string]bool
}

type CmdArgs struct {
	CmdArgKey string
	Netns     string
	IfName    string
	Path      string
	Args      string
	StdinData string
}

type CNIConfiguration struct {
	CniVersion string `json:"cniVersion"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Bridge     string `json:"bridge"`
	MTU        string `json:"mtu"`
	Subnet     string `json:"subnet"`
}

func GetArgsFromEnv() (string, *CmdArgs, error) {
	var errMsg, cmd, conID, netns, ifName, path, args string
	cmd = os.Getenv(CommandEnvKey)
	if len(cmd) == 0 {
		errMsg = fmt.Sprintf("Environment variable %s is missing!", CommandEnvKey)
		fmt.Fprint(os.Stderr, errMsg)
		return "", nil, errors.new(errMsg)
	}
}
