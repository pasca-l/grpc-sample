package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/pasca-l/grpc-sample/server/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GreetingServer struct {
	pb.UnimplementedGreetingServiceServer
}

func NewGreetingServer() *GreetingServer {
	return &GreetingServer{}
}

func (s *GreetingServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// get metadata from context from client
	if md, exist := metadata.FromIncomingContext(ctx); exist {
		log.Println(md)
	}

	// set header and trailer from server
	headerMd := metadata.New(
		map[string]string{"type": "unary", "from": "server", "in": "header"},
	)
	if err := grpc.SetHeader(ctx, headerMd); err != nil {
		return nil, err
	}
	trailerMd := metadata.New(
		map[string]string{"type": "unary", "from": "server", "in": "trailer"},
	)
	if err := grpc.SetTrailer(ctx, trailerMd); err != nil {
		return nil, err
	}

	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func (s *GreetingServer) HelloServerStream(req *pb.HelloRequest, stream pb.GreetingService_HelloServerStreamServer) error {
	count := 5
	for i := range count {
		if err := stream.Send(&pb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func (s *GreetingServer) HelloClientStream(stream pb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(&pb.HelloResponse{
				Message: fmt.Sprintf("Hello, %v!", nameList),
			})
		}
		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}

func (s *GreetingServer) HelloBiStreams(stream pb.GreetingService_HelloBiStreamsServer) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&pb.HelloResponse{
			Message: fmt.Sprintf("Hello, %v!", req.GetName()),
		}); err != nil {
			return err
		}
	}
}
