# gen protoc-gen-grpc
protoc -I ../ ../*.proto \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:../gen/grpc-gateway/src/gen/pb-go/

# gen protoc-gen-grpc-gateway
protoc -I ../ --grpc-gateway_out ../gen/grpc-gateway/src/gen/pb-go/ \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt generate_unbound_methods=true \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  ../*.proto

# gen protoc-gen-swagger
protoc -I ../ --swagger_out ../gen/grpc-gateway/src/gen/pb-go/ \
  --swagger_opt logtostderr=true \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  ../*.proto