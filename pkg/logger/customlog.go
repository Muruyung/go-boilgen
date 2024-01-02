package logger

type LogJSON struct {
	HeaderRequest string `json:"header"`
	Request       string `json:"request"`
	Response      string `json:"response"`
}
