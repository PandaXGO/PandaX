FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download
RUN go build -o pandax

FROM alpine:latest
LABEL MAINTAINER="PandaX"

WORKDIR /pandax
COPY --from=builder /app/pandax ./
COPY --from=builder /app/config.yml ./
COPY --from=builder /app/resource ./resource
COPY --from=builder /app/uploads ./uploads

RUN chmod 755 ./pandax

EXPOSE 7788
EXPOSE 9001
EXPOSE 9002
EXPOSE 9003
EXPOSE 9003/udp

ENTRYPOINT ./pandax