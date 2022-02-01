package main

import (
	"fmt"
	"gitHub/operacoes"
)

var op int
var v string

func main() {
	for {
		fmt.Println("\n\nDigite a opção que deseja\n1-Para verificar seu saldo\n2-Para sacar\n3-Para depositar\n0-Para finalizar")
		fmt.Scan(&op)

		if op == 1 {
			fmt.Println("Seu saldo -> ", operacoes.VerificarSaldo())
		}

		if op == 2 {
			fmt.Println("\n\nDigite o valor a sacar")
			fmt.Scan(&v)
			a, erro := operacoes.TradeToFloat64(v)

			if erro != "" {
				panic("Voce inseriu algo diferente de um numero, tente novamente\n" + erro)
			}

			fmt.Println("\n\n", operacoes.Sacar(a))
		}

		if op == 0 {
			break
		}
	}
}
