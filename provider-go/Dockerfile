# syntax=docker/dockerfile:1.7

FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS build

RUN apk add --no-cache ca-certificates git

WORKDIR /src
COPY go.mod go.sum ./

ARG METATUBE_SDK_VERSION=main
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go get github.com/metatube-community/metatube-sdk-go@${METATUBE_SDK_VERSION} && \
    go mod download

COPY *.go ./

ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /out/metatube-provider-go .

FROM alpine:3.22

RUN apk add --no-cache ca-certificates && \
    addgroup -S provider && \
    adduser -S -G provider provider && \
    mkdir -p /data && \
    chown provider:provider /data

COPY --from=build /out/metatube-provider-go /usr/local/bin/metatube-provider-go

ENV METATUBE_HOST=0.0.0.0 \
    METATUBE_PORT=8080 \
    METATUBE_DSN=/data/metatube-provider-go.db

WORKDIR /data
VOLUME ["/data"]
USER provider:provider

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/metatube-provider-go"]
