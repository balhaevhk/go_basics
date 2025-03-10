package two

import "fmt"

func Two() {
	fmt.Println("=== TWO ===")

outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLoopLabel
		}
	}
	fmt.Println("End")

outerLoopLabel2:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			continue outerLoopLabel2
		}
	}
	fmt.Println("End")

	group := 0
	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			if i%10 == 0 {
				group++
				break // break относится к ближайшему switch
			}
			fmt.Printf("%02d: %d\n", group, i)
		default:
		}
	}


	for i := 0; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}


	for i := 1; i <= 100; i++ {
		found := false

		// проверяем, что кратно 3, и выводим Fizz
		if i%3 == 0 {
				fmt.Printf("Fizz")
				found = true
		}
		// проверяем, что кратно 5, и выводим Buzz
		if i%5 == 0 {
				fmt.Printf("Buzz")
				found = true
		}
		// если число кратно 3 и 5, выводим FizzBuzz

		if !found {
				// если не выполнилось ни одно из условий, выводим число
				fmt.Println(i)
				continue
		}

		fmt.Println()
}
}
