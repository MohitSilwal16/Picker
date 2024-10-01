package service

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/rjeczalik/notify"
)

var newFileDirPath = ""

// While Renaming File,
// Rename Event is triggered twice
// 1st time it's triggered with file/dir name as New Name
// 2nd time it's triggered with file/dir name as Old Name
func watchChanges(c chan notify.EventInfo, allowedExtensions []string, dirToWatchAbsPath string) {
	log.Println("Watching For Changes ...")
	for {
		event := <-c

		old_file_name_extension := strings.Split(event.Path(), ".")
		if len(old_file_name_extension) == 2 {
			changed_file_extension := old_file_name_extension[1]
			if event.Event() != notify.Rename && !utils.Contains(allowedExtensions, changed_file_extension) {
				continue
			}
		}
		path := strings.TrimPrefix(event.Path(), dirToWatchAbsPath)
		dirToWatchBasePath := filepath.Base(dirToWatchAbsPath)

		log.Println(event.Event(), ":", path)

		switch event.Event() {
		case notify.Create:
			info, err := os.Stat(event.Path())
			if err != nil {
				log.Println("Error:", err)
				log.Println("Description: Error while checking Stat of", event.Path())
				continue
			}

			if info.IsDir() {
				grpcclient.CreateDirRequest(path, dirToWatchBasePath)
			} else {
				grpcclient.CreateFileRequest(path, dirToWatchBasePath)
			}
		case notify.Remove:
			grpcclient.RemoveFileDirRequest(path, dirToWatchBasePath)
		case notify.Write:
			file, err := os.OpenFile(event.Path(), os.O_RDONLY, 0644)
			if err != nil {
				log.Println("Error:", err)
				log.Println("Description: Error while Opening", event.Path())
				continue
			}

			data, err := io.ReadAll(file)
			if err != nil {
				log.Println("Error:", err)
				log.Println("Description: Error Reading File", event.Path())
				file.Close()
				continue
			}
			file.Close()
			grpcclient.WriteFileRequest(path, data, dirToWatchBasePath)
		case notify.Rename:
			if newFileDirPath == "" {
				newFileDirPath = event.Path()
				continue
			}
			// Check whether New Name is in Allowed Extensions or not
			// If not then Delete it
			new_file_name_extension := strings.Split(newFileDirPath, ".")
			if len(new_file_name_extension) == 2 {
				new_file_extension := new_file_name_extension[1]
				if !utils.Contains(allowedExtensions, new_file_extension) {
					grpcclient.RemoveFileDirRequest(path, dirToWatchBasePath)
					newFileDirPath = ""
					continue
				}
			}

			// Check whether Old Name is in Allowed Extensions or not
			// If not then Create & Write in that file
			old_file_name_extension = strings.Split(path, ".")
			if len(old_file_name_extension) == 2 {
				old_file_extension := old_file_name_extension[1]
				if !utils.Contains(allowedExtensions, old_file_extension) {
					// Check the Extension of New file, If it's not allowed then ignore (continue)
					if len(new_file_name_extension) == 2 {
						new_file_extension := new_file_name_extension[1]
						if !utils.Contains(allowedExtensions, new_file_extension) {
							log.Println("Debug:", old_file_name_extension, new_file_name_extension, "Ignored")
							log.Println("New File/Dir Path:", newFileDirPath)
							newFileDirPath = ""
							continue
						}
					}
					log.Println("Debug:", old_file_name_extension, new_file_name_extension, "New File Creation Request")
					log.Println("New File/Dir Path:", newFileDirPath)
					// If the New file's Extension is allowed then, Create File & Write
					createFileDirAndWrite(newFileDirPath, dirToWatchBasePath, dirToWatchAbsPath)
					newFileDirPath = ""
					continue
				}
			}

			newFileDirPath = strings.TrimPrefix(newFileDirPath, dirToWatchAbsPath)
			grpcclient.RenameFileDirRequest(path, newFileDirPath, dirToWatchBasePath)
			newFileDirPath = ""
		}
	}
}

func createFileDirAndWrite(path string, dirToWatchBasePath string, dirToWatchAbsPath string) {
	info, err := os.Stat(path)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error while checking Stat of", path)
		return
	}

	if info.IsDir() {
		grpcclient.CreateDirRequest(path, dirToWatchBasePath)
		return
	}
	basePath, err := filepath.Rel(dirToWatchAbsPath, path)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error while getting Relative Path")
	}
	log.Println("Base Path:", basePath)

	grpcclient.CreateFileRequest(basePath, dirToWatchBasePath)
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error while Opening", path)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Error Reading File", path)
		return
	}
	grpcclient.WriteFileRequest(basePath, data, dirToWatchBasePath)
}
