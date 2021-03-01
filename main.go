package main

import (
	"log"
	"fmt"
	"net/http"
	"os"
	"bytes"
	"flag"
	"io/ioutil"
	"encoding/json"
	"github.com/HenryVolkmer/go-agent-smith/models"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hello World!"
	p.SetRect(0, 0, 25, 5)

	buffer := make([]string,1)
	var result string

	ui.Render(p)

	uiEvents := ui.PollEvents()

	for {
		select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				case "<Space>":
					buffer = append(buffer," ")
					break
				case "<Backspace>":
					buffer = buffer[:len(buffer)-1]
					break
				case "<Enter>":
					result = ""
					result = for _,v := range buffer {
						result += v
					}
				default:
					buffer = append(buffer,e.ID)
				}
				p.Text = ""
				for _,v := range buffer {
					p.Text += v
				}
				ui.Render(p)

		}
	}

	fmt.Println(result)

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

			//var user models.UserInterface

			if os.Args[1] == "register" {
				user := models.GetUserForRegistration(*username,*password,*homeserver)
				register(user)
			} else {
				user := models.GetUserForLogin(*username,*password,*homeserver)
				login(user)
			}
		default:
			fmt.Println("Unknown subcommand " + os.Args[1])
			os.Exit(1)


	}
}

func register(user *models.UserReg) *models.UserSession {
	endpoint := user.GetHomeserver() + "/_matrix/client/r0/register?kind=user"
	return getSession(user,endpoint)
}

func login(user *models.UserLogin) *models.UserSession {
	endpoint := user.GetHomeserver() + "/_matrix/client/r0/login"
	return getSession(user,endpoint)
}

func getSession(user models.UserInterface,endpoint string) *models.UserSession {

	payload,err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(endpoint,"application/x-www-form-urlencoded",bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
	sess,err := models.NewUserSession(body)	

	return sess
}