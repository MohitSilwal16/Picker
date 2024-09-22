package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/MohitSilwal16/Picker/server/handler"
	"github.com/MohitSilwal16/Picker/server/pb"
	"github.com/MohitSilwal16/Picker/server/utils"
	"google.golang.org/grpc"
)

const BASE_URL = "0.0.0.0:8080"

func main() {
	utils.ClearScreen()

	lis, err := net.Listen("tcp", BASE_URL)
	if err != nil {
		log.Println("Error:", err)
		log.Println("Description: Cannot Listen TCP at", BASE_URL)
		return
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(utils.StructuredLoggerInterceptor()),
	)

	pb.RegisterFileWatcherServer(s, &handler.FileWatcherServer{})

	log.Println("Running gPRC Server on", lis.Addr())

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Println("Error:", err)
			log.Println("Description: Cannot Serve at", BASE_URL)
			return
		}
	}()

	var choice string

	for {
		fmt.Scanln(&choice)

		switch choice {
		case "h":
			fmt.Println("Commands:")
			fmt.Println("h - Show this help message")
			fmt.Println("c - Clear the screen")
			fmt.Println("q - Quit the program")
		case "c":
			utils.ClearScreen()
		case "q":
			fmt.Println("Terminating Server ...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Enter 'h' for help.")
		}
	}
}
