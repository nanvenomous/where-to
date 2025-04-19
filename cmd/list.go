/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show all the current aliases and there corresponding paths",
	Long:  `show all the current aliases and there corresponding paths`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var (
			err      error
			homePath string
		)
		homePath, err = os.UserHomeDir()
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"alias", "path"})

		for al, pth := range paths {
			t.AppendRow(table.Row{al, strings.Replace(pth, homePath, "~", -1)})
		}
		t.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
