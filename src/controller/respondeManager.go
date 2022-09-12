package controller

//clases
type ResponseManager struct {
	Msg        string                 "json:\"msg\""
	StatusCode int                    "json:\"statusCode\""
	Status     string                    "json:\"status\""
	Data       map[string]interface{} "json:\"data\""
}

func NewResponseManager() *ResponseManager {
	return &ResponseManager{
		Msg:        "Succes",
		StatusCode: 200,
		Status:     "Succes",
		Data:       make(map[string]interface{}),
	}
}
