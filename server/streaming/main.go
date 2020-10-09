package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_sample/proto/streaming"
	"io"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream pb.HelloService_ChannelServer) error {
	for {
		/*
			Server nhận dữ liệu được gửi từ client trong vòng lặp
		*/
		args, err := stream.Recv()
		if err != nil {
			// Nếu gặp `io.EOF`, client stream sẽ đóng.
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &pb.String{Value: "hello:" + args.GetValue()}

		/*
			Dữ liệu trả về được  gửi đến client thông qua stream
			và việc gửi nhận dữ liệu stream hai chiều là hoàn toàn độc lập
		*/
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	var host = flag.String("host", "127.0.0.1", "The server host")
	var port = flag.Int("port", 8080, "The server port")

	flag.Parse()
	fmt.Println(fmt.Sprintf("%s:%d", *host, *port))

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	grpcServer.Serve(ln)
}
