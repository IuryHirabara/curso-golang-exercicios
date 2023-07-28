package main

import "fmt"

func main() {
	// criando channel com buffer
	// sendo um buffer com capacidade 2, podemos enviar até dois parâmetros para o channel sem que o fluxo seja interrompido
	// também podemos interpretar que está criando duas posições no channel
	channel := make(chan string, 2)

	channel <- "Hello world!"
	channel <- "Hello my friend!!"
	// com mais um, causa deadlock
	// channel <- "How are you??"

	// recebendo as mensagens
	message := <-channel
	message2 := <-channel
	// nesse caso, também irá dar deadlock pois está esperando um valor que não será enviado
	// message3 := <-channel
	fmt.Println(message)
	fmt.Println(message2)
	// fmt.Println(message3)

	// sem o buffer, o fluxo é interrompido de imediato
}
