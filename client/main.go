package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MohitSilwal16/Picker/client/auth"
	grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	"github.com/MohitSilwal16/Picker/client/myservice"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/spf13/viper"
	"golang.org/x/sys/windows/svc"
)

func main() {
	utils.ClearScreen()

	isWindowsService, err := svc.IsWindowsService()
	if err != nil {
		fmt.Printf("Failed to Determine if Running in an Interactive Session(Shell): %v\n", err)
		return
	}

	configFilePath := utils.GetDirOfConfigFile()
	viper.SetConfigFile(configFilePath)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Read Config File")
		return
	}

	serviceURL := viper.GetString("service_url")
	logFilePath := viper.GetString("log_file_absolute_path")
	sessionToken := viper.GetString("session_token")

	logFile, err := utils.SetUpFileLogging(logFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Set Log File")
		return
	}
	defer logFile.Close()

	err = grpcclient.NewGRPCClients(serviceURL)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Connect to Server")
		log.Println("Error from Server:", utils.TrimGrpcErrorMessage(err.Error()))
		return
	}

	if isWindowsService {
		isSessionTokenValid := grpcclient.VerifySessionToken(sessionToken)
		if isSessionTokenValid {
			myservice.RunService()
		} else if auth.LoginWithoutInteraction() {
			myservice.RunService()
		}
		log.Println("Auth Unsuccessful, Try Modifying Credentials in", configFilePath)
		return
	}

	// If Auth Unsuccessful then exit
	if !auth.AuthManager() {
		fmt.Println("AUTH UNSUCCESSFUL")
		return
	}
	fmt.Println("Auth Successful")

	var choice string

	for {
		fmt.Print("\n\n")
		fmt.Println("Welcome to Picker")
		fmt.Println("Enter I to Install Service")
		fmt.Println("Enter U to Uinstall Service")
		fmt.Println("Enter S to Start Service")
		fmt.Println("Enter T to Stop Service")
		fmt.Println("Enter R to Refresh Requests")
		fmt.Println("Enter Q to Quit")

		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)

		utils.ClearScreen()

		switch choice {
		case "I":
			fmt.Println("Installing Service ...")
			myservice.InstallService()
		case "U":
			fmt.Println("Uinstalling Service ...")
			myservice.UinstallService()
		case "S":
			fmt.Println("Starting Service ...")
			myservice.StartService()
		case "T":
			fmt.Println("Stopping Service ...")
			myservice.StopService()
		case "R":
			// Refresh Requests to Check Update
			fmt.Println("Requesting Again to Server ...")
		case "Q":
			fmt.Println("Terminating ...")
			os.Exit(0)
		default:
			fmt.Println("Please Select a Valid Option ...")
		}
	}
}
