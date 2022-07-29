/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"fmt"
	"where-to/system"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type NavPaths map[string]string

func newCommand(nm string, pth string) *cobra.Command {
	return &cobra.Command{
		Use:   nm,
		Short: pth,
		Long:  pth,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			err = system.VerifyPath(pth)
			if err != nil {
				return err
			}
			fmt.Println(pth)
			return nil
		},
	}
}

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "<alias> converts the argument alias to the path specified in the config",
	Long:  `<alias> converts the argument alias to the path specified in the config`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	initConfig()
	var err error
	var pths NavPaths
	err = viper.Unmarshal(&pths)
	cobra.CheckErr(err)

	for k, v := range pths {
		convertCmd.AddCommand(newCommand(k, v))
	}

	rootCmd.AddCommand(convertCmd)

}
