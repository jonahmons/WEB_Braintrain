package model

import (
	"fmt"
)

//User data struct
type User struct {
	ID       string `json:"_id"`
	Rev      string `json:"_rev"`
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Date     string `json:"date"`
}

// Add User to DB
func (u User) Add() error {
	// Convert User struct to map[string]interface as required by Save() method
	user := doMagic(u)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(user, "_id")
	delete(user, "_rev")

	// Add user to DB
	_, _, err := btDB.Save(user, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

// GetUser with the provided id from DB
func GetUser(id string) (User, error) {
	t, err := btDB.Get(id, nil)
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:       t["_id"].(string),
		Rev:      t["_rev"].(string),
		Type:     t["type"].(string),
		Username: t["username"].(string),
		Password: t["password"].(string),
		Email:    t["email"].(string),
		Date:     t["date"].(string),
	}
	return user, nil
}

// GetAllUsers from DB
func GetAllUsers() ([]map[string]interface{}, error) {
	allUsers, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "user"
			 }
		}
	 }`)
	if err != nil {
		return nil, err
	}
	return allUsers, nil

}

// UpdateUser in DB
func (u User) UpdateUser() error {

	err := btDB.Set(u.ID, doMagic(u))
	if err != nil {
		fmt.Printf("[Update] error: %s", err)
	}
	return err
}

// Delete User with the provided id from DB
func (u User) Delete() error {
	err := btDB.Delete(u.ID)

	return err
}

//CountUsers Funktion
func CountUsers() int {

	count, err := btDB.QueryJSON(`
{
"selector": {
"type": {
"$eq": "user"
}
},
"fields": ["_id"]
}`)

	if err != nil {
		fmt.Printf("[User][CountUsers] error: %s", err)
		return -1
	}

	return len(count)
}
