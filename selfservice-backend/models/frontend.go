package models

//Message is the message for FrontEnd Actions
type Message struct {
	Name string      `json:"name"` //Name of the command e.g., fsm start/stop
	Data interface{} //Info about the FSM
}

//ReplyMessage is the message for FrontEnd Actions
type ReplyMessage struct {
	Name string      `json:"name"` //Name of the command e.g., fsm start/stop
	Data interface{} `json:"Data"` //Info about the FSM
}

//ServiceAction is the message received to do the action
type ServiceAction struct {
	Name string      `json:"name"`
	Data interface{} `json:"Data"`
}

//FSMAction Information about FSM to start/stop/configure
type FSMAction struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//FSMStatusUpdate update message from ssm
type FSMStatusUpdate struct {
	FsmID   string `json:"id"`
	FsmName string `json:"name"`
	State   string `json:"state"`
}

//ErrorMessage Reply about the status of the FMS
type ErrorMessage struct {
	Type  string `json:"id"`
	Value string `json:"value"`
}
