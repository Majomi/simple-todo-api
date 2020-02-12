package sqlite

import (
	"github.com/jinzhu/gorm"
	"github.com/majomi/simple-todo-api/todo"
)

type TodoDB struct {
	DB *gorm.DB
}

func NewTodoDB(connectionString string) (*TodoDB, error) {
	db, err := gorm.Open("sqlite3", connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &TodoDB{
		DB: db,
	}, nil
}

// Create
func (td *TodoDB) Create(todo *todo.Todo) error {
	return td.DB.Create(todo).Error
}

// Read
func (td *TodoDB) All() ([]*todo.Todo, error) {
	var todos []*todo.Todo
	if err := td.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (td *TodoDB) ByID(id uint) (*todo.Todo, error) {
	var todo todo.Todo
	if err := td.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (td *TodoDB) ByTitle(title string) (*todo.Todo, error) {
	var todo todo.Todo
	if err := td.DB.Where("title = ?", title).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (td *TodoDB) ByUsername(username string) ([]*todo.Todo, error) {
	panic("not implemented") // TODO: Implement
}

// Update
func (td *TodoDB) Update(todo *todo.Todo) error {
	return td.DB.Save(todo).Error
}

// Delete
func (td *TodoDB) Delete(id uint) error {
	return td.DB.Delete(&todo.Todo{ID: id}).Error
}

func (td *TodoDB) AutoMigrate() error {
	return td.DB.AutoMigrate(&todo.Todo{}).Error
}

func (td *TodoDB) DestructiveReset() error {
	return td.DB.DropTableIfExists(&todo.Todo{}).Error
}
