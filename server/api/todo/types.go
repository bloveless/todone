package todo

import "todone-backend/api"

// Model ...
type Model struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Complete bool      `json:"complete"`
	Active   bool      `json:"active"`
	TimeLog  []TimeLog `json:"timeLog"`
}

// TimeLog ...
type TimeLog struct {
	StartTime api.JSONTime `json:"startTime"`
	EndTime   api.JSONTime `json:"endTime"`
}
