package communication

import (
	r "github.com/GoRethink/gorethink"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/zanetworker/son-selfservice/selfservice-backend/database"
	"github.com/zanetworker/son-selfservice/selfservice-backend/models"
)

//FindHandler to find the right handler for the command
type FindHandler func(string) (Handler, bool)

//The Client Struct to communicate with the FE
type Client struct {
	socket      *websocket.Conn
	DB          *database.Database
	Send        chan models.Message
	findHandler FindHandler
}

func (client *Client) Write() {
	defer client.socket.Close()
	for msgToSend := range client.Send {
		if err := client.socket.WriteJSON(msgToSend); err != nil {
			log.Error(err.Error())
			break
		}
		log.Infof("%#v\n", msgToSend)
	}
}

func (client *Client) Read() {
	defer client.socket.Close()

	for {
		var actionMessage models.Message
		if err := client.socket.ReadJSON(&actionMessage); err != nil {
			log.Error(err.Error())
			break
		}
		log.Info("Action Message")
		log.Infof("%#v\n", actionMessage)
		command := actionMessage.Name
		if handler, found := client.findHandler(command); found {
			handler(client, actionMessage.Data)
		}
	}

}

//SubscribeToUpdates test client
func (client *Client) SubscribeToUpdates() {
	log.Info("Waiting for Database Updates")
	cursor, err := r.DB("fsms").Table("fsm_psa").Changes(r.ChangesOpts{
		IncludeInitial: true,
	}).Run(client.DB.Connection)

	if err != nil {
		log.Error(err.Error())
		return
	}
	var changeResponse r.ChangeResponse
	for cursor.Next(&changeResponse) {
		var newFSMValue models.FSM
		if err := mapstructure.Decode(changeResponse.NewValue, &newFSMValue); err != nil {
			log.Error(err.Error())
			continue
		}
		// states := []string{"started", "stopped"}
		// access := rand.Intn(len(states))
		messageToSend := models.Message{
			Name: "fsm update",
			Data: models.FSMStatusUpdate{
				FsmID: newFSMValue.FsmID,
				State: newFSMValue.State,
			},
		}
		client.Send <- messageToSend
	}

}

//NewClient Instantiate a new Client
func NewClient(socket *websocket.Conn, findHandler FindHandler, db *database.Database) *Client {
	return &Client{
		socket:      socket,
		Send:        make(chan models.Message),
		findHandler: findHandler,
		DB:          db,
	}
}
