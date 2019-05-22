package model

import (
	"encoding/json"
	couchdb "github.com/leesper/couchdb-golang"
)

// Db handle
var btDB *couchdb.Database

func init() {
	var err error
	btDB, err = couchdb.NewDatabase("http://localhost:5984/braintrain")
	if err != nil {
		panic(err)
	}
}

// Convert from struct to map[string]interface{} as required by Set() method
func doMagic(t interface{}) map[string]interface{} {
	var doc map[string]interface{}
	ret, _ := json.Marshal(t)
	json.Unmarshal(ret, &doc)

	return doc
}
