package myservice

import (
	"log"
	"strings"

	"github.com/MohitSilwal16/Picker/client/utils"
	"github.com/rjeczalik/notify"
	"github.com/spf13/viper"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
)

type MyService struct{}

func (m *MyService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
	eventlog.Remove(MY_SERVICE_NAME)
	elog, err := eventlog.Open(MY_SERVICE_NAME)
	if err != nil {
		s <- svc.Status{State: svc.StopPending}
		return false, 0
	}

	elog.Info(1, "Service Initializing ...")
	s <- svc.Status{State: svc.StartPending}
	s <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

	configFileDir := utils.GetDirOfConfigFile()
	viper.SetConfigFile(configFileDir)

	err = viper.ReadInConfig()
	if err != nil {
		elog.Error(1, "Cannot Read Config File: "+err.Error())
		s <- svc.Status{State: svc.StopPending}
		return false, 0
	}

	dirToWatch := viper.GetString("dir_to_watch_absolute_path")
	ignoreExtensionsString := viper.GetString("ignore_extensions")

	ignoreExtensions := strings.Split(ignoreExtensionsString, ",")

	c := make(chan notify.EventInfo, 1)

	err = notify.Watch(dirToWatch, c, notify.All)
	if err != nil {
		log.Println("Error:", err)
		elog.Error(1, "Cannot Read Config File: "+err.Error())
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
