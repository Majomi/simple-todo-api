package todo

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	todoService Service
}

func NewController(todoService Service) *Controller {
	return &Controller{
		todoService: todoService,
	}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := render.Bind(r, &todo); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusBadRequest), err.Error()), http.StatusBadRequest)
		return
	}
	if err := c.todoService.Create(&todo); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusBadRequest), err.Error()), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) All(w http.ResponseWriter, r *http.Request) {
	todos, err := c.todoService.All()
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err.Error()), http.StatusInternalServerError)
		return
	}
	render.Respond(w, r, &todos)
}

func (c *Controller) ByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err.Error()), http.StatusInternalServerError)
		return
	}
	todo, err := c.todoService.ByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusNotFound), err.Error()), http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err.Error()), http.StatusInternalServerError)
		return
	}
	render.Respond(w, r, todo)
}

func (c *Controller) ByTitle(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")
	todo, err := c.todoService.ByTitle(title)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err.Error()), http.StatusInternalServerError)
		return
	}
	render.Respond(w, r, todo)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err.Error()), http.StatusInternalServerError)
		return
	}

	if err := c.todoService.Delete(uint(id)); err != nil {
		http.Error(w, fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
