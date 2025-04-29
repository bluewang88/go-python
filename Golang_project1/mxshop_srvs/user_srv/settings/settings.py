from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin

class ReconnectMySQLDatabase(ReconnectMixin, PooledMySQLDatabase): # type: ignore
    pass

# Database configuration
MYSQL_DB = "mxshop_user_srv"
MYSQL_USER = "root"
MYSQL_PASSWORD = "root"
MYSQL_HOST = "127.0.0.1"
MYSQL_PORT = 13306
MYSQL_CONN = ReconnectMySQLDatabase(
    MYSQL_DB,
    max_connections=32,
    stale_timeout=300,
    user=MYSQL_USER,
    password=MYSQL_PASSWORD,
    host=MYSQL_HOST,
    port=MYSQL_PORT,
    charset="utf8mb4",
    use_unicode=True,
    autocommit=True,
)