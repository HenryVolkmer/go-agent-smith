package models


type AuthData struct {
	Type  *AuthType `json:"type"`
	Session string `json:"session"`
}

type AuthType struct {
	Type string `json:"type"`
}

func NewLoginDummyAuth() *AuthType {
	aT := AuthType{Type: "m.login.dummy"}
	
	return &aT
}
