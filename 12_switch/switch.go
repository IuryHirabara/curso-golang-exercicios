package main

import "fmt"

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabado"
	default:
		return "invalid number"
	}
}

func dayOfWeek2(number int) string {
	var day_of_week string
	switch {
	case number == 1:
		day_of_week = "domingo"
		fallthrough // faz com que depois de entrar no case 1, entre direto no próximo case sem analisar a condição dele
	case number == 2:
		day_of_week = "segunda"
	case number == 3:
		day_of_week = "terça"
	case number == 4:
		day_of_week = "quarta"
	case number == 5:
		day_of_week = "quinta"
	case number == 6:
		day_of_week = "sexta"
	case number == 7:
		day_of_week = "sabado"
	default:
		day_of_week = "invalid number"
	}
	return day_of_week
}

func main() {
	day := dayOfWeek(3)
	fmt.Println(day)

	day2 := dayOfWeek2(1)
	fmt.Println(day2)
}
