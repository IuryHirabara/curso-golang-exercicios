package main

import "fmt"

func recoverApprovation() {
	// nil == null
	if r := recover(); r != nil {
		fmt.Println("Execution has been recovered!")
	}
}

func approvedStudent(n1, n2 float64) bool {
	// Para recuperar execuções interrompidas pelo panic, é usado o recover com defer
	defer recoverApprovation()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// com o panic, a execução do programa é interrompida imediatamente
	// depois do panic, as unicas coisas que serão executadas são as que são marcadas com defer
	// panic não é erro
	panic("The media is exacly 6!!!")
}

func main() {
	fmt.Println(approvedStudent(7, 6))
	fmt.Println(approvedStudent(6, 6))
	fmt.Println("Post execution")
}
