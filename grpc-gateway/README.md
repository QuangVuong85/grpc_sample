# grpc-gateway

- [grpc-gateway](#grpc-gateway)
  - [Overview](#Overview)
  - [Architecture Demo](#Architecture-Demo)
  - [Run](#Run)
  - [Test](#Test)
  - [Ref](#Ref)

## Overview

Building grpc-gateway has caused many difficulties for developers such as the installation of the Golang programming environment, protobuf, how to build,.... There is a simpler, less time-consuming solution that uses `Build grpc-gateway`.
  
## Architecture Demo

<p align="center">
  <img src="./images/model.png"/>
</p>
  
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

## Ref
```link
https://github.com/grpc-ecosystem/grpc-gateway
```