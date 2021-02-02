## 1.环境准备

- kubectl tools下载
- minikube 搭建本地k8s集群环境，[详情见](https://minikube.sigs.k8s.io/docs/start/#arm)
- docker本地搭建

## 2.开始

- 构建基础镜像推送到镜像仓库

- 编写deployment.yaml，[编写规范参考](https://cloud.google.com/kubernetes-engine/docs/how-to/stateless-apps?hl=zh-cn#create)，[或者官方文档](https://kubernetes.io/docs/tutorials/stateless-application/expose-external-ip-address/)

- 执行kubectl apply -f  your_deployment_path.yaml

  若image为private,报错：docker login

  请执行下面命令进行注册,[参考](https://kubernetes.io/docs/concepts/configuration/secret/)

  ```go
  kubectl create secret docker-registry secret-tiger-docker \
    --docker-server=your docker server
    --docker-username=your username \
    --docker-password=your password \
    --docker-email=your email
  ```

- 若要暴露服务端口，提供本地测试，[参考](https://kubernetes.io/docs/tutorials/stateless-application/expose-external-ip-address/)

  例如：

  ```go
  kubectl port-forward service/test-service 7080:8080
  ```

  