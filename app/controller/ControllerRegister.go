package controller

import (
	//"05Prak/app/model"
	//"fmt"
	"html/template"
	"net/http"
)
//Register wireframe
func Register(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/register.tmpl")
	t.ExecuteTemplate(w, "layout", MainContent)
}
