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
	ToDo{
		ID:       "467f8e85-e125-4068-8280-a581a771ca2b",
		Text:     "Do the dishes",
		Complete: false,
	},
	ToDo{
		ID:       "9067ecc9-6c81-4e1b-9cbf-95eaed6efbba",
		Text:     "Feed the dogs",
		Complete: false,
	},
	ToDo{
		ID:       "c63073ee-77f9-4623-9eec-149bcabd7833",
		Text:     "Love Jamie more than ever",
		Complete: false,
	},
	ToDo{
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
func (m ToDoModel) MarkComplete(ID string) error {
	found := false
	for i := range toDos {
		if toDos[i].ID == ID {
			toDos[i].Complete = true
			found = true
		}
	}

	if found {
		return nil
	}

	return errors.New("Unable to find the requested ToDo")
}
