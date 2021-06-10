package main

import "fmt"

func printResult(result float64) {
	if result >= 7 {
		fmt.Println("Approved:", result)
	} else {
		fmt.Println("Disapproved:", result)
	}
}

func printPlayResult(play1, play2 int) {
	if play1 > play2 {
		fmt.Println("Play 1 win!")
	} else if play2 > play1 {
		fmt.Println("Play 2 win!")
	} else {
		fmt.Println("The game tied!")
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)
	printPlayResult(1, 2)
	printPlayResult(2, 1)
	printPlayResult(2, 2)
}
