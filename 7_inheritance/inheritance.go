package main

import "fmt"

type person struct {
	name      string
	last_name string
	years     uint8
	altura    uint8
}

type student struct {
	person // sem a declaração de tipo "person", estamos jogando todos os campos do struct "person" no struct "student"
	course string
	school string
}

func main() {
	person1 := person{"Joãozito", "Pereira", 24, 200}
	fmt.Println(person1)

	// instanciando student
	student1 := student{person1, "TI", "School of America"}
	// jogando todos os campos da struct person na struct student, podemos acessar de forma direta
	// ao invés de acessar com "student1.person.name" podemos acessar com "student1.name"
	fmt.Println(student1, "\n", student1.name)
}
