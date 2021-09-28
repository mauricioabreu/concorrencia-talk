package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	after := time.After(time.Second * 3)

	go func() {
		for {
			time.Sleep(time.Second / 2) // o que aconteceria se colocássemos o sleep embaixo do receive?
			ch1 <- "o que usamos para comunicar canais?"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			ch2 <- "é claro que é select!"
		}
	}()

	for {
		select {
		case question := <-ch1:
			fmt.Println(question)
		case answer := <-ch2:
			fmt.Println(answer)
		case <-after:
			fmt.Println("Done!")
			os.Exit(0)
		}
	}
}
