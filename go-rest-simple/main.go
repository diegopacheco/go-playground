package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Todo represent a todo list item model
type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is a array of todos with all todo
type Todos []Todo

var todos Todos
var currentId int

func main() {
	todos = make([]Todo, 0)
	currentId := 1
	todos = append(todos, Todo{
		Id:        currentId,
		Name:      "TODO 1",
		Completed: false,
		Due:       time.Now(),
	})
	fmt.Printf("TODOS initialized %s \n", todos)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", Index)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	var handler http.Handler
	handler = http.HandlerFunc(TodoCreate)
	router.Methods("POST").Path("/todos").Name("TodoCreate").Handler(handler)

	fmt.Println("TODO - REST Service running on 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index is the main function called when / rest endpoint hits.
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoCreate create todo in memory
// curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	currentId = currentId + 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

// TodoShow is the main function called when / rest endpoint hits.
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := RepoFindTodo(todoId)
	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}
}

// RepoFindTodo find the todo in the list of todo
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}
