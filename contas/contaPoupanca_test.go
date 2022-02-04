package contas

import (
	"goClass/clientes"
	"testing"
)

func TestContaPoupanca_Depositar(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Operacao      int
		saldo         float64
	}
	type args struct {
		valorDoDeposito float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaPoupanca{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Operacao:      tt.fields.Operacao,
				saldo:         tt.fields.saldo,
			}
			c.Depositar(tt.args.valorDoDeposito)
		})
	}
}

func TestContaPoupanca_Transferir(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Operacao      int
		saldo         float64
	}
	type args struct {
		valorDaTransferencia float64
		contaDestino         ContaCorrente
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaPoupanca{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Operacao:      tt.fields.Operacao,
				saldo:         tt.fields.saldo,
			}
			if got := c.Transferir(tt.args.valorDaTransferencia, tt.args.contaDestino); got != tt.want {
				t.Errorf("ContaPoupanca.Transferir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContaPoupanca_ObterSaldo(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Operacao      int
		saldo         float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaPoupanca{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Operacao:      tt.fields.Operacao,
				saldo:         tt.fields.saldo,
			}
			if got := c.ObterSaldo(); got != tt.want {
				t.Errorf("ContaPoupanca.ObterSaldo() = %v, want %v", got, tt.want)
			}
		})
	}
}
