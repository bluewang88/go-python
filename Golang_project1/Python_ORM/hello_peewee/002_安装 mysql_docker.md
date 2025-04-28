# MySQL Docker 安装指南

要安装运行 MySQL 的 Docker 容器来配合您的 Peewee 代码使用，请按照以下步骤操作：

## 1. 拉取 MySQL 镜像

```bash
docker pull mysql:8.0
```

## 2. 创建并运行 MySQL 容器

```bash
docker run --name mysql-peewee \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=peewee \
  -p 3306:3306 \
  -d mysql:8.0
```

这个命令会：
- 创建名为 `mysql-peewee` 的容器
- 设置 root 用户密码为 "root"
- 自动创建名为 "peewee" 的数据库
- 将容器的 3306 端口映射到主机的 3306 端口
- 后台运行容器

## 3. 验证容器是否正常运行

```bash
docker ps
```

您应该能看到 `mysql-peewee` 容器正在运行。

## 4. 连接到 MySQL 进行测试

```bash
docker exec -it mysql-peewee mysql -uroot -proot
```

进入 MySQL 命令行后，可以查看数据库：

```sql
SHOW DATABASES;
USE peewee;
```

## 5. 与您的 Peewee 代码连接

您的 Peewee 代码已正确配置连接参数：

```python
mysqldb = MySQLDatabase('peewee', user='root', password='root', host='127.0.0.1', port=3306)
```

这些参数与我们创建的 Docker 容器完全匹配。

## 常用管理命令

```bash
# 停止 MySQL 容器
docker stop mysql-peewee

# 启动已停止的容器
docker start mysql-peewee

# 查看容器日志
docker logs mysql-peewee

# 删除容器（先停止）
docker rm mysql-peewee
```

## 使用 Docker Compose (推荐)

更好的方式是使用 Docker Compose，创建 `docker-compose.yml` 文件：

```yaml
version: '3'

services:
  mysql:
    image: mysql:8.0
    container_name: mysql-peewee
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: peewee
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

volumes:
  mysql_data:
```

然后使用以下命令启动：

```bash
docker-compose up -d
```

这种方式更容易管理，特别是当您需要添加其他服务（如 phpMyAdmin）时。

## 注意事项

1. 生产环境中应该使用更强的密码
2. 考虑为数据创建持久卷以防容器被删除时数据丢失
3. 可能需要设置适当的 MySQL 配置参数以优化性能


# Docker Compose 文件查找规则

是的，执行 `docker-compose up -d` 命令时，默认情况下 Docker Compose 只会查找**当前工作目录**中的 docker-compose.yml 或 `docker-compose.yaml` 文件。

## 如何使用不同位置或文件名的配置文件

如果需要指定其他位置或不同名称的配置文件，可以使用 `-f` 或 `--file` 参数：

```bash
# 指定不同路径的文件
docker-compose -f /path/to/other-directory/docker-compose.yml up -d

# 指定不同名称的文件
docker-compose -f custom-compose.yml up -d

# 在其他目录中运行
cd /path/to/other-directory && docker-compose up -d
```

## 多文件组合

您还可以组合多个配置文件，它们会按顺序合并：

```bash
docker-compose -f docker-compose.yml -f docker-compose.override.yml up -d
```

## 实用场景

1. **不同环境配置**：
   ```bash
   # 开发环境
   docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
   
   # 生产环境
   docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
   ```

2. **从任何位置启动特定项目**：
   ```bash
   docker-compose -f /home/ubuntu/codespace/go-python/Golang_project1/Python_ORM/hello_peewee/docker-compose.yml up -d
   ```

因此，如果您想运行 docker-compose.yml 文件，必须先切换到该目录或使用 `-f` 参数指定完整路径。