package model

type ResponseJSON struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseFailJSON struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
