# Service Mesh

## Demo3

### Aim

- Demonstrate canary implementation with service mesh

### Deployment

Remove Previous VirtualService:

```sh
istioctl delete -f demo2/virtualservice-play.yml
```

Deploy virtual service:

```sh
istioctl create -f demo3/virtualservice-canary.yml
```

Deploy New application v2:

```sh
kubectl apply -f demo3/service_v2.yml
kubectl apply -f demo3/deployment_v002.yml
```

Apply new virtual service:

```sh

```

### Usage

Generate 200 with different service version

```sh
curl http://ac75b5b1.eu-central-1.elb.amazonaws.com/canary/ -vv
```

Check log on both application version

### Notes
