package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Channel 2"
		}
	}()

	for {
		// Isso gerará um delay desnecessário, pois em meio segundo a mensagem ficará pronta
		// message1 := <-channel1
		// fmt.Println(message1)

		// Enquanto isso, terá que esperar os dois segundos da segunda go routine pois está no laço de repetição. Então a próxima iteração só será executada quando o segundo Println ser executada

		// Ocasionará perda de tempo pois a primeira mensagem ficará pronta de meio em meio segundo, enquanto a segunda chamada só ficará pronta de dois em dois segundos
		// message2 := <-channel2
		// fmt.Println(message2)

		// para contornar isso podemos usar a estrutura select que serve exclusivamente para tratar go routines
		// declarando select
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		}
	}
}
