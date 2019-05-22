package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"web-ss19/todo-01/app/model"
)

var todoList = make(model.Todos)
var tmpl *template.Template

// Is executed automatically on package load
func init() {
	tmpl = template.Must(template.ParseGlob("template/*.tmpl"))
}

// Index controller
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "todo.tmpl", todoList)
}

// AddTodo controller
func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		action := r.FormValue("action")

		if len(action) != 0 {
			todoList.Add(action)
		}
	}

	tmpl.ExecuteTemplate(w, "todo.tmpl", todoList)
}

// ToggleDone controller
func ToggleDone(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	todoList.Toggle(id)

	tmpl.ExecuteTemplate(w, "todo.tmpl", todoList)
}

// DeleteTodo controller
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	todoList.Delete(id)

	tmpl.ExecuteTemplate(w, "todo.tmpl", todoList)
}
