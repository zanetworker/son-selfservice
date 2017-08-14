package handlers

import (
	"fmt"

	r "github.com/GoRethink/gorethink"
	"github.com/prometheus/common/log"

	"github.com/zanetworker/son-selfservice/selfservice-backend/models"
)

//AddFSM adding an FSM
func AddFSM(dbConnection *r.Session, fsmToInsert models.FSM) error {
	log.Info("iDDD", fsmToInsert.ID)
	res, err := r.DB("fsms").Table("fsm_psa").Filter(map[string]interface{}{
		"fsmId": fsmToInsert.FsmID,
	}).Run(dbConnection)
	if err != nil {
		fmt.Print(err)
		return err
	}
	if res.IsNil() {
		_, err := r.DB("fsms").Table("fsm_psa").Insert(fsmToInsert).RunWrite(dbConnection)
		if err != nil {
			log.Errorf("Failed to inser FSM: %s", err.Error())
			return err
		}
		log.Infof("FSM %s was already added to the database", fsmToInsert.FsmID)
		return nil
	}
	log.Infof("FSM %s was already added to the database", fsmToInsert.FsmID)
	return nil
}
