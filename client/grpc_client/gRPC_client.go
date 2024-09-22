package grpcclient

import (
	"context"
	"log"
	"time"

	"github.com/MohitSilwal16/Picker/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const CONTEXT_TIMEOUT = time.Second * 20

var fileWatcherClient pb.FileWatcherClient

func NewGRPCClients(address string) error {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	fileWatcherClient = pb.NewFileWatcherClient(conn)

	return nil
}

func CreateFileRequest(filePath string) {
	// Errors:
	// CANNOT CREATE FILE
	// FILE ALREADY EXISTS
	// GIVEN PATH IS DIRECTORY
	// INTERNAL SERVER ERROR

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.CreateFile(ctxTimeout, &pb.CreateFileRequest{
		FilePath: filePath,
	})
	if err != nil {
		log.Println("Error from Server during CreateFile:", err.Error())
	}
}

func CreateDirRequest(dirPath string) {
	// Errors:
	// CANNOT CREATE DIRECTORY
	// DIRECTORY ALREADY EXISTS

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.CreateDir(ctxTimeout, &pb.CreateDirRequest{
		DirPath: dirPath,
	})
	if err != nil {
		log.Println("Error from Server during CreateDir:", err.Error())
	}
}

func RemoveFileDirRequest(fileDirPath string) {
	// Errors:
	// CANNOT REMOVE FILE/DIRECTORY
	// FILE/DIR DOESN'T EXISTS

	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancelFunc()

	_, err := fileWatcherClient.RemoveFileDir(ctxTimeout, &pb.RemoveFileDirRequest{
		FileDirPath: fileDirPath,
	})
	if err != nil {
		log.Println("Error from Server during RemoveFileDir:", err.Error())
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
	})
	if err != nil {
		log.Println("Error from Server during RenameFileDir:", err.Error())
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
		FilePath:    filePath,
		FileContent: []byte(fileContent),
	})
	if err != nil {
		log.Println("Error from Server during WriteFile:", err.Error())
	}
}
