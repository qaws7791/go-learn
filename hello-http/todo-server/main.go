package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Todo Server!"))
	})

	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		todos := getTodos()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	})
	mux.HandleFunc("POST /todos", func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		createdTodo := createTodo(todo.Name)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdTodo)
	})

	mux.HandleFunc("GET /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {

			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		todo := getTodoByID(idInt)
		if todo == nil {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)

	})
	mux.HandleFunc("PUT /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		updatedTodo := updateTodoByID(idInt, todo.Name, todo.Status)
		if updatedTodo == nil {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedTodo)
	})
	mux.HandleFunc("DELETE /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		if !deleteTodoByID(idInt) {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: Logging(mux),
	}

	log.Println("Starting server on :8080")
	server.ListenAndServe()
}

type Todo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

var todos = []Todo{
	{ID: 1, Name: "Learn Go", Status: "completed"},
	{ID: 2, Name: "Build a web server", Status: "in-progress"},
}

func getTodos() []Todo {
	return todos
}

func getTodoByID(id int) *Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo
		}
	}
	return nil
}

func createTodo(name string) Todo {
	id := len(todos) + 1
	todo := Todo{ID: id, Name: name, Status: "in-progress"}
	todos = append(todos, todo)
	return todo
}

func updateTodoByID(id int, name string, status string) *Todo {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Name = name
			todos[i].Status = status
			return &todos[i]
		}
	}
	return nil
}

func deleteTodoByID(id int) bool {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}
