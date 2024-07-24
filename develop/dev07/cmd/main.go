package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
			log.Printf("Channel with durations %v was closed\n", after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		//sig(2*time.Hour),
		//sig(5*time.Minute),
		sig(1*time.Second),
		//sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}

func or(channels ...<-chan interface{}) <-chan interface{} {
	utilCh := make(chan interface{})

	wg := new(sync.WaitGroup)
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			defer wg.Done()
			for v := range ch {
				utilCh <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(utilCh)
	}()

	return utilCh
}
