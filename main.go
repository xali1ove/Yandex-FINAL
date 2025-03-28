package main

import (
	"log"
	"net/http"
	"os"

	dbfuncs "github.com/xali1ove/Yandex-FINAL/database"
	"github.com/xali1ove/Yandex-FINAL/handler"
)

func main() {

	db, err := dbfuncs.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	h := &handler.Handler{DB: db}

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = ":7540"
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("web")))

	mux.HandleFunc("GET /api/nextdate", handler.NextDateHandler)
	mux.HandleFunc("POST /api/task", h.CreateTaskHandler)
	mux.HandleFunc("GET /api/task", h.GetTaskHandler)
	mux.HandleFunc("PUT /api/task", h.UpdateTaskHandler)
	mux.HandleFunc("DELETE /api/task", h.DeleteTaskHandler)
	mux.HandleFunc("GET /api/tasks", h.TasksHandler)
	mux.HandleFunc("POST /api/task/done", h.TaskDoneHandler)
	
	log.Printf("Сервер слушает на порту %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
