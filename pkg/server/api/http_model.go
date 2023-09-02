package api

type RequestDTO struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

type ResponseDTO struct {
	Result string `json:"result"`
}
