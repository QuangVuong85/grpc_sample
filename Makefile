all: client server

serverpath = ${GOPATH}/src/grpc_sample/server/streaming
clientpath = ${GOPATH}/src/grpc_sample/client/streaming

protoc:
	@echo "Generating Go files..."
	protoc -I ./proto ./proto/math.proto --go_out=plugins=grpc:proto/math && \
    protoc -I ./proto ./proto/user.proto --go_out=plugins=grpc:proto/user && \
    protoc -I ./proto ./proto/streaming.proto --go_out=plugins=grpc:proto/streaming && \
    protoc -I ./proto ./proto/pubsubservice.proto --go_out=plugins=grpc:proto/pubsubservice && \
    protoc -I ./proto ./proto/chat.proto --go_out=plugins=grpc:proto/chat

server: protoc
	@echo "Building server streaming..."
	go build -o $(serverpath)/server-streaming \
		$(serverpath)

client: protoc
	@echo "Building client streaming..."
	go build -o $(clientpath)/client-streaming \
    		$(clientpath)

clean:
	rm -f $(serverpath)/server-streaming
	rm -f $(clientpath)/client-streaming

.PHONY: client server protoc