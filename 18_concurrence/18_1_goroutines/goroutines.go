package main

import (
	"fmt"
	"time"
)

func main() {
	// paralelismo
	// quando há duas ou mais tarefas sendo executadas ao mesmo tempo

	// concorrência
	// as tarefas podem estar sendo executadas ao mesmo tempo em diferentes núcleos ou em um só
	// onde uma tarefa ficaria revezando com a outro a sua execução, nesse caso, não estão sendo executadas ao mesmo tempo

	// invocando goroutine com o parâmetro "go"
	// com esse parâmetro, o programa não vai esperar a primeira chamada da função write terminar para só então executar a segunda chamada
	go write("Hello World!")

	// se colocarmos o parâmetro "go" nessa chamada da função write, o programa não esperará o término da execução dessa função e já irá encerrar a execução do programa
	write("Programming in GO!")
}

func write(text string) {
	// laço de repetição infinito
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
