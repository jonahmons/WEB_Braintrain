package model

import (
	"fmt"
)

//Fortschrittkasten data struct
type Fortschrittkasten struct {
	ID                string             `json:"_id"`
	Rev               string             `json:"_rev"`
	Type              string             `json:"type"`
	User              string             `json:"user"`
	Kasten            string             `json:"kasten"`
	Fortschrittgesamt string             `json:"fortschrittgesamt"`
	Fortschrittkarten []Fortschrittkarte `json:"fortschrittkarten"`
}
//Fortschrittkarte struct
type Fortschrittkarte struct {
	Nr    string `json:"nr"`
	Phase string `json:"phase"`
}

// Add Fortschrittkasten to DB
func (f Fortschrittkasten) Add() error {
	// Convert Fortschrittkasten struct to map[string]interface as required by Save() method
	fortschrittkasten := doMagic(f)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(fortschrittkasten, "_id")
	delete(fortschrittkasten, "_rev")

	// Add Fortschrittkasten to DB
	_, _, err := btDB.Save(fortschrittkasten, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

// GetFortschrittkasten with the provided id from DB
func GetFortschrittkasten(id string) (Fortschrittkasten, error) {
	f, err := btDB.Get(id, nil)
	if err != nil {
		return Fortschrittkasten{}, err
	}

	fkarten := []Fortschrittkarte{}

	for value := range doMagic(f["fortschrittkarten"]) {
		fk := doMagic(value)
		fkarten = append(fkarten, Fortschrittkarte{
			Nr:    fk["nr"].(string),
			Phase: fk["phase"].(string),
		})
	}

	Fortschrittkasten := Fortschrittkasten{
		ID:                f["_id"].(string),
		Rev:               f["_rev"].(string),
		Type:              f["type"].(string),
		User:              f["user"].(string),
		Kasten:            f["kasten"].(string),
		Fortschrittgesamt: f["fortschrittgesamt"].(string),
		Fortschrittkarten: fkarten,
	}
	return Fortschrittkasten, nil
}

// GetAllFortschrittkasten from DB
func GetAllFortschrittkasten() ([]map[string]interface{}, error) {
	allFortschrittkasten, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "fortschrittkasten"
			 }
		}
	 }`)
	if err != nil {
		return nil, err
	}
	return allFortschrittkasten, nil

}
// GetFortschrittkastenByKastenIDAndUserID function
func GetFortschrittkastenByKastenIDAndUserID(kastenID string, userID string) (Fortschrittkasten, error) {
	f, err := btDB.QueryJSON(`
	{
		"selector": {
			 "type": {
					"$eq": "fortschrittkasten"
			 },
			 "user": {
				 "$eq": "` + userID + `"
			},
			"kasten": {
				"$eq": "` + kastenID + `"
		 }
		}
	 }`)

	alleFortschritte := Fortschrittkasten{}
	for _, temp := range f {
		alleFortschritte = fuelleFortschritt(temp)
	}
	if err != nil {
		return alleFortschritte, err
	}
	return alleFortschritte, nil
}

func fuelleFortschritt(i map[string]interface{}) Fortschrittkasten {
	fkarten := []Fortschrittkarte{}
	for value := range doMagic(i["fortschrittkarten"]) {
		fk := doMagic(value)
		fkarten = append(fkarten, Fortschrittkarte{
			Nr:    fk["nr"].(string),
			Phase: fk["phase"].(string),
		})
	}
	fortschrittkasten := Fortschrittkasten{
		ID:                i["_id"].(string),
		Rev:               i["_rev"].(string),
		Type:              i["type"].(string),
		User:              i["user"].(string),
		Kasten:            i["kasten"].(string),
		Fortschrittgesamt: i["fortschrittgesamt"].(string),
		Fortschrittkarten: fkarten,
	}
	return fortschrittkasten
}

// UpdateFortschrittkasten in DB
func (f Fortschrittkasten) UpdateFortschrittkasten() error {

	err := btDB.Set(f.ID, doMagic(f))
	if err != nil {
		fmt.Printf("[Update] error: %s", err)
	}
	return err
}

// DeleteFortschrittkasten with the provided id from DB
func (f Fortschrittkasten) DeleteFortschrittkasten() error {
	err := btDB.Delete(f.ID)

	return err
}
