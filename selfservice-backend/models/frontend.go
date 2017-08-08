package models

//Message is the message for FrontEnd Actions
type Message struct {
	Name string      `json:"name"` //Name of the command e.g., fsm start/stop
	Data interface{} //Info about the FSM
}

//FSMAction Information about FSM to start/stop/configure
type FSMAction struct {
	ID     string `json:"id"`
	Action string `json:"action"`
}

//FSMStatusUpdate update message from ssm
type FSMStatusUpdate struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

//ErrorMessage Reply about the status of the FMS
type ErrorMessage struct {
	Type  string `json:"id"`
	Value string `json:"value"`
}
