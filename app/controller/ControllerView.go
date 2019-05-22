package controller

import (
	"05Prak/app/model"
	//"fmt"
	"html/template"
	"net/http"
)

//View wireframe
func View(w http.ResponseWriter, r *http.Request) {
	karteikastenID := "59ad19ed6ed193c4006530e8bb002d0f"
	karteikartenID := "59ad19ed6ed193c4006530e8bb03f875"
	karteikarten, _ := model.GetAllKarteikartenByKasten(karteikastenID)
	userID := "743548efdc643c8b9346ad38ac0008d5"
	fortschritkastenAnduser, _ := model.GetFortschrittkastenByKastenIDAndUserID(karteikastenID, userID)

	guiKarteikartenContent := model.GUIKarteikartenContent{
		Karteikarten:        map[string]model.GUIKarteikarten{},
		KarteikastenInfo:    map[string]model.GUIKarteikasten{},
		EinzelneKarteikarte: map[string]model.GUIEinzelneKarteikarte{},
		Fortschritt:         map[string]model.GUIFortschritt{},
	}
	count := 1
	for _, kk := range karteikarten {
		guiKarteikarten := model.GUIKarteikarten{
			Position:         count,
			ID:               kk.ID,
			KarteikartenName: kk.KarteikartenName,
			Frage:            kk.Frage,
			Antwort:          kk.Antwort,
			KastenID:         kk.Kasten,
		}
		guiKarteikartenContent.Karteikarten[kk.ID] = guiKarteikarten
		//fmt.Println(count)
		count++
	}
	karteikastenInfo, _ := model.GetKarteikasten(karteikastenID)
	_, category := model.GetUpperCategoryByCatID(karteikastenInfo.Kategorie)
	user, _ := model.GetUser(karteikastenInfo.Autor)
	userName := user.Username
	karteikarte, _ := model.GetKarteikarte(karteikartenID)

	guiKarteikartenContent.KarteikastenInfo[karteikastenID] = model.GUIKarteikasten{
		ID:             karteikastenInfo.ID,
		Oberkategorie:  category,
		Unterkategorie: model.GetCategoryByID(karteikastenInfo.Kategorie),
		Name:           karteikastenInfo.KarteikastenName,
		AnzahlKarten:   model.CountKarteikartenInKateikasten(karteikastenID),
		Beschreibung:   karteikastenInfo.Beschreibung,
		Ersteller:      userName,
	}

	guiKarteikartenContent.EinzelneKarteikarte[karteikartenID] = model.GUIEinzelneKarteikarte{
		ID:               karteikarte.ID,
		KarteikartenName: karteikarte.KarteikartenName,
		Frage:            karteikarte.Frage,
		Antwort:          karteikarte.Antwort,
		KastenID:         karteikarte.Kasten,
	}

	fortschrittkarten := []model.GUIFortschrittKarte{}
	for _, i := range fortschritkastenAnduser.Fortschrittkarten{
		fortschrittkarten = append(fortschrittkarten, model.GUIFortschrittKarte{
			Nr : i.Nr,
			Phase : i.Phase,
		})
	}
	guiKarteikartenContent.Fortschritt[karteikastenID] = model.GUIFortschritt{
		ID: fortschritkastenAnduser.ID,
		User: fortschritkastenAnduser.User,
		Kasten: fortschritkastenAnduser.Kasten,
		Fortschrittgesamt: fortschritkastenAnduser.Fortschrittgesamt,
		Fortschrittkarten: fortschrittkarten,
	}
	//fmt.Println("KarteikartenID: " + karteikartenID)
	//fmt.Printf("%+v\n", guiKarteikartenContent.Fortschritt[karteikastenID])

	//fmt.Printf("%+v\n", guiKarteikartenContent)
	//fmt.Printf("USERNAME: " + userName)
	MainContent.Content = guiKarteikartenContent

	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/karteikastenAnschauen.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", model.Main{})
	}
}
