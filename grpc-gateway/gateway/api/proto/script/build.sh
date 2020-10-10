go build -o ./build/grpc_gateway ../gen/grpc-gateway/src/pkg/main/main.go && \
cp ../gen/grpc-gateway/src/gen/pb-go/serviceA.swagger.json ./build && \
cp ../gen/grpc-gateway/src/gen/pb-go/serviceB.swagger.json ./build && \
cp ./config.yaml ./build && \
cd build