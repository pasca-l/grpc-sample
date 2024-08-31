# https://grpc.io/docs/languages/go/quickstart/
FROM --platform=linux/amd64 golang:1.23-bookworm

WORKDIR /home/usr/src/

RUN apt-get update && apt-get upgrade -y

# install protocol compiler
RUN apt install -y protobuf-compiler
# install go plugins for protocol compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# install grpcurl, for testing connection via CLI
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
