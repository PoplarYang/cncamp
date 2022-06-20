### Install ingress controller

```sh
kubectl create -f httpserver-deployment.yaml
kubectl create -f httpbin-deployment.yaml
```

### Generate key-cert

```sh
# use openssl (version >= 1.1.1f) on Linux, e.g. Ubuntu 20.04
# don't run on macOS, which is using LibreSSL
# instead, you can `brew install openssl` on macOS
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=tree.io/O=tree" -addext "subjectAltName = DNS:tree.io"
```

### Create secret

```sh
kubectl create secret tls tree-io-tls --cert=./tls.crt --key=./tls.key
```

### Create a ingress

```sh
kubectl create -f ingress.yaml
```

### 查看ingress地址
```bash
root@kubernetes-master:/jfs/cncamp101/module8/ingress# kubectl get svc -n ingress-nginx
NAME                                 TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx-controller             LoadBalancer   10.108.3.54     <pending>     80:31823/TCP,443:31098/TCP   48m
ingress-nginx-controller-admission   ClusterIP      10.111.165.87   <none>        443/TCP                      48m
```

### Test the result

```sh
# 访问 httpserver/get/200 --> /get/200
root@kubernetes-master:/jfs/cncamp/module8# curl -H "Host: tree.io" https://10.108.3.54/httpserver/get/200 -k -v
*   Trying 10.108.3.54:443...
* TCP_NODELAY set
* Connected to 10.108.3.54 (10.108.3.54) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: O=Acme Co; CN=Kubernetes Ingress Controller Fake Certificate
*  start date: Jun 19 16:14:34 2022 GMT
*  expire date: Jun 19 16:14:34 2023 GMT
*  issuer: O=Acme Co; CN=Kubernetes Ingress Controller Fake Certificate
*  SSL certificate verify result: unable to get local issuer certificate (20), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x561ca1b422f0)
> GET /httpserver/get/200 HTTP/2
> Host: tree.io
> user-agent: curl/7.68.0
> accept: */*
>
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* old SSL session ID is stale, removing
* Connection state changed (MAX_CONCURRENT_STREAMS == 128)!
< HTTP/2 200
< date: Mon, 20 Jun 2022 13:53:40 GMT
< content-type: text/plain; charset=utf-8
< content-length: 277
< x-real-ip: 172.31.0.173
< strict-transport-security: max-age=15724800; includeSubDomains
<
X-Real-Ip: 172.31.0.173
X-Forwarded-For: 172.31.0.173
X-Forwarded-Port: 443
X-Forwarded-Proto: https
X-Scheme: https
User-Agent: curl/7.68.0
X-Forwarded-Host: tree.io
X-Forwarded-Scheme: https
Accept: */*
X-Request-Id: b7838689d95d9890e14ac06f34956247
status ok, version: v1.2
* Connection #0 to host 10.108.3.54 left intact

# 访问 /httpbin/get --> /get
root@kubernetes-master:/jfs/cncamp/module8# curl -H "Host: tree.io" https://10.108.3.54/httpbin/get -k -v
*   Trying 10.108.3.54:443...
* TCP_NODELAY set
* Connected to 10.108.3.54 (10.108.3.54) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: O=Acme Co; CN=Kubernetes Ingress Controller Fake Certificate
*  start date: Jun 19 16:14:34 2022 GMT
*  expire date: Jun 19 16:14:34 2023 GMT
*  issuer: O=Acme Co; CN=Kubernetes Ingress Controller Fake Certificate
*  SSL certificate verify result: unable to get local issuer certificate (20), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x562e1b65a2f0)
> GET /httpbin/get HTTP/2
> Host: tree.io
> user-agent: curl/7.68.0
> accept: */*
>
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* old SSL session ID is stale, removing
* Connection state changed (MAX_CONCURRENT_STREAMS == 128)!
< HTTP/2 200
< date: Mon, 20 Jun 2022 13:54:56 GMT
< content-type: application/json; encoding=utf-8
< content-length: 396
< access-control-allow-credentials: true
< access-control-allow-origin: *
< strict-transport-security: max-age=15724800; includeSubDomains
<
* Connection #0 to host 10.108.3.54 left intact
{"args":{},"headers":{"Accept":["*/*"],"Host":["tree.io"],"User-Agent":["curl/7.68.0"],"X-Forwarded-For":["172.31.0.173"],"X-Forwarded-Host":["tree.io"],"X-Forwarded-Port":["443"],"X-Forwarded-Proto":["https"],"X-Forwarded-Scheme":["https"],"X-Real-Ip":["172.31.0.173"],"X-Request-Id":["a12a3af153ab1a33b6ca2573b0b62d24"],"X-Scheme":["https"]},"origin":"172.31.0.173","url":"https://tree.io/get"}
```
