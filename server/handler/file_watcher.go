package handler

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/MohitSilwal16/Picker/server/db"
	"github.com/MohitSilwal16/Picker/server/errs"
	"github.com/MohitSilwal16/Picker/server/pb"
	"github.com/MohitSilwal16/Picker/server/utils"
)

type FileWatcherServer struct {
	pb.UnimplementedFileWatcherServer
}

func (s *FileWatcherServer) InitDir(ctx context.Context, req *pb.InitDirRequest) (*pb.InitDirResponse, error) {
	// Errors:
	// USER IS ALREADY HOSTING dir1 DIR
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	username, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, errs.ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Get Username from Session Token")
		log.Println("Source: InitDir()")

		return nil, errs.ErrInternalServerError
	}

	err = db.AddUploader(username, req.DirName, req.ExistingData)
	if err != nil {
		if err.Error() == "USER IS ALREADY HOSTING "+req.DirName+" DIR" {
			return nil, err
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Add Uploader\nSource: InitDir()")
		return nil, errs.ErrInternalServerError
	}

	dirPath := "uploads/" + username + "/" + req.DirName
	err = os.MkdirAll(dirPath, 0644)
	if err != nil && !os.IsExist(err) {
		log.Println("Error:", err)
		log.Println("Description: Cannot Init Dir\nSource: InitDir()")
		return &pb.InitDirResponse{IsDirInitialized: false}, errs.ErrInternalServerError
	}

	zipPath := "temp/" + username + ".zip"
	// Create Zip File from Bytes
	initDirZip, err := os.Create(zipPath)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Create Zip File\nSource: InitDir()")
		return &pb.InitDirResponse{IsDirInitialized: false}, errs.ErrInternalServerError
	}
	defer initDirZip.Close()

	if _, err = initDirZip.Write(req.ExistingData); err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Write Content into Zip File\nSource: InitDir()")
		return &pb.InitDirResponse{IsDirInitialized: false}, errs.ErrInternalServerError
	}

	// Now Unzip that File & Store Existing Data
	err = utils.UnZipFile(zipPath, dirPath)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Unzip the Zip File\nSource: InitDir()")

		initDirZip.Close()
		err = os.RemoveAll(zipPath)
		if err != nil {
			log.Println("Error:", err)
			log.Println("Description: Cannot Delete Zip File after Failing to Unzip File\nSource: InitDir()")
			return &pb.InitDirResponse{IsDirInitialized: false}, errs.ErrInternalServerError
		}
		return &pb.InitDirResponse{IsDirInitialized: false}, errs.ErrInternalServerError
	}

	// Delete temp zip file
	// We need to close else we won't be able to delete since we marked to use this zip file above
	initDirZip.Close()
	err = os.Remove(zipPath)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Delete Zip File after Failing to Unzip File\nSource: InitDir()")
	}
	return &pb.InitDirResponse{IsDirInitialized: true}, nil
}

func (s *FileWatcherServer) CreateFile(ctx context.Context, req *pb.CreateFileRequest) (*pb.CreateFileResponse, error) {
	// Errors:
	// CANNOT CREATE FILE
	// FILE ALREADY EXISTS
	// GIVEN PATH IS DIRECTORY
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	sender, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, errs.ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Get Username from Session Token")
		log.Println("Source: CreateFile()")

		return nil, errs.ErrInternalServerError
	}

	filePath := "uploads/" + sender + "/" + req.SenderDirName + "/" + req.FilePath

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				log.Println("Error:", err)
				log.Println("Description: Given Path is a Directory or Incorrect Path")
				log.Println("Source: CreateFile()")
				return nil, errs.ErrGivenPathIsDir
			}
			log.Println("Error:", err)
			log.Println("Description: Cannot Create File Named", filePath)
			log.Println("Source: CreateFile()")

			return nil, errs.ErrCannotCreateFile
		}
		defer file.Close()

		return &pb.CreateFileResponse{FileCreated: true}, nil
	}

	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error while checking Stat of", filePath)
		log.Println("Source: CreateFile()")

		return nil, errs.ErrInternalServerError
	}

	return nil, errs.ErrFileAlreadyExists
}

func (s *FileWatcherServer) CreateDir(ctx context.Context, req *pb.CreateDirRequest) (*pb.CreateDirResponse, error) {
	// Errors:
	// CANNOT CREATE DIRECTORY
	// DIRECTORY ALREADY EXISTS
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	sender, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, errs.ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Get Username from Session Token")
		log.Println("Source: CreateFile()")

		return nil, errs.ErrInternalServerError
	}

	dirName := "uploads/" + sender + "/" + req.SenderDirName + "/" + req.DirPath

	err = os.Mkdir(dirName, 0644)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Println("Error:", err)
			log.Printf("Description: Directory '%s' Already Exists or Incorrect Path\n", dirName)
			log.Println("Source: CreateDir()")
			return nil, errs.ErrDirAlreadyExists
		}
		log.Println("Error:", err)
		log.Printf("Description: Cannot Create Directory Named '%s'\n", dirName)
		log.Println("Source: CreateDir()")
		return nil, errs.ErrCannotCreateDir
	}

	return &pb.CreateDirResponse{DirCreated: true}, nil
}

func (s *FileWatcherServer) RemoveFileDir(ctx context.Context, req *pb.RemoveFileDirRequest) (*pb.RemoveFileDirResponse, error) {
	// Errors:
	// CANNOT REMOVE FILE/DIRECTORY
	// FILE/DIR DOESN'T EXISTS
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	sender, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, errs.ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Get Username from Session Token")
		log.Println("Source: CreateFile()")

		return nil, errs.ErrInternalServerError
	}

	fileDirPath := "uploads/" + sender + "/" + req.SenderDirName + "/" + req.FileDirPath

	_, err = os.Stat(fileDirPath)
	if os.IsNotExist(err) {
		log.Printf("Error: Specified File/Dir '%s' Doesn't Exists\n", fileDirPath)
		log.Println("Source: RemoveFileDir()")
		return nil, errs.ErrFileDirDoesntExists
	}

	// os.RemoveAll() Doesn't Throws Error if File/Dir DOESN'T EXISTS
	err = os.RemoveAll(fileDirPath)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Remove File Named", fileDirPath)
		return nil, errs.ErrCannotRemoveFileDir
	}

	return &pb.RemoveFileDirResponse{FileRemoved: true}, nil
}

func (s *FileWatcherServer) RenameFileDir(ctx context.Context, req *pb.RenameFileDirRequest) (*pb.RenameFileDirResponse, error) {
	// Errors:
	// CANNOT RENAME FILE/DIRECTORY
	// OLD FILE DOESN'T EXISTS
	// NEW FILE ALREADY EXISTS
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	sender, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, errs.ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Get Username from Session Token")
		log.Println("Source: CreateFile()")

		return nil, errs.ErrInternalServerError
	}

	oldFileDirName := "uploads/" + sender + "/" + req.SenderDirName + "/" + req.OldFileDirName
	newFileDirName := "uploads/" + sender + "/" + req.SenderDirName + "/" + req.NewFileDirName

	doesNewFileDirAlreadyExists := true

	_, err = os.Stat(newFileDirName)
	if os.IsNotExist(err) {
		doesNewFileDirAlreadyExists = false
	} else if err != nil {
		log.Println("Error:", err)
		log.Printf("Description: Cannot Fetch Info of '%s'\n", newFileDirName)
		return nil, errs.ErrCannotRenameFileDir
	}

	if doesNewFileDirAlreadyExists {
		log.Printf("Error: New File/Dir '%s' Alredy Exists\n", newFileDirName)
		log.Println("Source: RenameFileDir()")
		return nil, errs.ErrNewFileAlreadyExists
	}

	err = os.Rename(oldFileDirName, newFileDirName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Error: Old File/Dir '%s' Doesn't Exists\n", oldFileDirName)
			log.Println("Source: RenameFileDir()")
			return nil, errs.ErrOldFileDoesntExists
		}

		log.Println("Error:", err)
		log.Println("Description: Cannot Rename File Named", oldFileDirName)
		return nil, errs.ErrCannotRenameFileDir
	}

	return &pb.RenameFileDirResponse{FileRenamed: true}, nil
}

func (s *FileWatcherServer) WriteFile(ctx context.Context, req *pb.WriteFileRequest) (*pb.WriteFileResponse, error) {
	// Errors:
	// CANNOT OPEN FILE
	// CANNOT WRITE INTO FILE
	// GIVEN PATH IS DIRECTORY
	// INVALID SESSION TOKEN
	// INTERNAL SERVER ERROR

	sender, err := db.GetUsernameBySessionToken(req.SessionToken)
	if err != nil {
		if err.Error() == "INVALID SESSION TOKEN" {
			return nil, errs.ErrInvalidSessionToken
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Get Username from Session Token")
		log.Println("Source: CreateFile()")

		return nil, errs.ErrInternalServerError
	}

	filePath := filepath.Join("uploads", sender, req.SenderDirName, req.FilePath)

	// Also Create File if doesn't Exists
	// Overwrite File's Content
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Println("Error: Given Path is Directory")
			log.Println("Source: WriteFile()")
			return nil, errs.ErrGivenPathIsDir
		}
		log.Println("Error:", err)
		log.Println("Description: Cannot Open File Named", filePath)
		log.Println("Source: WriteFile()")

		return nil, errs.ErrCannotOpenFile
	}
	defer file.Close()

	_, err = file.Write(req.FileContent)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Write in", filePath)
		return nil, errs.ErrCannotWriteIntoFile
	}

	return &pb.WriteFileResponse{FileWritten: true}, nil
}
