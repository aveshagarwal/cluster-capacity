package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cluster-capacity/cmd/options"
	"k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
)

func NewClusterCapacityCommand() *cobra.Command {
	s := options.NewClusterCapacityServer()
	cmd := &cobra.Command{
		Use:  "cluster-capacity",
		Long: `Cluster-capacity is used for emulating scheduling of one or multiple pods`,
		Run: func(cmd *cobra.Command, args []string) {
			err := Run(s)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	s.AddFlags(cmd.Flags())
	return cmd
}

func Run(s *options.ClusterCapacityServer) error {
	if len(s.Master) < 0 {
		return fmt.Errorf("master needs to be specified")
	}

	if len(s.Kubeconfig) < 0 {
		return fmt.Errorf("kubeconfig needs to be specified")
	}

	kubeconfig, err := clientcmd.BuildConfigFromFlags(s.Master, s.Kubeconfig)
	if err != nil {
		return fmt.Errorf("unable to build config from flags: %v", err)
	}

	kubeClient, err := unversioned.New(kubeconfig)
	if err != nil {
		return fmt.Errorf("Invalid API configuration: %v", err)
	}
	if v, err := kubeClient.Discovery().ServerVersion(); err != nil {
		return fmt.Errorf("Unable to get server version: %v\n", err)
	} else {
		return fmt.Errorf("Server version: %#v\n", v)
	}
	return nil
}
