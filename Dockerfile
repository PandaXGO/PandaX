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
EXPOSE 8801
EXPOSE 5060

ENTRYPOINT ./pandax