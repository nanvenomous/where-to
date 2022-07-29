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
		zsh_init := `
function t() {
  to_go_to="${1}";
  find-where-to-go isdir $to_go_to
  if [ $? -eq 0 ]; then
    clear
    if command -v exa &> /dev/null; then 
      exa --tree --level=1 --group-directories-first ${to_go_to}
    elif command -v tree &> /dev/null; then 
      tree -C -L 1 --dirsfirst ${to_go_to}
    else
      ls --color --group-directories-first -1 ${to_go_to}
    fi
  fi
}

function dn() {
  to_go_to="${1}";
  find-where-to-go isdir $to_go_to

  if [ $? -eq 0 ]; then
    if [ $# -eq 0 ]; then 
      to_go_to=${HOME};
    fi;
    builtin cd "${to_go_to}"; t
  fi
}
alias up="cd ..; t"

function to() {
  headed=$(find-where-to-go convert $1)
  if ! [ -z $headed ]; then
    builtin cd "${headed}"
    clear; exa --tree --level=1 --group-directories-first
  fi
}

`
		fmt.Println(zsh_init)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}
