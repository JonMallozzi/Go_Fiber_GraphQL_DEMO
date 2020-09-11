package postgres

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/go-pg/pg/v10"
)

func PostgresConnect() *pg.DB {
	//connecting to the db with orm using hidden JSON parameters
	file, _ := ioutil.ReadFile("./dbConnection.json")
	data := make(map[string]interface{})
	err := json.Unmarshal(file, &data)

	if err != nil {
		log.Fatal("Cannot unmarshal the json ", err)
	}

	db := pg.Connect(&pg.Options{
		Addr:     data["Addr"].(string),
		User:     data["User"].(string),
		Password: data["Password"].(string),
		Database: data["Database"].(string),
	})

	return db
}
