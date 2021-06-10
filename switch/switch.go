package main

import (
	"fmt"
	"time"
)

func conceptValue(n float64) string {
	// Por padrao no go, um bloco switch ja tem um break
	// Caso por algum motivo, seja necessario passar em um bloco e continuar
	// descendo, usar a palavra reservada fallthrough

	var value = int(n)
	switch value {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid value"
	}
}

// switch true, entra no primeiro bloco que for verdadeiro

func greeting() {
	t := time.Now()

	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good nigth!")
	}
}

// Switch baseado em tipo do parametro

func swithWitchType(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "func"
	default:
		return "i don't know!"
	}
}

func main() {
	fmt.Println(conceptValue(10))
	fmt.Println(conceptValue(7))
	fmt.Println(conceptValue(9))
	fmt.Println(conceptValue(100))
	fmt.Println(conceptValue(0))
	fmt.Println(conceptValue(1))
	greeting()
	fmt.Println(swithWitchType(1))
	fmt.Println(swithWitchType("1"))
	fmt.Println(swithWitchType(1.9))
	fmt.Println(swithWitchType(true))
	fmt.Println(swithWitchType(func() {}))
}
