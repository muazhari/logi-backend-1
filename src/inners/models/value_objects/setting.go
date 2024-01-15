package value_objects

type Setting struct {
	IndexRetainTime string `json:"index_retain_time" bson:"index_retain_time"`
}

func NewSetting(indexRetainTime string) *Setting {
	setting := &Setting{
		IndexRetainTime: indexRetainTime,
	}
	return setting
}

func (setting *Setting) Patch(from *Setting) {
	if from.IndexRetainTime != "" {
		setting.IndexRetainTime = from.IndexRetainTime
	}
}
