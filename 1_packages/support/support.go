package support

import "fmt"

// Quando uma função é exportada para outro arquivo, é uma boa prática colocar um comentário do que ela faz acima
func Write() { // o nome da função tem de começar com letra maiuscula, pois assim estará visível em outros arquivos
	fmt.Println("writing from function Write!")
	write2() // posso chamar diretamente pois o support.go e support2.go estão no mesmo pacote, no caso o package support
}
