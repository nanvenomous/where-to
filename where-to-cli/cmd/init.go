/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes the 'to' command in the shell with completions",
	Long:  `initializes the 'to' command in the shell with completions`,
	Run: func(cmd *cobra.Command, args []string) {
		zsh_init := `function go_to() {
  headed=$(find-where-to-go $1)
  if ! [ -z $headed ]; then
    builtin cd "${headed}"
    clear; exa --tree --level=1 --group-directories-first
  fi
}
alias to=go_to


setopt completealiases`
		fmt.Println(zsh_init)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}
