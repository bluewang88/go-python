# Docker 与 Docker Compose 安装部署详细指南

## 一、Docker 介绍

Docker 是一个开源的应用容器引擎，让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的 Linux 或 Windows 机器上。Docker Compose 则是一个用于定义和运行多容器 Docker 应用程序的工具。

## 二、Docker 安装

### Ubuntu 系统安装 Docker

1. **更新软件包索引**
   ```bash
   sudo apt update
   sudo apt upgrade -y
   ```

2. **安装docker必要的依赖包**
   ```bash
   sudo apt install -y apt-transport-https ca-certificates curl  gnupg lsb-release software-properties-common
   ```

3. **添加 Docker 官方 GPG 密钥**
   ```bash
   curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
   ```

4. **添加 Docker 软件源**
   ```bash
   sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
   ```

5. **再次更新软件包索引**
   ```bash
   sudo apt update
   ```

6. **安装最新版本的 Docker**
   ```bash
   sudo apt install -y docker-ce docker-ce-cli containerd.io
   ```

7. **启动并设置自启动**
   ```bash
   sudo systemctl start docker
   sudo systemctl enable docker
   ```

8. **验证安装**
   ```bash
   sudo docker run hello-world
   ```

9.  **设置当前用户无需 sudo 运行 Docker**（可选）
   ```bash
   sudo usermod -aG docker $USER
   # 重新登录以应用更改
   ```

### CentOS 系统安装 Docker

1. **安装必要的依赖包**
   ```bash
   sudo yum install -y yum-utils device-mapper-persistent-data lvm2
   ```

2. **添加 Docker 软件源**
   ```bash
   sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
   ```

3. **安装 Docker**
   ```bash
   sudo yum install -y docker-ce docker-ce-cli containerd.io
   ```

4. **启动并设置自启动**
   ```bash
   sudo systemctl start docker
   sudo systemctl enable docker
   ```

5. **验证安装**
   ```bash
   sudo docker run hello-world
   ```

### macOS 安装 Docker Desktop

1. 访问 [Docker Desktop 下载页面](https://www.docker.com/products/docker-desktop)
2. 下载 Mac 版本的 Docker Desktop 安装包
3. 双击安装包进行安装
4. 安装完成后，启动 Docker Desktop 应用
5. 验证安装：在终端运行 `docker run hello-world`

### Windows 安装 Docker Desktop

1. 访问 [Docker Desktop 下载页面](https://www.docker.com/products/docker-desktop)
2. 下载 Windows 版本的 Docker Desktop 安装包
3. 双击安装包进行安装
4. 可能需要启用 WSL 2 或 Hyper-V（根据系统提示操作）
5. 安装完成后，启动 Docker Desktop 应用
6. 验证安装：在 PowerShell 或命令提示符中运行 `docker run hello-world`

## 三、Docker Compose 安装

### Linux 安装 Docker Compose

1. **下载 Docker Compose**
   ```bash
   sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
   ```

2. **添加执行权限**
   ```bash
   sudo chmod +x /usr/local/bin/docker-compose
   ```

3. **验证安装**
   ```bash
   docker-compose --version
   ```

4. **如果上面的方法不可用，也可以使用 pip 安装**
   ```bash
   sudo pip install docker-compose
   ```

### macOS 安装 Docker Compose

如果您使用的是 Docker Desktop，Docker Compose 已经包含在其中。如果需要单独安装：

```bash
brew install docker-compose
```

### Windows 安装 Docker Compose

如果您使用的是 Docker Desktop，Docker Compose 已经包含在其中。如果需要单独安装：

1. 使用 pip 安装：
   ```
   pip install docker-compose
   ```

## 四、Docker 基本配置

### 配置国内镜像源（可选）

1. **创建或编辑 daemon.json 文件**：
   ```bash
   # Linux
   sudo mkdir -p /etc/docker
   sudo nano /etc/docker/daemon.json
   
   # macOS/Windows (Docker Desktop)
   # 在 Docker Desktop 的设置 -> Docker Engine 中编辑配置
   ```

2. **添加以下内容**：
   ```json
   {
     "registry-mirrors": [
       "https://registry.docker-cn.com",
       "https://docker.mirrors.ustc.edu.cn",
       "https://hub-mirror.c.163.com"
     ]
   }
   ```

3. **重启 Docker 服务**：
   ```bash
   # Linux
   sudo systemctl daemon-reload
   sudo systemctl restart docker
   
   # macOS/Windows
   # 在 Docker Desktop 中点击 Apply & Restart
   ```

### 调整资源限制（macOS 和 Windows）

在 Docker Desktop 的设置中，可以调整分配给 Docker 的 CPU、内存和磁盘资源。

## 五、Docker Compose 使用示例

1. **创建一个简单的 Docker Compose 项目**

   创建一个名为 `docker-compose.yml` 的文件，内容如下：
   ```yaml
   version: '3'
   services:
     web:
       image: nginx:latest
       ports:
         - "8080:80"
       volumes:
         - ./html:/usr/share/nginx/html
       restart: always
   ```

2. **创建 html 目录和测试文件**
   ```bash
   mkdir -p html
   echo "<h1>Hello from Docker Compose!</h1>" > html/index.html
   ```

3. **启动服务**
   ```bash
   docker-compose up -d
   ```

4. **验证服务**
   - 在浏览器中访问 `http://localhost:8080`
   - 应该会看到 "Hello from Docker Compose!"

5. **查看容器状态**
   ```bash
   docker-compose ps
   ```

6. **停止服务**
   ```bash
   docker-compose down
   ```

## 六、常用 Docker 命令

```bash
# 列出所有运行中的容器
docker ps

# 列出所有容器（包括已停止的）
docker ps -a

# 启动/停止容器
docker start <container_id_or_name>
docker stop <container_id_or_name>

# 删除容器
docker rm <container_id_or_name>

# 列出本地镜像
docker images

# 删除镜像
docker rmi <image_id_or_name>

# 构建镜像
docker build -t <image_name> .

# 查看日志
docker logs <container_id_or_name>
```

## 七、常用 Docker Compose 命令

```bash
# 启动服务
docker-compose up -d

# 停止服务
docker-compose down

# 查看日志
docker-compose logs

# 查看服务状态
docker-compose ps

# 重启服务
docker-compose restart

# 构建或重建服务
docker-compose build
```

## 八、常见问题排查

1. **权限不足**：
   - 错误：`Got permission denied...`
   - 解决方案：使用 `sudo` 或将用户添加到 docker 组

2. **端口被占用**：
   - 错误：`port is already allocated`
   - 解决方案：更改端口映射或停止占用端口的服务

3. **磁盘空间不足**：
   - 错误：`no space left on device`
   - 解决方案：清理未使用的镜像和容器，运行 `docker system prune`

4. **Docker Compose 版本兼容性**：
   - 错误：`version in "./docker-compose.yml" is unsupported`
   - 解决方案：更新 Docker Compose 或降低 YAML 文件中的版本号

## 九、最佳实践

1. 保持 Docker 和 Docker Compose 更新到最新版本
2. 使用官方镜像作为基础镜像
3. 为容器设置资源限制
4. 使用数据卷持久化重要数据
5. 利用 Docker 网络隔离容器
6. 为 Docker Compose 项目创建专用目录
7. 使用环境变量文件（.env）管理配置

通过以上步骤，您应该能够在各种操作系统上成功安装和使用 Docker 与 Docker Compose，开始容器化您的应用程序。

找到具有 1 个许可证类型的类似代码