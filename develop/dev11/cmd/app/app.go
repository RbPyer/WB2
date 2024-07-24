package main

import "dev11/internal/api"

func main() {
	s := api.NewServer()
	s.SetupHandlers()
	if err := s.Start(); err != nil {
		panic(err)
	}
}
