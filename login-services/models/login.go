package loginModels

import "time"

type Login struct {
	Username string
	Password string
}

type Logout struct {
	Token string
}

type RegisterLoging struct {
	Id         int32
	Username   string
	UserId     int32
	LogedAt    time.Time
	LoginLimit time.Time
	Token      string
}
