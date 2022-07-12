package main

import (
	"fmt"
	"os"

	"github.com/usmhk/mycni/pkg/args"
)

func main() {
	fmt.Println("hello mycni!")
	cmd, cmdArgs, err := args.GetArgsFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "getting cmd arguments with error: %v", err)
	}
	_, _ = cmd, cmdArgs

}
