# Service Mesh

## Demo1

### Aim

- Demonstrate tooling added by a service mesh

### Deployment



### Usage

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
