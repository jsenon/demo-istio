# Service Mesh

## Demo9

### Aim

- Demonstrate End-User authentication using JSON Web Token and Auth0

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
```

Deploy application v1:

```sh
kubectl apply -f demo9/service.yml
kubectl apply -f demo9/deployment_v001.yml
```

Deploy application v2:

```sh
kubectl apply -f demo9/service_v2.yml
kubectl apply -f demo9/deployment_v002.yml
```

Apply new virtual sercvice:

```sh
istioctl create -f demo9/virtualservice.yml
```
### Auth0 Configuration

Create a [free account](https://auth0.com/) if you don't have one.

Dashboard:

![alt text](img/dash.png)

Create new API:

![alt text](img/api1.png)

Complete information regarding your url api ie. http://<REDACTED>.amazonaws.com/v2 and select signing alrorithm to `rs256`:
![alt text](img/api2.png)

You should see:

![alt text](img/api3.png)

In your api define scope:

![alt text](img/scope.png)

Select application:

![alt text](img/appli1.png)

Create a new m2m application wich have access to api /v2 only:

![alt text](img/appli2.png)

Link this application to api demo-istio ONLY with scope read:

![alt text](img/appli3.png)

Define new Api, demo-istio-v1, with url api ie. http://<REDACTED>.amazonaws.com/v1:

![alt text](img/apiv1.png)

Define scope for this new api:

![alt text](img/scopev1.png)

Create a new m2m application which have access to api /v1 and /v2:

![alt text](img/appli1v1.png)

Select api and scope:

![alt text](img/appli2v1.png)

You should see:

![alt text](img/appli3v1.png)

Add new api and scope:

![alt text](img/appli4v1.png)

You should see:

![alt text](img/appli5v1.png)


### Auth0 Check

Get your access token with m2m machine that only have access to api /v2, replace `url` `client_id` `client_secret` and `audience`.

Where:

- `url` is your Auth0 Management API 
- `client_id` is your client id of m2m application which have access to audience api `/v2/`
- `client_secret` is your client secret of m2m application
- `audience` is your target api ie. `/v2`

```sh
curl --request POST \
  --url https://<REDACTED>.eu.auth0.com/oauth/token \
  --header 'content-type: application/json' \
  --data '{"client_id":"REDACTED>","client_secret":"<REDACTED>","audience":"http://<REDACTED>.amazonaws.com/v2/","grant_type":"client_credentials"}'
  ```
You should see reply:

```sh
{"access_token":"<REDACTED>","scope":"read:version","expires_in":86400,"token_type":"Bearer"}%
```

You can check your access token on [jwt debugger](https://jwt.io):

![alt text](img/checktoken.png)

### Istio Configuration

Apply jwt istio policy:

```sh
istioctl create -f demo9/jwt-policy.yaml
```

### Access to your api

#### Without any JWT token

Without any jwt you can't access to /v1 or /v2:

```sh
curl -X GET -vvv http://<REDACTED>.amazonaws.com/v2/
```

Response:
```
Origin authentication failed.%
```
#### With JWT token audience /v2 only

With a jwt token from m2m machine that only have acces to /v2

Get your access token for audience /v2:

```sh
curl --request POST \
  --url https://<REDACTED>.eu.auth0.com/oauth/token \
  --header 'content-type: application/json' \
  --data '{"client_id":"REDACTED>","client_secret":"<REDACTED>","audience":"http://<REDACTED>.amazonaws.com/v2/","grant_type":"client_credentials"}'
```

```sh
curl -X GET \
  http://<REDACTED>.amazonaws.com/v2/ \
  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGc<REDACTED>' \
```

Response:
```sh
Hello, Im Service version: v.0.0.2
My IP is: <REDACTED>
```

```sh
curl -X GET \
  http://<REDACTED>.amazonaws.com/v1/ \
  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGc<REDACTED>' \
```

Response:
```
Origin authentication failed.%
```

Try to get access token for audience /v1:

```sh
curl --request POST \
  --url https://<REDACTED>.eu.auth0.com/oauth/token \
  --header 'content-type: application/json' \
  --data '{"client_id":"REDACTED>","client_secret":"<REDACTED>","audience":"http://<REDACTED>.amazonaws.com/v1/","grant_type":"client_credentials"}'r access token for audience /v2:
```

Response:

```sh
{"error":"access_denied","error_description":"Client is not authorized to access \"http://<REDACTED>.amazonaws.com/v1/\". You might probably want to create a \"client-grant\" associated to this API. See: https://auth0.com/docs/api/v2#!/Client_Grants/post_client_grants"}%
```

#### With JWT token audience /v1 and /v2

With a jwt token form m2m machine that have access to /v1 and /v2.

:warning: **Because auth0 doesn't support multiple audience you need to retrieve your access token for each audience ie. /v1 and /v2**

```sh
curl -X GET \
  http://<REDACTED>.amazonaws.com/v2/ \
  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIs<REDACTED>' \
```
Response:
```sh
Hello, Im Service version: v.0.0.2
My IP is: <REDACTED>
```

```sh
curl -X GET \
  http://<REDACTED>.amazonaws.com/v1/ \
  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIs<REDACTED>' \
```
Response:
```sh
Hello, Im Service version: v.0.0.1
My IP is: <REDACTED>
```

### Further reading

Have a look on [RBAC filtering](../demo10/README.MD) with scope of your access token
