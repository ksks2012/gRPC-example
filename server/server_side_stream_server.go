package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/start-grpc/proto"
)

type MachineGreeterServer struct{}

func (s *MachineGreeterServer) SayList(r *pb.HelloRequest, stream pb.MachineGreeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "hello.list"})
	}
	return nil
}

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "啟動通訊埠編號")
	flag.Parse()
}

func main() {
	server := grpc.NewServer()
	pb.RegisterMachineGreeterServer(server, &MachineGreeterServer{})
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
