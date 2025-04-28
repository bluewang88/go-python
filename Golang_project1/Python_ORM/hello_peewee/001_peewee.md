# Peewee ORM 详细指南

## 1. 简介

Peewee 是一个简单而强大的 Python ORM (对象关系映射) 库，专为小型到中型项目设计。它提供了一种优雅的方式来操作关系型数据库，无需直接编写 SQL 语句，而是通过 Python 对象和方法来管理数据库操作。

## 2. 核心特点

- **轻量级**：代码库小、依赖少，安装简单
- **语法简洁**：API 设计直观，易于学习
- **支持多种数据库**：SQLite、MySQL、PostgreSQL 和 CockroachDB
- **丰富的查询构建器**：支持复杂查询、连接、聚合等
- **事务支持**：提供完整的数据库事务支持
- **模型定义灵活**：支持自定义字段类型、索引和约束
- **迁移支持**：内置简单的模式迁移工具

## 3. 安装方法

### 基本安装

```bash
pip install peewee
```

### 带数据库驱动安装

```bash
# SQLite (内置在Python中，无需额外安装)

# MySQL
pip install peewee pymysql

# PostgreSQL
pip install peewee psycopg2

# 完整安装（所有依赖）
pip install peewee[all]
```

## 4. 基础使用示例

### 定义模型

```python
from peewee import *

# 创建数据库连接
db = SqliteDatabase('my_app.db')

# 定义基础模型类
class BaseModel(Model):
    class Meta:
        database = db

# 定义用户模型
class User(BaseModel):
    username = CharField(unique=True)
    email = CharField(unique=True)
    password = CharField()
    is_active = BooleanField(default=True)
    created_at = DateTimeField(default=datetime.datetime.now)

# 定义博客文章模型
class Post(BaseModel):
    title = CharField()
    content = TextField()
    author = ForeignKeyField(User, backref='posts')
    created_at = DateTimeField(default=datetime.datetime.now)
```

### 创建表

```python
# 连接数据库
db.connect()

# 创建表
db.create_tables([User, Post])
```

### 数据操作

```python
# 创建
user = User.create(
    username='john_doe',
    email='john@example.com',
    password='secure_password'
)

# 查询单个记录
user = User.get(User.username == 'john_doe')

# 查询多个记录
active_users = User.select().where(User.is_active == True)
for user in active_users:
    print(user.username)

# 更新
query = User.update(is_active=False).where(User.username == 'john_doe')
query.execute()

# 删除
user = User.get(User.username == 'john_doe')
user.delete_instance()
```

## 5. 高级特性

### 复杂查询

```python
# 连接查询
query = (Post
         .select(Post, User)
         .join(User)
         .where(User.username == 'john_doe'))

# 聚合查询
post_count = (Post
              .select(fn.COUNT(Post.id).alias('count'))
              .where(Post.author == user)
              .scalar())
```

### 事务支持

```python
with db.atomic() as txn:
    try:
        # 执行数据库操作
        user = User.create(username='new_user', email='new@example.com')
        post = Post.create(title='My Post', content='Content', author=user)
    except:
        # 出错时回滚
        txn.rollback()
        raise
```

## 6. 与其他 ORM 对比

| 特性 | Peewee | SQLAlchemy | Django ORM |
|------|--------|------------|------------|
| 体积 | 轻量 | 重量级 | 中等 |
| 学习曲线 | 简单 | 较陡 | 中等 |
| 功能丰富度 | 适中 | 非常丰富 | 丰富 |
| 性能 | 良好 | 优秀 | 良好 |
| 独立使用 | 是 | 是 | 依赖Django |
| 社区支持 | 中等 | 强大 | 强大 |

## 7. 适用场景

- **小型到中型应用**：简单直接的 API 适合快速开发
- **原型开发**：快速设置数据库模型
- **学习 ORM 概念**：相比其他 ORM，更容易理解和掌握
- **简单数据库脚本**：轻量级特性非常适合数据处理脚本
- **微服务或 API 后端**：当不需要完整 Web 框架时很有用

## 8. 注意事项

- **大型项目可能受限**：对于非常复杂的查询和极高性能需求，SQLAlchemy 可能更合适
- **异步支持有限**：对于异步应用程序，可能需要其他库
- **不是全栈框架**：只专注于数据库交互，不提供表单验证、身份验证等功能

## 9. 文档资源

- 官方文档：http://docs.peewee-orm.com/
- GitHub 仓库：https://github.com/coleifer/peewee
- 示例项目：https://github.com/coleifer/peewee/tree/master/examples

Peewee 的名称灵感来自于小型鹦鹉，象征其轻量级和活泼的特性，是一个易于学习、快速实现数据库操作的优秀 Python ORM 库。