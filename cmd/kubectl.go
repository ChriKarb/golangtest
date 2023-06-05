package cmd

import (
	"github.com/ckone4you/golangtest/pkg/kubectl"
	"github.com/spf13/cobra"
)

func NewKubectlCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "kubectl [namespace]",
		Short: "List pods in the specified namespace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace := args[0]
			return kubectl.ListPods(namespace)
		}, // Implement kubectl subcommands here

	}
}
