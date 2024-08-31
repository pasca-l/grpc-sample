package server

import (
	"context"
	"errors"
	"io"
	"log"

	"google.golang.org/grpc"
)

func unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[interceptor] unary server interceptor: ", info.FullMethod)
	res, err := handler(ctx, req)
	return res, err
}

func streamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("[interceptor] stream server interceptor: ", info.FullMethod)
	err := handler(srv, &serverStreamWrapper{ss})
	return err
}

type serverStreamWrapper struct {
	grpc.ServerStream
}

func (s *serverStreamWrapper) RecvMsg(m interface{}) error {
	err := s.ServerStream.RecvMsg(m)
	if !errors.Is(err, io.EOF) {
		log.Println("[interceptor] intercept on receive")
	}
	return err
}

func (s *serverStreamWrapper) SendMsg(m interface{}) error {
	log.Println("[interceptor] intercept on transmission")
	return s.ServerStream.SendMsg(m)
}
