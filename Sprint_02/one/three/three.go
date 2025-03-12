package three

import "fmt"

func Three() {
	println("=== THREE ===")

	m := make(map[string]int, 5)
	m["three"] = 3
	m["one"] = 1
	m["two"] = 2

	fmt.Println(m)

	type MyMap map[string]string

	var m1 MyMap = make(MyMap, 5)

	// объект готов
	m1["foo"] = "bar"
	fmt.Println(m1)

	myStringMap := map[string]string{"first": "первый", "second": "второй"}
	fmt.Println(myStringMap)

	v, ok := myStringMap["first"]
	fmt.Println(v, ok)

	var yy = make(map[string]string)
	// yy["one"] = "one"
	fmt.Println(yy)
	// if yy != nil {            // если не проверить это условие,
	yy["first"] = "bar" // то здесь можно получить panic
	// }

	// предложение
	sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
	// инициализируем map
	// ключами будут знаки, а значениями — количество вхождений
	frequency := make(map[rune]int)
	// пройдёмся по знакам в предложении
	for _, v := range sentence {
		frequency[v]++ // встреченному знаку увеличиваем частоту
	}
	// печатаем
	for k, v := range frequency {
		fmt.Printf("Знак %c встречается %d раз \n", k, v)
	}

	mapThree := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}
	fmt.Println(1)
	for k, v := range mapThree {
		if v >= 500 {
			fmt.Println(k, v)
		}
	}
	fmt.Println(2)
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	sum := 0
	for _, v := range order {
		sum += mapThree[v]
	}
	fmt.Println("Стоимость заказа:", sum)

	fmt.Println(3)
	mapFour := map[int]int{
		1:  10,
		2:  20,
		3:  30,
		4:  40,
		5:  50,
		6:  60,
		7:  70,
		8:  40,
		9:  90,
		10: 100,
	}
	ke := 80
	sliceMap := make([]int, 0, 10)
	for k, v := range mapFour {
		if v+v == ke {
			sliceMap = append(sliceMap, k)
		}
	}
	fmt.Println(sliceMap)

	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(RemoveDuplicates(input))

}

func RemoveDuplicates(input []string) []string {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				input = append(input[:j], input[j+1:]...)
				j--
			}
		}
	}
	return input
}
