package webapp

type ToDoRepository interface {
	GetAllToDos() ([]*ToDo, error)
	GetToDoById(ID string) (*ToDo, error)
	ToggleToDoComplete(ID string) error
	ToggleToDoActive(ID string) error
}
