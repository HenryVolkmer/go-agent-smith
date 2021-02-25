package main

import (
	"fmt"
	"net/http"
	"os"
	"bytes"
	"flag"
	"io/ioutil"
	"github.com/HenryVolkmer/go-agent-smith/models"
)

func main() {


	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	username := registerCmd.String("username","","the username to register")
	password := registerCmd.String("password","","the password")
	homeserver := registerCmd.String("homeserver","","homeserver url")

	if len(os.Args) < 2 {
		fmt.Println("expected 'register' or 'login' subcommands")
        os.Exit(1)
	}

	switch os.Args[1] {
		case "register","login":
			registerCmd.Parse(os.Args[2:])
			if (*homeserver == "") {
				fmt.Println("homeserver missing")
				os.Exit(1)	
			}
			if (*username == "" || *password == "") {
				fmt.Println("username/password missing")
				os.Exit(1)	
			}
			user := models.GetUserInstance(*username,*password,*homeserver)
			requestSession(user,os.Args[1])
		default:
			fmt.Println("Unknown subcommand " + os.Args[1])
			os.Exit(1)


	}
}


func requestSession(user *models.User,kind string) *models.UserSession {

	endpoint := user.GetHomeserver() + "/_matrix/client/r0/login"

	if kind == "register" {
		fmt.Println("register new user")
		endpoint = user.GetHomeserver() + "/_matrix/client/r0/register?kind=user"
	} else {
		user.Type = "m.login.password"
		user.Auth = nil
		fmt.Println("login user")
	}

	resp, err := http.Post(endpoint,"application/x-www-form-urlencoded",bytes.NewBuffer(user.AsJson()))

	fmt.Println(string(user.AsJson()))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)

	sess,err := models.NewUserSession(body)	

	fmt.Println(resp)

	return sess
}