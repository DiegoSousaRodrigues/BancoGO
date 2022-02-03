package saldologs

var logs = make(map[int]string)

func AdicionarLogs(message string) {
	logs[len(logs)+1] = message
}

func RetornarLogs() map[int]string {
	return logs
}
