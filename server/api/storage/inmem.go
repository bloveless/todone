package storage

import (
	"errors"
	"sync"
	"time"
	"todone-backend/api"
	"todone-backend/api/todo"
)

type InMemoryStorage struct {
	mtx sync.Mutex
}

var toDos = []*todo.Model{
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
	{
		ID:       "4a0df66a-7036-4e0c-a141-b06cf9d3e713",
		Text:     "This is a simple test.",
		Complete: false,
		Active:   false,
	},
}

// GetAll ...
func (s InMemoryStorage) GetAllToDos() ([]*todo.Model, error) {
	return toDos, nil
}

func (s InMemoryStorage) GetToDoById(ID string) (*todo.Model, error) {
	for i := range toDos {
		if toDos[i].ID == ID {
			return toDos[i], nil
		}
	}

	return nil, errors.New("unable to find requested ToDo")
}

// ToggleComplete ...
func (s InMemoryStorage) ToggleToDoComplete(ID string) error {
	found := false
	for i := range toDos {
		if toDos[i].ID == ID {
			s.mtx.Lock()
			toDos[i].Complete = !toDos[i].Complete
			if toDos[i].Complete == true {
				// If the current ToDo was active then we need to deactivate it and record it.
				if toDos[i].Active == true {
					curTime := time.Now()
					toDos[i].TimeLog = append(toDos[i].TimeLog, todo.TimeLog{EndTime: api.JSONTime(curTime)})
					toDos[i].Active = false
				}
			}
			s.mtx.Unlock()

			found = true
		}
	}

	if found {
		return nil
	}

	return errors.New("unable to find the requested ToDo")
}

// ToggleActive ...
func (s InMemoryStorage) ToggleToDoActive(ID string) error {
	found := false
	for i := range toDos {
		curTime := time.Now()

		if toDos[i].ID == ID {
			s.mtx.Lock()
			toDos[i].Active = !toDos[i].Active
			if toDos[i].Active {
				toDos[i].TimeLog = append(toDos[i].TimeLog, todo.TimeLog{StartTime: api.JSONTime(curTime)})
			} else {
				toDos[i].TimeLog = append(toDos[i].TimeLog, todo.TimeLog{EndTime: api.JSONTime(curTime)})
			}
			s.mtx.Unlock()

			found = true
		} else {
			// If this active was active then we need to record that it is no longer active.
			if toDos[i].Active {
				s.mtx.Lock()
				toDos[i].TimeLog = append(toDos[i].TimeLog, todo.TimeLog{EndTime: api.JSONTime(curTime)})
				s.mtx.Unlock()
			}
			toDos[i].Active = false
		}
	}

	if found {
		return nil
	}

	return errors.New("unable to find the requested ToDo")
}
