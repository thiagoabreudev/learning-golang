package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Sub =>", a-b)
	fmt.Println("Div =>", a/b)
	fmt.Println("Mult =>", a*b)
	fmt.Println("Mod =>", a%b)
	fmt.Println("and =>", a&b) // 11 & 10 = 10, bit a bit
	fmt.Println("or => ", a|b)
	fmt.Println("xor => ", a^b)

	c := 3.0
	d := 2.0

	// other operations with math...

	fmt.Println("GT => ", math.Max(float64(a), float64(b)))
	fmt.Println("LT => ", math.Min(c, d))
	fmt.Println("Pow =>", math.Pow(c, d))
}
