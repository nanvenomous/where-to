/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type NavPaths map[string]string

var (
	cfgFile    string
	completion string
	shells     []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "find-where-to-go",
	Short: "helper to turn aliases into actual paths on the system",
	Long:  `helper to turn aliases into actual paths on the system`,
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
		}

		return nil
	},
}

func newCommand(nm string, pth string) *cobra.Command {
	return &cobra.Command{
		Use:   nm,
		Short: pth,
		Long:  pth,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			flInfo, err := os.Stat(pth)
			if err != nil {
				return err
			}
			if !flInfo.IsDir() {
				return errors.New(fmt.Sprintf("%s is not a directory", pth))
			}
			fmt.Println(pth)
			return nil
		},
	}
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	// cobra.OnInitialize(initConfig)
	initConfig()

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

	var err error
	var pths NavPaths
	err = viper.Unmarshal(&pths)
	cobra.CheckErr(err)
	for k, v := range pths {
		rootCmd.AddCommand(newCommand(k, v))
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		var err error
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name "adi" (without extension).
		viper.AddConfigPath(path.Join(home, ".config"))
		viper.SetConfigName("where-to")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	_ = viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
