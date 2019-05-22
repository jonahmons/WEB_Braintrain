package main

import (
	"05Prak/app/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/loggedin", controller.IndexLoggedIn)
	http.HandleFunc("/logout", controller.IndexLogOut)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/karteikaesten", controller.Karteikaesten)
	http.HandleFunc("/meinekarteikarten", controller.MeineKarteikarten)
	http.HandleFunc("/karteikaesten/edit", controller.Edit)
	http.HandleFunc("/karteikaesten/edit2", controller.Edit2)
	http.HandleFunc("/karteikaesten/view", controller.View)
	http.HandleFunc("/karteikaesten/lern", controller.Lernen)
	http.HandleFunc("/karteikaesten/lern2", controller.Lernen2)
	http.HandleFunc("/profil", controller.MeinProfil)
	http.HandleFunc("/profil2", controller.MeinProfilModal)

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/font/", http.StripPrefix("/font", http.FileServer(http.Dir("./static/fonts"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./static/js/"))))
	http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("./static/img/"))))

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
}
