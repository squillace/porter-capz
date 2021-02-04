package main

import (
	"github.com/spf13/cobra"
	"github.com/squillace/porter-capz/pkg/capz"
)

var (
	commandFile string
)

func buildInstallCommand(m *capz.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Execute the install functionality of this mixin",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			//Do something here if needed
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
