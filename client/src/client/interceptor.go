package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func unaryClientInterceptor(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("[interceptor] unary client interceptor: ", method, req)
	err := invoker(ctx, method, req, res, cc, opts...)
	return err
}

func streamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	log.Println("[interceptor] stream client interceptor: ", method)
	stream, err := streamer(ctx, desc, cc, method, opts...)
	return &clientStreamWrapper{stream}, err
}

type clientStreamWrapper struct {
	grpc.ClientStream
}

func (s *clientStreamWrapper) SendMsg(m interface{}) error {
	log.Println("[interceptor] intercept on transmission")
	return s.ClientStream.SendMsg(m)
}

func (s *clientStreamWrapper) RecvMsg(m interface{}) error {
	err := s.ClientStream.RecvMsg(m)
	if !errors.Is(err, io.EOF) {
		log.Println("[interceptor] intercept on receive")
	}
	return err
}
