package main

import "fmt"

type user struct {
	nome  string
	years uint8
}

// criando método
// obrigatoriamento se coloca o tipo como sendo a struct
// nome da função coloca-se como se fosse o retorno da função
func (user1 user) save() {
	fmt.Printf("Saving data of user %s in database\n", user1.nome)
}

func (user1 user) isOfLegalAge() bool {
	return user1.years > 17
}

// adicionando metodo set
// passando como ponteiro para alterar a idade do user1 da função main
// usado quando fazemos alguma alteração de informação na struct
func (user1 *user) birthday() {
	user1.years++
}

func main() {
	user1 := user{"User number 1", 17}
	fmt.Println(user1)
	user1.save()
	fmt.Println(user1.isOfLegalAge())
	user1.birthday()
	fmt.Println(user1.years)
	fmt.Println(user1.isOfLegalAge())
}
