package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	argsWithoutProg := os.Args[1:]

	execPath, err := exec.LookPath(argsWithoutProg[0])

	if err != nil {
		fmt.Println("Error: ", err)
	}

	cmd := &exec.Cmd{
		Path:   execPath,
		Args:   argsWithoutProg,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
