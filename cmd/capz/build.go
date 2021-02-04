package main

import (
	"github.com/spf13/cobra"
	"github.com/squillace/porter-capz/pkg/capz"
)

func buildBuildCommand(m *capz.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Build()
		},
	}
	return cmd
}
