package main

import (
	"fmt"
	"time"
)

func main() {
	// O que é executado antes? Pense sobre alocação de recursos ao se executar uma goroutine
	go say_something("Concorrência")
	say_something("em go")
}

func say_something(something string) {
	for i := 0; i < 5; i++ {
		fmt.Println(something)
		time.Sleep(200 * time.Millisecond)
	}
}
