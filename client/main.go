package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MohitSilwal16/Picker/client/auth"
	grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	"github.com/MohitSilwal16/Picker/client/service"
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

	configFilePath := utils.GetPathOfConfigFile()
	viper.SetConfigFile(configFilePath)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Read Config File")
		log.Println("Error:", err)
		log.Println("Description: Cannot Read Config File")
		return
	}

	serviceURL := viper.GetString("service_url")
	logFilePath := viper.GetString("log_file_absolute_path")
	sessionToken := viper.GetString("session_token")

	logFile, err := utils.SetUpFileLogging(logFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Set Log File")
		log.Println("Error:", err)
		log.Println("Description: Cannot Set Log File")
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
			service.RunService()
		} else if auth.LoginWithoutInteraction() {
			service.RunService()
		} else {
			log.Println("Auth Unsuccessful, Try Modifying Credentials in", configFilePath)
		}
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
		fmt.Println("Enter R to Restart Service")
		fmt.Println("Enter H to Host/Upload New Dir")
		fmt.Println("Enter F to Fetch/Subscribe New Dir")
		fmt.Println("Enter Q to Quit")

		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)

		utils.ClearScreen()

		switch choice {
		case "I":
			fmt.Println("Installing Service ...")
			service.InstallService()
		case "U":
			fmt.Println("Uinstalling Service ...")
			service.UinstallService()
			os.Exit(0)
		case "S":
			fmt.Println("Starting Service ...")
			service.StartService()
		case "T":
			fmt.Println("Stopping Service ...")
			service.StopService()
		case "R":
			// TODO: Restart Service
			fmt.Println("Restarting Service ...")
			service.StopService()
			service.StartService()
		case "H":
			var dirToHost, allowedExtensions string
			fmt.Println("Enter the Dir's Path you want to Host: ")
			fmt.Scanln(&dirToHost)
			fmt.Println("Enter Allowed Extensions: ")
			fmt.Scanln(&allowedExtensions)

			isDirsToWatch_ignoreExtensions_Valid := utils.IsDirsToWatch_IgnoreExtensions_Valid(dirToHost, allowedExtensions)
			if !isDirsToWatch_ignoreExtensions_Valid {
				fmt.Println("Invalid Configuration of Dirs to Watch & Ignore Extensions")
				continue
			}

			if grpcclient.InitDirRequest(dirToHost, allowedExtensions) {
				fmt.Println("Restart Service to Affect Changes")
			}

		case "F":
			// TODO: Subscribe to Uploader
		case "Q":
			fmt.Println("Terminating ...")
			os.Exit(0)
		default:
			fmt.Println("Please Select a Valid Option ...")
		}
	}
}
