#!/bin/sh
sudo snap install docker go
sudo systemctl start docker
sudo systemctl enable docker

docker run -itd --name mysql --restart=always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:8.0.23
docker run -itd --name tdengine --restart=always -p 6030:6030 -p 6041:6041 -p 6043-6049:6043-6049 -p 6043-6049:6043-6049/udp tdengine/tdengine:3.0.4.2
docker run -itd -p 6379:6379 --restart=always --name redis redis:7.0.4 --requirepass 123456
docker run -d --name emqx --restart=always -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx/emqx:5.1.0

docker cp pandax_iot.sql mysql:/var/lib/mysql/pandax_iot.sql
docker exec mysql mysql -uroot -p123456 pandax_iot < /var/lib/mysql/pandax_iot.sql

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go mod tidy
go run main.go