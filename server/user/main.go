package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_sample/proto/user"
	"log"
	"net"
)

type server struct {
	request    []*pb.GetRequest
	savedUsers []*pb.Model
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	ok, _ := s.FindUser(ctx, &pb.FindRequest{Id: req.User.Id})
	if ok.User != nil {
		return &pb.CreateResponse{
			Status:  "Fail exist user id=" + req.User.Id,
			Message: req.User.Id,
		}, nil
	}

	s.savedUsers = append(s.savedUsers, req.User)

	return &pb.CreateResponse{
		Status:  "Create success user id=" + req.User.Id,
		Message: req.User.Id,
	}, nil
}

func (s *server) FindUser(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	for _, user := range s.savedUsers {
		if req.Id != "" {
			if req.Id == user.Id {
				return &pb.FindResponse{
					Status: "Find success user id=" + req.Id,
					User:   user,
				}, nil
			}
		}
	}

	return &pb.FindResponse{
		Status: "Find fail user id=" + req.Id,
		User:   nil,
	}, nil
}

func (s *server) GetUsers(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	count := len(s.savedUsers)

	users := &pb.GetResponse{
		Status: "Get all users success!",
		Users:  s.savedUsers,
		Count:  int32(count),
	}

	return users, nil
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

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	s.Serve(ln)
}
