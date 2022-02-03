package saldologs

var saldo float64 = 100.00

func GetSaldo() float64 {

	return saldo
}

func setSaldo(saldo_ float64) {
	saldo = saldo_
}

func Sacar(v float64) {
	setSaldo(saldo - v)
}

func Depositar(v float64) {
	setSaldo(saldo + v)
}
