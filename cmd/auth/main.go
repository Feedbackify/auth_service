package main

import (
	"fmt"
	"github.com/Feedbackify/auth_service/insecure"
	authUseCase "github.com/Feedbackify/auth_service/internal/auth/application"
	authDelivery "github.com/Feedbackify/auth_service/internal/auth/delivery"
	config "github.com/Feedbackify/auth_service/internal/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	config.LoadConfig()
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer(
		// TODO: Replace certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	authUC := authUseCase.NewAuthUseCase()
	authDelivery.NewAuthHandler(s, authUC)
	fmt.Println("Serving gRPC on https://", addr)

	log.Info("Serving gRPC on https://", addr)
	log.Fatal(s.Serve(lis))
}
