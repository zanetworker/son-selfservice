package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"github.com/zanetworker/son-selfservice/selfservice-backend/models"
)

//StartFSM command to send the SSM to start a specific FSM
func StartFSM(client *Client, fsmInputData interface{}) {
	var fsmData models.FSMAction
	var fsmDataReply models.Message
	if err := mapstructure.Decode(fsmInputData, &fsmData); err != nil {
		log.Error(err)
	}

	//TODO send command to the SSM
	fsmDataReply.Name = "fsm start"
	fsmDataReply.Data = fsmData
	client.send <- fsmDataReply

	log.Infof("%#v\n", fsmData)
	log.Info("started FSM")
}

//StopFSM command to send the SSM to stop a specific FSM
func StopFSM(client *Client, fsmInputData interface{}) {
	var fsmData models.FSMAction
	var fsmDataReply models.Message
	if err := mapstructure.Decode(fsmInputData, &fsmData); err != nil {
		log.Error(err)
	}

	//TODO send command to the SSM
	fsmDataReply.Name = "fsm stop"
	fsmDataReply.Data = fsmData
	client.send <- fsmDataReply

	log.Infof("%#v\n", fsmData)
	log.Info("started FSM")
}

//UpdateFSM used to update the database with FSMs information received from SSM
func UpdateFSM(client *Client, fsmInputData interface{}) {

}
