package configuration

//Configuration holds Debugger Configuration Details
type Configuration struct {
	DBHostIP   string `json:"host"`
	DBHostPort string `json:"port"`
}
