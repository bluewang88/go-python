from peewee import Model, AutoField, CharField, DateField, TextField, IntegerField, BooleanField, DateTimeField
import datetime
import sys
import os
# 将当前文件的上级目录添加到 sys.path 中
# 导入路径的第一部分需要直接位于 sys.path 中的某个目录下，而非该目录本身
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '../..'))) # No longer needed with relative import
# print("项目根目录",os.path.abspath(os.path.join(os.path.dirname(__file__), '../..')))
# print("当前路径:", os.path.abspath(os.path.dirname(__file__)))
# Python 解释器对于绝对导入的处理是从当前目录、sys.path、环境变量 PYTHONPATH 中搜索需要导入的包和模块。
# print("系统 sys.path：",sys.path)
# from user_srv.settings import settings
# from mxshop_srvs.user_srv import settings
# from user_srv.settings import * 

# 移除 sys.path.append 相关代码，使用绝对导入
from user_srv.settings import settings
from passlib.hash import pbkdf2_sha256




class BaseModel(Model):
    class Meta:
        database = settings.MYSQL_CONN # 指定数据库连接
        # 这里的 database 是在 settings.py 中定义的数据库连接对象
        
        


class User(BaseModel):
    GENDER_CHOICES = (
        ('male', '男'),
        ('female', '女'),
    )

    ROLE_CHOICES = (
        (0, '普通用户'),
        (1, '管理员'),
    )
    id = AutoField()  # 自增主键
    username = CharField(max_length=20, unique=True)  # 用户名
    nick_name = CharField(max_length=20, null=True, unique=True, verbose_name="昵称")  # 昵称
    mobile = CharField(max_length=11, index=True, unique=True, verbose_name="手机号")  # 手机号
    head_img = CharField(max_length=100, null=True, verbose_name="头像")
    age = CharField(default=18, max_length=20, verbose_name="年龄")  # 年龄
    gender = CharField(max_length=6, null=True, verbose_name="性别", choices=GENDER_CHOICES)  # 性别
    birthday = DateField(null=True, verbose_name="生日")
    password = CharField(max_length=100, verbose_name="密码")  # 密码
    # salt = CharField(max_length=100, null=True, verbose_name="密码盐")
    email = CharField(max_length=100, unique=True)  # 邮箱
    address = CharField(max_length=100, null=True, verbose_name="地址")
    desc = TextField(null=True, verbose_name="个人简介")
    role = IntegerField(default=0, verbose_name="角色", choices=ROLE_CHOICES)  # 角色
    is_active = BooleanField(default=True, verbose_name="是否激活")  # 是否激活
    created_at = DateTimeField(default=datetime.datetime.now)  # 创建时间
    
if __name__ == "__main__":
    # 连接数据库
    settings.MYSQL_CONN.connect()
    # 创建表
    settings.MYSQL_CONN.create_tables([User], safe=True) #safe=True 忽略已经存在的表
    # 关闭数据库连接
    # settings.MYSQL_CONN.close()
    # 关闭数据库连接
    
    # import hashlib # hashlib 是 Python 内置的加密模块
    # 创建 md5 对象
    # md5 是一种不可逆的加密算法
    # 通过 update 方法更新要加密的字符串
    # 这里的字符串需要是字节类型，所以要加 b 前缀
    # 也可以使用 encode 方法将字符串转换为字节类型
    # 例如 "123456".encode('utf-8')
    # m = hashlib.md5() 
    # m.update(b"123456")
    # print(m.hexdigest())  # 输出加密后的密码
    
    #先执行 pip install passlib
    # from passlib.hash import pbkdf2_sha256
    # 创建一个加密对象
    # 加密对象可以使用不同的加密算法
    # 这里使用 pbkdf2_sha256 算法
    # 该算法是基于 sha256 的加密算法
    # hash = pbkdf2_sha256.hash("123456")
    # print(hash) # 输出加密后的密码
    # 验证密码
    # 验证密码是否正确
    # 这里的 hash 是加密后的密码
    # 这里的 "123456" 是用户输入的密码
    # 如果密码正确，返回 True
    # 如果密码错误，返回 False
    # print(pbkdf2_sha256.verify("123456", hash)) # 输出 True
    # print(pbkdf2_sha256.verify("1234567", hash)) # 输出 False
    
    #创建用户
    
    for i in range(10):
        user = User() # 创建一个用户对象
        user.username = f"user_name{i}"   
        user.nick_name = f"bobby{i}"
        user.age = 18 + i
        user.mobile = f"138123456{i}"
        user.password = pbkdf2_sha256.hash("admin123")
        user.email = f"user{i}@example.com"
        user.is_active = True
        user.save()