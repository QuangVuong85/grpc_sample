package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_sample/proto/math"
	"log"
	"net"
	"os"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *math.Request) (*math.SumResponse, error) {
	a := req.GetA()
	b := req.GetB()

	sum := a + b

	return &math.SumResponse{
		Sum: sum,
	}, nil
}

func (s *server) Mul(ctx context.Context, req *math.Request) (*math.MulResponse, error) {
	a := req.GetA()
	b := req.GetB()

	mul := a * b

	return &math.MulResponse{
		Mul: mul,
	}, nil
}

func main() {
	ln, err := net.Listen("tcp",
		fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	math.RegisterMathServer(s, &server{})

	log.Println("Starting GRPC server...\n" +
		"\t- GRPC_PORT=50001 go run main.go")

	err = s.Serve(ln)

	if err != nil {
		log.Fatal(err)
	}
}
