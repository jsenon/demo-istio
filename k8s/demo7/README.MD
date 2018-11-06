# Service Mesh

## Demo7

### Aim

- Demonstrate RateLimit

### Deployment

Delete previous virtualservice created:

```sh
istioctl delete virtualservice xxx
```

Delete any destination rule:

```sh
istioctl delete DestinationRule xxx
```

Create virtual service:

```sh
istioctl create -f demo7/virtualservice_rate.yml
```

Create rate limit:

```sh
istioctl create -f demo7/quotas.yml
```

### Usage

Force 1 connexion:

```sh
while true; do curl -s -o /dev/null http://aee0fc8d0c6f711e8a24d02c2dbb86c0-1989099266.eu-central-1.elb.amazonaws.com/rate/; done
```

Try another connexion:

```sh
curl http://aee0fc8d0c6f711e8a24d02c2dbb86c0-1989099266.eu-central-1.elb.amazonaws.com/rate/ -v
```

### Notes

## Draft - Not working

Check relevant [issue](https://github.com/istio/istio/issues/5692) to check evolution.

### Todo