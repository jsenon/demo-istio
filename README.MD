[![CircleCI](https://circleci.com/gh/jsenon/demo-istio.svg?style=svg)](https://circleci.com/gh/jsenon/demo-istio)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/demo-istio)](https://goreportcard.com/report/github.com/jsenon/demo-istio)

# Demo Istio
 
This is a Server in order to illustrate service mesh Istio

## Build

Build docker container
```
make build
```

### Launch Container

```
docker run -d -p 9010:9010 -e MY_VERSION=x.x.x jsenon/demo-istio:vX.X
```

### Func

- Web Server with IP and Version 
- API generate 200
- API generate 503 randomly
- API ping/pong

### Usage

Generate 200
```
curl -X GET 127.0.0.1:9010/healthz
```

Generate 503 randomly
```
curl -X GET --header "X-Custom-Error: true" 127.0.0.1:9010/healthz
```

Web Server
```
curl -X GET 127.0.0.1:9010
```

You can generate a ping/pong to another endpoint/server 

Defined target server by env MY_TARGET_PING_SVC/MY_TARGET_PING_PORT
```
curl -X POST 127.0.0.1:9010/play
```
You can answer to a ping from another andpoint/server by a pong

### Env

- `MY_VERSION` = Set Application version ie v.0.0.1
- `MY_TARGET_PING_SVC` = IP of target service for ping ie. 127.0.0.1 or mysvc.mynamspace.svc.cluster.local
- `MY_TARGET_PING_PORT` = Port of target service for ping ie. 9010

### TIPS

If you receive `cannot execute binary file: Exec format error`

Compile with  `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build demoserver.go`

And build `docker build -t jsenon/demo-istio:vx.x .`
