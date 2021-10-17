package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	pb "github.com/start-grpc/proto"
)

type GreeterClient struct{}

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "啟動通訊埠編號")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Request: "hong"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
