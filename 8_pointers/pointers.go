package main

import "fmt"

// ponteiros salvar o endereço de memória de algo, como variáveis, constantes, etc
func main() {
	var v1 int = 10
	// cópia da variável v1
	var v2 int = v1

	v1++
	// mudou o valor só da primeira variável
	fmt.Println(v1, v2)

	// ponteiros são uma referência de memória
	var v3 int
	// criando ponteiro
	var p1 *int

	// para jogar o valor da v3 para o ponteiro p1
	v3 = 100
	p1 = &v3

	// irá mostrar "100" e o endereço de memória da variável v3
	fmt.Println(v3, p1)
	// para mostrar o valor da variável v3 pelo ponteiro p1 usamos o processo conhecido como "desreferenciação"
	fmt.Println(*p1)

	// se alterarmos o valor da variável v3, isso não afetará o endereço de memória, com isso o ponteiro p1 permanece com o mesmo endereço, e se usado o processo de desreferenciação, ele mostrará o valor da variável v3
	v3 = 150
	fmt.Println(v3, p1, *p1)
	*p1 = 200
	fmt.Println(v3, *p1)
}
