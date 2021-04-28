package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func makeInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initialise new repository",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.Flags().String(
		pathFlag,
		".",
		"directory to operate in",
	)
	logIfError(cmd.MarkFlagRequired(pathFlag))
	logIfError(viper.BindPFlags(cmd.Flags()))
	return cmd
}
