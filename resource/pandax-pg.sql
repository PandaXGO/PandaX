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


CREATE SEQUENCE IF NOT EXISTS "casbin_rule_id_seq";

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS "public"."casbin_rule";
CREATE SEQUENCE IF NOT EXISTS "casbin_rule_id_seq";
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
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/gen/code/:tableId', 'GET', '', '', 2190);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/gen/configure/:tableId', 'GET', '', '', 2191);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/gen/preview/:tableId', 'GET', '', '', 2189);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/table', 'POST', '', '', 2186);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/table', 'PUT', '', '', 2187);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/table/:tableId', 'DELETE', '', '', 2188);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/table/db/list', 'GET', '', '', 2181);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/table/info/:tableId', 'GET', '', '', 2183);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/table/info/tableName', 'GET', '', '', 2184);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/table/list', 'GET', '', '', 2182);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/develop/code/table/tableTree', 'GET', '', '', 2185);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/job', 'POST', '', '', 2193);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/job', 'PUT', '', '', 2194);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/job/:jobId', 'DELETE', '', '', 2196);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/job/:jobId', 'GET', '', '', 2195);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/job/changeStatus', 'PUT', '', '', 2199);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/job/list', 'GET', '', '', 2192);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/job/start/:jobId', 'GET', '', '', 2198);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/job/stop/:jobId', 'GET', '', '', 2197);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logJob/:logId', 'DELETE', '', '', 2208);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logJob/all', 'DELETE', '', '', 2207);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logJob/list', 'GET', '', '', 2206);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logLogin/:infoId', 'DELETE', '', '', 2201);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logLogin/all', 'DELETE', '', '', 2202);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logLogin/list', 'GET', '', '', 2200);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logOper/:operId', 'DELETE', '', '', 2204);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logOper/all', 'DELETE', '', '', 2205);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/log/logOper/list', 'GET', '', '', 2203);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/email', 'POST', '', '', 2211);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/email', 'PUT', '', '', 2212);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/email/:mailId', 'DELETE', '', '', 2213);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/email/:mailId', 'GET', '', '', 2210);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/email/changeStatus', 'PUT', '', '', 2214);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/email/debugMail', 'POST', '', '', 2215);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/email/list', 'GET', '', '', 2209);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/oss', 'POST', '', '', 2231);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/oss', 'PUT', '', '', 2232);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/oss/:ossId', 'DELETE', '', '', 2233);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/oss/:ossId', 'GET', '', '', 2230);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/oss/changeStatus', 'PUT', '', '', 2234);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/oss/list', 'GET', '', '', 2229);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/resource/oss/uploadFile', 'POST', '', '', 2235);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/api', 'POST', '', '', 2153);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/api', 'PUT', '', '', 2154);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/api/:id', 'DELETE', '', '', 2155);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/api/:id', 'GET', '', '', 2152);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/api/all', 'GET', '', '', 2150);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/api/getPolicyPathByRoleId', 'GET', '', '', 2151);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/api/list', 'GET', '', '', 2149);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/config', 'POST', '', '', 2159);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/config', 'PUT', '', '', 2160);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/config/:configId', 'DELETE', '', '', 2161);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/config/:configId', 'GET', '', '', 2158);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/config/configKey', 'GET', '', '', 2157);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/config/list', 'GET', '', '', 2156);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dept', 'POST', '', '', 2166);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dept', 'PUT', '', '', 2167);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dept/:deptId', 'DELETE', '', '', 2168);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dept/:deptId', 'GET', '', '', 2163);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dept/deptTree', 'GET', '', '', 2165);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dept/list', 'GET', '', '', 2162);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', 2164);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/data', 'POST', '', '', 2178);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/data', 'PUT', '', '', 2179);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/data/:dictCode', 'DELETE', '', '', 2180);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/data/:dictCode', 'GET', '', '', 2177);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/data/list', 'GET', '', '', 2175);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/data/type', 'GET', '', '', 2176);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/type', 'POST', '', '', 2171);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/type', 'PUT', '', '', 2172);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/type/:dictId', 'DELETE', '', '', 2173);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/type/:dictId', 'GET', '', '', 2170);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/type/export', 'GET', '', '', 2174);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/dict/type/list', 'GET', '', '', 2169);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu', 'POST', '', '', 2222);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu', 'PUT', '', '', 2223);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu/:menuId', 'DELETE', '', '', 2224);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu/:menuId', 'GET', '', '', 2221);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu/list', 'GET', '', '', 2220);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu/menuPaths', 'GET', '', '', 2219);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu/menuRole', 'GET', '', '', 2217);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu/menuTreeSelect', 'GET', '', '', 2216);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', 2218);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/notice', 'POST', '', '', 2226);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/notice', 'PUT', '', '', 2227);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/notice/:noticeId', 'DELETE', '', '', 2228);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/notice/list', 'GET', '', '', 2225);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/post', 'POST', '', '', 2238);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/post', 'PUT', '', '', 2239);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/post/:postId', 'DELETE', '', '', 2240);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/post/:postId', 'GET', '', '', 2237);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/post/list', 'GET', '', '', 2236);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/role', 'POST', '', '', 2243);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/role', 'PUT', '', '', 2244);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/role/:roleId', 'DELETE', '', '', 2245);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/role/:roleId', 'GET', '', '', 2242);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/role/changeStatus', 'PUT', '', '', 2246);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/role/dataScope', 'PUT', '', '', 2247);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/role/export', 'GET', '', '', 2248);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/role/list', 'GET', '', '', 2241);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/tenant', 'POST', '', '', 2251);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/tenant', 'PUT', '', '', 2252);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/tenant/:tenantId', 'DELETE', '', '', 2253);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/tenant/:tenantId', 'GET', '', '', 2250);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/tenant/list', 'GET', '', '', 2249);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/tenant/lists', 'GET', '', '', 2254);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user', 'POST', '', '', 2263);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user', 'PUT', '', '', 2264);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/:userId', 'DELETE', '', '', 2257);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/avatar', 'POST', '', '', 2258);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/changeStatus', 'PUT', '', '', 2256);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/export', 'GET', '', '', 2265);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/getById/:userId', 'GET', '', '', 2260);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/getInit', 'GET', '', '', 2261);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/getRoPo', 'GET', '', '', 2262);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/list', 'GET', '', '', 2255);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'admin', '/system/user/pwd', 'PUT', '', '', 2259);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/develop/code/gen/preview/:tableId', 'GET', '', '', 2287);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/develop/code/table/db/list', 'GET', '', '', 2282);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/develop/code/table/info/:tableId', 'GET', '', '', 2284);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/develop/code/table/info/tableName', 'GET', '', '', 2285);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/develop/code/table/list', 'GET', '', '', 2283);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/develop/code/table/tableTree', 'GET', '', '', 2286);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/job/:jobId', 'GET', '', '', 2289);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/job/list', 'GET', '', '', 2288);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/log/logJob/list', 'GET', '', '', 2292);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/log/logLogin/list', 'GET', '', '', 2290);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/log/logOper/list', 'GET', '', '', 2291);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/resource/email/:mailId', 'GET', '', '', 2294);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/resource/email/list', 'GET', '', '', 2293);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/resource/oss/:ossId', 'GET', '', '', 2303);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/resource/oss/list', 'GET', '', '', 2302);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/api/:id', 'GET', '', '', 2269);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/api/all', 'GET', '', '', 2267);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/api/getPolicyPathByRoleId', 'GET', '', '', 2268);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/api/list', 'GET', '', '', 2266);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/config/:configId', 'GET', '', '', 2272);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/config/configKey', 'GET', '', '', 2271);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/config/list', 'GET', '', '', 2270);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dept/:deptId', 'GET', '', '', 2274);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dept/deptTree', 'GET', '', '', 2276);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dept/list', 'GET', '', '', 2273);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', 2275);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dict/data/:dictCode', 'GET', '', '', 2281);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dict/data/list', 'GET', '', '', 2279);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dict/data/type', 'GET', '', '', 2280);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dict/type/:dictId', 'GET', '', '', 2278);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/dict/type/list', 'GET', '', '', 2277);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/menu/:menuId', 'GET', '', '', 2300);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/menu/list', 'GET', '', '', 2299);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/menu/menuPaths', 'GET', '', '', 2298);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/menu/menuRole', 'GET', '', '', 2296);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/menu/menuTreeSelect', 'GET', '', '', 2295);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', 2297);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/notice/list', 'GET', '', '', 2301);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/post/:postId', 'GET', '', '', 2305);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/post/list', 'GET', '', '', 2304);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/role/:roleId', 'GET', '', '', 2307);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/role/list', 'GET', '', '', 2306);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/tenant/:tenantId', 'GET', '', '', 2309);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/tenant/list', 'GET', '', '', 2308);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/tenant/lists', 'GET', '', '', 2310);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/user/getById/:userId', 'GET', '', '', 2312);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/user/getInit', 'GET', '', '', 2313);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/user/getRoPo', 'GET', '', '', 2314);
INSERT INTO "public"."casbin_rule" VALUES ('p', '1', 'manage', '/system/user/list', 'GET', '', '', 2311);


-- ----------------------------
-- Table structure for dev_gen_table_columns
-- ----------------------------
DROP TABLE IF EXISTS "public"."dev_gen_table_columns";
CREATE SEQUENCE IF NOT EXISTS "dev_gen_table_columns_column_id_seq";
CREATE TABLE "public"."dev_gen_table_columns" (
   "column_id" int8 NOT NULL DEFAULT nextval('dev_gen_table_columns_column_id_seq'::regclass),
   "table_id" int8,
   "table_name" text COLLATE "pg_catalog"."default",
   "column_name" text COLLATE "pg_catalog"."default",
   "column_comment" text COLLATE "pg_catalog"."default",
   "column_type" text COLLATE "pg_catalog"."default",
   "column_key" text COLLATE "pg_catalog"."default",
   "go_type" text COLLATE "pg_catalog"."default",
   "go_field" text COLLATE "pg_catalog"."default",
   "json_field" text COLLATE "pg_catalog"."default",
   "html_field" text COLLATE "pg_catalog"."default",
   "is_pk" text COLLATE "pg_catalog"."default",
   "is_increment" text COLLATE "pg_catalog"."default",
   "is_required" text COLLATE "pg_catalog"."default",
   "is_insert" text COLLATE "pg_catalog"."default",
   "is_edit" text COLLATE "pg_catalog"."default",
   "is_list" text COLLATE "pg_catalog"."default",
   "is_query" text COLLATE "pg_catalog"."default",
   "query_type" text COLLATE "pg_catalog"."default",
   "html_type" text COLLATE "pg_catalog"."default",
   "dict_type" text COLLATE "pg_catalog"."default",
   "sort" int8,
   "link_table_name" text COLLATE "pg_catalog"."default",
   "link_table_class" text COLLATE "pg_catalog"."default",
   "link_table_package" text COLLATE "pg_catalog"."default",
   "link_label_id" text COLLATE "pg_catalog"."default",
   "link_label_name" text COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of dev_gen_table_columns
-- ----------------------------

-- ----------------------------
-- Table structure for dev_gen_tables
-- ----------------------------
DROP TABLE IF EXISTS "public"."dev_gen_tables";
CREATE SEQUENCE IF NOT EXISTS "dev_gen_tables_table_id_seq";
CREATE TABLE "public"."dev_gen_tables" (
  "table_id" int8 NOT NULL DEFAULT nextval('dev_gen_tables_table_id_seq'::regclass),
  "table_name" text COLLATE "pg_catalog"."default",
  "table_comment" text COLLATE "pg_catalog"."default",
  "class_name" text COLLATE "pg_catalog"."default",
  "tpl_category" text COLLATE "pg_catalog"."default",
  "package_name" text COLLATE "pg_catalog"."default",
  "module_name" text COLLATE "pg_catalog"."default",
  "business_name" text COLLATE "pg_catalog"."default",
  "function_name" text COLLATE "pg_catalog"."default",
  "function_author" text COLLATE "pg_catalog"."default",
  "options" text COLLATE "pg_catalog"."default",
  "remark" text COLLATE "pg_catalog"."default",
  "pk_column" text COLLATE "pg_catalog"."default",
  "pk_go_field" text COLLATE "pg_catalog"."default",
  "pk_json_field" text COLLATE "pg_catalog"."default",
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;

-- ----------------------------
-- Records of dev_gen_tables
-- ----------------------------


-- ----------------------------
-- Table structure for log_jobs
-- ----------------------------
DROP TABLE IF EXISTS "public"."log_jobs";
CREATE SEQUENCE IF NOT EXISTS "log_jobs_log_id_seq";
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
CREATE SEQUENCE IF NOT EXISTS "log_logins_info_id_seq";
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
CREATE SEQUENCE IF NOT EXISTS "log_opers_oper_id_seq";
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
CREATE SEQUENCE IF NOT EXISTS "sys_apis_id_seq";
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
INSERT INTO "public"."sys_apis" VALUES (1, '2021-12-09 09:21:04', '2021-12-09 09:21:04', NULL, '/system/user/list', '查询用户列表（分页）', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (2, '2021-12-09 09:29:36', '2021-12-09 09:29:36', NULL, '/system/user/changeStatus', '修改用户状态', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (3, '2021-12-09 09:34:37', '2021-12-09 09:34:37', NULL, '/system/user/:userId', '删除用户信息', 'user', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (4, '2021-12-09 09:36:43', '2021-12-09 09:36:43', NULL, '/system/dept/list', '获取部门列表', 'dept', 'GET');
INSERT INTO "public"."sys_apis" VALUES (5, '2021-12-09 09:37:31', '2021-12-09 09:37:31', NULL, '/system/dept/:deptId', '获取部门信息', 'dept', 'GET');
INSERT INTO "public"."sys_apis" VALUES (6, '2021-12-09 18:20:32', '2021-12-09 18:20:32', NULL, '/system/user/avatar', '上传头像', 'user', 'POST');
INSERT INTO "public"."sys_apis" VALUES (7, '2021-12-09 18:21:10', '2021-12-09 18:21:10', NULL, '/system/user/pwd', '修改密码', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (8, '2021-12-09 18:21:54', '2021-12-09 18:21:54', NULL, '/system/user/getById/:userId', '通过id获取用户信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (9, '2021-12-09 18:58:50', '2021-12-09 18:58:50', NULL, '/system/user/getInit', '获取初始化角色岗位信息(添加用户初始化)', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (10, '2021-12-09 18:59:43', '2021-12-09 18:59:43', NULL, '/system/user/getRoPo', '获取用户角色岗位信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (11, '2021-12-09 19:00:24', '2021-12-09 19:00:24', NULL, '/system/user', '添加用户信息', 'user', 'POST');
INSERT INTO "public"."sys_apis" VALUES (12, '2021-12-09 19:00:52', '2021-12-09 19:00:52', NULL, '/system/user', '修改用户信息', 'user', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (13, '2021-12-09 19:02:30', '2021-12-09 19:02:30', NULL, '/system/user/export', '导出用户信息', 'user', 'GET');
INSERT INTO "public"."sys_apis" VALUES (14, '2021-12-09 19:04:04', '2021-12-09 19:04:04', NULL, '/system/dept/roleDeptTreeSelect/:roleId', '获取角色部门树', 'dept', 'GET');
INSERT INTO "public"."sys_apis" VALUES (15, '2021-12-09 19:04:48', '2021-12-09 19:04:48', NULL, '/system/dept/deptTree', '获取所有部门树', 'dept', 'GET');
INSERT INTO "public"."sys_apis" VALUES (16, '2021-12-09 19:07:37', '2021-12-09 19:07:37', NULL, '/system/dept', '添加部门信息', 'dept', 'POST');
INSERT INTO "public"."sys_apis" VALUES (17, '2021-12-09 19:08:14', '2021-12-09 19:08:14', NULL, '/system/dept', '修改部门信息', 'dept', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (18, '2021-12-09 19:08:40', '2021-12-09 19:08:40', NULL, '/system/dept/:deptId', '删除部门信息', 'dept', 'DELETE');
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
INSERT INTO "public"."sys_apis" VALUES (79, '2021-12-24 15:45:03', '2021-12-24 15:46:36', NULL, '/log/logJob/list', '任务日志列表', 'log', 'GET');
INSERT INTO "public"."sys_apis" VALUES (80, '2021-12-24 15:45:33', '2021-12-24 15:46:43', NULL, '/log/logJob/all', '清空任务日志', 'log', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (81, '2021-12-24 15:46:08', '2021-12-24 16:33:13', NULL, '/log/logJob/:logId', '删除任务日志', 'log', 'DELETE');
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
INSERT INTO "public"."sys_apis" VALUES (99, '2022-01-13 16:44:44', '2022-01-13 16:45:27', NULL, '/resource/oss/list', '获取oss列表', 'oss', 'GET');
INSERT INTO "public"."sys_apis" VALUES (100, '2022-01-13 16:44:44', '2022-01-13 16:44:44', NULL, '/resource/oss/:ossId', '获取oss', 'oss', 'GET');
INSERT INTO "public"."sys_apis" VALUES (101, '2022-01-13 16:44:44', '2022-01-13 16:44:44', NULL, '/resource/oss', '添加oss', 'oss', 'POST');
INSERT INTO "public"."sys_apis" VALUES (102, '2022-01-13 16:44:44', '2022-01-13 16:44:44', NULL, '/resource/oss', '修改oss', 'oss', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (103, '2022-01-13 16:44:44', '2022-01-13 16:44:44', NULL, '/resource/oss/:ossId', '删除oss', 'oss', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (104, '2022-01-14 13:19:21', '2022-01-14 13:19:21', NULL, '/resource/oss/changeStatus', '修改状态', 'oss', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (105, '2022-01-14 13:20:14', '2022-01-14 13:20:14', NULL, '/resource/oss/uploadFile', '调试上传文件', 'oss', 'POST');
INSERT INTO "public"."sys_apis" VALUES (106, '2022-01-14 15:30:39', '2022-01-14 15:30:39', NULL, '/resource/email/list', '邮件分页列表', 'mail', 'GET');
INSERT INTO "public"."sys_apis" VALUES (107, '2022-01-14 15:31:20', '2022-01-14 15:31:20', NULL, '/resource/email/:mailId', '获取邮件', 'mail', 'GET');
INSERT INTO "public"."sys_apis" VALUES (108, '2022-01-14 15:31:54', '2022-01-14 15:31:54', NULL, '/resource/email', '添加邮件配置', 'mail', 'POST');
INSERT INTO "public"."sys_apis" VALUES (109, '2022-01-14 15:32:21', '2022-01-14 15:32:21', NULL, '/resource/email', '修改邮件配置', 'mail', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (110, '2022-01-14 15:32:53', '2022-01-14 15:32:53', NULL, '/resource/email/:mailId', '删除邮件', 'mail', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (111, '2022-01-14 17:11:42', '2022-01-14 17:11:42', NULL, '/resource/email/changeStatus', '修改状态', 'mail', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (112, '2022-01-14 17:12:17', '2022-01-14 17:12:17', NULL, '/resource/email/debugMail', '发送邮件调试', 'mail', 'POST');
INSERT INTO "public"."sys_apis" VALUES (113, '2022-07-15 18:06:27', '2022-07-15 18:06:27', NULL, '/system/tenant/list', '租户列表', 'tenant', 'GET');
INSERT INTO "public"."sys_apis" VALUES (114, '2022-07-15 18:07:16', '2022-07-15 18:07:16', NULL, '/system/tenant/:tenantId', '获取租户', 'tenant', 'GET');
INSERT INTO "public"."sys_apis" VALUES (115, '2022-07-15 18:07:43', '2022-07-15 18:07:43', NULL, '/system/tenant', '添加租户', 'tenant', 'POST');
INSERT INTO "public"."sys_apis" VALUES (116, '2022-07-15 18:08:08', '2022-07-15 18:08:08', NULL, '/system/tenant', '修改租户', 'tenant', 'PUT');
INSERT INTO "public"."sys_apis" VALUES (117, '2022-07-15 18:08:57', '2022-07-15 18:08:57', NULL, '/system/tenant/:tenantId', '删除租户', 'tenant', 'DELETE');
INSERT INTO "public"."sys_apis" VALUES (123, '2022-07-18 10:24:03', '2022-07-18 10:24:03', NULL, '/system/tenant/lists', '获取所有租户', 'tenant', 'GET');
-- ----------------------------
-- Table structure for sys_configs
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_configs";
CREATE SEQUENCE IF NOT EXISTS "sys_configs_config_id_seq";
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
CREATE SEQUENCE IF NOT EXISTS "sys_depts_dept_id_seq";
CREATE TABLE "public"."sys_depts" (
  "dept_id" int8 NOT NULL DEFAULT nextval('sys_depts_dept_id_seq'::regclass),
  "tenant_id" int8 NULL DEFAULT NULL,
  "parent_id" int8 NULL DEFAULT NULL,
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
INSERT INTO "public"."sys_depts" VALUES (2,1, 0, '/0/2', '熊猫科技', 0, 'xm', '18353366836', '342@qq.com', '0', 'admin', 'admin', '2021-12-01 17:31:53+00', '2021-12-02 08:56:19+00', NULL);
INSERT INTO "public"."sys_depts" VALUES (3,1, 2, '/0/2/3', '研发部', 1, 'panda', '18353366543', 'ewr@qq.com', '0', 'admin', 'admin', '2021-12-01 17:37:43+00', '2021-12-02 08:55:56+00', NULL);
INSERT INTO "public"."sys_depts" VALUES (7,1, 2, '/0/2/7', '营销部', 2, 'panda', '18353333333', '342@qq.com', '0', 'panda', 'panda', '2021-12-24 10:46:24+00', '2021-12-24 10:47:15+00', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dict_data";
CREATE SEQUENCE IF NOT EXISTS "sys_dict_data_dict_code_seq";
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
CREATE SEQUENCE IF NOT EXISTS "sys_dict_types_dict_id_seq";
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
CREATE SEQUENCE IF NOT EXISTS "sys_jobs_job_id_seq";
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
CREATE SEQUENCE IF NOT EXISTS "sys_menus_menu_id_seq";
CREATE TABLE "public"."public"."sys_menus" (
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
INSERT INTO "public"."sys_menus" VALUES (1, '系统设置', '', 0, 0, 'elementSetting', '/system', 'Layout', '1', '', 'M', '0', '0', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08', '2021-12-28 13:32:21', NULL);
INSERT INTO "public"."sys_menus" VALUES (3, '用户管理', '', 1, 1, 'elementUser', '/system/user', '/system/user/index', '1', '', 'C', '0', '1', '1', 'system:user:list', '0', 'admin', 'panda', '', '2021-12-02 14:07:56', '2021-12-28 13:32:44', NULL);
INSERT INTO "public"."sys_menus" VALUES (4, '添加用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:add', '0', 'admin', '', '', '2021-12-03 13:36:33', '2021-12-03 13:36:33', NULL);
INSERT INTO "public"."sys_menus" VALUES (5, '编辑用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:edit', '0', 'admin', '', '', '2021-12-03 13:48:13', '2021-12-03 13:48:13', NULL);
INSERT INTO "public"."sys_menus" VALUES (6, '角色管理', '', 1, 2, 'elementUserFilled', '/system/role', '/system/role/index', '1', '', 'C', '0', '1', '1', 'system:role:list', '0', '', 'panda', '', '2021-12-03 13:51:55', '2022-07-16 10:23:21', NULL);
INSERT INTO "public"."sys_menus" VALUES (7, '菜单管理', '', 1, 2, 'iconfont icon-juxingkaobei', '/system/menu', '/system/menu/index', '1', '', 'C', '0', '1', '1', 'system:menu:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:33:19', NULL);
INSERT INTO "public"."sys_menus" VALUES (8, '部门管理', '', 1, 3, 'iconfont icon-jiliandongxuanzeqi', '/system/dept', '/system/dept/index', '1', '', 'C', '0', '1', '1', 'system:dept:list', '0', 'admin', 'panda', '', '2021-12-03 13:58:36', '2021-12-28 13:40:20', NULL);
INSERT INTO "public"."sys_menus" VALUES (9, '岗位管理', '', 1, 4, 'iconfont icon-neiqianshujuchucun', '/system/post', '/system/post/index', '1', '', 'C', '0', '1', '1', 'system:post:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:40:31', NULL);
INSERT INTO "public"."sys_menus" VALUES (10, '字典管理', '', 1, 5, 'elementCellphone', '/system/dict', '/system/dict/index', '1', '', 'C', '0', '1', '1', 'system:dict:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:40:50', NULL);
INSERT INTO "public"."sys_menus" VALUES (11, '参数管理', '', 1, 6, 'elementDocumentCopy', '/system/config', '/system/config/index', '1', '', 'C', '0', '1', '1', 'system:config:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:41:05', NULL);
INSERT INTO "public"."sys_menus" VALUES (12, '个人中心', '', 0, 10, 'elementAvatar', '/personal', '/personal/index', '1', '', 'M', '0', '0', '0', '', '0', 'admin', 'panda', '', '2021-12-03 14:12:43', '2021-12-28 13:43:17', NULL);
INSERT INTO "public"."sys_menus" VALUES (13, '添加配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:add', '0', 'admin', '', '', '2021-12-06 17:19:19', '2021-12-06 17:19:19', NULL);
INSERT INTO "public"."sys_menus" VALUES (14, '修改配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:edit', '0', 'admin', '', '', '2021-12-06 17:20:30', '2021-12-06 17:20:30', NULL);
INSERT INTO "public"."sys_menus" VALUES (15, '删除配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:delete', '0', 'admin', '', '', '2021-12-06 17:23:52', '2021-12-06 17:23:52', NULL);
INSERT INTO "public"."sys_menus" VALUES (16, '导出配置', '', 11, 1, '', '', '', '', '', 'F', '', '', '', 'system:config:export', '0', 'admin', '', '', '2021-12-06 17:24:41', '2021-12-06 17:24:41', NULL);
INSERT INTO "public"."sys_menus" VALUES (17, '新增角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:add', '0', 'admin', '', '', '2021-12-06 17:43:35', '2021-12-06 17:43:35', NULL);
INSERT INTO "public"."sys_menus" VALUES (18, '删除角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:delete', '0', 'admin', '', '', '2021-12-06 17:44:10', '2021-12-06 17:44:10', NULL);
INSERT INTO "public"."sys_menus" VALUES (19, '修改角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:edit', '0', 'admin', '', '', '2021-12-06 17:44:48', '2021-12-06 17:44:48', NULL);
INSERT INTO "public"."sys_menus" VALUES (20, '导出角色', '', 6, 1, '', '', '', '', '', 'F', '', '', '', 'system:role:export', '0', 'admin', '', '', '2021-12-06 17:45:25', '2021-12-06 17:45:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (21, '添加菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:add', '0', 'admin', '', '', '2021-12-06 17:46:01', '2021-12-06 17:46:01', NULL);
INSERT INTO "public"."sys_menus" VALUES (22, '修改菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:edit', '0', 'admin', '', '', '2021-12-06 17:46:24', '2021-12-06 17:46:24', NULL);
INSERT INTO "public"."sys_menus" VALUES (23, '删除菜单', '', 7, 1, '', '', '', '', '', 'F', '', '', '', 'system:menu:delete', '0', 'admin', '', '', '2021-12-06 17:46:47', '2021-12-06 17:46:47', NULL);
INSERT INTO "public"."sys_menus" VALUES (24, '添加部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:add', '0', 'admin', '', '', '2021-12-07 09:33:58', '2021-12-07 09:33:58', NULL);
INSERT INTO "public"."sys_menus" VALUES (25, '编辑部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:edit', '0', 'admin', '', '', '2021-12-07 09:34:39', '2021-12-07 09:34:39', NULL);
INSERT INTO "public"."sys_menus" VALUES (26, '删除部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:delete', '0', 'admin', 'admin', '', '2021-12-07 09:35:09', '2021-12-07 09:36:26', NULL);
INSERT INTO "public"."sys_menus" VALUES (27, '导出部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:dept:export', '0', 'admin', '', '', '2021-12-07 09:35:51', '2021-12-07 09:35:51', '2021-12-07 09:36:37');
INSERT INTO "public"."sys_menus" VALUES (28, '添加岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:add', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (29, '编辑岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:edit', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (30, '删除岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:delete', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (31, '导出岗位', '', 9, 1, '', '', '', '', '', 'F', '', '', '', 'system:post:export', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (32, '添加字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:add', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (33, '编辑字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:edit', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (34, '删除字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:delete', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (35, '导出字典类型', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictT:export', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (36, '新增字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:add', '0', 'admin', '', '', '2021-12-07 09:35:09', '2021-12-07 09:35:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (37, '修改字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:edit', '0', 'admin', '', '', '2021-12-07 09:48:04', '2021-12-07 09:48:04', NULL);
INSERT INTO "public"."sys_menus" VALUES (38, '删除字典数据', '', 10, 1, '', '', '', '', '', 'F', '', '', '', 'system:dictD:delete', '0', 'admin', '', '', '2021-12-07 09:48:42', '2021-12-07 09:48:42', NULL);
INSERT INTO "public"."sys_menus" VALUES (39, 'API管理', '', 1, 2, 'iconfont icon-siweidaotu', '/system/api', '/system/api/index', '1', '', 'C', '0', '1', '1', 'system:api:list', '0', '', 'panda', '', '2021-12-09 09:09:13', '2022-07-16 10:23:42', NULL);
INSERT INTO "public"."sys_menus" VALUES (40, '添加api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:add', '0', 'admin', '', '', '2021-12-09 09:09:54', '2021-12-09 09:09:54', NULL);
INSERT INTO "public"."sys_menus" VALUES (41, '编辑api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:edit', '0', 'admin', '', '', '2021-12-09 09:10:38', '2021-12-09 09:10:38', NULL);
INSERT INTO "public"."sys_menus" VALUES (42, '删除api', '', 39, 1, '', '/system/api', '', '', '', 'F', '', '', '', 'system:api:delete', '0', 'admin', '', '', '2021-12-09 09:11:11', '2021-12-09 09:11:11', NULL);
INSERT INTO "public"."sys_menus" VALUES (43, '日志系统', '', 0, 1, 'iconfont icon-biaodan', '/log', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08', '2021-12-28 13:38:33', NULL);
INSERT INTO "public"."sys_menus" VALUES (44, '系统工具', '', 0, 2, 'iconfont icon-gongju', '/tool', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-16 16:35:15', '2021-12-28 13:38:46', NULL);
INSERT INTO "public"."sys_menus" VALUES (45, '操作日志', '', 43, 1, 'iconfont icon-bolangnengshiyanchang', '/log/operation', '/log/operation/index', '1', '', 'C', '0', '1', '1', 'log:operation:list', '0', 'admin', 'panda', '', '2021-12-16 16:42:03', '2021-12-28 13:39:44', NULL);
INSERT INTO "public"."sys_menus" VALUES (46, '登录日志', '', 43, 2, 'iconfont icon--chaifenlie', '/log/login', '/log/login/index', '1', '', 'C', '0', '1', '1', 'log:login:list', '0', 'admin', 'panda', '', '2021-12-16 16:43:28', '2021-12-28 13:39:58', NULL);
INSERT INTO "public"."sys_menus" VALUES (47, '服务监控', '', 44, 1, 'elementCpu', '/tool/monitor/', '/tool/monitor/index', '1', '', 'C', '0', '1', '1', 'tool:monitor:list', '0', 'admin', 'panda', '', '2021-12-03 14:12:43', '2021-12-28 13:41:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (48, '定时任务', '', 44, 2, 'elementAlarmClock', '/tool/job', '/tool/job/index', '1', '', 'C', '0', '1', '1', 'tool:job:list', '0', 'admin', 'panda', '', '2021-12-16 16:48:45', '2021-12-28 13:41:59', NULL);
INSERT INTO "public"."sys_menus" VALUES (49, '开发工具', '', 0, 3, 'iconfont icon-diannao', '/develop', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', '', '', '2021-12-16 16:53:11', '2021-12-16 16:53:11', NULL);
INSERT INTO "public"."sys_menus" VALUES (50, '表单构建', '', 49, 1, 'iconfont icon-zidingyibuju', '/develop/form', '/develop/form/index', '1', '', 'C', '0', '1', '1', 'develop:form:list', '0', 'admin', 'panda', '', '2021-12-16 16:55:01', '2022-07-12 18:56:18', NULL);
INSERT INTO "public"."sys_menus" VALUES (51, '代码生成', '', 49, 2, 'iconfont icon-zhongduancanshu', '/develop/code', '/develop/code/index', '1', '', 'C', '0', '1', '1', 'develop:code:list', '0', 'admin', '', '', '2021-12-16 16:56:48', '2021-12-16 16:56:48', NULL);
INSERT INTO "public"."sys_menus" VALUES (52, '系统接口', '', 49, 3, 'iconfont icon-wenducanshu-05', '/develop/apis', '/layout/routerView/iframes', '0', 'http://47.104.252.2:8080/swagger/index.html', 'C', '0', '1', '1', 'develop:apis:list', '0', '', 'panda', '', '2021-12-16 16:58:07', '2022-07-13 11:50:34', NULL);
INSERT INTO "public"."sys_menus" VALUES (53, '资源管理', '', 0, 4, 'iconfont icon-juxingkaobei', '/resource', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', '', '', '2021-12-16 17:02:06', '2021-12-16 17:02:06', NULL);
INSERT INTO "public"."sys_menus" VALUES (54, '对象存储', '', 53, 1, 'iconfont icon-chazhaobiaodanliebiao', '/resource/file', '/resource/file/index', '1', '', 'C', '0', '1', '1', 'resource:file:list', '0', 'admin', 'panda', '', '2021-12-16 17:06:04', '2022-01-13 17:30:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (55, '公告通知', '', 44, 3, 'elementTicket', '/tool/notice', '/tool/notice/index', '1', '', 'C', '0', '1', '1', 'tool:notice:list', '0', 'admin', 'panda', '', '2021-12-16 22:09:11', '2021-12-28 13:42:39', NULL);
INSERT INTO "public"."sys_menus" VALUES (56, '任务日志', '', 43, 1, 'iconfont icon--chaifenhang', '/log/job', '/log/job/index', '1', '', 'C', '0', '1', '1', 'log:job:list', '0', 'panda', 'panda', '', '2021-12-24 22:13:45', '2021-12-28 13:39:52', NULL);
INSERT INTO "public"."sys_menus" VALUES (57, '邮件管理', '', 53, 1, 'iconfont icon-shouye_dongtaihui', '/resource/mail', '/resource/mail/index', '1', '', 'C', '0', '1', '1', 'resource:mail:list', '0', 'panda', 'panda', '', '2021-12-28 15:12:37', '2021-12-28 15:12:45', NULL);
INSERT INTO "public"."sys_menus" VALUES (58, '短信管理', '', 53, 3, 'elementChatDotRound', '/resource/message', '/resource/message/index', '1', '', 'C', '0', '1', '1', 'resource:message:list', '0', 'panda', '', '', '2021-12-16 17:06:04', '2021-12-16 17:06:04', NULL);
INSERT INTO "public"."sys_menus" VALUES (59, '删除', '', 45, 1, '', '', '', '', '', 'F', '', '', '', 'log:operation:delete', '0', 'panda', '', '', '2022-01-14 13:28:25', '2022-01-14 13:28:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (60, '清空', '', 45, 1, '', '', '', '', '', 'F', '', '', '', 'log:operation:clean', '0', 'panda', '', '', '2022-01-14 13:29:24', '2022-01-14 13:29:24', NULL);
INSERT INTO "public"."sys_menus" VALUES (61, '删除', '', 56, 1, '', '', '', '', '', 'F', '', '', '', 'log:job:delete', '0', 'panda', '', '', '2022-01-14 13:29:57', '2022-01-14 13:29:57', NULL);
INSERT INTO "public"."sys_menus" VALUES (62, '清空', '', 56, 1, '', '', '', '', '', 'F', '', '', '', 'log:job:clean', '0', 'panda', '', '', '2022-01-14 13:30:15', '2022-01-14 13:30:15', NULL);
INSERT INTO "public"."sys_menus" VALUES (63, '删除', '', 46, 1, '', '', '', '', '', 'F', '', '', '', 'log:login:delete', '0', 'panda', '', '', '2022-01-14 13:30:46', '2022-01-14 13:30:46', NULL);
INSERT INTO "public"."sys_menus" VALUES (64, '清空', '', 46, 1, '', '', '', '', '', 'F', '', '', '', 'log:login:clean', '0', 'panda', '', '', '2022-01-14 13:31:06', '2022-01-14 13:31:06', NULL);
INSERT INTO "public"."sys_menus" VALUES (65, '新增', '', 48, 1, '', '', '', '', '', 'F', '', '', '', 'tool:job:add', '0', 'panda', '', '', '2022-01-14 13:32:48', '2022-01-14 13:32:48', NULL);
INSERT INTO "public"."sys_menus" VALUES (66, '编辑', '', 48, 1, '', '', '', '', '', 'F', '', '', '', 'tool:job:edit', '0', 'panda', '', '', '2022-01-14 13:33:17', '2022-01-14 13:33:17', NULL);
INSERT INTO "public"."sys_menus" VALUES (67, '删除', '', 48, 1, '', '', '', '', '', 'F', '', '', '', 'tool:job:delete', '0', 'panda', '', '', '2022-01-14 13:33:43', '2022-01-14 13:33:43', NULL);
INSERT INTO "public"."sys_menus" VALUES (68, '开关', '', 48, 1, '', '', '', '', '', 'F', '', '', '', 'tool:job:run', '0', 'panda', '', '', '2022-01-14 13:34:27', '2022-01-14 13:34:27', NULL);
INSERT INTO "public"."sys_menus" VALUES (69, '添加', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:add', '0', 'panda', '', '', '2022-01-14 13:35:23', '2022-01-14 13:35:23', NULL);
INSERT INTO "public"."sys_menus" VALUES (70, '编辑', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:edit', '0', 'panda', '', '', '2022-01-14 13:36:04', '2022-01-14 13:36:04', NULL);
INSERT INTO "public"."sys_menus" VALUES (71, '删除', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:delete', '0', 'panda', '', '', '2022-01-14 13:36:26', '2022-01-14 13:36:26', NULL);
INSERT INTO "public"."sys_menus" VALUES (72, '查看', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:view', '0', 'panda', '', '', '2022-01-14 13:36:51', '2022-01-14 13:36:51', NULL);
INSERT INTO "public"."sys_menus" VALUES (73, '导入', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:add', '0', 'panda', '', '', '2022-01-14 13:38:35', '2022-01-14 13:38:35', NULL);
INSERT INTO "public"."sys_menus" VALUES (74, '编辑', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:edit', '0', 'panda', '', '', '2022-01-14 13:41:25', '2022-01-14 13:41:25', NULL);
INSERT INTO "public"."sys_menus" VALUES (75, '删除', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:delete', '0', 'panda', '', '', '2022-01-14 13:41:42', '2022-01-14 13:41:42', NULL);
INSERT INTO "public"."sys_menus" VALUES (76, '预览', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:view', '0', 'panda', '', '', '2022-01-14 13:42:01', '2022-01-14 13:42:01', NULL);
INSERT INTO "public"."sys_menus" VALUES (77, '生成代码', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:code', '0', 'panda', '', '', '2022-01-14 13:42:48', '2022-01-14 13:42:48', NULL);
INSERT INTO "public"."sys_menus" VALUES (78, '添加', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:add', '0', 'panda', '', '', '2022-01-17 11:26:15', '2022-01-17 11:26:15', NULL);
INSERT INTO "public"."sys_menus" VALUES (79, '编辑', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:edit', '0', 'panda', '', '', '2022-01-17 11:26:39', '2022-01-17 11:26:39', NULL);
INSERT INTO "public"."sys_menus" VALUES (80, '删除', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:delete', '0', 'panda', '', '', '2022-01-17 11:26:56', '2022-01-17 11:26:56', NULL);
INSERT INTO "public"."sys_menus" VALUES (81, '调试', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:debug', '0', 'panda', '', '', '2022-01-17 11:27:14', '2022-01-17 11:27:14', NULL);
INSERT INTO "public"."sys_menus" VALUES (82, '调试', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:debug', '0', 'panda', '', '', '2022-01-17 11:27:17', '2022-01-17 11:27:17', '2022-01-17 11:27:25');
INSERT INTO "public"."sys_menus" VALUES (83, '添加', '', 57, 1, '', '', '', '', '', 'F', '', '', '', 'resource:mail:add', '0', 'panda', '', '', '2022-01-17 11:28:06', '2022-01-17 11:28:06', NULL);
INSERT INTO "public"."sys_menus" VALUES (84, '编辑', '', 57, 1, '', '', '', '', '', 'F', '', '', '', 'resource:mail:edit', '0', 'panda', '', '', '2022-01-17 11:28:37', '2022-01-17 11:28:37', NULL);
INSERT INTO "public"."sys_menus" VALUES (85, '删除', '', 57, 1, '', '', '', '', '', 'F', '', '', '', 'resource:mail:delete', '0', 'panda', '', '', '2022-01-17 11:29:09', '2022-01-17 11:29:09', NULL);
INSERT INTO "public"."sys_menus" VALUES (86, '调试', '', 57, 1, '', '', '', '', '', 'F', '', '', '', 'resource:mail:debug', '0', 'panda', '', '', '2022-01-17 11:29:46', '2022-01-17 11:29:46', NULL);
INSERT INTO "public"."sys_menus" VALUES (87, '租户管理', '', 1, 1, 'iconfont icon-quanxian', '/system/tenant', '/system/tenant/index', '1', '', 'C', '0', '1', '1', 'system:tenant:list', '0', 'panda', '', '', '2022-07-15 18:03:35', '2022-07-15 18:03:35', NULL);
INSERT INTO "public"."sys_menus" VALUES (88, '添加', '', 87, 1, '', '', '', '', '', 'F', '', '', '', 'system:tenant:add', '0', 'panda', '', '', '2022-07-15 18:28:58', '2022-07-15 18:28:58', NULL);
INSERT INTO "public"."sys_menus" VALUES (89, '编辑', '', 87, 1, '', '', '', '', '', 'F', '', '', '', 'system:tenant:edit', '0', 'panda', '', '', '2022-07-15 18:29:34', '2022-07-15 18:29:34', NULL);
INSERT INTO "public"."sys_menus" VALUES (90, '删除', '', 87, 1, '', '', '', '', '', 'F', '', '', '', 'system:tenant:delete', '0', 'panda', '', '', '2022-07-15 18:30:00', '2022-07-15 18:30:00', NULL);

-- ----------------------------
-- Table structure for sys_notices
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_notices";
CREATE SEQUENCE IF NOT EXISTS "sys_notices_notice_id_seq";
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
CREATE SEQUENCE IF NOT EXISTS "sys_posts_post_id_seq";
CREATE TABLE "public"."sys_posts" (
  "post_id" int8 NOT NULL DEFAULT nextval('sys_posts_post_id_seq'::regclass),
  "tenant_id" int8 NULL DEFAULT NULL,
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
INSERT INTO "public"."sys_posts" VALUES (1, 1,'首席执行官', 'CEO', 0, '0', '首席执行官', 'admin', '', '2021-12-02 09:21:44+00', '2021-12-02 09:24:25+00', NULL);
INSERT INTO "public"."sys_posts" VALUES (3, 1,'首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44+00', '2021-12-02 09:25:59+00', '2021-12-02 09:27:41+00');
INSERT INTO "public"."sys_posts" VALUES (4, 1,'首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44+00', '2021-12-02 09:25:59+00', NULL);
INSERT INTO "public"."sys_posts" VALUES (5, 1,'123', '123', 0, '0', '', 'admin', '', '2021-12-18 00:33:28+00', '2021-12-18 00:33:28+00', '2021-12-28 14:11:52+00');

-- ----------------------------
-- Table structure for sys_role_depts
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role_depts";
CREATE SEQUENCE IF NOT EXISTS "sys_role_depts_id_seq";
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
CREATE SEQUENCE IF NOT EXISTS "sys_role_menus_id_seq";
CREATE TABLE "public"."sys_role_menus" (
  "id" int8 NOT NULL DEFAULT nextval('sys_role_menus_id_seq'::regclass)
  "role_id" int8,
  "menu_id" int8,
  "role_name" varchar(128) COLLATE "pg_catalog"."default",

)
;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO "public"."sys_role_menus" VALUES (2870, 1, 1, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2871, 1, 3, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2872, 1, 4, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2873, 1, 5, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2874, 1, 6, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2875, 1, 7, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2876, 1, 8, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2877, 1, 9, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2878, 1, 10, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2879, 1, 11, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2880, 1, 12, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2881, 1, 13, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2882, 1, 14, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2883, 1, 15, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2884, 1, 16, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2885, 1, 17, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2886, 1, 18, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2887, 1, 19, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2888, 1, 20, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2889, 1, 21, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2890, 1, 22, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2891, 1, 23, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2892, 1, 24, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2893, 1, 25, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2894, 1, 26, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2895, 1, 28, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2896, 1, 29, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2897, 1, 30, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2898, 1, 31, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2899, 1, 32, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2900, 1, 33, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2901, 1, 34, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2902, 1, 35, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2903, 1, 36, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2904, 1, 37, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2905, 1, 38, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2906, 1, 39, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2907, 1, 40, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2908, 1, 41, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2909, 1, 42, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2910, 1, 43, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2911, 1, 44, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2912, 1, 45, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2913, 1, 46, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2914, 1, 47, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2915, 1, 48, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2916, 1, 49, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2917, 1, 50, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2918, 1, 51, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2919, 1, 52, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2920, 1, 53, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2921, 1, 54, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2922, 1, 55, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2923, 1, 56, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2924, 1, 57, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2925, 1, 58, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2926, 1, 59, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2927, 1, 60, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2928, 1, 61, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2929, 1, 62, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2930, 1, 63, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2931, 1, 64, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2932, 1, 65, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2933, 1, 66, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2934, 1, 67, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2935, 1, 68, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2936, 1, 69, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2937, 1, 70, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2938, 1, 71, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2939, 1, 72, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2940, 1, 73, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2941, 1, 74, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2942, 1, 75, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2943, 1, 76, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2944, 1, 77, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2945, 1, 78, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2946, 1, 79, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2947, 1, 80, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2948, 1, 81, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2949, 1, 83, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2950, 1, 84, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2951, 1, 85, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2952, 1, 86, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2953, 1, 87, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2954, 1, 88, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2955, 1, 89, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2956, 1, 90, 'admin');
INSERT INTO "public"."sys_role_menus" VALUES (2957, 2, 1, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2958, 2, 3, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2959, 2, 4, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2960, 2, 5, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2961, 2, 6, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2962, 2, 7, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2963, 2, 8, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2964, 2, 9, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2965, 2, 10, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2966, 2, 11, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2967, 2, 12, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2968, 2, 13, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2969, 2, 14, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2970, 2, 15, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2971, 2, 16, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2972, 2, 17, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2973, 2, 18, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2974, 2, 19, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2975, 2, 20, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2976, 2, 21, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2977, 2, 22, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2978, 2, 23, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2979, 2, 25, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2980, 2, 26, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2981, 2, 28, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2982, 2, 29, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2983, 2, 30, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2984, 2, 31, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2985, 2, 32, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2986, 2, 33, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2987, 2, 34, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2988, 2, 35, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2989, 2, 36, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2990, 2, 37, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2991, 2, 38, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2992, 2, 39, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2993, 2, 40, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2994, 2, 41, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2995, 2, 42, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2996, 2, 43, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2997, 2, 44, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2998, 2, 45, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (2999, 2, 46, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3000, 2, 47, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3001, 2, 48, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3002, 2, 49, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3003, 2, 50, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3004, 2, 51, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3005, 2, 52, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3006, 2, 53, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3007, 2, 54, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3008, 2, 55, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3009, 2, 56, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3010, 2, 57, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3011, 2, 58, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3012, 2, 59, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3013, 2, 60, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3014, 2, 61, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3015, 2, 62, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3016, 2, 63, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3017, 2, 64, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3018, 2, 65, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3019, 2, 66, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3020, 2, 67, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3021, 2, 68, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3022, 2, 69, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3023, 2, 70, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3024, 2, 71, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3025, 2, 72, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3026, 2, 73, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3027, 2, 74, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3028, 2, 75, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3029, 2, 76, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3030, 2, 77, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3031, 2, 78, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3032, 2, 79, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3033, 2, 80, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3034, 2, 81, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3035, 2, 83, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3036, 2, 84, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3037, 2, 85, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3038, 2, 86, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3039, 2, 87, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3040, 2, 88, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3041, 2, 89, 'manage');
INSERT INTO "public"."sys_role_menus" VALUES (3042, 2, 90, 'manage');

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_roles";
CREATE SEQUENCE IF NOT EXISTS "sys_roles_role_id_seq";
CREATE TABLE "public"."sys_roles" (
  "role_id" int8 NOT NULL DEFAULT nextval('sys_roles_role_id_seq'::regclass),
  "role_name" varchar(128) COLLATE "pg_catalog"."default",
  "tenant_id" int8 NULL DEFAULT NULL,
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
INSERT INTO "public"."sys_roles" VALUES (1, '超管理员',1, '0', 'admin', 2, '1', '', 'admin', 'panda', '', '2021-12-02 16:03:26+00', '2021-12-28 15:16:11+00', NULL);
INSERT INTO "public"."sys_roles" VALUES (2, '管理员',1, '0', 'manage', 2, '', '', 'panda', 'panda', '', '2021-12-19 16:06:20+00', '2021-12-28 15:16:23+00', NULL);


-- ----------------------------
-- Table structure for sys_tenants
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_tenants";
CREATE TABLE "public"."sys_tenants" (
  "id" int8 NOT NULL DEFAULT nextval('sys_tenants_id_seq'::regclass),
  "create_time" timestamptz(6),
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6),
  "tenant_name" varchar(255) COLLATE "pg_catalog"."default",
  "remark" varchar(255) COLLATE "pg_catalog"."default",
  "expire_time" timestamptz(6)
)
;
COMMENT ON COLUMN "public"."sys_tenants"."tenant_name" IS '租户名';
COMMENT ON COLUMN "public"."sys_tenants"."remark" IS '备注';
COMMENT ON COLUMN "public"."sys_tenants"."expire_time" IS '过期时间';

-- ----------------------------
-- Records of sys_tenants
-- ----------------------------
INSERT INTO "public"."sys_tenants" VALUES (1, '2022-07-16 18:28:33+00', '2022-07-16 18:28:33+00', NULL, '熊猫科技', '鹅鹅鹅', '2099-07-16 00:00:00+00');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_users";
CREATE SEQUENCE IF NOT EXISTS "sys_users_user_id_seq";
CREATE TABLE "public"."sys_users" (
  "user_id" int8 NOT NULL DEFAULT nextval('sys_users_user_id_seq'::regclass),
  "nick_name" varchar(128) COLLATE "pg_catalog"."default",
  "phone" varchar(11) COLLATE "pg_catalog"."default",
  "tenant_id" int8 NULL DEFAULT NULL,
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
INSERT INTO "public"."sys_users" VALUES (1, 'pandax', '13818888888',1, 1, NULL, '', '0', '1@qq.com', 2, 1, '1', '1', 'admin', '1', NULL, '0', '2021-12-03 09:46:55+00', '2021-12-03 10:51:34+00', NULL, 'panda', '$2a$10$EXMJ5huCCTlCmP2ITFkAJ.4Mgmq3JcZGUvtE.KLX8j7FmhiiTEEya');
INSERT INTO "public"."sys_users" VALUES (4, 'panda', '18353366912',1, 2, '', '', '0', '2417920382@qq.com', 2, 4, '2', '4,1', 'panda', '', '', '0', '2021-12-19 15:58:09+00', '2021-12-19 16:06:54+00', NULL, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2');

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

-- ----------------------------
-- Primary Key structure for table dev_gen_table_columns
-- ----------------------------
ALTER TABLE "public"."dev_gen_table_columns" ADD CONSTRAINT "dev_gen_table_columns_pkey" PRIMARY KEY ("column_id");

-- ----------------------------
-- Primary Key structure for table dev_gen_tables
-- ----------------------------
ALTER TABLE "public"."dev_gen_tables" ADD CONSTRAINT "dev_gen_tables_pkey" PRIMARY KEY ("table_id");

-- ----------------------------
-- Primary Key structure for table sys_tenants
-- ----------------------------
ALTER TABLE "public"."sys_tenants" ADD CONSTRAINT "sys_tenants_pkey" PRIMARY KEY ("id");
