# build gateway
src_gateway="./gateway/api/proto/gen/grpc-gateway/src" && \

go build -o ./build/grpc-gateway ${src_gateway}/pkg/main/main.go && \
cp ${src_gateway}/gen/pb-go/serviceA.swagger.json ./build && \
cp ${src_gateway}/gen/pb-go/serviceB.swagger.json ./build && \
cp ./gateway/api/proto/script/config.yaml ./build && \

# build services: ServiceA, ServiceB
src_services="./service/cmd/server" && \
go build -o ./build/services ${src_services}/main.go
