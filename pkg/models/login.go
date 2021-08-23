package models

type Login struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
