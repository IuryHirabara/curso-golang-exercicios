// um módulo é um arquivo do conjunto de pacotes do projeto, para compilar em um arquivo só
// para criar um módulo vá até a pasta do código e digite "go mod init nome_do_modulo"
// similiar ao package.json do node

package main

import (
	"fmt"
	"module/support"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("writing from function main!")
	support.Write()

	check := checkmail.ValidateFormat("devbook@gmail.com") // para referenciar pacotes, sempre se usa a(s) palavra(s) após a última barra, aqui no caso é "checkmail"
	fmt.Println(check)
}

// para instalar pacotes externos à aplicação, usa se "go get url"
// para remover dependencias não usadas na aplicação, usa se "go mod tidy"
