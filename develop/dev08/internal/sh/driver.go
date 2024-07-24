package sh

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type CmdDriver interface {
	Run(cmd string, args []string) (string, error)
	Cd(args []string) error
	Pwd() (string, error)
	Echo(args []string) string
	Kill(args []string) (string, error)
	Ps() (string, error)
	Exec(command string, args []string) (string, error)
	Quit()
}

func NewCmdDriver() CmdDriver {
	return &CurrentDriver{}
}

type CurrentDriver struct{}

func (driver *CurrentDriver) Quit() {
	log.Println("Have a good day, sir! Bye! :)")
	os.Exit(0)
}

func (driver *CurrentDriver) Run(cmd string, args []string) (string, error) {
	switch cmd {
	case "cd":
		if err := driver.Cd(args); err != nil {
			return "", err
		}
		return "", nil
	case "pwd":
		pwd, err := driver.Pwd()
		if err != nil {
			return "", err
		}
		return pwd, nil
	case "kill":
		result, err := driver.Kill(args)
		if err != nil {
			return "", err
		}
		return result, nil
	case "echo":
		return driver.Echo(args), nil
	case "exec":
		result, err := driver.Exec(args[0], args[1:])
		if err != nil {
			return "", err
		}
		return result, nil
	case "ps":
		result, err := driver.Ps()
		if err != nil {
			return "", err
		}
		return result, nil
	case "q", "exit":
		driver.Quit()
	default:
		return "", errors.New(fmt.Sprintf("Unknown command: %s", cmd))
	}

	return "", errors.New(fmt.Sprintf("Undefined behavior, pls contact me: @hero_of"))
}

func (driver *CurrentDriver) Ps() (string, error) {
	cmd := exec.Command("ps")

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("ps: %w", err)
	}

	return strings.TrimSpace(string(out)), nil
}

func (driver *CurrentDriver) Kill(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("kill: no process")
	}

	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("kill: %w", err)
	}

	if err = syscall.Kill(pid, syscall.SIGKILL); err != nil {
		return "", fmt.Errorf("kill: %w", err)
	}

	return fmt.Sprintf("kill: %d", pid), nil
}

func (driver *CurrentDriver) Exec(command string, args []string) (string, error) {
	cmd := exec.Command(command, args...)

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%s: %w", command, err)
	}

	return string(out), nil
}

func (driver *CurrentDriver) Echo(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return args[0]
}

func (driver *CurrentDriver) Pwd() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func (driver *CurrentDriver) Cd(args []string) error {
	if len(args) == 0 {
		if err := os.Chdir("/home"); err != nil {
			return err
		}
		return nil
	}

	if err := os.Chdir(args[0]); err != nil {
		return err
	}

	return nil
}
