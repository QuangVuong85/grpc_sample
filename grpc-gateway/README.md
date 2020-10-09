# grpc-gateway

- [grpc-gateway](#grpc-gateway)
  - [Overview](#Overview)
  - [Architecture Demo](#Architecture-Demo)
  - [Run](#Run)
  - [Test](#Test)

## Overview

Building grpc-gateway has caused many difficulties for developers such as the installation of the Golang programming environment, protobuf, how to build,.... There is a simpler, less time-consuming solution that uses `Build grpc-gateway`.
  
## Architecture Demo

<p align="center">
  <img src="./images/model.png"/>
</p>
  
## Test

- Test ServiceA
  
```sh
$ curl -X GET "http://localhost:9000/core/serviceA/ping/999999"
{"timestamp":"1560311912214","serviceName":"service A"}
```

- Test ServiceB

```sh
curl -X GET "http://localhost:9000/core/serviceB/ping/999999"
{"timestamp":"1560312187849","serviceName":"service B"}
```