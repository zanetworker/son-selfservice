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
	http.Handle("/frontend", router)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
