package handler

import (
	"context"
	"log"
	"os"

	"github.com/MohitSilwal16/Picker/server/pb"
)

type FileWatcherServer struct {
	pb.UnimplementedFileWatcherServer
}

func (s *FileWatcherServer) CreateFile(ctx context.Context, req *pb.CreateFileRequest) (*pb.CreateFileResponse, error) {
	// Errors:
	// CANNOT CREATE FILE
	// FILE ALREADY EXISTS
	// GIVEN PATH IS DIRECTORY
	// INTERNAL SERVER ERROR

	filePath := "uploads\\" + req.FilePath

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		_, err = os.Create(filePath)
		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				log.Println("Error: Given Path is a Directory")
				log.Println("Source: CreateFile()")
				return nil, ErrGivenPathIsDir
			}
			log.Println("Error:", err)
			log.Println("Description: Cannot Create File Named", filePath)
			log.Println("Source: CreateFile()")

			return nil, ErrCannotCreateFile
		}

		return &pb.CreateFileResponse{FileCreated: true}, nil
	}

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error while checking Stat of", filePath)
		log.Println("Source: CreateFile()")

		return nil, ErrInternalServer
	}

	return nil, ErrFileAlreadyExists
}

func (s *FileWatcherServer) CreateDir(ctx context.Context, req *pb.CreateDirRequest) (*pb.CreateDirResponse, error) {
	// Errors:
	// CANNOT CREATE DIRECTORY
	// DIRECTORY ALREADY EXISTS

	dirName := "uploads\\" + req.DirPath

	err := os.Mkdir(dirName, 0644)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Printf("Error: Directory '%s' Already Exists", dirName)
			log.Println("Source: CreateDir()")
			return nil, ErrDirAlreadyExists
		}
		log.Println("Error:", err)
		log.Printf("Description: Cannot Create Directory Named '%s'\n", dirName)
		log.Println("Source: CreateDir()")
		return nil, ErrCannotCreateDir
	}

	return &pb.CreateDirResponse{DirCreated: true}, nil
}

func (s *FileWatcherServer) RemoveFileDir(ctx context.Context, req *pb.RemoveFileDirRequest) (*pb.RemoveFileDirResponse, error) {
	// Errors:
	// CANNOT REMOVE FILE/DIRECTORY
	// FILE/DIR DOESN'T EXISTS

	fileDirPath := "uploads\\" + req.FileDirPath

	err := os.RemoveAll(fileDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Error: Specified File/Dir '%s' Doesn't Exists\n", fileDirPath)
			log.Println("Source: RemoveFileDir()")
			return nil, ErrFileDirDoesntExists
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Remove File Named", fileDirPath)
		return nil, ErrCannotRemoveFileDir
	}

	return &pb.RemoveFileDirResponse{FileRemoved: true}, nil
}

func (s *FileWatcherServer) RenameFileDir(ctx context.Context, req *pb.RenameFileDirRequest) (*pb.RenameFileDirResponse, error) {
	// Errors:
	// CANNOT RENAME FILE/DIRECTORY
	// OLD FILE DOESN'T EXISTS
	// NEW FILE ALREADY EXISTS

	oldFileDirName := "uploads\\" + req.OldFileDirName
	newFileDirName := "uploads\\" + req.NewFileDirName

	doesNewFileDirAlreadyExists := true

	_, err := os.Stat(newFileDirName)
	if os.IsNotExist(err) {
		doesNewFileDirAlreadyExists = false
	} else if err != nil {
		log.Println("Error:", err)
		log.Printf("Description: Cannot Fetch Info of '%s'\n", newFileDirName)
		return nil, ErrCannotRenameFileDir
	}

	if doesNewFileDirAlreadyExists {
		log.Printf("Error: New File/Dir '%s' Alredy Exists\n", newFileDirName)
		log.Println("Source: RenameFileDir()")
		return nil, ErrNewFileAlreadyExists
	}

	err = os.Rename(oldFileDirName, newFileDirName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Error: Old File/Dir '%s' Doesn't Exists\n", oldFileDirName)
			log.Println("Source: RenameFileDir()")
			return nil, ErrOldFileDoesntExists
		}

		log.Println("Error:", err)
		log.Println("Description: Cannot Rename File Named", oldFileDirName)
		return nil, ErrCannotRenameFileDir
	}

	return &pb.RenameFileDirResponse{FileRenamed: true}, nil
}

func (s *FileWatcherServer) WriteFile(ctx context.Context, req *pb.WriteFileRequest) (*pb.WriteFileResponse, error) {
	// Errors:
	// CANNOT OPEN FILE
	// CANNOT WRITE INTO FILE
	// FILE DOESN'T EXISTS
	// GIVEN PATH IS DIRECTORY

	filePath := "uploads\\" + req.FilePath

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Printf("Error: Given File '%s' Doesn't Exists\n", filePath)
		log.Println("Source: WriteFile()")

		return nil, ErrFileDoesntExists
	} else if err != nil {
		log.Println("Error:", err)
		log.Printf("Description: Cannot Fetch Stats of '%s'\n", filePath)
		log.Println("Source: WriteFile()")

		return nil, ErrGivenPathIsDir
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Println("Error: Given Path is Directory")
			log.Println("Source: WriteFile()")
			return nil, ErrGivenPathIsDir
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Open File Named", filePath)
		log.Println("Source: WriteFile()")

		return nil, ErrCannotOpenFile
	}

	_, err = file.Write(req.FileContent)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Write in", filePath)
		return nil, ErrCannotWriteIntoFile
	}

	return &pb.WriteFileResponse{FileWritten: true}, nil
}
