package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/banco/contas"
)

func main() {
	contaDoGabriel := contas.ContaCorrente{
		Titular:       "Gabriel",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         500,
	}

	contaDoPedro := contas.ContaCorrente{
		Titular:       "Pedro",
		NumeroAgencia: 172,
		NumeroConta:   654321,
		Saldo:         700,
	}

	status := contaDoPedro.Transferir(200, &contaDoGabriel)

	fmt.Println(status)
	fmt.Println(contaDoGabriel)
	fmt.Println(contaDoPedro)
}
