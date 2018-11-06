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
while true; do curl -s -o /dev/null http://<REDACTED>/rate/; done
```

Try another connexion:

```sh
curl http://<REDACTED>/rate/ -v
```

### Notes

## Draft - Not working

Check relevant [issue](https://github.com/istio/istio/issues/5692) to check evolution.

### Todo