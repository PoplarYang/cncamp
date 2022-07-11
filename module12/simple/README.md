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

## create gw
```bash
kubectl create -f gw.yaml
```

## create vs
```bash
kubectl create -f vs.yaml
```

## 访问
```bash
root@kubernetes-master:/jfs/cncamp/module12/simple# kubectl get svc -n istio-system
NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP      10.103.236.157   <none>        80/TCP,443/TCP                                                               30h
istio-ingressgateway   LoadBalancer   10.102.56.71     <pending>     15021:32592/TCP,80:32102/TCP,443:32290/TCP,31400:31932/TCP,15443:31967/TCP   30h
istiod                 ClusterIP      10.102.106.65    <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        30h
root@kubernetes-master:/jfs/cncamp/module12/simple# curl 10.102.56.71/healthz
status ok
```
