package main

import (
	"fmt"
	"math/rand"
	"time"
)

// o padrão multiplexer se propõe a juntar dois ou mais canais em um só
func main() {
	channel := multiplexer(write("Hello world!!"), write("Programming in go!!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(input_channel1, input_channel2 <-chan string) <-chan string {
	output_channel := make(chan string)

	go func() {
		for {
			// aqui no select, conforme os canais 1 e 2 de entrada receberem strings, eles serão tratados aqui
			select {
			case message := <-input_channel1:
				// jogando a mensagem para o canal de saida
				output_channel <- message
			case message := <-input_channel2:
				output_channel <- message
			}
		}
	}()

	return output_channel
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
