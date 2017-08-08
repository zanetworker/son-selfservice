package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
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
	rules map[string]Handler
}

//NewRouter Init for router
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
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
	client := NewClient(conn, router.FindHandler)
	go client.Write()
	go client.SubscribeToUpdates()
	client.Read()
}
