package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var errNotImplemented = errors.New("not implemented")

func New() *cobra.Command {
	cmd := cobra.Command{
		Use:          "yf",
		Short:        "yf is a tool to format YAML files",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errNotImplemented
		},
	}
	return &cmd
}
