## K3S
1. 安装  
curl -sfL http://rancher-mirror.cnrancher.com/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn INSTALL_K3S_VERSION="v1.23.1+k3s1" sh -
2. 卸载  
sudo /usr/local/bin/k3s-uninstall.sh 

## mysql 安装
1. kubectl kustomize ./deploy/mysql -o ./deploy/mysql.yaml
2. kubectl apply -f ./deploy/mysql.yaml

## k3s部署流程
1. 设置打包环境  
 go env -w GOOS=linux
 go env -w GOARCH=amd64
2. 构建Linux执行命令  
 go build -o pandax .
4. 构建docker镜像 (修改版本号 xmadmin/pandax:v1.0)  
 docker build -t xmadmin/pandax:v1.4 --rm .
5. 上传daocker镜像  
 docker push xmadmin/pandax:v1.4
6. 生成 deploy.yaml    
 kubectl kustomize deploy/manifest -o deploy/deploy.yaml
7. k8s安装yaml   
 kubectl apply -f deploy/deploy.yaml
 
## 查看部署状态
8. 查看 yaml 的安装状态  
 kubectl get pods -n pandax  
 kubectl get services -n pandax