package two

import "fmt"

func div(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("делитель равен 0")
    }
    return a / b, nil
}

func foo() {
	// паникуем
	panic("unexpected!")
}
//...
	// выполняется после срабатывания паники
	
	
func Two() {
    d, err := div(10, 0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("d = %d", d)
    }


		defer func() {
			if r := recover(); r != nil {
					// обработка паники, в переменной r будет лежать строка "unexpected"
			}
	}()
	// внутри foo срабатывает паника
	foo()
}