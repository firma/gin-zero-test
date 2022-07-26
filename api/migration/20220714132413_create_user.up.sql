-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict`  (
                              `id`  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
                              `data_type` varchar(320) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数字类型',
                              `item_key` varchar(320) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '键',
                              `item_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '值',
                              `xh` int(11) NULL DEFAULT NULL COMMENT '排序序号',
                              `out_data` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '外部数据',
                              `p_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0'  COMMENT '父ID',
                              `dict_desc` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典描述',
                              `dict_level` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '为空时则是全局字典 否则为某个学校的全局 字典权限',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
                              `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
                              `menu_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源名称',
                              `menu_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源类型 0：菜单  1：按钮 2：其他',
                              `p_menu_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0'  COMMENT '父资源ID',
                              `menu_icon` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源icon',
                              `xh` int(11) NULL DEFAULT NULL COMMENT '排序值',
                              `auth_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '位权限编码',
                              `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
                              `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '过滤URI',
                              `redirect` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '重定向URL',
                              `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
                              `if_leaf` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '是否是叶子节点',
                              PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'SAAS资源表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_org
-- ----------------------------
DROP TABLE IF EXISTS `sys_org`;
CREATE TABLE `sys_org`  (
                             `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    -- `org_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主键ID',
                             `org_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织名称',
                             `p_org_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0'  COMMENT '父节点ID',
                             `org_flag` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织标记',
                             `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             PRIMARY KEY (`org_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'SAAS组织机构' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
                              `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色权限ID',
                              `role_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色权限名称',
                              `parent_role_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '父角色权限ID',
                              `role_flag` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色标记',
                              `role_level` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色级别',
                              `role_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限编码',
                              `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'SAAS角色权限' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
                                   `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
                                   `role_id`   BIGINT UNSIGNED NOT NULL DEFAULT '0'  DEFAULT NULL,
                                   `menu_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0'  COMMENT '资源ID',
                                   `type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '角色菜单关系类型(0:该角色拥有此菜单或按钮,1:该角色拥有此菜单的子菜单或权限)'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'SAAS角色菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
                              `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
                              `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '手机号码',
                              `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '密码',
                              `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
                              `user_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户姓名',
                              `user_head_url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户头像',
                              `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '1' COMMENT '用户状态(0:禁用,1:正常)',
                              `user_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '1' COMMENT '用户类型(0:管理员,1:巡视员)',
                              `login_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
                              `org_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '组织ID',
                              `room_ids` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '警务室ID',
                              PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = 'SAAS用户管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
                                   `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
                                   `user_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0'  COMMENT '用户ID',
                                   `account_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0'  COMMENT '账户ID',
                                   `role_id`  BIGINT UNSIGNED NOT NULL DEFAULT '0'  COMMENT '角色权限ID',
                                   INDEX `IDX_USERID`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'SAAS用户权限表' ROW_FORMAT = DYNAMIC;

