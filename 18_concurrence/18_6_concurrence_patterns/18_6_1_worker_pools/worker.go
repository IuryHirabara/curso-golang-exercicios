package main

import "fmt"

func main() {
	// criando dois canais, um para armazenar as posições dos números de fibonacci (de 1 à 45)
	tasks := make(chan int, 45)
	// esse canal servirá para receber os resultados dos cálculos do número de fibonacci
	results := make(chan int, 45)

	// criando o processo de concorrência, não há problema em criar agora, vamos passar os números para o cálculo das posições de fibonacci no laço for seguinte
	go worker(tasks, results)
	// como a função concorrente worker está usando um único canal tasks para calcular as posições de fibonacci, ele consegue calcular "ao mesmo tempo" várias posições e terminar de calcular todas as 45 posições de forma mais rápida
	go worker(tasks, results)

	// passando as posições de fibonacci
	// assim que o canal tasks começar a receber valores, a posição de fibonacci é cálculada
	for i := 0; i < 45; i++ {
		tasks <- i
	}
	// fechando canal de tarefas
	close(tasks)

	for i := 0; i < 45; i++ {
		// assim que o canal results recebe um valor novo, ele é jogado para a variável result e em seguida é printado o resultado do cálculo da posição de fibonacci
		result := <-results
		fmt.Println(result)
	}
}

// o canal tasks só recebe dados "<-chan", o canal de results só envia dados "chan<-"
func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		// como a função worker foi declarada antes, a função worker ficará esperando um valor ser passado para as tasks
		// aqui também está jogando o cálculo da posição de fibonacci para o canal results
		results <- fibonacci(number)
	}
}

// função recursiva de fibonacci
func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
