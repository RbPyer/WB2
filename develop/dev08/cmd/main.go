package main

import (
	"dev08/internal/sh"
	"log"
)

func main() {
	shell := sh.NewShell()

	if err := shell.RunShell(); err != nil {
		log.Println(err)
	}
}
