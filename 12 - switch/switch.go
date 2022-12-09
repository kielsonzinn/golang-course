package main

import "fmt"

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "invalid number"
	}

}

func diaDaSemana2(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Sunday"
		//fallthrough
	case numero == 2:
		diaDaSemana = "Monday"
	case numero == 3:
		diaDaSemana = "Tuesday"
	case numero == 4:
		diaDaSemana = "Wednesday"
	case numero == 5:
		diaDaSemana = "Thursday"
	case numero == 6:
		diaDaSemana = "Friday"
	case numero == 7:
		diaDaSemana = "Saturday"
	default:
		diaDaSemana = "invalid number"
	}

	return diaDaSemana

}

func main() {

	fmt.Println("switch")

	day := diaDaSemana2(5)
	fmt.Println(day)

	fmt.Println("-------------")
	day2 := diaDaSemana2(1)
	fmt.Println(day2)

}
