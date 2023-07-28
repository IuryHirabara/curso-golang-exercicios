package main

import (
	"fmt"
	"reflect"
)

func main() {
	// para criar arrays
	// arrays tem tamanho fixo
	// não é tão utilizado em go, o slice é mais utilizado
	var arr1 [5]int
	arr1[0] = 60
	fmt.Println(arr1)

	arr2 := [5]int{1: 30}
	fmt.Println(arr2)

	// para ser um pouco mais flexivel, podemos passar "..." no tamanho do array, assim, o tamanho do array se dará pela quantidade de parâmetros declarados
	arr3 := [...]int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(arr3)

	// slice
	// tamanho dinâmico
	// o slice usa Arrays Internos (explicação mais abaixo)
	slc := []int{10, 20, 30}

	fmt.Println(reflect.TypeOf(slc))
	fmt.Println(reflect.TypeOf(arr3))

	// adicionando item ao slice
	slc = append(slc, 18)
	slc[0] = 1
	fmt.Println(slc)

	// podemos retonar intervalos de um array e colocar em um slice
	// o indice 1 é inclusivo e o indice 3 é exclusivo, ou seja irá retornar o 20 e 30 pois o indice 3 é o 40
	slc2 := arr3[1:3]
	fmt.Println(slc2)

	// slice de array (ou fatia de array) funciona como um ponteiro, então se alteraramos o arr3, a alteração irá refletir no slc2
	arr3[1] = 100
	fmt.Println(slc2)

	// Arrays Internos
	// make(tipo, tamanho, capacidade (opcional) )
	// se não especificado a capacidade, esse valor será o mesmo do tamanho
	slc3 := make([]float32, 10, 11)
	fmt.Println(len(slc3)) // tamanho do slice
	fmt.Println(cap(slc3)) // capacidade do slice

	// quando o slice vê que vai estourar o limite posto, ele dobra o tamanho do slice
	slc3 = append(slc3, 5)
	slc3 = append(slc3, 6)
	fmt.Println(len(slc3), cap(slc3))
}
