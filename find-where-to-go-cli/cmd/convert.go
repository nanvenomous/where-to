/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"errors"
	"fmt"
	"where-to/system"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "<alias> converts the argument alias to the path specified in the config",
	Long:  `<alias> converts the argument alias to the path specified in the config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		if len(args) < 1 {
			return errors.New("convert required one argument <alias>")
		}
		if pth, ok := paths[args[0]]; ok {
			err = system.VerifyPath(pth)
			if err != nil {
				return err
			}
			fmt.Println(pth)
		} else {
			return errors.New(fmt.Sprintf("the given alias, %s, is not in the config file: %s", args[0], viper.ConfigFileUsed()))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
