/*
 Navicat Premium Data Transfer

 Source Server Type    : PostgreSQL
 Source Server Version : 90204
 Source Host           : 10.100.1.1:5432
 Source Catalog        : postgres
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90204
 File Encoding         : 65001

 Date: 23/10/2023 08:40:14
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
  "id" int4 NOT NULL
)
;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------

-- ----------------------------
-- Table structure for dev_gen_table_columns
-- ----------------------------
DROP TABLE IF EXISTS "public"."dev_gen_table_columns";
CREATE TABLE "public"."dev_gen_table_columns" (
  "column_id" int8 NOT NULL,
  "table_id" int8,
  "table_name" varchar(191) COLLATE "pg_catalog"."default",
  "column_name" varchar(191) COLLATE "pg_catalog"."default",
  "column_comment" varchar(191) COLLATE "pg_catalog"."default",
  "column_type" varchar(191) COLLATE "pg_catalog"."default",
  "column_key" varchar(191) COLLATE "pg_catalog"."default",
  "go_type" varchar(191) COLLATE "pg_catalog"."default",
  "go_field" varchar(191) COLLATE "pg_catalog"."default",
  "json_field" varchar(191) COLLATE "pg_catalog"."default",
  "html_field" varchar(191) COLLATE "pg_catalog"."default",
  "is_pk" varchar(191) COLLATE "pg_catalog"."default",
  "is_increment" varchar(191) COLLATE "pg_catalog"."default",
  "is_required" varchar(191) COLLATE "pg_catalog"."default",
  "is_insert" varchar(191) COLLATE "pg_catalog"."default",
  "is_edit" varchar(191) COLLATE "pg_catalog"."default",
  "is_list" varchar(191) COLLATE "pg_catalog"."default",
  "is_query" varchar(191) COLLATE "pg_catalog"."default",
  "query_type" varchar(191) COLLATE "pg_catalog"."default",
  "html_type" varchar(191) COLLATE "pg_catalog"."default",
  "dict_type" varchar(191) COLLATE "pg_catalog"."default",
  "sort" int8,
  "link_table_name" varchar(191) COLLATE "pg_catalog"."default",
  "link_table_class" varchar(191) COLLATE "pg_catalog"."default",
  "link_table_package" varchar(191) COLLATE "pg_catalog"."default",
  "link_label_id" varchar(191) COLLATE "pg_catalog"."default",
  "link_label_name" varchar(191) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "owner" varchar(64) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."dev_gen_table_columns"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."dev_gen_table_columns"."owner" IS '创建者,所有者';

-- ----------------------------
-- Records of dev_gen_table_columns
-- ----------------------------

-- ----------------------------
-- Table structure for dev_gen_tables
-- ----------------------------
DROP TABLE IF EXISTS "public"."dev_gen_tables";
CREATE TABLE "public"."dev_gen_tables" (
  "table_id" int8 NOT NULL,
  "table_name" varchar(191) COLLATE "pg_catalog"."default",
  "table_comment" varchar(191) COLLATE "pg_catalog"."default",
  "class_name" varchar(191) COLLATE "pg_catalog"."default",
  "tpl_category" varchar(191) COLLATE "pg_catalog"."default",
  "package_name" varchar(191) COLLATE "pg_catalog"."default",
  "module_name" varchar(191) COLLATE "pg_catalog"."default",
  "business_name" varchar(191) COLLATE "pg_catalog"."default",
  "function_name" varchar(191) COLLATE "pg_catalog"."default",
  "function_author" varchar(191) COLLATE "pg_catalog"."default",
  "options" varchar(191) COLLATE "pg_catalog"."default",
  "remark" varchar(191) COLLATE "pg_catalog"."default",
  "pk_column" varchar(191) COLLATE "pg_catalog"."default",
  "pk_go_field" varchar(191) COLLATE "pg_catalog"."default",
  "pk_json_field" varchar(191) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6),
  "org_id" int4,
  "owner" varchar(64) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."dev_gen_tables"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."dev_gen_tables"."owner" IS '创建者,所有者';

-- ----------------------------
-- Records of dev_gen_tables
-- ----------------------------

-- ----------------------------
-- Table structure for device_alarms
-- ----------------------------
DROP TABLE IF EXISTS "public"."device_alarms";
CREATE TABLE "public"."device_alarms" (
  "id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "device_id" varchar(64) COLLATE "pg_catalog"."default",
  "product_id" varchar(64) COLLATE "pg_catalog"."default",
  "type" varchar(255) COLLATE "pg_catalog"."default",
  "level" varchar(64) COLLATE "pg_catalog"."default",
  "state" varchar(64) COLLATE "pg_catalog"."default",
  "details" varchar(1024) COLLATE "pg_catalog"."default",
  "time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."device_alarms"."name" IS '告警名称';
COMMENT ON COLUMN "public"."device_alarms"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."device_alarms"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."device_alarms"."type" IS '告警类型';
COMMENT ON COLUMN "public"."device_alarms"."level" IS '告警级别';
COMMENT ON COLUMN "public"."device_alarms"."state" IS '告警状态';
COMMENT ON COLUMN "public"."device_alarms"."details" IS '详情';
COMMENT ON COLUMN "public"."device_alarms"."time" IS '告警时间';

-- ----------------------------
-- Records of device_alarms
-- ----------------------------

-- ----------------------------
-- Table structure for device_cmd_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."device_cmd_log";
CREATE TABLE "public"."device_cmd_log" (
  "id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "device_id" varchar(64) COLLATE "pg_catalog"."default",
  "cmd_name" varchar(64) COLLATE "pg_catalog"."default",
  "cmd_content" text COLLATE "pg_catalog"."default",
  "state" varchar(1) COLLATE "pg_catalog"."default",
  "response_content" text COLLATE "pg_catalog"."default",
  "request_time" varchar(64) COLLATE "pg_catalog"."default",
  "response_time" varchar(64) COLLATE "pg_catalog"."default",
  "type" varchar(64) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of device_cmd_log
-- ----------------------------

-- ----------------------------
-- Table structure for device_groups
-- ----------------------------
DROP TABLE IF EXISTS "public"."device_groups";
CREATE TABLE "public"."device_groups" (
  "id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "name" varchar(128) COLLATE "pg_catalog"."default",
  "pid" varchar(64) COLLATE "pg_catalog"."default",
  "path" varchar(255) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "sort" int4,
  "ext" text COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."device_groups"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."device_groups"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."device_groups"."name" IS '设备分组名称';
COMMENT ON COLUMN "public"."device_groups"."pid" IS '上级设备分组类型';
COMMENT ON COLUMN "public"."device_groups"."path" IS '设备分组路径';
COMMENT ON COLUMN "public"."device_groups"."description" IS '设备分组说明';
COMMENT ON COLUMN "public"."device_groups"."sort" IS '排序';
COMMENT ON COLUMN "public"."device_groups"."ext" IS '扩展';
COMMENT ON COLUMN "public"."device_groups"."status" IS '状态';

-- ----------------------------
-- Records of device_groups
-- ----------------------------

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS "public"."devices";
CREATE TABLE "public"."devices" (
  "id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "name" varchar(128) COLLATE "pg_catalog"."default",
  "token" varchar(128) COLLATE "pg_catalog"."default",
  "alias" varchar(128) COLLATE "pg_catalog"."default",
  "pid" varchar(191) COLLATE "pg_catalog"."default",
  "gid" varchar(191) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "ota_version" varchar(64) COLLATE "pg_catalog"."default",
  "ext" text COLLATE "pg_catalog"."default",
  "parent_id" varchar(64) COLLATE "pg_catalog"."default",
  "device_type" varchar(64) COLLATE "pg_catalog"."default",
  "link_status" varchar(10) COLLATE "pg_catalog"."default",
  "last_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."devices"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."devices"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."devices"."name" IS '设备名称';
COMMENT ON COLUMN "public"."devices"."token" IS '设备token';
COMMENT ON COLUMN "public"."devices"."alias" IS '设备别名';
COMMENT ON COLUMN "public"."devices"."pid" IS '产品Id';
COMMENT ON COLUMN "public"."devices"."gid" IS '分组Id';
COMMENT ON COLUMN "public"."devices"."description" IS '产品说明';
COMMENT ON COLUMN "public"."devices"."status" IS '状态';
COMMENT ON COLUMN "public"."devices"."ota_version" IS '固件版本';
COMMENT ON COLUMN "public"."devices"."ext" IS '拓展';
COMMENT ON COLUMN "public"."devices"."parent_id" IS '父Id，子设备时，父设备为网关';
COMMENT ON COLUMN "public"."devices"."device_type" IS '设备类型';
COMMENT ON COLUMN "public"."devices"."link_status" IS '连接状态';
COMMENT ON COLUMN "public"."devices"."last_time" IS '最后一次在线时间';

-- ----------------------------
-- Records of devices
-- ----------------------------

-- ----------------------------
-- Table structure for job_logs
-- ----------------------------
DROP TABLE IF EXISTS "public"."job_logs";
CREATE TABLE "public"."job_logs" (
  "id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "org_id" varchar(64) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "entry_id" int4,
  "target_invoke" varchar(64) COLLATE "pg_catalog"."default",
  "log_info" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."job_logs"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."job_logs"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."job_logs"."name" IS '任务名称';
COMMENT ON COLUMN "public"."job_logs"."entry_id" IS '任务id';
COMMENT ON COLUMN "public"."job_logs"."target_invoke" IS '调用方法';
COMMENT ON COLUMN "public"."job_logs"."log_info" IS '日志信息';
COMMENT ON COLUMN "public"."job_logs"."status" IS '状态';

-- ----------------------------
-- Records of job_logs
-- ----------------------------

-- ----------------------------
-- Table structure for jobs
-- ----------------------------
DROP TABLE IF EXISTS "public"."jobs";
CREATE TABLE "public"."jobs" (
  "id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "job_name" varchar(64) COLLATE "pg_catalog"."default",
  "target_invoke" varchar(64) COLLATE "pg_catalog"."default",
  "target_args" varchar(64) COLLATE "pg_catalog"."default",
  "job_content" varchar(255) COLLATE "pg_catalog"."default",
  "cron_expression" varchar(64) COLLATE "pg_catalog"."default",
  "misfire_policy" varchar(1) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "entry_id" int4
)
;
COMMENT ON COLUMN "public"."jobs"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."jobs"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."jobs"."job_name" IS '名称';
COMMENT ON COLUMN "public"."jobs"."target_invoke" IS '调用目标';
COMMENT ON COLUMN "public"."jobs"."target_args" IS '目标传参';
COMMENT ON COLUMN "public"."jobs"."job_content" IS '目标传参 要执行的内容';
COMMENT ON COLUMN "public"."jobs"."cron_expression" IS 'cron表达式';
COMMENT ON COLUMN "public"."jobs"."misfire_policy" IS '执行策略';
COMMENT ON COLUMN "public"."jobs"."status" IS '状态';
COMMENT ON COLUMN "public"."jobs"."entry_id" IS 'job启动时返回的id';

-- ----------------------------
-- Records of jobs
-- ----------------------------

-- ----------------------------
-- Table structure for log_logins
-- ----------------------------
DROP TABLE IF EXISTS "public"."log_logins";
CREATE TABLE "public"."log_logins" (
  "info_id" int8 NOT NULL,
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
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
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

-- ----------------------------
-- Table structure for log_opers
-- ----------------------------
DROP TABLE IF EXISTS "public"."log_opers";
CREATE TABLE "public"."log_opers" (
  "oper_id" int8 NOT NULL,
  "title" varchar(128) COLLATE "pg_catalog"."default",
  "business_type" varchar(1) COLLATE "pg_catalog"."default",
  "method" varchar(255) COLLATE "pg_catalog"."default",
  "oper_name" varchar(255) COLLATE "pg_catalog"."default",
  "oper_url" varchar(255) COLLATE "pg_catalog"."default",
  "oper_ip" varchar(255) COLLATE "pg_catalog"."default",
  "oper_location" varchar(255) COLLATE "pg_catalog"."default",
  "oper_param" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
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

-- ----------------------------
-- Table structure for product_categories
-- ----------------------------
DROP TABLE IF EXISTS "public"."product_categories";
CREATE TABLE "public"."product_categories" (
  "id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "name" varchar(128) COLLATE "pg_catalog"."default",
  "pid" varchar(64) COLLATE "pg_catalog"."default",
  "path" varchar(255) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "sort" int4,
  "status" varchar(1) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."product_categories"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."product_categories"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."product_categories"."name" IS '产品类型名称';
COMMENT ON COLUMN "public"."product_categories"."pid" IS '上级产品类型';
COMMENT ON COLUMN "public"."product_categories"."path" IS '产品类型路径';
COMMENT ON COLUMN "public"."product_categories"."description" IS '产品类型说明';
COMMENT ON COLUMN "public"."product_categories"."sort" IS '排序';
COMMENT ON COLUMN "public"."product_categories"."status" IS '状态';

-- ----------------------------
-- Records of product_categories
-- ----------------------------

-- ----------------------------
-- Table structure for product_ota
-- ----------------------------
DROP TABLE IF EXISTS "public"."product_ota";
CREATE TABLE "public"."product_ota" (
  "id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "pid" varchar(191) COLLATE "pg_catalog"."default",
  "is_latest" int2,
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "version" varchar(64) COLLATE "pg_catalog"."default",
  "url" varchar(128) COLLATE "pg_catalog"."default",
  "check" varchar(255) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "org_id" int4
)
;
COMMENT ON COLUMN "public"."product_ota"."pid" IS '产品Id';
COMMENT ON COLUMN "public"."product_ota"."is_latest" IS '最新版本';
COMMENT ON COLUMN "public"."product_ota"."name" IS '固件名称';
COMMENT ON COLUMN "public"."product_ota"."version" IS '固件版本';
COMMENT ON COLUMN "public"."product_ota"."url" IS '下载地址';
COMMENT ON COLUMN "public"."product_ota"."check" IS 'md5校验值';
COMMENT ON COLUMN "public"."product_ota"."description" IS '说明';
COMMENT ON COLUMN "public"."product_ota"."org_id" IS '机构ID';

-- ----------------------------
-- Records of product_ota
-- ----------------------------

-- ----------------------------
-- Table structure for product_templates
-- ----------------------------
DROP TABLE IF EXISTS "public"."product_templates";
CREATE TABLE "public"."product_templates" (
  "id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "pid" varchar(64) COLLATE "pg_catalog"."default",
  "classify" varchar(64) COLLATE "pg_catalog"."default",
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "key" varchar(64) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "type" varchar(64) COLLATE "pg_catalog"."default",
  "define" text COLLATE "pg_catalog"."default",
  "org_id" int4
)
;
COMMENT ON COLUMN "public"."product_templates"."pid" IS '产品Id';
COMMENT ON COLUMN "public"."product_templates"."classify" IS '模型归类';
COMMENT ON COLUMN "public"."product_templates"."name" IS '名称';
COMMENT ON COLUMN "public"."product_templates"."key" IS '标识';
COMMENT ON COLUMN "public"."product_templates"."description" IS '属性说明';
COMMENT ON COLUMN "public"."product_templates"."type" IS '数据类型';
COMMENT ON COLUMN "public"."product_templates"."define" IS '数据约束';
COMMENT ON COLUMN "public"."product_templates"."org_id" IS '机构ID';

-- ----------------------------
-- Records of product_templates
-- ----------------------------

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS "public"."products";
CREATE TABLE "public"."products" (
  "id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "name" varchar(128) COLLATE "pg_catalog"."default",
  "photo_url" varchar(255) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "product_category_id" varchar(191) COLLATE "pg_catalog"."default",
  "protocol_name" varchar(64) COLLATE "pg_catalog"."default",
  "device_type" varchar(64) COLLATE "pg_catalog"."default",
  "rule_chain_id" varchar(64) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."products"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."products"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."products"."name" IS '产品名称';
COMMENT ON COLUMN "public"."products"."photo_url" IS '图片地址';
COMMENT ON COLUMN "public"."products"."description" IS '产品说明';
COMMENT ON COLUMN "public"."products"."product_category_id" IS '产品类型Id';
COMMENT ON COLUMN "public"."products"."protocol_name" IS '协议名称';
COMMENT ON COLUMN "public"."products"."device_type" IS '设备类型';
COMMENT ON COLUMN "public"."products"."rule_chain_id" IS '规则链Id';
COMMENT ON COLUMN "public"."products"."status" IS '状态';

-- ----------------------------
-- Records of products
-- ----------------------------

-- ----------------------------
-- Table structure for rule_chain
-- ----------------------------
DROP TABLE IF EXISTS "public"."rule_chain";
CREATE TABLE "public"."rule_chain" (
  "id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "root" varchar(1) COLLATE "pg_catalog"."default",
  "rule_name" varchar(50) COLLATE "pg_catalog"."default",
  "rule_base64" text COLLATE "pg_catalog"."default",
  "rule_remark" varchar(256) COLLATE "pg_catalog"."default",
  "rule_data_json" text COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."rule_chain"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."rule_chain"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."rule_chain"."root" IS '是否根节点1 根链 0 普通链';
COMMENT ON COLUMN "public"."rule_chain"."rule_name" IS '名称';
COMMENT ON COLUMN "public"."rule_chain"."rule_base64" IS 'Base64缩略图';
COMMENT ON COLUMN "public"."rule_chain"."rule_remark" IS '说明';
COMMENT ON COLUMN "public"."rule_chain"."rule_data_json" IS 'Json数据';

-- ----------------------------
-- Records of rule_chain
-- ----------------------------
INSERT INTO "public"."rule_chain" VALUES ('rulee765e9ef022812a8b89dfb4c', 'panda', 2, '2023-07-21 16:17:51', '2023-08-03 10:19:33', '1', 'Root Rule Chain', 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAABPIAAAMvCAYAAABY+f5KAAAAAXNSR0IArs4c6QAAIABJREFUeF7s3QeUXVd9L/7fjLqtYhWrWc2SJRckW5IlcMPgApaxE2KbLAjEhgcs0kghCayX/FMg5JG3QhJSgBC/UA0hCeCSUFwxNnHBliVhyRaWVWz1Xq1mlfmvfYY7vjOa0YykuXfuuedz1pqlkebcfX77s7e8lr7e++yGJf9vR1O4CBAgQIAAAQIECBAgQIAAAQIECBCoaYGGFOS97oNn1HSRiiNAgAABAgQIECBAgAABAgQIECBQZIHn/nVnCPKKPAP0nQABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgQIAAAQIECBAgQKDoAoK8os8A/SdAgAABAgQIECBAgAABAgQIEMiFgCAvF8OkSAIECBAgQIAAAQIECBAgQIAAgaILCPKKPgP0nwABAgQIECBAgAABAgQIECBAIBcCgrxcDJMiCRAgQIAAAQIECBAgQIAAAQIEii4gyCv6DNB/AgQIECBAgAABAgQIECBAgACBXAgI8nIxTIokQIAAAQIECBAgQIAAAQIECBAouoAgr+gzQP8JECBAgAABAgQIECBAgAABAgRyISDIy8UwKZIAAQIECBAgUB2Bw4cPx+rVq2PXrl1x9OjR4z60qampOkV5SpcEevXqFUOGDIkJEyZE7969u/QZNxEgQIAAAQL5EhDk5Wu8VEuAAAECBAgQqJjAwoULY+H8RTFi6KgYNGBoRDR027MaekVEY0T06r42u624ummoKXa/sj22btsUs2bNjFmzZtVNz3SEAAECBAgQaBYQ5JkJBAgQIECAAAEC8cMHH45Xdu6P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AAECBAgQIECAQM0ICPJqZigUQoAAAQIECHSXQFNTU9x5552tTqa97LLLYvr06d31CO0QIECAAAECBAgQqLqAIK/q5B5IgAABAgQIVFLgwIED8fWvfz2OHj3a8ph58+bFhAkTKvlYbRMgQIAAAQIECBCouIAgr+LEHkCAAAECBAhUS2DTpk1xzz33tHrcO97xjhg2bFi1SvAcAgQIECBAgAABAhUTEORVjFbDBAgQIECAQDUFVq9eHffee2/LIwcMGBDvete7ok+fPtUsw7MIECBAgAABAgQIVExAkFcxWg0TIECAAAEC1RJIB1qkgy1K16hRo+Ltb397tR7vOQQIECBAgAABAgSqIiDIqwqzhxAgQIAAAQKVEnjuuefisccea2l+xIgRcfPNN1fqcdolQIAAAQIECBAg0GMCgrweo/dgAgQIECBA4FQFFi1aFE899VRLM4MHD86207oIECBAgAABAgQI1KOAIK8eR1WfCBAgQIBAAQQWLlwYTz/9dEtP0zvxbr311gL0XBcJECBAgAABAgSKKiDIK+rI6zcBAgQIEMixwJIlS+Lxxx9v6UGvXr3iAx/4QI57pHQCBAgQIECAAAECnQsI8jo3cgcBAgQIECBQQwI/+9nP4tFHH21V0Yc+9KEaqlApBAgQIECAAAECBCojIMirjKtWCRAgQIAAgQoIrFixIh566KFWLb///e+P3r17V+BpmiRAgAABAgQIECBQWwKCvNoaD9UQIECAAAECHQjs2rUrli9f3uqnF110kRDPjCFAgAABAgQIECiMgCCvMEOtowQIECBAgAABAgQIECBAgAABAnkWEOTlefTUToAAAQIECBAgQIAAAQIECBAgUBgBQV5hhlpHCRAgQIBAbQl89icRq3dF/PVbI8q/r60qVUOAAAECBAgQIECgdgQEebUzFiohQIAAAQI1LTB/XcSf/jDiw2+IuGHaqZcqyDt1Qy0QIECAAAECBAgUS0CQV6zx1lsCBAgQIHDSAil4e3xNxGl9Ij4zL2JQv+am9hyM+Mi9ETed/1rA971lEXctbX3f8R58Kivy2nv+SXfSBwkQIECAAAECBAjUsIAgr4YHR2kECBAgQKBWBFJY9scPRsybGvFvz0Z85NKIOWcJ8mplfNRBgAABAgQIECBQDAFBXjHGWS8JECBAgMApCZSvsPvkI81NpXfbrd8T8eHvRezY3/xnQwdEvOGsiHuXv/a4Wy6IuGRcxGeeiJg6POKRlyL+8PKIVTuOfUde+tTT65o/mz6XtvG2t+KutILv9y499vmfvSFi/e7mbcAHDje3lZ6XtgOX2lqxvfUzTgnHhwkQIECAAAECBAhUSUCQVyVojyFAgAABAnkW+Nj9EROGNAdrKdT74oKIFJiNHdS1rbWl9+ulMC21ka6278j7zvOvBW6lgPADsyOunHjs1t3yz7YN+tJn/+yHEX9xdXN96fcfva95FeGTa18LD/M8HmonQIAAAQIECBAopoAgr5jjrtcECBAgQKDLAuWhWgri2v6+K+/IS0FeWpH36euaw7X2grzSCbalwlJYl673zjyxIC8FjX/z2LHdS6vyRp3evFLvrMFdf39fl6HcSIAAAQIECBAgQKDCAoK8CgNrngABAgQI5F0gBWpptVzba8qw5jAsXZ0ddnEyQV5pFeDJBHmdHbRRCvvmntW8RdhFgAABAgQIECBAIA8Cgrw8jJIaCRAgQIBADwmUVtsNG9A68ErB3P95NOL/uzLi3BHdE+SlcO2TVzcfolHefvp9CvXSVf5evnOGNf++va216b19aVtu6b14//x0xG/MjfjBixFXTGxeFXiiJ+v20BB4LAECBAgQIECAAIEWAUGeyUCAAAECBAh0KNA2UCu/sTxcK63aS4ddpHfnpat0CEb5YRfH21r7wtaI/YcjSgdRlA6oSG2V3rGXDq9Iz7hwVMS+Q6+Fi22fX37YRf/erwWE5YdzlP+5KUCAAAECBAgQIEAgDwKCvDyMkhoJECBAgAABAgQIECBAgAABAgQKLyDIK/wUAECAAAECBAgQIECAAAECBAgQIJAHAUFeHkZJjQQIECBAgAABAgQIECBAgAABAoUXEOQVfgoAIECAAAECBAgQIECAAAECBAgQyIOAIC8Po6RGAgQIECBQBwL79++P559/vlVPpk+fHv369auD3ukCAQIECBAgQIAAgcoLCPIqb+wJBAgQIECg8AK7d++Of//3f2/lcOONN8bYsWMLbwOAAAECBAgQIECAQFcFBHldlXIfAQIECBAgcFICaSXeHXfc0eqz8+bNiwkTJpxUez5EgAABAgQIECBAoKgCgryijrx+EyBAgACBKggcPnw4vvSlL7V60jXXXBNTpkypwtM9ggABAgQIECBAgEB9CQjy6ms89YYAAQIECNSUwO23396qniuvvDLOO++8mqpRMQQIECBAgAABAgTyIiDIy8tIqZMAAQIECORM4Itf/GIcOXKkperLLrss0uEWLgIECBAgQIAAAQIETk5AkHdybj5FgAABAgQIHEcgvRMvvRuvdM2dOzdmzZrFjAABAgQIECBAgACBUxAQ5J0Cno8SIECAAAECxwqk02nTKbWl6/Wvf33MnDkTFQECBAgQIECAAAECpyggyDtFQB8nQIAAAQIEXhO48847Y+vWrS1/cPnll8frXvc6RAQIECBAgAABAgQIdIOAIK8bEDVBgAABAgQIRNxzzz2xadOmFoqrr746zjnnHDQECBAgQIAAAQIECHSTgCCvmyA1Q4AAAQIEiipw8ODB+OY3vxmvvvpqC8G8efNiwoQJRSXRbwIECBAgQIAAAQIVERDkVYRVowQIECBAoBgC69ati+9973utOvuLv/iLMXr06GIA6CUBAgQIECBAgACBKgoI8qqI7VEECBAgQKCeBBYtWhRPPfVUS5caGhriXe96VwwaNKieuqkvBAgQIECAAAECBGpGQJBXM0OhEAIECBAgkB+BH//4x7F06dKWggcOHJiFeI2NjfnphEoJECBAgAABAgQI5ExAkJezAVMuAQIECBDoaYH7778/XnrppZYy0jbatJ3WRYAAAQIECBAgQIBAZQUEeZX11ToBAgQIEKgbgZ07d8YjjzzS6mTayZMnx7XXXls3fdQRAgQIECBAgAABArUsIMir5dFRGwECBAgQqBGB9evXx6OPPhq7d+9uqejCCy+MSy65pEYqVAYBAgQIECBAgACB+hcQ5NX/GOshAQIECBA4JYEXX3wx0jvxDh8+3NJOCvBSkOciQIAAAQIECBAgQKB6AoK86ll7EgECBAgQyJXAwYMH4+mnn47nn3++pe5+/frFG9/4xkhbal0ECBAgQIAAAQIECFRXQJBXXW9PI0CAAAECuRBYvXp1FuJt27atpd7hw4fH5ZdfHulwCxcBAgQIECBAgAABAtUXEORV39wTCRAgQIBATQvMnz8/FixY0KrG173udTFnzpxIK/JcBAgQIECAAAECBAj0jIAgr2fcPZUAAQIECNScwObNmyOFeGvXrm2pbdCgQTF37tw455xzaq5eBREgQIAAAQIECBAomoAgr2gjrr8ECBAgQKAdgSVLlmRbaQ8dOtTy0xTe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CRAgQIAAAQIE8iggyMvjqKmZAAECBAj0gEDbMK+hoSFuuOEGYV4PjIVHEiBAgAABAgQIFFNAkFfMcddrAgQIECBwUgKrVq2KBx54oNVnb7zxRmHeSWn6EAECBAgQIECAAIETExDknZiXuwkQIECAQOEF1q5dG9///veFeYWfCQAIECBAgAABAgSqLSDIq7a45xEgQIAAgToQ2Lx5c9x9993CvDoYS10gQIAAAQIECBDIj4AgLz9jpVICBAgQIFBTAjt37oz//M//FObV1KgohgABAgQIECBAoJ4FBHn1PLr6RoAAAQIEKiywd+/e+MY3vtHqKbfccksMHz68wk/WPAECBAgQIECAAIHiCQjyijfmekyAAAECBLpV4NVXX42vfOUrrdp897vfHQMHDuzW52iMAAECBAgQIECAQNEFBHlFnwH6T4AAAQIEukng9ttvb2mpX79+8c53vjP69+/fTa1rhgABAgQIECBAgAABQZ45QIAAAQIECHSLQNtttmPGjInrr78+evfu3S3ta4QAAQIECBAgQIBA0QUEeUWfAfpPgAABAgS6UWDTpk1xzz33tLQ4adKkeMtb3hINDQ3d+BRNESBAgAABAgQIECimgCCvmOOu1wQIECBAoGICK1eujAcffLCl/WnTpsWb3/zmij1PwwQIECBAgAABAgSKIiDIK8pI6ycBAgQIEKiiwOLFi+OJJ55oeeLs2bNjzpw5VazAowgQIECAAAECBAjUn4Agr/7GVI8IECBAgEBNCDz55JPx7LPPttTyxje+Mc4///yaqE0RBAgQIECAAAECBPIoIMjL46ipmQABAgQI5EQgbbFNW21L13XXXRcTJ07MSfXKJECAAAECBAgQIFBbAoK82hoP1RAgQIAAgboSOHjwYNx3332xcePGrF+nn356zJs3L4YPH15X/dQZAgQIECBAgAABAtUQEORVQ9kzCBAgQIBAgQW2b98e9957b7zyyiuZwsiRIyOtzBswYECBVXSdAAECBAgQIECAwIkLCPJO3MwnCBAgQIAAAQIECBAgQIAAAQIECFRdQJBXdXIPJECAAAECBAgQIECAAAECBAgQIHDiAoK8EzfzCQIECBAgQIAAAQIECBAgQIAAAQJVFxDkVZ3cAwkQIECAQH0LfOz+iAlDIj78hvrup94RIECAAAECBAgQqLaAIK/a4p5HgAABAgRyIrB+T8SHvxexY/+xBd9yQcdBnSAvJwOsTAIECBAgQIAAgdwJCPJyN2QKJkCAAAEC1Rf43rKIu5ZGfGZexKB+x39+tYO8E6mt+nKeSIAAAQIECBAgQKD7BAR53WepJQIECBAgULcCJxKWCfLqdhroGAECBAgQIECAQA8LCPJ6eAA8ngABAgQI5EGgbZC352DER+6NWLG9ufq5Z0X89Vubvy8P8jq7b8RpEcu2NbczZVjERy+P+KMHm7fzDh0Q8dkbIsYOeq3dp9c1f5/uTasDv7oo4jvPvyaYtvxeMi7iM09ETB0e8chLER+8uPmeD8yOuGFa873z1zXf8+nrXms/D+OgRgIECBAgQIAAgWILCPKKPf56T4AAAQIEuiTQNshLYd2bJr0WjJWHd22/P959izdFfPLqiHNHNAeD2/e/Ft6ldtKVAsLP/qT5+9IBGun3q3c1/6xtbSmk+9MfNtfW3v2pnWqvGuwSspsIECBAgAABAgQIdCIgyDNFCBAgQIAAgU4FysOyPa+2fwhGaVVeKSS7OR2I0c5hGW3v6yhsK4V1f/qm1qv/SsWWVuU9+nLr9/e1t9ouHdzx0fsiPnJpxNjBr30/56xOu+4GAgQIECBAgAABAjUjIMirmaFQCAECBAgQqF2BtkFeKRRrLwgrD/K6cl9Xg7ybzn9tBWC5VHsr8trbNluq6+yhzVtuS1uBa1ddZQQIECBAgAABAgRaCwjyzAgCBAgQIECgU4H2ttamD5XCsC8vjJgxMiIFe2231nblvnRP+XbZtr9PP1u08bVTc1M96UrbZ7sa5KWVel+YH7H/UMS7L2w/FOwUwg0ECBAgQIAAAQIEelBAkNeD+B5NgAABAgTyItDZYRfpkInSyrrjHXbR0X2dBXnp56nd0mEX5YdrpG2zpS285YddtD3IonTwRmorHZQxqF9e9NVJgAABAgQIECBAoFlAkGcmECBAgAABAoUQKAV5HW3RLQSCThIgQIAAAQIECORaQJCX6+FTPAECBAgQINBVgbSq8IsLXjsVt6ufcx8BAgQIECBAgACBWhEQ5NXKSKiDAAECBAgQqIhAaSXeut0Rn7y6+T1+LgIECBAgQIAAAQJ5FBDk5XHU1EyAAAECBAgQIECAAAECBAgQIFA4AUFe4YZchwkQIECAQH4EnnzyyXj22WdbCr7yyivjvPPOy08HVEqAAAECBAgQIECgGwUEed2IqSkCBAgQIECg+wUeeOCBWLVqVdZwQ0NDzJs3L8aPH9/9D9IiAQIECBAgQIAAgRoXEOTV+AApjwABAgQIFF1g//79cd9998XmzZszioEDB2Zh3rBhw4pOo/8ECBAgQIAAAQIFExDkFWzAdZcAAQIECORRYNu2bXHvvffG3r17s/JHjx4d1113XfTr1y+P3VEzAQIECBAgQIAAgZMSEOSdFJsPESBAgAABAtUWePnll7OVeaVr8uTJce2111a7DM8jQIAAAQIECBAg0GMCgrweo/dgAgQIECBA4EQFli5dGj/+8Y9bPnbhhRfGJZdccqLNuJ8AAQIECBAgQIBALgUEebkcNkUTIECAAIHiCsyfPz8WLFjQAjB37tyYNWtWcUH0nAABAgQIECBAoDACgrzCDLWOEiBAgACB+hH40Y9+FMuWLWvpUFqVl1bnuQgQIECAAAECBAjUs4Agr55HV98IECBAgECdCjQ1NcX3v//9WLduXUsPr7nmmpgyZUqd9li3CBAgQIAAAQIECEQI8swCAgQIECBAIJcCBw8ejHvuuSd27tzZUv873vGOGDZsWC77o2gCBAgQIECAAAECnQkI8joT8nMCBAgQIECgZgVeeeWV+Na3vhWHDh1qqfFDH/pQzdarMAIECBAgQIAAAQKnIiDIOxU9nyVAgAABAgR6XGD79u3x7W9/u1Ud73//+6N37949XpsCCBAgQIAAAQIECHSngCCvOzW1RYAAAQIECPSIwMaNG+O//uu/Wj37ve99b/Tr169H6vFQAgQIECBAgAABApUQEORVQlWbBAgQIECAQNUFNm/eHHfffXer577nPe+J008/veq1eCABAgQIECBAgACBSggI8iqhqk0CBAgQIECgRwR27NiRvTOv/HrXu94VgwcP7pF6PJQAAQIECBAgQIBAdwoI8rpTU1sECBAgQIBAjwu0F+bdcsstMXz48B6vTQEECBAgQIAAAQIETkVAkHcqej5LgAABAgQI1KRAe2HeL/3SL8XIkSNrsl5FESBAgAABAgQIEOiKgCCvK0ruIUCAAAECBHInkMK873znO3H06NGW2m+88cYYO3Zs7vqiYAIECBAgQIAAAQJJQJBnHhAgQIAAAQJ1K5DCvHSa7cGDB1v6eP3118f48ePrts86RoAAAQIECBAgUL8Cgrz6HVs9I0CAAAECBCIihXn33ntv7Nmzp8XjLW95S5x99tl8CBAgQIAAAQIECORKQJCXq+FSLAECBAgQIHAyAinMe+ihh2L79u0tH7/sssti+vTpJ9OczxAgQIAAAQIECBDoEQFBXo+weygBAgQIECBQbYH2wrwLLrggLr300ujVq1e1y/E8AgQIECBAgAABAicsIMg7YTIfIECAAAECBPIqkMK8xx57LNavX9/ShTFjxkRanTd8+PC8dkvdBAgQIECAAAECBREQ5BVkoHWTAAECBAgQaBY4cuRIPP7447F06dIWktNOOy1bmTdlyhRMBAgQIECAAAECBGpWQJBXs0OjMAIECBAgQKCSAosXL44nnnii1SPmzJkTs2fPruRjtU2AAAECBAgQIEDgpAUEeSdN54MECBAgQIBA3gVWr16drc7bvXt3S1fOOeecbHXegAED8t499RMgQIAAAQIECNSZgCCvzgZUdwgQIECAAIETE0ghXgrzUqhXukaMGJGFeen9eS4CBAgQIECAAAECtSIgyKuVkVAHAQIECBAg0KMCaZtt2m5bunr37h1pq+2FF17Yo3V5OAECBAgQIECAAIGSgCDPXCBAgAABAgQI/FwgHYCRVuelAzFK19ixY2PmzJkxbtw4TgQIECBAgAABAgR6VECQ16P8Hk6AAAECBAjUmsD69evjySefjK1bt7Yqbfr06Vmgl064rcfr8OHD2fbiXbt2xdGjR4/bxaampnokyG2fevXqFUOGDIkJEyZEWknqIkCAAAECBOpXQJBXv2OrZwQIECBAgMBJCqRQa9GiRfHss89G+r50DR48OAvzzjvvvJNsuTY/tnDhwlg0f0GMHtQ3hvU9Eg3RjUFdY0Q0NkVDr+a+N9QmQa6rOhoNse1Ar9i461DMnDU7Zs2alev+KJ4AAQIECBDoWECQZ3YQIECAAAECBDoQSKvyfvrTn8aKFSta3TFp0qQs0Bs5cmTu7R5+4P44sG1TzB0TMbhf5WK2FOQ19m6Kjdt3ZmZjRpyRe7ta68CuA03x5NqI/kNGx1XXvrXWylMPAQIECBAg0A0CgrxuQNQEAQIECBAgUN8CKchLgV75dtvGxsYszEtfed3OmFbibXxhSVwzqXIBXvnMSGHe5t07BHkV/uty3/KmGD15hpV5FXbWPAECBAgQ6AkBQV5PqHsmAQIECBAgkDuBjrbbjhgxIqZNmxbnnHNO9O/fPzf9Sv2548tfjV+Y1reiK/HagmzatSMae1uRV8mJklbm3f3cobj1tvfmNmSupI+2CRAgQIBAngUEeXkePbUTIECAAAECVRfoaLvt6aefHlOnTs0CvWHDhlW9rhN94MqVK+OFpx6r2mq8Un0bd+yMXn2bbK090QE7wfvTqrxzZ10RkydPPsFPup0AAQIECBCoZQFBXi2PjtoIECBAgACBmhVI223TYRhbtmw5psYU6KWvcePG1Wz9aVvtq6t+GrPGVveU041bd0avAYK8Sk+M+WsPR5+xM22vrTS09gkQIECAQJUFBHlVBvc4AgQIECBAoL4E0sq2F198MV5++eVjOnbWWWdlK/RSqJfeqVdL1zPPPBNHXn42Zo7tU9WyBHnV4V6w9lDEqAtjzpw51XmgpxAgQIAAAQJVERDkVYXZQwgQIECAAIF6F9iwYUMsX748C/XS++fKr6FDh2ZhXjrt9owzuv+01sWLF8eMGTNOiHj+/PlxdPViQd4JqeXnZkFefsZKpQQIECBA4EQEBHknouVeAgQIECBAgEAnArt27crCvPS1Z8+eY+4eOXJkjBkzJvuaMGFCt3imUG7ZsmVZmNfVQE+Q1y30NduIIK9mh0ZhBAgQIEDglAQEeafE58MECBAgQIAAgfYFDh06lIVraZXepk2b2r1pwIABWaCXtuCmUC8dmHEyVwrlFixYkH104MCBXQr0BHknI52fzwjy8jNWKiVAgAABAiciIMg7ES33EiBAgAABAgROQuCll16K9LV+/fp45ZVXOmxhxIgRMX78+CzUGzVqVJefVB7klT7UWaCXPtO0enFclON35P3b9x+PvfsPxoKlL8Xt3374GK8PveOq+LuPvicG9Ovb8rMXXtoQn/jCXfFPf3RbDB8ysF3jdM+X7nokPv6bN8fqDdviVz72ufjpstUdtr9g6cvx8FPPxx+89/r4/U9/o1Utn/jNW+JPPvT27LOp3oljz4zLZ07t8tie7I2CvJOV8zkCBAgQIFDbAoK82h4f1REgQIAAAQJ1JrBu3bpI79MrfXXUvV69esWQIUOyd+oNHjw4+770lVbylV/tBXmdBXp5D/L2H3w1Pv75O+NX3nZZfPP7j8f7b3pTnDtpTKvArhTG/e1XfxB//vnvdGkmlcK/ux6aHyvXbsnCufSc1H66fvA/P43f+9V58diiF+Pl9Vvi3W+7LPvz9Pv/eviZ7PtSLSkQTPfPnT4luzddgrwuDYObCBAgQIAAgQ4EBHmmBgECBAgQIECghwT2798fa9asibVr12ar9fbt29elSvr27dsS6qWQb/v27dmKv3Q1NTVFQ0PDMe20XaGX9yCvtGruY++/Mf76S989bpBXWpG3bdcr8dt/9bX481+/qVXo1xn62k3b4+UN27KVdKmNNRu3x4+efj4L9MqvUrj4i1ddHF//7v/Er954RTy9ZIUgrzNgPydAgAABAgS6LCDI6zKVGwkQIECAAAEClRXYunVrFuylVXs7duyIFPR15SoP70rfdxTo9e7dO9IW3nSNObwtt1tr0zbVdN10zZxjtrOWzNpurS2tmtu9d3+7W3HfetmM+Ppf/UarLbcpnGu7XbbUfun+f/6Ph2LyuDOzWtLqPUFeV2atewgQIECAAIGTERDknYyazxAgQIAAAQIEqiBw4MCB2LlzZxbqlX5N3+/du7fV09sL8ko3dBTopT+ffWbfXAZ5aVXcr/7RP8etN17eEp6l7awjhg6KP/mnb8Vf/vYvx9Yde1rec5dW5JVWy7XdgpucUsB35fs+GY9+5U9j9vkTW4K7Oz7169lW2NIW2vTcFNql7bbl790rhX23XDs3HnhiiSCvCn83PIIAAQIECPz/7d1d6N11HcDx7/5bawW1EGGbF8NGaQRL59NFGWqGEWhdiHXT000P1mW33kRdRRddSyBh1F30BKVdZKQRYVu2ChRzMHaxlVM3etjcxuJ34Iy5/adD9t9/7/N/DSKwds7nvD4fb96ch7UqIOSt1c173QQIECBAgEBWYPpF3CnozQPfvn37xtGjR2ev59xwt4ghb3o33s9/u2fcd8euM99RN732KbTNQ965P2Qx/Z1Hf/HUee+4m/7573Y/e96PYsy/A2/6nrzX+369KfbNvyfPR2uz/0oZnAABAgQIZASEvMyqDEqAAAECBAgQWF7g7B+7eKOP1m7atGls27ZtHD9+fGx59V/Jd+T97Ind49DhIzOMC/1i7Vxq+njtg5+6e3zhoYfHlqs3j499cOf4+nd+eMFTmn8cd/ol2umdeNPHZacfy3jw03fP/s78HXnTr9lO39P3iTtvmv3z+S/oPr//0OwdedOv2N57x67Zd+n5sQv/5hIgQIAAAQKXSkDIu1SSHocAAQIECBAgsEoCy4W8c0dZtB+7mH9H3tnvhpuC25N7nh3TR1y/eP9dZwi++4NfzWLa9AMU08duz3633vQ4y/2S7PwdeRcb8qbnuPPW989+Qfc927eMQ4ePjgfuuc2v1q7SvxOelgABAgQILKqAkLeom/W6CBAgQIAAgTUjcHbIe6OAN//f679ae27IOzvWffz2G8a//3vsoj52ezEh70I/djH/WO3ZH+l9+6aNr/mOvXlonNwv9Fwrcai7D5wYY8sHxi233LISD+8xCRAgQIAAgVUSEPJWCd7TEiBAgAABAgQulcByIe/cd+Cd+1yLEvLmv1r7mXtvH+9797Yz35H32FN/GdP32z30pU/OXvqFvj9vubg2//GL5X7FdrmdTR+x/eWTz4wvP/CRWcT78E3XzyLitx7+6ezXbOcxT8i7VBfvcQgQIECAwNoVEPLW7u69cgIECBAgQGBBBM4OeW8U8OYveRFC3n/+d3z2tSb+tAAAIABJREFUHXlTxPvQje89L9ZNQW76rrop5p0b8qbINv2IxQ3XbR8/+vbXxvXXbjtzDd/78ROz78ab/ky/jvv47/de8FK+8dX7x123vX/8/R8HXjPL9BfmP34xfdT2K998ZNnnWqkT/NOBE2Odd+StFK/HJUCAAAECqyYg5K0avScmQIAAAQIECFwagSnKPffcc2Pnzp2z/1zMn3rIu5jXuJb/P0LeWt6+106AAAECiywg5C3ydr02AgQIECBAYE0I7N2796ID3hxEyFvs0xDyFnu/Xh0BAgQIrF0BIW/t7t4rJ0CAAAECBNawgJC32MufQt7S1hvGzTffvNgv1KsjQIAAAQJrTEDIW2ML93IJECBAgAABApPAnj17xqv7nhm7rtlwWUEOvvjKWP+202Pb1e+6rM+71p7s6QMnx1uuuXHs2rVrrb10r5cAAQIECCy0gJC30Ov14ggQIECAAAECywu88MIL49k/PjXuvnbdZSU6+PIrY/1GIW+l0R97/vS4ftftY8eOHSv9VB6fAAECBAgQuIwCQt5lxPZUBAgQIECAAIErReDkyZPj0Ue+P+67buN451svX8w7dOTlsbRheEfeCh7CkWOnx0/+dmJ89nOfHxs2XN53XK7gy/LQBAgQIECAwBhDyHMGBAgQIECAAIE1KjB9vPbgs3+9bO/KW7d+jH8efXmm7aO1K3d007vxtu7Y6WO1K0fskQkQIECAwKoJCHmrRu+JCRAgQIAAAQKrL/CbXz8+jh0+NG7dNlb0nXlTxFvacHocfOkVIW+F1j69E+8PB8bYtHnruOuj96zQs3hYAgQIECBAYDUFhLzV1PfcBAgQIECAAIErQGB6Z96fn949tr5j47hq46mxNE5fuqmWxhhLp8cU8qY/L750ZPbfV1+1+dI9R+CRLqHoea/29Fg3Dh9bPw4eOTFu3HWTd+IF7sGIBAgQIEDgzQoIeW9Wzt8jQIAAAQIECCyQwPSdefv37x9HjhwZp06dWqBX1n8p69a9/ncYLi0tjc2bN4/t27f7Trz+ur0CAgQIECDwugJCngMhQIAAAQIECBAgQIAAAQIECBAgEBAQ8gJLMiIBAgQIECBAgAABAgQIECBAgAABIc8NECBAgAABAgQIECBAgAABAgQIEAgICHmBJRmRAAECBAgQIECAAAECBAgQIECAgJDnBggQIECAAAECBAgQIECAAAECBAgEBIS8wJKMSIAAAQIECBAgQIAAAQIECBAgQEDIcwMECBAgQIAAAQIECBAgQIAAAQIEAgJCXmBJRiRAgAABAgQIECBAgAABAgQIECAg5LkBAgQIECBAgAABAgQIECBAgAABAgEBIS+wJCMSIECAAAECBAgQIECAAAECBAgQEPLcAAECBAgQIECAAAECBAgQIECAAIGAgJAXWJIRCRAgQIAAAQIECBAgQIAAAQIECAh5boAAAQIECBAgQIAAAQIECBAgQIBAQEDICyzJiAQIECBAgAABAgQIECBAgAABAgSEPDdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQEPICSzIiAQIECBAgQIAAAQIECBAgQIAAASHPDRAgQIAAAQIECBAgQIAAAQIECBAICAh5gSUZkQABAgQIECBAgAABAgQIECBAgICQ5wYIECBAgAABAgQIECBAgAABAgQIBASEvMCSjEiAAAECBAgQIECAAAECBAgQIEBAyHMDBAgQIECAAAECBAgQIECAAAECBAICQl5gSUYkQIAAAQIECBAgQIAAAQIECBAgIOS5AQIECBAgQIAAAQIECBAgQIAAAQIBASEvsCQjEiBAgAABAgQIECBAgAABAgQIEBDy3AABAgQIECBAgAABAgQIECBAgACBgICQF1iSEQkQIECAAAECBAgQIECAAAECBAgIeW6AAAECBAgQIECAAAECBAgQIECAQEBAyAssyYgECBAgQIAAAQIECBAgQIAAAQIEhDw3QIAAAQIECBAgQIAAAQIECBAgQCAgIOQFlmREAgQIECBAgAABAgQIECBAgAABAkKeGyBAgAABAgQIECBAgAABAgQIECAQEBDyAksyIgECBAgQIECAAAECBAgQIECAAAEhzw0QIECAAAECBAgQIECAAAECBAgQCAgIeYElGZEAAQIECBAgQIAAAQIECBAgQICAkOcGCBAgQIAAAQIECBAgQIAAAQIECAQEhLzAkoxIgAABAgQIECBAgAABAgQIECBAQMhzAwQIECBAgAABAgQIECBAgAABAgQCAkJeYElGJECAAAECBAgQIECAAAECBAgQICDkuQECBAgQIECAAAECBAgQIECAAAECAQEhL7AkIxIgQIAAAQIECBAgQIAAAQIECBAQ8twAAQIECBAgQIAAAQIECBAgQIAAgYCAkBdYkhEJECBAgAABAgQIECBAgAABAgQICHlugAABAgQIECBAgAABAgQIECBAgEBAQMgLLMmIBAgQIECAAAECBAgQIECAAAECBIQ8N0CAAAECBAgQIECAAAECBAgQIEAgICDkBZZkRAIECBAgQIAAAQIECBAgQIAAAQJCnhsgQIAAAQIECBAgQIAAAQIECBAgEBAQ8gJLMiIBAgQIECBAgAABAgQIECBAgAABIc8NECBAgAABAgQIECBAgAABAgQIEAgICHmBJRmRAAECBAgQIECAAAECBAgQIECAgJDnBggQIECAAAECBAgQIECAAAECBAgEBIS8wJKMSIAAAQIECBAgQIAAAQIECBAgQEDIcwMECBAgQIAAAQIECBAgQIAAAQIEAgJCXmBJRiRAgAABAgQIECBAgAABAgQIECAg5LkBAgQIECBAgAABAgQIECBAgAABAgEBIS+wJCMSIECAAAECBAgQIECAAAECBAgQEPLcAAECBAgQIECAAAECBAgQIECAAIGAgJAXWJIRCRAgQIAAAQIECBAgQIAAAQIECAh5boAAAQIECBAgQIAAAQIECBAgQIBAQEDICyzJiAQIECBAgAABAgQIECBAgAABAgSEPDdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQEPICSzIiAQIECBAgQIAAAQIECBAgQIAAASHPDRAgQIAAAQIECBAgQIAAAQIECBAICAh5gSUZkQABAgQIECBAgAABAgQIECBAgICQ5wYIECBAgAABAgQIECBAgAABAgQIBASEvMCSjEiAAAECBAgQIECAAAECBAgQIEBAyHMDBAgQIECAAAECBAgQIECAAAECBAICQl5gSUYkQIAAAQIECBAgQIAAAQIECBAgIOS5AQIECBAgQIAAAQIECBAgQIAAAQIBASEvsCQjEiBAgAABAgQIECBAgAABAgQIEBDy3AABAgQIECBAgAABAgQIECBAgACBgICQF1iSEQkQIECAAAECBAgQIECAAAECBAgIeW6AAAECBAgQIECAAAECBAgQIECAQEBAyAssyYgECBAgQIAAAQIECBAgQIAAAQIEhDw3QIAAAQIECBAgQIAAAQIECBAgQCAgIOQFlmREAgQIECBAgAABAgQIECBAgAABAkKeGyBAgAABAgQIECBAgAABAgQIECAQEBDyAksyIgECBAgQIECAAAECBAgQIECAAAEhzw0QIECAAAECBAgQIECAAAECBAgQCAgIeYElGZEAAQIECBAgQIAAAQIECBAgQICAkOcGCBAgQIAAAQIECBAgQIAAAQIECAQEhLzAkoxIgAABAgQIECBAgAABAgQIECBAQMhzAwQIECBAgAABAgQIECBAgAABAgQCAkJeYElGJECAAAECBAgQIECAAAECBAgQICDkuQECBAgQIECAAAECBAgQIECAAAECAQEhL7AkIxIgQIAAAQIECBAgQIAAAQIECBAQ8twAAQIECBAgQIAAAQIECBAgQIAAgYCAkBdYkhEJECBAgAABAgQIECBAgAABAgQICHlugAABAgQIECBAgAABAgQIECBAgEBAQMgLLMmIBAgQIECAAAECBAgQIECAAAECBIQ8N0CAAAECBAgQIECAAAECBAgQIEAgICDkBZZkRAIECBAgQIAAAQIECBAgQIAAAQJCnhsgQIAAAQIECBAgQIAAAQIECBAgEBAQ8gJLMiIBAgQIECBAgAABAgQIECBAgAABIc8NECBAgAABAgQIECBAgAABAgQIEAgICHmBJRmRAAECBAgQIECAAAECBAgQIECAgJDnBggQIECAAAECBAgQIECAAAECBAgEBIS8wJKMSIAAAQIECBAgQIAAAQIECBAgQEDIcwMECBAgQIAAAQIECBAgQIAAAQIEAgJCXmBJRiRAgAABAgQIECBAgAABAgQIECAg5LkBAgQIECBAgAABAgQIECBAgAABAgEBIS+wJCMSIECAAAECBAgQIECAAAECBAgQEPLcAAECBAgQIECAAAECBAgQIECAAIGAgJAXWJIRCRAgQIAAAQIECBAgQIAAAQIECAh5boAAAQIECBAgQIAAAQIECBAgQIBAQEDICyzJiAQIECBAgAABAgQIECBAgAABAgSEPDdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQEPICSzIiAQIECBAgQIAAAQIECBAgQIAAASHPDRAgQIAAAQIECBAgQIAAAQIECBAICAh5gSUZkQABAgQIECBAgAABAgQIECBAgICQ5wYIECBAgAABAgQIECBAgAABAgQIBASEvMCSjEiAAAECBAgQIECAAAECBAgQIEBAyHMDBAgQIECAAAECBAgQIECAAAECBAICQl5gSUYkQIAAAQIECBAgQIAAAQIECBAgIOS5AQIECBAgQIAAAQIECBAgQIAAAQIBASEvsCQjEiBAgAABAgQIECBAgAABAgQIEBDy3AABAgQIECBAgAABAgQIECBAgACBgICQF1iSEQkQIECAAAECBAgQIECAAAECBAgIeW6AAAECBAgQIECAAAECBAgQIECAQEBAyAssyYgECBAgQIAAAQIECBAgQIAAAQIEhDw3QIAAAQIECBAgQIAAAQIECBAgQCAgIOQFlmREAgQIECBAgAABAgQIECBAgAABAkKeGyBAgAABAgQIECBAgAABAgQIECAQEBDyAksyIgECBAgQIECAAAECBAgQIECAAAEhzw0QIECAAAECBAgQIECAAAECBAgQCAgIeYElGZEAAQIECBAgQIAAAQIECBAgQICAkOcGCBAgQIAAAQIECBAgQIAAAQIECAQEhLzAkoxIgAABAgQIECBAgAABAgQIECBAQMhzAwQIECBAgAABAgQIECBAgAABAgQCAkJeYElGJECAAAECBAgQIECAAAECBAgQICDkuQECBAgQIECAAAECBAgQIECAAAECAQEhL7AkIxIgQIAAAQIECBAgQIAAAQIECBAQ8twAAQIECBAgQIAAAQIECBAgQIAAgYCAkBdYkhEJECBAgAABAgQIECBAgAABAgQICHlugAABAgQIECBAgAABAgQIECBAgEBAQMgLLMmIBAgQIECAAAECBAgQIECAAAECBIQ8N0CAAAECBAgQIECAAAECBAgQIEAgICDkBZZkRAIECBAgQIAAAQIECBAgQIAAAQJCnhsgQIAAAQIECBAgQIAAAQIECBAgEBAQ8gJLMiIBAgQIECBAgAABAgQIECBAgAABIc8NECBAgAABAgQIECBAgAABAgQIEAgICHmBJRmRAAECBAgQIECAAAECBAgQIECAgJDnBggQIECAAAECBAgQIECAAAECBAgEBIS8wJKMSIAAAQIECBAgQIAAAQIECBAgQEDIcwMECBAgQIAAAQIECBAgQIAAAQIEAgJCXmBJRiRAgAABAgQIECBAgAABAgQIECAg5LkBAgQIECBAgAABAgQIECBAgAABAgEBIS+wJCMSIECAAAECBAgQIECAAAECBAgQEPLcAAECBAgQIECAAAECBAgQIECAAIGAgJAXWJIRCRAgQIAAAQIECBAgQIAAAQIECAh5boAAAQIECBAgQIAAAQIECBAgQIBAQEDICyzJiAQIECBAgAABAgQIECBAgAABAgSEPDdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQEPICSzIiAQIECBAgQIAAAQIECBAgQIAAASHPDRAgQIAAAQIECBAgQIAAAQIECBAICAh5gSUZkQABAgQIECBAgAABAgQIECBAgICQ5wYIECBAgAABAgQIECBAgAABAgQIBASEvMCSjEiAAAECBAgQIECAAAECBAgQIEBAyHMDBAgQIECAAAECBAgQIECAAAECBAICQl5gSUYkQIAAAQIECBAgQIAAAQIECBAgIOS5AQIECBAgQIAAAQIECBAgQIAAAQIBASEvsCQjEiBAgAABAgQIECBAgAABAgQIEBDy3AABAgQIECBAgAABAgQIECBAgACBgICQF1iSEQkQIECAAAECBAgQIECAAAECBAgIeW6AAAECBAgQIECAAAECBAgQIECAQEBAyAssyYgECBAgQIAAAQIECBAgQIAAAQIEhDw3QIAAAQIECBAgQIAAAQIECBAgQCAgIOQFlmREAgQIECBAgAABAgQIECBAgAABAkKeGyBAgAABAgQIECBAgAABAgQIECAQEBDyAksyIgECBAgQIECAAAECBAgQIECAAAEhzw0QIECAAAECBAgQIECAAAECBAgQCAgIeYElGZEAAQIECBAgQIAAAQIECBAgQICAkOcGCBAgQIAAAQIECBAgQIAAAQIECAQEhLzAkoxIgAABAgQIECBAgAABAgQIECBAQMhzAwQIECBAgAABAgQIECBAgAABAgQCAkJeYElGJECAAAECBAgQIECAAAECBAgQICDkuQECBAgQIECAAAECBAgQIECAAAECAQEhL7AkIxIgQIAAAQIECBAgQIAAAQIECBAQ8twAAQIECBAgQIAAAQIECBAgQIAAgYCAkBdYkhEJECBAgAABAgQIECBAgAABAgQICHlugAABAgQIECBAgAABAgQIECBAgEBAQMgLLMmIBAgQIECAAAECBAgQIECAAAECBIQ8N0CAAAECBAgQIECAAAECBAgQIEAgICDkBZZkRAIECBAgQIAAAQIECBAgQIAAAQJCnhsgQIAAAQIECBAgQIAAAQIECBAgEBAQ8gJLMiIBAgQIECBAgAABAgQIECBAgAABIc8NECBAgAABAgQIECBAgAABAgQIEAgICHmBJRmRAAECBAgQIECAAAECBAgQIECAgJDnBggQIECAAAECBAgQIECAAAECBAgEBIS8wJKMSIAAAQIECBAgQIAAAQIECBAgQEDIcwMECBAgQIAAAQIECBAgQIAAAQIEAgJCXmBJRiRAgAABAgQIECBAgAABAgQIECAg5LkBAgQIECBAgAABAgQIECBAgAABAgEBIS+wJCMSIECAAAECBAgQIECAAAECBAgQEPLcAAECBAgQIECAAAECBAgQIECAAIGAgJAXWJIRCRAgQIAAAQIECBAgQIAAAQIECAh5boAAAQIECBAgQIAAAQIECBAgQIBAQEDICyzJiAQIECBAgAABAgQIECBAgAABAgSEPDdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQEPICSzIiAQIECBAgQIAAAQIECBAgQIAAASHPDRAgQIAAAQIECBAgQIAAAQIECBAICAh5gSUZkQABAgQIECBAgAABAgQIECBAgICQ5wYIECBAgAABAgQIECBAgAABAgQIBASEvMCSjEiAAAECBAgQIECAAAECBAgQIEBAyHMDBAgQIECAAAECBAgQIECAAAECBAICQl5gSUYkQIAAAQIECBAgQIAAAQIECBAgIOS5AQIECBAgQIAAAQIECBAgQIAAAQIBASEvsCQjEiBAgAABAgQIECBAgAABAgQIEBDy3AABAgQIECBAgAABAgQIECBAgACBgICQF1iSEQkQIECAAAECBAgQIECAAAECBAgIeW6AAAECBAgQIECAAAECBAgQIECAQEBAyAssyYgECBAgQIAAAQIECBAgQIAAAQIEhDw3QIAAAQIECBAgQIAAAQIECBAgQCAgIOQFlmREAgQIECBAgAABAgQIECBAgAABAkKeGyBAgAABAgQIECBAgAABAgQIECAQEBDyAksyIgECBAgQIECAAAECBAgQIECAAAEhzw0QIECAAAECBAgQIECAAAECBAgQCAgIeYElGZEAAQIECBAgQIAAAQIECBAgQICAkOcGCBAgQIAAAQIECBAgQIAAAQIECAQEhLzAkoxIgAABAgQIECBAgAABAgQIECBAQMhzAwQIECBAgAABAgQIECBAgAABAgQCAkJeYElGJECAAAECBAgQIECAAAECBAgQICDkuQECBAgQIECAAAECBAgQIECAAAECAQEhL7AkIxIgQIAAAQIECBAgQIAAAQIECBAQ8twAAQIECBAgQIAAAQIECBAgQIAAgYCAkBdYkhEJECBAgAABAgQIECBAgAABAgQICHlugAABAgQIECBAgAABAgQIECBAgEBAQMgLLMmIBAgQIECAAAECBAgQIECAAAECBIQ8N0CAAAECBAgQIECAAAECBAgQIEAgICDkBZZkRAIECBAgQIAAAQIECBAgQIAAAQJCnhsgQIAAAQIECBAgQIAAAQIECBAgEBAQ8gJLMiIBAgQIECBAgAABAgQIECBAgAABIc8NECBAgAABAgQIECBAgAABAgQIEAgICHmBJRmRAAECBAgQIECAAAECBAgQIECAgJDnBggQIECAAAECBAgQIECAAAECBAgEBIS8wJKMSIAAAQIECBAgQIAAAQIECBAgQEDIcwMECBAgQIAAAQIECBAgQIAAAQIEAgJCXmBJRiRAgAABAgQIECBAgAABAgQIECAg5LkBAgQIECBAgAABAgQIECBAgAABAgEBIS+wJCMSIECAAAECBAgQIECAAAECBAgQEPLcAAECBAgQIECAAAECBAgQIECAAIGAgJAXWJIRCRAgQIAAAQIECBAgQIAAAQIECAh5boAAAQIECBAgQIAAAQIECBAgQIBAQEDICyzJiAQIECBAgAABAgQIECBAgAABAgSEPDdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQEPICSzIiAQIECBAgQIAAAQIECBAgQIAAASHPDRAgQIAAAQIECBAgQIAAAQIECBAICAh5gSUZkQABAgQIECBAgAABAgQIECBAgICQ5wYIECBAgAABAgQIECBAgAABAgQIBASEvMCSjEiAAAECBAgQIECAAAECBAgQIEBAyHMDBAgQIECAAAECBAgQIECAAAECBAICQl5gSUYkQIAAAQIECBAgQIAAAQIECBAgIOS5AQIECBAgQIAAAQIECBAgQIAAAQIBASEvsCQjEiBAgAABAgQIECBAgAABAgQIEBDy3AABAgQIECBAgAABAgQIECBAgACBgICQF1iSEQkQIECAAAECBAgQIECAAAECBAgIeW6AAAECBAgQIECAAAECBAgQIECAQEBAyAssyYgECBAgQIAAAQIECBAgQIAAAQIEhDw3QIAAAQIECBAgQIAAAQIECBAgQCAgIOQFlmREAgQIECBAgAABAgQIECBAgAABAkKeGyBAgAABAgQIECBAgAABAgQIECAQEBDyAksyIgECBAgQIECAAAECBAgQIECAAAEhzw0QIECAAAECBAgQIECAAAECBAgQCAgIeYElGZEAAQIECBAgQIAAAQIECBAgQICAkOcGCBAgQIAAAQIECBAgQIAAAQIECAQEhLzAkoxIgAABAgQIECBAgAABAgQIECBAQMhzAwQIECBAgAABAgQIECBAgAABAgQCAkJeYElGJECAAAECBAgQIECAAAECBAgQICDkuQECBAgQIECAAAECBAgQIECAAAECAQEhL7AkIxIgQIAAAQIECBAgQIAAAQIECBAQ8twAAQIECBAgQIAAAQIECBAgQIAAgYCAkBdYkhEJECBAgAABAgQIECBAgAABAgQICHlugAABAgQIECBAgAABAgQIECBAgEBAQMgLLMmIBAgQIECAAAECBAgQIECAAAECBIQ8N0CAAAECBAgQIECAAAECBAgQIEAgICDkBZZkRAIECBAgQIAAAQIECBAgQIAAAQJCnhsgQIAAAQIECBAgQIAAAQIECBAgEBAQ8gJLMiIBAgQIECBAgAABAgQIECBAgAABIc8NECBAgAABAgQIECBAgAABAgQIEAgICHmBJRmRAAECBAgQIECAAAECBAgQIECAgJDnBggQIECAAAECBAgQIECAAAECBAgEBIS8wJKMSIAAAQIECBAgQIAAAQIECBAgQEDIcwMECBAgQIAAAQIECBAgQIAAAQIEAgJCXmBJRiRAgAABAgQIECBAgAABAgQIECAg5LkBAgQIECBAgAABAgQIECBAgAABAgEBIS+wJCMSIECAAAECBAgQIECAAAECBAgQEPLcAAECBAgQIECAAAECBAgQIECAAIGAgJAXWJIRCRAgQIAAAQIECBAgQIAAAQIECAh5boAAAQIECBAgQIAAAQIECBAgQIBAQEDICyzJiAQIECBAgAABAgQIECBAgAABAgSEPDdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQvxCEEAAAFj0lEQVQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQEPICSzIiAQIECBAgQIAAAQIECBAgQIAAASHPDRAgQIAAAQIECBAgQIAAAQIECBAICAh5gSUZkQABAgQIECBAgAABAgQIECBAgICQ5wYIECBAgAABAgQIECBAgAABAgQIBASEvMCSjEiAAAECBAgQIECAAAECBAgQIEBAyHMDBAgQIECAAAECBAgQIECAAAECBAICQl5gSUYkQIAAAQIECBAgQIAAAQIECBAgIOS5AQIECBAgQIAAAQIECBAgQIAAAQIBASEvsCQjEiBAgAABAgQIECBAgAABAgQIEBDy3AABAgQIECBAgAABAgQIECBAgACBgICQF1iSEQkQIECAAAECBAgQIECAAAECBAgIeW6AAAECBAgQIECAAAECBAgQIECAQEBAyAssyYgECBAgQIAAAQIECBAgQIAAAQIEhDw3QIAAAQIECBAgQIAAAQIECBAgQCAgIOQFlmREAgQIECBAgAABAgQIECBAgAABAkKeGyBAgAABAgQIECBAgAABAgQIECAQEBDyAksyIgECBAgQIECAAAECBAgQIECAAAEhzw0QIECAAAECBAgQIECAAAECBAgQCAgIeYElGZEAAQIECBAgQIAAAQIECBAgQICAkOcGCBAgQIAAAQIECBAgQIAAAQIECAQEhLzAkoxIgAABAgQIECBAgAABAgQIECBAQMhzAwQIECBAgAABAgQIECBAgAABAgQCAkJeYElGJECAAAECBAgQIECAAAECBAgQICDkuQECBAgQIECAAAECBAgQIECAAAECAQEhL7AkIxIgQIAAAQIECBAgQIAAAQIECBAQ8twAAQIECBAgQIAAAQIECBAgQIAAgYCAkBdYkhEJECBAgAABAgQIECBAgAABAgQICHlugAABAgQIECBAgAABAgQIECBAgEBAQMgLLMmIBAgQIECAAAECBAgQIECAAAECBIQ8N0CAAAECBAgQIECAAAECBAgQIEAgICDkBZZkRAIECBAgQIAAAQIECBAgQIAAAQJCnhsgQIAAAQIECBAgQIAAAQIECBAgEBAQ8gJLMiIBAgQIECBAgAABAgQIECBAgAABIc8NECBAgAABAgQIECBAgAABAgQIEAgICHmBJRmRAAECBAgQIECAAAECBAgQIECAgJDnBggQIECAAAECBAgQIECAAAECBAgEBIS8wJKMSIAAAQIECBAgQIAAAQIECBAgQEDIcwMECBAgQIAAAQIECBAgQIAAAQIEAgJCXmBJRiRAgAABAgQIECBAgAABAgQIECAg5LkBAgQIECBAgAABAgQIECBAgAABAgEBIS+wJCMSIECAAAECBAgQIECAAAECBAgQEPLcAAECBAgQIECAAAECBAgQIECAAIGAgJAXWJIRCRAgQIAAAQIECBAgQIAAAQIECAh5boAAAQIECBAgQIAAAQIECBAgQIBAQEDICyzJiAQIECBAgAABAgQIECBAgAABAgSEPDdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQEPICSzIiAQIECBAgQIAAAQIECBAgQIAAASHPDRAgQIAAAQIECBAgQIAAAQIECBAICAh5gSUZkQABAgQIECBAgAABAgQIECBAgMCZkIeCAAECBAgQIECAAAECBAgQIECAAIErW2DdlT2e6QgQIECAAAECBAgQIECAAAECBAgQmASEPHdAgAABAgQIECBAgAABAgQIECBAICAg5AWWZEQCBAgQIECAAAECBAgQIECAAAECQp4bIECAAAECBAgQIECAAAECBAgQIBAQ+D/tyO+BuOA5SAAAAABJRU5ErkJggg==', '根链1', '{\"lfData\":{\"globalColor\":\"#D49BEF\",\"dataCode\":{\"nodes\":[{\"id\":\"input\",\"type\":\"InputNode\",\"x\":116,\"y\":337,\"properties\":{\"debugMode\":false},\"zIndex\":1013,\"text\":{\"x\":126,\"y\":337,\"value\":\"输入\"}},{\"id\":\"b6365013-18f2-4361-85ed-d9db4b8144b5\",\"type\":\"SaveAttributesNode\",\"x\":702,\"y\":238,\"properties\":{\"debugMode\":false,\"name\":\"\"},\"zIndex\":1012,\"text\":{\"x\":712,\"y\":238,\"value\":\"保存参数\"}},{\"id\":\"0c2710b5-9714-4563-944c-8b1a78536814\",\"type\":\"SaveTimeSeriesNode\",\"x\":700,\"y\":437,\"properties\":{\"debugMode\":false,\"name\":\"\"},\"zIndex\":1014,\"text\":{\"x\":710,\"y\":437,\"value\":\"保存遥测\"}},{\"id\":\"74b514b9-9591-4ab2-b6f9-c6f2f005047f\",\"type\":\"MessageTypeSwitchNode\",\"x\":386,\"y\":332,\"properties\":{\"debugMode\":false},\"zIndex\":1002,\"text\":{\"x\":396,\"y\":332,\"value\":\"消息类型切换\"}}],\"edges\":[{\"id\":\"ba8d0084-9b85-437e-be3b-d83c644e8e09\",\"type\":\"bezier-link\",\"sourceNodeId\":\"input\",\"targetNodeId\":\"74b514b9-9591-4ab2-b6f9-c6f2f005047f\",\"startPoint\":{\"x\":176,\"y\":337},\"endPoint\":{\"x\":316,\"y\":332},\"properties\":{\"lineType\":[\"True\"]},\"text\":{\"x\":246,\"y\":334.5,\"value\":\"True\"},\"zIndex\":1003,\"pointsList\":[{\"x\":176,\"y\":337},{\"x\":276,\"y\":337},{\"x\":216,\"y\":332},{\"x\":316,\"y\":332}]},{\"id\":\"13cf5637-f995-4b56-9c41-530040418cdd\",\"type\":\"bezier-link\",\"sourceNodeId\":\"74b514b9-9591-4ab2-b6f9-c6f2f005047f\",\"targetNodeId\":\"0c2710b5-9714-4563-944c-8b1a78536814\",\"startPoint\":{\"x\":456,\"y\":332},\"endPoint\":{\"x\":640,\"y\":437},\"properties\":{\"lineType\":[\"Telemetry\"]},\"text\":{\"x\":548,\"y\":384.5,\"value\":\"Telemetry\"},\"zIndex\":1005,\"pointsList\":[{\"x\":456,\"y\":332},{\"x\":556,\"y\":332},{\"x\":540,\"y\":437},{\"x\":640,\"y\":437}]},{\"id\":\"0117eae6-8d1b-4eeb-96d6-6c42cddba1b4\",\"type\":\"bezier-link\",\"sourceNodeId\":\"74b514b9-9591-4ab2-b6f9-c6f2f005047f\",\"targetNodeId\":\"b6365013-18f2-4361-85ed-d9db4b8144b5\",\"startPoint\":{\"x\":456,\"y\":332},\"endPoint\":{\"x\":642,\"y\":238},\"properties\":{\"lineType\":[\"Attributes\"]},\"text\":{\"x\":549,\"y\":285,\"value\":\"Attributes\"},\"zIndex\":1006,\"pointsList\":[{\"x\":456,\"y\":332},{\"x\":556,\"y\":332},{\"x\":542,\"y\":238},{\"x\":642,\"y\":238}]}]},\"openRule\":false,\"setting\":{\"describe\":\"\",\"grid\":{\"size\":20,\"open\":false,\"type\":\"mesh\",\"config\":{\"color\":\"#cccccc\",\"thickness\":1}},\"backgroundColor\":\"#ffffff\"}}}');
INSERT INTO "public"."rule_chain" VALUES ('rule_a37571bb6c45378b57803793', 'panda', 2, '2023-07-21 16:17:51', '2023-08-03 11:08:22', '0', '高温告警规则', 'data:image/png;base64,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KE/Z/PCHP6xKKRs6XXTRRRnKUgMBAgTes4AVee+ZsOMBBHkNxK0bWpDXPc694SolyJs/f35MmjSp+tGVRzM9WtuVdRlrq0CzB3n1fXr55ZfjmWee2bYTYO33hg8fHlOnTo3x48drK4G0Arfeeuu2TY3e//73V6/icBAg0JwCq1atiu9973vV5Pbff//41V/91eacqFkRIECgmwUEeQ0EF+Q1ELdNkDd0xMfixBNP7J4LukpagTlz5nR5gFfDEuSlvW1aLsirdeL555+Pp59+eluoUfv62WefLczLfbv26upKiP3ggw9WBmUTmMsuuyz69Cnrgh0ECDSbwJtvvhnXX399Na2yevzKK69stimaDwECBHpEoFcEebNmzYqVr9wbEw5b3K3Igrzu4X72xdEx7KBzomxL7yDQrAKCvGbtTNfMq6zI6ztqckt+Q2HevHnx1FNPRVn5IMzrmvvBKM0rsHHjxrjpppvijTfeqCb54Q9/OI466qjmnbCZEejFAmVzmm984xuVQL9+/eJzn/tcL9ZQOgECBLYL9Iog74UXXog5M++KyROf69beL1uxLvoO2hwHHbhPt163t13siaePikmTP2YFSW9rfIvVW76h8PaLT8aUMd37DrLFy1dHv722xOgR+7eYWGtNd+bCTTFgzAkt+w2FNWvWVDsDLl26dafzcliZ11r3oNl2XqD+fahlg6zzzjuv8x92JgEC3SYgyOs2ahciQKDFBHpFkLdp06b49rf+OT4w6fnYe6+3uq1Fy1e/EX36bxHkNVB83fpB8fDjR8VVV302vKS9gdCGfs8C5RsK8x59KM4a172PcC1etTr6DRTkvecG7mKAu5/bEhOnnN7S31AojzCVMG/RokXbqr3gggti9OjRjeYzPoFuFSjBdVmVV0KCcnz84x+PQw45pFvn4GIECOxa4K233oprr722OrFsUHP11Vfv+kPOIECAQC8Q6BVBXuljWQ2z4NlHum1VXp++Ectff726hazIa9yfpLIab+y4U1t2FUzjZIzcbALlGwrXfevauOCogbHfoO4L85asWRV9+4cVeQ28IdZs2BI/eHpjfOaqq1v+GwrlPr3//vujbIhRjhJulJDDQYAAAQIEultg7dq1ccMNN1SX3WeffeKKK67o7im4HgECBJpSoNcEeUX/x/fdFW+sfCGOHPtKQ1fmlRCvz4DNUR6t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x9wJNJ3A3XffHS+//HI1r4985CNx+OGHN90cTYgAAQI9KSDI60l91yZAgAABAk0ocM0112yb1ec///kmnKEpEehZgWuvvTbeeuutahKf/OQnY7/99uvZCbk6gUQC9a93KDtDlx2iHQQIECCwXUCQ524gQIAAAQIEdhAou3KW3TnL8YlPfCKGDRtGiACBOoHbb789Fi1aVH3lox/9aIwbN44PAQJdJFC/oczFF18cBx54YBeNbBgCBAjkEBDk5eijKggQIECAQJcJ/OAHP4ilS5dW49m1tstYDZRI4Oc//3k89dRTVUVTp06NadOmJapOKQR6VuC73/1urF69upqEbyb1bC9cnQCB5hQQ5DVnX8yKAAECBAj0mMC9994bL774opCixzrgws0uUB/kfeADH4jjjz++2adsfgRaRqB+VfinP/3p2HfffVtm7iZKgACB7hAQ5HWHsmsQIECAAIEWEnjooYfi6aefrmY8fvz4alWegwCB7QI//elP49/+7d+qL3zwgx+MY445Bg8BAl0kUP8OyquvvjoGDRrURSMbhgABAjkEBHk5+qgKAgQIECDQZQLPPPNMPPjgg9V4w4cPj8suu6zLxjYQgQwC9913X7zwwgtVKWeddVZMmDAhQ1lqINAUAt/4xjdi8+bN1Vw+97nPRb9+/ZpiXiZBgACBZhEQ5DVLJ8yDAAECBAg0iUDZ6KI82lQ7LrnkkhgxYkSTzM40CPS8wB133BELFy6sJnLuuefG2LFje35SZkAgiYAgL0kjlUGAQMMEBHkNozUwAQIECBBoXYHbbrstXn311aqA4447Lk499dSWK2bTpk2xYMGCWLNmTbzzzjs7nf+WLVtarr7MEy4rcIYOHVoFZP3792+6Uq+77rpYv359Na/zzz8/xowZ03RzNCECrSpw/fXXx5tvvllN/8orr4y99967VUsxbwIECDREQJDXEFaDEiBAgACB1haYPXt2PProo1URJUi59NJLq2ClVY5Zs2bF7JlPxKghA2P4wM3RJ7owqOsbEX23RJ9/f9qrT6ugtNA834k+sWJDv1i8ZmOcMGVqTJkypWlm/9prr8WPfvSjaj4lYChBg4MAga4TsGtt11kaiQCBnAKCvJx9VRUBAgQIEHhPAitWrKjCirfffrsaZ/LkyXHyySe/pzG768Mz7r0nNqxYEieNjthvUONithLk9e2/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5QQr8dwQTScgyGu6lpgQAQIECBAgQIAAAQIECLSyQHln3oIFC2LNmjWxefPmVi4l3dz79Nl5sNq3b98YOnRojB071jvx0nU/R0GCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjoLq0AABUoSURBVAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIeAIC9HH1VBgAABAgQIECBAgAABAgQIECCQXECQl7zByiNAgAABAgQIECBAgAABAgQIEMghIMjL0UdVECBAgAABAgQIECBAgAABAgQIJBcQ5CVvsPIIECBAgAABAgQIECBAgAABAgRyCAjycvRRFQQIECBAgAABAgQIECBAgAABAskFBHnJG6w8AgQIECBAgAABAgQIECBAgACBHAKCvBx9VAUBAgQIECBAgAABAgQIECBAgEByAUFe8gYrjwABAgQIECBAgAABAgQIECBAIIfAtiAvRzmqIECAAAECBAgQIECAAAECBAgQIJBXoE/e0lRGgAABAgQIECBAgAABAgQIECBAII+AIC9PL1VCgAABAgQIECBAgAABAgQIECCQWECQl7i5SiNAgAABAgQIECBAgAABAgQIEMgjIMjL00uVECBAgAABAgQIECBAgAABAgQIJBb436ya7dpojYnCAAAAAElFTkSuQmCC', '根链1', '{\"lfData\":{\"globalColor\":\"#D49BEF\",\"dataCode\":{\"nodes\":[{\"id\":\"input\",\"type\":\"InputNode\",\"x\":116,\"y\":337,\"properties\":{\"debugMode\":false},\"zIndex\":1013,\"text\":{\"x\":126,\"y\":337,\"value\":\"输入\"}},{\"id\":\"b6365013-18f2-4361-85ed-d9db4b8144b5\",\"type\":\"SaveAttributesNode\",\"x\":624,\"y\":249,\"properties\":{\"debugMode\":false,\"name\":\"\"},\"zIndex\":1012,\"text\":{\"x\":634,\"y\":249,\"value\":\"保存参数\"}},{\"id\":\"0c2710b5-9714-4563-944c-8b1a78536814\",\"type\":\"SaveTimeSeriesNode\",\"x\":624,\"y\":409,\"properties\":{\"debugMode\":false,\"name\":\"\"},\"zIndex\":1014,\"text\":{\"x\":634,\"y\":409,\"value\":\"保存遥测\"}},{\"id\":\"74b514b9-9591-4ab2-b6f9-c6f2f005047f\",\"type\":\"MessageTypeSwitchNode\",\"x\":406,\"y\":338,\"properties\":{\"debugMode\":false},\"zIndex\":1002,\"text\":{\"x\":416,\"y\":338,\"value\":\"消息类型切换\"}},{\"id\":\"594f9c98-daaf-4348-a4c7-208feb413ff8\",\"type\":\"ScriptFilterNode\",\"x\":813,\"y\":308,\"properties\":{\"debugMode\":false,\"name\":\"验证温度大于20\",\"script\":\"return msg.temperature > 20;\"},\"zIndex\":1015,\"text\":{\"x\":823,\"y\":308,\"value\":\"验证温度大于20\"}},{\"id\":\"62882142-b992-490b-ad01-bd5f965c8570\",\"type\":\"CreateAlarmNode\",\"x\":1049,\"y\":212,\"properties\":{\"debugMode\":false,\"name\":\"创建设备告警信息\",\"alarmType\":\"高温报警\",\"alarmSeverity\":\"MAJOR\"},\"zIndex\":1002,\"text\":{\"x\":1059,\"y\":212,\"value\":\"创建设备告警信息\"}},{\"id\":\"b14a20cc-0369-4f91-8157-c98a25c19a04\",\"type\":\"ClearAlarmNode\",\"x\":1041,\"y\":431,\"properties\":{\"debugMode\":false,\"name\":\"清除告警\",\"alarmType\":\"高温报警\"},\"zIndex\":1006,\"text\":{\"x\":1051,\"y\":431,\"value\":\"清除告警\"}}],\"edges\":[{\"id\":\"ba8d0084-9b85-437e-be3b-d83c644e8e09\",\"type\":\"bezier-link\",\"sourceNodeId\":\"input\",\"targetNodeId\":\"74b514b9-9591-4ab2-b6f9-c6f2f005047f\",\"startPoint\":{\"x\":176,\"y\":337},\"endPoint\":{\"x\":336,\"y\":338},\"properties\":{\"lineType\":[\"True\"]},\"text\":{\"x\":256,\"y\":337.5,\"value\":\"True\"},\"zIndex\":1003,\"pointsList\":[{\"x\":176,\"y\":337},{\"x\":276,\"y\":337},{\"x\":236,\"y\":338},{\"x\":336,\"y\":338}]},{\"id\":\"13cf5637-f995-4b56-9c41-530040418cdd\",\"type\":\"bezier-link\",\"sourceNodeId\":\"74b514b9-9591-4ab2-b6f9-c6f2f005047f\",\"targetNodeId\":\"0c2710b5-9714-4563-944c-8b1a78536814\",\"startPoint\":{\"x\":476,\"y\":338},\"endPoint\":{\"x\":564,\"y\":409},\"properties\":{\"lineType\":[\"Telemetry\"]},\"text\":{\"x\":520,\"y\":373.5,\"value\":\"Telemetry\"},\"zIndex\":1005,\"pointsList\":[{\"x\":476,\"y\":338},{\"x\":576,\"y\":338},{\"x\":464,\"y\":409},{\"x\":564,\"y\":409}]},{\"id\":\"0117eae6-8d1b-4eeb-96d6-6c42cddba1b4\",\"type\":\"bezier-link\",\"sourceNodeId\":\"74b514b9-9591-4ab2-b6f9-c6f2f005047f\",\"targetNodeId\":\"b6365013-18f2-4361-85ed-d9db4b8144b5\",\"startPoint\":{\"x\":476,\"y\":338},\"endPoint\":{\"x\":564,\"y\":249},\"properties\":{\"lineType\":[\"Attributes\"]},\"text\":{\"x\":520,\"y\":293.5,\"value\":\"Attributes\"},\"zIndex\":1006,\"pointsList\":[{\"x\":476,\"y\":338},{\"x\":576,\"y\":338},{\"x\":464,\"y\":249},{\"x\":564,\"y\":249}]},{\"id\":\"5c2bc46c-10c8-4a96-a1c2-3f50f19777e4\",\"type\":\"bezier-link\",\"sourceNodeId\":\"0c2710b5-9714-4563-944c-8b1a78536814\",\"targetNodeId\":\"594f9c98-daaf-4348-a4c7-208feb413ff8\",\"startPoint\":{\"x\":684,\"y\":409},\"endPoint\":{\"x\":733,\"y\":308},\"properties\":{\"lineType\":[\"Success\"]},\"text\":{\"x\":708.5,\"y\":358.5,\"value\":\"Success\"},\"zIndex\":1016,\"pointsList\":[{\"x\":684,\"y\":409},{\"x\":784,\"y\":409},{\"x\":633,\"y\":308},{\"x\":733,\"y\":308}]},{\"id\":\"9a80d1c4-b0f6-4a57-89a4-af4fa9d459c4\",\"type\":\"bezier-link\",\"sourceNodeId\":\"594f9c98-daaf-4348-a4c7-208feb413ff8\",\"targetNodeId\":\"62882142-b992-490b-ad01-bd5f965c8570\",\"startPoint\":{\"x\":893,\"y\":308},\"endPoint\":{\"x\":969,\"y\":212},\"properties\":{\"lineType\":[\"True\"]},\"text\":{\"x\":931,\"y\":260,\"value\":\"True\"},\"zIndex\":1003,\"pointsList\":[{\"x\":893,\"y\":308},{\"x\":993,\"y\":308},{\"x\":869,\"y\":212},{\"x\":969,\"y\":212}]},{\"id\":\"2f659b99-c508-4e3e-a98c-fe87b9a79da1\",\"type\":\"bezier-link\",\"sourceNodeId\":\"594f9c98-daaf-4348-a4c7-208feb413ff8\",\"targetNodeId\":\"b14a20cc-0369-4f91-8157-c98a25c19a04\",\"startPoint\":{\"x\":893,\"y\":308},\"endPoint\":{\"x\":981,\"y\":431},\"properties\":{\"lineType\":[\"False\"]},\"text\":{\"x\":937,\"y\":369.5,\"value\":\"False\"},\"zIndex\":1007,\"pointsList\":[{\"x\":893,\"y\":308},{\"x\":993,\"y\":308},{\"x\":881,\"y\":431},{\"x\":981,\"y\":431}]}]},\"openRule\":false,\"setting\":{\"describe\":\"\",\"grid\":{\"size\":20,\"open\":false,\"type\":\"mesh\",\"config\":{\"color\":\"#cccccc\",\"thickness\":1}},\"backgroundColor\":\"#ffffff\"}}}');

-- ----------------------------
-- Table structure for rule_chain_msg_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."rule_chain_msg_log";
CREATE TABLE "public"."rule_chain_msg_log" (
  "message_id" varchar(191) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "msg_type" varchar(191) COLLATE "pg_catalog"."default",
  "device_id" varchar(64) COLLATE "pg_catalog"."default",
  "device_name" varchar(191) COLLATE "pg_catalog"."default",
  "ts" timestamp(6),
  "content" varchar(191) COLLATE "pg_catalog"."default",
  "create_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."rule_chain_msg_log"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."rule_chain_msg_log"."owner" IS '创建者,所有者';

-- ----------------------------
-- Records of rule_chain_msg_log
-- ----------------------------
INSERT INTO "public"."rule_chain_msg_log" VALUES ('3453d', 2, 'panda', 'Telemetry', 'd_1928b99619910dae5a001fa7', 'ws432', '2023-07-31 14:23:13', 'Incoming message', '2023-09-06 15:28:45');

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_apis";
CREATE TABLE "public"."sys_apis" (
  "id" int8 NOT NULL,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6),
  "path" varchar(191) COLLATE "pg_catalog"."default",
  "description" varchar(191) COLLATE "pg_catalog"."default",
  "api_group" varchar(191) COLLATE "pg_catalog"."default",
  "method" varchar(191) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_apis"."path" IS 'api路径';
COMMENT ON COLUMN "public"."sys_apis"."description" IS 'api中文描述';
COMMENT ON COLUMN "public"."sys_apis"."api_group" IS 'api组';
COMMENT ON COLUMN "public"."sys_apis"."method" IS '方法';

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO "public"."sys_apis" VALUES (1, '2021-12-09 09:21:04', '2021-12-09 09:21:04', NULL, '/system/user/list', '查询用户列表（分页）', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (2, '2021-12-09 09:29:36', '2021-12-09 09:29:36', NULL, '/system/user/changeStatus', '修改用户状态', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (3, '2021-12-09 09:34:37', '2021-12-09 09:34:37', NULL, '/system/user/:userId', '删除用户信息', 'user', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (4, '2021-12-09 09:36:43', '2023-09-14 14:05:54', NULL, '/system/organization/list', '获取组织列表', 'organization', 'GET');
INSERT INTO "public"."sys_apis" VALUES (5, '2021-12-09 09:37:31', '2023-09-14 14:06:51', NULL, '/system/organization/:organizationId', '获取组织信息', 'organization', 'GET');
INSERT INTO "public"."sys_apis" VALUES (6, '2021-12-09 18:20:32', '2021-12-09 18:20:32', NULL, '/system/user/avatar', '上传头像', 'user', 'POST');
INSERT INTO "public"."sys_apis" VALUES (7, '2021-12-09 18:21:10', '2021-12-09 18:21:10', NULL, '/system/user/pwd', '修改密码', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (8, '2021-12-09 18:21:54', '2021-12-09 18:21:54', NULL, '/system/user/getById/:userId', '通过id获取用户信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (9, '2021-12-09 18:58:50', '2021-12-09 18:58:50', NULL, '/system/user/getInit', '获取初始化角色岗位信息(添加用户初始化)', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (10, '2021-12-09 18:59:43', '2021-12-09 18:59:43', NULL, '/system/user/getRoPo', '获取用户角色岗位信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (11, '2021-12-09 19:00:24', '2021-12-09 19:00:24', NULL, '/system/user', '添加用户信息', 'user', 'POST');
INSERT INTO "public"."sys_apis" VALUES (12, '2021-12-09 19:00:52', '2021-12-09 19:00:52', NULL, '/system/user', '修改用户信息', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (13, '2021-12-09 19:02:30', '2021-12-09 19:02:30', NULL, '/system/user/export', '导出用户信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (14, '2021-12-09 19:04:04', '2023-09-14 14:06:35', NULL, '/system/organization/roleOrganizationTreeSelect/:roleId', '获取角色部门树', 'organization', 'GET');
INSERT INTO "public"."sys_apis" VALUES (15, '2021-12-09 19:04:48', '2023-09-14 14:07:06', NULL, '/system/organization/organizationTree', '获取所有组织树', 'organization', 'GET');
INSERT INTO "public"."sys_apis" VALUES (16, '2021-12-09 19:07:37', '2023-09-14 14:07:18', NULL, '/system/organization', '添加组织信息', 'organization', 'POST');
INSERT INTO "public"."sys_apis" VALUES (17, '2021-12-09 19:08:14', '2023-09-14 14:07:28', NULL, '/system/organization', '修改组织信息', 'organization', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (18, '2021-12-09 19:08:40', '2023-09-14 14:07:41', NULL, '/system/organization/:organizationId', '删除组织信息', 'organization', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (19, '2021-12-09 19:09:41', '2021-12-09 19:09:41', NULL, '/system/config/list', '获取配置分页列表', 'config', 'GET');
INSERT INTO "public"."sys_apis" VALUES (20, '2021-12-09 19:10:11', '2021-12-09 19:10:11', NULL, '/system/config/configKey', '获取配置列表通过ConfigKey', 'config', 'GET');
INSERT INTO "public"."sys_apis" VALUES (21, '2021-12-09 19:10:45', '2021-12-09 19:10:45', NULL, '/system/config/:configId', '获取配置信息', 'config', 'GET');
INSERT INTO "public"."sys_apis" VALUES (22, '2021-12-09 19:11:22', '2021-12-09 19:11:22', NULL, '/system/config', '添加配置信息', 'config', 'POST');
INSERT INTO "public"."sys_apis" VALUES (23, '2021-12-09 19:11:41', '2021-12-09 19:11:41', NULL, '/system/config', '修改配置信息', 'config', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (24, '2021-12-09 19:12:28', '2021-12-09 19:12:28', NULL, '/system/config/:configId', '删除配置信息', 'config', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (25, '2021-12-09 19:13:08', '2021-12-09 19:13:08', NULL, '/system/dict/type/list', '获取字典类型分页列表', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (26, '2021-12-09 19:13:55', '2021-12-09 19:13:55', NULL, '/system/dict/type/:dictId', '获取字典类型信息', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (27, '2021-12-09 19:14:28', '2021-12-09 19:14:28', NULL, '/system/dict/type', '添加字典类型信息', 'dict', 'POST');
INSERT INTO "public"."sys_apis" VALUES (28, '2021-12-09 19:14:55', '2021-12-09 19:14:55', NULL, '/system/dict/type', '修改字典类型信息', 'dict', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (29, '2021-12-09 19:15:17', '2021-12-09 19:15:17', NULL, '/system/dict/type/:dictId', '删除字典类型信息', 'dict', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (30, '2021-12-09 19:15:50', '2021-12-09 19:15:50', NULL, '/system/dict/type/export', '导出字典类型信息', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (31, '2021-12-09 19:16:26', '2021-12-09 19:16:26', NULL, '/system/dict/data/list', '获取字典数据分页列表', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (32, '2021-12-09 19:17:01', '2021-12-09 19:17:01', NULL, '/system/dict/data/type', '获取字典数据列表通过字典类型', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (33, '2021-12-09 19:17:39', '2021-12-09 19:17:39', NULL, '/system/dict/data/:dictCode', '获取字典数据信息', 'dict', 'GET');
INSERT INTO "public"."sys_apis" VALUES (34, '2021-12-09 19:18:20', '2021-12-09 19:18:20', NULL, '/system/dict/data', '添加字典数据信息', 'dict', 'POST');
INSERT INTO "public"."sys_apis" VALUES (35, '2021-12-09 19:18:44', '2021-12-09 19:18:44', NULL, '/system/dict/data', '修改字典数据信息', 'dict', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (36, '2021-12-09 19:19:16', '2021-12-09 19:19:16', NULL, '/system/dict/data/:dictCode', '删除字典数据信息', 'dict', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (37, '2021-12-09 19:21:18', '2021-12-09 19:21:18', NULL, '/system/menu/menuTreeSelect', '获取菜单树', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (38, '2021-12-09 19:21:47', '2021-12-09 19:21:47', NULL, '/system/menu/menuRole', '获取角色菜单', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (39, '2021-12-09 19:22:42', '2021-12-09 19:22:42', NULL, '/system/menu/roleMenuTreeSelect/:roleId', '获取角色菜单树', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (40, '2021-12-09 19:23:17', '2021-12-09 19:23:17', NULL, '/system/menu/menuPaths', '获取角色菜单路径列表', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (41, '2021-12-09 19:23:40', '2021-12-09 19:23:40', NULL, '/system/menu/list', '获取菜单列表', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (42, '2021-12-09 19:24:09', '2021-12-09 19:24:09', NULL, '/system/menu/:menuId', '获取菜单信息', 'menu', 'GET');
INSERT INTO "public"."sys_apis" VALUES (43, '2021-12-09 19:24:29', '2021-12-09 19:24:29', NULL, '/system/menu', '添加菜单信息', 'menu', 'POST');
INSERT INTO "public"."sys_apis" VALUES (44, '2021-12-09 19:24:48', '2021-12-09 19:24:48', NULL, '/system/menu', '修改菜单信息', 'menu', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (45, '2021-12-09 19:25:10', '2021-12-09 19:25:10', NULL, '/system/menu/:menuId', '删除菜单信息', 'menu', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (46, '2021-12-09 19:25:44', '2021-12-09 19:27:06', NULL, '/system/post/list', '获取岗位分页列表', 'post', 'GET');
INSERT INTO "public"."sys_apis" VALUES (47, '2021-12-09 19:26:55', '2021-12-09 19:26:55', NULL, '/system/post/:postId', '获取岗位信息', 'post', 'GET');
INSERT INTO "public"."sys_apis" VALUES (48, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/post', '添加岗位信息', 'post', 'POST');
INSERT INTO "public"."sys_apis" VALUES (49, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/post', '修改岗位信息', 'post', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (50, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/post/:postId', '删除岗位信息', 'post', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (51, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/list', '获取角色分页列表', 'role', 'GET');
INSERT INTO "public"."sys_apis" VALUES (52, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/:roleId', '获取角色信息', 'role', 'GET');
INSERT INTO "public"."sys_apis" VALUES (53, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role', '添加角色信息', 'role', 'POST');
INSERT INTO "public"."sys_apis" VALUES (54, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role', '修改角色信息', 'role', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (55, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/:roleId', '删除角色信息', 'role', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (56, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/changeStatus', '修改角色状态', 'role', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (57, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/dataScope', '修改角色部门权限', 'role', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (58, '2021-12-09 19:25:44', '2021-12-09 19:25:44', NULL, '/system/role/export', '导出角色信息', 'role', 'GET');
INSERT INTO "public"."sys_apis" VALUES (59, '2021-12-09 19:50:57', '2022-01-19 08:58:20', NULL, '/system/api/list', '获取api分页列表1', 'api', 'GET');
INSERT INTO "public"."sys_apis" VALUES (60, '2021-12-09 19:51:26', '2021-12-09 19:51:26', NULL, '/system/api/all', '获取所有api', 'api', 'GET');
INSERT INTO "public"."sys_apis" VALUES (61, '2021-12-09 19:51:54', '2021-12-09 19:51:54', NULL, '/system/api/getPolicyPathByRoleId', '获取角色拥有的api权限', 'api', 'GET');
INSERT INTO "public"."sys_apis" VALUES (62, '2021-12-09 19:52:14', '2021-12-09 19:52:14', NULL, '/system/api/:id', '获取api信息', 'api', 'GET');
INSERT INTO "public"."sys_apis" VALUES (63, '2021-12-09 19:52:35', '2021-12-09 19:52:35', NULL, '/system/api', '添加api信息', 'api', 'POST');
INSERT INTO "public"."sys_apis" VALUES (64, '2021-12-09 19:52:50', '2021-12-09 19:52:50', NULL, '/system/api', '修改api信息', 'api', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (65, '2021-12-09 19:53:07', '2021-12-09 19:53:07', NULL, '/system/api/:id', '删除api信息', 'api', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (66, '2021-12-17 10:51:05', '2021-12-17 10:54:22', NULL, '/log/logLogin/list', '获取登录日志', 'log', 'GET');
INSERT INTO "public"."sys_apis" VALUES (67, '2021-12-17 10:51:43', '2021-12-17 10:54:28', NULL, '/log/logLogin/:infoId', '删除日志', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (68, '2021-12-17 10:53:09', '2021-12-17 10:54:34', NULL, '/log/logLogin/all', '清空所有', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (69, '2021-12-17 10:54:07', '2021-12-17 10:54:07', NULL, '/log/logOper/list', '操作日志列表', 'log', 'GET');
INSERT INTO "public"."sys_apis" VALUES (70, '2021-12-17 10:53:09', '2021-12-17 10:53:09', NULL, '/log/logOper/:operId', '删除', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (71, '2021-12-17 10:53:09', '2021-12-17 10:53:09', NULL, '/log/logOper/all', '清空', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (72, '2021-12-24 15:41:23', '2021-12-24 15:41:23', NULL, '/job/list', '任务列表', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (73, '2021-12-24 15:41:54', '2021-12-24 15:41:54', NULL, '/job', '添加', 'job', 'POST');
INSERT INTO "public"."sys_apis" VALUES (74, '2021-12-24 15:42:11', '2021-12-24 15:42:11', NULL, '/job', '修改任务', 'job', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (75, '2021-12-24 15:42:37', '2021-12-24 16:32:21', NULL, '/job/:jobId', '获取任务', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (76, '2021-12-24 15:43:09', '2021-12-24 16:32:05', NULL, '/job/:jobId', '删除job', 'job', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (77, '2021-12-24 15:43:35', '2021-12-24 16:31:11', NULL, '/job/stop/:jobId', '停止任务', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (78, '2021-12-24 15:44:09', '2021-12-24 16:30:38', NULL, '/job/start/:jobId', '开始任务', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (79, '2021-12-24 15:45:03', '2023-08-08 14:15:59', NULL, '/job/log/list', '任务日志列表', 'job', 'GET');
INSERT INTO "public"."sys_apis" VALUES (80, '2021-12-24 15:45:33', '2023-08-08 14:16:07', NULL, '/job/log/all', '清空任务日志', 'job', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (81, '2021-12-24 15:46:08', '2023-08-08 14:16:15', NULL, '/job/log/:logId', '删除任务日志', 'job', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (82, '2021-12-24 15:45:33', '2021-12-24 15:45:33', NULL, '/system/notice/list', '获取通知分页列表', 'notice', 'GET');
INSERT INTO "public"."sys_apis" VALUES (83, '2021-12-24 15:45:33', '2021-12-24 15:45:33', NULL, '/system/notice', '添加通知信息', 'notice', 'POST');
INSERT INTO "public"."sys_apis" VALUES (84, '2021-12-24 15:45:33', '2021-12-24 15:45:33', NULL, '/system/notice', '修改通知信息', 'notice', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (85, '2021-12-24 15:45:33', '2021-12-24 16:33:48', NULL, '/system/notice/:noticeId', '删除通知信息', 'notice', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (86, '2021-12-24 22:40:19', '2021-12-24 22:40:19', NULL, '/job/changeStatus', '修改状态', 'job', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (88, '2022-01-02 13:53:06', '2022-07-18 10:57:58', NULL, '/develop/code/table/db/list', '数据库表列表', 'gen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (89, '2022-01-02 13:53:44', '2022-01-02 13:53:44', NULL, '/develop/code/table/list', '表列表', 'gen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (90, '2022-01-02 13:54:10', '2022-01-02 13:54:10', NULL, '/develop/code/table/info/:tableId', '表信息', 'gen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (91, '2022-01-02 13:54:42', '2022-07-18 10:58:35', NULL, '/develop/code/table/info/tableName', '表名获取表信息', 'gen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (92, '2022-01-02 13:55:13', '2022-01-02 13:55:13', NULL, '/develop/code/table/tableTree', '表树', 'gen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (93, '2022-01-02 13:56:37', '2022-01-02 13:56:37', NULL, '/develop/code/table', '导入表', 'gen', 'POST');
INSERT INTO "public"."sys_apis" VALUES (94, '2022-01-02 13:57:36', '2022-01-02 13:57:36', NULL, '/develop/code/table', '修改代码生成信息', 'gen', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (95, '2022-01-02 13:58:25', '2022-01-02 13:58:25', NULL, '/develop/code/table/:tableId', '删除表数据', 'gen', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (96, '2022-01-02 13:59:07', '2022-01-02 13:59:07', NULL, '/develop/code/gen/preview/:tableId', '预览代码', 'gen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (97, '2022-01-02 13:59:43', '2022-01-02 13:59:43', NULL, '/develop/code/gen/code/:tableId', '生成代码', 'gen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (98, '2022-01-02 14:00:10', '2022-07-17 01:19:42', NULL, '/develop/code/gen/configure/:tableId', '生成api菜单', 'gen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (124, '2023-06-29 16:59:08', '2023-06-29 17:00:17', NULL, '/device/product/category/list', '获取产品分类列表', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (125, '2023-06-29 17:00:08', '2023-06-29 17:00:08', NULL, '/device/product/category/list/all', '获取所有列表', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (126, '2023-06-29 17:00:56', '2023-06-29 17:00:56', NULL, '/device/product/category/list/tree', '获取树', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (127, '2023-06-29 17:01:44', '2023-06-29 17:01:44', NULL, '/device/product/category/:id', '查询单个', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (128, '2023-06-29 17:02:16', '2023-06-29 17:02:16', NULL, '/device/product/category', '添加分类', 'product', 'POST');
INSERT INTO "public"."sys_apis" VALUES (129, '2023-06-29 17:02:42', '2023-06-29 17:02:42', NULL, '/device/product/category', '修改分类', 'product', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (130, '2023-06-29 17:03:07', '2023-06-29 17:03:07', NULL, '/device/product/category/:id', '删除分类', 'product', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (131, '2023-06-29 16:59:08', '2023-06-29 17:00:17', NULL, '/device/group/list', '获取设备分组列表', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (132, '2023-06-29 17:00:08', '2023-06-29 17:00:08', NULL, '/device/group/list/all', '获取所有列表', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (133, '2023-06-29 17:00:56', '2023-06-29 17:00:56', NULL, '/device/group/list/tree', '获取树', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (134, '2023-06-29 17:01:44', '2023-06-29 17:01:44', NULL, '/device/group/:id', '查询单个', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (135, '2023-06-29 17:02:16', '2023-06-29 17:02:16', NULL, '/device/group', '添加分组', 'device', 'POST');
INSERT INTO "public"."sys_apis" VALUES (136, '2023-06-29 17:02:42', '2023-06-29 17:02:42', NULL, '/device/group', '修改分组', 'device', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (137, '2023-06-29 17:03:07', '2023-06-29 17:03:07', NULL, '/device/group/:id', '删除分组', 'device', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (138, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product/:id', '删除产品信息', 'product', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (139, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product/:id', '获取产品信息', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (140, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product', '修改产品信息', 'product', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (141, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product/list', '查询产品列表（分页）', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (142, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product', '添加产品信息', 'product', 'POST');
INSERT INTO "public"."sys_apis" VALUES (143, '2023-06-30 14:20:03', '2023-06-30 15:26:46', NULL, '/device/list', '查询设备列表（分页）', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (144, '2023-06-30 14:20:03', '2023-06-30 15:26:52', NULL, '/device/:id', '获取设备信息', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (145, '2023-06-30 14:20:03', '2023-06-30 15:26:57', NULL, '/device', '添加设备信息', 'device', 'POST');
INSERT INTO "public"."sys_apis" VALUES (146, '2023-06-30 14:20:03', '2023-06-30 15:27:04', NULL, '/device/:id', '删除设备信息', 'device', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (147, '2023-06-30 14:20:03', '2023-06-30 15:27:09', NULL, '/device', '修改设备信息', 'device', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (148, '2023-06-30 15:11:25', '2023-08-02 16:06:13', NULL, '/device/group/list/tree/label', '获取设备组label', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (149, '2023-06-30 15:14:08', '2023-09-22 16:58:04', NULL, '/device/product/category/list/tree/label', '获取设置分类树', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (150, '2023-07-01 10:44:25', '2023-07-01 10:44:25', NULL, '/upload/up/oss', '上传文件到oss', 'upload', 'POST');
INSERT INTO "public"."sys_apis" VALUES (151, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota/list', '查询产品固件列表（分页）', 'ota', 'GET');
INSERT INTO "public"."sys_apis" VALUES (152, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota', '添加产品固件信息', 'ota', 'POST');
INSERT INTO "public"."sys_apis" VALUES (153, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota', '修改产品固件信息', 'ota', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (154, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota/:id', '删除产品固件信息', 'ota', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (155, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota/:id', '获取产品固件信息', 'ota', 'GET');
INSERT INTO "public"."sys_apis" VALUES (156, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template/list', '查询产品模型列表（分页）', 'template', 'GET');
INSERT INTO "public"."sys_apis" VALUES (157, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template', '修改产品模型信息', 'template', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (158, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template/:id', '获取产品模型信息', 'template', 'GET');
INSERT INTO "public"."sys_apis" VALUES (159, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template/:id', '删除产品模型信息', 'template', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (160, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template', '添加产品模型信息', 'template', 'POST');
INSERT INTO "public"."sys_apis" VALUES (161, '2023-07-07 16:35:45', '2023-07-07 16:35:45', NULL, '/device/product/list/all', '获取所有列表', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (162, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen', '修改bi大屏信息', 'screen', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (163, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen/:screenId', '获取bi大屏信息', 'screen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (164, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen/list', '查询bi大屏列表（分页）', 'screen', 'GET');
INSERT INTO "public"."sys_apis" VALUES (165, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen/:screenId', '删除bi大屏信息', 'screen', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (166, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen', '添加bi大屏信息', 'screen', 'POST');
INSERT INTO "public"."sys_apis" VALUES (167, '2023-04-13 10:15:27', '2023-04-13 10:15:27', NULL, '/visual/screen/group/list', '大屏分组列表', 'screen_group', 'GET');
INSERT INTO "public"."sys_apis" VALUES (168, '2023-04-13 10:16:15', '2023-04-13 10:16:15', NULL, '/visual/screen/group/list/tree', '大屏分组列表树', 'screen_group', 'GET');
INSERT INTO "public"."sys_apis" VALUES (169, '2023-04-13 10:16:38', '2023-04-13 10:16:38', NULL, '/visual/screen/group/list/all', '获取所有分组', 'screen_group', 'GET');
INSERT INTO "public"."sys_apis" VALUES (170, '2023-04-13 10:17:34', '2023-04-13 10:17:34', NULL, '/visual/screen/group/:id', '获取分组', 'screen_group', 'GET');
INSERT INTO "public"."sys_apis" VALUES (171, '2023-04-13 10:18:10', '2023-04-13 10:18:10', NULL, '/visual/screen/group', '添加分组', 'screen_group', 'POST');
INSERT INTO "public"."sys_apis" VALUES (172, '2023-04-13 10:18:35', '2023-04-13 10:18:35', NULL, '/visual/screen/group', '修改分组', 'screen_group', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (173, '2023-04-13 10:19:09', '2023-04-13 10:19:09', NULL, '/visual/screen/group/:id', '删除分组', 'screen_group', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (174, '2023-04-13 15:49:39', '2023-04-13 15:49:39', NULL, '/visual/screen/changeStatus', '改变状态', 'screen', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (175, '2023-04-13 15:50:18', '2023-07-21 17:44:48', NULL, '/rule/chain/changeRoot', '改变规则链', 'rulechain', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (176, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain/list', '查询规则链列表（分页）', 'rulechain', 'GET');
INSERT INTO "public"."sys_apis" VALUES (177, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain/:ruleId', '删除规则链信息', 'rulechain', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (178, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain', '修改规则链信息', 'rulechain', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (179, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain', '添加规则链信息', 'rulechain', 'POST');
INSERT INTO "public"."sys_apis" VALUES (180, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain/:ruleId', '获取规则链信息', 'rulechain', 'GET');
INSERT INTO "public"."sys_apis" VALUES (181, '2023-07-24 11:51:10', '2023-07-24 11:51:10', NULL, '/rule/chain/list/label', '获取规则链label列表', 'rulechain', 'GET');
INSERT INTO "public"."sys_apis" VALUES (182, '2023-07-31 14:14:06', '2023-07-31 14:14:06', NULL, '/device/list/all', '获取所有设备', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (183, '2023-08-02 16:05:24', '2023-08-02 16:05:24', NULL, '/device/:id/status', '获取设备状态', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (184, '2023-08-03 09:50:41', '2023-08-03 09:50:41', NULL, '/rule/chain/clone/:ruleId', '克隆规则链', 'rulechain', 'POST');
INSERT INTO "public"."sys_apis" VALUES (185, '2023-08-03 14:16:55', '2023-08-03 14:16:55', NULL, '/device/alarm/list', '告警分页列表', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (186, '2023-08-03 14:17:23', '2023-08-03 14:17:23', NULL, '/device/alarm', '修改告警', 'device', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (187, '2023-08-03 14:18:14', '2023-08-03 14:18:14', NULL, '/device/alarm/:id', '删除告警信息', 'device', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (188, '2023-08-04 10:59:57', '2023-08-04 10:59:57', NULL, '/device/cmd/list', '设备命令日志列表', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (189, '2023-08-04 11:00:18', '2023-08-04 11:00:18', NULL, '/device/cmd', '下发命令', 'device', 'POST');
INSERT INTO "public"."sys_apis" VALUES (190, '2023-08-04 11:00:46', '2023-08-04 11:00:46', NULL, '/device/cmd/:id', '删除命令记录', 'device', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (191, '2023-08-04 14:16:06', '2023-08-04 14:16:06', NULL, '/device/template/list/all', '查询所有tsl', 'template', 'GET');
INSERT INTO "public"."sys_apis" VALUES (192, '2023-08-04 16:39:06', '2023-08-04 16:39:06', NULL, '/device/:id/attribute/down', '下发设备属性', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (193, '2023-08-19 09:12:31', '2023-08-19 09:12:31', NULL, '/upload/up', '上传文件到本地', 'upload', 'POST');
INSERT INTO "public"."sys_apis" VALUES (194, '2023-09-05 08:42:13', '2023-09-05 08:42:13', NULL, '/video/ys/device/list', '获取萤石设备列表', 'video', 'GET');
INSERT INTO "public"."sys_apis" VALUES (195, '2023-09-05 08:43:11', '2023-09-05 08:43:11', NULL, '/video/ys/:deviceSerial/channel', '获取指定设备通道', 'video', 'GET');
INSERT INTO "public"."sys_apis" VALUES (196, '2023-09-05 08:45:31', '2023-09-05 08:45:31', NULL, '/video/ys/:deviceSerial/channel/live', '设备通道直播地址', 'video', 'GET');
INSERT INTO "public"."sys_apis" VALUES (197, '2023-09-05 08:46:14', '2023-09-05 08:46:14', NULL, '/video/ys/:deviceSerial/ptz/start', '摄像头操作', 'video', 'GET');
INSERT INTO "public"."sys_apis" VALUES (198, '2023-09-05 08:46:47', '2023-09-05 08:46:47', NULL, '/video/ys/:deviceSerial/ptz/stop', '摄像头操作停止', 'video', 'GET');
INSERT INTO "public"."sys_apis" VALUES (199, '2023-09-06 15:55:44', '2023-09-06 15:55:44', NULL, '/rule/chain/log/list', '规则链审计日志', 'rulechain', 'GET');
INSERT INTO "public"."sys_apis" VALUES (200, '2023-09-06 15:56:35', '2023-09-06 15:56:35', NULL, '/rule/chain/log/delete', '条件删除规则链审计', 'rulechain', 'GET');
INSERT INTO "public"."sys_apis" VALUES (201, '2023-09-08 17:20:35', '2023-09-08 17:20:35', NULL, '/device/:id/property/history', '获取设备属性的遥测历史', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (202, '2023-09-15 17:29:07', '2023-09-15 17:29:07', NULL, '/device/:id/allot/org', '设备分配组织', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (203, '2023-09-22 16:56:56', '2023-09-22 16:56:56', NULL, '/device/product/:id/tsl', '获取产品tsl', 'product', 'GET');
INSERT INTO "public"."sys_apis" VALUES (204, '2023-09-23 14:25:58', '2023-09-23 14:25:58', NULL, '/device/panel', '获取设备统计面板', 'device', 'GET');
INSERT INTO "public"."sys_apis" VALUES (205, '2023-09-25 10:13:59', '2023-09-25 10:13:59', NULL, '/device/alarm/panel', '获取面板告警分组', 'device', 'GET');

-- ----------------------------
-- Table structure for sys_configs
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_configs";
CREATE TABLE "public"."sys_configs" (
  "config_id" int8 NOT NULL,
  "config_name" varchar(128) COLLATE "pg_catalog"."default",
  "config_key" varchar(128) COLLATE "pg_catalog"."default",
  "config_value" varchar(255) COLLATE "pg_catalog"."default",
  "config_type" varchar(64) COLLATE "pg_catalog"."default",
  "is_frontend" varchar(1) COLLATE "pg_catalog"."default",
  "remark" varchar(128) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
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
INSERT INTO "public"."sys_configs" VALUES (1, '账号初始密码', 'sys.user.initPassword', '123456', '0', '0', '初始密码', '2021-12-04 13:50:13', '2021-12-04 13:54:52', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dict_data";
CREATE TABLE "public"."sys_dict_data" (
  "dict_code" int8 NOT NULL,
  "dict_sort" int4,
  "dict_label" varchar(64) COLLATE "pg_catalog"."default",
  "dict_value" varchar(64) COLLATE "pg_catalog"."default",
  "dict_type" varchar(64) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "css_class" varchar(128) COLLATE "pg_catalog"."default",
  "list_class" varchar(128) COLLATE "pg_catalog"."default",
  "is_default" varchar(8) COLLATE "pg_catalog"."default",
  "create_by" varchar(191) COLLATE "pg_catalog"."default",
  "update_by" varchar(191) COLLATE "pg_catalog"."default",
  "remark" varchar(256) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
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
INSERT INTO "public"."sys_dict_data" VALUES (1, 0, '男', '0', 'sys_user_sex', '0', NULL, NULL, NULL, 'admin', NULL, '男', '2021-11-30 14:58:18', '2021-11-30 14:58:18', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (2, 1, '女', '1', 'sys_user_sex', '0', NULL, NULL, NULL, 'admin', NULL, '女生', '2021-11-30 15:09:11', '2021-11-30 15:10:28', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (3, 2, '未知', '2', 'sys_user_sex', '0', NULL, NULL, NULL, 'admin', NULL, '未知', '2021-11-30 15:09:11', '2021-11-30 15:10:28', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (4, 0, '正常', '0', 'sys_normal_disable', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-01 15:58:50', '2021-12-01 15:58:50', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (5, 1, '停用', '1', 'sys_normal_disable', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-01 15:59:08', '2021-12-01 15:59:08', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (6, 0, '目录', 'M', 'sys_menu_type', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-02 09:49:12', '2021-12-02 09:49:12', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (7, 1, '菜单', 'C', 'sys_menu_type', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-02 09:49:35', '2021-12-02 09:49:52', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (8, 2, '按钮', 'F', 'sys_menu_type', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-02 09:49:35', '2021-12-02 09:49:35', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (9, 0, '显示', '0', 'sys_show_hide', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-02 09:56:40', '2021-12-02 09:56:40', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (10, 0, '隐藏', '1', 'sys_show_hide', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-02 09:56:52', '2021-12-02 09:56:52', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (11, 0, '是', '0', 'sys_num_yes_no', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-02 10:16:16', '2021-12-02 10:16:16', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (12, 1, '否', '1', 'sys_num_yes_no', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-02 10:16:26', '2021-12-02 10:16:26', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (13, 0, '是', '0', 'sys_yes_no', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-04 13:48:15', '2021-12-04 13:48:15', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (14, 0, '否', '1', 'sys_yes_no', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-04 13:48:21', '2021-12-04 13:48:21', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (15, 0, '创建(POST)', 'POST', 'sys_method_api', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-08 17:22:05', '2021-12-09 09:29:52', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (16, 1, '查询(GET)', 'GET', 'sys_method_api', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-08 17:22:24', '2021-12-09 09:29:59', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (17, 2, '修改(PUT)', 'PUT', 'sys_method_api', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-08 17:22:40', '2021-12-09 09:30:06', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (18, 3, '删除(DELETE)', 'DELETE', 'sys_method_api', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-08 17:22:54', '2021-12-09 09:30:13', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (19, 0, '成功', '0', 'sys_common_status', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-17 11:01:52', '2021-12-17 11:01:52', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (20, 0, '失败', '1', 'sys_common_status', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-17 11:02:08', '2021-12-17 11:02:08', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (21, 0, '其他', '0', 'sys_oper_type', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-17 11:30:07', '2021-12-17 11:30:07', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (22, 0, '新增', '1', 'sys_oper_type', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-17 11:30:21', '2021-12-17 11:30:21', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (23, 0, '修改', '2', 'sys_oper_type', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-17 11:30:32', '2021-12-17 11:30:32', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (24, 0, '删除', '3', 'sys_oper_type', '0', NULL, NULL, NULL, 'admin', NULL, NULL, '2021-12-17 11:30:40', '2021-12-17 11:30:40', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (25, 0, '默认', 'DEFAULT', 'sys_job_group', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2021-12-24 15:15:31', '2021-12-24 15:15:31', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (26, 1, '系统', 'SYSTEM', 'sys_job_group', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2021-12-24 15:15:50', '2021-12-24 15:15:50', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (27, 0, '发布通知', '1', 'sys_notice_type', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2021-12-26 15:24:07', '2021-12-26 15:24:07', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (28, 0, '任免通知', '2', 'sys_notice_type', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2021-12-26 15:24:18', '2021-12-26 15:24:18', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (29, 0, '事物通知', '3', 'sys_notice_type', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2021-12-26 15:24:46', '2021-12-26 15:24:46', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (30, 0, '审批通知', '4', 'sys_notice_type', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2021-12-26 15:25:08', '2021-12-26 15:25:08', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (31, 0, '阿里云', '0', 'res_oss_category', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2022-01-13 15:44:01', '2022-01-13 15:44:01', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (32, 1, '七牛云', '1', 'res_oss_category', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2022-01-13 15:44:18', '2022-01-13 15:44:18', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (33, 2, '腾讯云', '2', 'res_oss_category', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2022-01-13 15:44:39', '2022-01-13 15:44:39', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (34, 0, '阿里云', '0', 'res_sms_category', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2022-01-13 15:47:30', '2022-01-13 15:47:30', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (35, 1, '腾讯云', '1', 'res_sms_category', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2022-01-13 15:47:39', '2022-01-13 15:47:39', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (36, 0, '163邮箱', '0', 'res_mail_category', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2022-01-14 15:43:42', '2022-01-14 15:43:42', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (37, 0, 'qq邮箱', '1', 'res_mail_category', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2022-01-14 15:44:08', '2022-01-14 15:44:08', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (38, 0, '企业邮箱', '2', 'res_mail_category', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2022-01-14 15:44:20', '2022-01-14 15:44:20', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (39, 1, '未发布', '0', 'sys_release_type', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2023-07-21 16:11:27', '2023-07-21 16:11:27', NULL);
INSERT INTO "public"."sys_dict_data" VALUES (40, 2, '已发布', '1', 'sys_release_type', '0', NULL, NULL, NULL, 'panda', NULL, NULL, '2023-07-21 16:11:44', '2023-07-21 16:11:44', NULL);

-- ----------------------------
-- Table structure for sys_dict_types
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dict_types";
CREATE TABLE "public"."sys_dict_types" (
  "dict_id" int8 NOT NULL,
  "dict_name" varchar(64) COLLATE "pg_catalog"."default",
  "dict_type" varchar(64) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_by" varchar(191) COLLATE "pg_catalog"."default",
  "update_by" varchar(191) COLLATE "pg_catalog"."default",
  "remark" varchar(256) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."sys_dict_types"."dict_name" IS '名称';
COMMENT ON COLUMN "public"."sys_dict_types"."dict_type" IS '类型';
COMMENT ON COLUMN "public"."sys_dict_types"."status" IS '状态';
COMMENT ON COLUMN "public"."sys_dict_types"."remark" IS '备注';

-- ----------------------------
-- Records of sys_dict_types
-- ----------------------------
INSERT INTO "public"."sys_dict_types" VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', NULL, '性别列表', '2021-11-30 14:02:52', '2021-11-30 14:07:55', '2021-11-30 14:11:54');
INSERT INTO "public"."sys_dict_types" VALUES (2, '用户性别', 'sys_user_sex', '0', 'admin', NULL, '用户性别列表', '2021-11-30 14:12:33', '2021-11-30 14:12:33', '2021-11-30 14:14:19');
INSERT INTO "public"."sys_dict_types" VALUES (3, '的心', 'sfd', '0', 'admin', NULL, 'fs', '2021-11-30 14:13:22', '2021-11-30 14:13:22', '2021-11-30 14:14:19');
INSERT INTO "public"."sys_dict_types" VALUES (4, '用户性别', 'sys_user_sex', '0', 'admin', NULL, '性别字典', '2021-11-30 14:15:25', '2021-11-30 14:15:25', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (5, 'df', 'da', '0', 'admin', NULL, 'sd', '2021-11-30 15:54:33', '2021-11-30 15:54:33', '2021-11-30 15:54:40');
INSERT INTO "public"."sys_dict_types" VALUES (6, '系统开关', 'sys_normal_disable', '0', 'admin', NULL, '开关列表', '2021-12-01 15:57:58', '2021-12-01 15:57:58', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (7, '菜单类型', 'sys_menu_type', '0', 'admin', NULL, '菜单类型列表', '2021-12-02 09:48:48', '2021-12-02 09:56:12', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (8, '菜单状态', 'sys_show_hide', '0', 'admin', NULL, '菜单状态列表', '2021-12-02 09:55:59', '2021-12-02 09:55:59', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (9, '数字是否', 'sys_num_yes_no', '0', 'admin', NULL, '数字是否列表', '2021-12-02 10:13:29', '2021-12-02 10:13:40', '2021-12-02 10:15:07');
INSERT INTO "public"."sys_dict_types" VALUES (10, '数字是否', 'sys_num_yes_no', '0', 'admin', NULL, '数字是否', '2021-12-02 10:13:29', '2021-12-02 10:13:29', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (11, '状态是否', 'sys_yes_no', '0', 'admin', NULL, '状态是否', '2021-12-04 13:47:57', '2021-12-04 13:47:57', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (12, '网络请求方法', 'sys_method_api', '0', 'admin', NULL, '网络请求方法列表', '2021-12-08 17:21:27', '2021-12-08 17:21:27', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (13, '成功失败', 'sys_common_status', '0', 'admin', NULL, '是否成功失败', '2021-12-17 10:10:03', '2021-12-17 10:10:03', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (27, '操作分类', 'sys_oper_type', '0', 'admin', NULL, '操作分类列表', '2021-12-17 11:29:50', '2021-12-17 11:29:50', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (28, '任务组', 'sys_job_group', '0', 'panda', NULL, '系统任务，开机自启', '2021-12-24 15:14:56', '2021-12-24 15:14:56', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (29, '通知类型', 'sys_notice_type', '0', 'panda', NULL, '通知类型列表', '2021-12-26 15:23:26', '2021-12-26 15:23:26', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (30, 'oss分类', 'res_oss_category', '0', 'panda', NULL, 'oss分类列表', '2022-01-13 15:43:29', '2022-01-13 15:43:29', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (31, 'sms分类', 'res_sms_category', '0', 'panda', NULL, 'sms分类列表', '2021-12-26 15:23:26', '2022-01-13 15:47:13', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (32, 'mail分类', 'res_mail_category', '0', 'panda', NULL, 'mail分类列表', '2022-01-14 15:43:17', '2022-01-14 15:43:17', NULL);
INSERT INTO "public"."sys_dict_types" VALUES (33, '发布状态', 'sys_release_type', '0', 'panda', NULL, '发布状态', '2023-07-21 16:10:38', '2023-07-21 16:10:38', NULL);

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_menus";
CREATE TABLE "public"."sys_menus" (
  "menu_id" int8 NOT NULL,
  "menu_name" varchar(128) COLLATE "pg_catalog"."default",
  "title" varchar(64) COLLATE "pg_catalog"."default",
  "parent_id" int4,
  "sort" int4,
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
  "status" varchar(191) COLLATE "pg_catalog"."default",
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "remark" varchar(191) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
)
;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO "public"."sys_menus" VALUES (1, '系统设置', NULL, 0, 0, 'elementSetting', '/system', 'Layout', '1', NULL, 'M', '0', '0', '1', NULL, '0', 'admin', 'panda', NULL, '2021-12-02 11:04:08', '2021-12-28 13:32:21', NULL);
INSERT INTO "public"."sys_menus" VALUES (3, '用户管理', NULL, 1, 1, 'elementUser', '/system/user', '/system/user/index', '1', NULL, 'C', '0', '1', '1', 'system:user:list', '0', 'admin', 'panda', NULL, '2021-12-02 14:07:56', '2021-12-28 13:32:44', NULL);
INSERT INTO "public"."sys_menus" VALUES (4, '添加用户', NULL, 3, 1, NULL, NULL, NULL, NULL, NULL, 'F', '0', NULL, NULL, 'system:user:add', '0', 'admin', NULL, NULL, '2021-12-03 13:36:33', '2021-12-03 13:36:33', NULL);
INSERT INTO "public"."sys_menus" VALUES (5, '编辑用户', NULL, 3, 1, NULL, NULL, NULL, NULL, NULL, 'F', '0', NULL, NULL, 'system:user:edit', '0', 'admin', NULL, NULL, '2021-12-03 13:48:13', '2021-12-03 13:48:13', NULL);
INSERT INTO "public"."sys_menus" VALUES (6, '角色管理', NULL, 1, 2, 'elementUserFilled', '/system/role', '/system/role/index', '1', NULL, 'C', '0', '1', '1', 'system:role:list', '0', NULL, 'panda', NULL, '2021-12-03 13:51:55', '2022-07-16 10:23:21', NULL);
INSERT INTO "public"."sys_menus" VALUES (7, '菜单管理', NULL, 1, 2, 'iconfont icon-juxingkaobei', '/system/menu', '/system/menu/index', '1', NULL, 'C', '0', '1', '1', 'system:menu:list', '0', 'admin', 'panda', NULL, '2021-12-03 13:54:44', '2021-12-28 13:33:19', NULL);
INSERT INTO "public"."sys_menus" VALUES (8, '组织管理', NULL, 1, 3, 'iconfont icon-jiliandongxuanzeqi', '/system/organization', '/system/organization/index', '1', NULL, 'C', '0', '1', '1', 'system:organization:list', '0', NULL, 'panda', NULL, '2021-12-03 13:58:36', '2023-09-14 14:05:07', NULL);
INSERT INTO "public"."sys_menus" VALUES (9, '岗位管理', NULL, 1, 4, 'iconfont icon-neiqianshujuchucun', '/system/post', '/system/post/index', '1', NULL, 'C', '0', '1', '1', 'system:post:list', '0', 'admin', 'panda', NULL, '2021-12-03 13:54:44', '2021-12-28 13:40:31', NULL);
INSERT INTO "public"."sys_menus" VALUES (10, '字典管理', NULL, 1, 5, 'elementCellphone', '/system/dict', '/system/dict/index', '1', NULL, 'C', '0', '1', '1', 'system:dict:list', '0', 'admin', 'panda', NULL, '2021-12-03 13:54:44', '2021-12-28 13:40:50', NULL);
INSERT INTO "public"."sys_menus" VALUES (11, '参数管理', NULL, 1, 6, 'elementDocumentCopy', '/system/config', '/system/config/index', '1', NULL, 'C', '0', '1', '1', 'system:config:list', '0', 'admin', 'panda', NULL, '2021-12-03 13:54:44', '2021-12-28 13:41:05', NULL);
INSERT INTO "public"."sys_menus" VALUES (12, '个人中心', NULL, 0, 10, 'elementAvatar', '/personal', '/personal/index', '1', NULL, 'M', '1', '0', '0', NULL, '0', 'admin', 'panda', NULL, '2021-12-03 14:12:43', '2023-06-27 10:09:26', NULL);
INSERT INTO "public"."sys_menus" VALUES (13, '添加配置', NULL, 11, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:config:add', '0', 'admin', NULL, NULL, '2021-12-06 17:19:19', '2021-12-06 17:19:19', NULL);
INSERT INTO "public"."sys_menus" VALUES (14, '修改配置', NULL, 11, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:config:edit', '0', 'admin', NULL, NULL, '2021-12-06 17:20:30', '2021-12-06 17:20:30', NULL);
INSERT INTO "public"."sys_menus" VALUES (15, '删除配置', NULL, 11, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:config:delete', '0', 'admin', NULL, NULL, '2021-12-06 17:23:52', '2021-12-06 17:23:52', NULL);
INSERT INTO "public"."sys_menus" VALUES (16, '导出配置', NULL, 11, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:config:export', '0', 'admin', NULL, NULL, '2021-12-06 17:24:41', '2021-12-06 17:24:41', NULL);
INSERT INTO "public"."sys_menus" VALUES (17, '新增角色', NULL, 6, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:role:add', '0', 'admin', NULL, NULL, '2021-12-06 17:43:35', '2021-12-06 17:43:35', NULL);
INSERT INTO "public"."sys_menus" VALUES (18, '删除角色', NULL, 6, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:role:delete', '0', 'admin', NULL, NULL, '2021-12-06 17:44:10', '2021-12-06 17:44:10', NULL);
INSERT INTO "public"."sys_menus" VALUES (19, '修改角色', NULL, 6, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:role:edit', '0', 'admin', NULL, NULL, '2021-12-06 17:44:48', '2021-12-06 17:44:48', NULL);
INSERT INTO "public"."sys_menus" VALUES (20, '导出角色', NULL, 6, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:role:export', '0', 'admin', NULL, NULL, '2021-12-06 17:45:25', '2021-12-06 17:45:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (21, '添加菜单', NULL, 7, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:menu:add', '0', 'admin', NULL, NULL, '2021-12-06 17:46:01', '2021-12-06 17:46:01', NULL);
INSERT INTO "public"."sys_menus" VALUES (22, '修改菜单', NULL, 7, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:menu:edit', '0', 'admin', NULL, NULL, '2021-12-06 17:46:24', '2021-12-06 17:46:24', NULL);
INSERT INTO "public"."sys_menus" VALUES (23, '删除菜单', NULL, 7, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:menu:delete', '0', 'admin', NULL, NULL, '2021-12-06 17:46:47', '2021-12-06 17:46:47', NULL);
INSERT INTO "public"."sys_menus" VALUES (24, '添加部门', NULL, 8, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:organization:add', '0', NULL, 'panda', NULL, '2021-12-07 09:33:58', '2023-09-14 14:05:20', NULL);
INSERT INTO "public"."sys_menus" VALUES (25, '编辑部门', NULL, 8, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:organization:edit', '0', NULL, 'panda', NULL, '2021-12-07 09:34:39', '2023-09-14 14:05:26', NULL);
INSERT INTO "public"."sys_menus" VALUES (26, '删除部门', NULL, 8, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:organization:delete', '0', NULL, 'panda', NULL, '2021-12-07 09:35:09', '2023-09-14 14:05:32', NULL);
INSERT INTO "public"."sys_menus" VALUES (27, '导出部门', NULL, 8, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:dept:export', '0', 'admin', NULL, NULL, '2021-12-07 09:35:51', '2021-12-07 09:35:51', '2021-12-07 09:36:37');
INSERT INTO "public"."sys_menus" VALUES (28, '添加岗位', NULL, 9, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:post:add', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (29, '编辑岗位', NULL, 9, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:post:edit', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (30, '删除岗位', NULL, 9, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:post:delete', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (31, '导出岗位', NULL, 9, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:post:export', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (32, '添加字典类型', NULL, 10, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:dictT:add', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (33, '编辑字典类型', NULL, 10, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:dictT:edit', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (34, '删除字典类型', NULL, 10, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:dictT:delete', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (35, '导出字典类型', NULL, 10, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:dictT:export', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (36, '新增字典数据', NULL, 10, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:dictD:add', '0', 'admin', NULL, NULL, '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (37, '修改字典数据', NULL, 10, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:dictD:edit', '0', 'admin', NULL, NULL, '2021-12-07 09:48:04', '2021-12-07 09:48:04', NULL);
INSERT INTO "public"."sys_menus" VALUES (38, '删除字典数据', NULL, 10, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:dictD:delete', '0', 'admin', NULL, NULL, '2021-12-07 09:48:42', '2021-12-07 09:48:42', NULL);
INSERT INTO "public"."sys_menus" VALUES (39, 'API管理', NULL, 1, 2, 'iconfont icon-siweidaotu', '/system/api', '/system/api/index', '1', NULL, 'C', '0', '1', '1', 'system:api:list', '0', NULL, 'panda', NULL, '2021-12-09 09:09:13', '2022-07-16 10:23:42', NULL);
INSERT INTO "public"."sys_menus" VALUES (40, '添加api', NULL, 39, 1, NULL, '/system/api', NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:api:add', '0', 'admin', NULL, NULL, '2021-12-09 09:09:54', '2021-12-09 09:09:54', NULL);
INSERT INTO "public"."sys_menus" VALUES (41, '编辑api', NULL, 39, 1, NULL, '/system/api', NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:api:edit', '0', 'admin', NULL, NULL, '2021-12-09 09:10:38', '2021-12-09 09:10:38', NULL);
INSERT INTO "public"."sys_menus" VALUES (42, '删除api', NULL, 39, 1, NULL, '/system/api', NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'system:api:delete', '0', 'admin', NULL, NULL, '2021-12-09 09:11:11', '2021-12-09 09:11:11', NULL);
INSERT INTO "public"."sys_menus" VALUES (43, '日志系统', NULL, 0, 11, 'iconfont icon-biaodan', '/log', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'admin', 'panda', NULL, '2021-12-02 11:04:08', '2023-06-30 08:57:08', NULL);
INSERT INTO "public"."sys_menus" VALUES (44, '运维监控', NULL, 0, 9, 'iconfont icon-gongju', '/tool', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'admin', 'panda', NULL, '2021-12-16 16:35:15', '2023-06-27 10:09:16', NULL);
INSERT INTO "public"."sys_menus" VALUES (45, '操作日志', NULL, 43, 1, 'iconfont icon-bolangnengshiyanchang', '/log/operation', '/log/operation/index', '1', NULL, 'C', '0', '1', '1', 'log:operation:list', '0', 'admin', 'panda', NULL, '2021-12-16 16:42:03', '2021-12-28 13:39:44', NULL);
INSERT INTO "public"."sys_menus" VALUES (46, '登录日志', NULL, 43, 2, 'iconfont icon--chaifenlie', '/log/login', '/log/login/index', '1', NULL, 'C', '0', '1', '1', 'log:login:list', '0', 'admin', 'panda', NULL, '2021-12-16 16:43:28', '2021-12-28 13:39:58', NULL);
INSERT INTO "public"."sys_menus" VALUES (47, '服务监控', NULL, 44, 1, 'elementCpu', '/tool/monitor/', '/tool/monitor/index', '1', NULL, 'C', '0', '1', '1', 'tool:monitor:list', '0', 'admin', 'panda', NULL, '2021-12-03 14:12:43', '2021-12-28 13:41:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (49, '开发工具', NULL, 0, 10, 'iconfont icon-zhongduancanshu', '/develop', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'admin', 'panda', NULL, '2021-12-16 16:53:11', '2023-06-29 16:29:23', NULL);
INSERT INTO "public"."sys_menus" VALUES (50, '表单构建', NULL, 49, 1, 'iconfont icon-zidingyibuju', '/develop/form', '/develop/form/index', '1', NULL, 'C', '0', '1', '1', 'develop:form:list', '0', 'admin', 'panda', NULL, '2021-12-16 16:55:01', '2022-07-12 18:56:18', NULL);
INSERT INTO "public"."sys_menus" VALUES (51, '代码生成', NULL, 49, 2, 'iconfont icon-zhongduancanshu', '/develop/code', '/develop/code/index', '1', NULL, 'C', '0', '1', '1', 'develop:code:list', '0', 'admin', NULL, NULL, '2021-12-16 16:56:48', '2021-12-16 16:56:48', NULL);
INSERT INTO "public"."sys_menus" VALUES (52, '系统接口', NULL, 49, 3, 'iconfont icon-wenducanshu-05', '/develop/apis', '/layout/routerView/iframes', '0', 'https://82200r6gti.apifox.cn', 'C', '0', '1', '1', 'develop:apis:list', '0', NULL, 'panda', NULL, '2021-12-16 16:58:07', '2023-09-04 11:02:29', NULL);
INSERT INTO "public"."sys_menus" VALUES (54, '对象存储', NULL, 53, 1, 'iconfont icon-chazhaobiaodanliebiao', '/resource/file', '/resource/file/index', '1', NULL, 'C', '0', '1', '1', 'resource:file:list', '0', 'admin', 'panda', NULL, '2021-12-16 17:06:04', '2022-01-13 17:30:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (55, '公告通知', NULL, 44, 3, 'elementTicket', '/tool/notice', '/tool/notice/index', '1', NULL, 'C', '0', '1', '1', 'tool:notice:list', '0', 'admin', 'panda', NULL, '2021-12-16 22:09:11', '2021-12-28 13:42:39', NULL);
INSERT INTO "public"."sys_menus" VALUES (59, '删除', NULL, 45, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'log:operation:delete', '0', 'panda', NULL, NULL, '2022-01-14 13:28:25', '2022-01-14 13:28:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (60, '清空', NULL, 45, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'log:operation:clean', '0', 'panda', NULL, NULL, '2022-01-14 13:29:24', '2022-01-14 13:29:24', NULL);
INSERT INTO "public"."sys_menus" VALUES (63, '删除', NULL, 46, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'log:login:delete', '0', 'panda', NULL, NULL, '2022-01-14 13:30:46', '2022-01-14 13:30:46', NULL);
INSERT INTO "public"."sys_menus" VALUES (64, '清空', NULL, 46, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'log:login:clean', '0', 'panda', NULL, NULL, '2022-01-14 13:31:06', '2022-01-14 13:31:06', NULL);
INSERT INTO "public"."sys_menus" VALUES (69, '添加', NULL, 55, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'tool:notice:add', '0', 'panda', NULL, NULL, '2022-01-14 13:35:23', '2022-01-14 13:35:23', NULL);
INSERT INTO "public"."sys_menus" VALUES (70, '编辑', NULL, 55, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'tool:notice:edit', '0', 'panda', NULL, NULL, '2022-01-14 13:36:04', '2022-01-14 13:36:04', NULL);
INSERT INTO "public"."sys_menus" VALUES (71, '删除', NULL, 55, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'tool:notice:delete', '0', 'panda', NULL, NULL, '2022-01-14 13:36:26', '2022-01-14 13:36:26', NULL);
INSERT INTO "public"."sys_menus" VALUES (72, '查看', NULL, 55, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'tool:notice:view', '0', 'panda', NULL, NULL, '2022-01-14 13:36:51', '2022-01-14 13:36:51', NULL);
INSERT INTO "public"."sys_menus" VALUES (73, '导入', NULL, 51, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'develop:code:add', '0', 'panda', NULL, NULL, '2022-01-14 13:38:35', '2022-01-14 13:38:35', NULL);
INSERT INTO "public"."sys_menus" VALUES (74, '编辑', NULL, 51, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'develop:code:edit', '0', 'panda', NULL, NULL, '2022-01-14 13:41:25', '2022-01-14 13:41:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (75, '删除', NULL, 51, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'develop:code:delete', '0', 'panda', NULL, NULL, '2022-01-14 13:41:42', '2022-01-14 13:41:42', NULL);
INSERT INTO "public"."sys_menus" VALUES (76, '预览', NULL, 51, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'develop:code:view', '0', 'panda', NULL, NULL, '2022-01-14 13:42:01', '2022-01-14 13:42:01', NULL);
INSERT INTO "public"."sys_menus" VALUES (77, '生成代码', NULL, 51, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'develop:code:code', '0', 'panda', NULL, NULL, '2022-01-14 13:42:48', '2022-01-14 13:42:48', NULL);
INSERT INTO "public"."sys_menus" VALUES (95, '设备管理', NULL, 0, 1, 'iconfont icon-dongtai', '/device', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'panda', 'panda', NULL, '2023-06-29 16:21:31', '2023-09-02 15:03:55', NULL);
INSERT INTO "public"."sys_menus" VALUES (96, '规则链库', NULL, 0, 2, 'iconfont icon-shuxingtu', '/rule', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'panda', 'panda', NULL, '2023-06-29 16:33:23', '2023-09-02 15:04:06', NULL);
INSERT INTO "public"."sys_menus" VALUES (97, '组态大屏', NULL, 0, 3, 'iconfont icon-diannaobangong', '/visual', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'panda', 'panda', NULL, '2023-06-29 16:34:50', '2023-07-21 14:41:55', NULL);
INSERT INTO "public"."sys_menus" VALUES (98, '产品分类', NULL, 95, 1, 'iconfont icon-jiliandongxuanzeqi', '/device/product_category', '/device/product_category/index', '1', NULL, 'C', '0', '1', '1', 'product:category:list', '0', NULL, 'panda', NULL, '2023-06-29 16:44:56', '2023-06-29 16:49:55', NULL);
INSERT INTO "public"."sys_menus" VALUES (99, '产品', NULL, 95, 2, 'iconfont icon-AIshiyanshi', '/device/product', '/device/product/index', '1', NULL, 'C', '0', '1', '1', 'product:list', '0', NULL, 'panda', NULL, '2023-06-29 16:46:16', '2023-06-29 16:50:34', '2023-06-30 14:16:11');
INSERT INTO "public"."sys_menus" VALUES (100, '设备分组', NULL, 95, 3, 'iconfont icon-zidingyibuju', '/device/device_group', '/device/device_group/index', '1', NULL, 'C', '0', '1', '1', 'device:group:list', '0', NULL, 'panda', NULL, '2023-06-29 16:48:05', '2023-06-29 16:50:49', NULL);
INSERT INTO "public"."sys_menus" VALUES (101, '设备', NULL, 95, 4, 'iconfont icon-dongtai', '/device/device', '/device/device/index', '1', NULL, 'C', '0', '1', '1', 'device:list', '0', NULL, 'panda', NULL, '2023-06-29 16:49:03', '2023-06-29 16:50:59', '2023-06-30 14:16:48');
INSERT INTO "public"."sys_menus" VALUES (102, '添加', NULL, 98, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'product:category:add', '0', 'panda', NULL, NULL, '2023-06-29 16:51:38', '2023-06-29 16:51:38', NULL);
INSERT INTO "public"."sys_menus" VALUES (103, '修改', NULL, 98, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'product:category:edit', '0', 'panda', NULL, NULL, '2023-06-29 16:52:00', '2023-06-29 16:52:00', NULL);
INSERT INTO "public"."sys_menus" VALUES (104, '删除', NULL, 98, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'product:category:delete', '0', 'panda', NULL, NULL, '2023-06-29 16:52:36', '2023-06-29 16:52:36', NULL);
INSERT INTO "public"."sys_menus" VALUES (105, '新增', NULL, 100, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:group:add', '0', 'panda', NULL, NULL, '2023-06-29 16:53:16', '2023-06-29 16:53:16', NULL);
INSERT INTO "public"."sys_menus" VALUES (106, '修改', NULL, 100, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:group:edit', '0', 'panda', NULL, NULL, '2023-06-29 16:53:37', '2023-06-29 16:53:37', NULL);
INSERT INTO "public"."sys_menus" VALUES (107, '删除', NULL, 100, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:group:delete', '0', 'panda', NULL, NULL, '2023-06-29 16:53:56', '2023-06-29 16:53:56', NULL);
INSERT INTO "public"."sys_menus" VALUES (108, '克隆', NULL, 96, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rulechain:clone', '0', 'panda', NULL, NULL, '2023-06-29 16:54:33', '2023-06-29 16:54:33', '2023-07-21 14:41:09');
INSERT INTO "public"."sys_menus" VALUES (109, '设计', NULL, 96, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rulechain:design', '0', 'panda', NULL, NULL, '2023-06-29 16:55:18', '2023-06-29 16:55:18', '2023-07-21 14:41:14');
INSERT INTO "public"."sys_menus" VALUES (110, '预览', NULL, 96, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rulechain:view', '0', 'panda', NULL, NULL, '2023-06-29 16:55:39', '2023-06-29 16:55:39', '2023-07-21 14:41:18');
INSERT INTO "public"."sys_menus" VALUES (111, '修改', NULL, 96, 4, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rulechain:edit', '0', 'panda', NULL, NULL, '2023-06-29 16:55:59', '2023-06-29 16:55:59', '2023-07-21 14:41:30');
INSERT INTO "public"."sys_menus" VALUES (112, '删除', NULL, 96, 5, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rulechain:delete', '0', 'panda', NULL, NULL, '2023-06-29 16:56:15', '2023-06-29 16:56:15', '2023-07-21 14:41:26');
INSERT INTO "public"."sys_menus" VALUES (113, '添加', NULL, 96, 6, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rulechain:add', '0', 'panda', NULL, NULL, '2023-06-29 16:56:37', '2023-06-29 16:56:37', '2023-07-21 14:41:21');
INSERT INTO "public"."sys_menus" VALUES (114, '产品管理', NULL, 95, 2, 'elementCpu', '/device/product', '/device/product/index', '1', NULL, 'C', '0', '1', '1', 'device:product:list', '0', NULL, 'panda', NULL, '2023-06-30 14:13:39', '2023-07-21 16:03:31', NULL);
INSERT INTO "public"."sys_menus" VALUES (115, '新增产品', NULL, 114, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:product:add', '0', 'admin', NULL, NULL, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL);
INSERT INTO "public"."sys_menus" VALUES (116, '修改产品', NULL, 114, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:product:edit', '0', 'admin', NULL, NULL, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL);
INSERT INTO "public"."sys_menus" VALUES (117, '删除产品', NULL, 114, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:product:delete', '0', 'admin', NULL, NULL, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL);
INSERT INTO "public"."sys_menus" VALUES (118, '设备管理', NULL, 95, 4, 'elementSetting', '/device/device', '/device/device/index', '1', NULL, 'C', '0', '1', '1', 'device:device:list', '0', NULL, 'panda', NULL, '2023-06-30 14:20:03', '2023-07-21 16:03:41', NULL);
INSERT INTO "public"."sys_menus" VALUES (119, '修改设备', NULL, 118, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:device:edit', '0', 'admin', NULL, NULL, '2023-06-30 14:20:03', '2023-06-30 14:20:03', NULL);
INSERT INTO "public"."sys_menus" VALUES (120, '新增设备', NULL, 118, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:device:add', '0', 'admin', NULL, NULL, '2023-06-30 14:20:03', '2023-06-30 14:20:03', NULL);
INSERT INTO "public"."sys_menus" VALUES (121, '删除设备', NULL, 118, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:device:delete', '0', 'admin', NULL, NULL, '2023-06-30 14:20:03', '2023-06-30 14:20:03', NULL);
INSERT INTO "public"."sys_menus" VALUES (122, '查看', NULL, 114, 4, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:product:view', '0', 'panda', NULL, NULL, '2023-07-05 17:14:20', '2023-07-05 17:14:20', NULL);
INSERT INTO "public"."sys_menus" VALUES (131, '查看设备', NULL, 118, 4, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:device:view', '0', 'panda', NULL, NULL, '2023-07-10 08:50:48', '2023-07-10 08:50:48', NULL);
INSERT INTO "public"."sys_menus" VALUES (132, '规则设计', NULL, 96, 1, 'iconfont icon-shuxingtu', '/rule/chain', '/rule/chain/index', '1', NULL, 'C', '0', '1', '1', 'rule:chain:list', '0', NULL, 'panda', NULL, '2023-07-21 14:38:54', '2023-09-02 14:33:03', NULL);
INSERT INTO "public"."sys_menus" VALUES (133, '克隆', NULL, 132, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rule:chain:clone', '0', NULL, 'panda', NULL, '2023-07-21 14:39:27', '2023-07-21 14:57:05', NULL);
INSERT INTO "public"."sys_menus" VALUES (134, '设计', NULL, 132, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rule:chain:design', '0', NULL, 'panda', NULL, '2023-07-21 14:39:53', '2023-07-21 14:57:13', NULL);
INSERT INTO "public"."sys_menus" VALUES (135, '预览', NULL, 132, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rule:chain:view', '0', NULL, 'panda', NULL, '2023-07-21 14:40:08', '2023-07-21 14:57:20', NULL);
INSERT INTO "public"."sys_menus" VALUES (136, '修改', NULL, 132, 4, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rule:chain:edit', '0', NULL, 'panda', NULL, '2023-07-21 14:40:31', '2023-07-21 14:57:26', NULL);
INSERT INTO "public"."sys_menus" VALUES (137, '删除', NULL, 132, 5, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rule:chain:delete', '0', NULL, 'panda', NULL, '2023-07-21 14:40:47', '2023-07-21 14:57:33', NULL);
INSERT INTO "public"."sys_menus" VALUES (138, '添加', NULL, 132, 6, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'rule:chain:add', '0', NULL, 'panda', NULL, '2023-07-21 14:41:04', '2023-07-21 14:57:39', NULL);
INSERT INTO "public"."sys_menus" VALUES (139, '大屏分组', NULL, 97, 1, 'iconfont icon-wenducanshu-05', '/visual/screen_group', '/visual/screen_group/index', '1', NULL, 'C', '0', '1', '1', 'screen:group:list', '0', 'panda', NULL, NULL, '2023-07-21 14:46:41', '2023-07-21 14:46:41', NULL);
INSERT INTO "public"."sys_menus" VALUES (140, '组态大屏', NULL, 97, 2, 'iconfont icon-diannaobangong', '/visual/screen', '/visual/screen/index', '1', NULL, 'C', '0', '1', '1', 'visual:screen:list', '0', 'panda', NULL, NULL, '2023-07-21 14:47:46', '2023-07-21 14:47:46', NULL);
INSERT INTO "public"."sys_menus" VALUES (141, '添加', NULL, 139, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'screen:group:add', '0', 'panda', NULL, NULL, '2023-07-21 14:50:40', '2023-07-21 14:50:40', NULL);
INSERT INTO "public"."sys_menus" VALUES (142, '编辑', NULL, 139, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'screen:group:edit', '0', 'panda', NULL, NULL, '2023-07-21 14:50:56', '2023-07-21 14:50:56', NULL);
INSERT INTO "public"."sys_menus" VALUES (143, '删除', NULL, 139, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, '	 screen:group:delete', '0', 'panda', NULL, NULL, '2023-07-21 14:51:22', '2023-07-21 14:51:22', NULL);
INSERT INTO "public"."sys_menus" VALUES (144, '新增组态', NULL, 140, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'visual:screen:add', '0', 'panda', NULL, NULL, '2023-07-21 14:53:26', '2023-07-21 14:53:26', NULL);
INSERT INTO "public"."sys_menus" VALUES (145, '修改大屏', NULL, 140, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'visual:screen:edit', '0', 'panda', NULL, NULL, '2023-07-21 14:53:50', '2023-07-21 14:53:50', NULL);
INSERT INTO "public"."sys_menus" VALUES (146, '删除大屏', NULL, 140, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'visual:screen:delete', '0', 'panda', NULL, NULL, '2023-07-21 14:54:14', '2023-07-21 14:54:14', NULL);
INSERT INTO "public"."sys_menus" VALUES (147, '克隆', NULL, 140, 4, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'visual:screen:clone', '0', 'panda', NULL, NULL, '2023-07-21 14:54:30', '2023-07-21 14:54:30', NULL);
INSERT INTO "public"."sys_menus" VALUES (148, '设计', NULL, 140, 5, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'visual:screen:design', '0', 'panda', NULL, NULL, '2023-07-21 14:54:57', '2023-07-21 14:54:57', NULL);
INSERT INTO "public"."sys_menus" VALUES (149, '预览', NULL, 140, 6, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'visual:screen:view', '0', 'panda', NULL, NULL, '2023-07-21 14:55:27', '2023-07-21 14:55:27', NULL);
INSERT INTO "public"."sys_menus" VALUES (150, '报表管理', NULL, 0, 4, 'iconfont icon-putong', '/report', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'panda', 'panda', NULL, '2023-07-24 10:12:26', '2023-07-24 10:13:54', '2023-08-18 14:01:16');
INSERT INTO "public"."sys_menus" VALUES (151, '报表设计', NULL, 150, 1, 'iconfont icon-dayin', '/report', '/layout/routerView/iframes', '0', 'http://101.35.247.125:9001/edit', 'C', '0', '1', '1', 'report:list', '0', 'panda', NULL, NULL, '2023-07-24 10:13:47', '2023-07-24 10:13:47', '2023-08-18 14:01:12');
INSERT INTO "public"."sys_menus" VALUES (152, '任务中心', NULL, 0, 5, 'iconfont icon-dayin', '/job', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'panda', 'panda', NULL, '2023-08-08 14:08:11', '2023-10-05 14:03:38', NULL);
INSERT INTO "public"."sys_menus" VALUES (153, '任务中心', NULL, 152, 1, 'elementAlarmClock', '/job/job', '/job/job/index', '1', NULL, 'C', '0', '1', '1', 'job:list', '0', NULL, 'panda', NULL, '2023-08-08 14:10:37', '2023-08-08 14:12:49', NULL);
INSERT INTO "public"."sys_menus" VALUES (154, '任务日志', NULL, 152, 2, 'elementDocument', '/job/log', '/job/log/index', '1', NULL, 'C', '0', '1', '1', 'job:log:list', '0', 'panda', NULL, NULL, '2023-08-08 14:12:37', '2023-08-08 14:12:37', NULL);
INSERT INTO "public"."sys_menus" VALUES (155, '新增', NULL, 153, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'job:add', '0', 'panda', NULL, NULL, '2023-08-08 14:20:17', '2023-08-08 14:20:17', NULL);
INSERT INTO "public"."sys_menus" VALUES (156, '编辑', NULL, 153, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'job:edit', '0', 'panda', NULL, NULL, '2023-08-08 14:20:44', '2023-08-08 14:20:44', NULL);
INSERT INTO "public"."sys_menus" VALUES (157, '删除', NULL, 153, 3, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'job:delete', '0', 'panda', NULL, NULL, '2023-08-08 14:21:03', '2023-08-08 14:21:03', NULL);
INSERT INTO "public"."sys_menus" VALUES (158, '运行启动', NULL, 153, 4, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'job:run', '0', 'panda', NULL, NULL, '2023-08-08 14:21:25', '2023-08-08 14:21:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (159, '删除', NULL, 154, 1, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'job:log:delete', '0', 'panda', NULL, NULL, '2023-08-08 14:22:05', '2023-08-08 14:22:05', NULL);
INSERT INTO "public"."sys_menus" VALUES (160, '清空', NULL, 154, 2, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'job:log:clean', '0', 'panda', NULL, NULL, '2023-08-08 14:22:33', '2023-08-08 14:22:33', NULL);
INSERT INTO "public"."sys_menus" VALUES (161, '视频监控', NULL, 0, 4, 'iconfont icon-step', '/video', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'panda', 'panda', NULL, '2023-09-02 13:52:17', '2023-10-05 14:04:19', NULL);
INSERT INTO "public"."sys_menus" VALUES (164, '视频广场', NULL, 161, 3, 'elementGrid', '/video/splitview', '/video/splitview/index', '1', NULL, 'C', '0', '1', '1', 'video:splitview:list', '0', 'panda', NULL, NULL, '2023-09-02 14:00:41', '2023-09-02 14:00:41', NULL);
INSERT INTO "public"."sys_menus" VALUES (165, '规则审计', NULL, 96, 2, 'iconfont icon--chaifenhang', '/rule/log', '/rule/log/index', '1', NULL, 'C', '0', '1', '1', 'rule:log:list', '0', 'panda', NULL, NULL, '2023-09-02 14:05:46', '2023-09-02 14:05:46', NULL);
INSERT INTO "public"."sys_menus" VALUES (166, '告警中心', NULL, 0, 6, 'iconfont icon-radio-off-full', '/alarm', 'Layout', '1', NULL, 'M', '0', '1', '1', NULL, '0', 'panda', 'panda', NULL, '2023-09-02 14:10:27', '2023-09-02 14:11:21', NULL);
INSERT INTO "public"."sys_menus" VALUES (167, '告警日志', NULL, 166, 1, 'iconfont icon-jiliandongxuanzeqi', '/alarm/log', '/alarm/log/index', '1', NULL, 'C', '0', '1', '1', 'alarm:log:list', '0', 'panda', NULL, NULL, '2023-09-02 14:12:13', '2023-09-02 14:12:13', NULL);
INSERT INTO "public"."sys_menus" VALUES (168, '设备地图', NULL, 95, 5, 'iconfont icon-ditu', 'device:map', '/device/map/index', '1', NULL, 'C', '0', '1', '1', 'device:map:list', '0', 'panda', NULL, NULL, '2023-09-02 14:14:00', '2023-09-02 14:14:00', NULL);
INSERT INTO "public"."sys_menus" VALUES (169, '边缘管理', NULL, 0, 7, 'iconfont icon-wendu', '/edge', 'Layout', '1', NULL, 'M', '1', '1', '1', NULL, '0', 'panda', 'panda', NULL, '2023-09-02 14:27:39', '2023-10-05 10:15:19', NULL);
INSERT INTO "public"."sys_menus" VALUES (170, '网关管理', NULL, 169, 1, 'iconfont icon-gouxuan-weixuanzhong-xianxingfangkuang', '/edge/gateway', '/edge/gateway/index', '1', NULL, 'C', '0', '1', '1', 'edge:gateway:list', '0', NULL, 'panda', NULL, '2023-09-02 14:44:13', '2023-09-19 10:20:34', NULL);
INSERT INTO "public"."sys_menus" VALUES (171, '采集器', NULL, 169, 2, 'iconfont icon-wendu', '/edge/collector', '/edge/collector/index', '1', NULL, 'C', '0', '1', '1', 'edge:collector:list', '0', NULL, 'panda', NULL, '2023-09-02 14:45:31', '2023-09-19 10:20:57', NULL);
INSERT INTO "public"."sys_menus" VALUES (172, '应用管理', NULL, 0, 8, 'iconfont icon-shoujidiannao', '/apply', 'Layout', '1', NULL, 'M', '1', '1', '1', NULL, '1', 'panda', 'panda', NULL, '2023-09-02 14:50:48', '2023-09-04 10:55:19', NULL);
INSERT INTO "public"."sys_menus" VALUES (173, '应用商店', NULL, 172, 1, 'iconfont icon-shoujidiannao', '/apply/common', '/apply/common/index', '1', NULL, 'C', '0', '1', '1', 'apply:common:list', '0', 'panda', NULL, NULL, '2023-09-02 14:51:56', '2023-09-02 14:51:56', NULL);
INSERT INTO "public"."sys_menus" VALUES (174, '我的应用', NULL, 172, 2, 'iconfont icon-LoggedinPC', '/apply/meapp', '/apply/meapp/index', '1', NULL, 'C', '0', '1', '1', 'apply:meapp:list', '0', 'panda', NULL, NULL, '2023-09-02 14:52:45', '2023-09-02 14:52:45', NULL);
INSERT INTO "public"."sys_menus" VALUES (175, '萤石设备', NULL, 161, 2, 'iconfont icon-gerenzhongxin', '/video/ezviz', '/video/ezviz/index', '1', NULL, 'C', '0', '1', '1', 'video:ezviz:list', '0', 'panda', NULL, NULL, '2023-09-05 10:05:27', '2023-09-05 10:05:27', NULL);
INSERT INTO "public"."sys_menus" VALUES (176, '国标设备', NULL, 161, 1, 'iconfont icon-wendu', '/video/gb28181', '/video/gb28181/index', '1', NULL, 'C', '0', '1', '1', 'video:gb28181:list', '0', 'panda', NULL, NULL, '2023-09-05 10:07:07', '2023-09-05 10:07:07', NULL);
INSERT INTO "public"."sys_menus" VALUES (177, '分配设备', NULL, 118, 5, NULL, NULL, NULL, NULL, NULL, 'F', NULL, NULL, NULL, 'device:device:allot', '0', 'panda', NULL, NULL, '2023-09-15 17:32:08', '2023-09-15 17:32:08', NULL);

-- ----------------------------
-- Table structure for sys_notices
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_notices";
CREATE TABLE "public"."sys_notices" (
  "notice_id" int8 NOT NULL,
  "title" varchar(128) COLLATE "pg_catalog"."default",
  "content" text COLLATE "pg_catalog"."default",
  "notice_type" varchar(1) COLLATE "pg_catalog"."default",
  "organization_id" int4,
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6),
  "user_name" varchar(64) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_notices"."title" IS '标题';
COMMENT ON COLUMN "public"."sys_notices"."content" IS '标题';
COMMENT ON COLUMN "public"."sys_notices"."notice_type" IS '通知类型';
COMMENT ON COLUMN "public"."sys_notices"."organization_id" IS '部门Id,部门及子部门';

-- ----------------------------
-- Records of sys_notices
-- ----------------------------
INSERT INTO "public"."sys_notices" VALUES (1, '关于学习交流的通知', '<p>发布<b>入群</b>通知&nbsp;<span style=\"color: var(--el-text-color-regular);\">467890197, 交流学习</span><span style=\"font-size: 14px; color: var(--el-text-color-regular);\"></span></p>', '1', 0, '2021-12-26 15:29:25', '2021-12-26 16:19:48', NULL, 'panda');

INSERT INTO "public"."sys_notices" VALUES (3, '版本更新通知：任务功能，通知功能完成', '<p><span style=\"font-size: 14px; color: var(--el-text-color-regular);\">', '1', 0, '2021-12-26 17:33:47', '2021-12-26 17:33:47', NULL, 'panda');

-- ----------------------------
-- Table structure for sys_organizations
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_organizations";
CREATE TABLE "public"."sys_organizations" (
  "organization_id" int8 NOT NULL,
  "parent_id" int4,
  "organization_path" varchar(255) COLLATE "pg_catalog"."default",
  "organization_name" varchar(128) COLLATE "pg_catalog"."default",
  "sort" int4,
  "leader" varchar(64) COLLATE "pg_catalog"."default",
  "phone" varchar(11) COLLATE "pg_catalog"."default",
  "email" varchar(64) COLLATE "pg_catalog"."default",
  "status" char(1) COLLATE "pg_catalog"."default",
  "create_by" varchar(64) COLLATE "pg_catalog"."default",
  "update_by" varchar(64) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."sys_organizations"."parent_id" IS '上级组织';
COMMENT ON COLUMN "public"."sys_organizations"."organization_path" IS '组织路径';
COMMENT ON COLUMN "public"."sys_organizations"."organization_name" IS '组织名称';
COMMENT ON COLUMN "public"."sys_organizations"."sort" IS '排序';
COMMENT ON COLUMN "public"."sys_organizations"."leader" IS '负责人';
COMMENT ON COLUMN "public"."sys_organizations"."phone" IS '手机';
COMMENT ON COLUMN "public"."sys_organizations"."email" IS '邮箱';
COMMENT ON COLUMN "public"."sys_organizations"."status" IS '状态';
COMMENT ON COLUMN "public"."sys_organizations"."create_by" IS '创建人';
COMMENT ON COLUMN "public"."sys_organizations"."update_by" IS '修改人';

-- ----------------------------
-- Records of sys_organizations
-- ----------------------------
INSERT INTO "public"."sys_organizations" VALUES (2, 0, '/0/2', '熊猫科技', 0, 'xm', '18353366836', '342@qq.com', '0', 'admin', 'admin', '2021-12-01 17:31:53', '2021-12-02 08:56:19', NULL);
INSERT INTO "public"."sys_organizations" VALUES (3, 2, '/0/2/3', '研发部', 1, 'panda', '18353366543', 'ewr@qq.com', '0', 'admin', 'admin', '2021-12-01 17:37:43', '2021-12-02 08:55:56', NULL);
INSERT INTO "public"."sys_organizations" VALUES (7, 2, '/0/2/7', '营销部', 2, 'panda', '18353333333', '342@qq.com', '0', 'panda', 'panda', '2021-12-24 10:46:24', '2021-12-24 10:47:15', NULL);

-- ----------------------------
-- Table structure for sys_posts
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_posts";
CREATE TABLE "public"."sys_posts" (
  "post_id" int8 NOT NULL,
  "post_name" varchar(128) COLLATE "pg_catalog"."default",
  "post_code" varchar(128) COLLATE "pg_catalog"."default",
  "sort" int4,
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
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
INSERT INTO "public"."sys_posts" VALUES (1, '首席执行官', 'CEO', 1, '0', '首席执行官', 'panda', NULL, '2021-12-02 09:21:44', '2022-07-16 17:36:32', NULL);
INSERT INTO "public"."sys_posts" VALUES (4, '首席技术执行官', 'CTO', 2, '0', NULL, 'panda', NULL, '2021-12-02 09:21:44', '2022-07-16 17:37:42', NULL);

-- ----------------------------
-- Table structure for sys_role_menus
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role_menus";
CREATE TABLE "public"."sys_role_menus" (
  "id" int8 NOT NULL,
  "role_id" int4,
  "menu_id" int4,
  "role_name" varchar(128) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO "public"."sys_role_menus" VALUES (6590, 5, 1, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6591, 5, 3, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6592, 5, 4, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6593, 5, 5, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6594, 5, 6, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6595, 5, 7, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6596, 5, 8, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6597, 5, 9, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6598, 5, 10, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6599, 5, 11, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6600, 5, 13, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6601, 5, 14, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6602, 5, 15, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6603, 5, 16, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6604, 5, 17, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6605, 5, 18, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6606, 5, 19, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6607, 5, 20, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6608, 5, 21, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6609, 5, 22, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6610, 5, 23, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6611, 5, 24, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6612, 5, 25, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6613, 5, 26, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6614, 5, 28, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6615, 5, 29, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6616, 5, 30, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6617, 5, 31, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6618, 5, 32, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6619, 5, 33, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6620, 5, 34, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6621, 5, 35, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6622, 5, 36, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6623, 5, 37, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6624, 5, 38, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6625, 5, 39, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6626, 5, 40, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6627, 5, 41, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (6628, 5, 42, 'test');
INSERT INTO "public"."sys_role_menus" VALUES (7377, 2, 1, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7378, 2, 3, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7379, 2, 4, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7380, 2, 5, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7381, 2, 6, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7382, 2, 7, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7383, 2, 8, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7384, 2, 9, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7385, 2, 10, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7386, 2, 11, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7387, 2, 12, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7388, 2, 13, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7389, 2, 14, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7390, 2, 15, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7391, 2, 16, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7392, 2, 17, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7393, 2, 18, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7394, 2, 19, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7395, 2, 20, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7396, 2, 21, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7397, 2, 22, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7398, 2, 23, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7399, 2, 24, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7400, 2, 25, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7401, 2, 26, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7402, 2, 28, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7403, 2, 29, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7404, 2, 30, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7405, 2, 31, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7406, 2, 32, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7407, 2, 33, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7408, 2, 34, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7409, 2, 35, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7410, 2, 36, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7411, 2, 37, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7412, 2, 38, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7413, 2, 39, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7414, 2, 40, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7415, 2, 41, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7416, 2, 42, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7417, 2, 43, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7418, 2, 44, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7419, 2, 45, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7420, 2, 46, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7421, 2, 47, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7422, 2, 49, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7423, 2, 50, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7424, 2, 51, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7425, 2, 52, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7426, 2, 55, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7427, 2, 59, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7428, 2, 60, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7429, 2, 63, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7430, 2, 64, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7431, 2, 69, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7432, 2, 70, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7433, 2, 71, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7434, 2, 72, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7435, 2, 73, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7436, 2, 74, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7437, 2, 75, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7438, 2, 76, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7439, 2, 77, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7440, 2, 95, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7441, 2, 96, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7442, 2, 97, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7443, 2, 98, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7444, 2, 100, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7445, 2, 102, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7446, 2, 103, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7447, 2, 104, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7448, 2, 105, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7449, 2, 106, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7450, 2, 107, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7451, 2, 114, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7452, 2, 115, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7453, 2, 116, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7454, 2, 117, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7455, 2, 118, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7456, 2, 119, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7457, 2, 120, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7458, 2, 121, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7459, 2, 122, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7460, 2, 131, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7461, 2, 132, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7462, 2, 133, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7463, 2, 134, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7464, 2, 135, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7465, 2, 136, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7466, 2, 137, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7467, 2, 138, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7468, 2, 139, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7469, 2, 140, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7470, 2, 141, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7471, 2, 142, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7472, 2, 143, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7473, 2, 144, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7474, 2, 145, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7475, 2, 146, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7476, 2, 147, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7477, 2, 148, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7478, 2, 149, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7479, 2, 152, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7480, 2, 153, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7481, 2, 154, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7482, 2, 155, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7483, 2, 156, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7484, 2, 157, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7485, 2, 158, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7486, 2, 159, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7487, 2, 160, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7488, 2, 161, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7489, 2, 164, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7490, 2, 165, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7491, 2, 166, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7492, 2, 167, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7493, 2, 168, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7494, 2, 169, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7495, 2, 170, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7496, 2, 171, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7497, 2, 175, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7498, 2, 176, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7499, 2, 177, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (7752, 1, 1, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7753, 1, 3, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7754, 1, 4, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7755, 1, 5, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7756, 1, 6, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7757, 1, 7, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7758, 1, 8, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7759, 1, 9, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7760, 1, 10, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7761, 1, 11, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7762, 1, 12, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7763, 1, 13, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7764, 1, 14, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7765, 1, 15, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7766, 1, 16, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7767, 1, 17, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7768, 1, 18, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7769, 1, 19, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7770, 1, 20, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7771, 1, 21, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7772, 1, 22, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7773, 1, 23, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7774, 1, 24, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7775, 1, 25, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7776, 1, 26, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7777, 1, 28, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7778, 1, 29, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7779, 1, 30, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7780, 1, 31, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7781, 1, 32, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7782, 1, 33, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7783, 1, 34, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7784, 1, 35, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7785, 1, 36, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7786, 1, 37, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7787, 1, 38, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7788, 1, 39, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7789, 1, 40, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7790, 1, 41, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7791, 1, 42, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7792, 1, 43, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7793, 1, 44, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7794, 1, 45, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7795, 1, 46, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7796, 1, 47, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7797, 1, 49, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7798, 1, 50, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7799, 1, 51, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7800, 1, 52, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7801, 1, 55, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7802, 1, 59, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7803, 1, 60, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7804, 1, 63, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7805, 1, 64, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7806, 1, 69, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7807, 1, 70, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7808, 1, 71, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7809, 1, 72, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7810, 1, 73, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7811, 1, 74, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7812, 1, 75, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7813, 1, 76, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7814, 1, 77, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7815, 1, 95, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7816, 1, 96, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7817, 1, 97, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7818, 1, 98, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7819, 1, 100, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7820, 1, 102, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7821, 1, 103, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7822, 1, 104, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7823, 1, 105, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7824, 1, 106, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7825, 1, 107, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7826, 1, 114, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7827, 1, 115, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7828, 1, 116, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7829, 1, 117, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7830, 1, 118, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7831, 1, 119, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7832, 1, 120, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7833, 1, 121, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7834, 1, 122, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7835, 1, 131, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7836, 1, 132, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7837, 1, 133, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7838, 1, 134, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7839, 1, 135, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7840, 1, 136, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7841, 1, 137, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7842, 1, 138, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7843, 1, 139, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7844, 1, 140, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7845, 1, 141, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7846, 1, 142, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7847, 1, 143, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7848, 1, 144, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7849, 1, 145, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7850, 1, 146, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7851, 1, 147, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7852, 1, 148, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7853, 1, 149, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7854, 1, 152, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7855, 1, 153, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7856, 1, 154, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7857, 1, 155, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7858, 1, 156, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7859, 1, 157, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7860, 1, 158, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7861, 1, 159, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7862, 1, 160, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7863, 1, 161, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7864, 1, 164, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7865, 1, 165, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7866, 1, 166, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7867, 1, 167, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7868, 1, 168, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7869, 1, 169, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7870, 1, 170, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7871, 1, 171, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7872, 1, 172, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7873, 1, 173, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7874, 1, 174, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7875, 1, 175, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7876, 1, 176, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (7877, 1, 177, 'admin');

-- ----------------------------
-- Table structure for sys_role_organizations
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role_organizations";
CREATE TABLE "public"."sys_role_organizations" (
  "role_id" int4,
  "organization_id" int4,
  "id" int8 NOT NULL
)
;

-- ----------------------------
-- Records of sys_role_organizations
-- ----------------------------
INSERT INTO "public"."sys_role_organizations" VALUES (1, 2, 9);
INSERT INTO "public"."sys_role_organizations" VALUES (1, 3, 10);
INSERT INTO "public"."sys_role_organizations" VALUES (1, 7, 11);
INSERT INTO "public"."sys_role_organizations" VALUES (2, 2, 12);
INSERT INTO "public"."sys_role_organizations" VALUES (2, 3, 13);
INSERT INTO "public"."sys_role_organizations" VALUES (2, 7, 14);

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_roles";
CREATE TABLE "public"."sys_roles" (
  "role_id" int8 NOT NULL,
  "role_name" varchar(128) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "role_key" varchar(128) COLLATE "pg_catalog"."default",
  "data_scope" varchar(1) COLLATE "pg_catalog"."default",
  "role_sort" int4,
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."sys_roles"."role_name" IS '角色名称';
COMMENT ON COLUMN "public"."sys_roles"."status" IS '状态';
COMMENT ON COLUMN "public"."sys_roles"."role_key" IS '角色代码';
COMMENT ON COLUMN "public"."sys_roles"."data_scope" IS '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）';
COMMENT ON COLUMN "public"."sys_roles"."role_sort" IS '角色排序';

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO "public"."sys_roles" VALUES (1, '超管理员', '0', 'admin', '4', 1, 'admin', 'panda', '超级管理', '2021-12-02 16:03:26', '2023-09-25 10:14:10', NULL);
INSERT INTO "public"."sys_roles" VALUES (2, '管理员', '0', 'manage', '1', 2, 'panda', 'panda', NULL, '2021-12-19 16:06:20', '2023-09-20 10:37:46', NULL);
INSERT INTO "public"."sys_roles" VALUES (5, '测试', '0', 'test', '0', 3, 'panda', NULL, NULL, '2023-09-14 11:34:44', '2023-09-14 11:34:44', NULL);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_users";
CREATE TABLE "public"."sys_users" (
  "user_id" int8 NOT NULL,
  "nick_name" varchar(128) COLLATE "pg_catalog"."default",
  "phone" varchar(11) COLLATE "pg_catalog"."default",
  "role_id" int4,
  "salt" varchar(255) COLLATE "pg_catalog"."default",
  "avatar" varchar(255) COLLATE "pg_catalog"."default",
  "sex" varchar(255) COLLATE "pg_catalog"."default",
  "email" varchar(128) COLLATE "pg_catalog"."default",
  "organization_id" int4,
  "post_id" int4,
  "create_by" varchar(128) COLLATE "pg_catalog"."default",
  "update_by" varchar(128) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "delete_time" timestamp(6),
  "username" varchar(64) COLLATE "pg_catalog"."default",
  "password" varchar(128) COLLATE "pg_catalog"."default",
  "role_ids" varchar(255) COLLATE "pg_catalog"."default",
  "post_ids" varchar(255) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_users"."role_ids" IS '多角色';
COMMENT ON COLUMN "public"."sys_users"."post_ids" IS '多岗位';

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO "public"."sys_users" VALUES (1, 'pandax', '13818888888', 1, NULL, NULL, '0', '1@qq.com', 2, 4, 'panda', '1', NULL, '0', '2021-12-03 09:46:55', '2022-02-09 13:28:49', NULL, 'panda', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '1', '1,4');
INSERT INTO "public"."sys_users" VALUES (3, '测试用户', '18435234356', 2, NULL, NULL, '0', '342@163.com', 3, 1, 'test', NULL, NULL, '0', '2021-12-06 15:16:53', '2022-05-10 19:19:25', NULL, 'test', '$2a$10$4cHTracxWJLdhMmazvbm1urKyD3v5N2AYxAFtNYh6juU39kgae73e', '2', '1,4');
INSERT INTO "public"."sys_users" VALUES (4, 'panda', '18353366912', 2, NULL, NULL, '0', '2417920382@qq.com', 2, 4, 'panda', NULL, NULL, '0', '2021-12-19 15:58:09', '2021-12-19 16:06:54', NULL, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '2', '4,1');

-- ----------------------------
-- Table structure for visual_data_set_field
-- ----------------------------
DROP TABLE IF EXISTS "public"."visual_data_set_field";
CREATE TABLE "public"."visual_data_set_field" (
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "field_id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "table_id" varchar(64) COLLATE "pg_catalog"."default",
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "group_type" varchar(1) COLLATE "pg_catalog"."default",
  "origin_name" varchar(255) COLLATE "pg_catalog"."default",
  "origin_type" varchar(50) COLLATE "pg_catalog"."default",
  "de_name" varchar(255) COLLATE "pg_catalog"."default",
  "de_type" varchar(1) COLLATE "pg_catalog"."default",
  "ext_field" varchar(1) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."visual_data_set_field"."field_id" IS '字段id';
COMMENT ON COLUMN "public"."visual_data_set_field"."table_id" IS '表id';
COMMENT ON COLUMN "public"."visual_data_set_field"."name" IS '字段名(描述)';
COMMENT ON COLUMN "public"."visual_data_set_field"."group_type" IS '维度/指标标识 d:维度，q:指标';
COMMENT ON COLUMN "public"."visual_data_set_field"."origin_name" IS '原始字段名称';
COMMENT ON COLUMN "public"."visual_data_set_field"."origin_type" IS '原始字段类型';
COMMENT ON COLUMN "public"."visual_data_set_field"."de_name" IS '数据源查询名';
COMMENT ON COLUMN "public"."visual_data_set_field"."de_type" IS '数据源字段类型：0-文本，1-时间，2-数值，3-数值小数';
COMMENT ON COLUMN "public"."visual_data_set_field"."ext_field" IS '是否扩展字段 0否 1是';

-- ----------------------------
-- Records of visual_data_set_field
-- ----------------------------

-- ----------------------------
-- Table structure for visual_data_set_table
-- ----------------------------
DROP TABLE IF EXISTS "public"."visual_data_set_table";
CREATE TABLE "public"."visual_data_set_table" (
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "table_id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "data_source_id" varchar(64) COLLATE "pg_catalog"."default",
  "table_type" varchar(64) COLLATE "pg_catalog"."default",
  "info" text COLLATE "pg_catalog"."default",
  "create_by" int8,
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "org_id" int4,
  "owner" varchar(64) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."visual_data_set_table"."table_id" IS '表id';
COMMENT ON COLUMN "public"."visual_data_set_table"."data_source_id" IS '数据圆ID';
COMMENT ON COLUMN "public"."visual_data_set_table"."table_type" IS 'db,sql,excel,union';
COMMENT ON COLUMN "public"."visual_data_set_table"."info" IS '原始表信息';
COMMENT ON COLUMN "public"."visual_data_set_table"."name" IS '名称';
COMMENT ON COLUMN "public"."visual_data_set_table"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."visual_data_set_table"."owner" IS '创建者,所有者';

-- ----------------------------
-- Records of visual_data_set_table
-- ----------------------------

-- ----------------------------
-- Table structure for visual_data_set_table_function
-- ----------------------------
DROP TABLE IF EXISTS "public"."visual_data_set_table_function";
CREATE TABLE "public"."visual_data_set_table_function" (
  "name" varchar(64) COLLATE "pg_catalog"."default",
  "func" varchar(500) COLLATE "pg_catalog"."default",
  "db_type" varchar(255) COLLATE "pg_catalog"."default",
  "desc" text COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."visual_data_set_table_function"."name" IS '函数名称';
COMMENT ON COLUMN "public"."visual_data_set_table_function"."func" IS '函数表达式';
COMMENT ON COLUMN "public"."visual_data_set_table_function"."db_type" IS '所属数据库';
COMMENT ON COLUMN "public"."visual_data_set_table_function"."desc" IS '描述';

-- ----------------------------
-- Records of visual_data_set_table_function
-- ----------------------------
INSERT INTO "public"."visual_data_set_table_function" VALUES ('ABS', 'ABS(x)', 'MySQL', '返回x的绝对值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('PI', 'PI()', 'MySQL', '返回圆周率π，默认显示6位小数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('SQRT', 'SQRT(x)', 'MySQL', '返回非负数的x的二次方根');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('MOD', 'MOD(x,y)', 'MySQL', '返回x被y除后的余数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('CEIL', 'CEIL(x)', 'MySQL', '返回不小于x的最小整数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('FLOOR', 'FLOOR(x)', 'MySQL', '返回不大于x的最大整数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('ROUND', 'ROUND(x)', 'MySQL', '返回离x最近的整数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('ROUND', 'ROUND(x,y)', 'MySQL', '保留x小数点后y位的值，但截断时要进行四舍五入');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('SIGN', 'SIGN(x)', 'MySQL', '返回参数x的符号，-1表示负数，0表示0，1表示正数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('POW', 'POW(x,y)', 'MySQL', '返回x的y次乘方的值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('EXP', 'EXP(x)', 'MySQL', '返回e的x乘方后的值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('LOG', 'LOG(x)', 'MySQL', '返回x的自然对数，x相对于基数e的对数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('LOG10', 'LOG10(x)', 'MySQL', '返回x的基数为10的对数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('RADIANS', 'RADIANS(x)', 'MySQL', '返回x由角度转化为弧度的值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('DEGREES', 'DEGREES(x)', 'MySQL', '返回x由弧度转化为角度的值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('SIN', 'SIN(x)', 'MySQL', '返回x的正弦，其中x为给定的弧度值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('ASIN', 'ASIN(x)', 'MySQL', '返回x的反正弦值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('COS', 'COS(x)', 'MySQL', '返回x的余弦，其中x为给定的弧度值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('ACOS', 'ACOS(x)', 'MySQL', '返回x的反余弦值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('TAN', 'TAN(x)', 'MySQL', '返回x的正切，其中x为给定的弧度值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('ATAN', 'ATAN(x)', 'MySQL', '返回x的反正切值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('COT', 'COT(x)', 'MySQL', '返回给定弧度值x的余切');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('CHAR_LENGTH', 'CHAR_LENGTH(str)', 'MySQL', '计算字符串字符个数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('TRIM', 'TRIM(s)', 'MySQL', '返回字符串s删除了两边空格之后的字符串');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('LTRIM', 'LTRIM(s)', 'MySQL', '返回字符串s，其左边所有空格被删除');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('RTRIM', 'RTRIM(s)', 'MySQL', '返回字符串s，其右边所有空格被删除');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('REPLACE', 'REPLACE(s,s1,s2)', 'MySQL', '返回一个字符串，用字符串s2替代字符串s中所有的字符串s1');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('SUBSTRING', 'SUBSTRING(s,n,len)', 'MySQL', '获取从字符串s中的第n个位置开始长度为len的字符串');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('CONCAT', 'CONCAT(s1,s2，...)', 'MySQL', '返回连接参数产生的字符串，一个或多个待拼接的内容，任意一个为NULL则返回值为NULL');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('INSERT', 'INSERT(s1,x,len,s2)', 'MySQL', '返回字符串s1，其子字符串起始于位置x，被字符串s2取代len个字符');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('LOWER', 'LOWER(str)', 'MySQL', '将str中的字母全部转换成小写');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('UPPER', 'UPPER(str)', 'MySQL', '将字符串中的字母全部转换成大写');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('LEFT', 'LEFT(s,n)', 'MySQL', '返回字符串s从最左边开始的n个字符');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('RIGHT', 'RIGHT(s,n)', 'MySQL', '返回字符串s从最右边开始的n个字符');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('REPEAT', 'REPEAT(s,n)', 'MySQL', '返回一个由重复字符串s组成的字符串，字符串s的数目等于n');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('SPACE', 'SPACE(n)', 'MySQL', '返回一个由n个空格组成的字符串');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('REVERSE', 'REVERSE(s)', 'MySQL', '将字符串s反转');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('CURDATE', 'CURDATE()', 'MySQL', '将当前日期按照\"YYYY-MM-DD\"或者\"YYYYMMDD\"格式的值返回，具体格式根据函数用在字符串或是数字语境中而定');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('CURRENT_DATE', 'CURRENT_DATE()', 'MySQL', '将当前日期按照\"YYYY-MM-DD\"或者\"YYYYMMDD\"格式的值返回，具体格式根据函数用在字符串或是数字语境中而定');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('NOW', 'NOW()', 'MySQL', '返回当前日期和时间值，格式为\"YYYY_MM-DD HH:MM:SS\"或\"YYYYMMDDHHMMSS\"，具体格式根据函数用在字符串或数字语境中而定');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('SYSDATE', 'SYSDATE()', 'MySQL', '返回当前日期和时间值，格式为\"YYYY_MM-DD HH:MM:SS\"或\"YYYYMMDDHHMMSS\"，具体格式根据函数用在字符串或数字语境中而定');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('LOCALTIME', 'LOCALTIME()', 'MySQL', '返回当前日期和时间值，格式为\"YYYY_MM-DD HH:MM:SS\"或\"YYYYMMDDHHMMSS\"，具体格式根据函数用在字符串或数字语境中而定');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('CURRENT_TIMESTAMP', 'CURRENT_TIMESTAMP()', 'MySQL', '返回当前日期和时间值，格式为\"YYYY_MM-DD HH:MM:SS\"或\"YYYYMMDDHHMMSS\"，具体格式根据函数用在字符串或数字语境中而定');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('UNIX_TIMESTAMP', 'UNIX_TIMESTAMP()', 'MySQL', '返回一个格林尼治标准时间1970-01-01 00:00:00到现在的秒数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('UNIX_TIMESTAMP', 'UNIX_TIMESTAMP(date)', 'MySQL', '返回一个格林尼治标准时间1970-01-01 00:00:00到指定时间的秒数');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('FROM_UNIXTIME', 'FROM_UNIXTIME(date)', 'MySQL', '把UNIX时间戳转换为普通格式的时间');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('CASE', 'CASE expr WHEN v1 THEN r1 [WHEN v2 THEN v2] [ELSE rn] END', 'MySQL', '如果expr等于某个vn，则返回对应位置THEN后面的结果，如果与所有值都不想等，则返回ELSE后面的rn');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('IF', 'IF(expr,v1,v2)', 'MySQL', '如果expr是TRUE则返回v1，否则返回v2');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('IFNULL', 'IFNULL(v1,v2)', 'MySQL', '如果v1不为NULL，则返回v1，否则返回v2');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('datalength', 'datalength(s)', 'SqlServer', '返回字符串包含字符数,但不包含后面的空格');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('substring', 'substring(expression,start,length)', 'SqlServer', '取子串');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('ltrim', 'ltrim(expression)', 'SqlServer', '把字符串头部的空格去掉');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('rtrim', 'rtrim(expression)', 'SqlServer', '把字符串尾部的空格去掉');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('upper', 'substring(expression)', 'SqlServer', '转为大写');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('lower', 'substring(expression)', 'SqlServer', '转为小写');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('reverse', 'reverse(expression)', 'SqlServer', '反转字符串');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('pi', 'pi()', 'SqlServer', 'π');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('abs', 'abs(numeric_expr)', 'SqlServer', '求绝对值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('sqrt', 'sqrt(float_expr)', 'SqlServer', '平方根');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('power', 'power(numeric_expr,power) ', 'SqlServer', '返回power次方');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('getdate', 'getdate() ', 'SqlServer', '返回日期');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('getutcdate', 'getutcdate() ', 'SqlServer', '获取utc时间');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('day', 'day(getdate()) ', 'SqlServer', '取出天');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('abs', 'abs(x)', 'PostgreSQL', '绝对值');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('sin', 'sin(x)', 'PostgreSQL', '正弦');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('sqrt', 'sqrt(double/numeric)', 'PostgreSQL', '平方根');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('lower', 'lower(string)', 'PostgreSQL', '把字串转化为小写');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('length', 'length(string text)', 'PostgreSQL', 'string中字符的数目');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('md5', 'md5(string text)', 'PostgreSQL', '计算给出string的MD5散列，以十六进制返回结果');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('strpos', 'strpos(string, substring)', 'PostgreSQL', '声明的子字串的位置');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('to_char', 'to_char(timestamp, text)', 'PostgreSQL', '将时间戳转换为字符串');
INSERT INTO "public"."visual_data_set_table_function" VALUES ('to_timestamp', 'to_timestamp(double precision)', 'PostgreSQL', '把UNIX纪元转换成时间戳');

-- ----------------------------
-- Table structure for visual_data_source
-- ----------------------------
DROP TABLE IF EXISTS "public"."visual_data_source";
CREATE TABLE "public"."visual_data_source" (
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "source_id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "source_type" varchar(50) COLLATE "pg_catalog"."default",
  "source_name" varchar(50) COLLATE "pg_catalog"."default",
  "source_comment" varchar(50) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "db" text COLLATE "pg_catalog"."default",
  "create_by" int8,
  "org_id" int4,
  "owner" varchar(64) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."visual_data_source"."source_id" IS '数据源Id';
COMMENT ON COLUMN "public"."visual_data_source"."source_type" IS '数据源类型';
COMMENT ON COLUMN "public"."visual_data_source"."source_name" IS '数据源名称';
COMMENT ON COLUMN "public"."visual_data_source"."source_comment" IS '数据源描述';
COMMENT ON COLUMN "public"."visual_data_source"."status" IS '数据源状态 1:在线 0连接异常';
COMMENT ON COLUMN "public"."visual_data_source"."db" IS '详细信息';
COMMENT ON COLUMN "public"."visual_data_source"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."visual_data_source"."owner" IS '创建者,所有者';

-- ----------------------------
-- Records of visual_data_source
-- ----------------------------

-- ----------------------------
-- Table structure for visual_screen
-- ----------------------------
DROP TABLE IF EXISTS "public"."visual_screen";
CREATE TABLE "public"."visual_screen" (
  "id" varchar(191) COLLATE "pg_catalog"."default" NOT NULL,
  "group_id" int4,
  "screen_name" varchar(50) COLLATE "pg_catalog"."default",
  "screen_data_json" text COLLATE "pg_catalog"."default",
  "screen_base64" text COLLATE "pg_catalog"."default",
  "screen_remark" varchar(255) COLLATE "pg_catalog"."default",
  "status" varchar(1) COLLATE "pg_catalog"."default",
  "create_time" timestamp(6),
  "update_time" timestamp(6),
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "org_id" int4
)
;
COMMENT ON COLUMN "public"."visual_screen"."group_id" IS '分组Id';
COMMENT ON COLUMN "public"."visual_screen"."screen_name" IS '名称';
COMMENT ON COLUMN "public"."visual_screen"."screen_data_json" IS 'Json数据';
COMMENT ON COLUMN "public"."visual_screen"."screen_base64" IS 'Base64缩略图';
COMMENT ON COLUMN "public"."visual_screen"."screen_remark" IS '说明';
COMMENT ON COLUMN "public"."visual_screen"."status" IS '状态';
COMMENT ON COLUMN "public"."visual_screen"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."visual_screen"."org_id" IS '机构ID';

-- ----------------------------
-- Records of visual_screen
-- ----------------------------

-- ----------------------------
-- Table structure for visual_screen_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."visual_screen_group";
CREATE TABLE "public"."visual_screen_group" (
  "id" int4 NOT NULL,
  "org_id" int4,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "pid" int4,
  "path" varchar(64) COLLATE "pg_catalog"."default",
  "sort" int4,
  "status" char(1) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."visual_screen_group"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."visual_screen_group"."owner" IS '创建者,所有者';
COMMENT ON COLUMN "public"."visual_screen_group"."name" IS '分组名';
COMMENT ON COLUMN "public"."visual_screen_group"."pid" IS '父Id';
COMMENT ON COLUMN "public"."visual_screen_group"."path" IS '路径';
COMMENT ON COLUMN "public"."visual_screen_group"."sort" IS '排序';
COMMENT ON COLUMN "public"."visual_screen_group"."status" IS '状态';

-- ----------------------------
-- Records of visual_screen_group
-- ----------------------------


-- ----------------------------
-- Table structure for visual_screen_image
-- ----------------------------
DROP TABLE IF EXISTS "public"."visual_screen_image";
CREATE TABLE "public"."visual_screen_image" (
  "org_id" int4,
  "owner" varchar(64) COLLATE "pg_catalog"."default",
  "file_name" varchar(255) COLLATE "pg_catalog"."default",
  "file_path" varchar(255) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."visual_screen_image"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."visual_screen_image"."owner" IS '创建者,所有者';

-- ----------------------------
-- Records of visual_screen_image
-- ----------------------------

-- ----------------------------
-- Indexes structure for table casbin_rule
-- ----------------------------
CREATE UNIQUE INDEX "idx_casbin_rule" ON "public"."casbin_rule" USING btree (
  "ptype" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v0" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v1" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v2" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v3" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "v4" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table casbin_rule
-- ----------------------------
ALTER TABLE "public"."casbin_rule" ADD CONSTRAINT "casbin_rule_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table dev_gen_table_columns
-- ----------------------------
ALTER TABLE "public"."dev_gen_table_columns" ADD CONSTRAINT "dev_gen_table_columns_pkey" PRIMARY KEY ("column_id");

-- ----------------------------
-- Primary Key structure for table dev_gen_tables
-- ----------------------------
ALTER TABLE "public"."dev_gen_tables" ADD CONSTRAINT "_copy_31" PRIMARY KEY ("table_id");

-- ----------------------------
-- Primary Key structure for table device_alarms
-- ----------------------------
ALTER TABLE "public"."device_alarms" ADD CONSTRAINT "_copy_30" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table device_cmd_log
-- ----------------------------
ALTER TABLE "public"."device_cmd_log" ADD CONSTRAINT "_copy_29" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table device_groups
-- ----------------------------
ALTER TABLE "public"."device_groups" ADD CONSTRAINT "_copy_28" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table devices
-- ----------------------------
CREATE INDEX "fk_devices_device_group" ON "public"."devices" USING btree (
  "gid" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "fk_devices_product" ON "public"."devices" USING btree (
  "pid" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table devices
-- ----------------------------
ALTER TABLE "public"."devices" ADD CONSTRAINT "_copy_27" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table job_logs
-- ----------------------------
ALTER TABLE "public"."job_logs" ADD CONSTRAINT "_copy_26" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table jobs
-- ----------------------------
ALTER TABLE "public"."jobs" ADD CONSTRAINT "_copy_25" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table log_logins
-- ----------------------------
ALTER TABLE "public"."log_logins" ADD CONSTRAINT "_copy_24" PRIMARY KEY ("info_id");

-- ----------------------------
-- Primary Key structure for table log_opers
-- ----------------------------
ALTER TABLE "public"."log_opers" ADD CONSTRAINT "_copy_23" PRIMARY KEY ("oper_id");

-- ----------------------------
-- Primary Key structure for table product_categories
-- ----------------------------
ALTER TABLE "public"."product_categories" ADD CONSTRAINT "_copy_22" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table product_ota
-- ----------------------------
ALTER TABLE "public"."product_ota" ADD CONSTRAINT "_copy_21" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table product_templates
-- ----------------------------
ALTER TABLE "public"."product_templates" ADD CONSTRAINT "_copy_20" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table products
-- ----------------------------
CREATE INDEX "fk_products_product_category" ON "public"."products" USING btree (
  "product_category_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD CONSTRAINT "_copy_19" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table rule_chain
-- ----------------------------
ALTER TABLE "public"."rule_chain" ADD CONSTRAINT "_copy_18" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_apis
-- ----------------------------
ALTER TABLE "public"."sys_apis" ADD CONSTRAINT "_copy_17" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_configs
-- ----------------------------
ALTER TABLE "public"."sys_configs" ADD CONSTRAINT "_copy_16" PRIMARY KEY ("config_id");

-- ----------------------------
-- Primary Key structure for table sys_dict_data
-- ----------------------------
ALTER TABLE "public"."sys_dict_data" ADD CONSTRAINT "_copy_15" PRIMARY KEY ("dict_code");

-- ----------------------------
-- Primary Key structure for table sys_dict_types
-- ----------------------------
ALTER TABLE "public"."sys_dict_types" ADD CONSTRAINT "_copy_14" PRIMARY KEY ("dict_id");

-- ----------------------------
-- Primary Key structure for table sys_menus
-- ----------------------------
ALTER TABLE "public"."sys_menus" ADD CONSTRAINT "_copy_13" PRIMARY KEY ("menu_id");

-- ----------------------------
-- Primary Key structure for table sys_notices
-- ----------------------------
ALTER TABLE "public"."sys_notices" ADD CONSTRAINT "_copy_12" PRIMARY KEY ("notice_id");

-- ----------------------------
-- Primary Key structure for table sys_organizations
-- ----------------------------
ALTER TABLE "public"."sys_organizations" ADD CONSTRAINT "_copy_11" PRIMARY KEY ("organization_id");

-- ----------------------------
-- Primary Key structure for table sys_posts
-- ----------------------------
ALTER TABLE "public"."sys_posts" ADD CONSTRAINT "_copy_10" PRIMARY KEY ("post_id");

-- ----------------------------
-- Primary Key structure for table sys_role_menus
-- ----------------------------
ALTER TABLE "public"."sys_role_menus" ADD CONSTRAINT "_copy_9" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_role_organizations
-- ----------------------------
ALTER TABLE "public"."sys_role_organizations" ADD CONSTRAINT "_copy_8" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_roles
-- ----------------------------
ALTER TABLE "public"."sys_roles" ADD CONSTRAINT "_copy_7" PRIMARY KEY ("role_id");

-- ----------------------------
-- Primary Key structure for table sys_users
-- ----------------------------
ALTER TABLE "public"."sys_users" ADD CONSTRAINT "_copy_6" PRIMARY KEY ("user_id");

-- ----------------------------
-- Primary Key structure for table visual_data_set_field
-- ----------------------------
ALTER TABLE "public"."visual_data_set_field" ADD CONSTRAINT "_copy_5" PRIMARY KEY ("field_id");

-- ----------------------------
-- Primary Key structure for table visual_data_set_table
-- ----------------------------
ALTER TABLE "public"."visual_data_set_table" ADD CONSTRAINT "_copy_4" PRIMARY KEY ("table_id");

-- ----------------------------
-- Primary Key structure for table visual_data_source
-- ----------------------------
ALTER TABLE "public"."visual_data_source" ADD CONSTRAINT "_copy_3" PRIMARY KEY ("source_id");

-- ----------------------------
-- Primary Key structure for table visual_screen
-- ----------------------------
ALTER TABLE "public"."visual_screen" ADD CONSTRAINT "_copy_2" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table visual_screen_group
-- ----------------------------
ALTER TABLE "public"."visual_screen_group" ADD CONSTRAINT "_copy_1" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table devices
-- ----------------------------
ALTER TABLE "public"."devices" ADD CONSTRAINT "fk_devices_device_group" FOREIGN KEY ("gid") REFERENCES "public"."device_groups" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."devices" ADD CONSTRAINT "fk_devices_product" FOREIGN KEY ("pid") REFERENCES "public"."products" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD CONSTRAINT "fk_products_product_category" FOREIGN KEY ("product_category_id") REFERENCES "public"."product_categories" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
