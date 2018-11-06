# Service Mesh

## Demo8

### Aim

- Demonstrate External communication

### Before Deployment

Check communication before deployment

```sh
kubectl exec  toolingcontainer-57d8869654-q45mz  -it -- curl http://httpbin.org -vvv
```

You will see:

```sh
* Rebuilt URL to: http://httpbin.org/
*   Trying 52.1.65.199...
* TCP_NODELAY set
* Connected to httpbin.org (52.1.65.199) port 80 (#0)
> GET / HTTP/1.1
> Host: httpbin.org
> User-Agent: curl/7.61.1
> Accept: */*
>
< HTTP/1.1 404 Not Found
< date: Tue, 06 Nov 2018 21:34:16 GMT
< server: envoy
< content-length: 0
<
* Connection #0 to host httpbin.org left intact
```

With tls:

```sh
kubectl exec  toolingcontainer-57d8869654-q45mz  -it -- curl https://httpbin.org -vvv
```

You will see:

```sh
* Rebuilt URL to: https://httpbin.org/
*   Trying 52.4.95.48...
* TCP_NODELAY set
* Connected to httpbin.org (52.4.95.48) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: none
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* LibreSSL SSL_connect: SSL_ERROR_SYSCALL in connection to httpbin.org:443
curl: (35) LibreSSL SSL_connect: SSL_ERROR_SYSCALL in connection to httpbin.org:443
command terminated with exit code 35
```

### Deployment

Create ServiceEntry and VirtualService:

```sh
istioctl create -f external.yaml
```

### Validation

Check that you are able to contact `http://httpbin.org`

```sh
kubectl exec  toolingcontainer-57d8869654-q45mz  -it -- curl http://httpbin.org -vvv
```

Check that you are able to contact `https://www.google.com`

```sh
kubectl exec  toolingcontainer-57d8869654-q45mz  -it -- curl https://www.google.com -vvv
```

Apply new ServiceEntry for `https://httpbin.org`

```sh
istioctl create -f external-bis.yaml
```

And check that you are able to contact `https://httpbin.org`

```sh
kubectl exec  toolingcontainer-57d8869654-q45mz  -it -- curl https://httpbin.org -vvv
```