package grpcclient

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MohitSilwal16/Picker/client/pb"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/spf13/viper"
)

const CONTEXT_TIMEOUT = time.Second * 20

func CreateFileRequest(filePath string) {
	// Errors:
	// CANNOT CREATE FILE
	// FILE ALREADY EXISTS
	// GIVEN PATH IS DIRECTORY
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.CreateFile(ctxTimeout, &pb.CreateFileRequest{
		FilePath:     filePath,
		SessionToken: viper.GetString("session_token"),
	})
	if err != nil {
		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during CreateFile:", trimmedErr)

		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				CreateFileRequest(filePath)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: CreateFileRequest")
			os.Exit(1)
		}
	}
}

func CreateDirRequest(dirPath string) {
	// Errors:
	// CANNOT CREATE DIRECTORY
	// DIRECTORY ALREADY EXISTS
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.CreateDir(ctxTimeout, &pb.CreateDirRequest{
		DirPath:      dirPath,
		SessionToken: viper.GetString("session_token"),
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during Register: REQUEST TIMED OUT")
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during CreateDir:", trimmedErr)

		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				CreateDirRequest(dirPath)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: CreateDirRequest()")
			os.Exit(1)
		}
	}
}

func RemoveFileDirRequest(fileDirPath string) {
	// Errors:
	// CANNOT REMOVE FILE/DIRECTORY
	// FILE/DIR DOESN'T EXISTS
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.RemoveFileDir(ctxTimeout, &pb.RemoveFileDirRequest{
		FileDirPath:  fileDirPath,
		SessionToken: viper.GetString("session_token"),
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during Register: REQUEST TIMED OUT")
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during RemoveFileDir:", trimmedErr)
		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				RemoveFileDirRequest(fileDirPath)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: RemoveFileDirPath()")
			os.Exit(1)
		}
	}
}

func RenameFileDirRequest(oldFileDirPath, newFileDirPath string) {
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
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during Register: REQUEST TIMED OUT")
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during RenameFileDir:", trimmedErr)
		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				RenameFileDirRequest(oldFileDirPath, newFileDirPath)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: RenameFileDirRequest()")
			os.Exit(1)
		}
	}
}

func WriteFileRequest(filePath string, fileContent []byte) {
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
		FilePath:     filePath,
		FileContent:  []byte(fileContent),
		SessionToken: viper.GetString("session_token"),
	})
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error from Server during Register: REQUEST TIMED OUT")
		}

		trimmedErr := utils.TrimGrpcErrorMessage(err.Error())
		log.Println("Error from Server during WriteFile:", trimmedErr)
		if trimmedErr == "INVALID SESSION TOKEN" {
			username := viper.GetString("username")
			password := viper.GetString("password")

			isLoginSuccessful := Login(username, password)
			if isLoginSuccessful {
				WriteFileRequest(filePath, fileContent)
				return
			}
			log.Println("Error: Login Without Interaction Failed\nSource: WriteFileRequest()")
			os.Exit(1)
		}
	}
}
