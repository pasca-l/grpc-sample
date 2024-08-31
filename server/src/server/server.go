package server

import (
	"fmt"
	"log"
	"net"

	pb "github.com/pasca-l/grpc-sample/server/pkg/grpc"
	"github.com/pasca-l/grpc-sample/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve(port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		// middlewares are called interceptors
		grpc.UnaryInterceptor(unaryServerInterceptor),
		grpc.StreamInterceptor(streamServerInterceptor),
	)
	pb.RegisterGreetingServiceServer(s, service.NewGreetingServer())

	// server reflection protocol setting, for grpcurl command to work
	// sends proto file information directly from the server
	reflection.Register(s)

	log.Printf("starting gRPC server on port: %v", port)
	err = s.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
