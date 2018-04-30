[![CircleCI](https://circleci.com/gh/jsenon/demo-istio.svg?style=svg)](https://circleci.com/gh/jsenon/demo-istio)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/demo-istio)](https://goreportcard.com/report/github.com/jsenon/demo-istio)

# Demo Istio
 
This is a Server in order to illustrate service mesh Istio

### Func

- Web Server with IP and Version 
- API generate 200
- API generate 503 randomly
- API generate custom answer code
- API ping/pong

## Build

Build docker container
```
make build
```

### Launch Container

```
docker run -d -p 9010:9010 -e MY_VERSION="v.x.x.x" -e MY_TARGET_PING_SVC="xxx" -e MY_TARGET_PING_PORT="80" jsenon/demo-istio:develop
```

### Usage

Generate Health, answer could be 200 or 5xx if `MY_ANSWER_CODE` is set
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

Target server is set by env `MY_TARGET_PING_SVC`/`MY_TARGET_PING_PORT`
```
curl -X POST 127.0.0.1:9010/play
```

### Env

- `MY_VERSION` = Set Application version ie v.0.0.1
- `MY_TARGET_PING_SVC` = IP of target service for ping ie. 127.0.0.1 or mysvc.mynamespace.svc.cluster.local
- `MY_TARGET_PING_PORT` = Port of target service for ping ie. 80
- `MY_JAEGER_AGENT` = Agent Jaeger ie. jaeger-agent.istio-system:5775, do not set if not needed
- `MY_SVC_SPAN_NAME` = Set Service span name
- `MY_ANSWER_CODE` = If set, define custom answer code ie. 503 (only 500, 501 and 503 are supported for now)

### TIPS

If you receive `cannot execute binary file: Exec format error`

Compile with  `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build demoserver.go`

And build `docker build -t jsenon/demo-istio:vx.x .`
