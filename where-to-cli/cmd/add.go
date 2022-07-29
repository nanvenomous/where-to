/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "<alias> <path> add alias to path so you can quickly jump there",
	Long:  `<alias> <path> add alias to path so you can quickly jump there`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("expected two arguments <alias> <path>")

		}
		viper.WriteConfig()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
