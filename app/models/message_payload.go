package models

type MessagePayload struct {
	From    string `json:"from"`
	Message string `json:"message"`
}
