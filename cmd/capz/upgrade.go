package main

import (
	"github.com/spf13/cobra"
	"github.com/squillace/porter-capz/pkg/capz"
)

func buildUpgradeCommand(m *capz.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
