package webapp

// ToDo ...
type ToDo struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Complete bool      `json:"complete"`
	Active   bool      `json:"active"`
	TimeLog  []TimeLog `json:"timeLog"`
}

// TimeLog ...
type TimeLog struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
