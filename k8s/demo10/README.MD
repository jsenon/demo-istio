# Service Mesh

## Demo10

### Aim

- Use RBAC with JWT access token

### Deployment

Use previous [demo9](../demo9/README.MD)

### Auth0 Configuration

Create a new scope on api audience /v2:

![alt text](img/api1.png)

Add new scope:

![alt text](img/api2.png)

Create a new m2m application which have access to api /v2:

![alt text](img/appli1.png)

Select audience api /v2 and scope created before:

![alt text](img/appli2.png)

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
{"access_token":"<REDACTED>","scope":"YOUR SCOPE","expires_in":86400,"token_type":"Bearer"}%
```

You can check your access token on [jwt debugger](https://jwt.io):

### Istio Configuration

Apply rbac policy:

```sh
istioctl create -f demo10/rbac.yaml
```

### Access to your api

#### With JWT token audience /v2 scope read:version

With a jwt token form m2m machine that have access to /v2 and scope read:version

`curl` your application with `GET` verbs:
```sh
curl -X GET \
  http://<REDACTED>.amazonaws.com/v2/ \
  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIs<REDACTED>' \
```

You should see, Response:
```sh
Hello, Im Service version: v.0.0.2
My IP is: <REDACTED>
```

`curl` your application with `POST` verbs:

```sh
curl -X POST \
  http://<REDACTED>.amazonaws.com/v2/ \
  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIs<REDACTED>' \
```

You should see, Response:
```sh
RBAC: access denied%
```

#### With JWT token audience /v2 scope write:version

With a jwt token form m2m machine that have access to /v2 and scope read:version

`curl` your application with `GET` verbs:

```sh
curl -X GET \
  http://<REDACTED>.amazonaws.com/v2/ \
  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIs<REDACTED>' \
```

You should see, Response:
```sh
RBAC: access denied%
```

`curl` your application with `GET` verbs:

```sh
curl -X POST \
  http://<REDACTED>.amazonaws.com/v2/ \
  -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIs<REDACTED>' \
```

You should see, Response:

```sh
Hello, Im Service version: v.0.0.2
My IP is: <REDACTED>
```