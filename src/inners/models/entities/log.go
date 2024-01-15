package entities

import "time"

type Log struct {
	Id        string    `json:"id" bson:"id"`
	Content   string    `json:"content" bson:"content"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

func NewLog(id string, content string, timestamp time.Time) *Log {
	log := &Log{
		Id:        id,
		Content:   content,
		Timestamp: timestamp,
	}
	return log
}

func (log *Log) Patch(from *Log) {
	if from.Id != "" {
		log.Id = from.Id
	}
	if from.Content != "" {
		log.Content = from.Content
	}
	if !from.Timestamp.IsZero() {
		log.Timestamp = from.Timestamp
	}
}
