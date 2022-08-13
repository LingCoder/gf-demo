CREATE DATABASE IF NOT EXISTS `demo` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_0900_ai_ci';

USE demo;


CREATE TABLE IF NOT EXISTS `sys_user`
(
    `id`             bigint       NOT NULL AUTO_INCREMENT COMMENT 'user id',
    `nick_name`      varchar(128) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `email`          varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱地址',
    `avatar_picture` varchar(100) NOT NULL DEFAULT '' COMMENT '用户头像',
    `is_locked`      tinyint(1)   NOT NULL DEFAULT 0 COMMENT '是否锁定(1-是,0-否)',
    `is_deleted`     tinyint(1)   NOT NULL DEFAULT '0' COMMENT '是否删除(1-是,0-否)',
    `locked_time`    datetime     NULL COMMENT '锁定时间',
    `deleted_time`   datetime     NULL COMMENT '删除时间',
    `create_time`    datetime     NULL     DEFAULT now() COMMENT '创建时间',
    `update_time`    datetime     NULL     DEFAULT now() COMMENT '更新时间',
    `remark`         varchar(500) NOT NULL DEFAULT '' COMMENT '备注',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT = '用户信息表'
  ROW_FORMAT = Dynamic;
