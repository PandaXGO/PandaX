/*
 Navicat Premium Data Transfer

 Source Server         : 阿里云
 Source Server Type    : MySQL
 Source Server Version : 50562
 Source Host           : 47.104.252.2:3306
 Source Schema         : pandax

 Target Server Type    : MySQL
 Target Server Version : 50562
 File Encoding         : 65001

 Date: 01/01/2022 11:19:40
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
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 125 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job', 'POST', '', '', '', 34);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job', 'PUT', '', '', '', 35);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/:jobId', 'DELETE', '', '', '', 37);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/:jobId', 'GET', '', '', '', 36);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/changeStatus', 'PUT', '', '', '', 40);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/list', 'GET', '', '', '', 33);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/start/:jobId', 'GET', '', '', '', 39);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/stop/:jobId', 'GET', '', '', '', 38);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logJob/:logId', 'DELETE', '', '', '', 49);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logJob/all', 'DELETE', '', '', '', 48);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logJob/list', 'GET', '', '', '', 47);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logLogin/:infoId', 'DELETE', '', '', '', 42);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logLogin/all', 'DELETE', '', '', '', 43);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logLogin/list', 'GET', '', '', '', 41);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logOper/:operId', 'DELETE', '', '', '', 45);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logOper/all', 'DELETE', '', '', '', 46);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logOper/list', 'GET', '', '', '', 44);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api', 'POST', '', '', '', 5);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api', 'PUT', '', '', '', 6);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/:id', 'DELETE', '', '', '', 7);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/:id', 'GET', '', '', '', 4);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/all', 'GET', '', '', '', 2);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/getPolicyPathByRoleId', 'GET', '', '', '', 3);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/list', 'GET', '', '', '', 1);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'POST', '', '', '', 11);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'PUT', '', '', '', 12);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/:configId', 'DELETE', '', '', '', 13);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/:configId', 'GET', '', '', '', 10);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/configKey', 'GET', '', '', '', 9);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/list', 'GET', '', '', '', 8);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept', 'POST', '', '', '', 18);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept', 'PUT', '', '', '', 19);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/:deptId', 'DELETE', '', '', '', 20);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/:deptId', 'GET', '', '', '', 15);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/deptTree', 'GET', '', '', '', 17);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/list', 'GET', '', '', '', 14);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', '', 16);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'POST', '', '', '', 30);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'PUT', '', '', '', 31);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/:dictCode', 'DELETE', '', '', '', 32);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/:dictCode', 'GET', '', '', '', 29);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/list', 'GET', '', '', '', 27);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/type', 'GET', '', '', '', 28);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'POST', '', '', '', 23);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'PUT', '', '', '', 24);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/:dictId', 'DELETE', '', '', '', 25);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/:dictId', 'GET', '', '', '', 22);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/export', 'GET', '', '', '', 26);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/list', 'GET', '', '', '', 21);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'POST', '', '', '', 56);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'PUT', '', '', '', 57);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/:menuId', 'DELETE', '', '', '', 58);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/:menuId', 'GET', '', '', '', 55);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/list', 'GET', '', '', '', 54);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuPaths', 'GET', '', '', '', 53);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuRole', 'GET', '', '', '', 51);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuTreeSelect', 'GET', '', '', '', 50);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', '', 52);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/notice', 'POST', '', '', '', 60);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/notice', 'PUT', '', '', '', 61);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/notice/:noticeId', 'DELETE', '', '', '', 62);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/notice/list', 'GET', '', '', '', 59);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'POST', '', '', '', 65);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'PUT', '', '', '', 66);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/:postId', 'DELETE', '', '', '', 67);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/:postId', 'GET', '', '', '', 64);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/list', 'GET', '', '', '', 63);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'POST', '', '', '', 70);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'PUT', '', '', '', 71);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/:roleId', 'DELETE', '', '', '', 72);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/:roleId', 'GET', '', '', '', 69);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/changeStatus', 'PUT', '', '', '', 73);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/dataScope', 'PUT', '', '', '', 74);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/export', 'GET', '', '', '', 75);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/list', 'GET', '', '', '', 68);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'POST', '', '', '', 84);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'PUT', '', '', '', 85);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/:userId', 'DELETE', '', '', '', 78);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/avatar', 'POST', '', '', '', 79);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/changeStatus', 'PUT', '', '', '', 77);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/export', 'GET', '', '', '', 86);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getById/:userId', 'GET', '', '', '', 81);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getInit', 'GET', '', '', '', 82);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getRoPo', 'GET', '', '', '', 83);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/list', 'GET', '', '', '', 76);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/pwd', 'PUT', '', '', '', 80);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/job/:jobId', 'GET', '', '', '', 104);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/job/changeStatus', 'PUT', '', '', '', 105);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/job/list', 'GET', '', '', '', 103);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/log/logJob/list', 'GET', '', '', '', 108);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/log/logLogin/list', 'GET', '', '', '', 106);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/log/logOper/list', 'GET', '', '', '', 107);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/api/:id', 'GET', '', '', '', 90);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/api/all', 'GET', '', '', '', 88);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/api/getPolicyPathByRoleId', 'GET', '', '', '', 89);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/api/list', 'GET', '', '', '', 87);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/config/:configId', 'GET', '', '', '', 93);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/config/configKey', 'GET', '', '', '', 92);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/config/list', 'GET', '', '', '', 91);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dept/:deptId', 'GET', '', '', '', 95);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dept/deptTree', 'GET', '', '', '', 97);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dept/list', 'GET', '', '', '', 94);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', '', 96);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/data/:dictCode', 'GET', '', '', '', 102);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/data/list', 'GET', '', '', '', 100);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/data/type', 'GET', '', '', '', 101);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/type/:dictId', 'GET', '', '', '', 99);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/type/list', 'GET', '', '', '', 98);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/:menuId', 'GET', '', '', '', 114);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/list', 'GET', '', '', '', 113);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/menuPaths', 'GET', '', '', '', 112);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/menuRole', 'GET', '', '', '', 110);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/menuTreeSelect', 'GET', '', '', '', 109);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', '', 111);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/notice/list', 'GET', '', '', '', 115);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/post/:postId', 'GET', '', '', '', 117);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/post/list', 'GET', '', '', '', 116);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/role/:roleId', 'GET', '', '', '', 119);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/role/list', 'GET', '', '', '', 118);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/avatar', 'POST', '', '', '', 121);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/getById/:userId', 'GET', '', '', '', 122);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/getInit', 'GET', '', '', '', 123);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/getRoPo', 'GET', '', '', '', 124);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/list', 'GET', '', '', '', 120);

-- ----------------------------
-- Table structure for log_jobs
-- ----------------------------
DROP TABLE IF EXISTS `log_jobs`;
CREATE TABLE `log_jobs`  (
  `log_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '任务名称',
  `job_group` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '分组',
  `entry_id` int(11) NULL DEFAULT NULL COMMENT '任务id',
  `invoke_target` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '调用方法',
  `log_info` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '日志信息',
  `status` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 260 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;


-- ----------------------------
-- Table structure for log_logins
-- ----------------------------
DROP TABLE IF EXISTS `log_logins`;
CREATE TABLE `log_logins`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT,
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
) ENGINE = InnoDB AUTO_INCREMENT = 383 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for log_opers
-- ----------------------------
DROP TABLE IF EXISTS `log_opers`;
CREATE TABLE `log_opers`  (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作的模块',
  `business_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '0其它 1新增 2修改 3删除',
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
) ENGINE = InnoDB AUTO_INCREMENT = 173 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 87 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `sys_apis` VALUES (66, '2021-12-17 10:51:05', '2021-12-17 10:54:22', NULL, '/log/logLogin/list', '获取登录日志', 'log', 'GET');
INSERT INTO `sys_apis` VALUES (67, '2021-12-17 10:51:43', '2021-12-17 10:54:28', NULL, '/log/logLogin/:infoId', '删除日志', 'log', 'DELETE');
INSERT INTO `sys_apis` VALUES (68, '2021-12-17 10:53:09', '2021-12-17 10:54:34', NULL, '/log/logLogin/all', '清空所有', 'log', 'DELETE');
INSERT INTO `sys_apis` VALUES (69, '2021-12-17 10:54:07', '2021-12-17 10:54:07', NULL, '/log/logOper/list', '操作日志列表', 'log', 'GET');
INSERT INTO `sys_apis` VALUES (70, '2021-12-17 10:53:09', '2021-12-17 10:53:09', NULL, '/log/logOper/:operId', '删除', 'log', 'DELETE');
INSERT INTO `sys_apis` VALUES (71, '2021-12-17 10:53:09', '2021-12-17 10:53:09', NULL, '/log/logOper/all', '清空', 'log', 'DELETE');
INSERT INTO `sys_apis` VALUES (72, '2021-12-24 15:41:23', '2021-12-24 15:41:23', NULL, '/job/list', '任务列表', 'job', 'GET');
INSERT INTO `sys_apis` VALUES (73, '2021-12-24 15:41:54', '2021-12-24 15:41:54', NULL, '/job', '添加', 'job', 'POST');
INSERT INTO `sys_apis` VALUES (74, '2021-12-24 15:42:11', '2021-12-24 15:42:11', NULL, '/job', '修改任务', 'job', 'PUT');
INSERT INTO `sys_apis` VALUES (75, '2021-12-24 15:42:37', '2021-12-24 16:32:21', NULL, '/job/:jobId', '获取任务', 'job', 'GET');
INSERT INTO `sys_apis` VALUES (76, '2021-12-24 15:43:09', '2021-12-24 16:32:05', NULL, '/job/:jobId', '删除job', 'job', 'DELETE');
INSERT INTO `sys_apis` VALUES (77, '2021-12-24 15:43:35', '2021-12-24 16:31:11', NULL, '/job/stop/:jobId', '停止任务', 'job', 'GET');
INSERT INTO `sys_apis` VALUES (78, '2021-12-24 15:44:09', '2021-12-24 16:30:38', NULL, '/job/start/:jobId', '开始任务', 'job', 'GET');
INSERT INTO `sys_apis` VALUES (79, '2021-12-24 15:45:03', '2021-12-24 15:46:36', NULL, '/log/logJob/list', '任务日志列表', 'log', 'GET');
INSERT INTO `sys_apis` VALUES (80, '2021-12-24 15:45:33', '2021-12-24 15:46:43', NULL, '/log/logJob/all', '清空任务日志', 'log', 'DELETE');
INSERT INTO `sys_apis` VALUES (81, '2021-12-24 15:46:08', '2021-12-24 16:33:13', NULL, '/log/logJob/:logId', '删除任务日志', 'log', 'DELETE');
INSERT INTO `sys_apis` VALUES (82, '2021-12-24 15:45:33', '2021-12-24 15:45:33', NULL, '/system/notice/list', '获取通知分页列表', 'notice', 'GET');
INSERT INTO `sys_apis` VALUES (83, '2021-12-24 15:45:33', '2021-12-24 15:45:33', NULL, '/system/notice', '添加通知信息', 'notice', 'POST');
INSERT INTO `sys_apis` VALUES (84, '2021-12-24 15:45:33', '2021-12-24 15:45:33', NULL, '/system/notice', '修改通知信息', 'notice', 'PUT');
INSERT INTO `sys_apis` VALUES (85, '2021-12-24 15:45:33', '2021-12-24 16:33:48', NULL, '/system/notice/:noticeId', '删除通知信息', 'notice', 'DELETE');
INSERT INTO `sys_apis` VALUES (86, '2021-12-24 22:40:19', '2021-12-24 22:40:19', NULL, '/job/changeStatus', '修改状态', 'job', 'PUT');

-- ----------------------------
-- Table structure for sys_configs
-- ----------------------------
DROP TABLE IF EXISTS `sys_configs`;
CREATE TABLE `sys_configs`  (
  `config_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_configs
-- ----------------------------
INSERT INTO `sys_configs` VALUES (1, '账号初始密码', 'sys.user.initPassword', '123456', '0', '0', '初始密码', '2021-12-04 13:50:13', '2021-12-04 13:54:52', NULL);

-- ----------------------------
-- Table structure for sys_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_depts`;
CREATE TABLE `sys_depts`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NULL DEFAULT NULL COMMENT '上级部门',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门路径',
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门名称',
  `sort` int(4) NULL DEFAULT NULL COMMENT '排序',
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
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_depts
-- ----------------------------
INSERT INTO `sys_depts` VALUES (2, 0, '/0/2', '熊猫科技', 0, 'xm', '18353366836', '342@qq.com', '0', 'admin', 'admin', '2021-12-01 17:31:53', '2021-12-02 08:56:19', NULL);
INSERT INTO `sys_depts` VALUES (3, 2, '/0/2/3', '研发部', 1, 'panda', '18353366543', 'ewr@qq.com', '0', 'admin', 'admin', '2021-12-01 17:37:43', '2021-12-02 08:55:56', NULL);
INSERT INTO `sys_depts` VALUES (7, 2, '/0/2/7', '营销部', 2, 'panda', '18353333333', '342@qq.com', '0', 'panda', 'panda', '2021-12-24 10:46:24', '2021-12-24 10:47:15', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT,
  `dict_sort` int(11) NULL DEFAULT NULL COMMENT '排序',
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
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `sys_dict_data` VALUES (19, 0, '成功', '0', 'sys_common_status', '0', '', '', '', 'admin', '', '', '2021-12-17 11:01:52', '2021-12-17 11:01:52', NULL);
INSERT INTO `sys_dict_data` VALUES (20, 0, '失败', '1', 'sys_common_status', '0', '', '', '', 'admin', '', '', '2021-12-17 11:02:08', '2021-12-17 11:02:08', NULL);
INSERT INTO `sys_dict_data` VALUES (21, 0, '其他', '0', 'sys_oper_type', '0', '', '', '', 'admin', '', '', '2021-12-17 11:30:07', '2021-12-17 11:30:07', NULL);
INSERT INTO `sys_dict_data` VALUES (22, 0, '新增', '1', 'sys_oper_type', '0', '', '', '', 'admin', '', '', '2021-12-17 11:30:21', '2021-12-17 11:30:21', NULL);
INSERT INTO `sys_dict_data` VALUES (23, 0, '修改', '2', 'sys_oper_type', '0', '', '', '', 'admin', '', '', '2021-12-17 11:30:32', '2021-12-17 11:30:32', NULL);
INSERT INTO `sys_dict_data` VALUES (24, 0, '删除', '3', 'sys_oper_type', '0', '', '', '', 'admin', '', '', '2021-12-17 11:30:40', '2021-12-17 11:30:40', NULL);
INSERT INTO `sys_dict_data` VALUES (25, 0, '默认', 'DEFAULT', 'sys_job_group', '0', '', '', '', 'panda', '', '', '2021-12-24 15:15:31', '2021-12-24 15:15:31', NULL);
INSERT INTO `sys_dict_data` VALUES (26, 1, '系统', 'SYSTEM', 'sys_job_group', '0', '', '', '', 'panda', '', '', '2021-12-24 15:15:50', '2021-12-24 15:15:50', NULL);
INSERT INTO `sys_dict_data` VALUES (27, 0, '发布通知', '1', 'sys_notice_type', '0', '', '', '', 'panda', '', '', '2021-12-26 15:24:07', '2021-12-26 15:24:07', NULL);
INSERT INTO `sys_dict_data` VALUES (28, 0, '任免通知', '2', 'sys_notice_type', '0', '', '', '', 'panda', '', '', '2021-12-26 15:24:18', '2021-12-26 15:24:18', NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, '事物通知', '3', 'sys_notice_type', '0', '', '', '', 'panda', '', '', '2021-12-26 15:24:46', '2021-12-26 15:24:46', NULL);
INSERT INTO `sys_dict_data` VALUES (30, 0, '审批通知', '4', 'sys_notice_type', '0', '', '', '', 'panda', '', '', '2021-12-26 15:25:08', '2021-12-26 15:25:08', NULL);

-- ----------------------------
-- Table structure for sys_dict_types
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_types`;
CREATE TABLE `sys_dict_types`  (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT,
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
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `sys_dict_types` VALUES (13, '成功失败', 'sys_common_status', '0', 'admin', '', '是否成功失败', '2021-12-17 10:10:03', '2021-12-17 10:10:03', NULL);
INSERT INTO `sys_dict_types` VALUES (27, '操作分类', 'sys_oper_type', '0', 'admin', '', '操作分类列表', '2021-12-17 11:29:50', '2021-12-17 11:29:50', NULL);
INSERT INTO `sys_dict_types` VALUES (28, '任务组', 'sys_job_group', '0', 'panda', '', '系统任务，开机自启', '2021-12-24 15:14:56', '2021-12-24 15:14:56', NULL);
INSERT INTO `sys_dict_types` VALUES (29, '通知类型', 'sys_notice_type', '0', 'panda', '', '通知类型列表', '2021-12-26 15:23:26', '2021-12-26 15:23:26', NULL);

-- ----------------------------
-- Table structure for sys_jobs
-- ----------------------------
DROP TABLE IF EXISTS `sys_jobs`;
CREATE TABLE `sys_jobs`  (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `job_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `job_group` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `job_type` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `invoke_target` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `args` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `misfire_policy` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `concurrent` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `entry_id` int(11) NULL DEFAULT NULL,
  `create_by` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`job_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_jobs
-- ----------------------------
INSERT INTO `sys_jobs` VALUES (1, 'testcron', 'SYSTEM', '2', ' 0/10 * * * * ?', 'cronHandle', 'aaa', '', '', '1', 0, 'panda', '', '2021-12-24 16:06:09', '2021-12-24 22:44:09', NULL);

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `parent_id` int(11) NULL DEFAULT NULL,
  `sort` int(4) NULL DEFAULT NULL,
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_iframe` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
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
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (1, '系统设置', '', 0, 0, 'elementSetting', '/system', 'Layout', '1', '', 'M', '0', '0', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08', '2021-12-28 13:32:21', NULL);
INSERT INTO `sys_menus` VALUES (3, '用户管理', '', 1, 1, 'elementUser', '/system/user', '/system/user/index', '1', '', 'C', '0', '1', '1', 'system:user:list', '0', 'admin', 'panda', '', '2021-12-02 14:07:56', '2021-12-28 13:32:44', NULL);
INSERT INTO `sys_menus` VALUES (4, '添加用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:add', '0', 'admin', '', '', '2021-12-03 13:36:33', '2021-12-03 13:36:33', NULL);
INSERT INTO `sys_menus` VALUES (5, '编辑用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:edit', '0', 'admin', '', '', '2021-12-03 13:48:13', '2021-12-03 13:48:13', NULL);
INSERT INTO `sys_menus` VALUES (6, '角色管理', '', 1, 1, 'elementUserFilled', '/system/role', '/system/role/index', '1', '', 'C', '0', '1', '1', 'system:role:list', '0', 'admin', 'panda', '', '2021-12-03 13:51:55', '2021-12-28 13:32:55', NULL);
INSERT INTO `sys_menus` VALUES (7, '菜单管理', '', 1, 2, 'iconfont icon-juxingkaobei', '/system/menu', '/system/menu/index', '1', '', 'C', '0', '1', '1', 'system:menu:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:33:19', NULL);
INSERT INTO `sys_menus` VALUES (8, '部门管理', '', 1, 3, 'iconfont icon-jiliandongxuanzeqi', '/system/dept', '/system/dept/index', '1', '', 'C', '0', '1', '1', 'system:dept:list', '0', 'admin', 'panda', '', '2021-12-03 13:58:36', '2021-12-28 13:40:20', NULL);
INSERT INTO `sys_menus` VALUES (9, '岗位管理', '', 1, 4, 'iconfont icon-neiqianshujuchucun', '/system/post', '/system/post/index', '1', '', 'C', '0', '1', '1', 'system:post:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:40:31', NULL);
INSERT INTO `sys_menus` VALUES (10, '字典管理', '', 1, 5, 'elementCellphone', '/system/dict', '/system/dict/index', '1', '', 'C', '0', '1', '1', 'system:dict:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:40:50', NULL);
INSERT INTO `sys_menus` VALUES (11, '参数管理', '', 1, 6, 'elementDocumentCopy', '/system/config', '/system/config/index', '1', '', 'C', '0', '1', '1', 'system:config:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:41:05', NULL);
INSERT INTO `sys_menus` VALUES (12, '个人中心', '', 0, 10, 'elementAvatar', '/personal', '/personal/index', '1', '', 'M', '0', '0', '0', '', '0', 'admin', 'panda', '', '2021-12-03 14:12:43', '2021-12-28 13:43:17', NULL);
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
INSERT INTO `sys_menus` VALUES (43, '日志系统', '', 0, 1, 'iconfont icon-biaodan', '/log', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08', '2021-12-28 13:38:33', NULL);
INSERT INTO `sys_menus` VALUES (44, '系统工具', '', 0, 2, 'iconfont icon-gongju', '/tool', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-16 16:35:15', '2021-12-28 13:38:46', NULL);
INSERT INTO `sys_menus` VALUES (45, '操作日志', '', 43, 1, 'iconfont icon-bolangnengshiyanchang', '/log/operation', '/log/operation/index', '1', '', 'C', '0', '1', '1', 'log:operation:list', '0', 'admin', 'panda', '', '2021-12-16 16:42:03', '2021-12-28 13:39:44', NULL);
INSERT INTO `sys_menus` VALUES (46, '登录日志', '', 43, 2, 'iconfont icon--chaifenlie', '/log/login', '/log/login/index', '1', '', 'C', '0', '1', '1', 'log:login:list', '0', 'admin', 'panda', '', '2021-12-16 16:43:28', '2021-12-28 13:39:58', NULL);
INSERT INTO `sys_menus` VALUES (47, '服务监控', '', 44, 1, 'elementCpu', '/tool/monitor/', '/tool/monitor/index', '1', '', 'C', '0', '1', '1', 'tool:monitor:list', '0', 'admin', 'panda', '', '2021-12-03 14:12:43', '2021-12-28 13:41:25', NULL);
INSERT INTO `sys_menus` VALUES (48, '定时任务', '', 44, 2, 'elementAlarmClock', '/tool/job', '/tool/job/index', '1', '', 'C', '0', '1', '1', 'tool:job:list', '0', 'admin', 'panda', '', '2021-12-16 16:48:45', '2021-12-28 13:41:59', NULL);
INSERT INTO `sys_menus` VALUES (49, '开发工具', '', 0, 3, 'iconfont icon-diannao', '/develop', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', '', '', '2021-12-16 16:53:11', '2021-12-16 16:53:11', NULL);
INSERT INTO `sys_menus` VALUES (50, '表单构建', '', 49, 1, 'iconfont icon-zidingyibuju', '/develop/form', '/layout/routerView/iframes', '0', 'http://127.0.0.1:7788/form-generator/', 'C', '0', '1', '1', 'develop:form:list', '0', 'admin', 'panda', '', '2021-12-16 16:55:01', '2021-12-31 08:23:07', NULL);
INSERT INTO `sys_menus` VALUES (51, '代码生成', '', 49, 2, 'iconfont icon-zhongduancanshu', '/develop/code', '/develop/code/index', '1', '', 'C', '0', '1', '1', 'develop:code:list', '0', 'admin', '', '', '2021-12-16 16:56:48', '2021-12-16 16:56:48', NULL);
INSERT INTO `sys_menus` VALUES (52, '系统接口', '', 49, 3, 'iconfont icon-wenducanshu-05', '/develop/apis', '/layout/routerView/iframes', '0', 'http://127.0.0.1:7788/swagger/index.html', 'C', '0', '1', '1', 'develop:apis:list', '0', 'admin', 'panda', '', '2021-12-16 16:58:07', '2021-12-25 00:09:23', NULL);
INSERT INTO `sys_menus` VALUES (53, '资源管理', '', 0, 4, 'iconfont icon-juxingkaobei', '/resource', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', '', '', '2021-12-16 17:02:06', '2021-12-16 17:02:06', NULL);
INSERT INTO `sys_menus` VALUES (54, '文件管理', '', 53, 1, 'iconfont icon-chazhaobiaodanliebiao', '/resource/file', '/resource/file/index', '1', '', 'C', '0', '1', '1', 'resource:file:list', '0', 'admin', 'panda', '', '2021-12-16 17:06:04', '2021-12-28 15:13:08', NULL);
INSERT INTO `sys_menus` VALUES (55, '公告通知', '', 44, 3, 'elementTicket', '/tool/notice', '/tool/notice/index', '1', '', 'C', '0', '1', '1', 'tool:notice:list', '0', 'admin', 'panda', '', '2021-12-16 22:09:11', '2021-12-28 13:42:39', NULL);
INSERT INTO `sys_menus` VALUES (56, '任务日志', '', 43, 1, 'iconfont icon--chaifenhang', '/log/job', '/log/job/index', '1', '', 'C', '0', '1', '1', 'log:job:list', '0', 'panda', 'panda', '', '2021-12-24 22:13:45', '2021-12-28 13:39:52', NULL);
INSERT INTO `sys_menus` VALUES (57, '邮件管理', '', 53, 1, 'iconfont icon-shouye_dongtaihui', '/resource/mail', '/resource/mail/index', '1', '', 'C', '0', '1', '1', 'resource:mail:list', '0', 'panda', 'panda', '', '2021-12-28 15:12:37', '2021-12-28 15:12:45', NULL);
INSERT INTO `sys_menus` VALUES (58, '短信管理', '', 53, 3, 'elementChatDotRound', '/resource/message', '/resource/message/index', '1', '', 'C', '0', '1', '1', 'resource:message:list', '0', 'panda', '', '', '2021-12-16 17:06:04', '2021-12-16 17:06:04', NULL);

-- ----------------------------
-- Table structure for sys_notices
-- ----------------------------
DROP TABLE IF EXISTS `sys_notices`;
CREATE TABLE `sys_notices`  (
  `notice_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '标题',
  `notice_type` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '通知类型',
  `dept_id` int(11) NULL DEFAULT NULL COMMENT '部门Id,部门及子部门',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  `user_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_notices
-- ----------------------------
INSERT INTO `sys_notices` VALUES (1, '关于学习交流的通知', '<p>发布<b>入群</b>通知&nbsp;<span style=\"color: var(--el-text-color-regular);\">467890197, 交流学习</span><span style=\"font-size: 14px; color: var(--el-text-color-regular);\"></span></p>', '1', 0, '2021-12-26 15:29:25', '2021-12-26 16:19:48', NULL, 'panda');
INSERT INTO `sys_notices` VALUES (2, 'test', '<p>sdsad</p>', '1', 2, '2021-12-26 16:23:13', '2021-12-26 16:23:13', '2021-12-26 16:31:31', 'panda');
INSERT INTO `sys_notices` VALUES (3, '版本更新通知：任务功能，通知功能完成', '<p><span style=\"font-size: 14px; color: var(--el-text-color-regular);\">', '1', 0, '2021-12-26 17:33:47', '2021-12-26 17:33:47', NULL, 'panda');

-- ----------------------------
-- Table structure for sys_posts
-- ----------------------------
DROP TABLE IF EXISTS `sys_posts`;
CREATE TABLE `sys_posts`  (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位名称',
  `post_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位代码',
  `sort` int(4) NULL DEFAULT NULL COMMENT '岗位排序',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_posts
-- ----------------------------
INSERT INTO `sys_posts` VALUES (1, '首席执行官', 'CEO', 0, '0', '首席执行官', 'admin', '', '2021-12-02 09:21:44', '2021-12-02 09:24:25', NULL);
INSERT INTO `sys_posts` VALUES (3, '首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44', '2021-12-02 09:25:59', '2021-12-02 09:27:41');
INSERT INTO `sys_posts` VALUES (4, '首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44', '2021-12-02 09:25:59', NULL);
INSERT INTO `sys_posts` VALUES (5, '123', '123', 0, '0', '', 'admin', '', '2021-12-18 00:33:28', '2021-12-18 00:33:28', '2021-12-28 14:11:52');

-- ----------------------------
-- Table structure for sys_role_depts
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_depts`;
CREATE TABLE `sys_role_depts`  (
  `role_id` int(11) NULL DEFAULT NULL,
  `dept_id` int(11) NULL DEFAULT NULL,
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

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
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NULL DEFAULT NULL,
  `menu_id` int(11) NULL DEFAULT NULL,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 937 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO `sys_role_menus` VALUES (826, 1, 1, 'admin');
INSERT INTO `sys_role_menus` VALUES (827, 1, 3, 'admin');
INSERT INTO `sys_role_menus` VALUES (828, 1, 4, 'admin');
INSERT INTO `sys_role_menus` VALUES (829, 1, 5, 'admin');
INSERT INTO `sys_role_menus` VALUES (830, 1, 6, 'admin');
INSERT INTO `sys_role_menus` VALUES (831, 1, 7, 'admin');
INSERT INTO `sys_role_menus` VALUES (832, 1, 8, 'admin');
INSERT INTO `sys_role_menus` VALUES (833, 1, 9, 'admin');
INSERT INTO `sys_role_menus` VALUES (834, 1, 10, 'admin');
INSERT INTO `sys_role_menus` VALUES (835, 1, 11, 'admin');
INSERT INTO `sys_role_menus` VALUES (836, 1, 12, 'admin');
INSERT INTO `sys_role_menus` VALUES (837, 1, 13, 'admin');
INSERT INTO `sys_role_menus` VALUES (838, 1, 14, 'admin');
INSERT INTO `sys_role_menus` VALUES (839, 1, 15, 'admin');
INSERT INTO `sys_role_menus` VALUES (840, 1, 16, 'admin');
INSERT INTO `sys_role_menus` VALUES (841, 1, 17, 'admin');
INSERT INTO `sys_role_menus` VALUES (842, 1, 18, 'admin');
INSERT INTO `sys_role_menus` VALUES (843, 1, 19, 'admin');
INSERT INTO `sys_role_menus` VALUES (844, 1, 20, 'admin');
INSERT INTO `sys_role_menus` VALUES (845, 1, 21, 'admin');
INSERT INTO `sys_role_menus` VALUES (846, 1, 22, 'admin');
INSERT INTO `sys_role_menus` VALUES (847, 1, 23, 'admin');
INSERT INTO `sys_role_menus` VALUES (848, 1, 24, 'admin');
INSERT INTO `sys_role_menus` VALUES (849, 1, 25, 'admin');
INSERT INTO `sys_role_menus` VALUES (850, 1, 26, 'admin');
INSERT INTO `sys_role_menus` VALUES (851, 1, 28, 'admin');
INSERT INTO `sys_role_menus` VALUES (852, 1, 29, 'admin');
INSERT INTO `sys_role_menus` VALUES (853, 1, 30, 'admin');
INSERT INTO `sys_role_menus` VALUES (854, 1, 31, 'admin');
INSERT INTO `sys_role_menus` VALUES (855, 1, 32, 'admin');
INSERT INTO `sys_role_menus` VALUES (856, 1, 33, 'admin');
INSERT INTO `sys_role_menus` VALUES (857, 1, 34, 'admin');
INSERT INTO `sys_role_menus` VALUES (858, 1, 35, 'admin');
INSERT INTO `sys_role_menus` VALUES (859, 1, 36, 'admin');
INSERT INTO `sys_role_menus` VALUES (860, 1, 37, 'admin');
INSERT INTO `sys_role_menus` VALUES (861, 1, 38, 'admin');
INSERT INTO `sys_role_menus` VALUES (862, 1, 39, 'admin');
INSERT INTO `sys_role_menus` VALUES (863, 1, 40, 'admin');
INSERT INTO `sys_role_menus` VALUES (864, 1, 41, 'admin');
INSERT INTO `sys_role_menus` VALUES (865, 1, 42, 'admin');
INSERT INTO `sys_role_menus` VALUES (866, 1, 43, 'admin');
INSERT INTO `sys_role_menus` VALUES (867, 1, 44, 'admin');
INSERT INTO `sys_role_menus` VALUES (868, 1, 45, 'admin');
INSERT INTO `sys_role_menus` VALUES (869, 1, 46, 'admin');
INSERT INTO `sys_role_menus` VALUES (870, 1, 47, 'admin');
INSERT INTO `sys_role_menus` VALUES (871, 1, 48, 'admin');
INSERT INTO `sys_role_menus` VALUES (872, 1, 49, 'admin');
INSERT INTO `sys_role_menus` VALUES (873, 1, 50, 'admin');
INSERT INTO `sys_role_menus` VALUES (874, 1, 51, 'admin');
INSERT INTO `sys_role_menus` VALUES (875, 1, 52, 'admin');
INSERT INTO `sys_role_menus` VALUES (876, 1, 53, 'admin');
INSERT INTO `sys_role_menus` VALUES (877, 1, 54, 'admin');
INSERT INTO `sys_role_menus` VALUES (878, 1, 55, 'admin');
INSERT INTO `sys_role_menus` VALUES (879, 1, 56, 'admin');
INSERT INTO `sys_role_menus` VALUES (880, 1, 57, 'admin');
INSERT INTO `sys_role_menus` VALUES (881, 1, 58, 'admin');
INSERT INTO `sys_role_menus` VALUES (882, 2, 1, 'manage');
INSERT INTO `sys_role_menus` VALUES (883, 2, 3, 'manage');
INSERT INTO `sys_role_menus` VALUES (884, 2, 4, 'manage');
INSERT INTO `sys_role_menus` VALUES (885, 2, 5, 'manage');
INSERT INTO `sys_role_menus` VALUES (886, 2, 6, 'manage');
INSERT INTO `sys_role_menus` VALUES (887, 2, 7, 'manage');
INSERT INTO `sys_role_menus` VALUES (888, 2, 8, 'manage');
INSERT INTO `sys_role_menus` VALUES (889, 2, 9, 'manage');
INSERT INTO `sys_role_menus` VALUES (890, 2, 10, 'manage');
INSERT INTO `sys_role_menus` VALUES (891, 2, 11, 'manage');
INSERT INTO `sys_role_menus` VALUES (892, 2, 12, 'manage');
INSERT INTO `sys_role_menus` VALUES (893, 2, 13, 'manage');
INSERT INTO `sys_role_menus` VALUES (894, 2, 14, 'manage');
INSERT INTO `sys_role_menus` VALUES (895, 2, 15, 'manage');
INSERT INTO `sys_role_menus` VALUES (896, 2, 16, 'manage');
INSERT INTO `sys_role_menus` VALUES (897, 2, 17, 'manage');
INSERT INTO `sys_role_menus` VALUES (898, 2, 18, 'manage');
INSERT INTO `sys_role_menus` VALUES (899, 2, 19, 'manage');
INSERT INTO `sys_role_menus` VALUES (900, 2, 20, 'manage');
INSERT INTO `sys_role_menus` VALUES (901, 2, 21, 'manage');
INSERT INTO `sys_role_menus` VALUES (902, 2, 22, 'manage');
INSERT INTO `sys_role_menus` VALUES (903, 2, 23, 'manage');
INSERT INTO `sys_role_menus` VALUES (904, 2, 24, 'manage');
INSERT INTO `sys_role_menus` VALUES (905, 2, 25, 'manage');
INSERT INTO `sys_role_menus` VALUES (906, 2, 26, 'manage');
INSERT INTO `sys_role_menus` VALUES (907, 2, 28, 'manage');
INSERT INTO `sys_role_menus` VALUES (908, 2, 29, 'manage');
INSERT INTO `sys_role_menus` VALUES (909, 2, 30, 'manage');
INSERT INTO `sys_role_menus` VALUES (910, 2, 31, 'manage');
INSERT INTO `sys_role_menus` VALUES (911, 2, 32, 'manage');
INSERT INTO `sys_role_menus` VALUES (912, 2, 33, 'manage');
INSERT INTO `sys_role_menus` VALUES (913, 2, 34, 'manage');
INSERT INTO `sys_role_menus` VALUES (914, 2, 35, 'manage');
INSERT INTO `sys_role_menus` VALUES (915, 2, 36, 'manage');
INSERT INTO `sys_role_menus` VALUES (916, 2, 37, 'manage');
INSERT INTO `sys_role_menus` VALUES (917, 2, 38, 'manage');
INSERT INTO `sys_role_menus` VALUES (918, 2, 39, 'manage');
INSERT INTO `sys_role_menus` VALUES (919, 2, 40, 'manage');
INSERT INTO `sys_role_menus` VALUES (920, 2, 41, 'manage');
INSERT INTO `sys_role_menus` VALUES (921, 2, 42, 'manage');
INSERT INTO `sys_role_menus` VALUES (922, 2, 43, 'manage');
INSERT INTO `sys_role_menus` VALUES (923, 2, 44, 'manage');
INSERT INTO `sys_role_menus` VALUES (924, 2, 45, 'manage');
INSERT INTO `sys_role_menus` VALUES (925, 2, 46, 'manage');
INSERT INTO `sys_role_menus` VALUES (926, 2, 47, 'manage');
INSERT INTO `sys_role_menus` VALUES (927, 2, 48, 'manage');
INSERT INTO `sys_role_menus` VALUES (928, 2, 49, 'manage');
INSERT INTO `sys_role_menus` VALUES (929, 2, 50, 'manage');
INSERT INTO `sys_role_menus` VALUES (930, 2, 51, 'manage');
INSERT INTO `sys_role_menus` VALUES (931, 2, 52, 'manage');
INSERT INTO `sys_role_menus` VALUES (932, 2, 53, 'manage');
INSERT INTO `sys_role_menus` VALUES (933, 2, 54, 'manage');
INSERT INTO `sys_role_menus` VALUES (934, 2, 55, 'manage');
INSERT INTO `sys_role_menus` VALUES (935, 2, 57, 'manage');
INSERT INTO `sys_role_menus` VALUES (936, 2, 58, 'manage');

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色代码',
  `data_scope` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `role_sort` int(4) NULL DEFAULT NULL COMMENT '角色排序',
  `flag` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '删除标识',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO `sys_roles` VALUES (1, '超管理员', '0', 'admin', '2', 1, '', 'admin', 'panda', '', '2021-12-02 16:03:26', '2021-12-28 15:16:11', NULL);
INSERT INTO `sys_roles` VALUES (2, '管理员', '0', 'manage', '', 2, '', 'panda', 'panda', '', '2021-12-19 16:06:20', '2021-12-28 15:16:23', NULL);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role_id` int(11) NULL DEFAULT NULL,
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `dept_id` int(11) NULL DEFAULT NULL,
  `post_id` int(11) NULL DEFAULT NULL,
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
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, 'pandax', '13818888888', 1, NULL, '', '0', '1@qq.com', 2, 1, 'admin', '1', NULL, '0', '2021-12-03 09:46:55', '2021-12-03 10:51:34', NULL, 'panda', '$2a$10$EXMJ5huCCTlCmP2ITFkAJ.4Mgmq3JcZGUvtE.KLX8j7FmhiiTEEya', '1', '1');
INSERT INTO `sys_users` VALUES (3, '测试用户', '18435234356', 1, '', '', '0', '342@163.com', 3, 1, 'admin', '', '', '0', '2021-12-06 15:16:53', '2021-12-06 15:29:28', NULL, 'test', '$2a$10$4cHTracxWJLdhMmazvbm1urKyD3v5N2AYxAFtNYh6juU39kgae73e', '1', '1,4');
INSERT INTO `sys_users` VALUES (4, 'panda', '18353366912', 2, '', '', '0', '2417920382@qq.com', 2, 4, 'panda', '', '', '0', '2021-12-19 15:58:09', '2021-12-19 16:06:54', NULL, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '2', '4,1');

SET FOREIGN_KEY_CHECKS = 1;
