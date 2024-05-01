package models

type Errors struct {
	PersError string `json:"pers_error"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
}

type DocumentsToSend struct {
	Document string `json:"document" bson:"document" binding:"required"`
}
