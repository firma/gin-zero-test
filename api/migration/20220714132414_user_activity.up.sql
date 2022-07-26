CREATE TABLE `c_user_activity`
(
    `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `latest_login_at` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '最近登陆',
    `create_time`   BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
    `update_time`   BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
    `modify_time`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '维护字段-更新时间',
    `delete_time`   BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_user_id` (`user_id`)
) ENGINE = InnoDB COMMENT ='用户活动';
