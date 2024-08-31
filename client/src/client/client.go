package client

import (
	pb "github.com/pasca-l/grpc-sample/client/pkg/grpc"
	"github.com/pasca-l/grpc-sample/client/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(address string) error {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// middlewares are called interceptors
		grpc.WithUnaryInterceptor(unaryClientInterceptor),
		grpc.WithStreamInterceptor(streamClientInterceptor),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewGreetingServiceClient(conn)

	// unary RPC request
	if err = usecase.Hello(client); err != nil {
		return err
	}
	// server streaming RPC request
	if err = usecase.HelloServerStream(client); err != nil {
		return err
	}
	// client streaming RPC request
	if err = usecase.HelloClientStream(client); err != nil {
		return err
	}
	// bidirectional streaming RPC request
	if err = usecase.HelloBiStreams(client); err != nil {
		return err
	}

	return nil
}
