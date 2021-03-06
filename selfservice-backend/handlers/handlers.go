package handlers

import (
	"fmt"
	"net/url"
	"reflect"

	r "github.com/GoRethink/gorethink"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/zanetworker/son-selfservice/selfservice-backend/communication"
	"github.com/zanetworker/son-selfservice/selfservice-backend/models"
)

var (
	//URL used for connecting to ssm
	// URL = url.URL{Scheme: "ws", Host: "localhost:9191", Path: "/echo"}
	URL = url.URL{Scheme: "ws", Host: "selfservice-ssm:9191", Path: "/echo"}
)

//StartServiceBasic command to send the SSM for starting the basic tier service
func StartServiceBasic(client *communication.Client, serviceInputData interface{}) {
	log.Infof("Starting Service Basic: %#v\n", serviceInputData)

	var serviceRequest models.Message
	var serviceReply models.Message

	serviceRequest.Name = "basic start"
	serviceReply.Data = ""

	//Initiate the websocket connection
	//u := url.URL{Scheme: "ws", Host: "selfservice-ssm:9191", Path: "/echo"}
	// u := url.URL{Scheme: "ws", Host: "localhost:9191", Path: "/echo"}

	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Error("Failed to Connect to SSM!")
	}
	defer c.Close()

	err = c.WriteJSON(serviceRequest)
	if err != nil {
		log.Println("write:", err)
		return
	}

	log.Info("Reading Reply...!")
	err = c.ReadJSON(&serviceReply)
	if err != nil {
		log.Println("Read:", err)
		return
	}

	log.Info(serviceReply)
	client.Send <- serviceReply
	log.Info("Started Service")
}

//StopServiceBasic command to send the SSM for starting the basic tier service
func StopServiceBasic(client *communication.Client, serviceInputData interface{}) {
	log.Infof("Starting Service Basic: %#v\n", serviceInputData)

	var serviceRequest models.Message
	var serviceReply models.Message

	serviceRequest.Name = "basic stop"
	serviceReply.Data = ""

	//Initiate the websocket connection
	//u := url.URL{Scheme: "ws", Host: "selfservice-ssm:9191", Path: "/echo"}
	// u := url.URL{Scheme: "ws", Host: "localhost:9191", Path: "/echo"}

	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Error("Failed to Connect to SSM!")
	}
	defer c.Close()

	err = c.WriteJSON(serviceRequest)
	if err != nil {
		log.Println("write:", err)
		return
	}

	log.Info("Reading Reply...!")
	err = c.ReadJSON(&serviceReply)
	if err != nil {
		log.Println("Read:", err)
		return
	}

	log.Info(serviceReply)
	client.Send <- serviceReply
	log.Info("Started Service")
}

//StartServiceAnon command to send the SSM for starting the Anon tier service
func StartServiceAnon(client *communication.Client, serviceInputData interface{}) {
	log.Infof("Starting Service Anon: %#v\n", serviceInputData)

	var serviceRequest models.Message
	var serviceReply models.Message

	serviceRequest.Name = "anon start"
	serviceReply.Data = ""

	//Initiate the websocket connection
	//u := url.URL{Scheme: "ws", Host: "selfservice-ssm:9191", Path: "/echo"}
	// u := url.URL{Scheme: "ws", Host: "localhost:9191", Path: "/echo"}

	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Error("Failed to Connect to SSM!")
	}
	defer c.Close()

	err = c.WriteJSON(serviceRequest)
	if err != nil {
		log.Println("write:", err)
		return
	}

	log.Info("Reading Reply...!")
	err = c.ReadJSON(&serviceReply)
	if err != nil {
		log.Println("Read:", err)
		return
	}

	log.Info(serviceReply)
	client.Send <- serviceReply
	log.Info("Started Service")
}

//StopServiceAnon command to send the SSM for starting the Anon tier service
func StopServiceAnon(client *communication.Client, serviceInputData interface{}) {
	log.Infof("Starting Service Anon: %#v\n", serviceInputData)

	var serviceRequest models.Message
	var serviceReply models.Message

	serviceRequest.Name = "anon stop"
	serviceReply.Data = ""

	//Initiate the websocket connection
	//u := url.URL{Scheme: "ws", Host: "selfservice-ssm:9191", Path: "/echo"}
	// u := url.URL{Scheme: "ws", Host: "localhost:9191", Path: "/echo"}

	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Error("Failed to Connect to SSM!")
	}
	defer c.Close()

	err = c.WriteJSON(serviceRequest)
	if err != nil {
		log.Println("write:", err)
		return
	}

	log.Info("Reading Reply...!")
	err = c.ReadJSON(&serviceReply)
	if err != nil {
		log.Println("Read:", err)
		return
	}

	log.Info(serviceReply)
	client.Send <- serviceReply
	log.Info("Started Service")
}

//StartFSM command to send the SSM to start a specific FSM
func StartFSM(client *communication.Client, fsmInputData interface{}) {
	var fsmData models.FSMAction
	var fsmDataRequest models.Message
	var fsmDataReply models.Message

	if err := mapstructure.Decode(fsmInputData, &fsmData); err != nil {
		log.Error(err)
	}

	fsmDataRequest.Name = "fsm start"
	fsmDataRequest.Data = fsmData

	//Initiate the websocket connection
	//u := url.URL{Scheme: "ws", Host: "selfservice-ssm:9191", Path: "/echo"}
	// u := url.URL{Scheme: "ws", Host: "localhost:9191", Path: "/echo"}

	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Error("Failed to Connect to SSM!")
	}
	defer c.Close()

	err = c.WriteJSON(fsmDataRequest)
	if err != nil {
		log.Println("write:", err)
		return
	}

	log.Info("Reading Reply...!")
	err = c.ReadJSON(&fsmDataReply)
	if err != nil {
		log.Println("Read:", err)
		return
	}

	log.Info(fsmDataReply)
	client.Send <- fsmDataReply
	log.Info("Started FSM")
}

//StopFSM command to send the SSM to stop a specific FSM
func StopFSM(client *communication.Client, fsmInputData interface{}) {
	var fsmData models.FSMAction
	var fsmDataRequest models.Message
	var fsmDataReply models.Message

	if err := mapstructure.Decode(fsmInputData, &fsmData); err != nil {
		log.Error(err)
	}

	fsmDataRequest.Name = "fsm stop"
	fsmDataRequest.Data = fsmData

	//Initiate the websocket connection
	// u := url.URL{Scheme: "ws", Host: "selfservice-ssm:1234", Path: "/echo"}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Error("Failded to Connect to SSM!")
	}
	defer c.Close()

	err = c.WriteJSON(fsmDataRequest)
	if err != nil {
		log.Println("write:", err)
		return
	}

	log.Info("Reading Reply...!")
	err = c.ReadJSON(&fsmDataReply)
	if err != nil {
		log.Println("Read:", err)
		return
	}

	log.Info(fsmDataReply)
	client.Send <- fsmDataReply
	log.Info("Stopped FSM")
}

//AddFSMs adds FSMs if they don't exists
func AddFSMs(client *communication.Client, fsmsInputData interface{}) {
	log.Info("Adding FSMs")
	var fsms models.FSMs

	switch reflect.TypeOf(fsmsInputData).Kind() {
	case reflect.Slice:
		log.Info("Its a slice")
		fsmsInputDataValue := reflect.ValueOf(fsmsInputData)
		// log.Infof("%#v\n", fsmsInputDataValue)

		// if err := mapstructure.Decode(fsmsInputDataValue, &fsms); err != nil {
		// 	log.Error(err.Error())
		// 	return
		// }

		test := fsmsInputDataValue.Interface().([]interface{})
		for _, vlaue := range test {
			var fsm models.FSM
			err := mapstructure.Decode(vlaue, &fsm)
			if err != nil {
				log.Error(err.Error())
			}
			fsms.FSMList = append(fsms.FSMList, fsm)
		}

		for _, fsm := range fsms.FSMList {
			err := AddFSM(client.DB.Connection, fsm)
			if err != nil {
				log.Error(err.Error())
			}
		}
	}

}

// //AddFSM adding an FSM
// func AddFSM(client *commclient.Client, fsmInputData interface{}) {
//
// 	var fsmToInsert models.FSM
// 	err := mapstructure.Decode(fsmInputData, &fsmToInsert)
// 	if err != nil {
// 		log.Error("Failed to map structure FSM")
// 	} else {
// 		response, err := r.DB("fsms").Table("fsm_psa").Insert(fsmToInsert).RunWrite(client.DB.Connection)
// 		if err != nil {
// 			log.Errorf("Failed to inser FSM: %s", err.Error())
// 			return
// 		}
// 		log.Info(response)
// 	}
//
// }

//UpdateFSM used to update the database with FSMs information received from SSM
func UpdateFSM(client *communication.Client, fsmInputData interface{}) {
	res, err := r.DB("fsms").Table("fsm_psa").Filter(map[string]interface{}{
		"name": "test",
	}).Run(client.DB.Connection)
	if err != nil {
		fmt.Print(err)
		return
	}
	// Scan query result into the person variable
	var results []interface{}
	err = res.All(&results)
	if len(results) == 0 {
		log.Infof("FSM does not exist in the database")
	}
	var fsmResults models.FSM
	err = mapstructure.Decode(results[0], &fsmResults)
	if err != nil {
		log.Error("Failed to find fsm", err.Error())
		return
	}
	log.Infof("%#v\n", fsmResults.ID)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
		return
	}
	var updateFSMwithValues models.FSM
	err = mapstructure.Decode(fsmInputData, &updateFSMwithValues)
	if err != nil {
		log.Error("Failed to map fsm to inputdata", err.Error())
	}
	resp, err := r.DB("fsms").Table("fsm_psa").Get(fsmResults.ID).Update(fsmInputData).RunWrite(client.DB.Connection)
	log.Infof("%#v\n", resp)
	if err != nil {
		fmt.Print(err)
		return
	}
	// log.Infof("%d users", len(results))

}
