package database

import (
	r "github.com/GoRethink/gorethink"
	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"github.com/zanetworker/son-selfservice/selfservice-backend/models"
)

//Database is the struct for connecting to the database
type Database struct {
	Connection *r.Session
	DBChannel  chan r.ChangeResponse
}

//DeleteAll deletes all entries in the datq
func (db *Database) DeleteAll(dbName, tableName string) {
	r.DB(dbName).Table("fsm_psa").Delete().Run(db.Connection)
}

//AddFSM asdasd
func (db *Database) AddFSM(dbName, tableName string, dataToAdd interface{}) error {
	var fsmToAdd models.FSM
	if err := mapstructure.Decode(dataToAdd, &fsmToAdd); err != nil {
		log.Error(err.Error())
		return err
	}
	_, err := r.DB(dbName).Table(tableName).Insert(dataToAdd).RunWrite(db.Connection)
	if err != nil {
		return err
	}

	return nil
}

//NewDB initialize DB
func NewDB() *Database {
	session, err := r.Connect(r.ConnectOpts{
		//TODO fetch from Config File
		Address:  "localhost:28015",
		Database: "fsms",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return &Database{
		Connection: session,
		DBChannel:  make(chan r.ChangeResponse),
	}
}
