## k8s部署流程
1. 设置打包环境  
 set GOOS=linux
 set GOARCH=amd64
2. 构建Linux执行命令  
 go build -o pandax .
4. 构建docker镜像 (修改版本号 例如pandax/pandax:v1.0)  
 docker build -t pandax/pandax:v1.0
5. 上传daocker镜像  
 docker push pandax/pandax:v1.0
6. 生成 deploy.yaml    
 kubectl kustomize deploy/manifest -o deploy/deploy.yaml
7. k8s安装yaml   
 kubectl apply -f deploy/deploy.yaml
8. 查看 yaml 的安装状态  
 kubectl get pods -n pandax
 kubectl get services -n pandax