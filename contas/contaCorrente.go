package contas

import (
	"fmt"
	"goClass/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque Realizado com Êxito"
	} else {
		return "saldo Insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	if valorDoDeposito >= 0 {
		c.saldo = c.saldo + valorDoDeposito
		fmt.Println("Deposito Realizado com Sucesso")
		fmt.Println("saldo atual na conta:", c.Titular, c.saldo)
	} else {
		fmt.Println("Valor do Deposito Menor ou Igual a Zero")
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo = c.saldo - valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false

	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
