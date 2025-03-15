package two

import (
	"fmt"
	"time"
)

var Global = 5

func Two() {
	println("===TWO===")
	fmt.Println(Global)
	defer changeGlobal()
	fmt.Println(Global)
	
	defer metricTime(time.Now())
}

func metricTime(start time.Time) {
	fmt.Println(time.Now().Sub(start))
}

func changeGlobal() {
	Global = 10
	fmt.Println(Global)
}