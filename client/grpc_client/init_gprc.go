package grpcclient

import (
	"github.com/MohitSilwal16/Picker/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var fileWatcherClient pb.FileWatcherClient
var authClient pb.AuthClient

func NewGRPCClients(address string) error {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	fileWatcherClient = pb.NewFileWatcherClient(conn)
	authClient = pb.NewAuthClient(conn)

	return nil
}
