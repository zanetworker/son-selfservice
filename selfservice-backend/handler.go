package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"github.com/zanetworker/son-selfservice/selfservice-backend/models"
)

//StartFSM command to send the SSM to start a specific FSM
func StartFSM(client *Client, fsmInputData interface{}) {
	var fsmData models.FSMAction
	if err := mapstructure.Decode(fsmInputData, &fsmData); err != nil {
		log.Error(err)
	}

	//TODO send command to the SSM
	log.Infof("%#v\n", fsmData)
	log.Info("started FSM")
}

//StopFSM command to send the SSM to stop a specific FSM
func StopFSM(fsmInputData interface{}) error {
	//TODO
	return nil
}
