package one

import (
	// "bufio"
	"fmt"
	// "os"
)

func One() {
	fmt.Println("=== ONE ===")

	var a int = 5
	p := &a

	fmt.Println(a, p) //a=5 p=0xc0000b2008

	type A struct {
		IntField int
	}
	// Литерал А{} создаёт в памяти переменную типа А. Затем от неё берётся указатель
	pp := &A{
		IntField: 10,
	}

	println(pp)

	i := 42
	ppp := &i
	fmt.Println(*ppp) // читаем значение переменной i через указатель p
	*ppp = 21         // записываем в переменную i значение 21 через указатель p
	fmt.Println(*ppp)
	fmt.Println(i)

	incrementCopy := func(ii int) {
		ii++
	}

	increment := func(ii *int) {
		(*ii)++
	}

	ii := 42

	incrementCopy(ii)
	fmt.Println(ii) // 42

	increment(&ii)
	fmt.Println(ii) // 43

	aaa := 1  //1
	q := &aaa // 1
	b := &q   // 1
	println(aaa)
	println(*q)
	println(**b)
	*q = 3  // 3
	**b = 5 // 5
	println(aaa)
	println(*q)
	println(**b)

	aaa += 4 + *q + **b // 14

	fmt.Printf("%d", *q)
	println()

	// Получаем читателя пользовательского ввода
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Interaction counter")

	// cnt := 0
	// for {
	// 	fmt.Print("-> ")
	// 	// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
	// 	_, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	f(&cnt)

	// 	fmt.Printf("User input %d lines\n", cnt)
	// }

}

// func f(cnt *int) {
// 	*cnt++
// }
