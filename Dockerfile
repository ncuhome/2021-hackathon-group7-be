FROM golang:1.15.5-alpine as builder

RUN go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=0 \
  && go env -w GOPROXY=https://goproxy.cn

WORKDIR /build

# 先把go.mod go.sum复制，这俩文件没变时构建时可以使用缓存，不需要重新go mod download
COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -ldflags '-extldflags "-static"' -o main

FROM alpine:latest

COPY --from=builder /build/main /usr/bin/main

COPY --from=builder /build/config /data/config

WORKDIR /data

ENTRYPOINT [ "/usr/bin/main" ]