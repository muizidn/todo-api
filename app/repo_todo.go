package app

type Todo struct {
	uuid        string
	title       string
	description string
	createdAt   int64
}

type TodoRepo interface {
	List() ([]*Todo, error)
	Get(uuid string) (*Todo, error)
	Create(title string, description string) (*Todo, error)
	Update(todo Todo) (*Todo, error)
	Delete(uuid string) error
}
