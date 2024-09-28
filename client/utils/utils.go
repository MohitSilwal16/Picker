package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	"github.com/spf13/viper"
)

func ClearScreen() {
	var cmd *exec.Cmd

	// Check the operating system to determine the appropriate clear command
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear") // for Unix-like systems
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	default:
		fmt.Println("Unsupported platform.")
		return
	}

	// Execute the clear command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func IsPasswordInFormat(s string) bool {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) < 8 || len(s) > 20 {
		return false
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}

func TrimGrpcErrorMessage(errMsg string) string {
	// Split the error message
	parts := strings.Split(errMsg, "desc = ")
	if len(parts) > 1 {
		// Return the part after "desc = "
		return parts[1]
	}
	// Return the original error message if "desc = " is not found
	return errMsg
}

func SetUpFileLogging(logFilePath string) (*os.File, error) {
	if logFilePath == "" {
		logFilePath = filepath.Dir(GetPathOfConfigFile()) + "/picker.log"
	}
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Println("Description: Error while Creating Log File at the Path Specified")
				return nil, err
			}
			logFilePath = GetPathOfConfigFile() + "/picker.log"
			viper.Set("log_file_absolute_path", logFilePath)
			viper.WriteConfig()

			log.SetOutput(logFile)
			return logFile, nil
		}
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Setup Log File")
		return nil, err
	}
	log.SetOutput(logFile)
	return logFile, nil
}

// Validations & Checks
func IsServiceExePathValid(serviceExecutablePath string, configPath string) bool {
	if serviceExecutablePath == "" {
		fmt.Println("Service's Executable Path is Empty in", configPath)
		return false
	}

	if len(serviceExecutablePath) >= 4 {
		lastFour := serviceExecutablePath[len(serviceExecutablePath)-4:]
		if lastFour != ".exe" {
			fmt.Println("Invalid Extension of Service Executable in", configPath)
			fmt.Println("Extension of Service Executable MUST be '.exe'")
			return false
		}
	} else {
		fmt.Println("Invalid Path of Service Executable in", configPath)
		return false
	}

	_, err := os.Stat(serviceExecutablePath)
	if os.IsNotExist(err) {
		fmt.Println("Invalid Path of Service Executable in", configPath)
		return false
	} else if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Error while Checking the Existance of Service Executable")
		return false
	}
	return true
}

func IsDirsToWatch_IgnoreExtensions_Valid(dirToWatchString string, ignoreExtensionsString string) bool {
	dirsToWatch := strings.Split(dirToWatchString, ";")
	ignoreExtensions := strings.Split(ignoreExtensionsString, ";")

	if len(dirsToWatch) != len(ignoreExtensions) {
		fmt.Println("Length of Dirs To Watch & Ignore Extensions MUST BE SAME")
		return false
	}

	for idx := range dirsToWatch {
		info, err := os.Stat(dirsToWatch[idx])
		if os.IsNotExist(err) {
			fmt.Println("Invalid Path of Dir To Watch")
			return false
		} else if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Description: Error while Checking the Existance of Dir to Watch")
			return false
		}

		if !info.IsDir() {
			fmt.Println("Dir to Watch MUST BE Directory")
			return false
		}

		if strings.Contains(ignoreExtensions[idx], ".") {
			fmt.Println("Extensions MUST not have '.'")
			return false
		}
	}
	return true
}

func IsLogFilePathValid(logFilePath string, configPath string) bool {
	if logFilePath == "" {
		logFilePath = filepath.Dir(GetPathOfConfigFile()) + "/picker.log"
	}

	if len(logFilePath) >= 4 {
		lastFour := logFilePath[len(logFilePath)-4:]
		if lastFour != ".log" {
			fmt.Println("Invalid Extension of Log File in", configPath)
			fmt.Println("Extension of Log File MUST be '.log'")
			return false
		}
	} else {
		fmt.Println("Invalid Path of Log File in", configPath)
		return false
	}

	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		_, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Description: Error while Creating Log File at the Path Specified in", configPath)
			return false
		}
		viper.Set("log_file_absolute_path", logFilePath)
		viper.WriteConfig()
		return true
	} else if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Error while Checking the Existance of Log File")
		return false
	}
	return true
}

func GetPathOfConfigFile() string {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Get Path of Service Executable using os.Executable()")
	}

	currDir := filepath.Dir(execPath)
	configPath := filepath.Join(currDir, "config.json")

	return configPath
}
