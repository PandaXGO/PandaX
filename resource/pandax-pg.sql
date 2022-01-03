/*
 Navicat Premium Data Transfer

 Source Server         : localhost_pg
 Source Server Type    : PostgreSQL
 Source Server Version : 120001
 Source Host           : localhost:5432
 Source Catalog        : pandax
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 120001
 File Encoding         : 65001

 Date: 01/01/2022 13:53:04
*/


-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS "public"."casbin_rule";
CREATE TABLE "public"."casbin_rule" (
  "ptype" varchar(100) COLLATE "pg_catalog"."default",
  "v0" varchar(100) COLLATE "pg_catalog"."default",
  "v1" varchar(100) COLLATE "pg_catalog"."default",
  "v2" varchar(100) COLLATE "pg_catalog"."default",
  "v3" varchar(100) COLLATE "pg_catalog"."default",
  "v4" varchar(100) COLLATE "pg_catalog"."default",
  "v5" varchar(100) COLLATE "pg_catalog"."default",
  "id" int8 NOT NULL DEFAULT nextval('casbin_rule_id_seq'::regclass)
)
;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/job', 'POST', '', '', '', 34);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/job', 'PUT', '', '', '', 35);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/job/:jobId', 'DELETE', '', '', '', 37);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/job/:jobId', 'GET', '', '', '', 36);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/job/changeStatus', 'PUT', '', '', '', 40);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/job/list', 'GET', '', '', '', 33);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/job/start/:jobId', 'GET', '', '', '', 39);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/job/stop/:jobId', 'GET', '', '', '', 38);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logJob/:logId', 'DELETE', '', '', '', 49);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logJob/all', 'DELETE', '', '', '', 48);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logJob/list', 'GET', '', '', '', 47);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logLogin/:infoId', 'DELETE', '', '', '', 42);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logLogin/all', 'DELETE', '', '', '', 43);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logLogin/list', 'GET', '', '', '', 41);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logOper/:operId', 'DELETE', '', '', '', 45);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logOper/all', 'DELETE', '', '', '', 46);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/log/logOper/list', 'GET', '', '', '', 44);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/api', 'POST', '', '', '', 5);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/api', 'PUT', '', '', '', 6);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/api/:id', 'DELETE', '', '', '', 7);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/api/:id', 'GET', '', '', '', 4);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/api/all', 'GET', '', '', '', 2);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/api/getPolicyPathByRoleId', 'GET', '', '', '', 3);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/api/list', 'GET', '', '', '', 1);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/config', 'POST', '', '', '', 11);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/config', 'PUT', '', '', '', 12);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/config/:configId', 'DELETE', '', '', '', 13);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/config/:configId', 'GET', '', '', '', 10);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/config/configKey', 'GET', '', '', '', 9);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/config/list', 'GET', '', '', '', 8);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dept', 'POST', '', '', '', 18);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dept', 'PUT', '', '', '', 19);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dept/:deptId', 'DELETE', '', '', '', 20);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dept/:deptId', 'GET', '', '', '', 15);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dept/deptTree', 'GET', '', '', '', 17);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dept/list', 'GET', '', '', '', 14);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', '', 16);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/data', 'POST', '', '', '', 30);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/data', 'PUT', '', '', '', 31);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/data/:dictCode', 'DELETE', '', '', '', 32);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/data/:dictCode', 'GET', '', '', '', 29);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/data/list', 'GET', '', '', '', 27);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/data/type', 'GET', '', '', '', 28);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/type', 'POST', '', '', '', 23);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/type', 'PUT', '', '', '', 24);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/type/:dictId', 'DELETE', '', '', '', 25);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/type/:dictId', 'GET', '', '', '', 22);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/type/export', 'GET', '', '', '', 26);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/dict/type/list', 'GET', '', '', '', 21);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu', 'POST', '', '', '', 56);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu', 'PUT', '', '', '', 57);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu/:menuId', 'DELETE', '', '', '', 58);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu/:menuId', 'GET', '', '', '', 55);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu/list', 'GET', '', '', '', 54);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu/menuPaths', 'GET', '', '', '', 53);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu/menuRole', 'GET', '', '', '', 51);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu/menuTreeSelect', 'GET', '', '', '', 50);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', '', 52);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/notice', 'POST', '', '', '', 60);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/notice', 'PUT', '', '', '', 61);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/notice/:noticeId', 'DELETE', '', '', '', 62);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/notice/list', 'GET', '', '', '', 59);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/post', 'POST', '', '', '', 65);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/post', 'PUT', '', '', '', 66);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/post/:postId', 'DELETE', '', '', '', 67);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/post/:postId', 'GET', '', '', '', 64);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/post/list', 'GET', '', '', '', 63);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/role', 'POST', '', '', '', 70);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/role', 'PUT', '', '', '', 71);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/role/:roleId', 'DELETE', '', '', '', 72);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/role/:roleId', 'GET', '', '', '', 69);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/role/changeStatus', 'PUT', '', '', '', 73);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/role/dataScope', 'PUT', '', '', '', 74);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/role/export', 'GET', '', '', '', 75);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/role/list', 'GET', '', '', '', 68);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user', 'POST', '', '', '', 84);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user', 'PUT', '', '', '', 85);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/:userId', 'DELETE', '', '', '', 78);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/avatar', 'POST', '', '', '', 79);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/changeStatus', 'PUT', '', '', '', 77);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/export', 'GET', '', '', '', 86);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/getById/:userId', 'GET', '', '', '', 81);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/getInit', 'GET', '', '', '', 82);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/getRoPo', 'GET', '', '', '', 83);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/list', 'GET', '', '', '', 76);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'admin', '/system/user/pwd', 'PUT', '', '', '', 80);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/job/:jobId', 'GET', '', '', '', 104);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/job/changeStatus', 'PUT', '', '', '', 105);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/job/list', 'GET', '', '', '', 103);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/log/logJob/list', 'GET', '', '', '', 108);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/log/logLogin/list', 'GET', '', '', '', 106);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/log/logOper/list', 'GET', '', '', '', 107);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/api/:id', 'GET', '', '', '', 90);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/api/all', 'GET', '', '', '', 88);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/api/getPolicyPathByRoleId', 'GET', '', '', '', 89);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/api/list', 'GET', '', '', '', 87);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/config/:configId', 'GET', '', '', '', 93);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/config/configKey', 'GET', '', '', '', 92);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/config/list', 'GET', '', '', '', 91);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dept/:deptId', 'GET', '', '', '', 95);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dept/deptTree', 'GET', '', '', '', 97);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dept/list', 'GET', '', '', '', 94);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', '', 96);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dict/data/:dictCode', 'GET', '', '', '', 102);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dict/data/list', 'GET', '', '', '', 100);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dict/data/type', 'GET', '', '', '', 101);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dict/type/:dictId', 'GET', '', '', '', 99);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/dict/type/list', 'GET', '', '', '', 98);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/menu/:menuId', 'GET', '', '', '', 114);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/menu/list', 'GET', '', '', '', 113);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/menu/menuPaths', 'GET', '', '', '', 112);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/menu/menuRole', 'GET', '', '', '', 110);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/menu/menuTreeSelect', 'GET', '', '', '', 109);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', '', 111);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/notice/list', 'GET', '', '', '', 115);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/post/:postId', 'GET', '', '', '', 117);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/post/list', 'GET', '', '', '', 116);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/role/:roleId', 'GET', '', '', '', 119);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/role/list', 'GET', '', '', '', 118);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/user/avatar', 'POST', '', '', '', 121);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/user/getById/:userId', 'GET', '', '', '', 122);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/user/getInit', 'GET', '', '', '', 123);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/user/getRoPo', 'GET', '', '', '', 124);
INSERT INTO "public"."casbin_rule" VALUES ('p', 'manage', '/system/user/list', 'GET', '', '', '', 120);

-- ----------------------------
-- Table structure for log_jobs
-- ----------------------------
DROP TABLE IF EXISTS "public"."log_jobs";
CREATE TABLE "public"."log_jobs" (
  "log_id" int8 NOT NULL DEFAULT nextval('log_jobs_log_id_seq'::regclass),
  "name" varchar(128) COLLATE "pg_catalog"."default",
  "job_group" varchar(128) COLLATE "pg_catalog"."default",
  "entry_id" int8,
  "invoke_target" varchar(128) COLLATE "pg_catalog"."default",
  "log_info" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."log_jobs"."name" IS '任务名称';
COMMENT ON COLUMN "public"."log_jobs"."job_group" IS '分组';
COMMENT ON COLUMN "public"."log_jobs"."entry_id" IS '任务id';
COMMENT ON COLUMN "public"."log_jobs"."invoke_target" IS '调用方法';
COMMENT ON COLUMN "public"."log_jobs"."log_info" IS '日志信息';
COMMENT ON COLUMN "public"."log_jobs"."status" IS '状态';

-- ----------------------------
-- Records of log_jobs
-- ----------------------------

-- ----------------------------
-- Table structure for log_logins
-- ----------------------------
DROP TABLE IF EXISTS "public"."log_logins";
CREATE TABLE "public"."log_logins" (
  "info_id" int8 NOT NULL DEFAULT nextval('log_logins_info_id_seq'::regclass),
  "username" varchar(128) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "ipaddr" varchar(255) COLLATE "pg_catalog"."default",
  "login_location" varchar(255) COLLATE "pg_catalog"."default",
  "browser" varchar(255) COLLATE "pg_catalog"."default",
  "os" varchar(255) COLLATE "pg_catalog"."default",
  "platform" varchar(255) COLLATE "pg_catalog"."default",
  "login_time" timestamp(6),
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "msg" varchar(255) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."log_logins"."username" IS '用户名';
COMMENT ON COLUMN "public"."log_logins"."status" IS '状态';
COMMENT ON COLUMN "public"."log_logins"."ipaddr" IS 'ip地址';
COMMENT ON COLUMN "public"."log_logins"."login_location" IS '归属地';
COMMENT ON COLUMN "public"."log_logins"."browser" IS '浏览器';
COMMENT ON COLUMN "public"."log_logins"."os" IS '系统';
COMMENT ON COLUMN "public"."log_logins"."platform" IS '固件';
COMMENT ON COLUMN "public"."log_logins"."login_time" IS '登录时间';
COMMENT ON COLUMN "public"."log_logins"."create_by" IS '创建人';
COMMENT ON COLUMN "public"."log_logins"."update_by" IS '更新者';

-- ----------------------------
-- Records of log_logins
-- ----------------------------
INSERT INTO "public"."log_logins" VALUES (1, 'panda', '0', '127.0.0.1', '内部IP', 'Chrome 96.0.4664.110', 'Windows 10', 'Windows', '2022-01-01 12:32:51.697576', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36', '登录成功', '2022-01-01 04:32:51.699189+00', '2022-01-01 04:32:51.699189+00', NULL);
INSERT INTO "public"."log_logins" VALUES (2, 'panda', '0', '127.0.0.1', '内部IP', 'Chrome 96.0.4664.110', 'Windows 10', 'Windows', '2022-01-01 13:50:11.578926', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36', '登录成功', '2022-01-01 05:50:11.580002+00', '2022-01-01 05:50:11.580002+00', NULL);

-- ----------------------------
-- Table structure for log_opers
-- ----------------------------
DROP TABLE IF EXISTS "public"."log_opers";
CREATE TABLE "public"."log_opers" (
  "oper_id" int8 NOT NULL DEFAULT nextval('log_opers_oper_id_seq'::regclass),
  "title" varchar(128) COLLATE "pg_catalog"."default",
  "business_type" varchar(1) COLLATE "pg_catalog"."default",
  "method" varchar(255) COLLATE "pg_catalog"."default",
  "oper_name" varchar(255) COLLATE "pg_catalog"."default",
  "oper_url" varchar(255) COLLATE "pg_catalog"."default",
  "oper_ip" varchar(255) COLLATE "pg_catalog"."default",
  "oper_location" varchar(255) COLLATE "pg_catalog"."default",
  "oper_param" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."log_opers"."title" IS '操作的模块';
COMMENT ON COLUMN "public"."log_opers"."business_type" IS '0其它 1新增 2修改 3删除';
COMMENT ON COLUMN "public"."log_opers"."method" IS '请求方法';
COMMENT ON COLUMN "public"."log_opers"."oper_name" IS '操作人员';
COMMENT ON COLUMN "public"."log_opers"."oper_url" IS '操作url';
COMMENT ON COLUMN "public"."log_opers"."oper_ip" IS '操作IP';
COMMENT ON COLUMN "public"."log_opers"."oper_location" IS '操作地点';
COMMENT ON COLUMN "public"."log_opers"."oper_param" IS '请求参数';
COMMENT ON COLUMN "public"."log_opers"."status" IS '0=正常,1=异常';

-- ----------------------------
-- Records of log_opers
-- ----------------------------
INSERT INTO "public"."log_opers" VALUES (1, '删除部门信息', '3', 'DELETE', 'admin', '/system/dept/7', '127.0.0.1', '内部IP', '', '0', '2022-01-01 03:50:22.688828+00', '2022-01-01 03:50:22.688828+00', NULL);

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_apis";
CREATE TABLE "public"."sys_apis" (
  "id" int8 NOT NULL DEFAULT nextval('sys_apis_id_seq'::regclass),
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6),
  "path" text COLLATE "pg_catalog"."default",
  "description" text COLLATE "pg_catalog"."default",
  "api_group" text COLLATE "pg_catalog"."default",
  "method" text COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_apis"."path" IS 'api路径';
COMMENT ON COLUMN "public"."sys_apis"."description" IS 'api中文描述';
COMMENT ON COLUMN "public"."sys_apis"."api_group" IS 'api组';
COMMENT ON COLUMN "public"."sys_apis"."method" IS '方法';

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO "public"."sys_apis" VALUES (1, '2021-12-09 09:21:04+00', '2021-12-09 09:21:04+00', NULL, '/system/user/list', '查询用户列表（分页）', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (2, '2021-12-09 09:29:36+00', '2021-12-09 09:29:36+00', NULL, '/system/user/changeStatus', '修改用户状态', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (3, '2021-12-09 09:34:37+00', '2021-12-09 09:34:37+00', NULL, '/system/user/:userId', '删除用户信息', 'user', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (4, '2021-12-09 09:36:43+00', '2021-12-09 09:36:43+00', NULL, '/system/dept/list', '获取部门列表', 'dept', 'GET');
INSERT INTO "public"."sys_apis" VALUES (5, '2021-12-09 09:37:31+00', '2021-12-09 09:37:31+00', NULL, '/system/dept/:deptId', '获取部门信息', 'dept', 'GET');
INSERT INTO "public"."sys_apis" VALUES (6, '2021-12-09 18:20:32+00', '2021-12-09 18:20:32+00', NULL, '/system/user/avatar', '上传头像', 'user', 'POST');
INSERT INTO "public"."sys_apis" VALUES (7, '2021-12-09 18:21:10+00', '2021-12-09 18:21:10+00', NULL, '/system/user/pwd', '修改密码', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (8, '2021-12-09 18:21:54+00', '2021-12-09 18:21:54+00', NULL, '/system/user/getById/:userId', '通过id获取用户信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (9, '2021-12-09 18:58:50+00', '2021-12-09 18:58:50+00', NULL, '/system/user/getInit', '获取初始化角色岗位信息(添加用户初始化)', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (10, '2021-12-09 18:59:43+00', '2021-12-09 18:59:43+00', NULL, '/system/user/getRoPo', '获取用户角色岗位信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (11, '2021-12-09 19:00:24+00', '2021-12-09 19:00:24+00', NULL, '/system/user', '添加用户信息', 'user', 'POST');
INSERT INTO "public"."sys_apis" VALUES (12, '2021-12-09 19:00:52+00', '2021-12-09 19:00:52+00', NULL, '/system/user', '修改用户信息', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (13, '2021-12-09 19:02:30+00', '2021-12-09 19:02:30+00', NULL, '/system/user/export', '导出用户信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (14, '2021-12-09 19:04:04+00', '2021-12-09 19:04:04+00', NULL, '/system/dept/roleDeptTreeSelect/:roleId', '获取角色部门树', 'dept', 'GET');
INSERT INTO "public"."sys_apis" VALUES (15, '2021-12-09 19:04:48+00', '2021-12-09 19:04:48+00', NULL, '/system/dept/deptTree', '获取所有部门树', 'dept', 'GET');
INSERT INTO "public"."sys_apis" VALUES (16, '2021-12-09 19:07:37+00', '2021-12-09 19:07:37+00', NULL, '/system/dept', '添加部门信息', 'dept', 'POST');
INSERT INTO "public"."sys_apis" VALUES (17, '2021-12-09 19:08:14+00', '2021-12-09 19:08:14+00', NULL, '/system/dept', '修改部门信息', 'dept', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (18, '2021-12-09 19:08:40+00', '2021-12-09 19:08:40+00', NULL, '/system/dept/:deptId', '删除部门信息', 'dept', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (19, '2021-12-09 19:09:41+00', '2021-12-09 19:09:41+00', NULL, '/system/config/list', '获取配置分页列表', 'config', 'GET');
INSERT INTO "public"."sys_apis" VALUES (20, '2021-12-09 19:10:11+00', '2021-12-09 19:10:11+00', NULL, '/system/config/configKey', '获取配置列表通过ConfigKey', 'config', 'GET');
INSERT INTO "public"."sys_apis" VALUES (21, '2021-12-09 19:10:45+00', '2021-12-09 19:10:45+00', NULL, '/system/config/:configId', '获取配置信息', 'config', 'GET');
INSERT INTO "public"."sys_apis" VALUES (22, '2021-12-09 19:11:22+00', '2021-12-09 19:11:22+00', NULL, '/system/config', '添加配置信息', 'config', 'POST');
INSERT INTO "public"."sys_apis" VALUES (23, '2021-12-09 19:11:41+00', '2021-12-09 19:11:41+00', NULL, '/system/config', '修改配置信息', 'config', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (24, '2021-12-09 19:12:28+00', '2021-12-09 19:12:28+00', NULL, '/system/config/:configId', '删除配置信息', 'config', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (25, '2021-12-09 19:13:08+00', '2021-12-09 19:13:08+00', NULL, '/system/dict/type/list', '获取字典类型分页列表', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (26, '2021-12-09 19:13:55+00', '2021-12-09 19:13:55+00', NULL, '/system/dict/type/:dictId', '获取字典类型信息', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (27, '2021-12-09 19:14:28+00', '2021-12-09 19:14:28+00', NULL, '/system/dict/type', '添加字典类型信息', 'dict', 'POST');
INSERT INTO "public"."sys_apis" VALUES (28, '2021-12-09 19:14:55+00', '2021-12-09 19:14:55+00', NULL, '/system/dict/type', '修改字典类型信息', 'dict', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (29, '2021-12-09 19:15:17+00', '2021-12-09 19:15:17+00', NULL, '/system/dict/type/:dictId', '删除字典类型信息', 'dict', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (30, '2021-12-09 19:15:50+00', '2021-12-09 19:15:50+00', NULL, '/system/dict/type/export', '导出字典类型信息', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (31, '2021-12-09 19:16:26+00', '2021-12-09 19:16:26+00', NULL, '/system/dict/data/list', '获取字典数据分页列表', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (32, '2021-12-09 19:17:01+00', '2021-12-09 19:17:01+00', NULL, '/system/dict/data/type', '获取字典数据列表通过字典类型', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (33, '2021-12-09 19:17:39+00', '2021-12-09 19:17:39+00', NULL, '/system/dict/data/:dictCode', '获取字典数据信息', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (34, '2021-12-09 19:18:20+00', '2021-12-09 19:18:20+00', NULL, '/system/dict/data', '添加字典数据信息', 'dict', 'POST');
INSERT INTO "public"."sys_apis" VALUES (35, '2021-12-09 19:18:44+00', '2021-12-09 19:18:44+00', NULL, '/system/dict/data', '修改字典数据信息', 'dict', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (36, '2021-12-09 19:19:16+00', '2021-12-09 19:19:16+00', NULL, '/system/dict/data/:dictCode', '删除字典数据信息', 'dict', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (37, '2021-12-09 19:21:18+00', '2021-12-09 19:21:18+00', NULL, '/system/menu/menuTreeSelect', '获取菜单树', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (38, '2021-12-09 19:21:47+00', '2021-12-09 19:21:47+00', NULL, '/system/menu/menuRole', '获取角色菜单', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (39, '2021-12-09 19:22:42+00', '2021-12-09 19:22:42+00', NULL, '/system/menu/roleMenuTreeSelect/:roleId', '获取角色菜单树', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (40, '2021-12-09 19:23:17+00', '2021-12-09 19:23:17+00', NULL, '/system/menu/menuPaths', '获取角色菜单路径列表', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (41, '2021-12-09 19:23:40+00', '2021-12-09 19:23:40+00', NULL, '/system/menu/list', '获取菜单列表', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (42, '2021-12-09 19:24:09+00', '2021-12-09 19:24:09+00', NULL, '/system/menu/:menuId', '获取菜单信息', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (43, '2021-12-09 19:24:29+00', '2021-12-09 19:24:29+00', NULL, '/system/menu', '添加菜单信息', 'menu', 'POST');
INSERT INTO "public"."sys_apis" VALUES (44, '2021-12-09 19:24:48+00', '2021-12-09 19:24:48+00', NULL, '/system/menu', '修改菜单信息', 'menu', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (45, '2021-12-09 19:25:10+00', '2021-12-09 19:25:10+00', NULL, '/system/menu/:menuId', '删除菜单信息', 'menu', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (46, '2021-12-09 19:25:44+00', '2021-12-09 19:27:06+00', NULL, '/system/post/list', '获取岗位分页列表', 'post', 'GET');
INSERT INTO "public"."sys_apis" VALUES (47, '2021-12-09 19:26:55+00', '2021-12-09 19:26:55+00', NULL, '/system/post/:postId', '获取岗位信息', 'post', 'GET');
INSERT INTO "public"."sys_apis" VALUES (48, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/post', '添加岗位信息', 'post', 'POST');
INSERT INTO "public"."sys_apis" VALUES (49, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/post', '修改岗位信息', 'post', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (50, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/post/:postId', '删除岗位信息', 'post', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (51, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/role/list', '获取角色分页列表', 'role', 'GET');
INSERT INTO "public"."sys_apis" VALUES (52, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/role/:roleId', '获取角色信息', 'role', 'GET');
INSERT INTO "public"."sys_apis" VALUES (53, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/role', '添加角色信息', 'role', 'POST');
INSERT INTO "public"."sys_apis" VALUES (54, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/role', '修改角色信息', 'role', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (55, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/role/:roleId', '删除角色信息', 'role', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (56, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/role/changeStatus', '修改角色状态', 'role', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (57, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/role/dataScope', '修改角色部门权限', 'role', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (58, '2021-12-09 19:25:44+00', '2021-12-09 19:25:44+00', NULL, '/system/role/export', '导出角色信息', 'role', 'GET');
INSERT INTO "public"."sys_apis" VALUES (59, '2021-12-09 19:50:57+00', '2021-12-09 19:50:57+00', NULL, '/system/api/list', '获取api分页列表', 'api', 'GET');
INSERT INTO "public"."sys_apis" VALUES (60, '2021-12-09 19:51:26+00', '2021-12-09 19:51:26+00', NULL, '/system/api/all', '获取所有api', 'api', 'GET');
INSERT INTO "public"."sys_apis" VALUES (61, '2021-12-09 19:51:54+00', '2021-12-09 19:51:54+00', NULL, '/system/api/getPolicyPathByRoleId', '获取角色拥有的api权限', 'api', 'GET');
INSERT INTO "public"."sys_apis" VALUES (62, '2021-12-09 19:52:14+00', '2021-12-09 19:52:14+00', NULL, '/system/api/:id', '获取api信息', 'api', 'GET');
INSERT INTO "public"."sys_apis" VALUES (63, '2021-12-09 19:52:35+00', '2021-12-09 19:52:35+00', NULL, '/system/api', '添加api信息', 'api', 'POST');
INSERT INTO "public"."sys_apis" VALUES (64, '2021-12-09 19:52:50+00', '2021-12-09 19:52:50+00', NULL, '/system/api', '修改api信息', 'api', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (65, '2021-12-09 19:53:07+00', '2021-12-09 19:53:07+00', NULL, '/system/api/:id', '删除api信息', 'api', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (66, '2021-12-17 10:51:05+00', '2021-12-17 10:54:22+00', NULL, '/log/logLogin/list', '获取登录日志', 'log', 'GET');
INSERT INTO "public"."sys_apis" VALUES (67, '2021-12-17 10:51:43+00', '2021-12-17 10:54:28+00', NULL, '/log/logLogin/:infoId', '删除日志', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (68, '2021-12-17 10:53:09+00', '2021-12-17 10:54:34+00', NULL, '/log/logLogin/all', '清空所有', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (69, '2021-12-17 10:54:07+00', '2021-12-17 10:54:07+00', NULL, '/log/logOper/list', '操作日志列表', 'log', 'GET');
INSERT INTO "public"."sys_apis" VALUES (70, '2021-12-17 10:53:09+00', '2021-12-17 10:53:09+00', NULL, '/log/logOper/:operId', '删除', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (71, '2021-12-17 10:53:09+00', '2021-12-17 10:53:09+00', NULL, '/log/logOper/all', '清空', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (72, '2021-12-24 15:41:23+00', '2021-12-24 15:41:23+00', NULL, '/job/list', '任务列表', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (73, '2021-12-24 15:41:54+00', '2021-12-24 15:41:54+00', NULL, '/job', '添加', 'job', 'POST');
INSERT INTO "public"."sys_apis" VALUES (74, '2021-12-24 15:42:11+00', '2021-12-24 15:42:11+00', NULL, '/job', '修改任务', 'job', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (75, '2021-12-24 15:42:37+00', '2021-12-24 16:32:21+00', NULL, '/job/:jobId', '获取任务', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (76, '2021-12-24 15:43:09+00', '2021-12-24 16:32:05+00', NULL, '/job/:jobId', '删除job', 'job', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (77, '2021-12-24 15:43:35+00', '2021-12-24 16:31:11+00', NULL, '/job/stop/:jobId', '停止任务', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (78, '2021-12-24 15:44:09+00', '2021-12-24 16:30:38+00', NULL, '/job/start/:jobId', '开始任务', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (79, '2021-12-24 15:45:03+00', '2021-12-24 15:46:36+00', NULL, '/log/logJob/list', '任务日志列表', 'log', 'GET');
INSERT INTO "public"."sys_apis" VALUES (80, '2021-12-24 15:45:33+00', '2021-12-24 15:46:43+00', NULL, '/log/logJob/all', '清空任务日志', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (81, '2021-12-24 15:46:08+00', '2021-12-24 16:33:13+00', NULL, '/log/logJob/:logId', '删除任务日志', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (82, '2021-12-24 15:45:33+00', '2021-12-24 15:45:33+00', NULL, '/system/notice/list', '获取通知分页列表', 'notice', 'GET');
INSERT INTO "public"."sys_apis" VALUES (83, '2021-12-24 15:45:33+00', '2021-12-24 15:45:33+00', NULL, '/system/notice', '添加通知信息', 'notice', 'POST');
INSERT INTO "public"."sys_apis" VALUES (84, '2021-12-24 15:45:33+00', '2021-12-24 15:45:33+00', NULL, '/system/notice', '修改通知信息', 'notice', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (85, '2021-12-24 15:45:33+00', '2021-12-24 16:33:48+00', NULL, '/system/notice/:noticeId', '删除通知信息', 'notice', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (86, '2021-12-24 22:40:19+00', '2021-12-24 22:40:19+00', NULL, '/job/changeStatus', '修改状态', 'job', 'PUT');

-- ----------------------------
-- Table structure for sys_configs
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_configs";
CREATE TABLE "public"."sys_configs" (
  "config_id" int8 NOT NULL DEFAULT nextval('sys_configs_config_id_seq'::regclass),
  "config_name" varchar(128) COLLATE "pg_catalog"."default",
  "config_key" varchar(128) COLLATE "pg_catalog"."default",
  "config_value" varchar(255) COLLATE "pg_catalog"."default",
  "config_type" varchar(64) COLLATE "pg_catalog"."default",
  "is_frontend" varchar(1) COLLATE "pg_catalog"."default",
  "remark" varchar(128) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_configs"."config_id" IS '主键编码';
COMMENT ON COLUMN "public"."sys_configs"."config_name" IS 'ConfigName';
COMMENT ON COLUMN "public"."sys_configs"."config_key" IS 'ConfigKey';
COMMENT ON COLUMN "public"."sys_configs"."config_value" IS 'ConfigValue';
COMMENT ON COLUMN "public"."sys_configs"."config_type" IS '是否系统内置0，1';
COMMENT ON COLUMN "public"."sys_configs"."is_frontend" IS '是否前台';
COMMENT ON COLUMN "public"."sys_configs"."remark" IS 'Remark';

-- ----------------------------
-- Records of sys_configs
-- ----------------------------
INSERT INTO "public"."sys_configs" VALUES (1, '账号初始密码', 'sys.user.initPassword', '123456', '0', '0', '初始密码', '2021-12-04 13:50:13+00', '2021-12-04 13:54:52+00', NULL);

-- ----------------------------
-- Table structure for sys_depts
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_depts";
CREATE TABLE "public"."sys_depts" (
  "dept_id" int8 NOT NULL DEFAULT nextval('sys_depts_dept_id_seq'::regclass),
  "parent_id" int8,
  "dept_path" varchar(255) COLLATE "pg_catalog"."default",
  "dept_name" varchar(128) COLLATE "pg_catalog"."default",
  "sort" int8,
  "leader" varchar(64) COLLATE "pg_catalog"."default",
  "phone" varchar(11) COLLATE "pg_catalog"."default",
  "email" varchar(64) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_by" varchar(64) COLLATE "pg_catalog"."default",
  "update_by" varchar(64) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_depts"."parent_id" IS '上级部门';
COMMENT ON COLUMN "public"."sys_depts"."dept_path" IS '部门路径';
COMMENT ON COLUMN "public"."sys_depts"."dept_name" IS '部门名称';
COMMENT ON COLUMN "public"."sys_depts"."sort" IS '排序';
COMMENT ON COLUMN "public"."sys_depts"."leader" IS '负责人';
COMMENT ON COLUMN "public"."sys_depts"."phone" IS '手机';
COMMENT ON COLUMN "public"."sys_depts"."email" IS '邮箱';
COMMENT ON COLUMN "public"."sys_depts"."status" IS '状态';
COMMENT ON COLUMN "public"."sys_depts"."create_by" IS '创建人';
COMMENT ON COLUMN "public"."sys_depts"."update_by" IS '修改人';

-- ----------------------------
-- Records of sys_depts
-- ----------------------------
INSERT INTO "public"."sys_depts" VALUES (2, 0, '/0/2', '熊猫科技', 0, 'xm', '18353366836', '342@qq.com', '0', 'admin', 'admin', '2021-12-01 17:31:53+00', '2021-12-02 08:56:19+00', NULL);
INSERT INTO "public"."sys_depts" VALUES (3, 2, '/0/2/3', '研发部', 1, 'panda', '18353366543', 'ewr@qq.com', '0', 'admin', 'admin', '2021-12-01 17:37:43+00', '2021-12-02 08:55:56+00', NULL);
INSERT INTO "public"."sys_depts" VALUES (7, 2, '/0/2/7', '营销部', 2, 'panda', '18353333333', '342@qq.com', '0', 'panda', 'panda', '2021-12-24 10:46:24+00', '2021-12-24 10:47:15+00', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dict_data";
CREATE TABLE "public"."sys_dict_data" (
  "dict_code" int8 NOT NULL DEFAULT nextval('sys_dict_data_dict_code_seq'::regclass),
  "dict_sort" int8,
  "dict_label" varchar(64) COLLATE "pg_catalog"."default",
  "dict_value" varchar(64) COLLATE "pg_catalog"."default",
  "dict_type" varchar(64) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "css_class" varchar(128) COLLATE "pg_catalog"."default",
  "list_class" varchar(128) COLLATE "pg_catalog"."default",
  "is_default" varchar(8) COLLATE "pg_catalog"."default",
  "create_by" text COLLATE "pg_catalog"."default",
  "update_by" text COLLATE "pg_catalog"."default",
  "remark" varchar(256) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_dict_data"."dict_sort" IS '排序';
COMMENT ON COLUMN "public"."sys_dict_data"."dict_label" IS '标签';
COMMENT ON COLUMN "public"."sys_dict_data"."dict_value" IS '值';
COMMENT ON COLUMN "public"."sys_dict_data"."dict_type" IS '字典类型';
COMMENT ON COLUMN "public"."sys_dict_data"."status" IS '状态（0正常 1停用）';
COMMENT ON COLUMN "public"."sys_dict_data"."css_class" IS 'CssClass';
COMMENT ON COLUMN "public"."sys_dict_data"."list_class" IS 'ListClass';
COMMENT ON COLUMN "public"."sys_dict_data"."is_default" IS 'IsDefault';
COMMENT ON COLUMN "public"."sys_dict_data"."remark" IS '备注';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO "public"."sys_dict_data" VALUES (1, 0, '男', '0', 'sys_user_sex', '0', '', '', '', 'admin', '', '男', '2021-11-30 14:58:18+00', '2021-11-30 14:58:18+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (2, 1, '女', '1', 'sys_user_sex', '0', '', '', '', 'admin', '', '女生', '2021-11-30 15:09:11+00', '2021-11-30 15:10:28+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (3, 2, '未知', '2', 'sys_user_sex', '0', '', '', '', 'admin', '', '未知', '2021-11-30 15:09:11+00', '2021-11-30 15:10:28+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (4, 0, '正常', '0', 'sys_normal_disable', '0', '', '', '', 'admin', '', '', '2021-12-01 15:58:50+00', '2021-12-01 15:58:50+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (5, 1, '停用', '1', 'sys_normal_disable', '0', '', '', '', 'admin', '', '', '2021-12-01 15:59:08+00', '2021-12-01 15:59:08+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (6, 0, '目录', 'M', 'sys_menu_type', '0', '', '', '', 'admin', '', '', '2021-12-02 09:49:12+00', '2021-12-02 09:49:12+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (7, 1, '菜单', 'C', 'sys_menu_type', '0', '', '', '', 'admin', '', '', '2021-12-02 09:49:35+00', '2021-12-02 09:49:52+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (8, 2, '按钮', 'F', 'sys_menu_type', '0', '', '', '', 'admin', '', '', '2021-12-02 09:49:35+00', '2021-12-02 09:49:35+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (9, 0, '显示', '0', 'sys_show_hide', '0', '', '', '', 'admin', '', '', '2021-12-02 09:56:40+00', '2021-12-02 09:56:40+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (10, 0, '隐藏', '1', 'sys_show_hide', '0', '', '', '', 'admin', '', '', '2021-12-02 09:56:52+00', '2021-12-02 09:56:52+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (11, 0, '是', '0', 'sys_num_yes_no', '0', '', '', '', 'admin', '', '', '2021-12-02 10:16:16+00', '2021-12-02 10:16:16+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (12, 1, '否', '1', 'sys_num_yes_no', '0', '', '', '', 'admin', '', '', '2021-12-02 10:16:26+00', '2021-12-02 10:16:26+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (13, 0, '是', '0', 'sys_yes_no', '0', '', '', '', 'admin', '', '', '2021-12-04 13:48:15+00', '2021-12-04 13:48:15+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (14, 0, '否', '1', 'sys_yes_no', '0', '', '', '', 'admin', '', '', '2021-12-04 13:48:21+00', '2021-12-04 13:48:21+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (15, 0, '创建(POST)', 'POST', 'sys_method_api', '0', '', '', '', 'admin', '', '', '2021-12-08 17:22:05+00', '2021-12-09 09:29:52+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (16, 1, '查询(GET)', 'GET', 'sys_method_api', '0', '', '', '', 'admin', '', '', '2021-12-08 17:22:24+00', '2021-12-09 09:29:59+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (17, 2, '修改(PUT)', 'PUT', 'sys_method_api', '0', '', '', '', 'admin', '', '', '2021-12-08 17:22:40+00', '2021-12-09 09:30:06+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (18, 3, '删除(DELETE)', 'DELETE', 'sys_method_api', '0', '', '', '', 'admin', '', '', '2021-12-08 17:22:54+00', '2021-12-09 09:30:13+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (19, 0, '成功', '0', 'sys_common_status', '0', '', '', '', 'admin', '', '', '2021-12-17 11:01:52+00', '2021-12-17 11:01:52+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (20, 0, '失败', '1', 'sys_common_status', '0', '', '', '', 'admin', '', '', '2021-12-17 11:02:08+00', '2021-12-17 11:02:08+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (21, 0, '其他', '0', 'sys_oper_type', '0', '', '', '', 'admin', '', '', '2021-12-17 11:30:07+00', '2021-12-17 11:30:07+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (22, 0, '新增', '1', 'sys_oper_type', '0', '', '', '', 'admin', '', '', '2021-12-17 11:30:21+00', '2021-12-17 11:30:21+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (23, 0, '修改', '2', 'sys_oper_type', '0', '', '', '', 'admin', '', '', '2021-12-17 11:30:32+00', '2021-12-17 11:30:32+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (24, 0, '删除', '3', 'sys_oper_type', '0', '', '', '', 'admin', '', '', '2021-12-17 11:30:40+00', '2021-12-17 11:30:40+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (25, 0, '默认', 'DEFAULT', 'sys_job_group', '0', '', '', '', 'panda', '', '', '2021-12-24 15:15:31+00', '2021-12-24 15:15:31+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (26, 1, '系统', 'SYSTEM', 'sys_job_group', '0', '', '', '', 'panda', '', '', '2021-12-24 15:15:50+00', '2021-12-24 15:15:50+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (27, 0, '发布通知', '1', 'sys_notice_type', '0', '', '', '', 'panda', '', '', '2021-12-26 15:24:07+00', '2021-12-26 15:24:07+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (28, 0, '任免通知', '2', 'sys_notice_type', '0', '', '', '', 'panda', '', '', '2021-12-26 15:24:18+00', '2021-12-26 15:24:18+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (29, 0, '事物通知', '3', 'sys_notice_type', '0', '', '', '', 'panda', '', '', '2021-12-26 15:24:46+00', '2021-12-26 15:24:46+00', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (30, 0, '审批通知', '4', 'sys_notice_type', '0', '', '', '', 'panda', '', '', '2021-12-26 15:25:08+00', '2021-12-26 15:25:08+00', NULL);

-- ----------------------------
-- Table structure for sys_dict_types
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dict_types";
CREATE TABLE "public"."sys_dict_types" (
  "dict_id" int8 NOT NULL DEFAULT nextval('sys_dict_types_dict_id_seq'::regclass),
  "dict_name" varchar(64) COLLATE "pg_catalog"."default",
  "dict_type" varchar(64) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_by" text COLLATE "pg_catalog"."default",
  "update_by" text COLLATE "pg_catalog"."default",
  "remark" varchar(256) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_dict_types"."dict_name" IS '名称';
COMMENT ON COLUMN "public"."sys_dict_types"."dict_type" IS '类型';
COMMENT ON COLUMN "public"."sys_dict_types"."status" IS '状态';
COMMENT ON COLUMN "public"."sys_dict_types"."remark" IS '备注';

-- ----------------------------
-- Records of sys_dict_types
-- ----------------------------
INSERT INTO "public"."sys_dict_types" VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '', '性别列表', '2021-11-30 14:02:52+00', '2021-11-30 14:07:55+00', '2021-11-30 14:11:54+00');
INSERT INTO "public"."sys_dict_types" VALUES (2, '用户性别', 'sys_user_sex', '0', 'admin', '', '用户性别列表', '2021-11-30 14:12:33+00', '2021-11-30 14:12:33+00', '2021-11-30 14:14:19+00');
INSERT INTO "public"."sys_dict_types" VALUES (3, '的心', 'sfd', '0', 'admin', '', 'fs', '2021-11-30 14:13:22+00', '2021-11-30 14:13:22+00', '2021-11-30 14:14:19+00');
INSERT INTO "public"."sys_dict_types" VALUES (4, '用户性别', 'sys_user_sex', '0', 'admin', '', '性别字典', '2021-11-30 14:15:25+00', '2021-11-30 14:15:25+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (5, 'df', 'da', '0', 'admin', '', 'sd', '2021-11-30 15:54:33+00', '2021-11-30 15:54:33+00', '2021-11-30 15:54:40+00');
INSERT INTO "public"."sys_dict_types" VALUES (6, '系统开关', 'sys_normal_disable', '0', 'admin', '', '开关列表', '2021-12-01 15:57:58+00', '2021-12-01 15:57:58+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (7, '菜单类型', 'sys_menu_type', '0', 'admin', '', '菜单类型列表', '2021-12-02 09:48:48+00', '2021-12-02 09:56:12+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (8, '菜单状态', 'sys_show_hide', '0', 'admin', '', '菜单状态列表', '2021-12-02 09:55:59+00', '2021-12-02 09:55:59+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (9, '数字是否', 'sys_num_yes_no', '0', 'admin', '', '数字是否列表', '2021-12-02 10:13:29+00', '2021-12-02 10:13:40+00', '2021-12-02 10:15:07+00');
INSERT INTO "public"."sys_dict_types" VALUES (10, '数字是否', 'sys_num_yes_no', '0', 'admin', '', '数字是否', '2021-12-02 10:13:29+00', '2021-12-02 10:13:29+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (11, '状态是否', 'sys_yes_no', '0', 'admin', '', '状态是否', '2021-12-04 13:47:57+00', '2021-12-04 13:47:57+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (12, '网络请求方法', 'sys_method_api', '0', 'admin', '', '网络请求方法列表', '2021-12-08 17:21:27+00', '2021-12-08 17:21:27+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (13, '成功失败', 'sys_common_status', '0', 'admin', '', '是否成功失败', '2021-12-17 10:10:03+00', '2021-12-17 10:10:03+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (27, '操作分类', 'sys_oper_type', '0', 'admin', '', '操作分类列表', '2021-12-17 11:29:50+00', '2021-12-17 11:29:50+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (28, '任务组', 'sys_job_group', '0', 'panda', '', '系统任务，开机自启', '2021-12-24 15:14:56+00', '2021-12-24 15:14:56+00', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (29, '通知类型', 'sys_notice_type', '0', 'panda', '', '通知类型列表', '2021-12-26 15:23:26+00', '2021-12-26 15:23:26+00', NULL);

-- ----------------------------
-- Table structure for sys_jobs
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_jobs";
CREATE TABLE "public"."sys_jobs" (
  "job_id" int8 NOT NULL DEFAULT nextval('sys_jobs_job_id_seq'::regclass),
  "job_name" varchar(255) COLLATE "pg_catalog"."default",
  "job_group" varchar(255) COLLATE "pg_catalog"."default",
  "job_type" varchar(1) COLLATE "pg_catalog"."default",
  "cron_expression" varchar(255) COLLATE "pg_catalog"."default",
  "invoke_target" varchar(255) COLLATE "pg_catalog"."default",
  "args" varchar(255) COLLATE "pg_catalog"."default",
  "misfire_policy" varchar(1) COLLATE "pg_catalog"."default",
  "concurrent" varchar(1) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "entry_id" int8,
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_jobs"."create_by" IS '创建人';
COMMENT ON COLUMN "public"."sys_jobs"."update_by" IS '更新者';

-- ----------------------------
-- Records of sys_jobs
-- ----------------------------
INSERT INTO "public"."sys_jobs" VALUES (1, 'testcron', 'SYSTEM', '2', ' 0/10 * * * * ?', 'cronHandle', 'aaa', '', '', '1', 0, 'panda', '', '2021-12-24 16:06:09+00', '2021-12-24 22:44:09+00', NULL);

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_menus";
CREATE TABLE "public"."sys_menus" (
  "menu_id" int8 NOT NULL DEFAULT nextval('sys_menus_menu_id_seq'::regclass),
  "menu_name" varchar(128) COLLATE "pg_catalog"."default",
  "title" varchar(64) COLLATE "pg_catalog"."default",
  "parent_id" int8,
  "sort" int8,
  "icon" varchar(128) COLLATE "pg_catalog"."default",
  "path" varchar(128) COLLATE "pg_catalog"."default",
  "component" varchar(255) COLLATE "pg_catalog"."default",
  "is_iframe" varchar(1) COLLATE "pg_catalog"."default",
  "is_link" varchar(255) COLLATE "pg_catalog"."default",
  "menu_type" varchar(1) COLLATE "pg_catalog"."default",
  "is_hide" varchar(1) COLLATE "pg_catalog"."default",
  "is_keep_alive" varchar(1) COLLATE "pg_catalog"."default",
  "is_affix" varchar(1) COLLATE "pg_catalog"."default",
  "permission" varchar(32) COLLATE "pg_catalog"."default",
  "status" text COLLATE "pg_catalog"."default",
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "remark" text COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO "public"."sys_menus" VALUES (1, '系统设置', '', 0, 0, 'elementSetting', '/system', 'Layout', '1', '', 'M', '0', '0', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08+00', '2021-12-28 13:32:21+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (3, '用户管理', '', 1, 1, 'elementUser', '/system/user', '/system/user/index', '1', '', 'C', '0', '1', '1', 'system:user:list', '0', 'admin', 'panda', '', '2021-12-02 14:07:56+00', '2021-12-28 13:32:44+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (4, '添加用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:add', '0', 'admin', '', '', '2021-12-03 13:36:33+00', '2021-12-03 13:36:33+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (5, '编辑用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:edit', '0', 'admin', '', '', '2021-12-03 13:48:13+00', '2021-12-03 13:48:13+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (6, '角色管理', '', 1, 1, 'elementUserFilled', '/system/role', '/system/role/index', '1', '', 'C', '0', '1', '1', 'system:role:list', '0', 'admin', 'panda', '', '2021-12-03 13:51:55+00', '2021-12-28 13:32:55+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (7, '菜单管理', '', 1, 2, 'iconfont icon-juxingkaobei', '/system/menu', '/system/menu/index', '1', '', 'C', '0', '1', '1', 'system:menu:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44+00', '2021-12-28 13:33:19+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (8, '部门管理', '', 1, 3, 'iconfont icon-jiliandongxuanzeqi', '/system/dept', '/system/dept/index', '1', '', 'C', '0', '1', '1', 'system:dept:list', '0', 'admin', 'panda', '', '2021-12-03 13:58:36+00', '2021-12-28 13:40:20+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (9, '岗位管理', '', 1, 4, 'iconfont icon-neiqianshujuchucun', '/system/post', '/system/post/index', '1', '', 'C', '0', '1', '1', 'system:post:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44+00', '2021-12-28 13:40:31+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (10, '字典管理', '', 1, 5, 'elementCellphone', '/system/dict', '/system/dict/index', '1', '', 'C', '0', '1', '1', 'system:dict:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44+00', '2021-12-28 13:40:50+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (11, '参数管理', '', 1, 6, 'elementDocumentCopy', '/system/config', '/system/config/index', '1', '', 'C', '0', '1', '1', 'system:config:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44+00', '2021-12-28 13:41:05+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (12, '个人中心', '', 0, 10, 'elementAvatar', '/personal', '/personal/index', '1', '', 'M', '0', '0', '0', '', '0', 'admin', 'panda', '', '2021-12-03 14:12:43+00', '2021-12-28 13:43:17+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (13, '添加配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:add', '0', 'admin', '', '', '2021-12-06 17:19:19+00', '2021-12-06 17:19:19+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (14, '修改配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:edit', '0', 'admin', '', '', '2021-12-06 17:20:30+00', '2021-12-06 17:20:30+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (15, '删除配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:delete', '0', 'admin', '', '', '2021-12-06 17:23:52+00', '2021-12-06 17:23:52+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (16, '导出配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:export', '0', 'admin', '', '', '2021-12-06 17:24:41+00', '2021-12-06 17:24:41+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (17, '新增角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:add', '0', 'admin', '', '', '2021-12-06 17:43:35+00', '2021-12-06 17:43:35+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (18, '删除角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:delete', '0', 'admin', '', '', '2021-12-06 17:44:10+00', '2021-12-06 17:44:10+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (19, '修改角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:edit', '0', 'admin', '', '', '2021-12-06 17:44:48+00', '2021-12-06 17:44:48+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (20, '导出角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:export', '0', 'admin', '', '', '2021-12-06 17:45:25+00', '2021-12-06 17:45:25+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (21, '添加菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:add', '0', 'admin', '', '', '2021-12-06 17:46:01+00', '2021-12-06 17:46:01+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (22, '修改菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:edit', '0', 'admin', '', '', '2021-12-06 17:46:24+00', '2021-12-06 17:46:24+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (23, '删除菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:delete', '0', 'admin', '', '', '2021-12-06 17:46:47+00', '2021-12-06 17:46:47+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (24, '添加部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:add', '0', 'admin', '', '', '2021-12-07 09:33:58+00', '2021-12-07 09:33:58+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (25, '编辑部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:edit', '0', 'admin', '', '', '2021-12-07 09:34:39+00', '2021-12-07 09:34:39+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (26, '删除部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:delete', '0', 'admin', 'admin', '', '2021-12-07 09:35:09+00', '2021-12-07 09:36:26+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (27, '导出部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:export', '0', 'admin', '', '', '2021-12-07 09:35:51+00', '2021-12-07 09:35:51+00', '2021-12-07 09:36:37+00');
INSERT INTO "public"."sys_menus" VALUES (28, '添加岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:add', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (29, '编辑岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:edit', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (30, '删除岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:delete', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (31, '导出岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:export', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (32, '添加字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:add', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (33, '编辑字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:edit', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (34, '删除字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:delete', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (35, '导出字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:export', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (36, '新增字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:add', '0', 'admin', '', '', '2021-12-07 09:35:09+00', '2021-12-07 09:35:09+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (37, '修改字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:edit', '0', 'admin', '', '', '2021-12-07 09:48:04+00', '2021-12-07 09:48:04+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (38, '删除字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:delete', '0', 'admin', '', '', '2021-12-07 09:48:42+00', '2021-12-07 09:48:42+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (39, 'API管理', '', 1, 1, 'iconfont icon-siweidaotu', '/system/api', '/system/api/index', '1', '', 'C', '0', '1', '1', 'system:api:list', '0', 'admin', '', '', '2021-12-09 09:09:13+00', '2021-12-09 09:09:13+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (40, '添加api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:add', '0', 'admin', '', '', '2021-12-09 09:09:54+00', '2021-12-09 09:09:54+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (41, '编辑api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:edit', '0', 'admin', '', '', '2021-12-09 09:10:38+00', '2021-12-09 09:10:38+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (42, '删除api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:delete', '0', 'admin', '', '', '2021-12-09 09:11:11+00', '2021-12-09 09:11:11+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (43, '日志系统', '', 0, 1, 'iconfont icon-biaodan', '/log', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08+00', '2021-12-28 13:38:33+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (44, '系统工具', '', 0, 2, 'iconfont icon-gongju', '/tool', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-16 16:35:15+00', '2021-12-28 13:38:46+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (45, '操作日志', '', 43, 1, 'iconfont icon-bolangnengshiyanchang', '/log/operation', '/log/operation/index', '1', '', 'C', '0', '1', '1', 'log:operation:list', '0', 'admin', 'panda', '', '2021-12-16 16:42:03+00', '2021-12-28 13:39:44+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (46, '登录日志', '', 43, 2, 'iconfont icon--chaifenlie', '/log/login', '/log/login/index', '1', '', 'C', '0', '1', '1', 'log:login:list', '0', 'admin', 'panda', '', '2021-12-16 16:43:28+00', '2021-12-28 13:39:58+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (47, '服务监控', '', 44, 1, 'elementCpu', '/tool/monitor/', '/tool/monitor/index', '1', '', 'C', '0', '1', '1', 'tool:monitor:list', '0', 'admin', 'panda', '', '2021-12-03 14:12:43+00', '2021-12-28 13:41:25+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (48, '定时任务', '', 44, 2, 'elementAlarmClock', '/tool/job', '/tool/job/index', '1', '', 'C', '0', '1', '1', 'tool:job:list', '0', 'admin', 'panda', '', '2021-12-16 16:48:45+00', '2021-12-28 13:41:59+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (49, '开发工具', '', 0, 3, 'iconfont icon-diannao', '/develop', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', '', '', '2021-12-16 16:53:11+00', '2021-12-16 16:53:11+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (50, '表单构建', '', 49, 1, 'iconfont icon-zidingyibuju', '/develop/form', '/layout/routerView/iframes', '0', 'http://127.0.0.1:7788/form-generator/', 'C', '0', '1', '1', 'develop:form:list', '0', 'admin', 'panda', '', '2021-12-16 16:55:01+00', '2021-12-31 08:23:07+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (51, '代码生成', '', 49, 2, 'iconfont icon-zhongduancanshu', '/develop/code', '/develop/code/index', '1', '', 'C', '0', '1', '1', 'develop:code:list', '0', 'admin', '', '', '2021-12-16 16:56:48+00', '2021-12-16 16:56:48+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (52, '系统接口', '', 49, 3, 'iconfont icon-wenducanshu-05', '/develop/apis', '/layout/routerView/iframes', '0', 'http://127.0.0.1:7788/swagger/index.html', 'C', '0', '1', '1', 'develop:apis:list', '0', 'admin', 'panda', '', '2021-12-16 16:58:07+00', '2021-12-25 00:09:23+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (53, '资源管理', '', 0, 4, 'iconfont icon-juxingkaobei', '/resource', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', '', '', '2021-12-16 17:02:06+00', '2021-12-16 17:02:06+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (54, '文件管理', '', 53, 1, 'iconfont icon-chazhaobiaodanliebiao', '/resource/file', '/resource/file/index', '1', '', 'C', '0', '1', '1', 'resource:file:list', '0', 'admin', 'panda', '', '2021-12-16 17:06:04+00', '2021-12-28 15:13:08+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (55, '公告通知', '', 44, 3, 'elementTicket', '/tool/notice', '/tool/notice/index', '1', '', 'C', '0', '1', '1', 'tool:notice:list', '0', 'admin', 'panda', '', '2021-12-16 22:09:11+00', '2021-12-28 13:42:39+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (56, '任务日志', '', 43, 1, 'iconfont icon--chaifenhang', '/log/job', '/log/job/index', '1', '', 'C', '0', '1', '1', 'log:job:list', '0', 'panda', 'panda', '', '2021-12-24 22:13:45+00', '2021-12-28 13:39:52+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (57, '邮件管理', '', 53, 1, 'iconfont icon-shouye_dongtaihui', '/resource/mail', '/resource/mail/index', '1', '', 'C', '0', '1', '1', 'resource:mail:list', '0', 'panda', 'panda', '', '2021-12-28 15:12:37+00', '2021-12-28 15:12:45+00', NULL);
INSERT INTO "public"."sys_menus" VALUES (58, '短信管理', '', 53, 3, 'elementChatDotRound', '/resource/message', '/resource/message/index', '1', '', 'C', '0', '1', '1', 'resource:message:list', '0', 'panda', '', '', '2021-12-16 17:06:04+00', '2021-12-16 17:06:04+00', NULL);

-- ----------------------------
-- Table structure for sys_notices
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_notices";
CREATE TABLE "public"."sys_notices" (
  "notice_id" int8 NOT NULL DEFAULT nextval('sys_notices_notice_id_seq'::regclass),
  "title" varchar(128) COLLATE "pg_catalog"."default",
  "content" text COLLATE "pg_catalog"."default",
  "notice_type" varchar(1) COLLATE "pg_catalog"."default",
  "dept_id" int8,
  "user_name" varchar(64) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_notices"."title" IS '标题';
COMMENT ON COLUMN "public"."sys_notices"."content" IS '标题';
COMMENT ON COLUMN "public"."sys_notices"."notice_type" IS '通知类型';
COMMENT ON COLUMN "public"."sys_notices"."dept_id" IS '部门Id,部门及子部门';
COMMENT ON COLUMN "public"."sys_notices"."user_name" IS '发布人';

-- ----------------------------
-- Records of sys_notices
-- ----------------------------
INSERT INTO "public"."sys_notices" VALUES (1, '关于学习交流的通知', '<p>发布<b>入群</b>通知&nbsp;<span style=\"color: var(--el-text-color-regular);\">467890197, 交流学习</span><span style=\"font-size: 14px; color: var(--el-text-color-regular);\"></span></p>', '1', 0, 'panda', '2021-12-26 15:29:25+00', '2021-12-26 16:19:48+00', NULL);
INSERT INTO "public"."sys_notices" VALUES (2, 'test', '<p>sdsad</p>', '1', 2, 'panda', '2021-12-26 16:23:13+00', '2021-12-26 16:23:13+00', '2021-12-26 16:31:31+00');
INSERT INTO "public"."sys_notices" VALUES (3, '版本更新通知：任务功能，通知功能完成', '<p><span style=\"font-size: 14px; color: var(--el-text-color-regular);\">', '1', 0, 'panda', '2021-12-26 17:33:47+00', '2021-12-26 17:33:47+00', NULL);

-- ----------------------------
-- Table structure for sys_posts
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_posts";
CREATE TABLE "public"."sys_posts" (
  "post_id" int8 NOT NULL DEFAULT nextval('sys_posts_post_id_seq'::regclass),
  "post_name" varchar(128) COLLATE "pg_catalog"."default",
  "post_code" varchar(128) COLLATE "pg_catalog"."default",
  "sort" int8,
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_posts"."post_name" IS '岗位名称';
COMMENT ON COLUMN "public"."sys_posts"."post_code" IS '岗位代码';
COMMENT ON COLUMN "public"."sys_posts"."sort" IS '岗位排序';
COMMENT ON COLUMN "public"."sys_posts"."status" IS '状态';
COMMENT ON COLUMN "public"."sys_posts"."remark" IS '描述';

-- ----------------------------
-- Records of sys_posts
-- ----------------------------
INSERT INTO "public"."sys_posts" VALUES (1, '首席执行官', 'CEO', 0, '0', '首席执行官', 'admin', '', '2021-12-02 09:21:44+00', '2021-12-02 09:24:25+00', NULL);
INSERT INTO "public"."sys_posts" VALUES (3, '首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44+00', '2021-12-02 09:25:59+00', '2021-12-02 09:27:41+00');
INSERT INTO "public"."sys_posts" VALUES (4, '首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44+00', '2021-12-02 09:25:59+00', NULL);
INSERT INTO "public"."sys_posts" VALUES (5, '123', '123', 0, '0', '', 'admin', '', '2021-12-18 00:33:28+00', '2021-12-18 00:33:28+00', '2021-12-28 14:11:52+00');

-- ----------------------------
-- Table structure for sys_role_depts
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role_depts";
CREATE TABLE "public"."sys_role_depts" (
  "role_id" int8,
  "dept_id" int8,
  "id" int8 NOT NULL DEFAULT nextval('sys_role_depts_id_seq'::regclass)
)
;

-- ----------------------------
-- Records of sys_role_depts
-- ----------------------------
INSERT INTO "public"."sys_role_depts" VALUES (1, 2, 1);
INSERT INTO "public"."sys_role_depts" VALUES (1, 3, 2);

-- ----------------------------
-- Table structure for sys_role_menus
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role_menus";
CREATE TABLE "public"."sys_role_menus" (
  "role_id" int8,
  "menu_id" int8,
  "role_name" varchar(128) COLLATE "pg_catalog"."default",
  "id" int8 NOT NULL DEFAULT nextval('sys_role_menus_id_seq'::regclass)
)
;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO "public"."sys_role_menus" VALUES (1, 1, 'admin', 1);
INSERT INTO "public"."sys_role_menus" VALUES (1, 3, 'admin', 2);
INSERT INTO "public"."sys_role_menus" VALUES (1, 4, 'admin', 3);
INSERT INTO "public"."sys_role_menus" VALUES (1, 5, 'admin', 4);
INSERT INTO "public"."sys_role_menus" VALUES (1, 6, 'admin', 5);
INSERT INTO "public"."sys_role_menus" VALUES (1, 7, 'admin', 6);
INSERT INTO "public"."sys_role_menus" VALUES (1, 8, 'admin', 7);
INSERT INTO "public"."sys_role_menus" VALUES (1, 9, 'admin', 8);
INSERT INTO "public"."sys_role_menus" VALUES (1, 10, 'admin', 9);
INSERT INTO "public"."sys_role_menus" VALUES (1, 11, 'admin', 10);
INSERT INTO "public"."sys_role_menus" VALUES (1, 12, 'admin', 11);
INSERT INTO "public"."sys_role_menus" VALUES (1, 13, 'admin', 12);
INSERT INTO "public"."sys_role_menus" VALUES (1, 14, 'admin', 13);
INSERT INTO "public"."sys_role_menus" VALUES (1, 15, 'admin', 14);
INSERT INTO "public"."sys_role_menus" VALUES (1, 16, 'admin', 15);
INSERT INTO "public"."sys_role_menus" VALUES (1, 17, 'admin', 16);
INSERT INTO "public"."sys_role_menus" VALUES (1, 18, 'admin', 17);
INSERT INTO "public"."sys_role_menus" VALUES (1, 19, 'admin', 18);
INSERT INTO "public"."sys_role_menus" VALUES (1, 20, 'admin', 19);
INSERT INTO "public"."sys_role_menus" VALUES (1, 21, 'admin', 20);
INSERT INTO "public"."sys_role_menus" VALUES (1, 22, 'admin', 21);
INSERT INTO "public"."sys_role_menus" VALUES (1, 23, 'admin', 22);
INSERT INTO "public"."sys_role_menus" VALUES (1, 24, 'admin', 23);
INSERT INTO "public"."sys_role_menus" VALUES (1, 25, 'admin', 24);
INSERT INTO "public"."sys_role_menus" VALUES (1, 26, 'admin', 25);
INSERT INTO "public"."sys_role_menus" VALUES (1, 28, 'admin', 26);
INSERT INTO "public"."sys_role_menus" VALUES (1, 29, 'admin', 27);
INSERT INTO "public"."sys_role_menus" VALUES (1, 30, 'admin', 28);
INSERT INTO "public"."sys_role_menus" VALUES (1, 31, 'admin', 29);
INSERT INTO "public"."sys_role_menus" VALUES (1, 32, 'admin', 30);
INSERT INTO "public"."sys_role_menus" VALUES (1, 33, 'admin', 31);
INSERT INTO "public"."sys_role_menus" VALUES (1, 34, 'admin', 32);
INSERT INTO "public"."sys_role_menus" VALUES (1, 35, 'admin', 33);
INSERT INTO "public"."sys_role_menus" VALUES (1, 36, 'admin', 34);
INSERT INTO "public"."sys_role_menus" VALUES (1, 37, 'admin', 35);
INSERT INTO "public"."sys_role_menus" VALUES (1, 38, 'admin', 36);
INSERT INTO "public"."sys_role_menus" VALUES (1, 39, 'admin', 37);
INSERT INTO "public"."sys_role_menus" VALUES (1, 40, 'admin', 38);
INSERT INTO "public"."sys_role_menus" VALUES (1, 41, 'admin', 39);
INSERT INTO "public"."sys_role_menus" VALUES (1, 42, 'admin', 40);
INSERT INTO "public"."sys_role_menus" VALUES (1, 43, 'admin', 41);
INSERT INTO "public"."sys_role_menus" VALUES (1, 44, 'admin', 42);
INSERT INTO "public"."sys_role_menus" VALUES (1, 45, 'admin', 43);
INSERT INTO "public"."sys_role_menus" VALUES (1, 46, 'admin', 44);
INSERT INTO "public"."sys_role_menus" VALUES (1, 47, 'admin', 45);
INSERT INTO "public"."sys_role_menus" VALUES (1, 48, 'admin', 46);
INSERT INTO "public"."sys_role_menus" VALUES (1, 49, 'admin', 47);
INSERT INTO "public"."sys_role_menus" VALUES (1, 50, 'admin', 48);
INSERT INTO "public"."sys_role_menus" VALUES (1, 51, 'admin', 49);
INSERT INTO "public"."sys_role_menus" VALUES (1, 52, 'admin', 50);
INSERT INTO "public"."sys_role_menus" VALUES (1, 53, 'admin', 51);
INSERT INTO "public"."sys_role_menus" VALUES (1, 54, 'admin', 52);
INSERT INTO "public"."sys_role_menus" VALUES (1, 55, 'admin', 53);
INSERT INTO "public"."sys_role_menus" VALUES (1, 56, 'admin', 54);
INSERT INTO "public"."sys_role_menus" VALUES (1, 57, 'admin', 55);
INSERT INTO "public"."sys_role_menus" VALUES (1, 58, 'admin', 56);
INSERT INTO "public"."sys_role_menus" VALUES (2, 1, 'manage', 57);
INSERT INTO "public"."sys_role_menus" VALUES (2, 3, 'manage', 58);
INSERT INTO "public"."sys_role_menus" VALUES (2, 4, 'manage', 59);
INSERT INTO "public"."sys_role_menus" VALUES (2, 5, 'manage', 60);
INSERT INTO "public"."sys_role_menus" VALUES (2, 6, 'manage', 61);
INSERT INTO "public"."sys_role_menus" VALUES (2, 7, 'manage', 62);
INSERT INTO "public"."sys_role_menus" VALUES (2, 8, 'manage', 63);
INSERT INTO "public"."sys_role_menus" VALUES (2, 9, 'manage', 64);
INSERT INTO "public"."sys_role_menus" VALUES (2, 10, 'manage', 65);
INSERT INTO "public"."sys_role_menus" VALUES (2, 11, 'manage', 66);
INSERT INTO "public"."sys_role_menus" VALUES (2, 12, 'manage', 67);
INSERT INTO "public"."sys_role_menus" VALUES (2, 13, 'manage', 68);
INSERT INTO "public"."sys_role_menus" VALUES (2, 14, 'manage', 69);
INSERT INTO "public"."sys_role_menus" VALUES (2, 15, 'manage', 70);
INSERT INTO "public"."sys_role_menus" VALUES (2, 16, 'manage', 71);
INSERT INTO "public"."sys_role_menus" VALUES (2, 17, 'manage', 72);
INSERT INTO "public"."sys_role_menus" VALUES (2, 18, 'manage', 73);
INSERT INTO "public"."sys_role_menus" VALUES (2, 19, 'manage', 74);
INSERT INTO "public"."sys_role_menus" VALUES (2, 20, 'manage', 75);
INSERT INTO "public"."sys_role_menus" VALUES (2, 21, 'manage', 76);
INSERT INTO "public"."sys_role_menus" VALUES (2, 22, 'manage', 77);
INSERT INTO "public"."sys_role_menus" VALUES (2, 23, 'manage', 78);
INSERT INTO "public"."sys_role_menus" VALUES (2, 24, 'manage', 79);
INSERT INTO "public"."sys_role_menus" VALUES (2, 25, 'manage', 80);
INSERT INTO "public"."sys_role_menus" VALUES (2, 26, 'manage', 81);
INSERT INTO "public"."sys_role_menus" VALUES (2, 28, 'manage', 82);
INSERT INTO "public"."sys_role_menus" VALUES (2, 29, 'manage', 83);
INSERT INTO "public"."sys_role_menus" VALUES (2, 30, 'manage', 84);
INSERT INTO "public"."sys_role_menus" VALUES (2, 31, 'manage', 85);
INSERT INTO "public"."sys_role_menus" VALUES (2, 32, 'manage', 86);
INSERT INTO "public"."sys_role_menus" VALUES (2, 33, 'manage', 87);
INSERT INTO "public"."sys_role_menus" VALUES (2, 34, 'manage', 88);
INSERT INTO "public"."sys_role_menus" VALUES (2, 35, 'manage', 89);
INSERT INTO "public"."sys_role_menus" VALUES (2, 36, 'manage', 90);
INSERT INTO "public"."sys_role_menus" VALUES (2, 37, 'manage', 91);
INSERT INTO "public"."sys_role_menus" VALUES (2, 38, 'manage', 92);
INSERT INTO "public"."sys_role_menus" VALUES (2, 39, 'manage', 93);
INSERT INTO "public"."sys_role_menus" VALUES (2, 40, 'manage', 94);
INSERT INTO "public"."sys_role_menus" VALUES (2, 41, 'manage', 95);
INSERT INTO "public"."sys_role_menus" VALUES (2, 42, 'manage', 96);
INSERT INTO "public"."sys_role_menus" VALUES (2, 43, 'manage', 97);
INSERT INTO "public"."sys_role_menus" VALUES (2, 44, 'manage', 98);
INSERT INTO "public"."sys_role_menus" VALUES (2, 45, 'manage', 99);
INSERT INTO "public"."sys_role_menus" VALUES (2, 46, 'manage', 100);
INSERT INTO "public"."sys_role_menus" VALUES (2, 47, 'manage', 101);
INSERT INTO "public"."sys_role_menus" VALUES (2, 48, 'manage', 102);
INSERT INTO "public"."sys_role_menus" VALUES (2, 49, 'manage', 103);
INSERT INTO "public"."sys_role_menus" VALUES (2, 50, 'manage', 104);
INSERT INTO "public"."sys_role_menus" VALUES (2, 51, 'manage', 105);
INSERT INTO "public"."sys_role_menus" VALUES (2, 52, 'manage', 106);
INSERT INTO "public"."sys_role_menus" VALUES (2, 53, 'manage', 107);
INSERT INTO "public"."sys_role_menus" VALUES (2, 54, 'manage', 108);
INSERT INTO "public"."sys_role_menus" VALUES (2, 55, 'manage', 109);
INSERT INTO "public"."sys_role_menus" VALUES (2, 57, 'manage', 110);
INSERT INTO "public"."sys_role_menus" VALUES (2, 58, 'manage', 111);

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_roles";
CREATE TABLE "public"."sys_roles" (
  "role_id" int8 NOT NULL DEFAULT nextval('sys_roles_role_id_seq'::regclass),
  "role_name" varchar(128) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "role_key" varchar(128) COLLATE "pg_catalog"."default",
  "role_sort" int8,
  "data_scope" varchar(1) COLLATE "pg_catalog"."default",
  "flag" varchar(128) COLLATE "pg_catalog"."default",
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_roles"."role_name" IS '角色名称';
COMMENT ON COLUMN "public"."sys_roles"."status" IS '状态';
COMMENT ON COLUMN "public"."sys_roles"."role_key" IS '角色代码';
COMMENT ON COLUMN "public"."sys_roles"."role_sort" IS '角色排序';
COMMENT ON COLUMN "public"."sys_roles"."data_scope" IS '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）';
COMMENT ON COLUMN "public"."sys_roles"."flag" IS '删除标识';

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO "public"."sys_roles" VALUES (1, '超管理员', '0', 'admin', 2, '1', '', 'admin', 'panda', '', '2021-12-02 16:03:26+00', '2021-12-28 15:16:11+00', NULL);
INSERT INTO "public"."sys_roles" VALUES (2, '管理员', '0', 'manage', 2, '', '', 'panda', 'panda', '', '2021-12-19 16:06:20+00', '2021-12-28 15:16:23+00', NULL);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_users";
CREATE TABLE "public"."sys_users" (
  "user_id" int8 NOT NULL DEFAULT nextval('sys_users_user_id_seq'::regclass),
  "nick_name" varchar(128) COLLATE "pg_catalog"."default",
  "phone" varchar(11) COLLATE "pg_catalog"."default",
  "role_id" int8,
  "salt" varchar(255) COLLATE "pg_catalog"."default",
  "avatar" varchar(255) COLLATE "pg_catalog"."default",
  "sex" varchar(255) COLLATE "pg_catalog"."default",
  "email" varchar(128) COLLATE "pg_catalog"."default",
  "dept_id" int8,
  "post_id" int8,
  "role_ids" varchar(255) COLLATE "pg_catalog"."default",
  "post_ids" varchar(255) COLLATE "pg_catalog"."default",
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6),
  "username" varchar(64) COLLATE "pg_catalog"."default",
  "password" varchar(128) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO "public"."sys_users" VALUES (1, 'pandax', '13818888888', 1, NULL, '', '0', '1@qq.com', 2, 1, '1', '1', 'admin', '1', NULL, '0', '2021-12-03 09:46:55+00', '2021-12-03 10:51:34+00', NULL, 'panda', '$2a$10$EXMJ5huCCTlCmP2ITFkAJ.4Mgmq3JcZGUvtE.KLX8j7FmhiiTEEya');
INSERT INTO "public"."sys_users" VALUES (4, 'panda', '18353366912', 2, '', '', '0', '2417920382@qq.com', 2, 4, '2', '4,1', 'panda', '', '', '0', '2021-12-19 15:58:09+00', '2021-12-19 16:06:54+00', NULL, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2');

-- ----------------------------
-- Indexes structure for table casbin_rule
-- ----------------------------
CREATE UNIQUE INDEX "idx_casbin_rule" ON "public"."casbin_rule" USING btree (
  "ptype" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v0" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v1" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v2" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v3" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v4" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v5" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table casbin_rule
-- ----------------------------
ALTER TABLE "public"."casbin_rule" ADD CONSTRAINT "casbin_rule_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table log_jobs
-- ----------------------------
ALTER TABLE "public"."log_jobs" ADD CONSTRAINT "log_jobs_pkey" PRIMARY KEY ("log_id");

-- ----------------------------
-- Primary Key structure for table log_logins
-- ----------------------------
ALTER TABLE "public"."log_logins" ADD CONSTRAINT "log_logins_pkey" PRIMARY KEY ("info_id");

-- ----------------------------
-- Primary Key structure for table log_opers
-- ----------------------------
ALTER TABLE "public"."log_opers" ADD CONSTRAINT "log_opers_pkey" PRIMARY KEY ("oper_id");

-- ----------------------------
-- Primary Key structure for table sys_apis
-- ----------------------------
ALTER TABLE "public"."sys_apis" ADD CONSTRAINT "sys_apis_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_configs
-- ----------------------------
ALTER TABLE "public"."sys_configs" ADD CONSTRAINT "sys_configs_pkey" PRIMARY KEY ("config_id");

-- ----------------------------
-- Primary Key structure for table sys_depts
-- ----------------------------
ALTER TABLE "public"."sys_depts" ADD CONSTRAINT "sys_depts_pkey" PRIMARY KEY ("dept_id");

-- ----------------------------
-- Primary Key structure for table sys_dict_data
-- ----------------------------
ALTER TABLE "public"."sys_dict_data" ADD CONSTRAINT "sys_dict_data_pkey" PRIMARY KEY ("dict_code");

-- ----------------------------
-- Primary Key structure for table sys_dict_types
-- ----------------------------
ALTER TABLE "public"."sys_dict_types" ADD CONSTRAINT "sys_dict_types_pkey" PRIMARY KEY ("dict_id");

-- ----------------------------
-- Primary Key structure for table sys_jobs
-- ----------------------------
ALTER TABLE "public"."sys_jobs" ADD CONSTRAINT "sys_jobs_pkey" PRIMARY KEY ("job_id");

-- ----------------------------
-- Primary Key structure for table sys_menus
-- ----------------------------
ALTER TABLE "public"."sys_menus" ADD CONSTRAINT "sys_menus_pkey" PRIMARY KEY ("menu_id");

-- ----------------------------
-- Primary Key structure for table sys_notices
-- ----------------------------
ALTER TABLE "public"."sys_notices" ADD CONSTRAINT "sys_notices_pkey" PRIMARY KEY ("notice_id");

-- ----------------------------
-- Primary Key structure for table sys_posts
-- ----------------------------
ALTER TABLE "public"."sys_posts" ADD CONSTRAINT "sys_posts_pkey" PRIMARY KEY ("post_id");

-- ----------------------------
-- Primary Key structure for table sys_role_depts
-- ----------------------------
ALTER TABLE "public"."sys_role_depts" ADD CONSTRAINT "sys_role_depts_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_role_menus
-- ----------------------------
ALTER TABLE "public"."sys_role_menus" ADD CONSTRAINT "sys_role_menus_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_roles
-- ----------------------------
ALTER TABLE "public"."sys_roles" ADD CONSTRAINT "sys_roles_pkey" PRIMARY KEY ("role_id");

-- ----------------------------
-- Primary Key structure for table sys_users
-- ----------------------------
ALTER TABLE "public"."sys_users" ADD CONSTRAINT "sys_users_pkey" PRIMARY KEY ("user_id");
