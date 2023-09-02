package server

type RequestDTO struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

type ResponseDTO struct {
	Result string `json:"result"`
}
