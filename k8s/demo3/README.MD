# Service Mesh

## Demo3

### Aim

- Demonstrate canary implementation with service mesh

### Deployment

### Usage

Generate 200 with different service version

```sh
curl http://ac75b5b1.eu-central-1.elb.amazonaws.com/canary/ -vv
```

### Notes

Remove Previous VirtualService

```sh
istioctl delete -f virtualservice-play.yml
```
