package controller

import (
	"05Prak/app/model"
	//"fmt"
	"html/template"
	"net/http"
)

// IndexLoggedIn Controller
func IndexLoggedIn(w http.ResponseWriter, r *http.Request) {
	user, _ := model.GetUser("743548efdc643c8b9346ad38ac0008d5")
	MainContent.Username = user.Username
	MainContent.LoggedIn = true
	//fmt.Print(user)
	t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/startpage.tmpl")

	t.ExecuteTemplate(w, "layout", MainContent)
}

// IndexLogOut Controller
func IndexLogOut(w http.ResponseWriter, r *http.Request) {
	MainContent.LoggedIn = false
	//fmt.Print(user)
	t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")

	t.ExecuteTemplate(w, "layout", MainContent)
}
