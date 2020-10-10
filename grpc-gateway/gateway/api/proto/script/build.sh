go build -o ./build/grpc_gateway ../gen/grpc-gateway/src/pkg/main/main.go & \
cp ../gen/grpc-gateway/src/gen/pb-go/*swagger.json ./build & \
cp ./config.yaml ./build