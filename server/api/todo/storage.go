package todo

type Storage interface {
	GetAllToDos() ([]*Model, error)
	GetToDoById(ID string) (*Model, error)
	ToggleToDoComplete(ID string) error
	ToggleToDoActive(ID string) error
}
