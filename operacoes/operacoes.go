package operacoes

import (
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

	saldologs.SetSaldo(v)
	return "Operação realizada"
}

func TradeToFloat64(v string) (float64, string) {
	newV, err := strconv.Atoi(v)
	if err != nil {
		return 0, err.Error()
	}
	return float64(newV), ""
}
