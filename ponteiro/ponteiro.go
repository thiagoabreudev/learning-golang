package main

import "fmt"

func main() {
	i := 1

	// Go nao tem aritmetica de ponteiro

	var p *int = nil

	p = &i // pegando o endereço da variável

	*p++ // desreferenciando (pegando o valor)

	i++

	// p++ // Nao pode pois p eh a referencia e ponteiro nao aceita artimetica diferentemente de *p++ que pega o valor desreferenciado

	fmt.Println(p, *p, i, &i)

}
