package main

import (
	"dev02/unpack"
	"fmt"
	"log"
)

func main() {
	var str string
	if _, err := fmt.Scanf("%s", &str); err != nil {
		log.Fatalf("failed to read input: %v", err)
	}
	result, err := unpack.StringUnpack(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
