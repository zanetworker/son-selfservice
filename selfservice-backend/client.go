package main

import (
	"math/rand"
	"time"

	"github.com/gorilla/websocket"

	log "github.com/Sirupsen/logrus"
	"github.com/zanetworker/son-selfservice/selfservice-backend/models"
)

//FindHandler to find the right handler for the command
type FindHandler func(string) (Handler, bool)

//The Client Struct to communicate with the FE
type Client struct {
	socket      *websocket.Conn
	send        chan models.Message
	findHandler FindHandler
}

func (client *Client) Write() {
	defer client.socket.Close()
	for msgToSend := range client.send {
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

		command := actionMessage.Name
		if handler, found := client.findHandler(command); found {
			handler(client, actionMessage.Data)
		}
	}

}

//SubscribeToUpdates test client
func (client *Client) SubscribeToUpdates() {
	for {
		time.Sleep(time.Second * 1)
		states := []string{"started", "stopped"}
		access := rand.Intn(len(states))
		messageToSend := models.Message{
			Name: "fsm update",
			Data: models.FSMStatusUpdate{
				ID:    "12345",
				State: states[access],
			},
		}
		client.send <- messageToSend
	}
}

//NewClient Instantiate a new Client
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		socket:      socket,
		send:        make(chan models.Message),
		findHandler: findHandler,
	}
}
