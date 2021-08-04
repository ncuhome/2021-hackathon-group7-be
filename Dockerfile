FROM alpine:latest
MAINTAINER pyf(316851756@qq.com)

COPY ./main /data/main

COPY ./config /data/config

WORKDIR /data

ENTRYPOINT [ "/data/main" ]