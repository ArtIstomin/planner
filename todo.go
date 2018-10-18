package planner

import "time"

// UUID uniquely identifies a particular todo.
type UUID string

// Todo is the central struct in the domain model.
type Todo struct {
	ID        UUID      `json:"id,omitempty" bson:"_id,omitempty"`
	Text      string    `json:"text"`
	Status    Status    `json:"status,number,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// TodoRepo provides access to todo.
type TodoRepo interface {
	FindByID(id UUID) (*Todo, error)
	FindAll() ([]*Todo, error)
	Create(todo *Todo) (UUID, error)
	Update(id UUID, todo *Todo) error
	DeleteByID(id UUID) error
	Exists(id UUID) (bool, error)
}

// Status describes status of todo
type Status uint8

// Valid todo statuses
const (
	Incomplete Status = iota + 1
	Inprogress
	Done
)

func (s Status) String() string {
	switch s {
	case Incomplete:
		return "To Do"
	case Inprogress:
		return "In Progress"
	case Done:
		return "Done"
	}

	return ""
}
