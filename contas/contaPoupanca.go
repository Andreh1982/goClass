package contas

import (
	"fmt"
	"goClass/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) {
	if valorDoDeposito >= 0 {
		c.saldo = c.saldo + valorDoDeposito
		fmt.Println("Deposito Realizado com Sucesso")
		fmt.Println("saldo atual na conta:", c.Titular, c.saldo)
	} else {
		fmt.Println("Valor do Deposito Menor ou Igual a Zero")
	}
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo = c.saldo - valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
