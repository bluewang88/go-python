from peewee import Model, AutoField, CharField, DateField, TextField, IntegerField, BooleanField, DateTimeField
import datetime
import sys
import os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '../..')))
from user_srv.settings import settings


class BaseModel(Model):
    class Meta:
        database = settings.DB  # 指定数据库连接
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
    email = CharField(max_length=100, unique=True)  # 邮箱
    address = CharField(max_length=100, null=True, verbose_name="地址")
    desc = TextField(null=True, verbose_name="个人简介")
    role = IntegerField(default=0, verbose_name="角色", choices=ROLE_CHOICES)  # 角色
    is_active = BooleanField(default=True, verbose_name="是否激活")  # 是否激活
    created_at = DateTimeField(default=datetime.datetime.now)  # 创建时间
    
if __name__ == "__main__":
    # 连接数据库
    settings.DB.connect()
    # 创建表
    settings.DB.create_tables([User], safe=True) #safe=True 忽略已经存在的表
    # 关闭数据库连接
    settings.DB.close()
    # 关闭数据库连接