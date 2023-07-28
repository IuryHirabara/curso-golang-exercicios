package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos de números inteiros
	// conforme número de bits
	// int8, int16, int32, int64
	// se não especificado o tipo de int, o compilador irá usar a arquitetura do sistema como base
	var number int = 10000000
	fmt.Println(number)

	// unsigned int
	// para números maiores ou iguais à 0
	var number2 uint32 = 1000
	fmt.Println(number2)

	// "rune" é exatamente igual ao int32, funciona como apelido
	var number3 rune = 1234
	fmt.Println(number3)

	// apelido para uint8
	var number4 byte = 123
	fmt.Println(number4)

	// numeros reais
	// float32, float64
	var number5 float32 = 123.45
	fmt.Println(number5)

	// por inferencia
	// irá aparecer somente como "float", mas na declaração de tipos não pode se utilizar como por exemplo "var number6 float"
	number6 := 1234.56
	fmt.Println(number6)

	// strings
	// não há o tipo "char", pois o compilador converte para números o char
	// para strings, sempre utilizar aspas duplas, se não o compilador considera como char e converte
	var string1 string = "string 1"
	fmt.Println(string1)

	// exemplo de char
	char := 'B'
	fmt.Println(char)
	// 66
	// corresponde ao número da tabela ASCII
	// se colocado mais de um caracter irá dar erro

	// valor zero
	// valor zero é o valor que é atribuido quando não inicializarmos a variável
	var text string        // ""
	var number_zero int    // 0
	var float_zero float32 // 0
	fmt.Println(text, number_zero, float_zero)

	// boleano
	var bool1 bool = true
	var bool2 bool // false
	fmt.Println(bool1, bool2)

	// tipo error
	var error1 error // <nil>
	var error2 error = errors.New("internal error")
	fmt.Println(error1, error2)
}
