package main

import (
	"dev05/internal/processor"
	"github.com/pborman/getopt"
	"os"
	"strings"
)

func main() {
	regPattern := getopt.String('p', "")
	after := getopt.IntLong("after", 'A', 0)
	before := getopt.IntLong("before", 'B', 0)
	grepContext := getopt.IntLong("context", 'C', 0)
	count := getopt.Bool('c')
	ignore := getopt.Bool('i')
	invert := getopt.Bool('v')
	n := getopt.Bool('n')

	getopt.Parse()
	var (
		path  string
		flags *processor.Flags
	)

	flags.After, flags.Before, flags.Context, flags.Count, flags.Ignore, flags.Invert, flags.Number = *after, *before,
		*grepContext, *count, *ignore, *invert, *n

	for _, arg := range os.Args {
		if !(strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--")) {
			path = arg
			break
		}
	}

	pr := processor.NewProcessor(path, *regPattern, flags)

	if err := pr.InputHandle(); err != nil {
		panic(err)
	}

	if err := pr.Process(); err != nil {
		panic(err)
	}
}
