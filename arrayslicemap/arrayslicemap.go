package main

import (
	"fmt"
	"reflect"
)

func workWithFixedArray() {
	// array homogenio (mesmo tipo) e estatico (fixo)
	var fixedArray [3]float64
	fmt.Println(fixedArray)

	fixedArray[0], fixedArray[1], fixedArray[2] = 7.8, 4.3, 9.1
	// fixedArray[3] = 9.9 // Error: Nao existe posicao 3 neste array fixo

	fmt.Println(fixedArray)
	fmt.Println(len(fixedArray)) // tamanho do array

}

func forWithRange() {
	numbers := [...]int{1, 2, 3, 4, 5} // ... fala pro compilador contar a partir dos valores

	for i, number := range numbers {
		fmt.Printf("%d) %d\n", i+1, number)
	}

	// ignore index with _

	for _, number := range numbers {
		fmt.Println(number)
	}

}

func workWithSlice() {
	array1 := [3]int{1, 2, 3} // array
	array2 := []int{1, 2, 3}  //slice

	fmt.Println(array1, array2)
	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(array2))

	array3 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de um array.
	slice1 := array3[1:3]
	fmt.Println(array3, slice1)

	// novo slice, mas aponta para o mesmo array
	slice2 := array3[:2]
	fmt.Println(array3, slice2)

	// voce pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	slice3 := array3[:1]
	fmt.Println(slice2, slice3)

}

func sliceWithMake() {
	slice := make([]int, 10)
	slice[9] = 12
	fmt.Println(slice)

	// slice com tamanho e capacidade

	slice2 := make([]int, 10, 20)
	fmt.Println(slice2, len(slice2), cap(slice2)) // cap => funcao para ver a capacidade

	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // append em go

	fmt.Println(slice2, len(slice2), cap(slice2)) // cap => funcao para ver a capacidade

	slice2 = append(slice2, 1) // Nao quebra, adiciona e automaticamente dobra a capacidade

	fmt.Println(slice2, len(slice2), cap(slice2)) // cap => funcao para ver a capacidade

}

func main() {
	workWithFixedArray()
	forWithRange()
	workWithSlice()
	sliceWithMake()
}
