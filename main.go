package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/majomi/simple-todo-api/sqlite"
	"github.com/majomi/simple-todo-api/todo"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowedMethods: []string{"DELETE", "GET", "POST"}})

	r.Use(cors.Handler, render.SetContentType(render.ContentTypeJSON))
	srv := &http.Server{
		Handler:      r,
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	storage, err := sqlite.NewTodoDB("todo.db")
	if err != nil {
		log.Fatal(err)
	}

	todoService := todo.NewService(storage)
	todoService.DestructiveReset()
	todoService.AutoMigrate()
	todoController := todo.NewController(todoService)

	r.Get("/todo", todoController.All)
	r.Get("/todo/{id:[0-9]+}", todoController.ByID)
	r.Get("/todo/{title:[a-zA-Z0-9_ ]+}", todoController.ByTitle)
	r.Post("/todo", todoController.Create)
	r.Delete("/todo/{id:[0-9]+}", todoController.Delete)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func test() {

	/* 	todoService.DestructiveReset()
	   	todoService.AutoMigrate()
	   	todoService.Create(&todo.Todo{Title: "My First Todo Item", Message: "Something really meaningfull", Completed: false, Priority: todo.PriorityHigh})

	   	todo, err := todoService.ByID(1)
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   	log.Println(todo) */
}
