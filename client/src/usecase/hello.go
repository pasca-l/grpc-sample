package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	pb "github.com/pasca-l/grpc-sample/client/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Hello(client pb.GreetingServiceClient) error {
	log.Println("calling unary RPC")

	// set metadata from client
	ctx := context.Background()
	md := metadata.New(map[string]string{"type": "unary", "from": "client"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	// get header and trailer from server
	var header, trailer metadata.MD

	req := &pb.HelloRequest{
		Name: "unary RPC",
	}
	res, err := client.Hello(
		ctx, req, grpc.Header(&header), grpc.Trailer(&trailer),
	)
	if err != nil {
		return err
	} else {
		fmt.Println(header, trailer)
		fmt.Println(res.GetMessage())
	}
	return nil
}

func HelloServerStream(client pb.GreetingServiceClient) error {
	log.Println("calling server streaming RPC")

	req := &pb.HelloRequest{
		Name: "server streaming RPC",
	}
	stream, err := client.HelloServerStream(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("all responses have been received!")
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(res)
	}
	return nil
}

func HelloClientStream(client pb.GreetingServiceClient) error {
	log.Println("calling client streaming RPC")

	stream, err := client.HelloClientStream(context.Background())
	if err != nil {
		return err
	}

	count := 5
	for _ = range count {
		if err := stream.Send(&pb.HelloRequest{
			Name: "client streaming RPC",
		}); err != nil {
			return err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	} else {
		fmt.Println(res.GetMessage())
	}
	return nil
}

func HelloBiStreams(client pb.GreetingServiceClient) error {
	log.Println("calling bidirectional streaming RPC")

	stream, err := client.HelloBiStreams(context.Background())
	if err != nil {
		return err
	}

	count := 5
	var sendEnd, recvEnd bool
	for !(sendEnd && recvEnd) {
		// transmission
		if !sendEnd {
			for _ = range count {
				if err := stream.Send(&pb.HelloRequest{
					Name: "bidirectional streaming RPC",
				}); err != nil {
					return err
				}
			}
			sendEnd = true
			if err := stream.CloseSend(); err != nil {
				return err
			}
		}

		// reception
		if !recvEnd {
			res, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				recvEnd = true
				fmt.Println("all responses have been received!")
				break
			}
			if err != nil {
				return err
			}
			fmt.Println(res.GetMessage())
		}
	}
	return nil
}
