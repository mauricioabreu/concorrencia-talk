package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go say_something("channels são usados para comunicação", ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}

func say_something(something string, ch chan string) {
	defer close(ch)

	for i := 0; i < 5; i++ {
		ch <- something
		time.Sleep(time.Second / 2)
	}
}
