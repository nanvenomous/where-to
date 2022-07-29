/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"where-to/system"

	"github.com/spf13/cobra"
)

// isdirCmd represents the isdir command
var isdirCmd = &cobra.Command{
	Use:   "isdir",
	Short: "<path> check if a path is a valid directory",
	Long:  `<path> check if a path is a valid directory`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if len(args) > 0 {
			err = system.VerifyPath(args[0])
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(isdirCmd)

}
