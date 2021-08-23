package models

type Login struct {
	Id    uint   `json:"id"`
	Token string `json:"token"`
}
