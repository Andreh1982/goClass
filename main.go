package main

import (
	"fmt"
	"goClass/clientes"
	"goClass/contas"
)

func main() {

	cc := contas.ContaCorrente{
		Titular:       clientes.Titular{"Denis Alves", "345CPF567", "Logistica"},
		NumeroAgencia: 10049,
		NumeroConta:   273660,
		Saldo:         500,
	}

	cc1 := contas.ContaCorrente{
		Titular:       clientes.Titular{"Maria Silva", "777CPF888", "Recursos Humanos"},
		NumeroAgencia: 14268,
		NumeroConta:   385660,
		Saldo:         400,
	}

	done := cc.Transferir(100, cc1)
	fmt.Println(done)

	fmt.Println(cc)
	fmt.Println(cc1)
}
