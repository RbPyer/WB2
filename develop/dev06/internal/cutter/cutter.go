package cutter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Cutter interface {
	Process() error
}

type CurrentCutter struct {
	fields    string
	delimiter string
	sep       bool
}

func (cc *CurrentCutter) Process() error {
	reader := bufio.NewReader(os.Stdin)
	var (
		str string
		err error
	)

	for err != io.EOF {
		str, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		result, cutErr := cc.cut(strings.TrimSpace(str))
		if cutErr != nil {
			return cutErr
		}
		fmt.Println(result)
	}
	return nil
}

func (cc *CurrentCutter) cut(str string) (string, error) {
	var response string

	cols := strings.Split(str, cc.delimiter)
	if len(cols) == 1 && !cc.sep {
		return str, nil
	}
	if len(cols) > 1 && cc.sep {
		digitFields, err := parseFields(cc.fields)
		if err != nil {
			return "", err
		}
		for _, num := range digitFields {
			response += fmt.Sprintf("%s ", cols[num-1])
		}
	}
	return response, nil
}

func parseFields(fields string) ([]int, error) {
	fields = strings.TrimSpace(fields)
	intStrList := strings.Split(fields, ",")
	res := make([]int, len(intStrList))
	for ind, intStr := range intStrList {
		num, err := strconv.Atoi(intStr)
		if err != nil {
			return nil, errors.New("bad field list")
		}
		res[ind] = num
	}
	return res, nil
}

func NewCutter(fields, delimiter string, sep bool) Cutter {
	return &CurrentCutter{
		fields:    fields,
		delimiter: delimiter,
		sep:       sep,
	}
}
