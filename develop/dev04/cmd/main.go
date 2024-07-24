package main

import (
	"dev04/internal/processor"
	"fmt"
)

func main() {
	testData := []string{"тяпка", "пятак", "пятка", "листок", "слиток", "столик"}
	p := processor.NewProcessor()
	p.SetInfo(testData)

	for k, v := range p.GetInfo() {
		fmt.Printf("%s = %s\n", k, v)
	}
}
