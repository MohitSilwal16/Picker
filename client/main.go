package main

import (
	"fmt"
	"os"

	grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	"github.com/MohitSilwal16/Picker/client/myservice"
	"github.com/MohitSilwal16/Picker/client/utils"
	"golang.org/x/sys/windows/svc"
)

const SERVICE_ADDRESS = "192.168.65.1:8080" // IP Address if Service is running inside VM
// const SERVICE_ADDRESS = "localhost:8080"

func main() {
	utils.ClearScreen()

	isWindowsService, err := svc.IsWindowsService()
	if err != nil {
		fmt.Printf("Failed to Determine if Running in an Interactive Session(Shell): %v", err)
		return
	}

	err = grpcclient.NewGRPCClients(SERVICE_ADDRESS)
	if err != nil {
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
