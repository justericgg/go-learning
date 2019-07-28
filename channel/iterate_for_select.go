package main

import (
	"fmt"
	"time"
)

func main() {
	members := []string{"Eric", "Joanne", "James", "River"}

	done := make(chan int)
	workChan := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()

	go func() {
		for w := range workChan {
			fmt.Println(w)
		}
	}()

	for _, m := range members {
		select {
		case <-done:
			fmt.Println("Done")
			return
		case workChan <- m:
			time.Sleep(1 * time.Second)
		}
	}
}
