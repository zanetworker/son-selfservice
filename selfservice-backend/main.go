package main

import (
	"log"
	"net/http"

	"github.com/zanetworker/son-selfservice/selfservice-backend/communication"
	"github.com/zanetworker/son-selfservice/selfservice-backend/configuration"
	"github.com/zanetworker/son-selfservice/selfservice-backend/database"
	"github.com/zanetworker/son-selfservice/selfservice-backend/handlers"
)

func main() {
	configuration.LoadAppConfig()

	db := database.NewDB()
	router := communication.NewRouter(db)

	router.Handle("fsm start", handlers.StartFSM)
	router.Handle("fsm add", handlers.AddFSMs)
	router.Handle("fsm update", handlers.UpdateFSM)
	router.Handle("fsm stop", handlers.StopFSM)

	http.Handle("/ws", router)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
