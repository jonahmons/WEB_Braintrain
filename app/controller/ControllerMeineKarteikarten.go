package controller

import (
	"05Prak/app/model"
	//"fmt"
	"html/template"
	"net/http"
)

//MeineKarteikarten wireframes
func MeineKarteikarten(w http.ResponseWriter, r *http.Request) {

	eigeneKarteikasten, _ := model.GetAllKarteikaestenByUserID("743548efdc643c8b9346ad38ac0008d5")
	andereKarteikasten, _ := model.GetAllPublicKarteikaesten()
	kategorien := map[int]model.GUIEigenKarteikastenKategorien{}

	guiEigenKarteikartenContent := model.GUIEigenKarteikastenContent{
		EigenKarteiKaesten: map[string]model.GUIEigenKarteikasten{},
		AndereKarteiKaesten:  map[string]model.GUIEigenKarteikasten{},
	}

	for _, kk := range eigeneKarteikasten {
		_, category := model.GetUpperCategoryByCatID(kk.Kategorie)

		GUIEigenKarteikasten := model.GUIEigenKarteikasten{
			ID:             kk.ID,
			Kategorie:      category,
			Unterkategorie: model.GetCategoryByID(kk.Kategorie),
			Name:           kk.KarteikastenName,
			AnzahlKarten:   model.CountKarteikartenInKateikasten(kk.ID),
			Beschreibung:   kk.Beschreibung,
			Sichtbarkeit:   kk.Sichtbarkeit,
			Fortschritt:    0,
		}
		id, name := model.GetUpperCategoryByCatID(kk.Kategorie)

		if _, ok := kategorien[id]; !ok {
			kategorien[id] = model.GUIEigenKarteikastenKategorien{
				Name:          name,
				Karteikaesten: map[string]model.GUIEigenKarteikasten{},
			}
		}
		guiEigenKarteikartenContent.EigenKarteiKaesten[kk.ID] = GUIEigenKarteikasten
	}

	for _, kk := range andereKarteikasten {
		_, category := model.GetUpperCategoryByCatID(kk.Kategorie)

		GUIAndereKarteikasten := model.GUIEigenKarteikasten{
			ID:             kk.ID,
			Kategorie:      category,
			Unterkategorie: model.GetCategoryByID(kk.Kategorie),
			Name:           kk.KarteikastenName,
			AnzahlKarten:   model.CountKarteikartenInKateikasten(kk.ID),
			Beschreibung:   kk.Beschreibung,
			Sichtbarkeit:   "Andere",
			Fortschritt:    0,
		}
		id, name := model.GetUpperCategoryByCatID(kk.Kategorie)

		if _, ok := kategorien[id]; !ok {
			kategorien[id] = model.GUIEigenKarteikastenKategorien{
				Name:          name,
				Karteikaesten: map[string]model.GUIEigenKarteikasten{},
			}
		}
		guiEigenKarteikartenContent.AndereKarteiKaesten[kk.ID] = GUIAndereKarteikasten
	}
	//fmt.Println(guiEigenKarteikartenContent.AndereKarteiKaesten)
	MainContent.Content = guiEigenKarteikartenContent
	//fmt.Println(MainContent)
	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/meineKarteikaesten.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", model.Main{})
	}
}
