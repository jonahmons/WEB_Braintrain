package model

import (
	"fmt"
)

//Karteikarte data struct
type Karteikarte struct {
	ID               string `json:"_id"`
	Rev              string `json:"_rev"`
	Type             string `json:"type"`
	KarteikartenName string `json:"KarteikartenName"`
	Frage            string `json:"frage"`
	Antwort          string `json:"antwort"`
	Kasten           string `json:"kasten"`
}

// Add Karteikarten to DB
func (k Karteikarte) Add() error {
	// Convert Karteikarten struct to map[string]interface as required by Save() method
	Karteikarten := doMagic(k)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(Karteikarten, "_id")
	delete(Karteikarten, "_rev")

	// Add Karteikarten to DB
	_, _, err := btDB.Save(Karteikarten, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

// GetKarteikarte with the provided id from DB
func GetKarteikarte(id string) (Karteikarte, error) {
	k, err := btDB.Get(id, nil)
	if err != nil {
		return Karteikarte{}, err
	}
	return fuelleKareteikarte(k), nil
}

// GetAllKarteikarten DB
func GetAllKarteikarten() ([]Karteikarte, error) {
	k, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "karteikarte"s
			 }
		}
	 }`)
	allKateikarten := []Karteikarte{}
	for _, temp := range k {
		allKateikarten = append(allKateikarten, fuelleKareteikarte(temp))
	}
	if err != nil {
		return nil, err
	}

	return allKateikarten, nil
}

// GetAllKarteikartenByKasten DB
func GetAllKarteikartenByKasten(id string) ([]Karteikarte, error) {
	k, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "karteikarte"
			 },
			 "kasten": {
				 "$eq": "`+ id +`"
			}
		}
	 }`)
	 //fmt.Printf("%+v\n", k)
	allKateikarten := []Karteikarte{}
	for _, temp := range k {
		allKateikarten = append(allKateikarten, fuelleKareteikarte(temp))
	}
	if err != nil {
		return nil, err
	}

	return allKateikarten, nil
}

// UpdateKarteikarte in DB
func (k Karteikarte) UpdateKarteikarte() error {

	err := btDB.Set(k.ID, doMagic(k))
	if err != nil {
		fmt.Printf("[Update] error: %s", err)
	}
	return err
}

// DeleteKarteikarte with the provided id from DB
func (k Karteikarte) DeleteKarteikarte() error {
	err := btDB.Delete(k.ID)

	return err
}

//fuelleKareteikarte map in Karteikarte
func fuelleKareteikarte(i map[string]interface{}) Karteikarte {
	karteikarte := Karteikarte{
		ID:               i["_id"].(string),
		Rev:              i["_rev"].(string),
		Type:             i["type"].(string),
		KarteikartenName: i["KarteikartenName"].(string),
		Frage:            i["frage"].(string),
		Antwort:          i["antwort"].(string),
		Kasten:           i["kasten"].(string),
	}
	return karteikarte
}

//CountKarteikarten
func CountKarteikarten() int {

	count, err := btDB.QueryJSON(`
{
"selector": {
"type": {
"$eq": "karteikarte"
}
},
"fields": ["_id"]
}`)

	if err != nil {
		fmt.Printf("[Karteikarte][CountKarteikarten] error: %s", err)
		return -1
	}

	return len(count)
}

func CountKarteikartenInKateikasten(box string) int {
	return countCards(`
		{
			"selector": {
				 "type": {
						"$eq": "karteikarte"
				 },
				 "box_id": "` + box + `"
			},
			"fields": ["_id"]
		 }`)
}
func CountCards() int {
	return countCards(`
		{
			"selector": {
				 "type": {
						"$eq": "karteikarte"
				 }
			},
			"fields": ["_id"]
		 }`)
}

//CountCards
func countCards(query string) int {

	count, err := btDB.QueryJSON(query)

	if err != nil {
		fmt.Printf("[User][COuntCards] error: %s", err)
		return -1
	}

	return len(count)
}
