## README

作业1

现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？
是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。
作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

作业拆解
- 优雅启动

  --> pod中程序启动前，需要准备资源等，这里通过touch文件体现
- 优雅终止

  --> pod中程序终止前，需要释放资源，这里通过删除启动时的touch文件体现

  --> 程序支持捕捉 SIGTREM 信号，捕捉到信号后，拒绝新的连接，已经建立的连接继续处理，处理完成后退出程序；避免程序被 SIGKILL

- 资源需求和 QoS 保证

  --> 通过 requests 和 limit 对pod所有资源进行控制

- 探活

  --> 通过 livenessProbe 探测 touch的文件

  --> 通过 readinessProbe 请求 /healthz

- 日常运维需求，日志等级

  --> glog 定义了日志等级

- 配置和代码分离
  --> 这里的配置比较简单，VERSION 通过环境变量传入

  --> 如果使用配置文件，则用configmap传入

涉及代码

deployment httpserver-deployment.yaml

module2/simpleHttpserver/simpleHttpServer.go


作业2

除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧。
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：

Service --> httpbin-deployment.yaml and httpserver-deployment.yaml

Ingress --> ingress.yaml

可以考虑的细节

如何确保整个应用的高可用。 --> 应用高可用通过多副本实现
如何通过证书保证 httpServer 的通讯安全。--> 添加https

> 通过使用annotation实现rewrite，但总的来看ingress-nginx的功能还是太弱。