package main

import (
	"os"

	"github.com/eduroyu/GolangBlockChain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
