# messengerBot

## 本地开发

### 环境准备

#### 0. go开发环境

http://docscn.studygolang.com

#### 1. 启动本地mysql，推荐docker方式

```shell
docker run -itd --name mine-mysql -p 13306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7.38
```

#### 2. 启动本地redis，推荐docker方式

```shell
docker run -itd --name mine-redis -p 16379:6379 redis:5.0.14
```

#### 3. 安装wire，用于管理依赖注入

```shell
go install github.com/google/wire/cmd/wire@latest
```

#### 4. 安装swagger，用于文档生成

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

#### 5. 配置更改
配置文件在`./configs/config.yaml`中维护

同时需要更新`./conf/config.go`中的配置结构体，即可全局读取，参考main.go中的使用

### CRUD

#### 1. 创建 DB Table

DDL语句在`./configs/ddl.sql`中维护
```sql
create
database IF NOT EXISTS `messengerBot`;

use
`messengerBot`;

CREATE TABLE `user`
(
    `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '',
    `user_id`       bigint(20) unsigned NOT NULL COMMENT '',
    `email`         varchar(128) NOT NULL COMMENT '',
    `password`      varchar(256) NOT NULL COMMENT '',
    `password_salt` varchar(32)  NOT NULL,
    `nickname`      varchar(36)  NOT NULL DEFAULT '' COMMENT '',
    `avatar_key`    varchar(512) NOT NULL DEFAULT '' COMMENT '',
    `created_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '',
    `updated_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_uid` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='user table';

```
#### 2. 生成dal代码

这个脚本会根据数据库配置，在dal目录下生成操作代码

```shell
sh scripts/dal.sh
```

#### 3. 编写service及router

参考user和post相关代码

#### 4. 通过wire生成依赖注入代码

```shell
wire
```

#### 5. 通过swag生成swagger文档

```shell
swag fmt
swag init
```

#### 6. 通过浏览器访问swagger

运行程序
```shell
sh scripts/dev.sh
```

#### 7. 通过浏览器访问swagger

http://localhost:8080/swagger/index.html

## 打包部署

### Dev

```shell
sh scripts/build.sh
```

### Docker

打包为Docker镜像

```shell
docker build -t messengerBot:v1 .
```
## 构建运行

### Dev

```shell
sh scripts/dev.sh
```

### Test

```shell
sh scripts/test.sh
```

### Docker

运行Docker容器(Prod)

```shell
docker run --name messengerBot -v /etc/localtime:/etc/localtime -v /var/logs/messengerBot:/go/messengerBot/logs -d -p 8888:80 messengerBot:v1 --env prod
```