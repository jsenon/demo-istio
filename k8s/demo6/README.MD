# Service Mesh

## Demo6

### Aim

- Demonstrate mTLS
- Demonstrate SecurityPolicy

mTLS could be implemented on:

- cluster side
- namespace
- or dedicated service

In order to demonstrate mTLS we will use mTLS on dedicated service and namespace.

### mTLS

#### mTLS Deployment

Remove Previous VirtualService:

```sh
istioctl delete -f demo5/virtualservice-mirror.yml
```

Create new ns with istio:

```sh
kubectl create ns mtls
kubectl label namespace mtls istio-injection=enabled
```

Deploy service, and application on new namespace:

```sh
kubectl apply -f demo6/service_v2.yml
kubectl apply -f demo6/deployment_v002.yml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml -n mtls
```

Create new ns without istio:

```sh
kubectl create ns legacy
```

Deploy Sleep application on legacy namespace:

```sh
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml -n legacy
```

Deploy Sleep application on default namesapce:

```sh
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/sleep/sleep.yaml -n default
```

#### mTLS Usage

Perform First test:

```sh
for from in "default" "mtls" "legacy"; do for to in "default" "mtls"; do kubectl exec $(kubectl get pod -l app=sleep -n ${from} -o jsonpath={.items..metadata.name}) -c sleep -n ${from} -- curl --silent http://demo-istio-v2.${to}/healthz  -o /dev/null -w "sleep.${from} to demo-istio-v2.${to}: %{http_code}\n"; done; done
```

You will see, that http_code are 200:

```sh
sleep.default to demo-istio-v2.default: 200
sleep.default to demo-istio-v2.mtls: 200
sleep.mtls to demo-istio-v2.default: 200
sleep.mtls to demo-istio-v2.mtls: 200
sleep.legacy to demo-istio-v2.default: 200
sleep.legacy to demo-istio-v2.mtls: 200
```

Check policies:

```sh
kubectl get policies.authentication.istio.io -n default
kubectl get policies.authentication.istio.io -n mtls
```

Now enforce mtls namespace with mTLS:

```sh
istioctl create -f demo6/policy_mtls.yml
```

Check if policy is correctly applied:

```sh
kubectl get policies.authentication.istio.io -n mtls
```

Perform First test:

```sh
for from in "default" "mtls" "legacy"; do for to in "default" "mtls"; do kubectl exec $(kubectl get pod -l app=sleep -n ${from} -o jsonpath={.items..metadata.name}) -c sleep -n ${from} -- curl --silent http://demo-istio-v2.${to}/healthz  -o /dev/null -w "sleep.${from} to demo-istio-v2.${to}: %{http_code}\n"; done; done
```

You will see legacy namespace without istio can't contact service with istio and mTLS activate (`sleep.legacy to demo-istio-v2.mtls: 000`):

```sh
sleep.default to demo-istio-v2.default: 200
sleep.default to demo-istio-v2.mtls: 200
sleep.mtls to demo-istio-v2.default: 200
sleep.mtls to demo-istio-v2.mtls: 200
sleep.legacy to demo-istio-v2.default: 200
sleep.legacy to demo-istio-v2.mtls: 000
command terminated with exit code 56
```

### Basic Access

If you acces demo-istio svc with /play, service demo-istio will contact demo-istio-pong svc. In order to demonstrate basic access we will deny access from demo-istio to demo-istio-pong.

#### Basic Access Deployment

Remove previous virtualservice:

```sh
istioctl delete -f demo5/virtualservice-mirror.yml
```

Apply previously virtualservice with /play:

```sh
istioctl create -f demo2/virtualservice-play.yml
```

#### Basic Access Usage

Check if ping/pong works fine:

```sh
curl http://aee0fc8d0c6f711e8a24d02c2dbb86c0-1989099266.eu-central-1.elb.amazonaws.com/play -vv
...
< HTTP/1.1 200 OK
< date: Wed, 30 May 2018 11:43:41 GMT
< content-length: 10
< content-type: text/plain; charset=utf-8
< x-envoy-upstream-service-time: 11
< server: envoy
<
* Connection #0 to host ac75-131881.eu-central-1.elb.amazonaws.com left intact
200 - Pong%
```

Apply denier:

```sh
istioctl create -f demo6/denier.yml
```

Check  ping/pong result:

```sh
curl http://aee0fc8d0c6f711e8a24d02c2dbb86c0-1989099266.eu-central-1.elb.amazonaws.com/play -vv
...
< HTTP/1.1 200 OK
< date: Wed, 30 May 2018 11:48:10 GMT
< content-length: 53
< content-type: text/plain; charset=utf-8
< x-envoy-upstream-service-time: 11
< server: envoy
<
* Connection #0 to host ac75b-81.eu-central-1.elb.amazonaws.com left intact
PERMISSION_DENIED:denypong.denier.default:Not allowed%
```

Remove denier:

```sh
istioctl delete -f demo6/denier.yml
```

### Notes

Issue [open](https://github.com/istio/istio/issues/5909) regarding force destination to *.mtls to be enforce with MUTUAL

### Todo