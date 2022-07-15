CREATE TABLE `user`
(
    `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `username` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户名',
    `create_time`   BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
    `update_time`   BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
    `modify_time`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '维护字段-更新时间',
    `delete_time`   BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_username` (`username`)
) ENGINE = InnoDB COMMENT ='用户表';
