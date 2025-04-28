# 1. 定义并生成表
from peewee import *
import datetime

# 使用 logger
import logging
logger  = logging.getLogger('peewee')
logger.setLevel(logging.DEBUG)
# 创建一个控制台处理器
consoleHandler = logging.StreamHandler()
logger.addHandler(consoleHandler)

# 2. 定义数据库连接
mysqldb = MySQLDatabase('peewee', user='root', password='root', host='127.0.0.1', port=3306)

# 3. 定义模型类
class User(Model):
    #在 peewee中如果没有设置主键，peewee会自动创建一个 id 主键
    id = AutoField()  # 自增主键
    username = CharField(max_length=20, unique=True)  # 用户名
    age = CharField(default=18, max_length=20,verbose_name="年龄")  # 年龄
    password = CharField(max_length=50)  # 密码
    email = CharField(max_length=100, unique=True)  # 邮箱
    created_at = DateTimeField(default=datetime.datetime.now)  # 创建时间

    class Meta:
        database = mysqldb  # 指定数据库连接
        table_name = 'user'  # 指定表名
        
class Tweet(Model):
    id = AutoField()  # 自增主键
    user = ForeignKeyField(User, backref='tweets')  # 外键关联用户
    message = TextField()  # 消息内容
    created_date = DateTimeField(default=datetime.datetime.now)  # 创建时间
    is_published = BooleanField(default=True)  # 是否发表

    class Meta:
        database = mysqldb  # 指定数据库连接
        table_name = 'tweet'  # 指定表名
        
# 4. 创建表

def create_tables():
    # 连接数据库
    mysqldb.connect()
    # 创建表
    mysqldb.create_tables([User, Tweet], safe=True)
    # 关闭数据库连接
    mysqldb.close()
    
# 5.添加数据
def add_data():
    # 连接数据库
    mysqldb.connect()
    # 添加数据
    user = User.create(username='testuser', password='password', email='testuser@example.com')

    mysqldb.close() 
    
# 6. 查询数据
def query_data():
    # 连接数据库
    mysqldb.connect()
    # select() 查询User 表中的所有数据
    # select不会抛异常
    # 主要用于组装 sql
    users = User.select()
    for user in users:
        print("使用 select 方式查询表数据",user.username, user.password, user.email, user.created_at)
        
    # get()查询 User 表
    # 1. get方法返回的是 user对象
    # 2. get方法查询条件是唯一的
    # 3. get方法如果查询不到会抛出异常
    try:
        getUser = User.get(User.username == 'testuser')
    except User.DoesNotExist as e:
        print("用户不存在")
    # getUser = User.get(User.username == 'testuser')
    print("使用 get 查询表数据",getUser.username, getUser.password, getUser.email, getUser.created_at)
    
    # 关闭数据库连接
    mysqldb.close()
    
# 7. 更新数据
def update_data():
    # 连接数据库
    mysqldb.connect()
    # 更新数据
    user = User.get(User.username == 'testuser')
    user.password = 'newpassword'
    user.save()  # 保存更新

#  运行
if __name__ == '__main__':
    create_tables()
    print("表创建成功")
    add_data()
    print("数据添加成功")
    query_data()
    print("数据查询成功")
        

