# Service Mesh

## Demo4

### Aim

- Demonstrate delay with cutom header
- Demonstrate circuit breaking with service mesh

### Deployment

Remove Previous VirtualService:

```sh
istioctl delete -f demo3/virtualservice-canary.yml
```

Deploy virtual service:

```sh
istioctl create -f demo4/virtualservice-delay.yml
```

Deploy outlier:

```sh
istioctl create -f demo4/outlier.yml
```

Be sure to have deployed v002:

```sh
kubectl apply -f demo3/service_v2.yml
kubectl apply -f demo3/deployment_v002.yml
```

### Usage

Delay with custom header:

```sh
curl http://<REDACTED>/delay/ -vv --cookie "user=julien" -H "uri: /delay/"
```

One concurent connection:

```sh
fortio load -c 1 -qps 0 -n 20 -loglevel Warning http://<REDACTED>/outlier/healthz
```

Two concurent connection with circuit breakerm you should see some 5xx errors:

```sh
fortio load -c 2 -qps 0 -n 20 -loglevel Warning http://<REDACTED>/outlier/healthz
```

Scale deployment demo-istio-v2:

```sh
kubectl scale deployment demo-istio-v2 --replicas=6
```

Outlier with 5xx error:

```sh
fortio load -H "X-Custom-Error: true" -c 1 -qps 0 -n 20 -loglevel Warning http://<REDACTED>/outlier/healthz
```

Observer envoy admin page for istio-ingressgateway:

```sh
kubectl port-forward -n istio-system istio-ingressgateway-57f76dc4db-j9nwd 15000:15000
```

You should see a number of ejection active, if not regenerate 5xx errors:

```sh
watch -n1 'curl --silent 127.0.0.1:15000/stats | grep "cluster.outbound|80||demo-istio-v2.default.svc.cluster.local.outlier_detection.ejections_active"'
```

Check logs on pod that have generate 5xx error to confirm that traffic is not sent to it.

### Notes

If Further information is needed [Outlier explanation](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/outlier.html)could be read on Envoy website.
[Outlier detection statistics](https://www.envoyproxy.io/docs/envoy/latest/configuration/cluster_manager/cluster_stats#outlier-detection-statistics) is also available on Envoy site.

### Todo
