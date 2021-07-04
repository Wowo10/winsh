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

	const tempFilePath = "./temp/temp.sh"

	ioutil.WriteFile(tempFilePath, []byte(contentStr), 0)

	execPath, err := exec.LookPath(tempFilePath)

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

}
