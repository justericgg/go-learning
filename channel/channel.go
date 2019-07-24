package main

import "fmt"

func main() {
	fmt.Println("Start...")

	// 這裡定義一個 function 回傳一個唯讀的 channel
	// 如此外界只能對這個 channel 取讀取的動作，我們也就
	// 確保只有在這個 function 內可以做寫以及關閉 channel 的動作
	buildChan := func() <-chan int {
		myChan := make(chan int)

		//將要對 channel 做寫動作的地方用一個獨立的 goroutine 來做操作
		go func() {
			defer close(myChan)
			for i := 0; i <= 5; i++ {
				myChan <- i
			}
		}()

		return myChan
	}

	myChan := buildChan()
	for result := range myChan {
		fmt.Println(result)
	}
	fmt.Println("Done")
}
