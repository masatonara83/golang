package main

import "fmt"

func main() {
	/*
		var (
				u8  uint      = 255
				i8  int8      = 127
				f32 float32   = 0.2
				c64 complex64 = -5 + 12i
		)
		fmt.Println(u8, i8, f32, c64)

		fmt.Printf("type=%T value=%v", u8, u8)
	*/
	// x := 1 + 1
	// fmt.Println(x)
	// fmt.Println("1 + 1 =", 1+1)
	// fmt.Println("10.0 / 3 =", 10/3.0)

	// y := 0
	// fmt.Println(y)

	fmt.Println(1 << 0) // 0001 0001 1
	fmt.Println(1 << 1) // 0001 0010 2
	fmt.Println(1 << 2) // 0001 0100 4
	fmt.Println(1 << 3) // 0001 1000 8
}
