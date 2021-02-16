package model

type UserVO struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

type UserQry struct {
	UserId string `json:"userId"`
}
