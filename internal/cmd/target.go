package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	targetFlag = "target"
	envFlag    = "env"
)

func makeTargetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "target",
		Short:         "manage target",
		SilenceErrors: true,
	}
	cmd.AddCommand(makeTargetAddCmd())
	return cmd
}

func makeTargetAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a target to an environment",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("target %s added to environment %s\n", viper.GetString(targetFlag), viper.GetString(envFlag))
		},
	}

	cmd.Flags().String(
		dirFlag,
		".",
		"directory to operate in",
	)

	cmd.Flags().String(
		targetFlag,
		".",
		"name of target to add to the environment",
	)

	cmd.Flags().String(
		envFlag,
		".",
		"name of the environment to add a target to",
	)

	logIfError(cmd.MarkFlagRequired(targetFlag))
	logIfError(viper.BindPFlags(cmd.Flags()))
	return cmd
}
