package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/zanetworker/son-selfservice/selfservice-backend/database"
)

func main() {
	db := database.NewDB()
	router := NewRouter(db)

	router.Handle("fsm start", StartFSM)
	router.Handle("fsm add", AddFSM)
	router.Handle("fsm update", UpdateFSM)
	router.Handle("fsm stop", StopFSM)

	http.Handle("/ws", router)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
