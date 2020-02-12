package todo

// DB specifies the interaction with the underlying Database
type DB interface {
	// Create
	Create(todo *Todo) error
	// Read
	All() ([]*Todo, error)
	ByID(id uint) (*Todo, error)
	ByTitle(title string) (*Todo, error)
	ByUsername(username string) ([]*Todo, error)
	// Update
	Update(todo *Todo) error
	// Delete
	Delete(id uint) error

	AutoMigrate() error
	DestructiveReset() error
}
