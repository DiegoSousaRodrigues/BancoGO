package operacoes

import (
	"fmt"
	saldologs "gitHub/saldoLogs"
	"strconv"
)

func VerificarSaldo() float64 {
	return saldologs.GetSaldo()
}

func Sacar(v float64) string {
	if VerificarSaldo() < v {
		return "Operação nao realizada, valor de saque maior que o da conta"
	}

	log := "Sacou " + fmt.Sprintf("%f", v)
	fmt.Println(log)
	saldologs.AdicionarLogs(log)

	saldologs.Sacar(v)
	return "Operação realizada"
}

func Depositar(v float64) string {
	log := "Depositou " + fmt.Sprintf("%f", v)
	fmt.Println(log)
	saldologs.AdicionarLogs(log)

	saldologs.Depositar(v)
	return "Operação realizada"

}

func TradeToFloat64(v string) (float64, string) {
	newV, err := strconv.Atoi(v)
	if err != nil {
		return 0, err.Error()
	}
	return float64(newV), ""
}
