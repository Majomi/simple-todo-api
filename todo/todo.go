package todo

import (
	"net/http"
	"time"
)

// Priority specifies the priority
type Priority uint

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

// Todo represents a Todo Item
type Todo struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	Completed bool       `json:"completed"`
	Priority  Priority   `json:"priority"`
}

func (t *Todo) Bind(r *http.Request) error {
	return nil
}

func (t *Todo) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
