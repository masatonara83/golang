package main

import (
	"fmt"
	// "strings"
)

//クロージャー
func circleArea(pai float64) func(num float64) float64 {
	return func(num float64) float64 {
		return pai * num * num
	}
}

func totalPrice(tax float64) func(num1, num2 float64) float64 {
	return func(num1, num2 float64) float64 {
		return (num1 + num2) * (tax)
	}
}

//可変長引数
func number(nums ...int) {
	fmt.Println(len(nums), nums)
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(string("Hello World"[0]))

	// var s string = "Hello World"

	// s = strings.Replace(s, "H", "x", 1)
	// fmt.Println(s)
	// fmt.Println(strings.Contains(s, "World"))

	// fmt.Println(`TEST
	// 		Test
	// TEST`)

	// fmt.Println(`"`)

	ci := circleArea(3.14)
	fmt.Println(ci(3))

	number(1, 2)
	number(1, 2, 3)

	c2 := totalPrice(1.08)
	fmt.Println(c2(500,1000))
}
