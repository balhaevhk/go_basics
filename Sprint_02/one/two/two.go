package two

import (
	"fmt"
	"reflect"
)

func Two() {
	fmt.Println("=== TWO ===")

	var week [7]int
	fmt.Println(week)

	thisWeekTemp := [7]int{-3, 5, 7}    // [-3 5 7 0 0 0 0]
	rgbColor := [3]uint8{255, 255, 128} // [255 255 128]

	fmt.Println(thisWeekTemp)
	fmt.Println(rgbColor)

	rgbColor1 := [...]uint8{255, 255, 128} // [255 255 128] len = 3
	rgbColor2 := []uint8{255, 255, 128}    // [255 255 128] len = 3

	fmt.Println(len(rgbColor1))
	rgbColor2 = append(rgbColor2, 29)
	fmt.Println(rgbColor2)

	thisWeek := [7]int{6: 11, 2: 3} // [0 0 3 0 0 0 11]
	fmt.Println(thisWeek)

	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

	sumTemp := 0

	for i := 0; i < len(weekTemp); i++ {
		sumTemp += weekTemp[i]
	}

	average := sumTemp / len(weekTemp)
	fmt.Println(average)

	var weekTemp2 = [7]int{5, 4, 6, 8, 11, 9, 5}

	sumTemp2 := 0

	for _, temp := range weekTemp2 {
		sumTemp2 += temp
	}

	average2 := sumTemp2 / len(weekTemp2)
	fmt.Println(average2)

	var weekTemp3 = [7]int{5, 4, 6, 8, 11, 9, 5}
	// weekTemp [5 4 6 8 11 9 5]
	for _, temp3 := range weekTemp3 {
		temp3 = 0
		println(temp3)
	}
	// weekTemp [5 4 6 8 11 9 5] ! — значения не изменились
	// если значение элемента не используется, можно опустить вторую переменную
	for i := range weekTemp3 {
		weekTemp3[i] = 0
	}
	// weekTemp [0 0 0 0 0 0 0] ! — значения изменились

	// mySlice := make([]TypeOfelement, LenOfslice, CapOfSlice)
	mySlice2 := make([]int, 0)     // слайс [], базовый массив []
	mySlice3 := make([]int, 5)     // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
	mySlice4 := make([]int, 5, 10) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)
	fmt.Println(mySlice4)

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7

	sli := []int{1, 9, 8, 7, 6}
	fmt.Println(sli)
	sli = append(sli, 9, 9, 9)
	fmt.Println(sli)

	s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
	// 1. Создаём слайс s с базовым массивом на 7 элементов.
	// Четыре первых элемента будут доступны в слайсе.

	slice1 := append(s[:2], 2, 3, 4)
	fmt.Println(s, slice1)                        // [0 0 2 3] [0 0 2 3 4]
	fmt.Println(s, len(s), cap(s))                // [0 0 2 3] 4 7
	fmt.Println(slice1, len(slice1), cap(slice1)) // [0 0 2 3] 4 7
	slice2 := append(s[:2], 5, 6, 7, 8)
	fmt.Println(s, len(s), cap(s))                // [0 0 2 3] 4 7
	fmt.Println(slice1, len(slice1), cap(slice1)) // [0 0 2 3] 4 7
	fmt.Println(slice2, len(slice2), cap(slice2)) // [0 0 2 3] 4 7

	// Полезные приёмы для работы со слайсами
	// Удаление последнего элемента слайса
	sl := []int{1, 2, 3}
	if len(sl) != 0 { // защищаемся от паники
		sl = sl[:len(sl)-1]
	}
	fmt.Println(sl) // [1 2]
	// Удаление первого элемента слайса
	slic := []int{1, 2, 3}
	if len(slic) != 0 { // защищаемся от паники
		slic = slic[1:]
	}
	fmt.Println(slic) // [2 3]

	// Удаление элемента слайса с индексом i:
	slice := []int{1, 2, 3, 4, 5}
	i := 2

	if len(slice) != 0 && i < len(slice) { // защищаемся от паники
		slice = append(slice[:i], slice[i+1:]...)
	}
	fmt.Println(slice) // [1 2 4 5]

	// Сравнение двух слайсов:

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []string{"1", "2", "3"}
	s4 := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(s1, s2)) // false
	fmt.Println(reflect.DeepEqual(s1, s3)) // false
	fmt.Println(reflect.DeepEqual(s1, s4)) // true

	newSlice := []int{}
	for i := 1; i <= 100; i++ {
		newSlice = append(newSlice, i)
	}
	fmt.Println(newSlice)
	newSlice = append(newSlice[:10], newSlice[90:]...)
	fmt.Println(newSlice)
	reverse(newSlice)
	fmt.Println(newSlice)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
