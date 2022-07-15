package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/usmhk/mycni/pkg/args"
	"github.com/usmhk/mycni/pkg/hanlder"
)

const (
	IPStore = "/root"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fmt.Println("hello mycni!")
	cmd, cmdArgs, err := args.GetArgsFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "getting cmd arguments with error: %v", err)
	}
	_, _ = cmd, cmdArgs
	fh := hanlder.NewFileHandler(IPStore)

	switch cmd {
	case "ADD":
		err = fh.HandleAdd(cmdArgs)
	case "DEL":
		err = fh.HandleDel(cmdArgs)
	case "CHECK":
		err = fh.HandleCheck(cmdArgs)

	case "VERSION", "version", "v":
		err = fh.HandleVersion(cmdArgs)
	default:
		err = fmt.Errorf("unknow CNI_COMMAND: %s", cmd)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to handle CNI_COMMAND %q: %v", cmd, err)
		os.Exit(1)
	}
}
