package entities

type Log struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

func NewLog(id string, content string) *Log {
	log := &Log{
		Id:      id,
		Content: content,
	}
	return log
}
