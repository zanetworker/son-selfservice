package communication

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zanetworker/son-selfservice/selfservice-backend/database"
)

//Handler the function handler type
type Handler func(*Client, interface{})

//Upgrader to switch protocols
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//Router is the handler we will use
type Router struct {
	rules    map[string]Handler
	database *database.Database
}

//NewRouter Init for router
func NewRouter(db *database.Database) *Router {
	return &Router{
		rules:    make(map[string]Handler),
		database: db,
	}
}

//FindHandler a method to find the right Handler
func (router *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := router.rules[msgName]
	return handler, found
}

//Handle maps commands to functions to execute
func (router *Router) Handle(msgName string, handler Handler) {
	router.rules[msgName] = handler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	client := NewClient(conn, router.FindHandler, router.database)
	// router.database.DeleteAll("fsms", "fsm_psa")
	go client.Write()
	go client.SubscribeToUpdates()
	go client.Read()
}
