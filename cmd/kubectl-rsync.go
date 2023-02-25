package main

import (
	"kubectl-rsync/pkg/cmd"
)

func main() {
	cmd.RootCmd().Execute()
}
