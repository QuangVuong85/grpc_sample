package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_sample/proto/user"
	"log"
)

var (
	c   pb.UserClient
	ctx = context.Background()
)

func init() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s",
			"localhost", "8080",
		),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatal(err)
	}

	c = pb.NewUserClient(conn)
}

func main() {
	resp_c, _ := c.CreateUser(ctx, &pb.CreateRequest{
		User: &pb.Model{
			Id:   "1",
			Name: "grpc 01",
		},
	})
	log.Println(resp_c)

	find_1, _ := c.FindUser(ctx, &pb.FindRequest{Id: "1"})
	log.Println(find_1)

	find_2, _ := c.FindUser(ctx, &pb.FindRequest{Id: "2"})
	log.Println(find_2)

	getall, _ := c.GetUsers(ctx, &pb.GetRequest{})
	log.Println(getall)
}
