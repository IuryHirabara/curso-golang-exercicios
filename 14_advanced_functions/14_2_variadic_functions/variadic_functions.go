package main

import "fmt"

// possibilita passar quantos ints quiser
// no caso, "numbers" será um slice com todos os números passados para a função multSum
// tem de ser o ultimo parametro passado para a função e não pode haver mais de um por funçõa
func multSum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := multSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}
