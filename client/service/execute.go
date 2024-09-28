package service

import (
	"log"
	"strings"

	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/rjeczalik/notify"
	"github.com/spf13/viper"
	"golang.org/x/sys/windows/svc"
)

type MyService struct{}

func (m *MyService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
	s <- svc.Status{State: svc.StartPending}
	s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

	configFilePath := utils.GetPathOfConfigFile()
	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		s <- svc.Status{State: svc.StopPending}
		return false, 0
	}

	dirsToWatchString := viper.GetString("dirs_to_watch_absolute_path")
	allowedExtensionsString := viper.GetString("allowed_extensions")

	allowedExtensions := strings.Split(allowedExtensionsString, ";")
	dirsToWatch := strings.Split(dirsToWatchString, ";")

	for idx := range dirsToWatch {
		c := make(chan notify.EventInfo, 1)
		allowedExtensForParticularDir := strings.Split(allowedExtensions[idx], ",")

		err = notify.Watch(dirsToWatch[idx], c, notify.All)
		if err != nil {
			log.Println("Error:", err)
			s <- svc.Status{State: svc.StopPending}
			return false, 0
		}
		defer notify.Stop(c)

		dirsToWatch[idx] = dirsToWatch[idx][:len(dirsToWatch[idx])-3]
		go watchChanges(c, allowedExtensForParticularDir, dirsToWatch[idx])
	}

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
