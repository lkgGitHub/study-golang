package models

type Question struct {
	Id       uint   `json:"Id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}


