package myservice

import (
	"log"
	"os"
	"strings"

	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/joho/godotenv"
	"github.com/rjeczalik/notify"
	"golang.org/x/sys/windows/svc"
)

type MyService struct{}

func (m *MyService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
	s <- svc.Status{State: svc.StartPending}
	s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

	configPath := utils.GetPathOfConfigFile()

	err := godotenv.Load(configPath)
	if err != nil {
		s <- svc.Status{State: svc.StopPending}
		return false, 0
	}

	dirToWatch := os.Getenv("DIR_TO_WATCH_ABSOLUTE_PATH")
	ignoreExtensions := strings.Split(os.Getenv("IGNORE_EXTENSIONS"), ",")
	logFilePath := os.Getenv("LOG_FILE_ABSOLUTE_PATH")

	logFile, err := utils.SetUpFileLogging(logFilePath)
	if err != nil {
		s <- svc.Status{State: svc.StopPending}
		return false, 0
	}
	defer logFile.Close()
	log.Println("Initializing Service ...")

	c := make(chan notify.EventInfo, 1)

	err = notify.Watch(dirToWatch, c, notify.All)
	if err != nil {
		log.Println("Error:", err)
		s <- svc.Status{State: svc.StopPending}
		return false, 0
	}
	defer notify.Stop(c)

	dirToWatch = dirToWatch[:len(dirToWatch)-3]
	go watchChanges(c, ignoreExtensions, dirToWatch)

	log.Println("Starting Service ...")
	for {
		change := <-r
		switch change.Cmd {
		case svc.Interrogate:
			s <- change.CurrentStatus
		case svc.Stop, svc.Shutdown:
			log.Println("Stopping Service ...")
			s <- svc.Status{State: svc.StopPending}
			return false, 0
		}
	}
}
