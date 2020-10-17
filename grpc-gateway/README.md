# grpc-gateway

- [grpc-gateway](#grpc-gateway)
  - [Overview](#Overview)
  - [Architecture Demo](#Architecture-Demo)
  - [Run](#Run)
  - [Test](#Test)
  - [Tree project](#Tree-project)
  - [Ref](#Ref)
  - [Lib](#L)

## Overview

Building grpc-gateway has caused many difficulties for developers such as the installation of the Golang programming environment, protobuf, how to build,.... There is a simpler, less time-consuming solution that uses `Build grpc-gateway`.
  
## Architecture Demo
![](./images/kientruc.png)
  
## Run
 - Build use script bash:
 ```shell script
    ./build.sh
 ```
 - Build use script bash:
 ```shell script
    Step 1: make
    Step 2: make run-gateway
    Step 3 (open terminal other): make run-services
 ```
 
 - Run grpc-gateway
 ```shell script
    cd ./build
    ./grpc-gateway
 ```
 - Run services
 ```shell script
    cd ./build
    ./services
 ```

## Test

- Test ServiceA
  
```sh
curl -X GET "http://localhost:9000/core/serviceA/ping/999"
{"timestamp":"1602319774563","serviceName":"Service A! Wellcome"}
```

- Test ServiceB

```sh
curl -X GET "http://localhost:9000/core/serviceB/ping/999"
{"timestamp":"1602319782972","serviceName":"Service B! Wellcome"}
```

### Tree project
```
├── build.sh
├── client
│   ├── math
│   │   └── main.go
│   ├── pubsubservice
│   │   ├── clientpub
│   │   │   └── main.go
│   │   └── clientsub
│   │       └── main.go
│   ├── streaming
│   │   └── main.go
│   └── user
│       └── main.go
├── grpc-gateway
│   ├── build
│   │   ├── config.yaml
│   │   ├── grpc-gateway
│   │   ├── serviceA.swagger.json
│   │   ├── serviceB.swagger.json
│   │   └── services
│   ├── build.sh
│   ├── gateway
│   │   └── api
│   │       └── proto
│   │           ├── common.proto
│   │           ├── gen
│   │           │   └── grpc-gateway
│   │           │       └── src
│   │           │           ├── gen
│   │           │           │   └── pb-go
│   │           │           │       ├── common.pb.go
│   │           │           │       ├── common.swagger.json
│   │           │           │       ├── serviceA.pb.go
│   │           │           │       ├── serviceA.pb.gw.go
│   │           │           │       ├── serviceA.swagger.json
│   │           │           │       ├── serviceB.pb.go
│   │           │           │       ├── serviceB.pb.gw.go
│   │           │           │       └── serviceB.swagger.json
│   │           │           └── pkg
│   │           │               └── main
│   │           │                   └── main.go
│   │           ├── script
│   │           │   ├── config.yaml
│   │           │   └── gen.sh
│   │           ├── serviceA.proto
│   │           └── serviceB.proto
│   ├── images
│   │   └── model.png
│   ├── Makefile
│   ├── README.md
│   └── service
│       ├── api
│       │   └── proto
│       │       ├── common.proto
│       │       ├── serviceA.proto
│       │       └── serviceB.proto
│       ├── cmd
│       │   └── server
│       │       └── main.go
│       ├── pkg
│       │   ├── api
│       │   │   ├── common.pb.go
│       │   │   ├── serviceA.pb.go
│       │   │   └── serviceB.pb.go
│       │   ├── cmd
│       │   │   └── server
│       │   │       └── server.go
│       │   ├── serviceA
│       │   │   └── serviceA.go
│       │   ├── serviceB
│       │   │   └── serviceB.go
│       │   └── utils
│       │       └── getlocalIP.go
│       └── script
│           └── gen.sh
├── proto
│   ├── math
│   │   └── math.pb.go
│   ├── math.proto
│   ├── pubsubservice
│   │   └── pubsubservice.pb.go
│   ├── pubsubservice.proto
│   ├── streaming
│   │   └── streaming.pb.go
│   ├── streaming.proto
│   ├── user
│   │   └── user.pb.go
│   └── user.proto
├── README.md
├── server
│   ├── math
│   │   └── main.go
│   ├── pubsubservice
│   │   ├── main.go
│   │   └── pubsub
│   │       └── publisher.go
│   ├── streaming
│   │   └── main.go
│   └── user
│       └── main.go
└── utils
    └── executor
        ├── executor.go
        ├── README.md
        └── test
            └── executor_test.go
```

## Ref
```link
https://github.com/grpc-ecosystem/grpc-gateway
```
  
## Lib
```
go get -u github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/gengo/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```