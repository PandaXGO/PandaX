/*
 Navicat Premium Data Transfer

 Source Server         : pandax
 Source Server Type    : MySQL
 Source Server Version : 50562

 Target Server Type    : MySQL
 Target Server Version : 50562
 File Encoding         : 65001

 Date: 23/07/2022 10:27:43
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
) ENGINE = InnoDB AUTO_INCREMENT = 2315 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/gen/code/:tableId', 'GET', '', '', 2190);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/gen/configure/:tableId', 'GET', '', '', 2191);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/gen/preview/:tableId', 'GET', '', '', 2189);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/table', 'POST', '', '', 2186);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/table', 'PUT', '', '', 2187);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/table/:tableId', 'DELETE', '', '', 2188);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/table/db/list', 'GET', '', '', 2181);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/table/info/:tableId', 'GET', '', '', 2183);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/table/info/tableName', 'GET', '', '', 2184);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/table/list', 'GET', '', '', 2182);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/develop/code/table/tableTree', 'GET', '', '', 2185);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/job', 'POST', '', '', 2193);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/job', 'PUT', '', '', 2194);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/job/:jobId', 'DELETE', '', '', 2196);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/job/:jobId', 'GET', '', '', 2195);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/job/changeStatus', 'PUT', '', '', 2199);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/job/list', 'GET', '', '', 2192);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/job/start/:jobId', 'GET', '', '', 2198);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/job/stop/:jobId', 'GET', '', '', 2197);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logJob/:logId', 'DELETE', '', '', 2208);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logJob/all', 'DELETE', '', '', 2207);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logJob/list', 'GET', '', '', 2206);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logLogin/:infoId', 'DELETE', '', '', 2201);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logLogin/all', 'DELETE', '', '', 2202);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logLogin/list', 'GET', '', '', 2200);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logOper/:operId', 'DELETE', '', '', 2204);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logOper/all', 'DELETE', '', '', 2205);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/log/logOper/list', 'GET', '', '', 2203);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/email', 'POST', '', '', 2211);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/email', 'PUT', '', '', 2212);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/email/:mailId', 'DELETE', '', '', 2213);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/email/:mailId', 'GET', '', '', 2210);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/email/changeStatus', 'PUT', '', '', 2214);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/email/debugMail', 'POST', '', '', 2215);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/email/list', 'GET', '', '', 2209);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/oss', 'POST', '', '', 2231);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/oss', 'PUT', '', '', 2232);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/oss/:ossId', 'DELETE', '', '', 2233);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/oss/:ossId', 'GET', '', '', 2230);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/oss/changeStatus', 'PUT', '', '', 2234);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/oss/list', 'GET', '', '', 2229);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/resource/oss/uploadFile', 'POST', '', '', 2235);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/api', 'POST', '', '', 2153);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/api', 'PUT', '', '', 2154);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/api/:id', 'DELETE', '', '', 2155);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/api/:id', 'GET', '', '', 2152);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/api/all', 'GET', '', '', 2150);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/api/getPolicyPathByRoleId', 'GET', '', '', 2151);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/api/list', 'GET', '', '', 2149);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/config', 'POST', '', '', 2159);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/config', 'PUT', '', '', 2160);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/config/:configId', 'DELETE', '', '', 2161);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/config/:configId', 'GET', '', '', 2158);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/config/configKey', 'GET', '', '', 2157);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/config/list', 'GET', '', '', 2156);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dept', 'POST', '', '', 2166);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dept', 'PUT', '', '', 2167);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dept/:deptId', 'DELETE', '', '', 2168);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dept/:deptId', 'GET', '', '', 2163);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dept/deptTree', 'GET', '', '', 2165);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dept/list', 'GET', '', '', 2162);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', 2164);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/data', 'POST', '', '', 2178);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/data', 'PUT', '', '', 2179);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/data/:dictCode', 'DELETE', '', '', 2180);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/data/:dictCode', 'GET', '', '', 2177);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/data/list', 'GET', '', '', 2175);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/data/type', 'GET', '', '', 2176);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/type', 'POST', '', '', 2171);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/type', 'PUT', '', '', 2172);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/type/:dictId', 'DELETE', '', '', 2173);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/type/:dictId', 'GET', '', '', 2170);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/type/export', 'GET', '', '', 2174);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/dict/type/list', 'GET', '', '', 2169);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu', 'POST', '', '', 2222);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu', 'PUT', '', '', 2223);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu/:menuId', 'DELETE', '', '', 2224);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu/:menuId', 'GET', '', '', 2221);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu/list', 'GET', '', '', 2220);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu/menuPaths', 'GET', '', '', 2219);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu/menuRole', 'GET', '', '', 2217);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu/menuTreeSelect', 'GET', '', '', 2216);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', 2218);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/notice', 'POST', '', '', 2226);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/notice', 'PUT', '', '', 2227);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/notice/:noticeId', 'DELETE', '', '', 2228);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/notice/list', 'GET', '', '', 2225);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/post', 'POST', '', '', 2238);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/post', 'PUT', '', '', 2239);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/post/:postId', 'DELETE', '', '', 2240);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/post/:postId', 'GET', '', '', 2237);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/post/list', 'GET', '', '', 2236);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/role', 'POST', '', '', 2243);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/role', 'PUT', '', '', 2244);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/role/:roleId', 'DELETE', '', '', 2245);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/role/:roleId', 'GET', '', '', 2242);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/role/changeStatus', 'PUT', '', '', 2246);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/role/dataScope', 'PUT', '', '', 2247);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/role/export', 'GET', '', '', 2248);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/role/list', 'GET', '', '', 2241);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/tenant', 'POST', '', '', 2251);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/tenant', 'PUT', '', '', 2252);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/tenant/:tenantId', 'DELETE', '', '', 2253);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/tenant/:tenantId', 'GET', '', '', 2250);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/tenant/list', 'GET', '', '', 2249);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/tenant/lists', 'GET', '', '', 2254);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user', 'POST', '', '', 2263);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user', 'PUT', '', '', 2264);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/:userId', 'DELETE', '', '', 2257);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/avatar', 'POST', '', '', 2258);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/changeStatus', 'PUT', '', '', 2256);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/export', 'GET', '', '', 2265);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/getById/:userId', 'GET', '', '', 2260);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/getInit', 'GET', '', '', 2261);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/getRoPo', 'GET', '', '', 2262);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/list', 'GET', '', '', 2255);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'admin', '/system/user/pwd', 'PUT', '', '', 2259);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/develop/code/gen/preview/:tableId', 'GET', '', '', 2287);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/develop/code/table/db/list', 'GET', '', '', 2282);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/develop/code/table/info/:tableId', 'GET', '', '', 2284);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/develop/code/table/info/tableName', 'GET', '', '', 2285);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/develop/code/table/list', 'GET', '', '', 2283);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/develop/code/table/tableTree', 'GET', '', '', 2286);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/job/:jobId', 'GET', '', '', 2289);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/job/list', 'GET', '', '', 2288);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/log/logJob/list', 'GET', '', '', 2292);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/log/logLogin/list', 'GET', '', '', 2290);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/log/logOper/list', 'GET', '', '', 2291);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/resource/email/:mailId', 'GET', '', '', 2294);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/resource/email/list', 'GET', '', '', 2293);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/resource/oss/:ossId', 'GET', '', '', 2303);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/resource/oss/list', 'GET', '', '', 2302);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/api/:id', 'GET', '', '', 2269);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/api/all', 'GET', '', '', 2267);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/api/getPolicyPathByRoleId', 'GET', '', '', 2268);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/api/list', 'GET', '', '', 2266);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/config/:configId', 'GET', '', '', 2272);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/config/configKey', 'GET', '', '', 2271);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/config/list', 'GET', '', '', 2270);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dept/:deptId', 'GET', '', '', 2274);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dept/deptTree', 'GET', '', '', 2276);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dept/list', 'GET', '', '', 2273);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dept/roleDeptTreeSelect/:roleId', 'GET', '', '', 2275);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dict/data/:dictCode', 'GET', '', '', 2281);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dict/data/list', 'GET', '', '', 2279);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dict/data/type', 'GET', '', '', 2280);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dict/type/:dictId', 'GET', '', '', 2278);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/dict/type/list', 'GET', '', '', 2277);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/menu/:menuId', 'GET', '', '', 2300);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/menu/list', 'GET', '', '', 2299);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/menu/menuPaths', 'GET', '', '', 2298);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/menu/menuRole', 'GET', '', '', 2296);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/menu/menuTreeSelect', 'GET', '', '', 2295);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', 2297);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/notice/list', 'GET', '', '', 2301);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/post/:postId', 'GET', '', '', 2305);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/post/list', 'GET', '', '', 2304);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/role/:roleId', 'GET', '', '', 2307);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/role/list', 'GET', '', '', 2306);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/tenant/:tenantId', 'GET', '', '', 2309);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/tenant/list', 'GET', '', '', 2308);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/tenant/lists', 'GET', '', '', 2310);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/user/getById/:userId', 'GET', '', '', 2312);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/user/getInit', 'GET', '', '', 2313);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/user/getRoPo', 'GET', '', '', 2314);
INSERT INTO `casbin_rule` VALUES ('p', '1', 'manage', '/system/user/list', 'GET', '', '', 2311);

-- ----------------------------
-- Table structure for dev_gen_table_columns
-- ----------------------------
DROP TABLE IF EXISTS `dev_gen_table_columns`;
CREATE TABLE `dev_gen_table_columns`  (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `table_id` bigint(20) NULL DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `column_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `column_comment` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `column_type` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `column_key` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `go_type` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `go_field` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `json_field` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `html_field` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_pk` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_increment` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_required` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_insert` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_edit` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_list` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_query` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `query_type` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `html_type` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `dict_type` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` bigint(20) NULL DEFAULT NULL,
  `link_table_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_table_class` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_table_package` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_label_id` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_label_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 89 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of dev_gen_table_columns
-- ----------------------------
INSERT INTO `dev_gen_table_columns` VALUES (1, 1, '', 'log_id', '', 'bigint(20)', '', 'int64', 'LogId', 'logId', '', '1', '', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (2, 1, '', 'name', '任务名称', 'varchar(128)', '', 'string', 'Name', 'name', '', '0', '', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (3, 1, '', 'job_group', '分组', 'varchar(128)', '', 'string', 'JobGroup', 'jobGroup', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (4, 1, '', 'entry_id', '任务id', 'int(11)', '', 'int', 'EntryId', 'entryId', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (5, 1, '', 'invoke_target', '调用方法', 'varchar(128)', '', 'string', 'InvokeTarget', 'invokeTarget', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (6, 1, '', 'log_info', '日志信息', 'varchar(255)', '', 'string', 'LogInfo', 'logInfo', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (7, 1, '', 'status', '状态', 'varchar(1)', '', 'string', 'Status', 'status', '', '0', '', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 7, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (8, 1, '', 'create_time', '', 'datetime', '', 'Time', 'CreateTime', 'createTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 8, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (9, 1, '', 'update_time', '', 'datetime', '', 'Time', 'UpdateTime', 'updateTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (10, 1, '', 'delete_time', '', 'datetime', '', 'Time', 'DeleteTime', 'deleteTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 10, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (27, 3, '', 'delete_time', '', 'datetime', '', 'Time', 'DeleteTime', 'deleteTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 13, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (28, 4, '', 'config_value', 'ConfigValue', 'varchar(255)', '', 'string', 'ConfigValue', 'configValue', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (29, 3, '', 'business_type', '0其它 1新增 2修改 3删除', 'varchar(1)', '', 'string', 'BusinessType', 'businessType', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'select', '', 3, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (30, 4, '', 'config_type', '是否系统内置0，1', 'varchar(64)', '', 'string', 'ConfigType', 'configType', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'select', '', 5, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (31, 3, '', 'method', '请求方法', 'varchar(255)', '', 'string', 'Method', 'method', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (32, 4, '', 'is_frontend', '是否前台', 'varchar(1)', '', 'string', 'IsFrontend', 'isFrontend', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (33, 3, '', 'oper_name', '操作人员', 'varchar(255)', '', 'string', 'OperName', 'operName', '', '0', '', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (34, 4, '', 'remark', 'Remark', 'varchar(128)', '', 'string', 'Remark', 'remark', '', '0', '', '0', '1', '1', '1', '0', 'EQ', 'input', '', 7, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (35, 3, '', 'oper_url', '操作url', 'varchar(255)', '', 'string', 'OperUrl', 'operUrl', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (36, 4, '', 'create_time', '', 'datetime', '', 'Time', 'CreateTime', 'createTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 8, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (37, 3, '', 'oper_ip', '操作IP', 'varchar(255)', '', 'string', 'OperIp', 'operIp', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (38, 4, '', 'update_time', '', 'datetime', '', 'Time', 'UpdateTime', 'updateTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (39, 3, '', 'create_time', '', 'datetime', '', 'Time', 'CreateTime', 'createTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 11, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (40, 3, '', 'update_time', '', 'datetime', '', 'Time', 'UpdateTime', 'updateTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 12, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (41, 3, '', 'status', '0=正常,1=异常', 'varchar(1)', '', 'string', 'Status', 'status', '', '0', '', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 10, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (42, 3, '', 'oper_id', '', 'bigint(20)', '', 'int64', 'OperId', 'operId', '', '1', '', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (43, 5, '', 'delete_time', '', 'datetime', '', 'Time', 'DeleteTime', 'deleteTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 14, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (44, 5, '', 'category', '', 'varchar(191)', '', 'string', 'Category', 'category', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (45, 5, '', 'app_id', 'AppId', 'varchar(128)', '', 'string', 'AppId', 'appId', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (46, 5, '', 'access_key', 'accessKey', 'varchar(128)', '', 'string', 'AccessKey', 'accessKey', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (47, 5, '', 'secret_key', 'secretKey', 'varchar(128)', '', 'string', 'SecretKey', 'secretKey', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (48, 5, '', 'bucket_name', 'bucketName', 'varchar(128)', '', 'string', 'BucketName', 'bucketName', '', '0', '', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 6, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (49, 5, '', 'endpoint', 'endpoint', 'varchar(128)', '', 'string', 'Endpoint', 'endpoint', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (50, 5, '', 'oss_code', 'ossCode', 'varchar(128)', '', 'string', 'OssCode', 'ossCode', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (51, 5, '', 'region', '地址', 'varchar(128)', '', 'string', 'Region', 'region', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (52, 5, '', 'remark', '说明', 'varchar(128)', '', 'string', 'Remark', 'remark', '', '0', '', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (53, 5, '', 'update_time', '', 'datetime', '', 'Time', 'UpdateTime', 'updateTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 13, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (54, 5, '', 'status', '状态', 'varchar(1)', '', 'string', 'Status', 'status', '', '0', '', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 11, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (55, 5, '', 'create_time', '', 'datetime', '', 'Time', 'CreateTime', 'createTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 12, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (56, 5, '', 'oss_id', '主键编码', 'bigint(20)', '', 'int64', 'OssId', 'ossId', '', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (57, 6, '', 'app_id', 'AppId', 'varchar(128)', '', 'string', 'AppId', 'appId', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (58, 6, '', 'category', '', 'varchar(191)', '', 'string', 'Category', 'category', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (59, 6, '', 'endpoint', 'endpoint', 'varchar(128)', '', 'string', 'Endpoint', 'endpoint', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (60, 6, '', 'access_key', 'accessKey', 'varchar(128)', '', 'string', 'AccessKey', 'accessKey', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (61, 6, '', 'secret_key', 'secretKey', 'varchar(128)', '', 'string', 'SecretKey', 'secretKey', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (62, 6, '', 'bucket_name', 'bucketName', 'varchar(128)', '', 'string', 'BucketName', 'bucketName', '', '0', '', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 6, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (63, 6, '', 'remark', '说明', 'varchar(128)', '', 'string', 'Remark', 'remark', '', '0', '', '0', '1', '1', '1', '0', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (64, 6, '', 'oss_code', 'ossCode', 'varchar(128)', '', 'string', 'OssCode', 'ossCode', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (65, 6, '', 'region', '地址', 'varchar(128)', '', 'string', 'Region', 'region', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (66, 6, '', 'create_time', '', 'datetime', '', 'Time', 'CreateTime', 'createTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 12, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (67, 6, '', 'delete_time', '', 'datetime', '', 'Time', 'DeleteTime', 'deleteTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 14, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (68, 6, '', 'status', '状态', 'varchar(1)', '', 'string', 'Status', 'status', '', '0', '', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 11, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (69, 6, '', 'update_time', '', 'datetime', '', 'Time', 'UpdateTime', 'updateTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 13, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (70, 6, '', 'oss_id', '主键编码', 'bigint(20)', '', 'int64', 'OssId', 'ossId', '', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (71, 7, '', 'delete_time', '', 'datetime', '', 'Time', 'DeleteTime', 'deleteTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 10, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (72, 7, '', 'host', '主机', 'varchar(191)', '', 'string', 'Host', 'host', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (73, 7, '', 'secret', '密码', 'varchar(191)', '', 'string', 'Secret', 'secret', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (74, 7, '', 'update_time', '', 'datetime', '', 'Time', 'UpdateTime', 'updateTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (75, 7, '', 'from', '发件人', 'varchar(191)', '', 'string', 'From', 'from', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (76, 7, '', 'is_ssl', '开启ssl', 'tinyint(1)', '', 'int', 'IsSsl', 'isSsl', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (77, 7, '', 'mail_id', '主键编码', 'bigint(20)', 'Json', 'int64', 'MailId', 'mailId', '', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (82, 9, '', 'expire_time', '过期时间', 'datetime', '', 'Time', 'ExpireTime', 'expireTime', '', '0', '', '0', '1', '1', '1', '1', 'EQ', 'datetime', '', 7, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (83, 9, '', 'delete_time', 'DeleteTime', 'datetime', '', 'Time', 'DeleteTime', 'deleteTime', '', '0', '', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 3, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (84, 9, '', 'create_time', 'CreateTime', 'datetime', '', 'Time', 'CreateTime', 'createTime', '', '0', '', '0', '0', '0', '1', '0', 'EQ', 'datetime', '', 1, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (85, 9, '', 'update_time', 'UpdateTime', 'datetime', '', 'Time', 'UpdateTime', 'updateTime', '', '0', '', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 2, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (86, 9, '', 'tenant_name', '租户名', 'varchar(255)', '', 'string', 'TenantName', 'tenantName', '', '0', '', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 5, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (87, 9, '', 'remark', '备注', 'varchar(255)', '', 'string', 'Remark', 'remark', '', '0', '', '0', '1', '1', '1', '0', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `dev_gen_table_columns` VALUES (88, 9, '', 'id', 'Id', 'bigint(20)', '', 'int64', 'Id', 'id', '', '1', '1', '0', '0', '0', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');

-- ----------------------------
-- Table structure for dev_gen_tables
-- ----------------------------
DROP TABLE IF EXISTS `dev_gen_tables`;
CREATE TABLE `dev_gen_tables`  (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `table_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `table_comment` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `class_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tpl_category` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `package_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `module_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `business_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `function_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `function_author` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `options` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `remark` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pk_column` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pk_go_field` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pk_json_field` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of dev_gen_tables
-- ----------------------------
INSERT INTO `dev_gen_tables` VALUES (1, 'log_jobs', 'LogJobs', 'LogJobs', 'crud', 'admin', 'log-jobs', 'logJobs', 'LogJobs', 'panda', '', '', 'log_id', 'LogId', 'logId', '2022-01-06 14:44:14', '2022-01-06 14:44:14', '2022-01-10 11:58:43');
INSERT INTO `dev_gen_tables` VALUES (2, 'log_logins', 'LogLogins', 'LogLogins', 'crud', 'log', 'log-logins', 'logLogins', 'logLogins', 'pandax', '', '', 'info_id', 'InfoId', 'infoId', '2022-01-06 14:44:27', '2022-01-26 14:45:46', '2022-07-16 18:24:04');
INSERT INTO `dev_gen_tables` VALUES (3, 'log_opers', 'LogOpers', 'LogOpers', 'crud', 'admin', 'log-opers', 'logOpers', 'LogOpers', 'panda', '', '', 'oper_id', 'OperId', 'operId', '2022-01-06 15:02:45', '2022-01-06 15:02:45', NULL);
INSERT INTO `dev_gen_tables` VALUES (4, 'sys_configs', 'SysConfigs', 'SysConfigs', 'crud', 'admin', 'sys-configs', 'sysConfigs', 'SysConfigs', 'panda', '', '', 'config_id', 'ConfigId', 'configId', '2022-01-06 15:02:45', '2022-01-06 15:02:45', '2022-01-12 15:56:49');
INSERT INTO `dev_gen_tables` VALUES (5, 'res_osses', 'ResOsses', 'ResOsses', 'crud', 'resource', 'res-osses', 'oss', 'ResOsses', 'panda', '', '', 'oss_id', 'OssId', 'ossId', '2022-01-13 15:12:14', '2022-01-13 15:39:11', NULL);
INSERT INTO `dev_gen_tables` VALUES (6, 'res_osses', 'ResOsses', 'ResOsses', 'crud', 'admin', 'res-osses', 'resOsses', 'ResOsses', 'panda', '', '', 'oss_id', 'OssId', 'ossId', '2022-01-13 15:12:27', '2022-01-13 15:12:27', '2022-01-13 15:12:41');
INSERT INTO `dev_gen_tables` VALUES (7, 'res_emails', 'ResEmails', 'ResEmails', 'crud', 'resource', 'res-emails', 'email', 'ResEmails', 'panda', '', '', 'mail_id', 'MailId', 'mailId', '2022-01-14 15:20:27', '2022-01-26 15:56:37', NULL);
INSERT INTO `dev_gen_tables` VALUES (8, 'sys_tenants', '租户', 'SysTenants', 'crud', 'system', 'sys-tenants', 'tenant', 'Tenant', 'panda', '', '', 'tenant_id', 'TenantId', 'tenantId', '2022-07-14 16:52:28', '2022-07-16 16:09:50', '2022-07-16 18:09:40');
INSERT INTO `dev_gen_tables` VALUES (9, 'sys_tenants', '租户', 'SysTenants', 'crud', 'system', 'sys-tenants', 'tenants', 'Tenants', 'panda', '', '', 'id', 'Id', 'id', '2022-07-16 18:14:53', '2022-07-16 22:45:33', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of log_jobs
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 3481 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_logins
-- ----------------------------
INSERT INTO `log_logins` VALUES (3402, 'admin', '0', '127.0.0.1', '内部IP', 'Chrome 92.0.4515.107', 'Windows 10', 'Windows', '2022-07-18 11:47:32', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36', '登录成功', '2022-07-18 11:47:32', '2022-07-18 11:47:32', NULL);
INSERT INTO `log_logins` VALUES (3403, 'admin', '0', '111.15.24.100', 'ɽ��ʡ������ �ƶ�', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-18 11:57:19', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-18 11:57:19', '2022-07-18 11:57:19', NULL);
INSERT INTO `log_logins` VALUES (3404, 'admin', '0', '127.0.0.1', '内部IP', 'Edge 92.0.902.84', 'Windows 10', 'Windows', '2022-07-18 11:57:39', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84', '登录成功', '2022-07-18 11:57:39', '2022-07-18 11:57:39', NULL);
INSERT INTO `log_logins` VALUES (3405, 'admin', '0', '124.92.142.49', '����ʡ������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-18 13:40:13', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-18 13:40:13', '2022-07-18 13:40:13', NULL);
INSERT INTO `log_logins` VALUES (3406, 'admin', '0', '120.231.240.109', '�㶫ʡ ��ͨ', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-18 13:59:58', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-18 13:59:58', '2022-07-18 13:59:58', NULL);
INSERT INTO `log_logins` VALUES (3407, 'admin', '0', '127.0.0.1', '内部IP', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-18 16:43:34', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-18 16:43:35', '2022-07-18 16:43:35', NULL);
INSERT INTO `log_logins` VALUES (3408, 'admin', '0', '27.17.253.211', '����ʡ�人�� ����', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-18 15:35:13', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-18 15:35:13', '2022-07-18 15:35:13', NULL);
INSERT INTO `log_logins` VALUES (3409, 'admin', '0', '27.17.253.211', '����ʡ�人�� ����', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-18 16:09:52', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-18 16:09:52', '2022-07-18 16:09:52', NULL);
INSERT INTO `log_logins` VALUES (3410, 'admin', '0', '119.97.203.46', '����ʡ�人�� ����', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-18 16:32:23', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-18 16:32:23', '2022-07-18 16:32:23', NULL);
INSERT INTO `log_logins` VALUES (3411, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-18 16:46:06', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-18 16:46:06', '2022-07-18 16:46:06', NULL);
INSERT INTO `log_logins` VALUES (3412, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-18 16:56:15', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-18 16:56:15', '2022-07-18 16:56:15', NULL);
INSERT INTO `log_logins` VALUES (3413, 'admin', '0', '182.35.53.116', 'ɽ��ʡ̩���� ����', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 00:09:41', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 00:09:41', '2022-07-19 00:09:41', NULL);
INSERT INTO `log_logins` VALUES (3414, 'admin', '0', '121.204.36.125', '����ʡ������ ����', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 09:59:02', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 09:59:02', '2022-07-19 09:59:02', NULL);
INSERT INTO `log_logins` VALUES (3415, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 10:50:42', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 10:50:42', '2022-07-19 10:50:42', NULL);
INSERT INTO `log_logins` VALUES (3416, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 10:51:14', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 10:51:14', '2022-07-19 10:51:14', NULL);
INSERT INTO `log_logins` VALUES (3417, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 11:39:18', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 11:39:18', '2022-07-19 11:39:18', NULL);
INSERT INTO `log_logins` VALUES (3418, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 13:42:19', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 13:42:19', '2022-07-19 13:42:19', NULL);
INSERT INTO `log_logins` VALUES (3419, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 13:42:53', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 13:42:53', '2022-07-19 13:42:53', NULL);
INSERT INTO `log_logins` VALUES (3420, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 13:43:59', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 13:43:59', '2022-07-19 13:43:59', NULL);
INSERT INTO `log_logins` VALUES (3421, 'admin', '0', '222.175.246.74', 'ɽ��ʡ�Ͳ��� ����', 'Chrome 92.0.4515.107', 'Windows 10', 'Windows', '2022-07-19 13:45:28', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36', '登录成功', '2022-07-19 13:45:28', '2022-07-19 13:45:28', NULL);
INSERT INTO `log_logins` VALUES (3422, 'admin', '0', '222.175.246.74', 'ɽ��ʡ�Ͳ��� ����', 'Chrome 92.0.4515.107', 'Windows 10', 'Windows', '2022-07-19 14:02:06', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36', '登录成功', '2022-07-19 14:02:06', '2022-07-19 14:02:06', NULL);
INSERT INTO `log_logins` VALUES (3423, 'panda', '0', '222.175.246.74', 'ɽ��ʡ�Ͳ��� ����', 'Chrome 92.0.4515.107', 'Windows 10', 'Windows', '2022-07-19 14:03:04', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36', '登录成功', '2022-07-19 14:03:04', '2022-07-19 14:03:04', NULL);
INSERT INTO `log_logins` VALUES (3424, 'admin', '0', '113.251.17.192', '������ ����', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-19 14:21:34', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-19 14:21:34', '2022-07-19 14:21:34', NULL);
INSERT INTO `log_logins` VALUES (3425, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 14:59:14', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 14:59:14', '2022-07-19 14:59:14', NULL);
INSERT INTO `log_logins` VALUES (3426, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 15:00:51', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 15:00:51', '2022-07-19 15:00:51', NULL);
INSERT INTO `log_logins` VALUES (3427, 'admin', '0', '222.94.127.75', '����ʡ�Ͼ��� (��ϼ��)����ADSL', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-19 15:03:38', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 15:03:38', '2022-07-19 15:03:38', NULL);
INSERT INTO `log_logins` VALUES (3428, 'admin', '0', '121.69.84.210', '������ ����ͨ', 'Chrome 104.0.0.0', 'Windows 10', 'Windows', '2022-07-19 15:05:33', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36', '登录成功', '2022-07-19 15:05:33', '2022-07-19 15:05:33', NULL);
INSERT INTO `log_logins` VALUES (3429, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 15:13:01', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 15:13:01', '2022-07-19 15:13:01', NULL);
INSERT INTO `log_logins` VALUES (3430, 'admin', '0', '171.35.156.50', '����ʡ������ ��ͨ', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 16:20:56', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 16:20:56', '2022-07-19 16:20:56', NULL);
INSERT INTO `log_logins` VALUES (3431, 'admin', '0', '182.116.65.22', '����ʡ������ ��ͨ', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-19 17:00:10', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 17:00:10', '2022-07-19 17:00:10', NULL);
INSERT INTO `log_logins` VALUES (3432, 'admin', '0', '114.102.130.30', '����ʡ�Ϸ��� ����', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-19 18:42:10', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 18:42:10', '2022-07-19 18:42:10', NULL);
INSERT INTO `log_logins` VALUES (3433, 'admin', '0', '218.75.42.2', '�㽭ʡ������ ����', 'Edge 105.0.1300.0', 'Windows 10', 'Windows', '2022-07-19 20:06:22', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.0 Safari/537.36 Edg/105.0.1300.0', '登录成功', '2022-07-19 20:06:22', '2022-07-19 20:06:22', NULL);
INSERT INTO `log_logins` VALUES (3434, 'admin', '0', '180.167.217.166', '�Ϻ��� ����', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-19 23:57:59', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-19 23:57:59', '2022-07-19 23:57:59', NULL);
INSERT INTO `log_logins` VALUES (3435, 'admin', '0', '222.209.109.51', '�Ĵ�ʡ�ɶ��� ����ADSL', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-20 09:08:05', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-20 09:08:05', '2022-07-20 09:08:05', NULL);
INSERT INTO `log_logins` VALUES (3436, 'admin', '0', '221.207.12.12', '�ຣʡ������ ��ͨ', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-20 09:33:15', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-20 09:33:15', '2022-07-20 09:33:15', NULL);
INSERT INTO `log_logins` VALUES (3437, 'admin', '0', '119.131.144.47', '�㶫ʡ������ ����', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-20 09:38:26', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 09:38:26', '2022-07-20 09:38:26', NULL);
INSERT INTO `log_logins` VALUES (3438, 'admin', '0', '218.108.15.68', '�㽭ʡ������ ��������', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-20 10:08:04', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-20 10:08:04', '2022-07-20 10:08:04', NULL);
INSERT INTO `log_logins` VALUES (3439, 'admin', '0', '39.170.37.78', '�㽭ʡ������ �ƶ�', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-20 11:22:22', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 11:22:22', '2022-07-20 11:22:22', NULL);
INSERT INTO `log_logins` VALUES (3440, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-20 12:02:55', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 12:02:55', '2022-07-20 12:02:55', NULL);
INSERT INTO `log_logins` VALUES (3441, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-20 13:10:13', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 13:10:13', '2022-07-20 13:10:13', NULL);
INSERT INTO `log_logins` VALUES (3442, 'admin', '0', '222.129.54.58', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-20 14:47:53', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 14:47:53', '2022-07-20 14:47:53', NULL);
INSERT INTO `log_logins` VALUES (3443, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-20 14:51:58', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 14:51:58', '2022-07-20 14:51:58', NULL);
INSERT INTO `log_logins` VALUES (3444, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-20 14:52:15', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 14:52:15', '2022-07-20 14:52:15', NULL);
INSERT INTO `log_logins` VALUES (3445, 'admin', '0', '175.9.209.183', '����ʡ��ɳ�� ����', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-20 15:14:28', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 15:14:28', '2022-07-20 15:14:28', NULL);
INSERT INTO `log_logins` VALUES (3446, 'admin', '0', '175.9.209.183', '����ʡ��ɳ�� ����', 'Firefox 102.0', 'Intel Mac OS X 10.15', 'Macintosh', '2022-07-20 17:41:36', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0', '登录成功', '2022-07-20 17:41:36', '2022-07-20 17:41:36', NULL);
INSERT INTO `log_logins` VALUES (3447, 'admin', '0', '39.128.79.75', '����ʡ��ͨ�� ��ͨ', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-20 18:26:13', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 18:26:13', '2022-07-20 18:26:13', NULL);
INSERT INTO `log_logins` VALUES (3448, 'admin', '0', '113.251.17.192', '������ ����', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-20 18:35:08', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-20 18:35:08', '2022-07-20 18:35:08', NULL);
INSERT INTO `log_logins` VALUES (3449, 'admin', '0', '106.201.241.21', ' ӡ��BHARTIAIRTEL�û�', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-20 20:49:47', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-20 20:49:47', '2022-07-20 20:49:47', NULL);
INSERT INTO `log_logins` VALUES (3450, 'admin', '0', '211.97.121.88', '����ʡ������ ��ͨ', 'Chrome 100.0.4896.75', 'Windows 10', 'Windows', '2022-07-20 22:01:27', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36', '登录成功', '2022-07-20 22:01:27', '2022-07-20 22:01:27', NULL);
INSERT INTO `log_logins` VALUES (3451, 'admin', '0', '115.60.131.32', '����ʡ֣���� ��ͨ', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-20 22:54:53', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-20 22:54:53', '2022-07-20 22:54:53', NULL);
INSERT INTO `log_logins` VALUES (3452, 'admin', '0', '120.235.110.204', '�㶫ʡ������ �ƶ�', 'Chrome 100.0.4896.127', 'Windows 10', 'Windows', '2022-07-21 00:19:51', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36', '登录成功', '2022-07-21 00:19:51', '2022-07-21 00:19:51', NULL);
INSERT INTO `log_logins` VALUES (3453, 'admin', '0', '61.144.53.34', '�㶫ʡ��������ɳ�� ����ADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 01:54:29', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 01:54:29', '2022-07-21 01:54:29', NULL);
INSERT INTO `log_logins` VALUES (3454, 'admin', '0', '111.197.236.161', '������ ��ͨ', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 09:14:52', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 09:14:52', '2022-07-21 09:14:52', NULL);
INSERT INTO `log_logins` VALUES (3455, 'admin', '0', '222.129.41.224', '������ ��ͨADSL', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 13:12:35', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 13:12:35', '2022-07-21 13:12:35', NULL);
INSERT INTO `log_logins` VALUES (3456, 'admin', '0', '110.87.82.45', '����ʡ������ ����', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 13:46:34', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 13:46:34', '2022-07-21 13:46:34', NULL);
INSERT INTO `log_logins` VALUES (3457, 'admin', '0', '103.239.206.29', '�㶫ʡ������ BGP', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 13:47:11', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 13:47:11', '2022-07-21 13:47:11', NULL);
INSERT INTO `log_logins` VALUES (3458, 'admin', '0', '106.2.199.3', '������ �����ƴ�ʵҵ���޹�˾', 'Safari 15.5', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-21 13:57:02', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.5 Safari/605.1.15', '登录成功', '2022-07-21 13:57:02', '2022-07-21 13:57:02', NULL);
INSERT INTO `log_logins` VALUES (3459, 'admin', '0', '182.150.25.236', '�Ĵ�ʡ�ɶ��� ����', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-21 14:15:08', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-21 14:15:08', '2022-07-21 14:15:08', NULL);
INSERT INTO `log_logins` VALUES (3460, 'admin', '0', '106.85.137.109', '������ ����', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 14:25:27', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 14:25:27', '2022-07-21 14:25:27', NULL);
INSERT INTO `log_logins` VALUES (3461, 'admin', '0', '114.84.116.248', '�Ϻ��� ����', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_14_6', 'Macintosh', '2022-07-21 14:27:46', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 14:27:46', '2022-07-21 14:27:46', NULL);
INSERT INTO `log_logins` VALUES (3462, 'admin', '0', '61.242.128.44', '����ʡ ��ͨ', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-21 15:34:55', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 15:34:55', '2022-07-21 15:34:55', NULL);
INSERT INTO `log_logins` VALUES (3463, 'admin', '0', '183.66.202.50', '������ ����', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 15:48:06', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 15:48:06', '2022-07-21 15:48:06', NULL);
INSERT INTO `log_logins` VALUES (3464, 'admin', '0', '36.152.100.44', '����ʡ�Ͼ��� �ƶ�', 'Chrome 102.0.0.0', 'Windows 10', 'Windows', '2022-07-21 15:48:43', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36', '登录成功', '2022-07-21 15:48:43', '2022-07-21 15:48:43', NULL);
INSERT INTO `log_logins` VALUES (3465, 'admin', '0', '36.152.100.44', '����ʡ�Ͼ��� �ƶ�', 'Chrome 102.0.0.0', 'Windows 10', 'Windows', '2022-07-21 16:26:32', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36', '登录成功', '2022-07-21 16:26:32', '2022-07-21 16:26:32', NULL);
INSERT INTO `log_logins` VALUES (3466, 'admin', '0', '175.42.2.18', '����ʡ������ ��ͨ', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 17:31:16', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 17:31:16', '2022-07-21 17:31:16', NULL);
INSERT INTO `log_logins` VALUES (3467, 'admin', '0', '113.108.77.68', '�㶫ʡ������ ����', 'Edge 103.0.1264.62', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-21 17:34:53', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-21 17:34:53', '2022-07-21 17:34:53', NULL);
INSERT INTO `log_logins` VALUES (3468, 'test', '0', '112.49.232.149', '����ʡ������ ��ͨ', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-21 18:29:57', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-21 18:29:57', '2022-07-21 18:29:57', NULL);
INSERT INTO `log_logins` VALUES (3469, 'admin', '0', '60.217.68.103', 'ɽ��ʡ������ ��ͨ', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-21 18:44:51', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-21 18:44:51', '2022-07-21 18:44:51', NULL);
INSERT INTO `log_logins` VALUES (3470, 'admin', '0', '58.60.1.47', '�㶫ʡ������ ����', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-22 00:35:34', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-22 00:35:34', '2022-07-22 00:35:34', NULL);
INSERT INTO `log_logins` VALUES (3471, 'admin', '0', '60.217.68.103', 'ɽ��ʡ������ ��ͨ', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-22 09:01:29', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-22 09:01:29', '2022-07-22 09:01:29', NULL);
INSERT INTO `log_logins` VALUES (3472, 'admin', '0', '159.138.32.243', '���� ��Ϊ��˾', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-22 11:11:36', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-22 11:11:36', '2022-07-22 11:11:36', NULL);
INSERT INTO `log_logins` VALUES (3473, 'admin', '0', '112.49.232.149', '����ʡ������ ��ͨ', 'Chrome 103.0.0.0', 'Windows 10', 'Windows', '2022-07-22 11:36:46', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-22 11:36:46', '2022-07-22 11:36:46', NULL);
INSERT INTO `log_logins` VALUES (3474, 'admin', '0', '165.154.224.13', ' ���ô�', 'Chrome 103.0.0.0', 'Linux x86_64', 'X11', '2022-07-22 13:26:15', '0', '0', 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-22 13:26:15', '2022-07-22 13:26:15', NULL);
INSERT INTO `log_logins` VALUES (3475, 'admin', '0', '120.40.98.110', '����ʡ������ ����', 'Chrome 94.0.4606.81', 'Windows 10', 'Windows', '2022-07-22 18:03:31', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36', '登录成功', '2022-07-22 18:03:31', '2022-07-22 18:03:31', NULL);
INSERT INTO `log_logins` VALUES (3476, 'admin', '0', '115.60.128.186', '����ʡ֣���� ��ͨ', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-22 21:53:52', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-22 21:53:52', '2022-07-22 21:53:52', NULL);
INSERT INTO `log_logins` VALUES (3477, 'admin', '0', '218.5.163.180', '����ʡȪ���� ����ADSL', 'Edge 103.0.1264.62', 'Windows 10', 'Windows', '2022-07-23 00:50:22', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36 Edg/103.0.1264.62', '登录成功', '2022-07-23 00:50:22', '2022-07-23 00:50:22', NULL);
INSERT INTO `log_logins` VALUES (3478, 'admin', '0', '120.229.69.82', '�㶫ʡ������ ��ͨ', 'Chrome 97.0.4692.71', 'Windows 10', 'Windows', '2022-07-23 07:59:03', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36', '登录成功', '2022-07-23 07:59:03', '2022-07-23 07:59:03', NULL);
INSERT INTO `log_logins` VALUES (3479, 'admin', '0', '118.113.5.72', '�Ĵ�ʡ�ɶ��� ����', 'Chrome 103.0.0.0', 'Intel Mac OS X 10_15_7', 'Macintosh', '2022-07-23 09:41:24', '0', '0', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36', '登录成功', '2022-07-23 09:41:24', '2022-07-23 09:41:24', NULL);
INSERT INTO `log_logins` VALUES (3480, 'admin', '0', '127.0.0.1', '内部IP', 'Chrome 92.0.4515.107', 'Windows 10', 'Windows', '2022-07-23 10:18:32', '0', '0', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36', '登录成功', '2022-07-23 10:18:32', '2022-07-23 10:18:32', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 928 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_opers
-- ----------------------------
INSERT INTO `log_opers` VALUES (912, '清空操作日志信息', '3', 'DELETE', 'panda', '/log/logOper/all', '127.0.0.1', '内部IP', '', '0', '2022-07-18 11:36:32', '2022-07-18 11:36:32', NULL);
INSERT INTO `log_opers` VALUES (913, '清空登录日志信息', '3', 'DELETE', 'panda', '/log/logLogin/all', '127.0.0.1', '内部IP', '', '0', '2022-07-18 11:37:43', '2022-07-18 11:37:43', NULL);
INSERT INTO `log_opers` VALUES (914, '修改角色信息', '2', 'PUT', 'panda', '/system/role', '222.175.246.74', 'ɽ��ʡ�Ͳ��� ����', '', '0', '2022-07-19 14:03:34', '2022-07-19 14:03:34', NULL);
INSERT INTO `log_opers` VALUES (915, '修改状态', '2', 'PUT', 'admin', '/job/changeStatus', '121.69.84.210', '������ ����ͨ', '', '0', '2022-07-19 15:07:09', '2022-07-19 15:07:09', NULL);
INSERT INTO `log_opers` VALUES (916, '添加通知信息', '1', 'POST', 'admin', '/system/notice', '182.116.65.22', '����ʡ������ ��ͨ', '', '0', '2022-07-19 17:06:20', '2022-07-19 17:06:20', NULL);
INSERT INTO `log_opers` VALUES (917, '修改状态', '2', 'PUT', 'admin', '/job/changeStatus', '121.69.84.210', '������ ����ͨ', '', '0', '2022-07-19 18:04:42', '2022-07-19 18:04:42', NULL);
INSERT INTO `log_opers` VALUES (918, '修改SysTenant信息', '2', 'PUT', 'admin', '/system/tenant', '114.102.130.30', '����ʡ�Ϸ��� ����', '', '0', '2022-07-19 18:45:18', '2022-07-19 18:45:18', NULL);
INSERT INTO `log_opers` VALUES (919, '修改状态', '2', 'PUT', 'admin', '/resource/oss/changeStatus', '180.167.217.166', '�Ϻ��� ����', '', '0', '2022-07-20 00:03:56', '2022-07-20 00:03:56', NULL);
INSERT INTO `log_opers` VALUES (920, '添加SysTenant信息', '1', 'POST', 'admin', '/system/tenant', '119.131.144.47', '�㶫ʡ������ ����', '', '0', '2022-07-20 09:59:02', '2022-07-20 09:59:02', NULL);
INSERT INTO `log_opers` VALUES (921, '修改角色信息', '2', 'PUT', 'admin', '/system/role', '222.129.41.224', '������ ��ͨADSL', '', '0', '2022-07-20 16:39:14', '2022-07-20 16:39:14', NULL);
INSERT INTO `log_opers` VALUES (922, '添加用户信息', '1', 'POST', 'admin', '/system/user', '120.235.110.204', '�㶫ʡ������ �ƶ�', '', '0', '2022-07-21 00:20:54', '2022-07-21 00:20:54', NULL);
INSERT INTO `log_opers` VALUES (923, '修改用户状态', '2', 'PUT', 'admin', '/system/user/changeStatus', '120.235.110.204', '�㶫ʡ������ �ƶ�', '', '0', '2022-07-21 00:21:05', '2022-07-21 00:21:05', NULL);
INSERT INTO `log_opers` VALUES (924, '修改用户状态', '2', 'PUT', 'admin', '/system/user/changeStatus', '120.235.110.204', '�㶫ʡ������ �ƶ�', '', '0', '2022-07-21 00:21:07', '2022-07-21 00:21:07', NULL);
INSERT INTO `log_opers` VALUES (925, '修改用户状态', '2', 'PUT', 'admin', '/system/user/changeStatus', '120.235.110.204', '�㶫ʡ������ �ƶ�', '', '0', '2022-07-21 00:21:07', '2022-07-21 00:21:07', NULL);
INSERT INTO `log_opers` VALUES (926, '添加用户信息', '1', 'POST', 'admin', '/system/user', '103.239.206.29', '�㶫ʡ������ BGP', '', '0', '2022-07-21 13:51:57', '2022-07-21 13:51:57', NULL);
INSERT INTO `log_opers` VALUES (927, '新增表', '1', 'POST', 'admin', '/develop/code/table', '159.138.32.243', '���� ��Ϊ��˾', '', '0', '2022-07-22 11:13:38', '2022-07-22 11:13:38', NULL);

-- ----------------------------
-- Table structure for res_emails
-- ----------------------------
DROP TABLE IF EXISTS `res_emails`;
CREATE TABLE `res_emails`  (
  `mail_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `category` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `host` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `port` bigint(20) NULL DEFAULT NULL,
  `from` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nickname` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `secret` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_ssl` tinyint(1) NULL DEFAULT NULL,
  `status` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`mail_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of res_emails
-- ----------------------------
INSERT INTO `res_emails` VALUES (1, '0', 'smtp.163.com', 25, '18610165312@163.com', '熊猫', 'DCXZCAGTCMSEGPZL', 0, '0', '2022-01-14 16:14:37', '2022-01-17 10:27:57', NULL);

-- ----------------------------
-- Table structure for res_osses
-- ----------------------------
DROP TABLE IF EXISTS `res_osses`;
CREATE TABLE `res_osses`  (
  `oss_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `category` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `app_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'AppId',
  `access_key` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'accessKey',
  `secret_key` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'secretKey',
  `bucket_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'bucketName',
  `endpoint` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'endpoint',
  `oss_code` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ossCode',
  `region` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地址',
  `remark` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '说明',
  `status` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '状态',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`oss_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of res_osses
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 124 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `sys_apis` VALUES (59, '2021-12-09 19:50:57', '2022-01-19 08:58:20', NULL, '/system/api/list', '获取api分页列表1', 'api', 'GET');
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
INSERT INTO `sys_apis` VALUES (88, '2022-01-02 13:53:06', '2022-07-18 10:57:58', NULL, '/develop/code/table/db/list', '数据库表列表', 'gen', 'GET');
INSERT INTO `sys_apis` VALUES (89, '2022-01-02 13:53:44', '2022-01-02 13:53:44', NULL, '/develop/code/table/list', '表列表', 'gen', 'GET');
INSERT INTO `sys_apis` VALUES (90, '2022-01-02 13:54:10', '2022-01-02 13:54:10', NULL, '/develop/code/table/info/:tableId', '表信息', 'gen', 'GET');
INSERT INTO `sys_apis` VALUES (91, '2022-01-02 13:54:42', '2022-07-18 10:58:35', NULL, '/develop/code/table/info/tableName', '表名获取表信息', 'gen', 'GET');
INSERT INTO `sys_apis` VALUES (92, '2022-01-02 13:55:13', '2022-01-02 13:55:13', NULL, '/develop/code/table/tableTree', '表树', 'gen', 'GET');
INSERT INTO `sys_apis` VALUES (93, '2022-01-02 13:56:37', '2022-01-02 13:56:37', NULL, '/develop/code/table', '导入表', 'gen', 'POST');
INSERT INTO `sys_apis` VALUES (94, '2022-01-02 13:57:36', '2022-01-02 13:57:36', NULL, '/develop/code/table', '修改代码生成信息', 'gen', 'PUT');
INSERT INTO `sys_apis` VALUES (95, '2022-01-02 13:58:25', '2022-01-02 13:58:25', NULL, '/develop/code/table/:tableId', '删除表数据', 'gen', 'DELETE');
INSERT INTO `sys_apis` VALUES (96, '2022-01-02 13:59:07', '2022-01-02 13:59:07', NULL, '/develop/code/gen/preview/:tableId', '预览代码', 'gen', 'GET');
INSERT INTO `sys_apis` VALUES (97, '2022-01-02 13:59:43', '2022-01-02 13:59:43', NULL, '/develop/code/gen/code/:tableId', '生成代码', 'gen', 'GET');
INSERT INTO `sys_apis` VALUES (98, '2022-01-02 14:00:10', '2022-07-17 01:19:42', NULL, '/develop/code/gen/configure/:tableId', '生成api菜单', 'gen', 'GET');
INSERT INTO `sys_apis` VALUES (99, '2022-01-13 16:44:44', '2022-01-13 16:45:27', NULL, '/resource/oss/list', '获取oss列表', 'oss', 'GET');
INSERT INTO `sys_apis` VALUES (100, '2022-01-13 16:44:44', '2022-01-13 16:44:44', NULL, '/resource/oss/:ossId', '获取oss', 'oss', 'GET');
INSERT INTO `sys_apis` VALUES (101, '2022-01-13 16:44:44', '2022-01-13 16:44:44', NULL, '/resource/oss', '添加oss', 'oss', 'POST');
INSERT INTO `sys_apis` VALUES (102, '2022-01-13 16:44:44', '2022-01-13 16:44:44', NULL, '/resource/oss', '修改oss', 'oss', 'PUT');
INSERT INTO `sys_apis` VALUES (103, '2022-01-13 16:44:44', '2022-01-13 16:44:44', NULL, '/resource/oss/:ossId', '删除oss', 'oss', 'DELETE');
INSERT INTO `sys_apis` VALUES (104, '2022-01-14 13:19:21', '2022-01-14 13:19:21', NULL, '/resource/oss/changeStatus', '修改状态', 'oss', 'PUT');
INSERT INTO `sys_apis` VALUES (105, '2022-01-14 13:20:14', '2022-01-14 13:20:14', NULL, '/resource/oss/uploadFile', '调试上传文件', 'oss', 'POST');
INSERT INTO `sys_apis` VALUES (106, '2022-01-14 15:30:39', '2022-01-14 15:30:39', NULL, '/resource/email/list', '邮件分页列表', 'mail', 'GET');
INSERT INTO `sys_apis` VALUES (107, '2022-01-14 15:31:20', '2022-01-14 15:31:20', NULL, '/resource/email/:mailId', '获取邮件', 'mail', 'GET');
INSERT INTO `sys_apis` VALUES (108, '2022-01-14 15:31:54', '2022-01-14 15:31:54', NULL, '/resource/email', '添加邮件配置', 'mail', 'POST');
INSERT INTO `sys_apis` VALUES (109, '2022-01-14 15:32:21', '2022-01-14 15:32:21', NULL, '/resource/email', '修改邮件配置', 'mail', 'PUT');
INSERT INTO `sys_apis` VALUES (110, '2022-01-14 15:32:53', '2022-01-14 15:32:53', NULL, '/resource/email/:mailId', '删除邮件', 'mail', 'DELETE');
INSERT INTO `sys_apis` VALUES (111, '2022-01-14 17:11:42', '2022-01-14 17:11:42', NULL, '/resource/email/changeStatus', '修改状态', 'mail', 'PUT');
INSERT INTO `sys_apis` VALUES (112, '2022-01-14 17:12:17', '2022-01-14 17:12:17', NULL, '/resource/email/debugMail', '发送邮件调试', 'mail', 'POST');
INSERT INTO `sys_apis` VALUES (113, '2022-07-15 18:06:27', '2022-07-15 18:06:27', NULL, '/system/tenant/list', '租户列表', 'tenant', 'GET');
INSERT INTO `sys_apis` VALUES (114, '2022-07-15 18:07:16', '2022-07-15 18:07:16', NULL, '/system/tenant/:tenantId', '获取租户', 'tenant', 'GET');
INSERT INTO `sys_apis` VALUES (115, '2022-07-15 18:07:43', '2022-07-15 18:07:43', NULL, '/system/tenant', '添加租户', 'tenant', 'POST');
INSERT INTO `sys_apis` VALUES (116, '2022-07-15 18:08:08', '2022-07-15 18:08:08', NULL, '/system/tenant', '修改租户', 'tenant', 'PUT');
INSERT INTO `sys_apis` VALUES (117, '2022-07-15 18:08:57', '2022-07-15 18:08:57', NULL, '/system/tenant/:tenantId', '删除租户', 'tenant', 'DELETE');
INSERT INTO `sys_apis` VALUES (123, '2022-07-18 10:24:03', '2022-07-18 10:24:03', NULL, '/system/tenant/lists', '获取所有租户', 'tenant', 'GET');

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
  `is_frontend` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '是否前台',
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
  `tenant_id` int(11) NULL DEFAULT NULL COMMENT '租户id',
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
INSERT INTO `sys_depts` VALUES (2, NULL, 0, '/0/2', '熊猫科技', 0, 'xm', '18353366836', '342@qq.com', '0', 'admin', 'admin', '2021-12-01 17:31:53', '2021-12-02 08:56:19', NULL);
INSERT INTO `sys_depts` VALUES (3, NULL, 2, '/0/2/3', '研发部', 1, 'panda', '18353366543', 'ewr@qq.com', '0', 'admin', 'admin', '2021-12-01 17:37:43', '2021-12-02 08:55:56', NULL);
INSERT INTO `sys_depts` VALUES (7, NULL, 2, '/0/2/7', '营销部', 2, 'panda', '18353333333', '342@qq.com', '0', 'panda', 'panda', '2021-12-24 10:46:24', '2021-12-24 10:47:15', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `sys_dict_data` VALUES (31, 0, '阿里云', '0', 'res_oss_category', '0', '', '', '', 'panda', '', '', '2022-01-13 15:44:01', '2022-01-13 15:44:01', NULL);
INSERT INTO `sys_dict_data` VALUES (32, 1, '七牛云', '1', 'res_oss_category', '0', '', '', '', 'panda', '', '', '2022-01-13 15:44:18', '2022-01-13 15:44:18', NULL);
INSERT INTO `sys_dict_data` VALUES (33, 2, '腾讯云', '2', 'res_oss_category', '0', '', '', '', 'panda', '', '', '2022-01-13 15:44:39', '2022-01-13 15:44:39', NULL);
INSERT INTO `sys_dict_data` VALUES (34, 0, '阿里云', '0', 'res_sms_category', '0', '', '', '', 'panda', '', '', '2022-01-13 15:47:30', '2022-01-13 15:47:30', NULL);
INSERT INTO `sys_dict_data` VALUES (35, 1, '腾讯云', '1', 'res_sms_category', '0', '', '', '', 'panda', '', '', '2022-01-13 15:47:39', '2022-01-13 15:47:39', NULL);
INSERT INTO `sys_dict_data` VALUES (36, 0, '163邮箱', '0', 'res_mail_category', '0', '', '', '', 'panda', '', '', '2022-01-14 15:43:42', '2022-01-14 15:43:42', NULL);
INSERT INTO `sys_dict_data` VALUES (37, 0, 'qq邮箱', '1', 'res_mail_category', '0', '', '', '', 'panda', '', '', '2022-01-14 15:44:08', '2022-01-14 15:44:08', NULL);
INSERT INTO `sys_dict_data` VALUES (38, 0, '企业邮箱', '2', 'res_mail_category', '0', '', '', '', 'panda', '', '', '2022-01-14 15:44:20', '2022-01-14 15:44:20', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `sys_dict_types` VALUES (30, 'oss分类', 'res_oss_category', '0', 'panda', '', 'oss分类列表', '2022-01-13 15:43:29', '2022-01-13 15:43:29', NULL);
INSERT INTO `sys_dict_types` VALUES (31, 'sms分类', 'res_sms_category', '0', 'panda', '', 'sms分类列表', '2021-12-26 15:23:26', '2022-01-13 15:47:13', NULL);
INSERT INTO `sys_dict_types` VALUES (32, 'mail分类', 'res_mail_category', '0', 'panda', '', 'mail分类列表', '2022-01-14 15:43:17', '2022-01-14 15:43:17', NULL);

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
INSERT INTO `sys_jobs` VALUES (1, 'testcron', 'SYSTEM', '2', ' 0/10 * * * * ?', 'cronHandle', 'aaa', '', '', '1', 0, 'panda', '', '2021-12-24 16:06:09', '2022-01-22 08:11:24', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 95 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (1, '系统设置', '', 0, 0, 'elementSetting', '/system', 'Layout', '1', '', 'M', '0', '0', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08', '2021-12-28 13:32:21', NULL);
INSERT INTO `sys_menus` VALUES (3, '用户管理', '', 1, 1, 'elementUser', '/system/user', '/system/user/index', '1', '', 'C', '0', '1', '1', 'system:user:list', '0', 'admin', 'panda', '', '2021-12-02 14:07:56', '2021-12-28 13:32:44', NULL);
INSERT INTO `sys_menus` VALUES (4, '添加用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:add', '0', 'admin', '', '', '2021-12-03 13:36:33', '2021-12-03 13:36:33', NULL);
INSERT INTO `sys_menus` VALUES (5, '编辑用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:edit', '0', 'admin', '', '', '2021-12-03 13:48:13', '2021-12-03 13:48:13', NULL);
INSERT INTO `sys_menus` VALUES (6, '角色管理', '', 1, 2, 'elementUserFilled', '/system/role', '/system/role/index', '1', '', 'C', '0', '1', '1', 'system:role:list', '0', '', 'panda', '', '2021-12-03 13:51:55', '2022-07-16 10:23:21', NULL);
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
INSERT INTO `sys_menus` VALUES (39, 'API管理', '', 1, 2, 'iconfont icon-siweidaotu', '/system/api', '/system/api/index', '1', '', 'C', '0', '1', '1', 'system:api:list', '0', '', 'panda', '', '2021-12-09 09:09:13', '2022-07-16 10:23:42', NULL);
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
INSERT INTO `sys_menus` VALUES (50, '表单构建', '', 49, 1, 'iconfont icon-zidingyibuju', '/develop/form', '/develop/form/index', '1', '', 'C', '0', '1', '1', 'develop:form:list', '0', 'admin', 'panda', '', '2021-12-16 16:55:01', '2022-07-12 18:56:18', NULL);
INSERT INTO `sys_menus` VALUES (51, '代码生成', '', 49, 2, 'iconfont icon-zhongduancanshu', '/develop/code', '/develop/code/index', '1', '', 'C', '0', '1', '1', 'develop:code:list', '0', 'admin', '', '', '2021-12-16 16:56:48', '2021-12-16 16:56:48', NULL);
INSERT INTO `sys_menus` VALUES (52, '系统接口', '', 49, 3, 'iconfont icon-wenducanshu-05', '/develop/apis', '/layout/routerView/iframes', '0', 'http://47.104.252.2:8080/swagger/index.html', 'C', '0', '1', '1', 'develop:apis:list', '0', '', 'panda', '', '2021-12-16 16:58:07', '2022-07-13 11:50:34', NULL);
INSERT INTO `sys_menus` VALUES (53, '资源管理', '', 0, 4, 'iconfont icon-juxingkaobei', '/resource', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', '', '', '2021-12-16 17:02:06', '2021-12-16 17:02:06', NULL);
INSERT INTO `sys_menus` VALUES (54, '对象存储', '', 53, 1, 'iconfont icon-chazhaobiaodanliebiao', '/resource/file', '/resource/file/index', '1', '', 'C', '0', '1', '1', 'resource:file:list', '0', 'admin', 'panda', '', '2021-12-16 17:06:04', '2022-01-13 17:30:09', NULL);
INSERT INTO `sys_menus` VALUES (55, '公告通知', '', 44, 3, 'elementTicket', '/tool/notice', '/tool/notice/index', '1', '', 'C', '0', '1', '1', 'tool:notice:list', '0', 'admin', 'panda', '', '2021-12-16 22:09:11', '2021-12-28 13:42:39', NULL);
INSERT INTO `sys_menus` VALUES (56, '任务日志', '', 43, 1, 'iconfont icon--chaifenhang', '/log/job', '/log/job/index', '1', '', 'C', '0', '1', '1', 'log:job:list', '0', 'panda', 'panda', '', '2021-12-24 22:13:45', '2021-12-28 13:39:52', NULL);
INSERT INTO `sys_menus` VALUES (57, '邮件管理', '', 53, 1, 'iconfont icon-shouye_dongtaihui', '/resource/mail', '/resource/mail/index', '1', '', 'C', '0', '1', '1', 'resource:mail:list', '0', 'panda', 'panda', '', '2021-12-28 15:12:37', '2021-12-28 15:12:45', NULL);
INSERT INTO `sys_menus` VALUES (58, '短信管理', '', 53, 3, 'elementChatDotRound', '/resource/message', '/resource/message/index', '1', '', 'C', '0', '1', '1', 'resource:message:list', '0', 'panda', '', '', '2021-12-16 17:06:04', '2021-12-16 17:06:04', NULL);
INSERT INTO `sys_menus` VALUES (59, '删除', '', 45, 1, '', '', '', '', '', 'F', '', '', '', 'log:operation:delete', '0', 'panda', '', '', '2022-01-14 13:28:25', '2022-01-14 13:28:25', NULL);
INSERT INTO `sys_menus` VALUES (60, '清空', '', 45, 1, '', '', '', '', '', 'F', '', '', '', 'log:operation:clean', '0', 'panda', '', '', '2022-01-14 13:29:24', '2022-01-14 13:29:24', NULL);
INSERT INTO `sys_menus` VALUES (61, '删除', '', 56, 1, '', '', '', '', '', 'F', '', '', '', 'log:job:delete', '0', 'panda', '', '', '2022-01-14 13:29:57', '2022-01-14 13:29:57', NULL);
INSERT INTO `sys_menus` VALUES (62, '清空', '', 56, 1, '', '', '', '', '', 'F', '', '', '', 'log:job:clean', '0', 'panda', '', '', '2022-01-14 13:30:15', '2022-01-14 13:30:15', NULL);
INSERT INTO `sys_menus` VALUES (63, '删除', '', 46, 1, '', '', '', '', '', 'F', '', '', '', 'log:login:delete', '0', 'panda', '', '', '2022-01-14 13:30:46', '2022-01-14 13:30:46', NULL);
INSERT INTO `sys_menus` VALUES (64, '清空', '', 46, 1, '', '', '', '', '', 'F', '', '', '', 'log:login:clean', '0', 'panda', '', '', '2022-01-14 13:31:06', '2022-01-14 13:31:06', NULL);
INSERT INTO `sys_menus` VALUES (65, '新增', '', 48, 1, '', '', '', '', '', 'F', '', '', '', 'tool:job:add', '0', 'panda', '', '', '2022-01-14 13:32:48', '2022-01-14 13:32:48', NULL);
INSERT INTO `sys_menus` VALUES (66, '编辑', '', 48, 1, '', '', '', '', '', 'F', '', '', '', 'tool:job:edit', '0', 'panda', '', '', '2022-01-14 13:33:17', '2022-01-14 13:33:17', NULL);
INSERT INTO `sys_menus` VALUES (67, '删除', '', 48, 1, '', '', '', '', '', 'F', '', '', '', 'tool:job:delete', '0', 'panda', '', '', '2022-01-14 13:33:43', '2022-01-14 13:33:43', NULL);
INSERT INTO `sys_menus` VALUES (68, '开关', '', 48, 1, '', '', '', '', '', 'F', '', '', '', 'tool:job:run', '0', 'panda', '', '', '2022-01-14 13:34:27', '2022-01-14 13:34:27', NULL);
INSERT INTO `sys_menus` VALUES (69, '添加', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:add', '0', 'panda', '', '', '2022-01-14 13:35:23', '2022-01-14 13:35:23', NULL);
INSERT INTO `sys_menus` VALUES (70, '编辑', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:edit', '0', 'panda', '', '', '2022-01-14 13:36:04', '2022-01-14 13:36:04', NULL);
INSERT INTO `sys_menus` VALUES (71, '删除', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:delete', '0', 'panda', '', '', '2022-01-14 13:36:26', '2022-01-14 13:36:26', NULL);
INSERT INTO `sys_menus` VALUES (72, '查看', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:view', '0', 'panda', '', '', '2022-01-14 13:36:51', '2022-01-14 13:36:51', NULL);
INSERT INTO `sys_menus` VALUES (73, '导入', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:add', '0', 'panda', '', '', '2022-01-14 13:38:35', '2022-01-14 13:38:35', NULL);
INSERT INTO `sys_menus` VALUES (74, '编辑', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:edit', '0', 'panda', '', '', '2022-01-14 13:41:25', '2022-01-14 13:41:25', NULL);
INSERT INTO `sys_menus` VALUES (75, '删除', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:delete', '0', 'panda', '', '', '2022-01-14 13:41:42', '2022-01-14 13:41:42', NULL);
INSERT INTO `sys_menus` VALUES (76, '预览', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:view', '0', 'panda', '', '', '2022-01-14 13:42:01', '2022-01-14 13:42:01', NULL);
INSERT INTO `sys_menus` VALUES (77, '生成代码', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:code', '0', 'panda', '', '', '2022-01-14 13:42:48', '2022-01-14 13:42:48', NULL);
INSERT INTO `sys_menus` VALUES (78, '添加', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:add', '0', 'panda', '', '', '2022-01-17 11:26:15', '2022-01-17 11:26:15', NULL);
INSERT INTO `sys_menus` VALUES (79, '编辑', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:edit', '0', 'panda', '', '', '2022-01-17 11:26:39', '2022-01-17 11:26:39', NULL);
INSERT INTO `sys_menus` VALUES (80, '删除', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:delete', '0', 'panda', '', '', '2022-01-17 11:26:56', '2022-01-17 11:26:56', NULL);
INSERT INTO `sys_menus` VALUES (81, '调试', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:debug', '0', 'panda', '', '', '2022-01-17 11:27:14', '2022-01-17 11:27:14', NULL);
INSERT INTO `sys_menus` VALUES (82, '调试', '', 54, 1, '', '', '', '', '', 'F', '', '', '', 'resource:file:debug', '0', 'panda', '', '', '2022-01-17 11:27:17', '2022-01-17 11:27:17', '2022-01-17 11:27:25');
INSERT INTO `sys_menus` VALUES (83, '添加', '', 57, 1, '', '', '', '', '', 'F', '', '', '', 'resource:mail:add', '0', 'panda', '', '', '2022-01-17 11:28:06', '2022-01-17 11:28:06', NULL);
INSERT INTO `sys_menus` VALUES (84, '编辑', '', 57, 1, '', '', '', '', '', 'F', '', '', '', 'resource:mail:edit', '0', 'panda', '', '', '2022-01-17 11:28:37', '2022-01-17 11:28:37', NULL);
INSERT INTO `sys_menus` VALUES (85, '删除', '', 57, 1, '', '', '', '', '', 'F', '', '', '', 'resource:mail:delete', '0', 'panda', '', '', '2022-01-17 11:29:09', '2022-01-17 11:29:09', NULL);
INSERT INTO `sys_menus` VALUES (86, '调试', '', 57, 1, '', '', '', '', '', 'F', '', '', '', 'resource:mail:debug', '0', 'panda', '', '', '2022-01-17 11:29:46', '2022-01-17 11:29:46', NULL);
INSERT INTO `sys_menus` VALUES (87, '租户管理', '', 1, 1, 'iconfont icon-quanxian', '/system/tenant', '/system/tenant/index', '1', '', 'C', '0', '1', '1', 'system:tenant:list', '0', 'panda', '', '', '2022-07-15 18:03:35', '2022-07-15 18:03:35', NULL);
INSERT INTO `sys_menus` VALUES (88, '添加', '', 87, 1, '', '', '', '', '', 'F', '', '', '', 'system:tenant:add', '0', 'panda', '', '', '2022-07-15 18:28:58', '2022-07-15 18:28:58', NULL);
INSERT INTO `sys_menus` VALUES (89, '编辑', '', 87, 1, '', '', '', '', '', 'F', '', '', '', 'system:tenant:edit', '0', 'panda', '', '', '2022-07-15 18:29:34', '2022-07-15 18:29:34', NULL);
INSERT INTO `sys_menus` VALUES (90, '删除', '', 87, 1, '', '', '', '', '', 'F', '', '', '', 'system:tenant:delete', '0', 'panda', '', '', '2022-07-15 18:30:00', '2022-07-15 18:30:00', NULL);

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
  `tenant_id` int(11) NULL DEFAULT NULL COMMENT '租户id',
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
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_posts
-- ----------------------------
INSERT INTO `sys_posts` VALUES (1, 1, '首席执行官', 'CEO', 1, '0', '首席执行官', 'panda', '', '2021-12-02 09:21:44', '2022-07-16 17:36:32', NULL);
INSERT INTO `sys_posts` VALUES (3, 1, '首席技术执行官', 'CTO', 1, '0', '', 'admin', '', '2021-12-02 09:21:44', '2021-12-02 09:25:59', '2021-12-02 09:27:41');
INSERT INTO `sys_posts` VALUES (4, 1, '首席技术执行官', 'CTO', 2, '0', '', 'panda', '', '2021-12-02 09:21:44', '2022-07-16 17:37:42', NULL);
INSERT INTO `sys_posts` VALUES (5, 1, '123', '123', 0, '0', '', 'admin', '', '2021-12-18 00:33:28', '2021-12-18 00:33:28', '2021-12-28 14:11:52');
INSERT INTO `sys_posts` VALUES (6, 1, 'wr', 'rw', 0, '0', '', 'panda', '', '2022-07-16 16:55:57', '2022-07-16 16:55:57', '2022-07-16 16:56:08');

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
) ENGINE = InnoDB AUTO_INCREMENT = 3043 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO `sys_role_menus` VALUES (2870, 1, 1, 'admin');
INSERT INTO `sys_role_menus` VALUES (2871, 1, 3, 'admin');
INSERT INTO `sys_role_menus` VALUES (2872, 1, 4, 'admin');
INSERT INTO `sys_role_menus` VALUES (2873, 1, 5, 'admin');
INSERT INTO `sys_role_menus` VALUES (2874, 1, 6, 'admin');
INSERT INTO `sys_role_menus` VALUES (2875, 1, 7, 'admin');
INSERT INTO `sys_role_menus` VALUES (2876, 1, 8, 'admin');
INSERT INTO `sys_role_menus` VALUES (2877, 1, 9, 'admin');
INSERT INTO `sys_role_menus` VALUES (2878, 1, 10, 'admin');
INSERT INTO `sys_role_menus` VALUES (2879, 1, 11, 'admin');
INSERT INTO `sys_role_menus` VALUES (2880, 1, 12, 'admin');
INSERT INTO `sys_role_menus` VALUES (2881, 1, 13, 'admin');
INSERT INTO `sys_role_menus` VALUES (2882, 1, 14, 'admin');
INSERT INTO `sys_role_menus` VALUES (2883, 1, 15, 'admin');
INSERT INTO `sys_role_menus` VALUES (2884, 1, 16, 'admin');
INSERT INTO `sys_role_menus` VALUES (2885, 1, 17, 'admin');
INSERT INTO `sys_role_menus` VALUES (2886, 1, 18, 'admin');
INSERT INTO `sys_role_menus` VALUES (2887, 1, 19, 'admin');
INSERT INTO `sys_role_menus` VALUES (2888, 1, 20, 'admin');
INSERT INTO `sys_role_menus` VALUES (2889, 1, 21, 'admin');
INSERT INTO `sys_role_menus` VALUES (2890, 1, 22, 'admin');
INSERT INTO `sys_role_menus` VALUES (2891, 1, 23, 'admin');
INSERT INTO `sys_role_menus` VALUES (2892, 1, 24, 'admin');
INSERT INTO `sys_role_menus` VALUES (2893, 1, 25, 'admin');
INSERT INTO `sys_role_menus` VALUES (2894, 1, 26, 'admin');
INSERT INTO `sys_role_menus` VALUES (2895, 1, 28, 'admin');
INSERT INTO `sys_role_menus` VALUES (2896, 1, 29, 'admin');
INSERT INTO `sys_role_menus` VALUES (2897, 1, 30, 'admin');
INSERT INTO `sys_role_menus` VALUES (2898, 1, 31, 'admin');
INSERT INTO `sys_role_menus` VALUES (2899, 1, 32, 'admin');
INSERT INTO `sys_role_menus` VALUES (2900, 1, 33, 'admin');
INSERT INTO `sys_role_menus` VALUES (2901, 1, 34, 'admin');
INSERT INTO `sys_role_menus` VALUES (2902, 1, 35, 'admin');
INSERT INTO `sys_role_menus` VALUES (2903, 1, 36, 'admin');
INSERT INTO `sys_role_menus` VALUES (2904, 1, 37, 'admin');
INSERT INTO `sys_role_menus` VALUES (2905, 1, 38, 'admin');
INSERT INTO `sys_role_menus` VALUES (2906, 1, 39, 'admin');
INSERT INTO `sys_role_menus` VALUES (2907, 1, 40, 'admin');
INSERT INTO `sys_role_menus` VALUES (2908, 1, 41, 'admin');
INSERT INTO `sys_role_menus` VALUES (2909, 1, 42, 'admin');
INSERT INTO `sys_role_menus` VALUES (2910, 1, 43, 'admin');
INSERT INTO `sys_role_menus` VALUES (2911, 1, 44, 'admin');
INSERT INTO `sys_role_menus` VALUES (2912, 1, 45, 'admin');
INSERT INTO `sys_role_menus` VALUES (2913, 1, 46, 'admin');
INSERT INTO `sys_role_menus` VALUES (2914, 1, 47, 'admin');
INSERT INTO `sys_role_menus` VALUES (2915, 1, 48, 'admin');
INSERT INTO `sys_role_menus` VALUES (2916, 1, 49, 'admin');
INSERT INTO `sys_role_menus` VALUES (2917, 1, 50, 'admin');
INSERT INTO `sys_role_menus` VALUES (2918, 1, 51, 'admin');
INSERT INTO `sys_role_menus` VALUES (2919, 1, 52, 'admin');
INSERT INTO `sys_role_menus` VALUES (2920, 1, 53, 'admin');
INSERT INTO `sys_role_menus` VALUES (2921, 1, 54, 'admin');
INSERT INTO `sys_role_menus` VALUES (2922, 1, 55, 'admin');
INSERT INTO `sys_role_menus` VALUES (2923, 1, 56, 'admin');
INSERT INTO `sys_role_menus` VALUES (2924, 1, 57, 'admin');
INSERT INTO `sys_role_menus` VALUES (2925, 1, 58, 'admin');
INSERT INTO `sys_role_menus` VALUES (2926, 1, 59, 'admin');
INSERT INTO `sys_role_menus` VALUES (2927, 1, 60, 'admin');
INSERT INTO `sys_role_menus` VALUES (2928, 1, 61, 'admin');
INSERT INTO `sys_role_menus` VALUES (2929, 1, 62, 'admin');
INSERT INTO `sys_role_menus` VALUES (2930, 1, 63, 'admin');
INSERT INTO `sys_role_menus` VALUES (2931, 1, 64, 'admin');
INSERT INTO `sys_role_menus` VALUES (2932, 1, 65, 'admin');
INSERT INTO `sys_role_menus` VALUES (2933, 1, 66, 'admin');
INSERT INTO `sys_role_menus` VALUES (2934, 1, 67, 'admin');
INSERT INTO `sys_role_menus` VALUES (2935, 1, 68, 'admin');
INSERT INTO `sys_role_menus` VALUES (2936, 1, 69, 'admin');
INSERT INTO `sys_role_menus` VALUES (2937, 1, 70, 'admin');
INSERT INTO `sys_role_menus` VALUES (2938, 1, 71, 'admin');
INSERT INTO `sys_role_menus` VALUES (2939, 1, 72, 'admin');
INSERT INTO `sys_role_menus` VALUES (2940, 1, 73, 'admin');
INSERT INTO `sys_role_menus` VALUES (2941, 1, 74, 'admin');
INSERT INTO `sys_role_menus` VALUES (2942, 1, 75, 'admin');
INSERT INTO `sys_role_menus` VALUES (2943, 1, 76, 'admin');
INSERT INTO `sys_role_menus` VALUES (2944, 1, 77, 'admin');
INSERT INTO `sys_role_menus` VALUES (2945, 1, 78, 'admin');
INSERT INTO `sys_role_menus` VALUES (2946, 1, 79, 'admin');
INSERT INTO `sys_role_menus` VALUES (2947, 1, 80, 'admin');
INSERT INTO `sys_role_menus` VALUES (2948, 1, 81, 'admin');
INSERT INTO `sys_role_menus` VALUES (2949, 1, 83, 'admin');
INSERT INTO `sys_role_menus` VALUES (2950, 1, 84, 'admin');
INSERT INTO `sys_role_menus` VALUES (2951, 1, 85, 'admin');
INSERT INTO `sys_role_menus` VALUES (2952, 1, 86, 'admin');
INSERT INTO `sys_role_menus` VALUES (2953, 1, 87, 'admin');
INSERT INTO `sys_role_menus` VALUES (2954, 1, 88, 'admin');
INSERT INTO `sys_role_menus` VALUES (2955, 1, 89, 'admin');
INSERT INTO `sys_role_menus` VALUES (2956, 1, 90, 'admin');
INSERT INTO `sys_role_menus` VALUES (2957, 2, 1, 'manage');
INSERT INTO `sys_role_menus` VALUES (2958, 2, 3, 'manage');
INSERT INTO `sys_role_menus` VALUES (2959, 2, 4, 'manage');
INSERT INTO `sys_role_menus` VALUES (2960, 2, 5, 'manage');
INSERT INTO `sys_role_menus` VALUES (2961, 2, 6, 'manage');
INSERT INTO `sys_role_menus` VALUES (2962, 2, 7, 'manage');
INSERT INTO `sys_role_menus` VALUES (2963, 2, 8, 'manage');
INSERT INTO `sys_role_menus` VALUES (2964, 2, 9, 'manage');
INSERT INTO `sys_role_menus` VALUES (2965, 2, 10, 'manage');
INSERT INTO `sys_role_menus` VALUES (2966, 2, 11, 'manage');
INSERT INTO `sys_role_menus` VALUES (2967, 2, 12, 'manage');
INSERT INTO `sys_role_menus` VALUES (2968, 2, 13, 'manage');
INSERT INTO `sys_role_menus` VALUES (2969, 2, 14, 'manage');
INSERT INTO `sys_role_menus` VALUES (2970, 2, 15, 'manage');
INSERT INTO `sys_role_menus` VALUES (2971, 2, 16, 'manage');
INSERT INTO `sys_role_menus` VALUES (2972, 2, 17, 'manage');
INSERT INTO `sys_role_menus` VALUES (2973, 2, 18, 'manage');
INSERT INTO `sys_role_menus` VALUES (2974, 2, 19, 'manage');
INSERT INTO `sys_role_menus` VALUES (2975, 2, 20, 'manage');
INSERT INTO `sys_role_menus` VALUES (2976, 2, 21, 'manage');
INSERT INTO `sys_role_menus` VALUES (2977, 2, 22, 'manage');
INSERT INTO `sys_role_menus` VALUES (2978, 2, 23, 'manage');
INSERT INTO `sys_role_menus` VALUES (2979, 2, 25, 'manage');
INSERT INTO `sys_role_menus` VALUES (2980, 2, 26, 'manage');
INSERT INTO `sys_role_menus` VALUES (2981, 2, 28, 'manage');
INSERT INTO `sys_role_menus` VALUES (2982, 2, 29, 'manage');
INSERT INTO `sys_role_menus` VALUES (2983, 2, 30, 'manage');
INSERT INTO `sys_role_menus` VALUES (2984, 2, 31, 'manage');
INSERT INTO `sys_role_menus` VALUES (2985, 2, 32, 'manage');
INSERT INTO `sys_role_menus` VALUES (2986, 2, 33, 'manage');
INSERT INTO `sys_role_menus` VALUES (2987, 2, 34, 'manage');
INSERT INTO `sys_role_menus` VALUES (2988, 2, 35, 'manage');
INSERT INTO `sys_role_menus` VALUES (2989, 2, 36, 'manage');
INSERT INTO `sys_role_menus` VALUES (2990, 2, 37, 'manage');
INSERT INTO `sys_role_menus` VALUES (2991, 2, 38, 'manage');
INSERT INTO `sys_role_menus` VALUES (2992, 2, 39, 'manage');
INSERT INTO `sys_role_menus` VALUES (2993, 2, 40, 'manage');
INSERT INTO `sys_role_menus` VALUES (2994, 2, 41, 'manage');
INSERT INTO `sys_role_menus` VALUES (2995, 2, 42, 'manage');
INSERT INTO `sys_role_menus` VALUES (2996, 2, 43, 'manage');
INSERT INTO `sys_role_menus` VALUES (2997, 2, 44, 'manage');
INSERT INTO `sys_role_menus` VALUES (2998, 2, 45, 'manage');
INSERT INTO `sys_role_menus` VALUES (2999, 2, 46, 'manage');
INSERT INTO `sys_role_menus` VALUES (3000, 2, 47, 'manage');
INSERT INTO `sys_role_menus` VALUES (3001, 2, 48, 'manage');
INSERT INTO `sys_role_menus` VALUES (3002, 2, 49, 'manage');
INSERT INTO `sys_role_menus` VALUES (3003, 2, 50, 'manage');
INSERT INTO `sys_role_menus` VALUES (3004, 2, 51, 'manage');
INSERT INTO `sys_role_menus` VALUES (3005, 2, 52, 'manage');
INSERT INTO `sys_role_menus` VALUES (3006, 2, 53, 'manage');
INSERT INTO `sys_role_menus` VALUES (3007, 2, 54, 'manage');
INSERT INTO `sys_role_menus` VALUES (3008, 2, 55, 'manage');
INSERT INTO `sys_role_menus` VALUES (3009, 2, 56, 'manage');
INSERT INTO `sys_role_menus` VALUES (3010, 2, 57, 'manage');
INSERT INTO `sys_role_menus` VALUES (3011, 2, 58, 'manage');
INSERT INTO `sys_role_menus` VALUES (3012, 2, 59, 'manage');
INSERT INTO `sys_role_menus` VALUES (3013, 2, 60, 'manage');
INSERT INTO `sys_role_menus` VALUES (3014, 2, 61, 'manage');
INSERT INTO `sys_role_menus` VALUES (3015, 2, 62, 'manage');
INSERT INTO `sys_role_menus` VALUES (3016, 2, 63, 'manage');
INSERT INTO `sys_role_menus` VALUES (3017, 2, 64, 'manage');
INSERT INTO `sys_role_menus` VALUES (3018, 2, 65, 'manage');
INSERT INTO `sys_role_menus` VALUES (3019, 2, 66, 'manage');
INSERT INTO `sys_role_menus` VALUES (3020, 2, 67, 'manage');
INSERT INTO `sys_role_menus` VALUES (3021, 2, 68, 'manage');
INSERT INTO `sys_role_menus` VALUES (3022, 2, 69, 'manage');
INSERT INTO `sys_role_menus` VALUES (3023, 2, 70, 'manage');
INSERT INTO `sys_role_menus` VALUES (3024, 2, 71, 'manage');
INSERT INTO `sys_role_menus` VALUES (3025, 2, 72, 'manage');
INSERT INTO `sys_role_menus` VALUES (3026, 2, 73, 'manage');
INSERT INTO `sys_role_menus` VALUES (3027, 2, 74, 'manage');
INSERT INTO `sys_role_menus` VALUES (3028, 2, 75, 'manage');
INSERT INTO `sys_role_menus` VALUES (3029, 2, 76, 'manage');
INSERT INTO `sys_role_menus` VALUES (3030, 2, 77, 'manage');
INSERT INTO `sys_role_menus` VALUES (3031, 2, 78, 'manage');
INSERT INTO `sys_role_menus` VALUES (3032, 2, 79, 'manage');
INSERT INTO `sys_role_menus` VALUES (3033, 2, 80, 'manage');
INSERT INTO `sys_role_menus` VALUES (3034, 2, 81, 'manage');
INSERT INTO `sys_role_menus` VALUES (3035, 2, 83, 'manage');
INSERT INTO `sys_role_menus` VALUES (3036, 2, 84, 'manage');
INSERT INTO `sys_role_menus` VALUES (3037, 2, 85, 'manage');
INSERT INTO `sys_role_menus` VALUES (3038, 2, 86, 'manage');
INSERT INTO `sys_role_menus` VALUES (3039, 2, 87, 'manage');
INSERT INTO `sys_role_menus` VALUES (3040, 2, 88, 'manage');
INSERT INTO `sys_role_menus` VALUES (3041, 2, 89, 'manage');
INSERT INTO `sys_role_menus` VALUES (3042, 2, 90, 'manage');

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `tenant_id` int(11) NULL DEFAULT NULL,
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色代码',
  `data_scope` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `role_sort` int(4) NULL DEFAULT NULL COMMENT '角色排序',
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
INSERT INTO `sys_roles` VALUES (1, '超管理员', 1, '0', 'admin', '2', 1, 'admin', 'panda', '超级管理', '2021-12-02 16:03:26', '2022-07-18 10:57:14', NULL);
INSERT INTO `sys_roles` VALUES (2, '管理员', 1, '0', 'manage', '', 2, 'panda', 'panda', '', '2021-12-19 16:06:20', '2022-07-19 14:03:34', NULL);

-- ----------------------------
-- Table structure for sys_tenants
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenants`;
CREATE TABLE `sys_tenants`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  `delete_time` datetime NULL DEFAULT NULL,
  `tenant_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '租户名',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `expire_time` datetime NULL DEFAULT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_tenants
-- ----------------------------
INSERT INTO `sys_tenants` VALUES (1, '2022-07-16 18:28:33', '2022-07-16 18:28:33', NULL, '熊猫科技', '鹅鹅鹅', '2099-07-16 00:00:00');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `tenant_id` int(11) NULL DEFAULT NULL,
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
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, 'pandax', '13818888888', 1, 1, NULL, '', '0', '1@qq.com', 2, 4, 'panda', '1', NULL, '0', '2021-12-03 09:46:55', '2022-02-09 13:28:49', NULL, 'panda', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '1', '1,4');
INSERT INTO `sys_users` VALUES (3, '测试用户', '18435234356', 1, 2, '', '', '0', '342@163.com', 3, 1, 'test', '', '', '0', '2021-12-06 15:16:53', '2022-05-10 19:19:25', NULL, 'test', '$2a$10$4cHTracxWJLdhMmazvbm1urKyD3v5N2AYxAFtNYh6juU39kgae73e', '2', '1,4');
INSERT INTO `sys_users` VALUES (4, 'panda', '18353366912', 1, 2, '', '', '0', '2417920382@qq.com', 2, 4, 'panda', '', '', '0', '2021-12-19 15:58:09', '2021-12-19 16:06:54', NULL, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '2', '4,1');
INSERT INTO `sys_users` VALUES (5, 'tenant', '', 1, 1, '', '', '0', '', 3, 1, 'panda', '1', '', '0', '2021-12-03 09:46:55', '2022-02-09 13:28:49', NULL, 'tenant', '$2a$10$ycRsRdsrNQInLB2Ib0maOetsWZ0kFctmF6ytAErWTjOx5cWdeJMcK', '1', '1,4');

SET FOREIGN_KEY_CHECKS = 1;
