package three

import "fmt"

func Three() {
    // определите переменные ver, id, pi
    // ...
		ver, id, pi := 1, 2, 3.14

    fmt.Println("ver =", ver, "id =", id, "pi =", pi)


		// iota используется для перечисления констант
		const (
			Black = iota
			Gray
			White
	)
	
	// счётчик обнуляется
	const (
			Yellow = iota
			Red
			Green = iota // это присваивание не обнулит iota
			Blue
	)
	
			fmt.Println(Black, Gray, White) 
			fmt.Println(Yellow, Red, Green, Blue)

			const (
				_ = iota*10  // обратите внимание, что можно пропускать константы 
				ten
				twenty
				thirty
		)
		
		const (
				hello = "Hello, world!"  // iota равна 0
				one = 1                  // iota равна 1
		
				black = iota   // iota равна 2
				gray
		)
		
				fmt.Println(ten, twenty, thirty)
				fmt.Println(black, gray)

				const (
					one1 = iota*2+1// укажите здесь формулу с iota
					three3
					five5
					seven7
					nine9
					eleven11
			)
			
					fmt.Println(one1, three3, five5, seven7, nine9, eleven11)
	}
