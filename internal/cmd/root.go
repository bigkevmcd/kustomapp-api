package cmd

import (
	"log"
	"os"

	"github.com/bigkevmcd/kustomapp-api/pkg/tree"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	pathFlag = "path"
)

func init() {
	cobra.OnInitialize(initConfig)
}

func logIfError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}

func makeRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kapp",
		Short: "application manager",
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
	logIfError(viper.BindPFlags(cmd.Flags()))
	return cmd
}

// Execute is the main entry point into this component.
func Execute() {
	if err := makeRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

func makeTable(header ...string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	return table
}
