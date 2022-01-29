package saldologs

var logs map[int]string

func adicionarLogs(types string) {
	logs[len(logs)+1] = types
}
