# Service Mesh

## Demo11

### Aim

- Demonstrate OPA Agent and Istio Mixer

### Deployment

Clean Previous Demo

```sh
kubectl delete -f demo1/
kubectl delete -f demo2/
kubectl delete -f demo3/
kubectl delete -f demo4/
kubectl delete -f demo5/
kubectl delete -f demo6/
kubectl delete -f demo7/
kubectl delete -f demo8/
kubectl delete -f demo9/
kubectl delete -f demo10/

```

Deploy application v1:

```sh
kubectl apply -f demo11/service.yml
kubectl apply -f demo11/deployment_v001.yml
```

Deploy application v2:

```sh
kubectl apply -f demo11/service_v2.yml
kubectl apply -f demo11/deployment_v002.yml
```

Apply new virtual sercvice:

```sh
istioctl create -f demo11/virtualservice.yml
```

### Istio Configuration

Apply istio rules policy:

```sh
istioctl create -f demo11/rules.yml
```

### OPA Policy

Apply OPA Policy

```sh
istioctl create -f demo11/opa.yml
```

Wait few seconds before access to your api

### Access to your api

We have defined Policy that allow method `PUT` but denied method `GET`

#### Method PUT

Check that you can access to the api with `PUT` method:

```sh
curl -X PUT <REDACTED>.us-east-2.elb.amazonaws.com/v2
```

#### Method GET

Check that you can't have access to the api with `GET` method

```sh
curl -X PUT <REDACTED>.us-east-2.elb.amazonaws.com/v2
```

you will see

```sh
PERMISSION_DENIED:opa-handler.opa.istio-system:opa: request was rejected^C%
```