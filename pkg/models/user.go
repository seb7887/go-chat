package models

type User struct {
	ID       uint
	Username string
	Password string
}

type NewUserReq struct {
	Username string
	Password string
}

type NewUserResp struct {
	Id uint `json:"id"`
}
