package api

import (
	"html/template"
	"net/http"

	"github.com/Alfred-Jijo/todoApp-api/storage"
)

type Server struct {
	listenAddr string
	storage    storage.PageData
}

var tmpl *template.Template

func NewServer(listenAddr string, store storage.PageData) *Server {
	return &Server{
		listenAddr: listenAddr,
		storage:    store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.home)
	http.HandleFunc("/todo", s.todo)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func (s *Server) todo(w http.ResponseWriter, r *http.Request) {
	data := storage.PageData{
		Title: "Todo List",
		Todos: []storage.Todo{
			{Item: "Jump off a cliff", IfDone: false},
			{Item: "Learn HTMX", IfDone: false},
			{Item: "Swim", IfDone: false},
			{Item: "Revise Physics", IfDone: false},
			{Item: "Drown ", IfDone: false},
		},
	}

	tmpl = template.Must(template.ParseFiles("templates/todo.html"))
	tmpl.Execute(w, data)
}
