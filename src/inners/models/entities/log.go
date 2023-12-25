package entities

type Log struct {
	id      string
	content string
}

func NewLog(id string, content string) *Log {
	log := &Log{
		id:      id,
		content: content,
	}
	return log
}
