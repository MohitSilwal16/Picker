package grpcclient

import (
	"context"
	"log"
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
	}
}

func CreateDirRequest(dirPath string) {
	// Errors:
	// CANNOT CREATE DIRECTORY
	// DIRECTORY ALREADY EXISTS

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
	}
}

func RemoveFileDirRequest(fileDirPath string) {
	// Errors:
	// CANNOT REMOVE FILE/DIRECTORY
	// FILE/DIR DOESN'T EXISTS

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
	}
}

func RenameFileDirRequest(oldFileDirPath, newFileDirPath string) {
	// Errors:
	// CANNOT RENAME FILE/DIRECTORY
	// OLD FILE DOESN'T EXISTS
	// NEW FILE ALREADY EXISTS

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
	}
}

func WriteFileRequest(filePath string, fileContent []byte) {
	// Errors:
	// CANNOT OPEN FILE
	// CANNOT WRITE INTO FILE
	// FILE DOESN'T EXISTS
	// GIVEN PATH IS DIRECTORY

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
	}
}
