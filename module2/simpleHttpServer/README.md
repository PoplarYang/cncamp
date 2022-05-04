# 本地方式运行
## http server demo
- 接收客户端 request，并将 request 中带的 header 写入 response header
- 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
- Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
- 当访问 localhost/healthz 时，应返回 200

## 运行程序
```bash
~# export VERSION=1.0
~# go run simpleHttpServer.go -logtostderr=true
```
> -logtostderr=true 开启日志输出到控制台

## 测试
```bash
~# curl 127.0.0.1:8080/healthz
```

# 容器方式运行

## 构建镜像
```bash
~# docker build -t httpserver:v1.0 .
```

## 运行容器
```bash
~# docker run -P -d httpserver:v1.0
```

## 测试
```bash
~# ps aux | grep httpserver
root      820451  0.0  0.0 708292  4696 ?        Ssl  19:34   0:00 ./httpserver -logtostderr=true
root      820859  0.0  0.0   9032   672 pts/4    S+   19:37   0:00 grep --color=auto httpserver

~# nsenter -t 820451 -n ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
60: eth0@if61: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:04 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.4/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever

# curl 172.17.0.4:8080/healthz
status ok
```

## push 到 hub.docker.com
```bash
~# docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: hiyang
Password:
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded

~# docker tag c517a465f4ef  docker.io/hiyang/httpserver:v1

~# docker push hiyang/httpserver
Using default tag: latest
The push refers to repository [docker.io/hiyang/httpserver]
tag does not exist: hiyang/httpserver:latest
root@aliyun:~# docker push hiyang/httpserver:v1
The push refers to repository [docker.io/hiyang/httpserver]
b2b94fbf38ee: Pushed
5e7c68d682a7: Pushed
v1: digest: sha256:f85704af7c74cd6f967096071c1dfcad504c7160aa42270025d0bc38c39a1f10 size: 735
```
