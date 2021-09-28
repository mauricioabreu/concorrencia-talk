package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// um  waitgroup funciona como um contador
	// você diz ao runtime do go quantas goroutines você espera concluir
	// e espera por elas
	var wg sync.WaitGroup
	wg.Add(1) // troque para 2 e veja um problema ao esperar por uma goroutine que não existe

	go func() {
		say_something("waitgroup é uma espécie de contador")
		wg.Done()
	}()

	wg.Wait()
}

func say_something(something string) {
	for i := 0; i < 5; i++ {
		fmt.Println(something)
		time.Sleep(time.Second / 2)
	}
}
