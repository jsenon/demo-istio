# Service Mesh

## Demo2

### Aim

- Demonstrate microservice chaining with service mesh

### Deployment

Remove Previous VirtualService

```sh
istioctl delete -f demo1/virtualservice.yml
```

Deploy application ping:

```sh
kubectl apply -f demo2/service.yml
kubectl apply -f demo2/deployment_v001.yml
```

Deploy application pong:

```sh
kubectl apply -f demo2/service_pong.yml
kubectl apply -f demo2/deployment_pong_srv.yml
```

Apply new virtual sercvice:

```sh
istioctl create -f demo2/virtualservice-play.yml
```

### Usage

Generate Ping-Pong

```sh
curl http://<REDACTED>/play -vv
```

Check rendering on Jaegger, service graph.

### Notes