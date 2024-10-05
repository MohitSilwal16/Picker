package grpcclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/MohitSilwal16/Picker/client/pb"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/spf13/viper"
)

const CONTEXT_TIMEOUT = time.Second * 20

func InitDirRequest(dirName string, allowedExtensions string) bool {
	// Errors:
	// USER IS ALREADY HOSTING dir1 DIR
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	fmt.Println("Initializing Directory into Server")

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	dirName, err := filepath.Abs(dirName)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Convert Relative Path into Absolute Path")
		return false
	}

	allowedExtensionsSlice := strings.Split(allowedExtensions, ",")

	err = utils.ZipFile("temp.zip", dirName, allowedExtensionsSlice)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Cannot Convert Dir into Zip")
		return false
	}

	data, err := os.ReadFile("temp.zip")
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Description: Failed to Read Zip File")
		return false
	}

	_, err = fileWatcherClient.InitDir(ctxTimeout, &pb.InitDirRequest{
		DirName:      filepath.Base(dirName),
		ExistingData: data,
		SessionToken: viper.GetString("session_token"),
	})
	if err != nil {
		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		fmt.Println("Error from Server during InitDir:", trimmedErr)

		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				InitDirRequest(dirName, allowedExtensions)
				return false
			}
			os.Exit(1)
		}
	}

	err = os.Remove("temp.zip")
	if err != nil {
		fmt.Println("Cannot delete temp.zip:", err)
	}

	fmt.Println("Directory Hosting/Upload Initialized")

	updatedDirsToWatch := viper.GetString("dirs_to_watch_absolute_path")
	updatedAllowedExt := viper.GetString("allowed_extensions")

	if updatedDirsToWatch == "" {
		updatedDirsToWatch = dirName
		updatedAllowedExt = allowedExtensions
	} else {
		updatedDirsToWatch += ";" + dirName
		updatedAllowedExt += ";" + allowedExtensions
	}

	viper.Set("allowed_extensions", updatedAllowedExt)
	viper.Set("dirs_to_watch_absolute_path", filepath.Join(updatedDirsToWatch, "..."))
	viper.WriteConfig()

	return true
}

func CreateFileRequest(filePath string, dirToWatchBasePath string) bool {
	// Errors:
	// CANNOT CREATE FILE
	// FILE ALREADY EXISTS
	// GIVEN PATH IS DIRECTORY
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.CreateFile(ctxTimeout, &pb.CreateFileRequest{
		FilePath:      filePath,
		SessionToken:  viper.GetString("session_token"),
		SenderDirName: dirToWatchBasePath,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during File Creation: REQUEST TIMED OUT")
			return false
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during CreateFile:", trimmedErr)

		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				return CreateFileRequest(filePath, dirToWatchBasePath)
			}
			log.Println("Error: Login Without Interaction Failed\nSource: CreateFileRequest")
			os.Exit(1)
		}
		return false
	}
	return true
}

func CreateDirRequest(dirPath string, dirToWatchBasePath string) {
	// Errors:
	// CANNOT CREATE DIRECTORY
	// DIRECTORY ALREADY EXISTS
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.CreateDir(ctxTimeout, &pb.CreateDirRequest{
		DirPath:       dirPath,
		SessionToken:  viper.GetString("session_token"),
		SenderDirName: dirToWatchBasePath,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during Dir Creation: REQUEST TIMED OUT")
			return
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during CreateDir:", trimmedErr)

		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				CreateDirRequest(dirPath, dirToWatchBasePath)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: CreateDirRequest()")
			os.Exit(1)
		}
	}
}

func RemoveFileDirRequest(fileDirPath string, dirToWatchBasePath string) {
	// Errors:
	// CANNOT REMOVE FILE/DIRECTORY
	// FILE/DIR DOESN'T EXISTS
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.RemoveFileDir(ctxTimeout, &pb.RemoveFileDirRequest{
		FileDirPath:   fileDirPath,
		SessionToken:  viper.GetString("session_token"),
		SenderDirName: dirToWatchBasePath,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during Deletion of File/Dir: REQUEST TIMED OUT")
			return
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during RemoveFileDir:", trimmedErr)
		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				RemoveFileDirRequest(fileDirPath, dirToWatchBasePath)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: RemoveFileDirPath()")
			os.Exit(1)
		}
	}
}

func RenameFileDirRequest(oldFileDirPath, newFileDirPath string, dirToWatchBasePath string) {
	// Errors:
	// CANNOT RENAME FILE/DIRECTORY
	// OLD FILE DOESN'T EXISTS
	// NEW FILE ALREADY EXISTS
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.RenameFileDir(ctxTimeout, &pb.RenameFileDirRequest{
		OldFileDirName: oldFileDirPath,
		NewFileDirName: newFileDirPath,
		SessionToken:   viper.GetString("session_token"),
		SenderDirName:  dirToWatchBasePath,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server while Renaming File/Dir: REQUEST TIMED OUT")
			return
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during RenameFileDir:", trimmedErr)
		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				RenameFileDirRequest(oldFileDirPath, newFileDirPath, dirToWatchBasePath)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: RenameFileDirRequest()")
			os.Exit(1)
		}
	}
}

func WriteFileRequest(filePath string, fileContent []byte, dirToWatchBasePath string) {
	// Errors:
	// CANNOT OPEN FILE
	// CANNOT WRITE INTO FILE
	// FILE DOESN'T EXISTS
	// GIVEN PATH IS DIRECTORY
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.WriteFile(ctxTimeout, &pb.WriteFileRequest{
		FilePath:      filePath,
		FileContent:   []byte(fileContent),
		SessionToken:  viper.GetString("session_token"),
		SenderDirName: dirToWatchBasePath,
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server while Writing in File: REQUEST TIMED OUT")
			return
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during WriteFile:", trimmedErr)
		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				WriteFileRequest(filePath, fileContent, dirToWatchBasePath)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: WriteFileRequest()")
			os.Exit(1)
		}
	}
}
