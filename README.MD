[![CircleCI](https://circleci.com/gh/jsenon/demo-istio.svg?style=svg)](https://circleci.com/gh/jsenon/demo-istio)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/demo-istio)](https://goreportcard.com/report/github.com/jsenon/demo-istio)

# Demo Istio
 
This is a Server in order to illustrate service mesh Istio

## Build

Build version
```
go build demoserver.go
```

Check name of your binairies against name used in dockerfile


### Create Container

```
docker build -t jsenon/demo-istio:vX.X .
```

### Launch Container

```
docker run -d -p 80:9010 jsenon/demo-istio:vX.X
```

### Func

- Web Server with IP and Version 
- API generate 200
- API generate 503 randomly

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

### TIPS

If you receive `cannot execute binary file: Exec format error`

Compile with  `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build demoserver.go`

And build `docker build -t jsenon/demo-istio:vx.x .`
