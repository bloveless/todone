package models

import "errors"

// ToDoModel ...
type ToDoModel struct{}

// ToDo ...
type ToDo struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	Complete bool   `json:"complete"`
}

var toDos = []ToDo{
	{
		ID:       "467f8e85-e125-4068-8280-a581a771ca2b",
		Text:     "Do the dishes",
		Complete: false,
	},
	{
		ID:       "9067ecc9-6c81-4e1b-9cbf-95eaed6efbba",
		Text:     "Feed the dogs",
		Complete: false,
	},
	{
		ID:       "c63073ee-77f9-4623-9eec-149bcabd7833",
		Text:     "Love Jamie more than ever",
		Complete: false,
	},
	{
		ID:       "96d99097-6545-487a-b767-7bfe75592d77",
		Text:     "Get Jamie LASIK",
		Complete: false,
	},
}

// GetAll ...
func (m ToDoModel) GetAll() []ToDo {
	return toDos
}

// MarkComplete ...
func (m ToDoModel) ToggleComplete(ID string) error {
	found := false
	for i := range toDos {
		if toDos[i].ID == ID {
			toDos[i].Complete = !toDos[i].Complete
			found = true
		}
	}

	if found {
		return nil
	}

	return errors.New("unable to find the requested ToDo")
}
