package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// criando variável wait group
	var waitGroup sync.WaitGroup

	// adionando duas go routines que o wait group precisa esperar terminar
	waitGroup.Add(4)

	// função anônima
	go func() {
		write("Hello world!")

		// chamando váriavel wait group para dizer que a função terminou
		// decrementa do wait group
		waitGroup.Done()
	}()

	go func() {
		write("Programming")
		waitGroup.Done()
	}()

	go func() {
		write("in")
		waitGroup.Done()
	}()

	go func() {
		write("Go!!")
		waitGroup.Done()
	}()

	// para dizer à função main para esperar a contagem das go routines do wait group
	waitGroup.Wait()

	// dessa forma, as duas go routines serão executadas de forma concorrente, mas mesmo assim, a execução do programa não será encerrada até que as duas funções anônimas sejam completamente executadas
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
