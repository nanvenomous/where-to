/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"fmt"
	"os"
	"where-to/system"

	"github.com/spf13/cobra"
)

var (
	cfgFile    string
	completion string
	shells     []string
	paths      system.NavPaths
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "find-where-to-go",
	Short: "worker utilities to support the to shell function",
	Long:  `worker utilities to support the to shell function`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if completion != "" {
			switch completion {
			case shells[0]:
				cmd.Root().GenBashCompletion(os.Stdout)
			case shells[1]:
				cmd.Root().GenZshCompletion(os.Stdout)
			case shells[2]:
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case shells[3]:
				cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			default:
				fmt.Println("not a recognized shell")
				os.Exit(1)
			}
			os.Exit(0)
		} else {
			cmd.Help()
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(func() {
		system.InitConfig(cfgFile, &paths)
	})

	shells = []string{"bash", "zsh", "fish", "powershell"}
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	completionFlag := "completion"
	rootCmd.PersistentFlags().StringVar(&completion, completionFlag, "", "generate shell completion")
	rootCmd.RegisterFlagCompletionFunc(completionFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return shells, cobra.ShellCompDirectiveDefault
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/where-to.yaml)")
}
