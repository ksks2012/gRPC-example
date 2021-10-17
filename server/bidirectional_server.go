package main

import (
	"flag"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/start-grpc/proto"
)

type BidirectionalGreeter struct{}

func (s *BidirectionalGreeter) SayRoute(stream pb.BidirectionalGreeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "say.route"})

		resp, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("resp: %v", resp)
	}
}

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "啟動通訊埠編號")
	flag.Parse()
}

func main() {
	server := grpc.NewServer()
	pb.RegisterBidirectionalGreeterServer(server, &BidirectionalGreeter{})
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
