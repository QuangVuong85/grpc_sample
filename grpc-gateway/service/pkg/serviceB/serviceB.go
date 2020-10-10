package serviceB

import (
	"context"
	"github.com/prometheus/common/log"
	service_proto "grpc_sample/grpc-gateway/service/pkg/api"
	"time"
)

// ServiceB struct
type ServiceB struct {
}

// NewServiceB create service
func NewServiceB() service_proto.ServiceBServer {
	return &ServiceB{}
}

// getTimestamp get timestamp
func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Ping ping
func (service *ServiceB) Ping(ctx context.Context, msgPing *service_proto.MessagePing) (*service_proto.MessagePong, error) {
	log.Info(msgPing)
	return &service_proto.MessagePong{
		Timestamp:   getTimestamp(),
		ServiceName: "Service B! Wellcome",
	}, nil
}
