package controller

import (
	"05Prak/app/model"
	//"fmt"
	"html/template"
	"net/http"
)

//Edit wireframe
func Edit(w http.ResponseWriter, r *http.Request) {
	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/edit.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", model.Main{})
	}
}

//Edit2 wireframe
func Edit2(w http.ResponseWriter, r *http.Request) {
	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/edit2.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", model.Main{})
	}
}
