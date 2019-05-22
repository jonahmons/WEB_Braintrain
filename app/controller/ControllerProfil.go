package controller

import (
	"05Prak/app/model"
	//"fmt"
	"html/template"
	"net/http"
)

//MeinProfil wireframe
func MeinProfil(w http.ResponseWriter, r *http.Request) {
	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/meinProfil.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", model.Main{})
	}
}

//MeinProfilModal wireframe
func MeinProfilModal(w http.ResponseWriter, r *http.Request) {
	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/meinProfilModal.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", model.Main{})
	}
}
