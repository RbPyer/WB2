package parser

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type Parser interface {
	Read() error
	Write() error
}

type CurrentParser struct {
	Path  string
	Flags *Flags
	Data  []string
}

func (p *CurrentParser) Read() error {
	file, err := os.Open(p.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("open file %s error", p.Path))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.Data = append(p.Data, scanner.Text())
	}
	return nil
}

func (p *CurrentParser) Write() error {
	file, err := os.Create("sorted_data")
	if err != nil {
		return errors.New(fmt.Sprintf("open file %s error", p.Path))
	}
	defer file.Close()

	for _, line := range p.Data {
		_, writeErr := file.WriteString(line + "\n")
		if writeErr != nil {
			return writeErr
		}
	}

	log.Println("A new file with sorted data was created.")

	return nil
}

type Flags struct {
	U bool
	K int
	N bool
	R bool
}

func NewParser(path string) *CurrentParser {
	return &CurrentParser{Path: path, Flags: &Flags{}}
}
