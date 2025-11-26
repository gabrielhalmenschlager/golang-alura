package main

import "fmt"

type Pizza struct {
	ID    int
	Nome  string
	Preco float64
}

func main() {
	nomePizzaria := "Pizzaria Go"
	instagram, telefone := "@pizzaria_go", 555199999999

	fmt.Println(nomePizzaria, instagram, telefone)
}
