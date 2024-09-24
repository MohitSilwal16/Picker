package myservice

import (
	"fmt"
	"log"

	grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/spf13/viper"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

const MY_SERVICE_NAME = "Picker"
const MY_SERVICE_DESC = "Picker, File Transfer Without User Interaction"

func InstallService() {
	configPath := utils.GetDirOfConfigFile()
	serviceExePath := viper.GetString("service_executable_absolute_path")
	serviceStartTypeString := viper.GetString("service_start_type")
	dirToWatch := viper.GetString("dir_to_watch_absolute_path")
	logFilePath := viper.GetString("log_file_absolute_path")
	ignoreExtensions := viper.GetString("ignore_extensions")

	isServiceExeValid := utils.IsServiceExePathValid(serviceExePath, configPath)
	if !isServiceExeValid {
		fmt.Println("Invalid Configuration of Service Executable in", configPath)
		return
	}

	isDirToWatchValid := utils.IsDirToWatchValid(dirToWatch, configPath)
	if !isDirToWatchValid {
		fmt.Println("Invalid Configuration of Dir to Watch in", configPath)
		return
	}

	isLogFilePathValid := utils.IsLogFilePathValid(logFilePath, configPath)
	if !isLogFilePathValid {
		fmt.Println("Invalid Configuration of Log Path in", configPath)
		return
	}

	areIgnoreFileExtensionsValid := utils.AreIgnoreFileExtensionsValid(ignoreExtensions, configPath)
	if !areIgnoreFileExtensionsValid {
		fmt.Println("Invalid Configuration of Ignore Extensions in", configPath)
		return
	}

	var serviceStartType uint32
	if serviceStartTypeString == "StartManual" {
		serviceStartType = mgr.StartManual
	} else if serviceStartTypeString == "StartAutomatic" {
		serviceStartType = mgr.StartAutomatic
	} else if serviceStartTypeString == "StartDisabled" {
		serviceStartType = mgr.StartDisabled
	} else {
		fmt.Println("Unknown SERVICE_START_TYPE:", serviceStartTypeString)
		fmt.Println("It MUST be one of StartManual, StartAutomatic or StartDisabled")
		return
	}

	m, err := mgr.Connect()
	if err != nil {
		fmt.Printf("Could not Connect to Service Manager: %v\n", err)
		return
	}
	defer m.Disconnect()

	s, err := m.CreateService(MY_SERVICE_NAME, serviceExePath, mgr.Config{
		DisplayName: MY_SERVICE_DESC,
		StartType:   serviceStartType,
	})
	if err != nil {
		fmt.Printf("Could not Create Service: %v\n", err)
		return
	}
	defer s.Close()

	fmt.Printf("Service %s Installed Successfully\n", s.Name)
}

func UinstallService() {
	// Even if we delete Service it's still gonna run if the service was already running
	StopService()

	m, err := mgr.Connect()
	if err != nil {
		fmt.Printf("Failed to Connect to Service Manager: %v\n", err)
		return
	}
	defer m.Disconnect()

	s, err := m.OpenService(MY_SERVICE_NAME)
	if err != nil {
		fmt.Printf("Service %s Not Found: %v\n", MY_SERVICE_NAME, err)
		return
	}
	defer s.Close()

	err = s.Delete()
	if err != nil {
		fmt.Printf("Failed to Delete Service: %v\n", err)
		return
	}

	fmt.Printf("Service %s Uninstalled Successfully\n", MY_SERVICE_NAME)

	// Delete Users Session Token from Server when he/she uninstalls Service
	grpcclient.Logout(viper.GetString("session_token"))
}

func StartService() {
	m, err := mgr.Connect()
	if err != nil {
		fmt.Printf("Could not Connect to Service Manager: %v\n", err)
		return
	}

	defer m.Disconnect()

	s, err := m.OpenService(MY_SERVICE_NAME)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Error while Opening Service")
		return
	}
	defer s.Close()

	err = s.Start()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Error while Starting Service")
		return
	}

	fmt.Println("Service Started Successfully")
}

func StopService() {
	m, err := mgr.Connect()
	if err != nil {
		fmt.Printf("Could not Connect to Service Manager: %v\n", err)
		return
	}

	defer m.Disconnect()

	s, err := m.OpenService(MY_SERVICE_NAME)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer s.Close()

	_, err = s.Control(svc.Stop)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Service Stopped Successfully")
}

func RunService() {
	err := svc.Run(MY_SERVICE_NAME, &MyService{})
	if err != nil {
		log.Println("Error:", err)
		return
	}
}
