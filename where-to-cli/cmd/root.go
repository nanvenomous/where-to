/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"os"
	"where-to/system"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	paths system.NavPaths
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "where-to",
	Short: "your personal driver around your os",
	Long:  `your personal driver around your os`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return system.CompletionsOrHelp(cmd)
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		err = viper.WriteConfig()
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		system.InitConfig(&paths)
	})

	system.CommonFlagsAndCompletions(rootCmd)
}
