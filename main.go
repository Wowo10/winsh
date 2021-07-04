package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	content, err := ioutil.ReadFile(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}

	contentStr := string(content)
	contentStr = strings.ReplaceAll(contentStr, "\r\n", "\n")

	tempFile, err := ioutil.TempFile("", "winsh")
	if err != nil {
		panic(err)
	}

	err = tempFile.Chmod(0700)
	if err != nil {
		panic(err)
	}

	err = tempFile.Close()
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(tempFile.Name(), []byte(contentStr), 0)

	execPath, err := exec.LookPath(tempFile.Name())
	if err != nil {
		panic(err)
	}

	cmd := &exec.Cmd{
		Path:   execPath,
		Args:   append([]string{execPath}, argsWithoutProg[1:]...),
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	if err := os.Remove(tempFile.Name()); err != nil {
		panic(err)
	}
}
