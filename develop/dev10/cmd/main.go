package main

import (
	"context"
	tc "dev10/internal/telnetClient"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	const MustDie = 1
	timeout := flag.String("timeout", "10", "duration of time to wait before exit")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatalln("Usage: telnet <command> <args>")
	}
	ctx, cancel := context.WithTimeout(context.Background(), MustDie*time.Minute)
	timeoutDur, err := time.ParseDuration(fmt.Sprintf("%ss", *timeout))
	if err != nil {
		log.Fatalf("Invalid timeout duration: %s", err.Error())
	}

	host, port := args[0], args[1]
	client := tc.NewClient()
	err = client.Connect(host, port, timeoutDur, ctx, cancel)
	if err != nil {
		log.Fatalf("Connect failed: %s", err.Error())
	}

	log.Println("End of connection...")
}
