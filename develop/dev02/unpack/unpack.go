package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var InvalidStringErr = errors.New("invalid string for parsing")

func StringUnpack(data string) (string, error) {
	runes := []rune(data)
	sb := strings.Builder{}

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			if i == 0 || unicode.IsDigit(runes[i-1]) {
				return "", InvalidStringErr
			}
			num, err := strconv.Atoi(string(runes[i : i+1]))
			if err != nil {
				return "", err
			}
			sb.WriteString(strings.Repeat(string(runes[i-1]), num-1))
		} else {
			sb.WriteString(string(runes[i]))
		}
	}

	return sb.String(), nil
}
