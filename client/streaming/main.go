package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_sample/proto/streaming"
	"io"
	"log"
	"strconv"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &pb.String{Value: "hello room 1"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	//grpc stream
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			value := strconv.FormatInt(time.Now().UnixNano(), 10)
			if err := stream.Send(&pb.String{Value: value}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
