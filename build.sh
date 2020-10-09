# golang
protoc -I ./proto ./proto/math.proto --go_out=plugins=grpc:proto/math
protoc -I ./proto ./proto/user.proto --go_out=plugins=grpc:proto/user
protoc -I ./proto ./proto/streaming.proto --go_out=plugins=grpc:proto/streaming
protoc -I ./proto ./proto/pubsubservice.proto --go_out=plugins=grpc:proto/pubsubservice
