FROM alpine:latest
LABEL MAINTAINER="PandaX"

WORKDIR /go/src/panda
COPY ./dziot ./
COPY ./config.yaml ./
COPY ./kubeconfig ./
COPY ./web ./web
COPY ./resource ./resource

EXPOSE 8080

ENTRYPOINT ./dziot -c ./config.yaml -kc ./kubeconfig
