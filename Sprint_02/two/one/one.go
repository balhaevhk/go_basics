package one

import (
	"fmt"
	"os"
	"path/filepath"
)

func One() {
	println("===ONE===")
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum2(1, 2, 3, 4, 5))

	fmt.Println(Index("Hello", 'l'))

	fmt.Println(fact(5))

	fmt.Println(Fib(5))

	fmt.Println(fact2(5))

	fmt.Println(Fib2(5))

	PrintAllFiles(".")



}

func Sum(x, y int) int{
	return x + y
}

func Sum2(x ...int) (res int) {
	for _, v := range x {
			res += v
	}
	return
}
// если указано имя возвращаемого значения, то оно будет возвращено по умолчанию
func Index(st string, a rune) (index int, ok bool) {
	for i, c := range st {
			if c == a {
					return i, true
			}
	}
	return // вернутся значения по умолчанию
}

func fact(n int) int {
	if n == 0 {    // терминальная ветка — то есть условие выхода из рекурсии
			return 1
	} else {    // рекурсивная ветка 
			return n * fact(n-1)
	}
}

func fact2(n int) int {
	res := 1
	for n > 0 {
			res *= n
			n--
	}
	return res
}

func Fib(n int) int {
	switch {
	case n <= 1:    // терминальная ветка 
			return n
	default:        // рекурсивная ветка
			return Fib(n-1) + Fib(n-2)
	}
}

func Fib2(n int) int {
	a, b := 0, 1
	for n > 0 {
			a, b = b, a+b
			n--
	}
	return a
}

func PrintAllFiles(path string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
			fmt.Println("unable to get list of files", err)
			return
	}
	//  проходим по списку
	for _, f := range files {
			// получаем имя элемента
			// filepath.Join — функция, которая собирает путь к элементу с разделителями
			filename := filepath.Join(path, f.Name())
			// печатаем имя элемента
			fmt.Println(filename)
			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
			if f.IsDir() {
					PrintAllFiles(filename)
			}
	}
} 

