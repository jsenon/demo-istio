# Service Mesh

## Demo4

### Aim

- Demonstrate delay with cutom header
- Demonstrate circuit breaking with service mesh

### Deployment

### Usage

Delay with custom header:

```sh
curl http://ac75b5b82626.eu-central-1.elb.amazonaws.com/delay/ -vv --cookie "user=julien" -H "uri: /delay/"
```

One concurent connection:

```sh
fortio load -c 1 -qps 0 -n 20 -loglevel Warning http://ac75b5b81318031881.eu-central-1.elb.amazonaws.com/outlier/healthz
```

Two concurent connection with circuit breakerm you should see some 5xx errors:

```sh
fortio load -c 2 -qps 0 -n 20 -loglevel Warning http://ac75b5b82-1318031881.eu-central-1.elb.amazonaws.com/outlier/healthz
```

Scale deployment demo-istio-v2:

```sh
kubectl scale deployment demo-istio-v2 --replicas=6
```

Outlier with 5xx error:

```sh
fortio load -H "X-Custom-Error: true" -c 1 -qps 0 -n 6 -loglevel Warning http://ac75b5b82626.eu-central-1.elb.amazonaws.com/outlier/healthz
```

You should see a number of ejection active, if not regenerate 5xx errors:

```sh
watch -n1 'curl --silent 127.0.0.1:15000/stats | grep "cluster.outbound|80||demo-istio-v2.default.svc.cluster.local.outlier_detection.ejections_active"'
```

Check logs on pod that have generate 5xx error to confirm that traffic is not sent to it.

### Notes

Remove Previous VirtualService:

```sh
istioctl delete -f virtualservice-canary.yml
```

### Todo
