package system

import (
	"errors"
	"fmt"
	"os"
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
