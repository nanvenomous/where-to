/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"where-to/system"

	"github.com/spf13/cobra"
)

var (
	paths system.NavPaths
)

var rootCmd = &cobra.Command{
	Use:   "find-where-to-go",
	Short: "worker utilities to support the to shell function",
	Long:  `worker utilities to support the to shell function`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return system.CompletionsOrHelp(cmd)
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(func() {
		system.InitConfig(&paths)
	})

	system.CommonFlagsAndCompletions(rootCmd)
}
