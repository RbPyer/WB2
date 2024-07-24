package main

import (
	"dev03/internal/parser"
	"dev03/internal/processor"
	"fmt"
	"github.com/pborman/getopt"
	"log"
	"os"
	"strings"
)

func main() {
	p := parser.NewParser("")

	flagK := getopt.Int('k', 0)
	flagN := getopt.Bool('n')
	flagR := getopt.Bool('r')
	flagU := getopt.Bool('u')
	getopt.Parse()

	p.Flags.K, p.Flags.N, p.Flags.R, p.Flags.U = *flagK, *flagN, *flagR, *flagU

	if len(os.Args) == 0 {
		fmt.Println("Please give me a command.")
		os.Exit(1)
	}

	for _, arg := range os.Args {
		if !strings.HasPrefix(arg, "-") {
			p.Path = arg
		}
	}

	log.Printf("New information received:\nPath: %s; Flags: k - %d; n - %v, u - %v, r - %v",
		p.Path, p.Flags.K, p.Flags.N, p.Flags.U, p.Flags.R)

	prc := processor.NewProcessor()
	if err := p.Read(); err != nil {
		panic(err)
	}
	prc.Sort(p)
	if err := p.Write(); err != nil {
		panic(err)
	}
}
