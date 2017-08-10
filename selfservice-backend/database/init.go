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

//SubscribeToChanges asdasd
func (db *Database) SubscribeToChanges(dbName, tableName string) {
	log.Info("Waiting for Database Updates")
	cursor, err := r.DB(dbName).Table(tableName).Changes(r.ChangesOpts{
		IncludeInitial: true,
	}).Run(db.Connection)

	if err != nil {
		log.Error(err.Error())
		return
	}
	var changeResponse r.ChangeResponse
	for cursor.Next(&changeResponse) {
		log.Infof("%#v\n", changeResponse)
		db.DBChannel <- changeResponse
	}

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
