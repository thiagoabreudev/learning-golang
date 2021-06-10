package main

import (
	"fmt"
	"time"
)

func main() {
	//  So existe for no GO!

	i := 1

	// Exemplo 1 com condicao de parada
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// for tradicional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// loop infinito

	for {
		fmt.Println("for every...")
		time.Sleep(time.Second * 5) // every 5 seconds
	}

}
