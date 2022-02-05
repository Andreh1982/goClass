package contas

import (
	"goClass/clientes"
	"testing"
)

func TestContaCorrente_Sacar(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Saldo         float64
	}
	type args struct {
		valorDoSaque float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"Teste01", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 300}, args{140}, 160},
		{"Teste02", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 800}, args{220}, 580},
		{"Teste03", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 200}, args{20}, 180},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaCorrente{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}

			c.Sacar(tt.args.valorDoSaque)

			if c.Saldo != tt.want {
				t.Errorf("Saldo em conta após saque: %v, esperado %v", c.Saldo, tt.want)
			}
		})
	}
}

func TestContaCorrente_Depositar(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Saldo         float64
	}
	type args struct {
		valorDoDeposito float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"Teste01", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 100}, args{120}, 220},
		{"Teste02", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 100}, args{300}, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaCorrente{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			c.Depositar(tt.args.valorDoDeposito)

			checkTotal := tt.fields.Saldo + tt.args.valorDoDeposito
			if tt.want != checkTotal {
				t.Errorf("ContaCorrente.Depoositar() = %v, want %v", checkTotal, tt.want)

			}
		})
	}
}

func TestContaCorrente_Transferir(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Saldo         float64
	}
	type args struct {
		contaDestino         ContaCorrente
		valorDaTransferencia float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"Teste01", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 300}, args{}, 800},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaCorrente{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}

			c.Transferir(tt.args.valorDaTransferencia, tt.args.contaDestino)

			checkTotal := c.Saldo - tt.args.valorDaTransferencia

			if checkTotal != tt.want {
				t.Errorf("ContaCorrente.Transferir() = %v, want %v", checkTotal, tt.want)
			}
		})
	}
}

func TestContaCorrente_ObterSaldo(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Saldo         float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Teste01", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 120}, 120},
		{"Teste02", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 205}, 205},
		{"Teste03", fields{clientes.Titular{"Joao", "12CPF30", "Analista"}, 10049, 273660, 200}, 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaCorrente{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			if got := c.ObterSaldo(); got != tt.want {
				t.Errorf("ContaCorrente.ObterSaldo() = %v, want %v", got, tt.want)
			}
		})
	}
}
