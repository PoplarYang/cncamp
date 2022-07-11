
## install istio
refer: https://istio.io/latest/zh/docs/setup/getting-started/

```bash
curl -L https://istio.io/downloadIstio | sh -
istioctl install --set profile=demo -y
```

## 验证
正常运行
```bash
root@kubernetes-master:/jfs/cncamp/module12/simple# kubectl get pod -n istio-system
NAME                                    READY   STATUS    RESTARTS   AGE
istio-egressgateway-7fcb98978c-h8znr    1/1     Running   0          24m
istio-ingressgateway-55b6cffcbc-xlksv   1/1     Running   0          27m
istiod-56d9c5557-wx4ld                  1/1     Running   0          26m
```

## 创建namespace
```bash
root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl create ns httpserver
namespace/httpserver created

root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl label namespace httpserver istio-injection=enabled
namespace/httpserver labeled
```

## 创建 deployment 和 svc
```bash
root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl create -n httpserver -f httpserver-deployment.yaml
deployment.apps/httpserver-deployment created
root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl create -n httpserver -f httpserver-service.yaml
service/httpsvc created

root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl get svc -n httpserver
NAME      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
httpsvc   ClusterIP   10.110.133.21   <none>        8080/TCP   24s

root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl get pod -n httpserver
NAME                                   READY   STATUS    RESTARTS   AGE
httpserver-deployment-6dd57654-2mv8h   2/2     Running   0          49s
httpserver-deployment-6dd57654-g986j   2/2     Running   0          49s
```

## 创建tls
```bash
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/CN=*.tree.io' -keyout tree.io.key -out tree.io.crt

kubectl create -n istio-system secret tls httpserver --key=tree.io.key --cert=tree.io.crt
```

## create gw
```bash
root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl create -n httpserver -f gw.yaml
gateway.networking.istio.io/httpserver-gw created

root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl get gw -n httpserver
NAME            AGE
httpserver-gw   30s
```

## create vs
```bash
root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl create -n httpserver -f vs.yaml
virtualservice.networking.istio.io/httpserver created

root@kubernetes-master:/jfs/cncamp/module12/normal# kubectl get vs -n httpserver
NAME         GATEWAYS            HOSTS                    AGE
httpserver   ["httpserver-gw"]   ["httpserver.tree.io"]   30s
```

## 访问验证
```bash
root@kubernetes-master:/jfs/cncamp/module12/simple# kubectl get svc -n istio-system
NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP      10.103.236.157   <none>        80/TCP,443/TCP                                                               30h
istio-ingressgateway   LoadBalancer   10.102.56.71     <pending>     15021:32592/TCP,80:32102/TCP,443:32290/TCP,31400:31932/TCP,15443:31967/TCP   30h
istiod                 ClusterIP      10.102.106.65    <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        30h

root@kubernetes-master:/jfs/cncamp/module12/normal# curl --resolve httpserver.tree.io:443:10.102.56.71 https://httpserver.tree.io/healthz -v -k
* Added httpserver.tree.io:443:10.102.56.71 to DNS cache
* Hostname httpserver.tree.io was found in DNS cache
*   Trying 10.102.56.71:443...
* TCP_NODELAY set
* Connected to httpserver.tree.io (10.102.56.71) port 443 (#0)
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
*  subject: CN=*.tree.io
*  start date: Jul 11 15:00:15 2022 GMT
*  expire date: Jul 11 15:00:15 2023 GMT
*  issuer: CN=*.tree.io
*  SSL certificate verify result: self signed certificate (18), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x5600c19482f0)
> GET /healthz HTTP/2
> Host: httpserver.tree.io
> user-agent: curl/7.68.0
> accept: */*
>
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* old SSL session ID is stale, removing
* Connection state changed (MAX_CONCURRENT_STREAMS == 2147483647)!
< HTTP/2 200
< date: Mon, 11 Jul 2022 15:59:33 GMT
< content-length: 9
< content-type: text/plain; charset=utf-8
< x-envoy-upstream-service-time: 2
< server: istio-envoy
<
* Connection #0 to host httpserver.tree.io left intact
status ok
```

## TODO
```bash
curl https://10.102.56.71/healthz -H 'host: httpserver.tree.io' -v -k
```