package main

import "fmt"

func main() {

	data := make([]int, 5)

	iterate := func(handleData chan<- int) {
		defer close(handleData)
		for d := range data {
			handleData <- d
		}
	}

	handleData := make(chan int)
	go iterate(handleData)
	for num := range handleData {
		fmt.Println(num)
	}
}
