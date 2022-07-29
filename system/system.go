package system

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func InitConfig(cf string, pths *NavPaths) {
	var err error
	if cf != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cf)
	} else {
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

	err = viper.Unmarshal(pths)
	cobra.CheckErr(err)
}
