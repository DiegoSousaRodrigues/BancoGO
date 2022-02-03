package inteface

import (
	"fmt"
	"gitHub/operacoes"
	saldologs "gitHub/saldoLogs"
)

var op int
var v string

func Choose() {
	for {
		fmt.Println("\n\nDigite a opção que deseja\n1-Para verificar seu saldo e transações\n2-Para sacar\n3-Para depositar\n0-Para finalizar")
		fmt.Scan(&op)
		if next(op) {
			break
		}
	}
}

func next(op int) bool {

	if op == 1 {
		op1()
	}

	if op == 2 {
		op2()
	}

	if op == 3 {
		op3()
	}

	if op == 0 {
		return true
	}
	return false
}

func op1() {
	fmt.Println("Seu saldo -> ", operacoes.VerificarSaldo())

	for k, v := range saldologs.RetornarLogs() {
		fmt.Println("\n\nOperação numero: ", k, "\nVoce Realizou: "+v)
	}

}

func op2() {
	fmt.Println("\n\nDigite o valor a sacar")
	fmt.Scan(&v)
	a, err := operacoes.TradeToFloat64(v)

	if err != "" {
		panic("Voce inseriu algo diferente de um numero, tente novamente\n" + err)
	}

	fmt.Println("\n\n", operacoes.Sacar(a))
}

func op3() {
	fmt.Println("\n\nDigite o valor a depositar")
	fmt.Scan(&v)
	a, err := operacoes.TradeToFloat64(v)

	if err != "" {
		panic("Voce inseriu algo diferente de um numero, tente novamente\n" + err)
	}

	fmt.Println("\n\n", operacoes.Depositar(a))
}
