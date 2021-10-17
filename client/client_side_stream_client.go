package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	pb "github.com/start-grpc/proto"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "啟動通訊埠編號")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewReverseMachineGreeterClient(conn)
	_ = SayRecord(client, &pb.HelloRequest{Request: "hong"})
}

func SayRecord(client pb.ReverseMachineGreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
	}

	resp, _ := stream.CloseAndRecv()
	log.Printf("resp: %v", resp)
	return nil
}
