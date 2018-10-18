package todo

import "github.com/artistomin/planner"

// Service is the interface that provides todo methods
type Service interface {
	Todos() ([]*planner.Todo, error)
	Todo(id planner.UUID) (*planner.Todo, error)
	CreateTodo(todo *planner.Todo) (planner.UUID, error)
}

type service struct {
	todos planner.TodoRepo
}

func (s *service) Todos() ([]*planner.Todo, error) {
	todos, err := s.todos.FindAll()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *service) Todo(id planner.UUID) (*planner.Todo, error) {
	todo, err := s.todos.FindById(id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *service) CreateTodo(todo *planner.Todo) (planner.UUID, error) {
	uuid, err := s.todos.Create(todo)

	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (s *service) UpdateTodo(id planner.UUID, todo planner.Todo) error {
	return nil
}

// NewService creates a todo service with necessary dependencies.
func NewService(todos planner.TodoRepo) Service {
	return &service{
		todos: todos,
	}
}
