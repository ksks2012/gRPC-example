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

	client := pb.NewBidirectionalGreeterClient(conn)
	_ = SayRoute(client, &pb.HelloRequest{Request: "hong"})
}

func SayRoute(client pb.BidirectionalGreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp err: %v", resp)
	}

	_ = stream.CloseSend()

	return nil
}
