package model

//GUIKarteikasten data struct
type GUIKarteikasten struct {
	ID             string
	Oberkategorie  string
	Unterkategorie string
	Name           string
	AnzahlKarten   int
	Beschreibung   string
	Ersteller      string
}

//GUIEigenKarteikasten data struct
type GUIEigenKarteikasten struct {
	ID             string
	Kategorie      string
	Unterkategorie string
	Name           string
	AnzahlKarten   int
	Beschreibung   string
	Sichtbarkeit   string
	Fortschritt    int
}

//GUIStartpage struct
type GUIStartpage struct {
	Nutzer     int
	Lernkarten int
	Karteien   int
}

//GUIKarteikastenKategorien struct
type GUIKarteikastenKategorien struct {
	Name          string
	Karteikaesten map[string]GUIKarteikasten
}

//GUIKarteikastenContent struct
type GUIKarteikastenContent struct {
	Kategorien map[int]GUIKarteikastenKategorien
}

//GUIEigenKarteikastenKategorien struct
type GUIEigenKarteikastenKategorien struct {
	Name          string
	Karteikaesten map[string]GUIEigenKarteikasten
}

//GUIEigenKarteikastenContent struct
type GUIEigenKarteikastenContent struct {
	EigenKarteiKaesten  map[string]GUIEigenKarteikasten
	AndereKarteiKaesten map[string]GUIEigenKarteikasten
}

//GUIKarteikarten struct
type GUIKarteikarten struct {
	Position         int
	ID               string
	KarteikartenName string
	Frage            string
	Antwort          string
	KastenID         string
}

//GUIKarteikartenContent struct
type GUIKarteikartenContent struct {
	Karteikarten        map[string]GUIKarteikarten
	KarteikastenInfo    map[string]GUIKarteikasten
	EinzelneKarteikarte map[string]GUIEinzelneKarteikarte
	Fortschritt         map[string]GUIFortschritt
}

//GUILernenContent struct
type GUILernenContent struct {
	Karteikarte      map[string]GUIEinzelneKarteikarte
	KarteikastenInfo map[string]GUIKarteikasten
}

//GUIEinzelneKarteikarte struct
type GUIEinzelneKarteikarte struct {
	ID               string
	KarteikartenName string
	Frage            string
	Antwort          string
	KastenID         string
}

//GUIFortschritt struct
type GUIFortschritt struct {
	ID                string
	User              string
	Kasten            string
	Fortschrittgesamt string
	Fortschrittkarten []GUIFortschrittKarte
}

//GUIFortschrittKarte struct
type GUIFortschrittKarte struct {
	Nr    string
	Phase string
}
