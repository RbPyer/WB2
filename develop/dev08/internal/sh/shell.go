package sh

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell interface {
	RunShell() error
}

func NewShell() Shell {
	return &CurrentShell{
		driver: NewCmdDriver(),
	}
}

type CurrentShell struct {
	driver CmdDriver
}

func (ce *CurrentShell) RunShell() error {
	userInput := bufio.NewReader(os.Stdin)

	for {
		pwd, err := ce.driver.Pwd()
		if err != nil {
			return err
		}
		fmt.Printf("oswyndel~%s> ", pwd)

		input, err := userInput.ReadString('\n')
		if err != nil {
			return err
		}

		args := strings.Fields(input)
		output, err := ce.driver.Run(args[0], args[1:])

		if err != nil {
			return err
		}

		fmt.Println(output)
	}
}
