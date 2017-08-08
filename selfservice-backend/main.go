package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/zanetworker/son-selfservice/selfservice-backend/models"
)

//Upgrader to switch protocols
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	subscribeToUpdates(conn)

	if err != nil {
		log.Error(err)
		return
	}

	for {
		var actionMessage models.Message
		var outMessage models.Message
		if err = conn.ReadJSON(&actionMessage); err != nil {
			log.Error(err)
			break
		}

		switch actionMessage.Name {
		case "fsm start":
			log.Info("Starting FSM")
			if err = startFSM(actionMessage.Data); err != nil {
				outMessage = models.Message{
					Name: "fsm start",
					Data: models.ErrorMessage{
						Type:  "error",
						Value: err.Error(),
					},
				}
				if err = conn.WriteJSON(outMessage); err != nil {
					log.Error(err)
					break
				}
			}
			if err = conn.WriteJSON(outMessage); err != nil {
				log.Error(err)
				break
			}
		case "fsm stop":
			//TODO add code to stop FSM
		}
	}
}

func startFSM(fsmInputData interface{}) error {
	var fsmData models.FSMAction
	if err := mapstructure.Decode(fsmInputData, &fsmData); err != nil {
		return err
	}
	log.Infof("%#v\n", fsmData)
	return nil
}

func stopFSM(fsmInputData interface{}) error {
	//TODO
	return nil
}

func subscribeToUpdates(socket *websocket.Conn) {
	//TODO: rethinkDB Query / changefeed
	for {
		time.Sleep(time.Second * 1)
		messageToSend := models.Message{
			Name: "fsm update",
			Data: models.FSMStatusUpdate{
				ID:    "12345",
				State: "started",
			},
		}
		if err := socket.WriteJSON(messageToSend); err != nil {
			log.Error(err.Error())
		}
	}
}

func main() {
	http.HandleFunc("/wsecho", handler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
