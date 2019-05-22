package model

//Main data structure
type Main struct {
	LoggedIn bool
	Username string
	Karteikaesten []map[string]interface {}
	Content interface{}
}
