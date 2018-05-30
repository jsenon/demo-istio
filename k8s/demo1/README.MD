# Service Mesh

## Demo1

### Aim

- Demonstrate tooling added by a service mesh
- Demonstrate sidecar automatic injection

### Deployment

Deploy service 200:

```sh
kubectl apply -f demo1/service_200.yml
```

Deploy application 200:

```sh
kubectl apply -f demo1/deployment_200.yml
```

Deploy service 503:

```sh
kubectl apply -f demo1/service_503.yml
```

Deploy application 503:

```sh
kubectl apply -f demo1/deployment_503.yml
```

Deploy service 200 w/o span:

```sh
kubectl apply -f demo1/service_200_stdspan.yml
```

Deploy application 200 w/o span:

```sh
kubectl apply -f demo1/deployment_200_stdspan.yml
```

Deploy virtual service:

```sh
istioctl create -f demo1/virtualservice.yml
```

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

Observe Grafana metrics, in your browser `http://127.0.0.1:3000` by using port-forward:

```sh
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 3000:3000 &
```

Observe Opentracing, in your browser `http://127.0.0.1:16686` by using port-forward:

```sh
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=jaeger -o jsonpath='{.items[0].metadata.name}') 16686:16686 &
```

Observer Service Graph, in your browser `http://127.0.0.1:8088/force/forcegraph.html` by using port-forward:

```sh
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=servicegraph -o jsonpath='{.items[0].metadata.name}') 8088:8088 &
```