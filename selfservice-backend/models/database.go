package models

//FSM  update message to the database
type FSM struct {
	ID    string `gorethink:"id,omitempty"`
	FsmID string `gorethink:"fsmId,omitempty"`
	Name  string `gorethink:"name"`
	State string `gorethink:"state"`
}

//DBChangeResponse struct to save changes from the databse
type DBChangeResponse struct {
	NewValue interface{} `gorethink:"new_val,omitempty"`
	OldValue interface{} `gorethink:"old_val,omitempty"`
	State    string      `gorethink:"state,omitempty"`
	Error    string      `gorethink:"error,omitempty"`
}
