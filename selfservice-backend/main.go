package main

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/zanetworker/selfservice-backend/models"
)

//Upgrader to switch protocols
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Error(err)
			return
		}
		log.Info(messageType, string(msg))

		var actionMessage models.ActionMessage
		if err = json.Unmarshal(msg, &actionMessage); err != nil {
			log.Error(err)
			return
		}

		log.Infof("%#v\n", actionMessage)

		switch actionMessage.Name {
		case "fsm add":
			//TODO contact the communication block to send the command to SSMs
		}
		if err = conn.WriteMessage(messageType, msg); err != nil {
			log.Error(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/wsecho", handler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
