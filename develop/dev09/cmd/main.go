package main

import (
	"dev09/internal/wget"
	"fmt"
	"log"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("Please provide a url")
	case 2:
		wgetImpl := wget.NewWget()
		err := wgetImpl.Wget(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Undefined behavior, please, contact me: @hero_of")
	}
}
