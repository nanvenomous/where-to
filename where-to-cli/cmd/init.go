/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"html/template"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type TemplateValues struct {
	T                   string
	Dn                  string
	Up                  string
	To                  string
	VerticalListCommand string
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func getListCommand() string {
	if commandExists("exa") {
		return "exa --tree --level=1 --group-directories-first"
	} else if commandExists("tree") {
		return "tree -C -L 1 --dirsfirst"
	}
	// TODO : handle osx
	return "ls --color --group-directories-first -1"
}

const UNIX_INIT = `
function {{.T}}() {
  find-where-to-go isdir ${@}
  if [ $? -eq 0 ]; then
    clear; {{.VerticalListCommand}} ${@}
  fi
}

function {{.Dn}}() {
  to_go_to="${1}";
  find-where-to-go isdir $to_go_to

  if [ $? -eq 0 ]; then
    if [ $# -eq 0 ]; then 
      to_go_to=${HOME};
    fi;
    builtin cd "${to_go_to}"; clear; {{.VerticalListCommand}}
  fi
}
alias {{.Up}}="cd ..; clear; {{.VerticalListCommand}}"

function {{.To}}() {
  headed=$(find-where-to-go convert $1)
  if ! [ -z $headed ]; then
    builtin cd "${headed}"
    clear; {{.VerticalListCommand}}
  fi
}`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes the shell builtins",
	Long:  `initializes the shell builtins t, to, dn, up`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err        error
			unxInitTmp *template.Template
		)
		unxInitTmp = template.New("Unix_Init_Template")
		unxInitTmp, err = unxInitTmp.Parse(UNIX_INIT)
		if err != nil {
			return err
		}
		return unxInitTmp.Execute(os.Stdout, TemplateValues{
			T:                   "t",
			Dn:                  "dn",
			Up:                  "up",
			To:                  "to",
			VerticalListCommand: getListCommand(),
		})
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
