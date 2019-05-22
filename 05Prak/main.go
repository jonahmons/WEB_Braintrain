package main

import (
	"net/http"
	"05 Prak/app/controller"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/add-todo", controller.AddTodo)
	http.HandleFunc("/toggle-done", controller.ToggleDone)
	http.HandleFunc("/delete-todo", controller.DeleteTodo)

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts", http.FileServer(http.Dir("./static/fonts"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./static/js"))))

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
}
