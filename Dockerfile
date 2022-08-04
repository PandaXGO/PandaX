FROM alpine:latest
LABEL MAINTAINER="PandaX"

WORKDIR /go/src/panda
COPY ./pandax ./
COPY ./config.yml ./
COPY ./resource ./resource

RUN chmod 755 ./pandax

EXPOSE 7788

ENTRYPOINT ./pandax