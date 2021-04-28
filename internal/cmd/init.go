package cmd

import (
	"fmt"
	"log"

	"github.com/bigkevmcd/kustomapp-api/pkg/tree"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envsFlag   = "env"
	outputFlag = "output"
)

func makeInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initialise new repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := tree.InitialiseDirectory(viper.GetString(outputFlag), viper.GetStringSlice(envsFlag)); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("app initialised")
		},
	}

	cmd.Flags().String(
		outputFlag,
		".",
		"directory to write files to",
	)
	cmd.Flags().StringSlice(envsFlag, []string{}, "environments to generate")

	logIfError(cmd.MarkFlagRequired(envsFlag))
	logIfError(viper.BindPFlags(cmd.Flags()))
	return cmd
}
