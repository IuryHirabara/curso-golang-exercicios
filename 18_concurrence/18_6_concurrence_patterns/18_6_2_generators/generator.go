package main

import (
	"fmt"
	"time"
)

func main() {
	// dessa forma podemos chamar como se fosse uma função normal
	channel := write("Hello world")

	// com isso, abre a possibilidade de fazer um laço de repetição
	for text := range channel {
		fmt.Println(text)
	}
}

// nesse padrão, estamos encapsulando a go routine em uma função, que retornará essa go routine
// usamos esse padrão para esconder a complexidade da go routine
func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
