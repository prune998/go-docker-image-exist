FROM golang:1.11-alpine as builder

LABEL vendor="Prune - prune@lecentre.net" \
      content="go-docker-image-exist"

ARG VERSION="0.1"
ARG BUILDTIME="20190107"

COPY . /go/src/github.com/prune998/go-docker-image-exist
WORKDIR /go/src/github.com/prune998/go-docker-image-exist

RUN    CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-X main.version=${VERSION}-${BUILDTIME}"

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
WORKDIR /root/
COPY --from=0 /go/src/github.com/prune998/go-docker-image-exist/go-docker-image-exist .

ENTRYPOINT ["/root/go-docker-image-exist"]