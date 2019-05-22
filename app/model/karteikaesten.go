package model

import (
	//	"encoding/json"
	"fmt"
)

//Karteikasten data struct
type Karteikasten struct {
	ID               string `json:"_id"`
	Rev              string `json:"_rev"`
	Type             string `json:"type"`
	KarteikastenName string `json:"karteikastenName"`
	Beschreibung     string `json:"beschreibung"`
	Kategorie        int    `json:"kategorie"`
	Autor            string `json:"autor"`
	Sichtbarkeit     string `json:"sichtbarkeit"`
}

// Add Karteikasten to DB
func (k Karteikasten) Add() error {
	// Convert Karteikasten struct to map[string]interface as required by Save() method
	Karteikasten := doMagic(k)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(Karteikasten, "_id")
	delete(Karteikasten, "_rev")

	// Add Karteikasten to DB
	_, _, err := btDB.Save(Karteikasten, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

// GetKarteikasten with the provided id from DB
func GetKarteikasten(id string) (Karteikasten, error) {
	k, err := btDB.Get(id, nil)
	if err != nil {
		return Karteikasten{}, err
	}

	return fuelleKasten(k), nil
}

// GetAllKarteikaesten from DB
func GetAllKarteikaesten() ([]Karteikasten, error) {
	k, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "karteikasten"
			 }
		}
	 }`)
	allKarteikaesten := []Karteikasten{}
	for _, temp := range k {
		allKarteikaesten = append(allKarteikaesten, fuelleKasten(temp))
	}
	if err != nil {
		return nil, err
	}
	return allKarteikaesten, nil
}

// GetAllKarteikaestenByUserID from DB
func GetAllKarteikaestenByUserID(id string) ([]Karteikasten, error) {
	k, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "karteikasten"
			 },
			 "autor": {
				 	"$eq": "`+ id +`"
			 }
		}
	 }`)
	allKarteikaesten := []Karteikasten{}
	for _, temp := range k {
		allKarteikaesten = append(allKarteikaesten, fuelleKasten(temp))
	}
	if err != nil {
		return nil, err
	}
	return allKarteikaesten, nil
}

// GetAllPublicKarteikaesten from DB
func GetAllPublicKarteikaesten() ([]Karteikasten, error) {
	k, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "karteikasten"
			 },
			 "sichtbarkeit": {
				 	"$eq": "öffentlich"
			 }
		}
	 }`)
	allKarteikaesten := []Karteikasten{}
	for _, temp := range k {
		allKarteikaesten = append(allKarteikaesten, fuelleKasten(temp))
	}
	if err != nil {
		return nil, err
	}
	return allKarteikaesten, nil
}

// UpdateKarteikasten in DB
func (k Karteikasten) UpdateKarteikasten() error {

	err := btDB.Set(k.ID, doMagic(k))
	if err != nil {
		fmt.Printf("[Update] error: %s", err)
	}
	return err
}

// Delete Karteikasten with the provided id from DB
func (k Karteikasten) Delete() error {
	err := btDB.Delete(k.ID)

	return err
}

//fuelleKasten map in Karteikasten
func fuelleKasten(i map[string]interface{}) Karteikasten {
	karteikasten := Karteikasten{
		ID:               i["_id"].(string),
		Rev:              i["_rev"].(string),
		Type:             i["type"].(string),
		KarteikastenName: i["karteikastenName"].(string),
		Beschreibung:     i["beschreibung"].(string),
		Kategorie:        int(i["kategorie"].(float64)),
		Autor:            i["autor"].(string),
		Sichtbarkeit:     i["sichtbarkeit"].(string),
	}
	return karteikasten
}

//CountKarteikaesten func
func CountKarteikaesten() int {

	count, err := btDB.QueryJSON(`
{
"selector": {
"type": {
"$eq": "karteikasten"
}
},
"fields": ["_id"]
}`)

	if err != nil {
		fmt.Printf("[karteikasten][CountKarteikaesten] error: %s", err)
		return -1
	}

	return len(count)
}
