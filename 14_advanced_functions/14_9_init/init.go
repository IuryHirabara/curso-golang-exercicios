package main

import "fmt"

// a função init é executada antes da função main
// a diferença dela para a main é que ela alem de não representar a função principal "main", ela pode ser usada por arquivo e não por pacote

// variveis no escopo global não podem ser iniciadas com o short sintax ":="
var n int

func init() {
	n = 10
	fmt.Println("Executing init function")
}

func main() {
	fmt.Println(n)
	fmt.Println("Executing main function")
}
