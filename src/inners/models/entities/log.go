package entities

type Log struct {
	Id        string `json:"id" bson:"id"`
	Content   string `json:"content" bson:"content"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
	Timezone  string `json:"timezone" bson:"timezone"`
}

func NewLog(id string, content string, timestamp int64, timezone string) *Log {
	log := &Log{
		Id:        id,
		Content:   content,
		Timestamp: timestamp,
		Timezone:  timezone,
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
	if from.Timestamp != 0 {
		log.Timestamp = from.Timestamp
	}
	if from.Timezone != "" {
		log.Timezone = from.Timezone
	}
}
