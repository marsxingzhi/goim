// 创建user表
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `uid` bigint unsigned NOT NULL COMMENT '用户ID 系统生成',
  `aid` varchar(32) NOT NULL DEFAULT '' COMMENT '账户ID 用户设置',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `did` varchar(40) NOT NULL DEFAULT '' COMMENT '注册设备唯一标识',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '用户状态',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称',
  `firstname` varchar(32) NOT NULL DEFAULT '' COMMENT 'firstname',
  `lastname` varchar(32) NOT NULL DEFAULT '' COMMENT 'lastname',
  `gender` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '性别',
  `birth` bigint NOT NULL DEFAULT '0' COMMENT '生日',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT 'Email',
  `mobile` varchar(32) NOT NULL DEFAULT '' COMMENT '手机号',
  `platform` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '注册平台/登录平台',
  `server_id` int NOT NULL DEFAULT '0' COMMENT '分配的ws服务器',
  `city_id` int NOT NULL DEFAULT '0' COMMENT '城市ID',
  `avatar_key` varchar(64) NOT NULL DEFAULT '' COMMENT '小图 72*72',
  `created` bigint NOT NULL DEFAULT '0',
  `updated` bigint NOT NULL DEFAULT '0',
  `deleted` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `uniq_aid_deleted` (`aid`,`deleted`),
  UNIQUE KEY `uniq_mobile_deleted` (`mobile`,`deleted`),
  KEY `idx_gender` (`gender`),
  KEY `idx_cityId` (`city_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



