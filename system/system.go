package system

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	shells     = []string{"bash", "zsh", "fish", "powershell"}
	completion string
	cfgFile    string
)

func VerifyPath(pth string) error {
	var err error
	flInfo, err := os.Stat(pth)
	if err != nil {
		return err
	}
	if !flInfo.IsDir() {
		return errors.New(fmt.Sprintf("%s is not a directory", pth))
	}
	return nil
}

type NavPaths map[string]string

func InitConfig(pths *NavPaths) {
	var err error
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(path.Join(home, ".config"))
		viper.SetConfigName("where-to")
	}

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	err = viper.Unmarshal(pths)
	cobra.CheckErr(err)
}

func CompletionsOrHelp(cmd *cobra.Command) error {
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
}

func CommonFlagsAndCompletions(cmd *cobra.Command) {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	cmd.CompletionOptions.DisableDefaultCmd = true

	completionFlag := "completion"
	cmd.PersistentFlags().StringVar(&completion, completionFlag, "", "generate shell completion")
	cmd.RegisterFlagCompletionFunc(completionFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return shells, cobra.ShellCompDirectiveDefault
	})

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/where-to.yaml)")
}
