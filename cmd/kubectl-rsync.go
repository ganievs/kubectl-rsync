package main

import (
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
	"kubectl-rsync/pkg/cmd"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		kcmdutil.CheckErr(err)
	}
}
