package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	done := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("Done")
				return
			default:
			}
			fmt.Println("Hi Eric")
			time.Sleep(1 * time.Second)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		reader.ReadString('\n')
		close(done)
		break
	}

	wg.Wait()
}
