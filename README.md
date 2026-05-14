本项目是论文《Design and Implementation of a Data Sharing System Based on Hyperledger Fabric》的实验代码, 实现了一种基于区块链技术的数据共享系统。

The code in this project is the experimental implementation for the paper titled 'Design and Implementation of a Data Sharing System Based on Hyperledger Fabric', which realizes a blockchain-based data sharing system.

## 目录 Table of Contents

- [项目目录 Project Directory](#项目目录-project-directory)
- [使用方法 Getting Started](https://github.com/guihaorui/-Hyperledger-Fabric-/blob/main/README.md)
- [项目声明 Project Statement](https://github.com/guihaorui/-Hyperledger-Fabric-/blob/main/%E9%A1%B9%E7%9B%AE%E5%A3%B0%E6%98%8E%20Project%20Statement)

部署步骤

1. 安装docker 

	```bash
	#下载docker 
	# 官方脚本当前已无法下载，使用gitee备份的脚本:
	curl -fsSL https://gitee.com/real__cool/fabric_install/raw/main/docker_install.sh | bash -s docker --mirror Aliyun
	#添加当前用户到docker用户组 
	sudo usermod -aG docker $USER 
	newgrp docker 
	sudo mkdir -p /etc/docker
	#配置docker镜像加速
	sudo tee /etc/docker/daemon.json <<-'EOF'
	{
	    "registry-mirrors": [
	        "https://docker.m.daocloud.io",
	        "https://docker.1panel.live",
	        "https://hub.rat.dev"
	    ]
	}
	EOF

	#重启docker 
	sudo systemctl restart docker
	```

2. 安装开发使用的go、node、jq

	```bash
	#下载二进制包
	wget https://golang.google.cn/dl/go1.19.linux-amd64.tar.gz
	#将下载的二进制包解压至 /usr/local目录
	sudo tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
	mkdir $HOME/go
	#将以下内容添加至环境变量 ~/.bashrc
	export GOPATH=$HOME/go
	export GOROOT=/usr/local/go
	export PATH=$GOROOT/bin:$PATH
	export PATH=$GOPATH/bin:$PATH
	#更新环境变量
	source  ~/.bashrc 
	#设置代理
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	
	#下载nvm安装脚本
	wget https://gitee.com/real__cool/fabric_install/raw/main/nvminstall.sh
	#安装nvm；屏幕输出内容添加环境变量
	chmod +x nvminstall.sh
	./nvminstall.sh
	# 将环境变量写入.bashrc
	export NVM_DIR="$HOME/.nvm"
	[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
	[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
	export NVM_NODEJS_ORG_MIRROR=http://npmmirror.com/mirrors/node/ #更换阿里云nvm node源
 
	# 更新环境变量
	source  ~/.bashrc
	# 安装node16
	nvm install 16
	#换源
	npm config set registry https://registry.npmmirror.com
	
	#安装jq 
	sudo apt install jq
	```



3. 解压本项目 

	```bash
	tar -xzvf dataex.tar.gz
	```

4. 启动区块链部分。在dataex/blockchain/network目录下:

	```bash
	# 仅在首次使用执行：下载Fabric Docker镜像。如果拉取速度过慢或失败请检查是否完成docker换源并执行了重启docker命令。
	./install-fabric.sh -f 2.5.6 d 
	```
	```bash
	# 启动区块链网络
	./start.sh
	```	
 	 **如果在启动区块链网络时遇到报错可以尝试:**
	```bash
	# 执行清理所有的容器指令：
	docker rm -f $(docker ps -aq)
	```
	**然后再重新启动区块链网络**

5. 运行dataex/application/replaceip.sh 按照提示更换服务器IP。启动后端 在dataex/application/backend目录下： 执行： `go run main.go`


6. 新开一个窗口，启动前端 在dataex/application/web目录下： 执行： 

	```bash
	# 仅在首次运行执行：安装依赖
	npm install 
	```

	```bash
	# 启动前端
	npm run dev
	```

7. 在浏览器中打开：http://服务器IP:9528 即可看到前端页面。需要在服务器防火墙放行9090、9528端口

#### 六、关闭项目与重新运行步骤
##### 关闭项目：
1. 前端（`npm run dev`界面）与后端（`go run main.go`界面：

	使用键盘组合键：`ctrl+c`

2. 区块链部分：

	在`dataex/blockchain/network`目录`./stop.sh`，此步骤会清理所有的区块链、Mysql中的数据。

##### 开发模式启动项目：
1. 在`dataex/blockchain/network`目录
`./start.sh` 如果遇到报错可以执行以下命令后再试：
执行清理所有的容器指令：
`docker rm -f $(docker ps -aq)`
2. 在`dataex/application/backend`目录下： 执行： `go run main.go`
3. 在`dataex/application/web`目录下： 执行：
`npm run dev`
4. 在http://服务器IP:9528打开

## 项目声明 Project Statement
本项目的作者及单位:
The author and affiliation of this project:

项目名称(Project Name):基于Hyperledger Fabric的数据共享系统的设计与实现
项目作者 (Author) : Haorui Gui, Anjia Yang
作者单位(Affiliation):暨南大学网络空间安全学院(College of cyber Security,Jinan University)
