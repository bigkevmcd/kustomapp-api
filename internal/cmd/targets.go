package cmd

import (
	"github.com/bigkevmcd/kustomapp-api/pkg/tree"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func makeTargetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "targets",
		Short: "list target environments",
		Run: func(cmd *cobra.Command, args []string) {
			fs := osfs.New(viper.GetString(pathFlag))
			targets, err := tree.Targets(fs)
			logIfError(err)

			table := makeTable("Environment", "Cluster")
			for _, v := range targets {
				table.Append([]string{v.Environment, v.Cluster})
			}
			table.Render()
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
