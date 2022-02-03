package inteface

import (
	"fmt"
	"gitHub/operacoes"
)

var op int
var v string

func Choose() {
	for {
		fmt.Println("\n\nDigite a opção que deseja\n1-Para verificar seu saldo\n2-Para sacar\n3-Para depositar\n0-Para finalizar")
		fmt.Scan(&op)
		if next(op) == true {
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

	if op == 0 {
		return true
	}
	return false
}

func op1() {
	fmt.Println("Seu saldo -> ", operacoes.VerificarSaldo())
}

func op2() {
	fmt.Println("\n\nDigite o valor a sacar")
	fmt.Scan(&v)
	a, erro := operacoes.TradeToFloat64(v)

	if erro != "" {
		panic("Voce inseriu algo diferente de um numero, tente novamente\n" + erro)
	}

	fmt.Println("\n\n", operacoes.Sacar(a))
}
