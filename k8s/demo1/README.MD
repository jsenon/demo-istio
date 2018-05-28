# Service Mesh

## Demo1

### Aim

- Demonstrate tooling added by a service mesh
- Demonstrate sidecar automatic injection

### Deployment

### Usage

Observe automnatic sidecar injection

```sh
kubectl get pods
```

Inside PODS

```sh
kubectl exec -it demo-istio-200-6cb68774c8-bs67w -c istio-proxy bash
$ ps -e
```

```sh
kubectl describe pods demo-istio-200-6cb68774c8-bs67w
```

Generate 200

```sh
curl http://ac75b5b1.eu-central-1.elb.amazonaws.com/200/healthz -vv
```

Generate 503

```sh
curl http://ac75b5b1.eu-central-1.elb.amazonaws.com/503/healthz -vv
```

Generate 200 w/o span

```sh
curl http://ac75b5b1.eu-central-1.elb.amazonaws.com/stdspan/healthz -vv
```
