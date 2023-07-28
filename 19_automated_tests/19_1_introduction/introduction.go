package main

import (
	"fmt"
	"tests_introduction/adresses"
)

func main() {
	typeOfAdress := adresses.IsValidType("Avenue Paulista")
	fmt.Println(typeOfAdress)
}
