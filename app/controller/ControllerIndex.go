package controller

import (
	"05Prak/app/model"
	//"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

//MainContent Global variable
var MainContent = model.Main{}

// Is executed automatically on package load
func init() {
	//tmpl = template.Must(template.ParseGlob("template/*.tmpl")
}

// Index Controller
func Index(w http.ResponseWriter, r *http.Request) {
	// InitDatenbankKaesten()
	// InitDatenbankUser()
	// InitDatenbankKarteikarten()
	//InitDatenbankFortschritt()

	MainContent.Content = model.GUIStartpage{
		Nutzer:     model.CountUsers(),
		Lernkarten: model.CountKarteikarten(),
		Karteien:   model.CountKarteikaesten(),
	}

	t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
	t.ExecuteTemplate(w, "layout", MainContent)
}
//InitDatenbankFortschritt function
func InitDatenbankFortschritt(){
	fortschrittkarten:= []model.Fortschrittkarte{}

	fortschrittkarten = append(fortschrittkarten, model.Fortschrittkarte{
		Nr: "59ad19ed6ed193c4006530e8bb03f875",
		Phase: "0",
	})
	fortschrittkarten = append(fortschrittkarten, model.Fortschrittkarte{
		Nr: "59ad19ed6ed193c4006530e8bb03f9f3",
		Phase: "1",
	})
	fortschrittkarten = append(fortschrittkarten, model.Fortschrittkarte{
		Nr: "59ad19ed6ed193c4006530e8bb03fca1",
		Phase: "2",
	})


	fortschrittKasten := model.Fortschrittkasten{
		ID: "",
		Rev: "",
		Type: "fortschrittkasten",
		User: "743548efdc643c8b9346ad38ac0008d5",
		Kasten: "59ad19ed6ed193c4006530e8bb002d0f",
		Fortschrittgesamt: "99",
		Fortschrittkarten: fortschrittkarten,
	}
	fortschrittKasten.Add()
}

// InitDatenbankKaesten inizialisiert die Datenbank mit Anfangswerten für Karteikästen
func InitDatenbankKaesten() {
	karteikasten := model.Karteikasten{
		ID:               "",
		Rev:              "",
		Type:             "karteikasten",
		KarteikastenName: "KateikastenName0",
		Beschreibung:     "EineTolleBeschreibung 0 oder?",
		Kategorie:        1,
		Autor:            "Jonah",
		Sichtbarkeit:     "öffentlich",
	}
	karteikasten.Add()
	karteikasten1 := model.Karteikasten{
		ID:               "",
		Rev:              "",
		Type:             "karteikasten",
		KarteikastenName: "KateikastenName1",
		Beschreibung:     "EineTolleBeschreibung 1 oder?",
		Kategorie:        2,
		Autor:            "Jonah",
		Sichtbarkeit:     "öffentlich",
	}
	karteikasten1.Add()
	karteikasten2 := model.Karteikasten{
		ID:               "",
		Rev:              "",
		Type:             "karteikasten",
		KarteikastenName: "KateikastenName2",
		Beschreibung:     "EineTolleBeschreibung  2 oder?",
		Kategorie:        3,
		Autor:            "Jonah",
		Sichtbarkeit:     "privat",
	}
	karteikasten2.Add()

	//test, _ := model.GetAllKarteikaesten()
	//fmt.Println(test)
}

// InitDatenbankUser inizialisiert die Datenbank mit Anfangswerten für User
func InitDatenbankUser() {
	user := model.User{
		ID:       "",
		Rev:      "",
		Type:     "user",
		Username: "Jonah",
		Password: "sicher",
		Email:    "j@j.de",
		Date:     "10.05.2018",
	}
	user.Add()
}

// InitDatenbankKarteikarten inizialisiert die Datenbank mit Anfangswerten für Karten
func InitDatenbankKarteikarten() {
	karteikarte := model.Karteikarte{
		ID: "",
		Rev: "",
		Type: "karteikarte",
		KarteikartenName: "Titel der Karte",
		Frage: "Text von der Frage",
		Antwort : "Text von der Antwort",
		Kasten: "59ad19ed6ed193c4006530e8bb002d0f",
	}
	karteikarte.Add()
	karteikarte2 := model.Karteikarte{
		ID: "",
		Rev: "",
		Type: "karteikarte",
		KarteikartenName: "Titel der Karte2",
		Frage: "Text von der Frage3",
		Antwort : "Text von der Antwort3",
		Kasten: "59ad19ed6ed193c4006530e8bb002d0f",
	}
	karteikarte2.Add()
	karteikarte3 := model.Karteikarte{
		ID: "",
		Rev: "",
		Type: "karteikarte",
		KarteikartenName: "Titel der Karte3",
		Frage: "Text von der Frage3",
		Antwort : "Text von der Antwort3",
		Kasten: "59ad19ed6ed193c4006530e8bb002d0f",
	}
	karteikarte3.Add()

	//test, _ := model.GetAllKarteikarten()
	//fmt.Println(test)
}
