package one

import "fmt"

func One() {
	fmt.Println("=== ONE ===")

	num := 10
	if num < 10 {
		fmt.Println("num < 10")
	} else if num == 10 {
		fmt.Println("num == 10")
	} else {
		fmt.Println("num > 10")
	}

	a := 0.10000001 // float64
	// инициализация и основное условие
	if b := float32(a); b > float32(0.1) {
		fmt.Println("Var a is GT float32(0.1)")
	}

	y := 100
	switch y {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("Default case")
	}

	switch {
	case y == 100:
		fmt.Println("EQ 100")
		fallthrough
	case y > 0:
		fmt.Println("GT 0 AND NEQ 100")
	case y < 0:
		fmt.Println("LT 0 AND NEQ 100")
	}

	yearOfBirthday := 1993
	switch {
	case 1946 < yearOfBirthday && yearOfBirthday < 1964:
		println("Привет, бумер!")
	case 1965 < yearOfBirthday && yearOfBirthday < 1980:
		println("«Привет, представитель X!")
	case 1981 < yearOfBirthday && yearOfBirthday < 1996:
		println("«Привет, миллениал!")
	case 1997 < yearOfBirthday && yearOfBirthday < 2012:
		println("«Привет, зумер!")
	case 2013 < yearOfBirthday:
		println("«Привет, альфа!")
	}

}
