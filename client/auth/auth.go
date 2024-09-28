package auth

import (
	"fmt"
	"regexp"

	// grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/spf13/viper"
)

func LoginRegister() bool {
	var choice string
	var name string
	var pass string

	for {
		fmt.Println("\n\nEnter R to Register")
		fmt.Println("Enter L to Login")
		fmt.Println("Enter Q to Quit")
		fmt.Scanln(&choice)

		utils.ClearScreen()

		switch choice {
		case "R":
			for {
				fmt.Println("\nEnter Username: ")
				fmt.Scanln(&name)
				if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(name) {
					fmt.Println("USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC")
					fmt.Println("Ex - User123")
					continue
				}
				break
			}

			for {
				fmt.Println("\nEnter Password: ")
				fmt.Scanln(&pass)

				if !utils.IsPasswordInFormat(pass) {
					fmt.Println("PASSWORD MUST BE B'TWIN 8-20 CHARS, ATLEAST 1 UPPER, LOWER CASE & SPECIAL CHARACTER")
					fmt.Println("Ex - User123@")
					continue
				}
				break
			}

			if grpcclient.Register(name, pass) {
				return true
			}
		case "L":
			for {
				fmt.Println("\nEnter Username: ")
				fmt.Scanln(&name)
				if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(name) {
					fmt.Println("USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC")
					fmt.Println("Ex - User123")
					continue
				}
				break
			}

			for {
				fmt.Println("\nEnter Password: ")
				fmt.Scanln(&pass)

				if !utils.IsPasswordInFormat(pass) {
					fmt.Println("PASSWORD MUST BE B'TWIN 8-20 CHARS, ATLEAST 1 UPPER, LOWER CASE & SPECIAL CHARACTER")
					fmt.Println("Ex - User123@")
					continue
				}
				break
			}

			if grpcclient.Login(name, pass) {
				return true
			}
		case "Q":
			return false
		default:
			fmt.Println("Please Select Valid Option")
		}
	}
}

func AuthManager() bool {
	sessionToken := viper.GetString("session_token")
	if sessionToken == "" {
		return LoginRegister()
	}
	isSessionTokenValid := grpcclient.VerifySessionToken(sessionToken)
	if isSessionTokenValid {
		return true
	}
	if LoginWithoutInteraction() {
		return true
	}
	return LoginRegister()
}

func LoginWithoutInteraction() bool {
	username := viper.GetString("username")
	password := viper.GetString("password")

	return grpcclient.Login(username, password)
}
