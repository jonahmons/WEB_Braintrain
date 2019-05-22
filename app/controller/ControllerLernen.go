package controller

import (
	"05Prak/app/model"
	"fmt"
	"html/template"
	"net/http"
)

//Lernen wireframe
func Lernen(w http.ResponseWriter, r *http.Request) {
	karteikartenID := "59ad19ed6ed193c4006530e8bb03f875"
	karteikarte, _ := model.GetKarteikarte(karteikartenID)
	karteikastenID := karteikarte.Kasten
	kasten, _ := model.GetKarteikasten(karteikastenID)
	_, category := model.GetUpperCategoryByCatID(kasten.Kategorie)

	guiLernenContent := model.GUILernenContent{
		Karteikarte:      map[string]model.GUIEinzelneKarteikarte{},
		KarteikastenInfo: map[string]model.GUIKarteikasten{},
	}

	guiLernenContent.Karteikarte[karteikartenID] = model.GUIEinzelneKarteikarte{
		ID:               karteikarte.ID,
		KarteikartenName: karteikarte.KarteikartenName,
		Frage:            karteikarte.Frage,
		Antwort:          karteikarte.Antwort,
		KastenID:         karteikarte.Kasten,
	}

	guiLernenContent.KarteikastenInfo[karteikartenID] = model.GUIKarteikasten{
		ID             : kasten.ID,
		Oberkategorie  : category,
		Unterkategorie : model.GetCategoryByID(kasten.Kategorie),
		Name           : kasten.KarteikastenName,
		AnzahlKarten   : model.CountKarteikartenInKateikasten(karteikastenID),
		Beschreibung   : kasten.Beschreibung,
		Ersteller      : kasten.Autor,
	}

	fmt.Println(guiLernenContent)

	MainContent.Content = guiLernenContent

	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/lernen.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", model.Main{})
	}
}

//Lernen2 wireframe
func Lernen2(w http.ResponseWriter, r *http.Request) {
	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/lernen2.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", model.Main{})
	}
}
