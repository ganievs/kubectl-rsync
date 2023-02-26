package cmd

import (
	"github.com/openshift/oc/pkg/cli/rsync"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"
	"os"
)

func RootCmd() *cobra.Command {
	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	matchVersionKubeConfigFlags := kcmdutil.NewMatchVersionFlags(kubeConfigFlags)
	f := kcmdutil.NewFactory(matchVersionKubeConfigFlags)
	ioStreams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	rsyncCmd := rsync.NewCmdRsync(f, ioStreams)

	kubeConfigFlags.AddFlags(rsyncCmd.PersistentFlags())

	return rsyncCmd
}
