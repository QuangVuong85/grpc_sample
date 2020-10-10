package cmd

import (
	"context"
	service_proto "grpc_sample/grpc-gateway/service/pkg/api"
	"grpc_sample/grpc-gateway/service/pkg/serviceA"
	"grpc_sample/grpc-gateway/service/pkg/serviceB"
	"net"
	"os"
	"os/signal"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

// RunServer run server
func RunServer(ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// init services
	servA := serviceA.NewServiceA()
	servB := serviceB.NewServiceB()

	// init server gRPC
	server := grpc.NewServer()

	// register service
	service_proto.RegisterServiceAServer(server, servA)
	service_proto.RegisterServiceBServer(server, servB)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Info("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Info("Start service A, service B port " + port + " ...")
	return server.Serve(listen)
}
