package main

import (
	"get.porter.sh/mixin/aws/pkg/aws"
	"github.com/spf13/cobra"
)

func buildInstallCommand(m *aws.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Execute the install functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute(cmd.Context())
		},
	}
	return cmd
}
