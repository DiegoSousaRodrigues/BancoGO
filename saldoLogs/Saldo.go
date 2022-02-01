package saldologs

var saldo float64 = 100.00

func GetSaldo() float64 {

	return saldo
}

func SetSaldo(saldo_ float64) {
	saldo -= saldo_
}
