/*
Copyright © 2022 nanvenomous mrgarelli@gmail.com

*/
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"where-to/system"

	"github.com/spf13/viper"
)

var (
	paths system.NavPaths
)

func convert(args []string) error {
	var err error

	if len(args) < 1 {
		return errors.New("convert required one argument <alias>")
	}
	if pth, ok := paths[args[0]]; ok {
		err = system.VerifyPath(pth)
		if err != nil {
			return err
		}
		fmt.Println(pth)
	} else {
		return errors.New(fmt.Sprintf("the given alias, %s, is not in the config file: %s", args[0], viper.ConfigFileUsed()))
	}
	return nil
}

func isdir(args []string) error {
	if len(args) > 0 {
		toChk := args[len(args)-1]
		if !strings.HasPrefix(toChk, "-") {
			return system.VerifyPath(toChk)
		}
	}
	return nil
}

func findWhereToGo(args []string) error {
	if len(args) < 2 {
		return errors.New("required at least one argument.")
	}
	switch args[1] {
	case "isdir":
		return isdir(args[2:])
	case "convert":
		system.InitConfig(&paths)
		return convert(args[2:])
	default:
		return errors.New(fmt.Sprintf("Not a recognized command %s", args[1]))
	}
}

func main() {
	var err error
	err = findWhereToGo(os.Args)
	if err != nil {
		panic(err)
	}
}
