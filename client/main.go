package main

import (
	"fmt"
	"os"

	grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	"github.com/MohitSilwal16/Picker/client/myservice"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/joho/godotenv"
	"golang.org/x/sys/windows/svc"
)

func main() {
	utils.ClearScreen()

	isWindowsService, err := svc.IsWindowsService()
	if err != nil {
		fmt.Printf("Failed to Determine if Running in an Interactive Session(Shell): %v\n", err)
		return
	}

	configPath := utils.GetPathOfConfigFile()

	err = godotenv.Load(configPath)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Load", configPath)
	}
	serviceURL := os.Getenv("SERVICE_URL")

	err = grpcclient.NewGRPCClients(serviceURL)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Connect to Server")
		return
	}

	if isWindowsService {
		myservice.RunService()
		return
	}

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
