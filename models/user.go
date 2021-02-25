package models

import (
	"encoding/json"
)

/**
 * used for Register of new Users
 */
type UserReg struct {
	Username string `json:"username"`
	Password string `json:"password"`
	homeserver string
	Auth *AuthType `json:"auth"`
}

/**
 * used for Login with existing User
 * diffs from UserReg with json-key of Username ("user" <=> "username")
 */
type UserLogin struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Type string `json:"type"`
}

func (u *User) GetHomeserver() string {
	return u.homeserver
}

func (u *User) AsJson() []byte {

	json,err := json.Marshal(*u)

	if (err != nil) {
		panic(err)
	}

	return json
}

func GetUserForRegistration(username string,password string,homeserver string) *UserReg {
	user := UserReg{Username: username,Password: password,homeserver: homeserver,Auth: NewLoginDummyAuth()}
	return &user
}

func GetUserForLogin(username string,password string) *UserLogin {
	user := UserLogin{Username: username,Password: password,Type: "m.login.password"}
	return &user
}

/**
 * provides access_token and user-id
 */
type UserSession struct {
	AccessToken string `json:"access_token"`
	UserId string `json:"user_id"`
}

func NewUserSession(sessionData []byte) (*UserSession,error) {

	sess := UserSession{}

	if err := json.Unmarshal(sessionData, &sess); err != nil {
		return nil,err
	}

	return &sess,nil
}