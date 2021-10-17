package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/start-grpc/proto"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello.world"}, nil
}

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "啟動通訊埠編號")
	flag.Parse()
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
