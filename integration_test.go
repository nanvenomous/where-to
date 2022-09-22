package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func cap(command string, args []string) (string, error) {
	var outb, errb bytes.Buffer
	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		return errb.String(), err
	}
	return "", nil
}

func executeWithWhereTo(shell string, toEx string) (string, string, error) {
	var outb, errb bytes.Buffer
	cmd := exec.Command(
		shell,
		"-c",
		"eval \"$(where-to init)\""+"; "+toEx,
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		return "", errb.String(), err
	}
	return outb.String(), "", nil
}

func integrationTest(t *testing.T, shell string) error {
	var (
		err            error
		outStr, errStr string
	)
	err = os.MkdirAll("intg_tst/ex/mp/dr/", os.ModePerm)
	if err != nil {
		return err
	}

	_, err = os.OpenFile("intg_tst/ex/mp/dr/expected.txt", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	errStr, err = cap("./where-to", []string{"add", "als", "intg_tst/ex/mp/dr/"})
	if err != nil {
		fmt.Println(errStr)
		return err
	}

	outStr, errStr, err = executeWithWhereTo(shell, "to als")
	if err != nil {
		fmt.Println(errStr)
		return err
	}

	assert.True(t, strings.Contains(outStr, "expected.txt"))

	err = os.RemoveAll("./intg_tst")
	return nil
}

func TestIntegrationBash(t *testing.T) {
	err := integrationTest(t, "bash")
	if err != nil {
		t.Error(err)
	}
}

func TestIntegrationZsh(t *testing.T) {
	err := integrationTest(t, "zsh")
	if err != nil {
		t.Error(err)
	}
}
