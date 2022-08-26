/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err  error
			fl   *os.File
			byts []byte
			pth  string
			ok   bool
		)
		if pth, ok = paths[args[0]]; ok {
			delete(paths, args[0])
			fl, err = os.Create(viper.ConfigFileUsed())
			defer fl.Close()
			if err != nil {
				return err
			}
			byts, err = yaml.Marshal(paths)
			if err != nil {
				return err
			}
			_, err = fl.Write(byts)
			if err != nil {
				return err
			}
			fmt.Println("Successfully removed element from config file:")
			fmt.Printf("\t%s: %s\n", args[0], pth)
		} else {
			return errors.New(fmt.Sprintf("the given alias, %s, was never in the config file: %s", args[0], viper.ConfigFileUsed()))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
