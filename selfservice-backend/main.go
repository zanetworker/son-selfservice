package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func main() {
	router := NewRouter()
	router.Handle("fsm start", StartFSM)
	http.Handle("/wsecho", router)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
