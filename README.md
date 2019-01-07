
# go-docker-image-exist

a Go application to check if a Docker image exist in a remote repository

## Building

### Locally

Just clone this repository into your `GOPATH` (who cares about `GOPATH` anymore ?) and `go build`

### Docker Image

just use `docker build -t go-docker-image-exist:latest .`

## Usage

```go
./go-docker-image-exist -registryURL="https://us.gcr.io" -username='_token' -password=$(gcloud auth print-access-token) -logLevel=debug  -project=kube00-xxx -image=caddy/caddy -tag 0.11.0
```

- registry URL: the URL to reach the registry. Do not add the `/v2` endpoint
  - Google: `https://us.gcr.io` for the US registry. Can also be `https://gcr.io`
- project: the GCloud project name (keep empty is no project is used)
  - Google: find your project with `gcloud projects list`
- username: your Registry username.
  - Google: `_token`
- password: your password
  - Google: get your token password with `gcloud auth print-access-token`
- logLevel: one of debug, info, warning, error. When using `debug` you will also get the HTTP calls to the Registry.
- image: full name of the remote image, including directories
- tag: the tag to search for, defaults to `latest`

## example

```bash
./go-docker-image-exist -registryURL="https://us.gcr.io" -username='_token' -password=$(gcloud auth print-access-token) -logLevel=debug  -project=kube00-xxx -image=caddy/caddy -tag 0.11.0 
{"level":"debug","msg":"searching image kube00-xxx/caddy/caddy","time":"2019-01-07T11:26:02-05:00"}
{"level":"debug","msg":"registry.ping url=https://us.gcr.io/v2/","time":"2019-01-07T11:26:02-05:00"}
{"level":"debug","msg":"registry.tags url=https://us.gcr.io/v2/kube00-xxx/caddy/caddy/tags/list repository=kube00-xxx/caddy/caddy","time":"2019-01-07T11:26:03-05:00"}
{"level":"debug","msg":"image kube00-xx/caddy/caddy found","time":"2019-01-07T11:26:03-05:00"}
```

### Using the docker image

```bash
docker run -ti go-docker-image-exist:latest -registryURL="https://us.gcr.io" -username='_token' -password=$(gcloud auth print-access-token) -project=kube00-xxx -image=coyotelab/caddy -tag 0.11.0 ; echo $?
0
```