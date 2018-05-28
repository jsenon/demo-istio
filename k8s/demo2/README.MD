# Service Mesh

## Demo2

### Aim

- Demonstrate microservice chaining with service mesh

### Deployment

### Usage

Generate Ping-Pong

```sh
curl http://ac75b5b1.eu-central-1.elb.amazonaws.com/play -vv
```

### Notes

Remove Previous VirtualService

```sh
istioctl delete -f virtualservice.yml
```

