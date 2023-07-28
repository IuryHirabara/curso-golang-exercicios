package main

import (
	"fmt"
	"time"
)

func main() {
	// channels são para envio e recebimento de dados

	// criando channel
	channel := make(chan string)

	go write("Hello world!!", channel)

	// laço infinito
	// for {
	// erro deadlock
	// esses erros só são identificados em execução e não são pegos pelo compilador
	// vai esperar infinitos valor pois está em um laço de repetição infinito, porém, a função irá retornar 5 vezes por causa do laçõ de repetição dela
	// ou seja, o deadlock é basicamente a espera de retorno de uma função que não irá retornar nada, com isso, o channel permanecerá aberto.

	// a variável isOpen irá servir para verificar se o canal está aberto
	// 	message, isOpen := <-channel
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	// refatorando o código acima
	// dessa forma não precisamos fazer a verificação se o channel está aberto
	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("End of program!!!")

	// esperando valor do canal
	// message := <-channel
	// fmt.Println(message)
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		// mandando valor para o channel
		channel <- text
		time.Sleep(time.Second)
	}

	// fechando canal
	close(channel)
}
