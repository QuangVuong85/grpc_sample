package main

import (
	"context"
	cmd "grpc_sample/grpc-gateway/service/pkg/cmd/server"
	"log"
	"os"
)

func main() {
	if err := cmd.RunServer(context.Background(), "8000"); err != nil {
		log.Fatal("run server err: ", err)
		os.Exit(1)
	}
}
