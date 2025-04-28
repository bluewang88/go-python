# 1. 定义并生成表
from peewee import *
import datetime

# 2. 定义数据库连接
mysqldb = MySQLDatabase('peewee', user='root', password='root', host='127.0.0.1', port=3306)