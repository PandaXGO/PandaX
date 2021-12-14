/*
 Navicat Premium Data Transfer

 Source Server         : localhost_mysql
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : pandax

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 12/12/2021 17:44:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 134 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api', 'POST', '', '', '', 71);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api', 'PUT', '', '', '', 70);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/:id', 'DELETE', '', '', '', 69);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/:id', 'GET', '', '', '', 72);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/all', 'GET', '', '', '', 74);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/getPolicyPathByRoleId', 'GET', '', '', '', 73);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/list', 'GET', '', '', '', 75);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'POST', '', '', '', 78);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'PUT', '', '', '', 77);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/:configId', 'DELETE', '', '', '', 76);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/:configId', 'GET', '', '', '', 79);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/configKey', 'GET', '', '', '', 80);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/list', 'GET', '', '', '', 81);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept', 'POST', '', '', '', 85);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept', 'PUT', '', '', '', 84);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/:deptId', 'DELETE', '', '', '', 83);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/:deptId', 'GET', '', '', '', 87);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/deptTree', 'GET', '', '', '', 82);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/list', 'GET', '', '', '', 88);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', '', 86);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'POST', '', '', '', 90);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'PUT', '', '', '', 91);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/:dictCode', 'DELETE', '', '', '', 92);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/:dictCode', 'GET', '', '', '', 89);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/list', 'GET', '', '', '', 94);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/type', 'GET', '', '', '', 93);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'POST', '', '', '', 97);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'PUT', '', '', '', 100);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/:dictId', 'DELETE', '', '', '', 96);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/:dictId', 'GET', '', '', '', 98);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/export', 'GET', '', '', '', 95);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/list', 'GET', '', '', '', 99);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'POST', '', '', '', 103);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'PUT', '', '', '', 102);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/:menuId', 'DELETE', '', '', '', 101);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/:menuId', 'GET', '', '', '', 104);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/list', 'GET', '', '', '', 105);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuPaths', 'GET', '', '', '', 106);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuRole', 'GET', '', '', '', 108);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuTreeSelect', 'GET', '', '', '', 109);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', '', 107);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'POST', '', '', '', 110);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'PUT', '', '', '', 111);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/:postId', 'DELETE', '', '', '', 112);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/:postId', 'GET', '', '', '', 113);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/list', 'GET', '', '', '', 114);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'POST', '', '', '', 117);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'PUT', '', '', '', 118);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/:roleId', 'DELETE', '', '', '', 119);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/:roleId', 'GET', '', '', '', 116);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/changeStatus', 'PUT', '', '', '', 120);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/dataScope', 'PUT', '', '', '', 121);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/export', 'GET', '', '', '', 122);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/list', 'GET', '', '', '', 115);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'POST', '', '', '', 133);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'PUT', '', '', '', 132);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/:userId', 'DELETE', '', '', '', 124);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/avatar', 'POST', '', '', '', 125);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/changeStatus', 'PUT', '', '', '', 123);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/export', 'GET', '', '', '', 131);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getById/:userId', 'GET', '', '', '', 127);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getInit', 'GET', '', '', '', 128);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getRoPo', 'GET', '', '', '', 129);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/list', 'GET', '', '', '', 130);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/pwd', 'PUT', '', '', '', 126);

-- ----------------------------
-- Table structure for log_logins
-- ----------------------------
DROP TABLE IF EXISTS `log_logins`;
CREATE TABLE `log_logins`  (
  `info_id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `ipaddr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '系统',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '固件',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 153 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for log_opers
-- ----------------------------
DROP TABLE IF EXISTS `log_opers`;
CREATE TABLE `log_opers`  (
  `oper_id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作的模块',
  `business_type` int NULL DEFAULT NULL COMMENT '0其它 1新增 2修改 3删除',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `oper_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作人员',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作url',
  `oper_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作IP',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作地点',
  `oper_param` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求参数',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '0=正常,1=异常',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log_opers
-- ----------------------------

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2021-12-09 09:21:04', '2021-12-09 09:21:04', NULL, '/system/user/list', '查询用户列表（分页）', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (2, '2021-12-09 09:29:36', '2021-12-09 09:29:36', NULL, '/system/user/changeStatus', '修改用户状态', 'user', 'PUT');
INSERT INTO `sys_apis` VALUES (3, '2021-12-09 09:34:37', '2021-12-09 09:34:37', NULL, '/system/user/:userId', '删除用户信息', 'user', 'DELETE');
INSERT INTO `sys_apis` VALUES (4, '2021-12-09 09:36:43', '2021-12-09 09:36:43', NULL, '/system/dept/list', '获取部门列表', 'dept', 'GET');
INSERT INTO `sys_apis` VALUES (5, '2021-12-09 09:37:31', '2021-12-09 09:37:31', NULL, '/system/dept/:deptId', '获取部门信息', 'dept', 'GET');
INSERT INTO `sys_apis` VALUES (6, '2021-12-09 18:20:32', '2021-12-09 18:20:32', NULL, '/system/user/avatar', '上传头像', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (7, '2021-12-09 18:21:10', '2021-12-09 18:21:10', NULL, '/system/user/pwd', '修改密码', 'user', 'PUT');
INSERT INTO `sys_apis` VALUES (8, '2021-12-09 18:21:54', '2021-12-09 18:21:54', NULL, '/system/user/getById/:userId', '通过id获取用户信息', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (9, '2021-12-09 18:58:50', '2021-12-09 18:58:50', NULL, '/system/user/getInit', '获取初始化角色岗位信息(添加用户初始化)', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (10, '2021-12-09 18:59:43', '2021-12-09 18:59:43', NULL, '/system/user/getRoPo', '获取用户角色岗位信息', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (11, '2021-12-09 19:00:24', '2021-12-09 19:00:24', NULL, '/system/user', '添加用户信息', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2021-12-09 19:00:52', '2021-12-09 19:00:52', NULL, '/system/user', '修改用户信息', 'user', 'PUT');
INSERT INTO `sys_apis` VALUES (13, '2021-12-09 19:02:30', '2021-12-09 19:02:30', NULL, '/system/user/export', '导出用户信息', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (14, '2021-12-09 19:04:04', '2021-12-09 19:04:04', NULL, '/system/dept/roleDeptTreeSelect/:roleId', '获取角色部门树', 'dept', 'GET');
INSERT INTO `sys_apis` VALUES (15, '2021-12-09 19:04:48', '2021-12-09 19:04:48', NULL, '/system/dept/deptTree', '获取所有部门树', 'dept', 'GET');
INSERT INTO `sys_apis` VALUES (16, '2021-12-09 19:07:37', '2021-12-09 19:07:37', NULL, '/system/dept', '添加部门信息', 'dept', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2021-12-09 19:08:14', '2021-12-09 19:08:14', NULL, '/system/dept', '修改部门信息', 'dept', 'PUT');
INSERT INTO `sys_apis` VALUES (18, '2021-12-09 19:08:40', '2021-12-09 19:08:40', NULL, '/system/dept/:deptId', '删除部门信息', 'dept', 'DELETE');
INSERT INTO `sys_apis` VALUES (19, '2021-12-09 19:09:41', '2021-12-09 19:09:41', NULL, '/system/config/list', '获取配置分页列表', 'config', 'GET');
INSERT INTO `sys_apis` VALUES (20, '2021-12-09 19:10:11', '2021-12-09 19:10:11', NULL, '/system/config/configKey', '获取配置列表通过ConfigKey', 'config', 'GET');
INSERT INTO `sys_apis` VALUES (21, '2021-12-09 19:10:45', '2021-12-09 19:10:45', NULL, '/system/config/:configId', '获取配置信息', 'config', 'GET');
INSERT INTO `sys_apis` VALUES (22, '2021-12-09 19:11:22', '2021-12-09 19:11:22', NULL, '/system/config', '添加配置信息', 'config', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2021-12-09 19:11:41', '2021-12-09 19:11:41', NULL, '/system/config', '修改配置信息', 'config', 'PUT');
INSERT INTO `sys_apis` VALUES (24, '2021-12-09 19:12:28', '2021-12-09 19:12:28', NULL, '/system/config/:configId', '删除配置信息', 'config', 'DELETE');
INSERT INTO `sys_apis` VALUES (25, '2021-12-09 19:13:08', '2021-12-09 19:13:08', NULL, '/system/dict/type/list', '获取字典类型分页列表', 'dict', 'GET');
INSERT INTO `sys_apis` VALUES (26, '2021-12-09 19:13:55', '2021-12-09 19:13:55', NULL, '/system/dict/type/:dictId', '获取字典类型信息', 'dict', 'GET');
INSERT INTO `sys_apis` VALUES (27, '2021-12-09 19:14:28', '2021-12-09 19:14:28', NULL, '/system/dict/type', '添加字典类型信息', 'dict', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2021-12-09 19:14:55', '2021-12-09 19:14:55', NULL, '/system/dict/type', '修改字典类型信息', 'dict', 'PUT');
INSERT INTO `sys_apis` VALUES (29, '2021-12-09 19:15:17', '2021-12-09 19:15:17', NULL, '/system/dict/type/:dictId', '删除字典类型信息', 'dict', 'DELETE');
INSERT INTO `sys_apis` VALUES (30, '2021-12-09 19:15:50', '2021-12-09 19:15:50', NULL, '/system/dict/type/export', '导出字典类型信息', 'dict', 'GET');
INSERT INTO `sys_apis` VALUES (31, '2021-12-09 19:16:26', '2021-12-09 19:16:26', NULL, '/system/dict/data/list', '获取字典数据分页列表', 'dict', 'GET');
INSERT INTO `sys_apis` VALUES (32, '2021-12-09 19:17:01', '2021-12-09 19:17:01', NULL, '/system/dict/data/type', '获取字典数据列表通过字典类型', 'dict', 'GET');
INSERT INTO `sys_apis` VALUES (33, '2021-12-09 19:17:39', '2021-12-09 19:17:39', NULL, '/system/dict/data/:dictCode', '获取字典数据信息', 'dict', 'GET');
INSERT INTO `sys_apis` VALUES (34, '2021-12-09 19:18:20', '2021-12-09 19:18:20', NULL, '/system/dict/data', '添加字典数据信息', 'dict', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2021-12-09 19:18:44', '2021-12-09 19:18:44', NULL, '/system/dict/data', '修改字典数据信息', 'dict', 'PUT');
INSERT INTO `sys_apis` VALUES (36, '2021-12-09 19:19:16', '2021-12-09 19:19:16', NULL, '/system/dict/data/:dictCode', '删除字典数据信息', 'dict', 'DELETE');
INSERT INTO `sys_apis` VALUES (37, '2021-12-09 19:21:18', '2021-12-09 19:21:18', NULL, '/system/menu/menuTreeSelect', '获取菜单树', 'menu', 'GET');
INSERT INTO `sys_apis` VALUES (38, '2021-12-09 19:21:47', '2021-12-09 19:21:47', NULL, '/system/menu/menuRole', '获取角色菜单', 'menu', 'GET');
INSERT INTO `sys_apis` VALUES (39, '2021-12-09 19:22:42', '2021-12-09 19:22:42', NULL, '/system/menu/roleMenuTreeSelect/:roleId', '获取角色菜单树', 'menu', 'GET');
INSERT INTO `sys_apis` VALUES (40, '2021-12-09 19:23:17', '2021-12-09 19:23:17', NULL, '/system/menu/menuPaths', '获取角色菜单路径列表', 'menu', 'GET');
INSERT INTO `sys_apis` VALUES (41, '2021-12-09 19:23:40', '2021-12-09 19:23:40', NULL, '/system/menu/list', '获取菜单列表', 'menu', 'GET');
INSERT INTO `sys_apis` VALUES (42, '2021-12-09 19:24:09', '2021-12-09 19:24:09', NULL, '/system/menu/:menuId', '获取菜单信息', 'menu', 'GET');
INSERT INTO `sys_apis` VALUES (43, '2021-12-09 19:24:29', '2021-12-09 19:24:29', NULL, '/system/menu', '添加菜单信息', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (44, '2021-12-09 19:24:48', '2021-12-09 19:24:48', NULL, '/system/menu', '修改菜单信息', 'menu', 'PUT');
INSERT INTO `sys_apis` VALUES (45, '2021-12-09 19:25:10', '2021-12-09 19:25:10', NULL, '/system/menu/:menuId', '删除菜单信息', 'menu', 'DELETE');
INSERT INTO `sys_apis` VALUES (46, '2021-12-09 19:25:44', '2021-12-09 19:27:06', NULL, '/system/post/list', '获取岗位分页列表', 'post', 'GET');
INSERT INTO `sys_apis` VALUES (47, '2021-12-09 19:26:55', '2021-12-09 19:26:55', NULL, '/system/post/:postId', '获取岗位信息', 'post', 'GET');
INSERT INTO `sys_apis` VALUES (48, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/post', '添加岗位信息', 'post', 'POST');
INSERT INTO `sys_apis` VALUES (49, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/post', '修改岗位信息', 'post', 'PUT');
INSERT INTO `sys_apis` VALUES (50, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/post/:postId', '删除岗位信息', 'post', 'DELETE');
INSERT INTO `sys_apis` VALUES (51, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/list', '获取角色分页列表', 'role', 'GET');
INSERT INTO `sys_apis` VALUES (52, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/:roleId', '获取角色信息', 'role', 'GET');
INSERT INTO `sys_apis` VALUES (53, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role', '添加角色信息', 'role', 'POST');
INSERT INTO `sys_apis` VALUES (54, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role', '修改角色信息', 'role', 'PUT');
INSERT INTO `sys_apis` VALUES (55, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/:roleId', '删除角色信息', 'role', 'DELETE');
INSERT INTO `sys_apis` VALUES (56, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/changeStatus', '修改角色状态', 'role', 'PUT');
INSERT INTO `sys_apis` VALUES (57, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/dataScope', '修改角色部门权限', 'role', 'PUT');
INSERT INTO `sys_apis` VALUES (58, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/export', '导出角色信息', 'role', 'GET');
INSERT INTO `sys_apis` VALUES (59, '2021-12-09 19:50:57', '2021-12-09 19:50:57', NULL, '/system/api/list', '获取api分页列表', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (60, '2021-12-09 19:51:26', '2021-12-09 19:51:26', NULL, '/system/api/all', '获取所有api', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (61, '2021-12-09 19:51:54', '2021-12-09 19:51:54', NULL, '/system/api/getPolicyPathByRoleId', '获取角色拥有的api权限', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (62, '2021-12-09 19:52:14', '2021-12-09 19:52:14', NULL, '/system/api/:id', '获取api信息', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (63, '2021-12-09 19:52:35', '2021-12-09 19:52:35', NULL, '/system/api', '添加api信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (64, '2021-12-09 19:52:50', '2021-12-09 19:52:50', NULL, '/system/api', '修改api信息', 'api', 'PUT');
INSERT INTO `sys_apis` VALUES (65, '2021-12-09 19:53:07', '2021-12-09 19:53:07', NULL, '/system/api/:id', '删除api信息', 'api', 'DELETE');

-- ----------------------------
-- Table structure for sys_configs
-- ----------------------------
DROP TABLE IF EXISTS `sys_configs`;
CREATE TABLE `sys_configs`  (
  `config_id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '是否系统内置0，1',
  `is_frontend` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Remark',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_configs
-- ----------------------------
INSERT INTO `sys_configs` VALUES (1, '账号初始密码', 'sys.user.initPassword', '123456', '0', '0', '初始密码', '2021-12-04 13:50:13', '2021-12-04 13:54:52', NULL);

-- ----------------------------
-- Table structure for sys_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_depts`;
CREATE TABLE `sys_depts`  (
  `dept_id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` int NULL DEFAULT NULL COMMENT '上级部门',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门路径',
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门名称',
  `sort` int NULL DEFAULT NULL COMMENT '排序',
  `leader` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '修改人',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_depts
-- ----------------------------
INSERT INTO `sys_depts` VALUES (2, 0, '/0/2', '熊猫科技', 0, 'xm', '18353366836', '342@qq.com', '0', 'admin', 'admin', '2021-12-01 17:31:53', '2021-12-02 08:56:19', NULL);
INSERT INTO `sys_depts` VALUES (3, 2, '/3', '研发部', 1, 'panda', '18353366543', 'ewr@qq.com', '0', 'admin', 'admin', '2021-12-01 17:37:43', '2021-12-02 08:55:56', NULL);
INSERT INTO `sys_depts` VALUES (4, 2, '/0/2/4', '营销部', 0, 'lisa', '18653366543', '', '1', 'admin', 'admin', '2021-12-01 17:31:53', '2021-12-02 09:02:20', '2021-12-02 09:02:46');

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint NOT NULL AUTO_INCREMENT,
  `dict_sort` int NULL DEFAULT NULL COMMENT '排序',
  `dict_label` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `dict_value` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '值',
  `dict_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典类型',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态（0正常 1停用）',
  `css_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'CssClass',
  `list_class` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ListClass',
  `is_default` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'IsDefault',
  `create_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '男', '0', 'sys_user_sex', '0', '', '', '', 'admin', '', '男', '2021-11-30 14:58:18', '2021-11-30 14:58:18', NULL);
INSERT INTO `sys_dict_data` VALUES (2, 1, '女', '1', 'sys_user_sex', '0', '', '', '', 'admin', '', '女生', '2021-11-30 15:09:11', '2021-11-30 15:10:28', NULL);
INSERT INTO `sys_dict_data` VALUES (3, 2, '未知', '2', 'sys_user_sex', '0', '', '', '', 'admin', '', '未知', '2021-11-30 15:09:11', '2021-11-30 15:10:28', NULL);
INSERT INTO `sys_dict_data` VALUES (4, 0, '正常', '0', 'sys_normal_disable', '0', '', '', '', 'admin', '', '', '2021-12-01 15:58:50', '2021-12-01 15:58:50', NULL);
INSERT INTO `sys_dict_data` VALUES (5, 1, '停用', '1', 'sys_normal_disable', '0', '', '', '', 'admin', '', '', '2021-12-01 15:59:08', '2021-12-01 15:59:08', NULL);
INSERT INTO `sys_dict_data` VALUES (6, 0, '目录', 'M', 'sys_menu_type', '0', '', '', '', 'admin', '', '', '2021-12-02 09:49:12', '2021-12-02 09:49:12', NULL);
INSERT INTO `sys_dict_data` VALUES (7, 1, '菜单', 'C', 'sys_menu_type', '0', '', '', '', 'admin', '', '', '2021-12-02 09:49:35', '2021-12-02 09:49:52', NULL);
INSERT INTO `sys_dict_data` VALUES (8, 2, '按钮', 'F', 'sys_menu_type', '0', '', '', '', 'admin', '', '', '2021-12-02 09:49:35', '2021-12-02 09:49:35', NULL);
INSERT INTO `sys_dict_data` VALUES (9, 0, '显示', '0', 'sys_show_hide', '0', '', '', '', 'admin', '', '', '2021-12-02 09:56:40', '2021-12-02 09:56:40', NULL);
INSERT INTO `sys_dict_data` VALUES (10, 0, '隐藏', '1', 'sys_show_hide', '0', '', '', '', 'admin', '', '', '2021-12-02 09:56:52', '2021-12-02 09:56:52', NULL);
INSERT INTO `sys_dict_data` VALUES (11, 0, '是', '0', 'sys_num_yes_no', '0', '', '', '', 'admin', '', '', '2021-12-02 10:16:16', '2021-12-02 10:16:16', NULL);
INSERT INTO `sys_dict_data` VALUES (12, 1, '否', '1', 'sys_num_yes_no', '0', '', '', '', 'admin', '', '', '2021-12-02 10:16:26', '2021-12-02 10:16:26', NULL);
INSERT INTO `sys_dict_data` VALUES (13, 0, '是', '0', 'sys_yes_no', '0', '', '', '', 'admin', '', '', '2021-12-04 13:48:15', '2021-12-04 13:48:15', NULL);
INSERT INTO `sys_dict_data` VALUES (14, 0, '否', '1', 'sys_yes_no', '0', '', '', '', 'admin', '', '', '2021-12-04 13:48:21', '2021-12-04 13:48:21', NULL);
INSERT INTO `sys_dict_data` VALUES (15, 0, '创建(POST)', 'POST', 'sys_method_api', '0', '', '', '', 'admin', '', '', '2021-12-08 17:22:05', '2021-12-09 09:29:52', NULL);
INSERT INTO `sys_dict_data` VALUES (16, 1, '查询(GET)', 'GET', 'sys_method_api', '0', '', '', '', 'admin', '', '', '2021-12-08 17:22:24', '2021-12-09 09:29:59', NULL);
INSERT INTO `sys_dict_data` VALUES (17, 2, '修改(PUT)', 'PUT', 'sys_method_api', '0', '', '', '', 'admin', '', '', '2021-12-08 17:22:40', '2021-12-09 09:30:06', NULL);
INSERT INTO `sys_dict_data` VALUES (18, 3, '删除(DELETE)', 'DELETE', 'sys_method_api', '0', '', '', '', 'admin', '', '', '2021-12-08 17:22:54', '2021-12-09 09:30:13', NULL);

-- ----------------------------
-- Table structure for sys_dict_types
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_types`;
CREATE TABLE `sys_dict_types`  (
  `dict_id` bigint NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `dict_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `create_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`dict_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_types
-- ----------------------------
INSERT INTO `sys_dict_types` VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '', '性别列表', '2021-11-30 14:02:52', '2021-11-30 14:07:55', '2021-11-30 14:11:54');
INSERT INTO `sys_dict_types` VALUES (2, '用户性别', 'sys_user_sex', '0', 'admin', '', '用户性别列表', '2021-11-30 14:12:33', '2021-11-30 14:12:33', '2021-11-30 14:14:19');
INSERT INTO `sys_dict_types` VALUES (3, '的心', 'sfd', '0', 'admin', '', 'fs', '2021-11-30 14:13:22', '2021-11-30 14:13:22', '2021-11-30 14:14:19');
INSERT INTO `sys_dict_types` VALUES (4, '用户性别', 'sys_user_sex', '0', 'admin', '', '性别字典', '2021-11-30 14:15:25', '2021-11-30 14:15:25', NULL);
INSERT INTO `sys_dict_types` VALUES (5, 'df', 'da', '0', 'admin', '', 'sd', '2021-11-30 15:54:33', '2021-11-30 15:54:33', '2021-11-30 15:54:40');
INSERT INTO `sys_dict_types` VALUES (6, '系统开关', 'sys_normal_disable', '0', 'admin', '', '开关列表', '2021-12-01 15:57:58', '2021-12-01 15:57:58', NULL);
INSERT INTO `sys_dict_types` VALUES (7, '菜单类型', 'sys_menu_type', '0', 'admin', '', '菜单类型列表', '2021-12-02 09:48:48', '2021-12-02 09:56:12', NULL);
INSERT INTO `sys_dict_types` VALUES (8, '菜单状态', 'sys_show_hide', '0', 'admin', '', '菜单状态列表', '2021-12-02 09:55:59', '2021-12-02 09:55:59', NULL);
INSERT INTO `sys_dict_types` VALUES (9, '数字是否', 'sys_num_yes_no', '0', 'admin', '', '数字是否列表', '2021-12-02 10:13:29', '2021-12-02 10:13:40', '2021-12-02 10:15:07');
INSERT INTO `sys_dict_types` VALUES (10, '数字是否', 'sys_num_yes_no', '0', 'admin', '', '数字是否', '2021-12-02 10:13:29', '2021-12-02 10:13:29', NULL);
INSERT INTO `sys_dict_types` VALUES (11, '状态是否', 'sys_yes_no', '0', 'admin', '', '状态是否', '2021-12-04 13:47:57', '2021-12-04 13:47:57', NULL);
INSERT INTO `sys_dict_types` VALUES (12, '网络请求方法', 'sys_method_api', '0', 'admin', '', '网络请求方法列表', '2021-12-08 17:21:27', '2021-12-08 17:21:27', NULL);

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `menu_id` bigint NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `parent_id` int NULL DEFAULT NULL,
  `sort` int NULL DEFAULT NULL,
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_frame` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_link` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `menu_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_hide` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_keep_alive` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_affix` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `permission` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (1, '系统设置', '', 0, 0, 'el-icon-s-tools', '/system', 'Layout', '1', '', 'M', '0', '0', '1', '', '0', 'admin', 'admin', '', '2021-12-02 11:04:08', '2021-12-04 11:04:45', NULL);
INSERT INTO `sys_menus` VALUES (3, '用户管理', '', 1, 1, 'el-icon-user', '/system/user', '/system/user/index', '1', '', 'C', '0', '1', '1', 'system:user:list', '0', 'admin', '', '', '2021-12-02 14:07:56', '2021-12-02 14:07:56', NULL);
INSERT INTO `sys_menus` VALUES (4, '添加用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:add', '0', 'admin', '', '', '2021-12-03 13:36:33', '2021-12-03 13:36:33', NULL);
INSERT INTO `sys_menus` VALUES (5, '编辑用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:edit', '0', 'admin', '', '', '2021-12-03 13:48:13', '2021-12-03 13:48:13', NULL);
INSERT INTO `sys_menus` VALUES (6, '角色管理', '', 1, 1, 'el-icon-user-solid', '/system/role', '/system/role/index', '1', '', 'C', '0', '1', '1', 'system:role:list', '0', 'admin', 'admin', '', '2021-12-03 13:51:55', '2021-12-03 13:55:04', NULL);
INSERT INTO `sys_menus` VALUES (7, '菜单管理', '', 1, 2, 'el-icon-menu', '/system/menu', '/system/menu/index', '1', '', 'C', '0', '1', '1', 'system:menu:list', '0', 'admin', 'admin', '', '2021-12-03 13:54:44', '2021-12-03 13:58:55', NULL);
INSERT INTO `sys_menus` VALUES (8, '部门管理', '', 1, 3, 'el-icon-set-up', '/system/dept', '/system/dept/index', '1', '', 'C', '0', '1', '1', 'system:dept:list', '0', 'admin', '', '', '2021-12-03 13:58:36', '2021-12-03 13:58:36', NULL);
INSERT INTO `sys_menus` VALUES (9, '岗位管理', '', 1, 4, 'el-icon-s-flag', '/system/post', '/system/post/index', '1', '', 'C', '0', '1', '1', 'system:post:list', '0', 'admin', '', '', '2021-12-03 13:54:44', '2021-12-03 13:54:44', NULL);
INSERT INTO `sys_menus` VALUES (10, '字典管理', '', 1, 5, 'el-icon-s-order', '/system/dict', '/system/dict/index', '1', '', 'C', '0', '1', '1', 'system:dict:list', '0', 'admin', '', '', '2021-12-03 13:54:44', '2021-12-03 13:54:44', NULL);
INSERT INTO `sys_menus` VALUES (11, '参数管理', '', 1, 6, 'el-icon-document', '/system/config', '/system/config/index', '1', '', 'C', '0', '1', '1', 'system:config:list', '0', 'admin', 'admin', '', '2021-12-03 13:54:44', '2021-12-03 17:06:35', NULL);
INSERT INTO `sys_menus` VALUES (12, '个人中心', '', 0, 2, 'el-icon-s-custom', '/personal', '/personal/index', '1', '', 'M', '0', '0', '0', '', '0', 'admin', 'admin', '', '2021-12-03 14:12:43', '2021-12-04 12:14:46', NULL);
INSERT INTO `sys_menus` VALUES (13, '添加配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:add', '0', 'admin', '', '', '2021-12-06 17:19:19', '2021-12-06 17:19:19', NULL);
INSERT INTO `sys_menus` VALUES (14, '修改配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:edit', '0', 'admin', '', '', '2021-12-06 17:20:30', '2021-12-06 17:20:30', NULL);
INSERT INTO `sys_menus` VALUES (15, '删除配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:delete', '0', 'admin', '', '', '2021-12-06 17:23:52', '2021-12-06 17:23:52', NULL);
INSERT INTO `sys_menus` VALUES (16, '导出配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:export', '0', 'admin', '', '', '2021-12-06 17:24:41', '2021-12-06 17:24:41', NULL);
INSERT INTO `sys_menus` VALUES (17, '新增角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:add', '0', 'admin', '', '', '2021-12-06 17:43:35', '2021-12-06 17:43:35', NULL);
INSERT INTO `sys_menus` VALUES (18, '删除角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:delete', '0', 'admin', '', '', '2021-12-06 17:44:10', '2021-12-06 17:44:10', NULL);
INSERT INTO `sys_menus` VALUES (19, '修改角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:edit', '0', 'admin', '', '', '2021-12-06 17:44:48', '2021-12-06 17:44:48', NULL);
INSERT INTO `sys_menus` VALUES (20, '导出角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:export', '0', 'admin', '', '', '2021-12-06 17:45:25', '2021-12-06 17:45:25', NULL);
INSERT INTO `sys_menus` VALUES (21, '添加菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:add', '0', 'admin', '', '', '2021-12-06 17:46:01', '2021-12-06 17:46:01', NULL);
INSERT INTO `sys_menus` VALUES (22, '修改菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:edit', '0', 'admin', '', '', '2021-12-06 17:46:24', '2021-12-06 17:46:24', NULL);
INSERT INTO `sys_menus` VALUES (23, '删除菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:delete', '0', 'admin', '', '', '2021-12-06 17:46:47', '2021-12-06 17:46:47', NULL);
INSERT INTO `sys_menus` VALUES (24, '添加部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:add', '0', 'admin', '', '', '2021-12-07 09:33:58', '2021-12-07 09:33:58', NULL);
INSERT INTO `sys_menus` VALUES (25, '编辑部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:edit', '0', 'admin', '', '', '2021-12-07 09:34:39', '2021-12-07 09:34:39', NULL);
INSERT INTO `sys_menus` VALUES (26, '删除部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:delete', '0', 'admin', 'admin', '', '2021-12-07 09:35:09', '2021-12-07 09:36:26', NULL);
INSERT INTO `sys_menus` VALUES (27, '导出部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:export', '0', 'admin', '', '', '2021-12-07 09:35:51', '2021-12-07 09:35:51', '2021-12-07 09:36:37');
INSERT INTO `sys_menus` VALUES (28, '添加岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:add', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (29, '编辑岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:edit', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (30, '删除岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:delete', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (31, '导出岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:export', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (32, '添加字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:add', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (33, '编辑字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:edit', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (34, '删除字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:delete', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (35, '导出字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:export', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (36, '新增字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:add', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO `sys_menus` VALUES (37, '修改字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:edit', '0', 'admin', '', '', '2021-12-07 09:48:04', '2021-12-07 09:48:04', NULL);
INSERT INTO `sys_menus` VALUES (38, '删除字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:delete', '0', 'admin', '', '', '2021-12-07 09:48:42', '2021-12-07 09:48:42', NULL);
INSERT INTO `sys_menus` VALUES (39, 'API管理', '', 1, 1, 'iconfont icon-siweidaotu', '/system/api', '/system/api/index', '1', '', 'C', '0', '1', '1', 'system:api:list', '0', 'admin', '', '', '2021-12-09 09:09:13', '2021-12-09 09:09:13', NULL);
INSERT INTO `sys_menus` VALUES (40, '添加api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:add', '0', 'admin', '', '', '2021-12-09 09:09:54', '2021-12-09 09:09:54', NULL);
INSERT INTO `sys_menus` VALUES (41, '编辑api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:edit', '0', 'admin', '', '', '2021-12-09 09:10:38', '2021-12-09 09:10:38', NULL);
INSERT INTO `sys_menus` VALUES (42, '删除api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:delete', '0', 'admin', '', '', '2021-12-09 09:11:11', '2021-12-09 09:11:11', NULL);

-- ----------------------------
-- Table structure for sys_posts
-- ----------------------------
DROP TABLE IF EXISTS `sys_posts`;
CREATE TABLE `sys_posts`  (
  `post_id` bigint NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位名称',
  `post_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位代码',
  `sort` int NULL DEFAULT NULL COMMENT '岗位排序',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_posts
-- ----------------------------
INSERT INTO `sys_posts` VALUES (1, '首席执行官', 'CEO', 0, '0', '首席执行官', 'admin', '', '2021-12-02 09:21:44', '2021-12-02 09:24:25', NULL);
INSERT INTO `sys_posts` VALUES (3, '首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44', '2021-12-02 09:25:59', '2021-12-02 09:27:41');
INSERT INTO `sys_posts` VALUES (4, '首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44', '2021-12-02 09:25:59', NULL);

-- ----------------------------
-- Table structure for sys_role_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_depts`;
CREATE TABLE `sys_role_depts`  (
  `role_id` int NULL DEFAULT NULL,
  `dept_id` int NULL DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_depts
-- ----------------------------
INSERT INTO `sys_role_depts` VALUES (1, 2, 1);
INSERT INTO `sys_role_depts` VALUES (1, 3, 2);

-- ----------------------------
-- Table structure for sys_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menus`;
CREATE TABLE `sys_role_menus`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `role_id` int NULL DEFAULT NULL,
  `menu_id` int NULL DEFAULT NULL,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 349 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO `sys_role_menus` VALUES (309, 1, 1, 'admin');
INSERT INTO `sys_role_menus` VALUES (310, 1, 3, 'admin');
INSERT INTO `sys_role_menus` VALUES (311, 1, 4, 'admin');
INSERT INTO `sys_role_menus` VALUES (312, 1, 5, 'admin');
INSERT INTO `sys_role_menus` VALUES (313, 1, 6, 'admin');
INSERT INTO `sys_role_menus` VALUES (314, 1, 7, 'admin');
INSERT INTO `sys_role_menus` VALUES (315, 1, 8, 'admin');
INSERT INTO `sys_role_menus` VALUES (316, 1, 9, 'admin');
INSERT INTO `sys_role_menus` VALUES (317, 1, 10, 'admin');
INSERT INTO `sys_role_menus` VALUES (318, 1, 11, 'admin');
INSERT INTO `sys_role_menus` VALUES (319, 1, 12, 'admin');
INSERT INTO `sys_role_menus` VALUES (320, 1, 13, 'admin');
INSERT INTO `sys_role_menus` VALUES (321, 1, 14, 'admin');
INSERT INTO `sys_role_menus` VALUES (322, 1, 15, 'admin');
INSERT INTO `sys_role_menus` VALUES (323, 1, 16, 'admin');
INSERT INTO `sys_role_menus` VALUES (324, 1, 17, 'admin');
INSERT INTO `sys_role_menus` VALUES (325, 1, 18, 'admin');
INSERT INTO `sys_role_menus` VALUES (326, 1, 19, 'admin');
INSERT INTO `sys_role_menus` VALUES (327, 1, 20, 'admin');
INSERT INTO `sys_role_menus` VALUES (328, 1, 21, 'admin');
INSERT INTO `sys_role_menus` VALUES (329, 1, 22, 'admin');
INSERT INTO `sys_role_menus` VALUES (330, 1, 23, 'admin');
INSERT INTO `sys_role_menus` VALUES (331, 1, 24, 'admin');
INSERT INTO `sys_role_menus` VALUES (332, 1, 25, 'admin');
INSERT INTO `sys_role_menus` VALUES (333, 1, 26, 'admin');
INSERT INTO `sys_role_menus` VALUES (334, 1, 28, 'admin');
INSERT INTO `sys_role_menus` VALUES (335, 1, 29, 'admin');
INSERT INTO `sys_role_menus` VALUES (336, 1, 30, 'admin');
INSERT INTO `sys_role_menus` VALUES (337, 1, 31, 'admin');
INSERT INTO `sys_role_menus` VALUES (338, 1, 32, 'admin');
INSERT INTO `sys_role_menus` VALUES (339, 1, 33, 'admin');
INSERT INTO `sys_role_menus` VALUES (340, 1, 34, 'admin');
INSERT INTO `sys_role_menus` VALUES (341, 1, 35, 'admin');
INSERT INTO `sys_role_menus` VALUES (342, 1, 36, 'admin');
INSERT INTO `sys_role_menus` VALUES (343, 1, 37, 'admin');
INSERT INTO `sys_role_menus` VALUES (344, 1, 38, 'admin');
INSERT INTO `sys_role_menus` VALUES (345, 1, 39, 'admin');
INSERT INTO `sys_role_menus` VALUES (346, 1, 40, 'admin');
INSERT INTO `sys_role_menus` VALUES (347, 1, 41, 'admin');
INSERT INTO `sys_role_menus` VALUES (348, 1, 42, 'admin');

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `role_id` bigint NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色代码',
  `data_scope` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `role_sort` int NULL DEFAULT NULL COMMENT '角色排序',
  `flag` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除标识',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO `sys_roles` VALUES (1, '超管理员', '0', 'admin', '2', 1, '', 'admin', 'admin', '', '2021-12-02 16:03:26', '2021-12-09 19:53:20', NULL);

-- ----------------------------
-- Records of sys_settings
-- ----------------------------

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role_id` int NULL DEFAULT NULL,
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `dept_id` int NULL DEFAULT NULL,
  `post_id` int NULL DEFAULT NULL,
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '多角色',
  `post_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '多岗位',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, 'pandax', '13818888888', 1, NULL, '', '0', '1@qq.com', 2, 1, 'admin', '1', NULL, '0', '2021-12-03 09:46:55', '2021-12-03 10:51:34', NULL, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '1', '1');
INSERT INTO `sys_users` VALUES (2, '测试用户', '18353366754', 1, '', '', '0', '1234@163.com', 3, 4, 'admin', '', 'test', '0', '2021-12-03 10:11:11', '2021-12-03 11:23:59', '2021-12-03 11:24:38', 'test', '$2a$10$Xj35fNwCNbqcXWByNbd8/e0CrN0vqoDToqFxpOQDzhk4pN59Cm9Gu', NULL, '');
INSERT INTO `sys_users` VALUES (3, '测试用户', '18435234356', 1, '', '', '0', '342@163.com', 3, 1, 'admin', '', '', '0', '2021-12-06 15:16:53', '2021-12-06 15:29:28', NULL, 'test', '$2a$10$4cHTracxWJLdhMmazvbm1urKyD3v5N2AYxAFtNYh6juU39kgae73e', '1', '1,4');

SET FOREIGN_KEY_CHECKS = 1;
