package main

import "fmt"

func func1() {
	fmt.Println("func 1")
}

func func2() {
	fmt.Println("func 2")
}

func main() {
	// o defer adia a execução de uma função
	// adia a execução da função até o ultimo momento possível
	// se a função atual tiver um return, a função com defer será chamada imediatamente antes do return
	defer func1()
	func2()
}
