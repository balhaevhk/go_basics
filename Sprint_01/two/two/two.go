package two

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func Two() {
	var str string
	str = "hello мой друг"
	fmt.Println(str)
	fmt.Println(str[2])
	fmt.Println(len(str))
	println(utf8.RuneCountInString(str))

	type Name string
	type Fruit string

	var fruit Fruit = "apple"
	var name Name = "Alex"
	fmt.Println(fruit)
	fmt.Println(name)
	name = Name(fruit)
	fmt.Println(fruit)
	fmt.Println(name)

	var word string
	word = "Hello, world!"

	println(string(word[0]))
	println(strconv.Itoa(len(word)))



}
