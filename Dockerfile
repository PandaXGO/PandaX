FROM alpine:latest
LABEL MAINTAINER="PandaX"

WORKDIR /go/src/panda
COPY ./pandax ./
COPY ./config.yml ./
COPY ./resource ./resource
COPY ./uploads ./uploads

RUN chmod 755 ./pandax

EXPOSE 7788
EXPOSE 9001
EXPOSE 9002
EXPOSE 9003
EXPOSE 9003/udp

ENTRYPOINT ./pandax