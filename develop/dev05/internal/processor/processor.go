package processor

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Processor interface {
	InputHandle() error
	Process() error
}

func NewProcessor(path string, pattern string, flags *Flags) Processor {
	return &CurrentProcessor{
		Path:    path,
		Pattern: pattern,
		Flags:   flags,
	}
}

type CurrentProcessor struct {
	Pattern string
	Path    string
	Flags   *Flags
	Data    []string
}

func (c *CurrentProcessor) getExp(pattern string, ignore bool) (*regexp.Regexp, error) {
	var ignorePrefix = ""
	if ignore {
		ignorePrefix = "(?i)"
	}
	compExp, err := regexp.Compile(ignorePrefix + pattern)
	if err != nil {
		return nil, err
	}
	return compExp, nil
}

func (c *CurrentProcessor) GetInterSecNum(data []string, exp *regexp.Regexp) int {
	var result int
	for _, str := range data {
		if match := exp.Match([]byte(str)); match {
			result++
		}
	}
	return result
}

func (c *CurrentProcessor) InputHandle() error {
	file, err := os.Open(c.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		c.Data = append(c.Data, sc.Text())
	}
	return nil
}

func (c *CurrentProcessor) Process() error {
	exp, err := c.getExp(c.Pattern, c.Flags.Ignore)
	if err != nil {
		return err
	}

	if c.Flags.Count {
		result := c.GetInterSecNum(c.Data, exp)
		if c.Flags.Invert {
			result = len(c.Data) - result
		}
		fmt.Println(result)
		return nil
	}

	if c.Flags.After == 0 && c.Flags.Before == 0 && c.Flags.Context != 0 {
		tmp := c.Flags.Context / 2
		c.Flags.After, c.Flags.Before = tmp, tmp
		c.reg(exp)
	}
	return nil
}

func (c *CurrentProcessor) reg(expression *regexp.Regexp) {
	for i, str := range c.Data {
		match := expression.Match([]byte(str))
		if c.Flags.Invert && !match {
			c.print(i)
		} else if !c.Flags.Invert && match {
			c.print(i)
		}
	}
}

func (c *CurrentProcessor) print(i int) {
	var start, end = 0, len(c.Data)
	if i-c.Flags.After > 0 {
		start = i - c.Flags.After
	}
	if i+c.Flags.Before < len(c.Data) {
		end = i + c.Flags.Before
	}
	if end != len(c.Data) {
		end++
	}

	for num := start; num < end; num++ {
		if c.Flags.Number {
			fmt.Printf("%d: ", num+1)
		}
		fmt.Printf("%s\n", c.Data[num])
	}
}

type Flags struct {
	After   int
	Before  int
	Context int
	Count   bool
	Ignore  bool
	Invert  bool
	Number  bool
}
