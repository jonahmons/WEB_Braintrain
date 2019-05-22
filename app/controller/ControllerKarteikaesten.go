package controller

import (
	"05Prak/app/model"
	//"fmt"
	"html/template"
	"net/http"
)


//Karteikaesten wireframe
func Karteikaesten(w http.ResponseWriter, r *http.Request) {

	karteikaesten, _ := model.GetAllKarteikaesten()
	//fmt.Println(karteikaesten)
	kategorien := map[int]model.GUIKarteikastenKategorien{}

	for _, kk := range karteikaesten {
		GUIkarteikasten := model.GUIKarteikasten{
			ID:             kk.ID,
			Unterkategorie: model.GetCategoryByID(kk.Kategorie),
			Name:           kk.KarteikastenName,
			AnzahlKarten:   model.CountKarteikartenInKateikasten(kk.ID),
			Beschreibung:   kk.Beschreibung,
		}

		id, name := model.GetUpperCategoryByCatID(kk.Kategorie)

		if _, ok := kategorien[id]; !ok {
			kategorien[id] = model.GUIKarteikastenKategorien{
				Name:          name,
				Karteikaesten: map[string]model.GUIKarteikasten{},
			}
		}
		kategorien[id].Karteikaesten[kk.ID] = GUIkarteikasten
}
		GUIkarteikasten := model.GUIKarteikastenContent{
			Kategorien: kategorien,
		}
		MainContent.Content = GUIkarteikasten
		//fmt.Println(MainContent)
	if MainContent.LoggedIn {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/topLoggedIn.tmpl", "template/navbarLoggedIn.tmpl", "template/karteikaesten.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	} else {
		t, _ := template.ParseFiles("template/layout.tmpl", "template/top.tmpl", "template/navbar.tmpl", "template/startpage.tmpl")
		t.ExecuteTemplate(w, "layout", MainContent)
	}

}
