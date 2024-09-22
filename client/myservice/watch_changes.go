package myservice

import (
	"io"
	"log"
	"os"
	"strings"

	grpcclient "github.com/MohitSilwal16/Picker/client/grpc_client"
	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/rjeczalik/notify"
)

var newFileDirPath = ""

func watchChanges(c chan notify.EventInfo, ignoreExtensions []string, dirToWatch string) {
	log.Println("Watching For Changes ...")
	for {
		event := <-c

		temp := strings.Split(event.Path(), ".")
		if len(temp) == 2 {
			changed_file_extension := temp[1]
			if utils.Contains(ignoreExtensions, changed_file_extension) {
				continue
			}
		}
		path := strings.TrimPrefix(event.Path(), dirToWatch)
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
				grpcclient.CreateDirRequest(path)
			} else {
				grpcclient.CreateFileRequest(path)
			}
		case notify.Remove:
			grpcclient.RemoveFileDirRequest(path)
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
			}

			grpcclient.WriteFileRequest(path, data)
			file.Close()
		case notify.Rename:
			if newFileDirPath == "" {
				newFileDirPath = path
			} else {
				grpcclient.RenameFileDirRequest(path, newFileDirPath)
				newFileDirPath = ""
			}
		}
	}
}
