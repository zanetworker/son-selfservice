package models

//ActionMessage is the message for FrontEnd Actions
type ActionMessage struct {
	// The fields of this struct must be exported so that the json module will be
	// able to write into them. Therefore we need field tags to specify the names
	// by which these fields go in the JSON representation of events.
	// X int `json:"x"`
	// Y int `json:"y"`
	Name string `json:"name"`
	Data interface{}
}
