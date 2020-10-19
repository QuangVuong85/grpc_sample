protoc -I ../api/proto/ \
  ../api/proto/*.proto \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:../pkg/api/
