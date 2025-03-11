package cmd

import (
	"fmt"
	"io"

	"github.com/nokamoto/yf/internal/yaml"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := cobra.Command{
		Use:          "yf",
		Short:        "yf is a tool to format YAML files",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			in, err := io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}
			out, err := yaml.Format(in, 2)
			if err != nil {
				return fmt.Errorf("failed to format: %w", err)
			}
			cmd.OutOrStdout().Write(out)
			return nil
		},
	}
	return &cmd
}
