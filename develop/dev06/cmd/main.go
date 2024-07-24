package main

import (
	"dev06/internal/cutter"
	"github.com/pborman/getopt"
)

func main() {
	f := getopt.StringLong("fields", 'f', "")
	d := getopt.StringLong("delimiter", 'd', "\t")
	s := getopt.BoolLong("separated", 's')
	getopt.Parse()

	cc := cutter.NewCutter(*f, *d, *s)
	if err := cc.Process(); err != nil {
		panic(err)
	}
}
