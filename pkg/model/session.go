package model

type Session struct {
	sessionId    string
	sessionValue int
}

type SessionInput struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}
