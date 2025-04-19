/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"where-to/system"

	"github.com/nanvenomous/exfs"
	"github.com/spf13/cobra"
)

var (
	FS    *exfs.FileSystem
	paths system.NavPaths
)

var rootCmd = &cobra.Command{
	Use:   "where-to",
	Short: "your personal driver around your os",
	Long:  `your personal driver around your os`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		FS = exfs.NewFileSystem()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return system.CompletionsOrHelp(cmd)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(func() {
		system.InitConfig(&paths)
	})

	system.CommonFlagsAndCompletions(rootCmd)
}
