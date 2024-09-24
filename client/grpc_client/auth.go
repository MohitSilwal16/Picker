package grpcclient

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/MohitSilwal16/Picker/client/pb"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/spf13/viper"
)

func Register(name string, pass string) bool {
	// Errors:
	// USERNAME IS ALREADY USED
	// USER REGISTERED BUT FAILED TO INIT DIR
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := authClient.Register(ctxTimeout, &pb.AuthRequest{
		Name: name,
		Pass: pass,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("REQUEST TIMED OUT")
			log.Println("Error from Server during Register: REQUEST TIMED OUT")
			return false
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during Register:", trimmedErr)

		if trimmedErr == "USERNAME IS ALREADY USED" {
			fmt.Println("USERNAME IS ALREADY USED")
			return false
		}
		if trimmedErr == "USER REGISTERED BUT FAILED TO INIT DIR" {
			fmt.Println("USER REGISTERED BUT FAILED TO INIT DIR(It's fine unless you want to Upload Files)")
			return true
		}
		fmt.Println("INTERNAL SERVER ERROR")
		return false
	}

	viper.Set("username", name)
	viper.Set("password", pass)
	viper.Set("session_token", res.SessionToken)

	viper.WriteConfig()

	return true
}

func Login(name string, pass string) bool {
	// Errors:
	// INVALID CREDENTIALS
	// INTERNAL SERVER ERROR

	if !regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(name) {
		fmt.Println("USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC")
		fmt.Println("Ex - User123")
		return false
	}
	if !utils.IsPasswordInFormat(pass) {
		fmt.Println("PASSWORD MUST BE B'TWIN 8-20 CHARS, ATLEAST 1 UPPER, LOWER CASE & SPECIAL CHARACTER")
		fmt.Println("Ex - User123@")
		return false
	}

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := authClient.Login(ctxTimeout, &pb.AuthRequest{
		Name: name,
		Pass: pass,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("REQUEST TIMED OUT")
			log.Println("Error from Server during Login: REQUEST TIMED OUT")
			return false
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during Login:", trimmedErr)

		if trimmedErr == "INVALID CREDENTIALS" {
			fmt.Println("INVALID CREDENTIALS")
			return false
		}
		fmt.Println("INTERNAL SERVER ERROR")
		return false
	}

	viper.Set("username", name)
	viper.Set("password", pass)
	viper.Set("session_token", res.SessionToken)

	viper.WriteConfig()

	return true
}

func VerifySessionToken(sessionToken string) bool {
	// Errors:
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := authClient.VerifySessionToken(ctxTimeout, &pb.VerifySessionTokenRequest{
		SessionToken: sessionToken,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during Verify Session Token: REQUEST TIMED OUT")
			return false
		}
		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during Verifying Session Token:", trimmedErr)
		return false
	}
	return res.IsSessionTokenValid
}

func Logout(sessionToken string) bool {
	// Errors:
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	res, err := authClient.Logout(ctxTimeout, &pb.LogOutRequest{
		SessionToken: sessionToken,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during Logout: REQUEST TIMED OUT")
			return false
		}
		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during Logout:", trimmedErr)
		return false
	}

	viper.Set("session_token", "")
	viper.WriteConfig()

	return res.IsUserLoggedOut
}
