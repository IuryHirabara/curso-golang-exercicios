package main

import "fmt"

func closure() func() {
	text := "Inside of closure function"

	ffunction := func() {
		fmt.Println(text)
	}

	return ffunction
}

func main() {
	text := "Inside of main function"
	fmt.Println(text)

	newFunc := closure()
	newFunc()
}
