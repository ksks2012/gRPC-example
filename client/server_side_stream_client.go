package main

import (
	"context"
	"flag"
	"io"
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

	client := pb.NewMachineGreeterClient(conn)
	_ = SayList(client, &pb.HelloRequest{Request: "hong"})
}

func SayList(client pb.MachineGreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	return nil
}
