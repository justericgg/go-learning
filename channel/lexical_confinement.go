package main

import "fmt"

func main() {

	channelOwner := func() <-chan int {
		dataHandler := make(chan int)

		go func() {
			defer close(dataHandler)
			data := make([]int, 5)
			for d := range data {
				dataHandler <- d
			}
		}()

		return dataHandler
	}

	myChan := channelOwner()
	for num := range myChan {
		fmt.Println(num)
	}
}
