package models

import (
	"errors"
	"time"
)

// ToDoModel ...
type ToDoModel struct{}

// TimeLog ...
type TimeLog struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// ToDo ...
type ToDo struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Complete bool      `json:"complete"`
	Active   bool      `json:"active"`
	TimeLog  []TimeLog `json:"timeLog"`
}

var toDos = []ToDo{
	{
		ID:       "467f8e85-e125-4068-8280-a581a771ca2b",
		Text:     "Do the dishes",
		Complete: false,
		Active:   false,
	},
	{
		ID:       "9067ecc9-6c81-4e1b-9cbf-95eaed6efbba",
		Text:     "Feed the dogs",
		Complete: false,
		Active:   false,
	},
	{
		ID:       "c63073ee-77f9-4623-9eec-149bcabd7833",
		Text:     "Love Jamie more than ever",
		Complete: false,
		Active:   false,
	},
	{
		ID:       "96d99097-6545-487a-b767-7bfe75592d77",
		Text:     "Get Jamie LASIK",
		Complete: false,
		Active:   false,
	},
}

// GetAll ...
func (m ToDoModel) GetAll() []ToDo {
	return toDos
}

// ToggleComplete ...
func (m ToDoModel) ToggleComplete(ID string) error {
	found := false
	for i := range toDos {
		if toDos[i].ID == ID {
			toDos[i].Complete = !toDos[i].Complete

			if toDos[i].Complete == true {
				toDos[i].Active = false
			}

			found = true
		}
	}

	if found {
		return nil
	}

	return errors.New("unable to find the requested ToDo")
}

// ToggleActive ...
func (m ToDoModel) ToggleActive(ID string) error {
	found := false
	for i := range toDos {
		if toDos[i].ID == ID {
			toDos[i].Active = !toDos[i].Active
			curTime := time.Now()

			if toDos[i].Active {
				toDos[i].TimeLog = append(toDos[i].TimeLog, TimeLog{curTime.String(), ""})
			} else {
				toDos[i].TimeLog = append(toDos[i].TimeLog, TimeLog{"", curTime.String()})
			}

			found = true
		} else {
			toDos[i].Active = false
		}
	}

	if found {
		return nil
	}

	return errors.New("unable to find the requested ToDo")
}

