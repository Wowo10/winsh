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
	checkError(err)

	contentStr := string(content)
	contentStr = strings.ReplaceAll(contentStr, "\r\n", "\n")

	tempFile, err := ioutil.TempFile("", "winsh")
	checkError(err)

	err = tempFile.Chmod(0700)
	checkError(err)

	err = tempFile.Close()
	checkError(err)

	ioutil.WriteFile(tempFile.Name(), []byte(contentStr), 0)

	execPath, err := exec.LookPath(tempFile.Name())
	checkError(err)

	cmd := &exec.Cmd{
		Path:   execPath,
		Args:   append([]string{execPath}, argsWithoutProg[1:]...),
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	err = cmd.Run()
	checkError(err)

	err = os.Remove(tempFile.Name())
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
