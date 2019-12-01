# user_manage

- 安装依赖
```
git clone git@github.com:yuzhikuan/user_manage.git
cd user_manage
go mod tidy
```

- mysql配置
创建user表,并确保正确配置 mysql_map.toml：
```
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
  `addr` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
  `age` smallint(4) NOT NULL DEFAULT '0' COMMENT '年龄',
  `birth` varchar(100) NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '生日',
  `sex` smallint(4) NOT NULL DEFAULT '0' COMMENT '性别',
  `update_at` varchar(100) NOT NULL DEFAULT '1970-01-01 00:00:00',
  `create_at` varchar(100) NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 COMMENT='user'
```