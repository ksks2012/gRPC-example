package main

import (
	"flag"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/start-grpc/proto"
)

type ReverseMachineGreeterServer struct{}

func (s *ReverseMachineGreeterServer) SayRecord(stream pb.ReverseMachineGreeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			message := &pb.HelloReply{Message: "say.record"}
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
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
	pb.RegisterReverseMachineGreeterServer(server, &ReverseMachineGreeterServer{})
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
