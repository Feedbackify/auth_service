package main

import (
	"fmt"
	"github.com/Feedbackify/auth_service/insecure"
	server "github.com/Feedbackify/auth_service/internal/auth/delivery"
	authv1 "github.com/Feedbackify/auth_service/proto/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net"
	"os"
)

func main() {
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
	authv1.RegisterAuthServiceServer(s, server.New())
	fmt.Println("Serving gRPC on https://", addr)

	s.Serve(lis)
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()
}
