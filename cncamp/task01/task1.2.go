package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)
	defer close(messages)

	go func() {
		consumerTicker := time.NewTicker(1 * time.Second)
		for _ = range consumerTicker.C {
			select {
			case <-done:
				fmt.Println("customer interrupt...")
				return
			default:
				fmt.Printf("get message: %d\n", <-messages)
			}

		}
	}()

	go func() {
		i := 0
		producerTicker := time.NewTicker(1 * time.Second)
		for _ = range producerTicker.C {
			messages <- i
			fmt.Printf("enter message: %d\n", i)
			i++
		}
	}()

	time.Sleep(30 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main exit!")
}
