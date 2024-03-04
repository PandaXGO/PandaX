/*
 Navicat Premium Data Transfer

 Source Server         : localhost_mysql-8.0
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : pandax_iot

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 24/10/2023 16:19:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS pandax_iot;

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
  `id` int(0) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10604 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'test', '/system/api/list', 'GET', '', '', '', 8679);
INSERT INTO `casbin_rule` VALUES ('p', 'test', '/system/api/all', 'GET', '', '', '', 8680);
INSERT INTO `casbin_rule` VALUES ('p', 'test', '/system/api/getPolicyPathByRoleId', 'GET', '', '', '', 8681);
INSERT INTO `casbin_rule` VALUES ('p', 'test', '/system/api/:id', 'GET', '', '', '', 8682);
INSERT INTO `casbin_rule` VALUES ('p', 'test', '/system/api', 'POST', '', '', '', 8683);
INSERT INTO `casbin_rule` VALUES ('p', 'test', '/system/api', 'PUT', '', '', '', 8684);
INSERT INTO `casbin_rule` VALUES ('p', 'test', '/system/api/:id', 'DELETE', '', '', '', 8685);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/api/list', 'GET', '', '', '', 10162);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/api/all', 'GET', '', '', '', 10163);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/api/getPolicyPathByRoleId', 'GET', '', '', '', 10164);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/api/:id', 'GET', '', '', '', 10165);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/config/list', 'GET', '', '', '', 10166);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/config/configKey', 'GET', '', '', '', 10167);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/config/:configId', 'GET', '', '', '', 10168);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/group/list', 'GET', '', '', '', 10169);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/group/list/all', 'GET', '', '', '', 10170);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/group/list/tree', 'GET', '', '', '', 10171);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/group/:id', 'GET', '', '', '', 10172);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/list', 'GET', '', '', '', 10173);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/:id', 'GET', '', '', '', 10174);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/group/list/tree/label', 'GET', '', '', '', 10175);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/list/all', 'GET', '', '', '', 10176);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/:id/status', 'GET', '', '', '', 10177);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/alarm/list', 'GET', '', '', '', 10178);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/cmd/list', 'GET', '', '', '', 10179);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/:id/property/history', 'GET', '', '', '', 10180);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/panel', 'GET', '', '', '', 10181);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/alarm/panel', 'GET', '', '', '', 10182);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/type/list', 'GET', '', '', '', 10183);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/type/:dictId', 'GET', '', '', '', 10184);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/data/list', 'GET', '', '', '', 10185);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/data/type', 'GET', '', '', '', 10186);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/dict/data/:dictCode', 'GET', '', '', '', 10187);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/develop/code/table/db/list', 'GET', '', '', '', 10188);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/develop/code/table/list', 'GET', '', '', '', 10189);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/develop/code/table/info/:tableId', 'GET', '', '', '', 10190);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/develop/code/table/info/tableName', 'GET', '', '', '', 10191);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/develop/code/table/tableTree', 'GET', '', '', '', 10192);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/develop/code/gen/preview/:tableId', 'GET', '', '', '', 10193);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/job/list', 'GET', '', '', '', 10194);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/job/:jobId', 'GET', '', '', '', 10195);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/log/logLogin/list', 'GET', '', '', '', 10196);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/log/logOper/list', 'GET', '', '', '', 10197);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/menuTreeSelect', 'GET', '', '', '', 10198);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/menuRole', 'GET', '', '', '', 10199);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', '', 10200);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/menuPaths', 'GET', '', '', '', 10201);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/list', 'GET', '', '', '', 10202);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/menu/:menuId', 'GET', '', '', '', 10203);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/notice/list', 'GET', '', '', '', 10204);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/organization/list', 'GET', '', '', '', 10205);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/organization/:organizationId', 'GET', '', '', '', 10206);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/organization/roleOrganizationTreeSelect/:roleId', 'GET', '', '', '', 10207);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/organization/organizationTree', 'GET', '', '', '', 10208);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/ota/list', 'GET', '', '', '', 10209);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/ota/:id', 'GET', '', '', '', 10210);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/post/list', 'GET', '', '', '', 10211);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/post/:postId', 'GET', '', '', '', 10212);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/category/list', 'GET', '', '', '', 10213);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/category/list/all', 'GET', '', '', '', 10214);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/category/list/tree', 'GET', '', '', '', 10215);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/category/:id', 'GET', '', '', '', 10216);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/:id', 'GET', '', '', '', 10217);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/list', 'GET', '', '', '', 10218);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/category/list/tree/label', 'GET', '', '', '', 10219);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/list/all', 'GET', '', '', '', 10220);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/product/:id/tsl', 'GET', '', '', '', 10221);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/role/list', 'GET', '', '', '', 10222);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/role/:roleId', 'GET', '', '', '', 10223);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/rule/chain/list', 'GET', '', '', '', 10224);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/rule/chain/:ruleId', 'GET', '', '', '', 10225);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/rule/chain/list/label', 'GET', '', '', '', 10226);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/rule/chain/log/list', 'GET', '', '', '', 10227);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/visual/screen/:screenId', 'GET', '', '', '', 10228);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/visual/screen/list', 'GET', '', '', '', 10229);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/visual/screen/group/list', 'GET', '', '', '', 10230);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/visual/screen/group/list/tree', 'GET', '', '', '', 10231);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/visual/screen/group/list/all', 'GET', '', '', '', 10232);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/visual/screen/group/:id', 'GET', '', '', '', 10233);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/template/list', 'GET', '', '', '', 10234);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/template/:id', 'GET', '', '', '', 10235);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/device/template/list/all', 'GET', '', '', '', 10236);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/list', 'GET', '', '', '', 10237);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/getById/:userId', 'GET', '', '', '', 10238);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/getInit', 'GET', '', '', '', 10239);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/system/user/getRoPo', 'GET', '', '', '', 10240);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/video/ys/device/list', 'GET', '', '', '', 10241);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/video/ys/:deviceSerial/channel', 'GET', '', '', '', 10242);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/video/ys/:deviceSerial/channel/live', 'GET', '', '', '', 10243);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/video/ys/:deviceSerial/ptz/start', 'GET', '', '', '', 10244);
INSERT INTO `casbin_rule` VALUES ('p', 'manage', '/video/ys/:deviceSerial/ptz/stop', 'GET', '', '', '', 10245);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/list', 'GET', '', '', '', 10425);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/all', 'GET', '', '', '', 10426);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/getPolicyPathByRoleId', 'GET', '', '', '', 10427);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/:id', 'GET', '', '', '', 10428);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api', 'POST', '', '', '', 10429);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api', 'PUT', '', '', '', 10430);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/api/:id', 'DELETE', '', '', '', 10431);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/list', 'GET', '', '', '', 10432);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/configKey', 'GET', '', '', '', 10433);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/:configId', 'GET', '', '', '', 10434);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'POST', '', '', '', 10435);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'PUT', '', '', '', 10436);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/:configId', 'DELETE', '', '', '', 10437);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/group/list', 'GET', '', '', '', 10438);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/group/list/all', 'GET', '', '', '', 10439);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/group/list/tree', 'GET', '', '', '', 10440);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/group/:id', 'GET', '', '', '', 10441);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/group', 'POST', '', '', '', 10442);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/group', 'PUT', '', '', '', 10443);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/group/:id', 'DELETE', '', '', '', 10444);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/list', 'GET', '', '', '', 10445);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/:id', 'GET', '', '', '', 10446);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device', 'POST', '', '', '', 10447);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/:id', 'DELETE', '', '', '', 10448);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device', 'PUT', '', '', '', 10449);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/group/list/tree/label', 'GET', '', '', '', 10450);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/list/all', 'GET', '', '', '', 10451);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/:id/status', 'GET', '', '', '', 10452);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/alarm/list', 'GET', '', '', '', 10453);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/alarm', 'PUT', '', '', '', 10454);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/alarm/:id', 'DELETE', '', '', '', 10455);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/cmd/list', 'GET', '', '', '', 10456);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/cmd', 'POST', '', '', '', 10457);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/cmd/:id', 'DELETE', '', '', '', 10458);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/:id/attribute/down', 'GET', '', '', '', 10459);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/:id/property/history', 'GET', '', '', '', 10460);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/:id/allot/org', 'GET', '', '', '', 10461);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/panel', 'GET', '', '', '', 10462);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/alarm/panel', 'GET', '', '', '', 10463);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/list', 'GET', '', '', '', 10464);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/:dictId', 'GET', '', '', '', 10465);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'POST', '', '', '', 10466);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'PUT', '', '', '', 10467);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/:dictId', 'DELETE', '', '', '', 10468);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/export', 'GET', '', '', '', 10469);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/list', 'GET', '', '', '', 10470);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/type', 'GET', '', '', '', 10471);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/:dictCode', 'GET', '', '', '', 10472);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'POST', '', '', '', 10473);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'PUT', '', '', '', 10474);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/:dictCode', 'DELETE', '', '', '', 10475);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/table/db/list', 'GET', '', '', '', 10476);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/table/list', 'GET', '', '', '', 10477);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/table/info/:tableId', 'GET', '', '', '', 10478);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/table/info/tableName', 'GET', '', '', '', 10479);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/table/tableTree', 'GET', '', '', '', 10480);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/table', 'POST', '', '', '', 10481);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/table', 'PUT', '', '', '', 10482);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/table/:tableId', 'DELETE', '', '', '', 10483);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/gen/preview/:tableId', 'GET', '', '', '', 10484);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/gen/code/:tableId', 'GET', '', '', '', 10485);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/develop/code/gen/configure/:tableId', 'GET', '', '', '', 10486);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/list', 'GET', '', '', '', 10487);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job', 'POST', '', '', '', 10488);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job', 'PUT', '', '', '', 10489);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/:jobId', 'GET', '', '', '', 10490);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/:jobId', 'DELETE', '', '', '', 10491);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/stop/:jobId', 'GET', '', '', '', 10492);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/start/:jobId', 'GET', '', '', '', 10493);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/log/list', 'GET', '', '', '', 10494);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/log/all', 'DELETE', '', '', '', 10495);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/log/:logId', 'DELETE', '', '', '', 10496);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/job/changeStatus', 'PUT', '', '', '', 10497);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logLogin/list', 'GET', '', '', '', 10498);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logLogin/:infoId', 'DELETE', '', '', '', 10499);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logLogin/all', 'DELETE', '', '', '', 10500);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logOper/list', 'GET', '', '', '', 10501);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logOper/:operId', 'DELETE', '', '', '', 10502);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/log/logOper/all', 'DELETE', '', '', '', 10503);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuTreeSelect', 'GET', '', '', '', 10504);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuRole', 'GET', '', '', '', 10505);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/roleMenuTreeSelect/:roleId', 'GET', '', '', '', 10506);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/menuPaths', 'GET', '', '', '', 10507);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/list', 'GET', '', '', '', 10508);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/:menuId', 'GET', '', '', '', 10509);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'POST', '', '', '', 10510);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'PUT', '', '', '', 10511);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/:menuId', 'DELETE', '', '', '', 10512);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/notice/list', 'GET', '', '', '', 10513);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/notice', 'POST', '', '', '', 10514);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/notice', 'PUT', '', '', '', 10515);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/notice/:noticeId', 'DELETE', '', '', '', 10516);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/organization/list', 'GET', '', '', '', 10517);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/organization/:organizationId', 'GET', '', '', '', 10518);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/organization/roleOrganizationTreeSelect/:roleId', 'GET', '', '', '', 10519);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/organization/organizationTree', 'GET', '', '', '', 10520);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/organization', 'POST', '', '', '', 10521);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/organization', 'PUT', '', '', '', 10522);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/organization/:organizationId', 'DELETE', '', '', '', 10523);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/ota/list', 'GET', '', '', '', 10524);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/ota', 'POST', '', '', '', 10525);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/ota', 'PUT', '', '', '', 10526);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/ota/:id', 'DELETE', '', '', '', 10527);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/ota/:id', 'GET', '', '', '', 10528);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/list', 'GET', '', '', '', 10529);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/:postId', 'GET', '', '', '', 10530);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'POST', '', '', '', 10531);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'PUT', '', '', '', 10532);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/:postId', 'DELETE', '', '', '', 10533);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/category/list', 'GET', '', '', '', 10534);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/category/list/all', 'GET', '', '', '', 10535);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/category/list/tree', 'GET', '', '', '', 10536);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/category/:id', 'GET', '', '', '', 10537);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/category', 'POST', '', '', '', 10538);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/category', 'PUT', '', '', '', 10539);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/category/:id', 'DELETE', '', '', '', 10540);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/:id', 'DELETE', '', '', '', 10541);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/:id', 'GET', '', '', '', 10542);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product', 'PUT', '', '', '', 10543);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/list', 'GET', '', '', '', 10544);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product', 'POST', '', '', '', 10545);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/category/list/tree/label', 'GET', '', '', '', 10546);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/list/all', 'GET', '', '', '', 10547);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/product/:id/tsl', 'GET', '', '', '', 10548);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/list', 'GET', '', '', '', 10549);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/:roleId', 'GET', '', '', '', 10550);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'POST', '', '', '', 10551);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'PUT', '', '', '', 10552);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/:roleId', 'DELETE', '', '', '', 10553);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/changeStatus', 'PUT', '', '', '', 10554);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/dataScope', 'PUT', '', '', '', 10555);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/export', 'GET', '', '', '', 10556);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain/changeRoot', 'PUT', '', '', '', 10557);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain/list', 'GET', '', '', '', 10558);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain/:ruleId', 'DELETE', '', '', '', 10559);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain', 'PUT', '', '', '', 10560);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain', 'POST', '', '', '', 10561);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain/:ruleId', 'GET', '', '', '', 10562);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain/list/label', 'GET', '', '', '', 10563);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain/clone/:ruleId', 'POST', '', '', '', 10564);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain/log/list', 'GET', '', '', '', 10565);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/rule/chain/log/delete', 'GET', '', '', '', 10566);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen', 'PUT', '', '', '', 10567);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/:screenId', 'GET', '', '', '', 10568);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/list', 'GET', '', '', '', 10569);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/:screenId', 'DELETE', '', '', '', 10570);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen', 'POST', '', '', '', 10571);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/changeStatus', 'PUT', '', '', '', 10572);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/group/list', 'GET', '', '', '', 10573);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/group/list/tree', 'GET', '', '', '', 10574);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/group/list/all', 'GET', '', '', '', 10575);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/group/:id', 'GET', '', '', '', 10576);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/group', 'POST', '', '', '', 10577);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/group', 'PUT', '', '', '', 10578);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/visual/screen/group/:id', 'DELETE', '', '', '', 10579);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/template/list', 'GET', '', '', '', 10580);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/template', 'PUT', '', '', '', 10581);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/template/:id', 'GET', '', '', '', 10582);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/template/:id', 'DELETE', '', '', '', 10583);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/template', 'POST', '', '', '', 10584);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/device/template/list/all', 'GET', '', '', '', 10585);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/upload/up/oss', 'POST', '', '', '', 10586);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/upload/up', 'POST', '', '', '', 10587);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/list', 'GET', '', '', '', 10588);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/changeStatus', 'PUT', '', '', '', 10589);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/:userId', 'DELETE', '', '', '', 10590);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/avatar', 'POST', '', '', '', 10591);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/pwd', 'PUT', '', '', '', 10592);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getById/:userId', 'GET', '', '', '', 10593);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getInit', 'GET', '', '', '', 10594);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/getRoPo', 'GET', '', '', '', 10595);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'POST', '', '', '', 10596);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'PUT', '', '', '', 10597);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/export', 'GET', '', '', '', 10598);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/video/ys/device/list', 'GET', '', '', '', 10599);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/video/ys/:deviceSerial/channel', 'GET', '', '', '', 10600);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/video/ys/:deviceSerial/channel/live', 'GET', '', '', '', 10601);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/video/ys/:deviceSerial/ptz/start', 'GET', '', '', '', 10602);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/video/ys/:deviceSerial/ptz/stop', 'GET', '', '', '', 10603);

-- ----------------------------
-- Table structure for demo_new_trend_of_diagnosis
-- ----------------------------
DROP TABLE IF EXISTS `demo_new_trend_of_diagnosis`;
CREATE TABLE `demo_new_trend_of_diagnosis`  (
  `date` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '日期',
  `new_diagnosis` bigint(0) NULL DEFAULT NULL COMMENT '新增确诊',
  `current_diagnosis` bigint(0) NULL DEFAULT NULL COMMENT '现有确诊'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_new_trend_of_diagnosis
-- ----------------------------
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-10', 33, 505);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-11', 28, 506);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-12', 32, 512);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-13', 35, 523);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-14', 49, 542);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-15', 206, 727);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-16', 236, 935);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-17', 358, 1262);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-18', 258, 1497);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-19', 286, 1759);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-20', 317, 2097);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-21', 325, 2365);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-22', 743, 3098);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-23', 480, 3561);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-24', 612, 4143);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-25', 554, 4675);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-26', 655, 5036);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-27', 677, 5948);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-28', 570, 6480);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-29', 503, 6951);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-30', 381, 7303);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-5-31', 378, 7652);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-6-1', 362, 7983);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-6-2', 571, 8535);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-6-3', 610, 9110);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-6-4', 497, 9674);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-6-5', 541, 10049);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-6-6', 368, 10372);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-6-7', 233, 10552);
INSERT INTO `demo_new_trend_of_diagnosis` VALUES ('2021-6-8', 232, 10740);

-- ----------------------------
-- Table structure for dev_gen_table_columns
-- ----------------------------
DROP TABLE IF EXISTS `dev_gen_table_columns`;
CREATE TABLE `dev_gen_table_columns`  (
  `column_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `table_id` bigint(0) NULL DEFAULT NULL,
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
  `sort` bigint(0) NULL DEFAULT NULL,
  `link_table_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_table_class` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_table_package` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_label_id` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_label_name` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `org_id` int(0) NULL DEFAULT NULL COMMENT '机构ID',
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 138 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of dev_gen_table_columns
-- ----------------------------

-- ----------------------------
-- Table structure for dev_gen_tables
-- ----------------------------
DROP TABLE IF EXISTS `dev_gen_tables`;
CREATE TABLE `dev_gen_tables`  (
  `table_id` bigint(0) NOT NULL AUTO_INCREMENT,
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
  `pk_go_type` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pk_json_field` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  `org_id` int(0) NULL DEFAULT NULL COMMENT '机构ID',
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of dev_gen_tables
-- ----------------------------

-- ----------------------------
-- Table structure for device_alarms
-- ----------------------------
DROP TABLE IF EXISTS `device_alarms`;
CREATE TABLE `device_alarms`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '告警名称',
  `org_id` int(0) NULL DEFAULT NULL COMMENT '机构ID',
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '告警类型',
  `level` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '告警级别',
  `state` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '告警状态',
  `details` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '详情',
  `time` datetime(0) NULL DEFAULT NULL COMMENT '告警时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of device_alarms
-- ----------------------------


-- ----------------------------
-- Table structure for device_cmd_log
-- ----------------------------
DROP TABLE IF EXISTS `device_cmd_log`;
CREATE TABLE `device_cmd_log`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cmd_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cmd_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `state` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `response_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `request_time` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `response_time` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `mode` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of device_cmd_log
-- ----------------------------


-- ----------------------------
-- Table structure for device_groups
-- ----------------------------
DROP TABLE IF EXISTS `device_groups`;
CREATE TABLE `device_groups`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` int(0) NULL DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备分组名称',
  `pid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '上级设备分组类型',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备分组路径',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备分组说明',
  `sort` int(0) NULL DEFAULT NULL COMMENT '排序',
  `ext` json NULL COMMENT '扩展',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of device_groups
-- ----------------------------
INSERT INTO `device_groups` VALUES ('5h2eEVcqLw', 'panda', 2, '2023-06-30 08:52:37', '2023-06-30 08:52:37', '一层', 'eiAx7ZgWKg', '/0/eiAx7ZgWKg/5h2eEVcqLw', '', 1, 'null', '0');
INSERT INTO `device_groups` VALUES ('9uOQ1Ku0PQ', 'panda', 2, '2023-10-14 17:43:25', '2023-10-17 10:10:07', '默认分组', '0', '/0/9uOQ1Ku0PQ', '未定义分组的设备都在这里面', 1, 'null', '0');
INSERT INTO `device_groups` VALUES ('eiAx7ZgWKg', 'panda', 2, '2023-06-30 08:52:16', '2023-06-30 08:53:47', '1号楼', '0', '/0/eiAx7ZgWKg', '1号楼，位于园区东南角，安保人：张三，电话：11111', 1, 'null', '0');

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` int(0) NULL DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备名称',
  `token` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备token',
  `alias` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备别名',
  `pid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品Id',
  `gid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '分组Id',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品说明',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '状态',
  `ota_version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '固件版本',
  `ext` json NULL COMMENT '拓展',
  `parent_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '父Id，子设备时，父设备为网关',
  `device_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备类型',
  `link_status` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '连接状态',
  `last_time` datetime(0) NULL DEFAULT NULL COMMENT '最后一次在线时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_devices_product`(`pid`) USING BTREE,
  INDEX `fk_devices_device_group`(`gid`) USING BTREE,
  CONSTRAINT `fk_devices_device_group` FOREIGN KEY (`gid`) REFERENCES `device_groups` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_devices_product` FOREIGN KEY (`pid`) REFERENCES `products` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of devices
-- ----------------------------
INSERT INTO `devices` VALUES ('2HbCMj8WOQ', 'panda', 2, '2023-10-12 19:22:25', '2023-10-12 19:22:25', 'sparksiiot', 'ZGRlMTE2NmEtYWY5MS0zZDRmLTlhYTktZWE1Njg5Yjk0MTlm', '星原网关', 'uqNNwYJ5rw', '5h2eEVcqLw', '', '0', '', 'null', '', 'gateway', 'online', '2023-10-14 21:32:48');
INSERT INTO `devices` VALUES ('68zSC94dFQ', 'panda', 2, '2023-07-31 14:23:13', '2023-07-31 14:23:13', 'ctr453', '', '智能控制器453', 'p_bf52caf91f7cdd2abb52eaaf', 'eiAx7ZgWKg', '', '0', '', 'null', 'rC82hwE6iw', 'gatewayS', 'offline', '2023-09-28 10:08:07');
INSERT INTO `devices` VALUES ('9GOIPOI6GQ', 'panda', 2, '2023-07-26 22:23:16', '2023-09-07 11:35:22', 'ws432', 'YWRlMTA0MmYtMzc2MS0zZTljLThjNjAtMzNhMzg4ZjdkOGQ3', '温湿度器', 'p_3ba460634520cf4590dc90e5', 'eiAx7ZgWKg', '设备说明1', '0', '', '{\"location\": {\"lat\": 37.037581, \"lng\": 118.18431, \"address\": \"山东省淄博市张店区傅家镇淄博市植物园志愿者阅览室\", \"position\": [118.027698, 36.791573]}}', '', 'direct', 'online', '2023-10-14 12:27:55');
INSERT INTO `devices` VALUES ('k2opRSpr-g', 'panda', 7, '2023-10-14 22:21:02', '2023-10-17 10:09:05', 'testsub', '', '测试子设备', 'I_HlHDdh_Q', '9uOQ1Ku0PQ', '', '0', '', 'null', 'rC82hwE6iw', 'gatewayS', 'online', '2023-10-17 16:10:48');
INSERT INTO `devices` VALUES ('l7HF7UZCEA', 'panda', 2, '2023-09-28 09:22:41', '2023-09-28 09:22:41', 'zhilian01', 'MTZlZDM3OGItODdiOS0zZDIwLWJmZjQtMWY3ODM3YzRhN2Ji', '直连设备', 'p_3ba460634520cf4590dc90e5', '5h2eEVcqLw', '', '0', '', 'null', '', 'direct', 'inactive', '2023-09-28 09:22:41');
INSERT INTO `devices` VALUES ('lCtIzLLdIQ', 'panda', 2, '2023-09-27 11:47:47', '2023-09-27 11:47:47', 'TestTcp', 'OTYwNTE3ODUtYTFhMy0zOTIwLWIwZmItYzc3OWVkZWZjOTUw', 'TCP透传', 'mSOWuiA97g', '5h2eEVcqLw', '', '0', '', 'null', '', 'direct', 'offline', '2023-10-08 13:52:06');
INSERT INTO `devices` VALUES ('qmWqYlY6-w', 'panda', 2, '2023-09-27 15:08:56', '2023-09-27 15:08:56', 'httpde1', 'MDVlY2MyNzYtMzczMS0zN2Y2LTk1MWMtMDMwM2ZjNmQyNjlm', 'HTTP设备测试', 'ek2WUADl6g', '5h2eEVcqLw', '', '0', '', 'null', '', 'direct', 'offline', '2023-10-07 15:03:32');
INSERT INTO `devices` VALUES ('rC82hwE6iw', 'panda', 7, '2023-09-23 14:22:18', '2023-09-16 10:03:12', 'gateway4353', 'ZTg0ZDNkZDItOWQ1Mi0zYjM2LTg1NWQtYTI0NmE0NDcyOTM2', '智能网关4353', 'p_cdbb1eccd902018d51fe062e', 'eiAx7ZgWKg', '', '0', '', '{\"location\": {\"label\": \"\", \"address\": \"天津市西青区中北镇现快速处理中心\", \"content\": \"\", \"position\": [117.100495, 39.135469]}}', '', 'gateway', 'offline', '2023-10-17 16:35:39');
INSERT INTO `devices` VALUES ('YbWKD905pQ', 'panda', 2, '2023-10-12 19:23:29', '2023-10-12 19:23:29', 'Panasonic', '', '松下PLC', 'M32969chcw', '5h2eEVcqLw', '', '0', '', 'null', '2HbCMj8WOQ', 'gatewayS', 'online', '2023-10-12 19:53:59');

-- ----------------------------
-- Table structure for job_logs
-- ----------------------------
DROP TABLE IF EXISTS `job_logs`;
CREATE TABLE `job_logs`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务名称',
  `entry_id` int(0) NULL DEFAULT NULL COMMENT '任务id',
  `target_invoke` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '调用方法',
  `log_info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '日志信息',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of job_logs
-- ----------------------------

-- ----------------------------
-- Table structure for jobs
-- ----------------------------
DROP TABLE IF EXISTS `jobs`;
CREATE TABLE `jobs`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` int(0) NULL DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '名称',
  `target_invoke` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '调用目标',
  `target_args` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目标传参',
  `job_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '目标传参 要执行的内容',
  `cron_expression` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'cron表达式',
  `misfire_policy` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '执行策略',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '状态',
  `entry_id` int(0) NULL DEFAULT NULL COMMENT 'job启动时返回的id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jobs
-- ----------------------------
INSERT INTO `jobs` VALUES ('wvz4D6CXSw', 'panda', 2, '2023-08-08 17:29:30', '2023-08-08 17:42:58', 'adsa', 'cronDevice', 'd_1928b99619910dae5a001fa7', '{\"设备下发\":\"asdas\"}', ' 0/10 * * * * ?', '1', '0', 0);

-- ----------------------------
-- Table structure for log_logins
-- ----------------------------
DROP TABLE IF EXISTS `log_logins`;
CREATE TABLE `log_logins`  (
  `info_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `ipaddr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '系统',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '固件',
  `login_time` timestamp(0) NULL DEFAULT NULL COMMENT '登录时间',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新者',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3622 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log_logins
-- ----------------------------

-- ----------------------------
-- Table structure for log_opers
-- ----------------------------
DROP TABLE IF EXISTS `log_opers`;
CREATE TABLE `log_opers`  (
  `oper_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作的模块',
  `business_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '0其它 1新增 2修改 3删除',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `oper_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作人员',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作url',
  `oper_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作IP',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作地点',
  `oper_param` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求参数',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '0=正常,1=异常',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1759 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log_opers
-- ----------------------------

-- ----------------------------
-- Table structure for product_categories
-- ----------------------------
DROP TABLE IF EXISTS `product_categories`;
CREATE TABLE `product_categories`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` int(0) NULL DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品类型名称',
  `pid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '上级产品类型',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品类型路径',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品类型说明',
  `sort` int(0) NULL DEFAULT NULL COMMENT '排序',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product_categories
-- ----------------------------
INSERT INTO `product_categories` VALUES ('-_FMhNN1QA', 'panda', 2, '2023-10-10 14:55:42', '2023-10-10 14:55:43', '温度传感器', '0', '/0/-_FMhNN1QA', '', 3, '0');
INSERT INTO `product_categories` VALUES ('4BqMUN37_g', 'panda', 2, '2023-10-10 14:56:41', '2023-10-10 14:56:42', 'xx温度传感器', '-_FMhNN1QA', '/0/-_FMhNN1QA/4BqMUN37_g', '', 1, '0');
INSERT INTO `product_categories` VALUES ('8f_oaHIo9A', 'panda', 2, '2023-10-12 19:17:09', '2023-10-12 19:17:09', 'SBOX-G系列网关', '0', '/0/8f_oaHIo9A', '', 4, '0');
INSERT INTO `product_categories` VALUES ('KVys13MMsA', 'panda', 2, '2023-10-14 17:36:30', '2023-10-17 10:09:57', '平台默认产品', '0', '/0/KVys13MMsA', '未定义产品的设备，默认继承的产品', 1, '0');
INSERT INTO `product_categories` VALUES ('oKQcjqY8ZQ', 'panda', 2, '2023-10-12 19:17:30', '2023-10-12 19:17:30', 'SBOX-G系列网关', '8f_oaHIo9A', '/0/8f_oaHIo9A/oKQcjqY8ZQ', '', 1, '0');
INSERT INTO `product_categories` VALUES ('pc_8e12a1ec7ba3bffc1337e163', 'panda', 2, '2023-08-09 11:04:37', '2023-08-09 11:04:37', '海康摄像头', 'pc_d31572a0ceaa070f18cb669a', '/0/pc_d31572a0ceaa070f18cb669a/pc_8e12a1ec7ba3bffc1337e163', '', 1, '0');
INSERT INTO `product_categories` VALUES ('pc_d31572a0ceaa070f18cb669a', 'panda', 2, '2023-08-09 11:04:00', '2023-08-09 11:04:00', '视频产品', '0', '/0/pc_d31572a0ceaa070f18cb669a', '', 1, '0');
INSERT INTO `product_categories` VALUES ('pc61058315302171445335c3d5', 'panda', 2, '2023-06-29 17:50:30', '2023-06-29 17:50:31', ' 测试', '0', '/0/pc61058315302171445335c3d5', '', 1, '0');
INSERT INTO `product_categories` VALUES ('pcd2e673d2cd92e860cff5d958', 'panda', 2, '2023-06-29 17:52:18', '2023-06-29 17:52:18', '啊实打实', 'pc61058315302171445335c3d5', '/0/pc61058315302171445335c3d5/pcd2e673d2cd92e860cff5d958', '', 2, '0');

-- ----------------------------
-- Table structure for product_ota
-- ----------------------------
DROP TABLE IF EXISTS `product_ota`;
CREATE TABLE `product_ota`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `pid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品Id',
  `is_latest` tinyint(0) NULL DEFAULT NULL COMMENT '最新版本',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '固件名称',
  `version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '固件版本',
  `url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '下载地址',
  `check` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'md5校验值',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '说明',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product_ota
-- ----------------------------
INSERT INTO `product_ota` VALUES ('FlwNLfdNWg', '2023-10-05 12:09:26', '2023-10-05 12:09:26', 'p_3ba460634520cf4590dc90e5', 0, '测试固件', 'v1.1', '0683c172cdf300720c55ae418a8e83fc_20231005120900.zip', '4bb850e9e7ceb3e9f12fc35b8073eca3', '');

-- ----------------------------
-- Table structure for product_templates
-- ----------------------------
DROP TABLE IF EXISTS `product_templates`;
CREATE TABLE `product_templates`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `pid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品Id',
  `classify` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '模型归类',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '名称',
  `key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标识',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '属性说明',
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '数据类型',
  `define` json NULL COMMENT '数据约束',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product_templates
-- ----------------------------
INSERT INTO `product_templates` VALUES ('5QH8bNo7iA', '2023-10-12 19:21:20', '2023-10-17 15:48:44', 'M32969chcw', 'telemetry', 'X00', 'X00', '', 'bool', '{\"boolDefine\": [{\"key\": \"0\", \"value\": \"正确\"}, {\"key\": \"1\", \"value\": \"失败\"}]}');
INSERT INTO `product_templates` VALUES ('8AcZGXTL5A', '2023-10-12 19:21:30', '2023-10-12 19:21:30', 'M32969chcw', 'telemetry', 'Y00', 'Y00', '', 'bool', '{\"boolDefine\": [{\"key\": \"0\", \"value\": \"\"}, {\"key\": \"1\", \"value\": \"\"}]}');
INSERT INTO `product_templates` VALUES ('iEPLrCC1gA', '2023-09-27 11:35:25', '2023-09-27 11:35:25', 'mSOWuiA97g', 'telemetry', '温度', 'temperature', '', 'float64', '{\"max\": 100, \"min\": 0, \"unit\": \"摄氏度\"}');
INSERT INTO `product_templates` VALUES ('jvvLVElnLg', '2023-09-28 09:26:46', '2023-09-28 09:44:50', 'p_bf52caf91f7cdd2abb52eaaf', 'telemetry', '温度1', 'temperature', '', 'float64', '{}');
INSERT INTO `product_templates` VALUES ('n_yL_KvBOg', '2023-10-17 10:46:14', '2023-10-17 10:46:14', 'I_HlHDdh_Q', 'telemetry', 'test', 'test', '', 'string', 'null');
INSERT INTO `product_templates` VALUES ('R83jjHlUog', '2023-10-12 19:21:45', '2023-10-12 19:21:45', 'M32969chcw', 'telemetry', 'D0', 'D0', '', 'int64', '{}');
INSERT INTO `product_templates` VALUES ('SrVmTxfd5A', '2023-10-17 15:50:22', '2023-10-17 16:11:39', 'I_HlHDdh_Q', 'telemetry', 'bolt', 'bolt', '', 'bool', '{\"boolDefine\": [{\"key\": \"0\", \"value\": \"失败\"}, {\"key\": \"1\", \"value\": \"正确\"}]}');
INSERT INTO `product_templates` VALUES ('tm_14732f0fa234453e328bbd30', '2023-08-01 09:09:51', '2023-08-01 10:26:05', 'p_3ba460634520cf4590dc90e5', 'telemetry', '开关', 'open', '', 'bool', '{\"boolDefine\": [{\"key\": \"0\", \"value\": \"开\"}, {\"key\": \"1\", \"value\": \"关1\"}]}');
INSERT INTO `product_templates` VALUES ('tm_377e0b1cc9812ab11464e2b4', '2023-08-01 09:22:54', '2023-08-01 09:23:09', 'p_3ba460634520cf4590dc90e5', 'telemetry', '测试参数', 'test', '', 'enum', '{\"enumDefine\": [{\"key\": \"0\", \"value\": \"开\"}, {\"key\": \"1\", \"value\": \"关\"}]}');
INSERT INTO `product_templates` VALUES ('tm_43fa702e0c3aa6bb91d79e95', '2023-07-26 22:21:59', '2023-07-26 22:21:59', 'p_3ba460634520cf4590dc90e5', 'attributes', '编号', 'num', '', 'string', '{\"rw\": \"rw\", \"default_value\": \"23332442\"}');
INSERT INTO `product_templates` VALUES ('tm_538231f64592eb53b6d46d12', '2023-09-08 13:57:19', '2023-09-08 13:57:19', 'p_cdbb1eccd902018d51fe062e', 'attributes', '版本号', 'version', '', 'string', '{\"rw\": \"r\", \"default_value\": \"v1.0\"}');
INSERT INTO `product_templates` VALUES ('tm_925cec0662102b40fe33b7bb', '2023-07-26 22:20:45', '2023-07-26 22:20:45', 'p_3ba460634520cf4590dc90e5', 'telemetry', '湿度', 'humidity', '', 'float64', '{\"max\": \"100\", \"min\": \"1\", \"step\": 0.01, \"unit\": \"G\"}');
INSERT INTO `product_templates` VALUES ('tm_ac52beea237bb9009f1029af', '2023-07-26 22:20:08', '2023-07-26 22:20:08', 'p_3ba460634520cf4590dc90e5', 'telemetry', '温度', 'temperature', '', 'float64', '{\"max\": \"100\", \"min\": \"1\", \"step\": 0.01, \"unit\": \"度\"}');
INSERT INTO `product_templates` VALUES ('tm_e815087669adc6f9fcf6bcf4', '2023-08-01 14:14:47', '2023-08-01 14:14:47', 'p_3ba460634520cf4590dc90e5', 'commands', '重启', 'restart', '设备重启指令', '', '{\"input\": [{\"key\": \"aa\", \"name\": \"重启参数\", \"type\": \"int64\", \"define\": {\"max\": 100, \"min\": 1, \"step\": 1, \"unit\": \"KW\"}}], \"output\": []}');
INSERT INTO `product_templates` VALUES ('tm-4991928839c4dec5c08109f5', '2023-07-21 10:38:02', '2023-07-21 10:38:02', 'p03d9a6fb450e8443456f41b0', 'telemetry', '电流', 'i', '', 'float64', '{\"max\": \"100\", \"min\": \"1\", \"step\": \"1\", \"unit\": \"A\"}');
INSERT INTO `product_templates` VALUES ('tm-9e922dad5c325348c123103d', '2023-07-21 10:37:05', '2023-07-21 10:37:05', 'p03d9a6fb450e8443456f41b0', 'attributes', '序列号', 'ns', '', 'string', '{\"rw\": \"r\", \"default_value\": \"NS42342\"}');
INSERT INTO `product_templates` VALUES ('tm-a2998852fd8c1507cfc8d0e1', '2023-07-21 10:37:39', '2023-07-21 10:57:49', 'p03d9a6fb450e8443456f41b0', 'telemetry', '电压', 'u', '', 'float64', '{\"max\": 100, \"min\": 1, \"step\": 1, \"unit\": \"V\"}');
INSERT INTO `product_templates` VALUES ('UksLt1hVdQ', '2023-10-05 11:43:31', '2023-10-05 11:43:31', 'p_3ba460634520cf4590dc90e5', 'commands', '固件升级', 'ota', '', '', '{\"input\": [{\"key\": \"version\", \"name\": \"版本\", \"type\": \"string\", \"define\": {}, \"description\": \"要升级的版本\"}, {\"key\": \"url\", \"name\": \"固件路径\", \"type\": \"string\", \"define\": {}, \"description\": \"固件路径\"}, {\"key\": \"module\", \"name\": \"固件模块\", \"type\": \"string\", \"define\": {}}, {\"key\": \"check\", \"name\": \"校验和\", \"type\": \"string\", \"define\": {}, \"description\": \"md5校验，校验和\"}], \"output\": []}');
INSERT INTO `product_templates` VALUES ('VFuqZIlNnQ', '2023-09-27 11:36:15', '2023-09-27 11:36:15', 'mSOWuiA97g', 'telemetry', '湿度', 'humidity', '', 'float64', '{\"max\": 100, \"min\": 0, \"unit\": \"RH\"}');
INSERT INTO `product_templates` VALUES ('wR1s2TfugA', '2023-09-27 15:07:54', '2023-09-27 15:07:54', 'ek2WUADl6g', 'telemetry', '温度', 'temperature', '', 'float64', '{\"max\": 100, \"min\": 0, \"unit\": \"摄氏度\"}');
INSERT INTO `product_templates` VALUES ('YoDVrJFyAg', '2023-09-27 16:43:14', '2023-09-27 16:43:14', 'mSOWuiA97g', 'commands', '关闭指示灯', 'closeD', '', '', '{\"input\": [{\"key\": \"close\", \"name\": \"关闭\", \"type\": \"string\", \"define\": {}}], \"output\": []}');

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `owner` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` int(0) NULL DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品名称',
  `photo_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '图片地址',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品说明',
  `product_category_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '产品类型Id',
  `protocol_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '协议名称',
  `device_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备类型',
  `rule_chain_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '规则链Id',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_products_product_category`(`product_category_id`) USING BTREE,
  CONSTRAINT `fk_products_product_category` FOREIGN KEY (`product_category_id`) REFERENCES `product_categories` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES ('ek2WUADl6g', 'panda', 2, '2023-09-27 15:07:02', '2023-09-27 15:07:02', 'HTTP设备产品', '', '', 'pcd2e673d2cd92e860cff5d958', 'HTTP', 'direct', 'rulee765e9ef022812a8b89dfb4c', '0');
INSERT INTO `products` VALUES ('Fb3DlUja_Q', 'panda', 2, '2023-09-27 11:27:24', '2023-09-27 11:27:24', 'MQTT透传解析', '', '', 'pcd2e673d2cd92e860cff5d958', 'MQTT', 'direct', '', '0');
INSERT INTO `products` VALUES ('I_HlHDdh_Q', 'panda', 2, '2023-10-14 17:37:14', '2023-10-17 11:36:32', '默认子网关', '', '', 'KVys13MMsA', 'MQTT', 'gatewayS', 'rulee765e9ef022812a8b89dfb4c', '0');
INSERT INTO `products` VALUES ('kqEUXwsU9w', 'panda', 2, '2023-08-19 09:26:50', '2023-10-05 10:41:22', '测试产品1', '9b37cd4ca37090649adcee8bf17cfdcc_20230414141350.png', '', 'pcd2e673d2cd92e860cff5d958', 'MQTT', 'direct', 'rulee765e9ef022812a8b89dfb4c', '0');
INSERT INTO `products` VALUES ('M32969chcw', 'panda', 2, '2023-10-12 19:20:20', '2023-10-12 19:20:20', '松下PLC', '', '', 'oKQcjqY8ZQ', 'MQTT', 'gatewayS', 'rulee765e9ef022812a8b89dfb4c', '0');
INSERT INTO `products` VALUES ('mSOWuiA97g', 'panda', 2, '2023-09-27 11:25:11', '2023-09-27 11:26:20', 'TCP透传测试产品', '', '', 'pcd2e673d2cd92e860cff5d958', 'TCP', 'direct', 'mq1YRZbUgQ', '0');
INSERT INTO `products` VALUES ('p_3ba460634520cf4590dc90e5', 'panda', 2, '2023-07-26 22:17:27', '2023-08-03 10:13:45', '测试产品', '', '', 'pcd2e673d2cd92e860cff5d958', 'MQTT', 'direct', 'rule_a37571bb6c45378b57803793', '0');
INSERT INTO `products` VALUES ('p_bf52caf91f7cdd2abb52eaaf', 'panda', 2, '2023-07-31 14:16:29', '2023-07-31 14:16:29', '智能控制器', '', '', 'pcd2e673d2cd92e860cff5d958', 'MQTT', 'gatewayS', 'rulee765e9ef022812a8b89dfb4c', '0');
INSERT INTO `products` VALUES ('p_cdbb1eccd902018d51fe062e', 'panda', 2, '2023-07-31 14:15:35', '2023-07-31 14:15:35', '网关设备', '', '网关设备', 'pcd2e673d2cd92e860cff5d958', 'MQTT', 'gateway', 'rulee765e9ef022812a8b89dfb4c', '0');
INSERT INTO `products` VALUES ('uqNNwYJ5rw', 'panda', 2, '2023-10-12 19:19:17', '2023-10-12 19:19:17', '星原网关', '1df420e901be965018e95bac136ec17f_20231012191851.jpg', '', 'oKQcjqY8ZQ', 'MQTT', 'gateway', 'rulee765e9ef022812a8b89dfb4c', '0');

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `sys_apis_udk` (`path`,`method`)
) ENGINE = InnoDB AUTO_INCREMENT = 206 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2021-12-09 09:21:04', '2021-12-09 09:21:04', NULL, '/system/user/list', '查询用户列表（分页）', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (2, '2021-12-09 09:29:36', '2021-12-09 09:29:36', NULL, '/system/user/changeStatus', '修改用户状态', 'user', 'PUT');
INSERT INTO `sys_apis` VALUES (3, '2021-12-09 09:34:37', '2021-12-09 09:34:37', NULL, '/system/user/:userId', '删除用户信息', 'user', 'DELETE');
INSERT INTO `sys_apis` VALUES (4, '2021-12-09 09:36:43', '2023-09-14 14:05:54', NULL, '/system/organization/list', '获取组织列表', 'organization', 'GET');
INSERT INTO `sys_apis` VALUES (5, '2021-12-09 09:37:31', '2023-09-14 14:06:51', NULL, '/system/organization/:organizationId', '获取组织信息', 'organization', 'GET');
INSERT INTO `sys_apis` VALUES (6, '2021-12-09 18:20:32', '2021-12-09 18:20:32', NULL, '/system/user/avatar', '上传头像', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (7, '2021-12-09 18:21:10', '2021-12-09 18:21:10', NULL, '/system/user/pwd', '修改密码', 'user', 'PUT');
INSERT INTO `sys_apis` VALUES (8, '2021-12-09 18:21:54', '2021-12-09 18:21:54', NULL, '/system/user/getById/:userId', '通过id获取用户信息', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (9, '2021-12-09 18:58:50', '2021-12-09 18:58:50', NULL, '/system/user/getInit', '获取初始化角色岗位信息(添加用户初始化)', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (10, '2021-12-09 18:59:43', '2021-12-09 18:59:43', NULL, '/system/user/getRoPo', '获取用户角色岗位信息', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (11, '2021-12-09 19:00:24', '2021-12-09 19:00:24', NULL, '/system/user', '添加用户信息', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2021-12-09 19:00:52', '2021-12-09 19:00:52', NULL, '/system/user', '修改用户信息', 'user', 'PUT');
INSERT INTO `sys_apis` VALUES (13, '2021-12-09 19:02:30', '2021-12-09 19:02:30', NULL, '/system/user/export', '导出用户信息', 'user', 'GET');
INSERT INTO `sys_apis` VALUES (14, '2021-12-09 19:04:04', '2023-09-14 14:06:35', NULL, '/system/organization/roleOrganizationTreeSelect/:roleId', '获取角色部门树', 'organization', 'GET');
INSERT INTO `sys_apis` VALUES (15, '2021-12-09 19:04:48', '2023-09-14 14:07:06', NULL, '/system/organization/organizationTree', '获取所有组织树', 'organization', 'GET');
INSERT INTO `sys_apis` VALUES (16, '2021-12-09 19:07:37', '2023-09-14 14:07:18', NULL, '/system/organization', '添加组织信息', 'organization', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2021-12-09 19:08:14', '2023-09-14 14:07:28', NULL, '/system/organization', '修改组织信息', 'organization', 'PUT');
INSERT INTO `sys_apis` VALUES (18, '2021-12-09 19:08:40', '2023-09-14 14:07:41', NULL, '/system/organization/:organizationId', '删除组织信息', 'organization', 'DELETE');
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
INSERT INTO `sys_apis` VALUES (79, '2021-12-24 15:45:03', '2023-08-08 14:15:59', NULL, '/job/log/list', '任务日志列表', 'job', 'GET');
INSERT INTO `sys_apis` VALUES (80, '2021-12-24 15:45:33', '2023-08-08 14:16:07', NULL, '/job/log/all', '清空任务日志', 'job', 'DELETE');
INSERT INTO `sys_apis` VALUES (81, '2021-12-24 15:46:08', '2023-08-08 14:16:15', NULL, '/job/log/:logId', '删除任务日志', 'job', 'DELETE');
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
INSERT INTO `sys_apis` VALUES (124, '2023-06-29 16:59:08', '2023-06-29 17:00:17', NULL, '/device/product/category/list', '获取产品分类列表', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (125, '2023-06-29 17:00:08', '2023-06-29 17:00:08', NULL, '/device/product/category/list/all', '获取所有列表', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (126, '2023-06-29 17:00:56', '2023-06-29 17:00:56', NULL, '/device/product/category/list/tree', '获取树', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (127, '2023-06-29 17:01:44', '2023-06-29 17:01:44', NULL, '/device/product/category/:id', '查询单个', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (128, '2023-06-29 17:02:16', '2023-06-29 17:02:16', NULL, '/device/product/category', '添加分类', 'product', 'POST');
INSERT INTO `sys_apis` VALUES (129, '2023-06-29 17:02:42', '2023-06-29 17:02:42', NULL, '/device/product/category', '修改分类', 'product', 'PUT');
INSERT INTO `sys_apis` VALUES (130, '2023-06-29 17:03:07', '2023-06-29 17:03:07', NULL, '/device/product/category/:id', '删除分类', 'product', 'DELETE');
INSERT INTO `sys_apis` VALUES (131, '2023-06-29 16:59:08', '2023-06-29 17:00:17', NULL, '/device/group/list', '获取设备分组列表', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (132, '2023-06-29 17:00:08', '2023-06-29 17:00:08', NULL, '/device/group/list/all', '获取所有列表', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (133, '2023-06-29 17:00:56', '2023-06-29 17:00:56', NULL, '/device/group/list/tree', '获取树', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (134, '2023-06-29 17:01:44', '2023-06-29 17:01:44', NULL, '/device/group/:id', '查询单个', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (135, '2023-06-29 17:02:16', '2023-06-29 17:02:16', NULL, '/device/group', '添加分组', 'device', 'POST');
INSERT INTO `sys_apis` VALUES (136, '2023-06-29 17:02:42', '2023-06-29 17:02:42', NULL, '/device/group', '修改分组', 'device', 'PUT');
INSERT INTO `sys_apis` VALUES (137, '2023-06-29 17:03:07', '2023-06-29 17:03:07', NULL, '/device/group/:id', '删除分组', 'device', 'DELETE');
INSERT INTO `sys_apis` VALUES (138, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product/:id', '删除产品信息', 'product', 'DELETE');
INSERT INTO `sys_apis` VALUES (139, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product/:id', '获取产品信息', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (140, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product', '修改产品信息', 'product', 'PUT');
INSERT INTO `sys_apis` VALUES (141, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product/list', '查询产品列表（分页）', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (142, '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL, '/device/product', '添加产品信息', 'product', 'POST');
INSERT INTO `sys_apis` VALUES (143, '2023-06-30 14:20:03', '2023-06-30 15:26:46', NULL, '/device/list', '查询设备列表（分页）', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (144, '2023-06-30 14:20:03', '2023-06-30 15:26:52', NULL, '/device/:id', '获取设备信息', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (145, '2023-06-30 14:20:03', '2023-06-30 15:26:57', NULL, '/device', '添加设备信息', 'device', 'POST');
INSERT INTO `sys_apis` VALUES (146, '2023-06-30 14:20:03', '2023-06-30 15:27:04', NULL, '/device/:id', '删除设备信息', 'device', 'DELETE');
INSERT INTO `sys_apis` VALUES (147, '2023-06-30 14:20:03', '2023-06-30 15:27:09', NULL, '/device', '修改设备信息', 'device', 'PUT');
INSERT INTO `sys_apis` VALUES (148, '2023-06-30 15:11:25', '2023-08-02 16:06:13', NULL, '/device/group/list/tree/label', '获取设备组label', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (149, '2023-06-30 15:14:08', '2023-09-22 16:58:04', NULL, '/device/product/category/list/tree/label', '获取设置分类树', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (150, '2023-07-01 10:44:25', '2023-07-01 10:44:25', NULL, '/upload/up/oss', '上传文件到oss', 'upload', 'POST');
INSERT INTO `sys_apis` VALUES (151, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota/list', '查询产品固件列表（分页）', 'ota', 'GET');
INSERT INTO `sys_apis` VALUES (152, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota', '添加产品固件信息', 'ota', 'POST');
INSERT INTO `sys_apis` VALUES (153, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota', '修改产品固件信息', 'ota', 'PUT');
INSERT INTO `sys_apis` VALUES (154, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota/:id', '删除产品固件信息', 'ota', 'DELETE');
INSERT INTO `sys_apis` VALUES (155, '2023-07-06 15:31:15', '2023-07-06 15:31:15', NULL, '/device/ota/:id', '获取产品固件信息', 'ota', 'GET');
INSERT INTO `sys_apis` VALUES (156, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template/list', '查询产品模型列表（分页）', 'template', 'GET');
INSERT INTO `sys_apis` VALUES (157, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template', '修改产品模型信息', 'template', 'PUT');
INSERT INTO `sys_apis` VALUES (158, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template/:id', '获取产品模型信息', 'template', 'GET');
INSERT INTO `sys_apis` VALUES (159, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template/:id', '删除产品模型信息', 'template', 'DELETE');
INSERT INTO `sys_apis` VALUES (160, '2023-07-06 15:32:10', '2023-07-06 15:32:10', NULL, '/device/template', '添加产品模型信息', 'template', 'POST');
INSERT INTO `sys_apis` VALUES (161, '2023-07-07 16:35:45', '2023-07-07 16:35:45', NULL, '/device/product/list/all', '获取所有列表', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (162, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen', '修改bi大屏信息', 'screen', 'PUT');
INSERT INTO `sys_apis` VALUES (163, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen/:screenId', '获取bi大屏信息', 'screen', 'GET');
INSERT INTO `sys_apis` VALUES (164, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen/list', '查询bi大屏列表（分页）', 'screen', 'GET');
INSERT INTO `sys_apis` VALUES (165, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen/:screenId', '删除bi大屏信息', 'screen', 'DELETE');
INSERT INTO `sys_apis` VALUES (166, '2023-04-13 09:03:47', '2023-04-13 09:03:47', NULL, '/visual/screen', '添加bi大屏信息', 'screen', 'POST');
INSERT INTO `sys_apis` VALUES (167, '2023-04-13 10:15:27', '2023-04-13 10:15:27', NULL, '/visual/screen/group/list', '大屏分组列表', 'screen_group', 'GET');
INSERT INTO `sys_apis` VALUES (168, '2023-04-13 10:16:15', '2023-04-13 10:16:15', NULL, '/visual/screen/group/list/tree', '大屏分组列表树', 'screen_group', 'GET');
INSERT INTO `sys_apis` VALUES (169, '2023-04-13 10:16:38', '2023-04-13 10:16:38', NULL, '/visual/screen/group/list/all', '获取所有分组', 'screen_group', 'GET');
INSERT INTO `sys_apis` VALUES (170, '2023-04-13 10:17:34', '2023-04-13 10:17:34', NULL, '/visual/screen/group/:id', '获取分组', 'screen_group', 'GET');
INSERT INTO `sys_apis` VALUES (171, '2023-04-13 10:18:10', '2023-04-13 10:18:10', NULL, '/visual/screen/group', '添加分组', 'screen_group', 'POST');
INSERT INTO `sys_apis` VALUES (172, '2023-04-13 10:18:35', '2023-04-13 10:18:35', NULL, '/visual/screen/group', '修改分组', 'screen_group', 'PUT');
INSERT INTO `sys_apis` VALUES (173, '2023-04-13 10:19:09', '2023-04-13 10:19:09', NULL, '/visual/screen/group/:id', '删除分组', 'screen_group', 'DELETE');
INSERT INTO `sys_apis` VALUES (174, '2023-04-13 15:49:39', '2023-04-13 15:49:39', NULL, '/visual/screen/changeStatus', '改变状态', 'screen', 'PUT');
INSERT INTO `sys_apis` VALUES (175, '2023-04-13 15:50:18', '2023-07-21 17:44:48', NULL, '/rule/chain/changeRoot', '改变规则链', 'rulechain', 'PUT');
INSERT INTO `sys_apis` VALUES (176, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain/list', '查询规则链列表（分页）', 'rulechain', 'GET');
INSERT INTO `sys_apis` VALUES (177, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain/:ruleId', '删除规则链信息', 'rulechain', 'DELETE');
INSERT INTO `sys_apis` VALUES (178, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain', '修改规则链信息', 'rulechain', 'PUT');
INSERT INTO `sys_apis` VALUES (179, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain', '添加规则链信息', 'rulechain', 'POST');
INSERT INTO `sys_apis` VALUES (180, '2023-04-11 02:05:25', '2023-04-11 02:05:25', NULL, '/rule/chain/:ruleId', '获取规则链信息', 'rulechain', 'GET');
INSERT INTO `sys_apis` VALUES (181, '2023-07-24 11:51:10', '2023-07-24 11:51:10', NULL, '/rule/chain/list/label', '获取规则链label列表', 'rulechain', 'GET');
INSERT INTO `sys_apis` VALUES (182, '2023-07-31 14:14:06', '2023-07-31 14:14:06', NULL, '/device/list/all', '获取所有设备', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (183, '2023-08-02 16:05:24', '2023-08-02 16:05:24', NULL, '/device/:id/status', '获取设备状态', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (184, '2023-08-03 09:50:41', '2023-08-03 09:50:41', NULL, '/rule/chain/clone/:ruleId', '克隆规则链', 'rulechain', 'POST');
INSERT INTO `sys_apis` VALUES (185, '2023-08-03 14:16:55', '2023-08-03 14:16:55', NULL, '/device/alarm/list', '告警分页列表', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (186, '2023-08-03 14:17:23', '2023-08-03 14:17:23', NULL, '/device/alarm', '修改告警', 'device', 'PUT');
INSERT INTO `sys_apis` VALUES (187, '2023-08-03 14:18:14', '2023-08-03 14:18:14', NULL, '/device/alarm/:id', '删除告警信息', 'device', 'DELETE');
INSERT INTO `sys_apis` VALUES (188, '2023-08-04 10:59:57', '2023-08-04 10:59:57', NULL, '/device/cmd/list', '设备命令日志列表', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (189, '2023-08-04 11:00:18', '2023-08-04 11:00:18', NULL, '/device/cmd', '下发命令', 'device', 'POST');
INSERT INTO `sys_apis` VALUES (190, '2023-08-04 11:00:46', '2023-08-04 11:00:46', NULL, '/device/cmd/:id', '删除命令记录', 'device', 'DELETE');
INSERT INTO `sys_apis` VALUES (191, '2023-08-04 14:16:06', '2023-08-04 14:16:06', NULL, '/device/template/list/all', '查询所有tsl', 'template', 'GET');
INSERT INTO `sys_apis` VALUES (192, '2023-08-04 16:39:06', '2023-08-04 16:39:06', NULL, '/device/:id/attribute/down', '下发设备属性', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (193, '2023-08-19 09:12:31', '2023-08-19 09:12:31', NULL, '/upload/up', '上传文件到本地', 'upload', 'POST');
INSERT INTO `sys_apis` VALUES (194, '2023-09-05 08:42:13', '2023-09-05 08:42:13', NULL, '/video/ys/device/list', '获取萤石设备列表', 'video', 'GET');
INSERT INTO `sys_apis` VALUES (195, '2023-09-05 08:43:11', '2023-09-05 08:43:11', NULL, '/video/ys/:deviceSerial/channel', '获取指定设备通道', 'video', 'GET');
INSERT INTO `sys_apis` VALUES (196, '2023-09-05 08:45:31', '2023-09-05 08:45:31', NULL, '/video/ys/:deviceSerial/channel/live', '设备通道直播地址', 'video', 'GET');
INSERT INTO `sys_apis` VALUES (197, '2023-09-05 08:46:14', '2023-09-05 08:46:14', NULL, '/video/ys/:deviceSerial/ptz/start', '摄像头操作', 'video', 'GET');
INSERT INTO `sys_apis` VALUES (198, '2023-09-05 08:46:47', '2023-09-05 08:46:47', NULL, '/video/ys/:deviceSerial/ptz/stop', '摄像头操作停止', 'video', 'GET');
INSERT INTO `sys_apis` VALUES (199, '2023-09-06 15:55:44', '2023-09-06 15:55:44', NULL, '/rule/chain/log/list', '规则链审计日志', 'rulechain', 'GET');
INSERT INTO `sys_apis` VALUES (200, '2023-09-06 15:56:35', '2023-09-06 15:56:35', NULL, '/rule/chain/log/delete', '条件删除规则链审计', 'rulechain', 'GET');
INSERT INTO `sys_apis` VALUES (201, '2023-09-08 17:20:35', '2023-09-08 17:20:35', NULL, '/device/:id/property/history', '获取设备属性的遥测历史', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (202, '2023-09-15 17:29:07', '2023-09-15 17:29:07', NULL, '/device/:id/allot/org', '设备分配组织', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (203, '2023-09-22 16:56:56', '2023-09-22 16:56:56', NULL, '/device/product/:id/tsl', '获取产品tsl', 'product', 'GET');
INSERT INTO `sys_apis` VALUES (204, '2023-09-23 14:25:58', '2023-09-23 14:25:58', NULL, '/device/panel', '获取设备统计面板', 'device', 'GET');
INSERT INTO `sys_apis` VALUES (205, '2023-09-25 10:13:59', '2023-09-25 10:13:59', NULL, '/device/alarm/panel', '获取面板告警分组', 'device', 'GET');

-- ----------------------------
-- Table structure for sys_configs
-- ----------------------------
DROP TABLE IF EXISTS `sys_configs`;
CREATE TABLE `sys_configs`  (
  `config_id` bigint(0) NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '是否系统内置0，1',
  `is_frontend` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Remark',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_configs
-- ----------------------------
INSERT INTO `sys_configs` VALUES (1, '账号初始密码', 'sys.user.initPassword', '123456', '0', '0', '初始密码', '2021-12-04 13:50:13', '2021-12-04 13:54:52', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(0) NOT NULL AUTO_INCREMENT,
  `dict_sort` int(0) NULL DEFAULT NULL COMMENT '排序',
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
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

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
INSERT INTO `sys_dict_data` VALUES (39, 1, '未发布', '0', 'sys_release_type', '0', '', '', '', 'panda', '', '', '2023-07-21 16:11:27', '2023-07-21 16:11:27', NULL);
INSERT INTO `sys_dict_data` VALUES (40, 2, '已发布', '1', 'sys_release_type', '0', '', '', '', 'panda', '', '', '2023-07-21 16:11:44', '2023-07-21 16:11:44', NULL);

-- ----------------------------
-- Table structure for sys_dict_types
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_types`;
CREATE TABLE `sys_dict_types`  (
  `dict_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `dict_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `create_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`dict_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

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
INSERT INTO `sys_dict_types` VALUES (33, '发布状态', 'sys_release_type', '0', 'panda', '', '发布状态', '2023-07-21 16:10:38', '2023-07-21 16:10:38', NULL);

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `menu_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `parent_id` int(0) NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
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
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 179 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (1, '系统设置', '', 0, 0, 'elementSetting', '/system', 'Layout', '1', '', 'M', '0', '0', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08', '2021-12-28 13:32:21', NULL);
INSERT INTO `sys_menus` VALUES (3, '用户管理', '', 1, 1, 'elementUser', '/system/user', '/system/user/index', '1', '', 'C', '0', '1', '1', 'system:user:list', '0', 'admin', 'panda', '', '2021-12-02 14:07:56', '2021-12-28 13:32:44', NULL);
INSERT INTO `sys_menus` VALUES (4, '添加用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:add', '0', 'admin', '', '', '2021-12-03 13:36:33', '2021-12-03 13:36:33', NULL);
INSERT INTO `sys_menus` VALUES (5, '编辑用户', '', 3, 1, '', '', '', '', '', 'F', '0', '', '', 'system:user:edit', '0', 'admin', '', '', '2021-12-03 13:48:13', '2021-12-03 13:48:13', NULL);
INSERT INTO `sys_menus` VALUES (6, '角色管理', '', 1, 2, 'elementUserFilled', '/system/role', '/system/role/index', '1', '', 'C', '0', '1', '1', 'system:role:list', '0', '', 'panda', '', '2021-12-03 13:51:55', '2022-07-16 10:23:21', NULL);
INSERT INTO `sys_menus` VALUES (7, '菜单管理', '', 1, 2, 'iconfont icon-juxingkaobei', '/system/menu', '/system/menu/index', '1', '', 'C', '0', '1', '1', 'system:menu:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:33:19', NULL);
INSERT INTO `sys_menus` VALUES (8, '组织管理', '', 1, 3, 'iconfont icon-jiliandongxuanzeqi', '/system/organization', '/system/organization/index', '1', '', 'C', '0', '1', '1', 'system:organization:list', '0', '', 'panda', '', '2021-12-03 13:58:36', '2023-09-14 14:05:07', NULL);
INSERT INTO `sys_menus` VALUES (9, '岗位管理', '', 1, 4, 'iconfont icon-neiqianshujuchucun', '/system/post', '/system/post/index', '1', '', 'C', '0', '1', '1', 'system:post:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:40:31', NULL);
INSERT INTO `sys_menus` VALUES (10, '字典管理', '', 1, 5, 'elementCellphone', '/system/dict', '/system/dict/index', '1', '', 'C', '0', '1', '1', 'system:dict:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:40:50', NULL);
INSERT INTO `sys_menus` VALUES (11, '参数管理', '', 1, 6, 'elementDocumentCopy', '/system/config', '/system/config/index', '1', '', 'C', '0', '1', '1', 'system:config:list', '0', 'admin', 'panda', '', '2021-12-03 13:54:44', '2021-12-28 13:41:05', NULL);
INSERT INTO `sys_menus` VALUES (12, '个人中心', '', 0, 10, 'elementAvatar', '/personal', '/personal/index', '1', '', 'M', '1', '0', '0', '', '0', 'admin', 'panda', '', '2021-12-03 14:12:43', '2023-06-27 10:09:26', NULL);
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
INSERT INTO `sys_menus` VALUES (24, '添加部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:organization:add', '0', '', 'panda', '', '2021-12-07 09:33:58', '2023-09-14 14:05:20', NULL);
INSERT INTO `sys_menus` VALUES (25, '编辑部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:organization:edit', '0', '', 'panda', '', '2021-12-07 09:34:39', '2023-09-14 14:05:26', NULL);
INSERT INTO `sys_menus` VALUES (26, '删除部门', '', 8, 1, '', '', '', '', '', 'F', '', '', '', 'system:organization:delete', '0', '', 'panda', '', '2021-12-07 09:35:09', '2023-09-14 14:05:32', NULL);
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
INSERT INTO `sys_menus` VALUES (43, '日志系统', '', 0, 11, 'iconfont icon-biaodan', '/log', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-02 11:04:08', '2023-06-30 08:57:08', NULL);
INSERT INTO `sys_menus` VALUES (44, '告警监控', '', 0, 9, 'iconfont icon-gongju', '/tool', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-16 16:35:15', '2023-10-18 10:17:52', NULL);
INSERT INTO `sys_menus` VALUES (45, '操作日志', '', 43, 1, 'iconfont icon-bolangnengshiyanchang', '/log/operation', '/log/operation/index', '1', '', 'C', '0', '1', '1', 'log:operation:list', '0', 'admin', 'panda', '', '2021-12-16 16:42:03', '2021-12-28 13:39:44', NULL);
INSERT INTO `sys_menus` VALUES (46, '登录日志', '', 43, 2, 'iconfont icon--chaifenlie', '/log/login', '/log/login/index', '1', '', 'C', '0', '1', '1', 'log:login:list', '0', 'admin', 'panda', '', '2021-12-16 16:43:28', '2021-12-28 13:39:58', NULL);
INSERT INTO `sys_menus` VALUES (47, '服务监控', '', 44, 1, 'elementCpu', '/tool/monitor/', '/tool/monitor/index', '1', '', 'C', '0', '1', '1', 'tool:monitor:list', '0', 'admin', 'panda', '', '2021-12-03 14:12:43', '2021-12-28 13:41:25', NULL);
INSERT INTO `sys_menus` VALUES (49, '开发工具', '', 0, 10, 'iconfont icon-zhongduancanshu', '/develop', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'admin', 'panda', '', '2021-12-16 16:53:11', '2023-06-29 16:29:23', NULL);
INSERT INTO `sys_menus` VALUES (50, '表单构建', '', 49, 1, 'iconfont icon-zidingyibuju', '/develop/form', '/develop/form/index', '1', '', 'C', '0', '1', '1', 'develop:form:list', '0', 'admin', 'panda', '', '2021-12-16 16:55:01', '2022-07-12 18:56:18', NULL);
INSERT INTO `sys_menus` VALUES (51, '代码生成', '', 49, 2, 'iconfont icon-zhongduancanshu', '/develop/code', '/develop/code/index', '1', '', 'C', '0', '1', '1', 'develop:code:list', '0', 'admin', '', '', '2021-12-16 16:56:48', '2021-12-16 16:56:48', NULL);
INSERT INTO `sys_menus` VALUES (52, '系统接口', '', 49, 3, 'iconfont icon-wenducanshu-05', '/develop/apis', '/layout/routerView/iframes', '0', 'https://82200r6gti.apifox.cn', 'C', '0', '1', '1', 'develop:apis:list', '0', '', 'panda', '', '2021-12-16 16:58:07', '2023-09-04 11:02:29', NULL);
INSERT INTO `sys_menus` VALUES (54, '对象存储', '', 53, 1, 'iconfont icon-chazhaobiaodanliebiao', '/resource/file', '/resource/file/index', '1', '', 'C', '0', '1', '1', 'resource:file:list', '0', 'admin', 'panda', '', '2021-12-16 17:06:04', '2022-01-13 17:30:09', NULL);
INSERT INTO `sys_menus` VALUES (55, '公告通知', '', 44, 3, 'elementTicket', '/tool/notice', '/tool/notice/index', '1', '', 'C', '0', '1', '1', 'tool:notice:list', '0', 'admin', 'panda', '', '2021-12-16 22:09:11', '2021-12-28 13:42:39', NULL);
INSERT INTO `sys_menus` VALUES (59, '删除', '', 45, 1, '', '', '', '', '', 'F', '', '', '', 'log:operation:delete', '0', 'panda', '', '', '2022-01-14 13:28:25', '2022-01-14 13:28:25', NULL);
INSERT INTO `sys_menus` VALUES (60, '清空', '', 45, 1, '', '', '', '', '', 'F', '', '', '', 'log:operation:clean', '0', 'panda', '', '', '2022-01-14 13:29:24', '2022-01-14 13:29:24', NULL);
INSERT INTO `sys_menus` VALUES (63, '删除', '', 46, 1, '', '', '', '', '', 'F', '', '', '', 'log:login:delete', '0', 'panda', '', '', '2022-01-14 13:30:46', '2022-01-14 13:30:46', NULL);
INSERT INTO `sys_menus` VALUES (64, '清空', '', 46, 1, '', '', '', '', '', 'F', '', '', '', 'log:login:clean', '0', 'panda', '', '', '2022-01-14 13:31:06', '2022-01-14 13:31:06', NULL);
INSERT INTO `sys_menus` VALUES (69, '添加', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:add', '0', 'panda', '', '', '2022-01-14 13:35:23', '2022-01-14 13:35:23', NULL);
INSERT INTO `sys_menus` VALUES (70, '编辑', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:edit', '0', 'panda', '', '', '2022-01-14 13:36:04', '2022-01-14 13:36:04', NULL);
INSERT INTO `sys_menus` VALUES (71, '删除', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:delete', '0', 'panda', '', '', '2022-01-14 13:36:26', '2022-01-14 13:36:26', NULL);
INSERT INTO `sys_menus` VALUES (72, '查看', '', 55, 1, '', '', '', '', '', 'F', '', '', '', 'tool:notice:view', '0', 'panda', '', '', '2022-01-14 13:36:51', '2022-01-14 13:36:51', NULL);
INSERT INTO `sys_menus` VALUES (73, '导入', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:add', '0', 'panda', '', '', '2022-01-14 13:38:35', '2022-01-14 13:38:35', NULL);
INSERT INTO `sys_menus` VALUES (74, '编辑', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:edit', '0', 'panda', '', '', '2022-01-14 13:41:25', '2022-01-14 13:41:25', NULL);
INSERT INTO `sys_menus` VALUES (75, '删除', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:delete', '0', 'panda', '', '', '2022-01-14 13:41:42', '2022-01-14 13:41:42', NULL);
INSERT INTO `sys_menus` VALUES (76, '预览', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:view', '0', 'panda', '', '', '2022-01-14 13:42:01', '2022-01-14 13:42:01', NULL);
INSERT INTO `sys_menus` VALUES (77, '生成代码', '', 51, 1, '', '', '', '', '', 'F', '', '', '', 'develop:code:code', '0', 'panda', '', '', '2022-01-14 13:42:48', '2022-01-14 13:42:48', NULL);
INSERT INTO `sys_menus` VALUES (95, '设备管理', '', 0, 1, 'iconfont icon-dongtai', '/device', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'panda', 'panda', '', '2023-06-29 16:21:31', '2023-09-02 15:03:55', NULL);
INSERT INTO `sys_menus` VALUES (96, '规则链库', '', 0, 2, 'iconfont icon-shuxingtu', '/rule', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'panda', 'panda', '', '2023-06-29 16:33:23', '2023-09-02 15:04:06', NULL);
INSERT INTO `sys_menus` VALUES (97, '组态大屏', '', 0, 3, 'iconfont icon-diannaobangong', '/visual', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'panda', 'panda', '', '2023-06-29 16:34:50', '2023-07-21 14:41:55', NULL);
INSERT INTO `sys_menus` VALUES (98, '产品分类', '', 95, 1, 'iconfont icon-jiliandongxuanzeqi', '/device/product_category', '/device/product_category/index', '1', '', 'C', '0', '1', '1', 'product:category:list', '0', '', 'panda', '', '2023-06-29 16:44:56', '2023-06-29 16:49:55', NULL);
INSERT INTO `sys_menus` VALUES (100, '设备分组', '', 95, 3, 'iconfont icon-zidingyibuju', '/device/device_group', '/device/device_group/index', '1', '', 'C', '0', '1', '1', 'device:group:list', '0', '', 'panda', '', '2023-06-29 16:48:05', '2023-06-29 16:50:49', NULL);
INSERT INTO `sys_menus` VALUES (102, '添加', '', 98, 1, '', '', '', '', '', 'F', '', '', '', 'product:category:add', '0', 'panda', '', '', '2023-06-29 16:51:38', '2023-06-29 16:51:38', NULL);
INSERT INTO `sys_menus` VALUES (103, '修改', '', 98, 2, '', '', '', '', '', 'F', '', '', '', 'product:category:edit', '0', 'panda', '', '', '2023-06-29 16:52:00', '2023-06-29 16:52:00', NULL);
INSERT INTO `sys_menus` VALUES (104, '删除', '', 98, 3, '', '', '', '', '', 'F', '', '', '', 'product:category:delete', '0', 'panda', '', '', '2023-06-29 16:52:36', '2023-06-29 16:52:36', NULL);
INSERT INTO `sys_menus` VALUES (105, '新增', '', 100, 1, '', '', '', '', '', 'F', '', '', '', 'device:group:add', '0', 'panda', '', '', '2023-06-29 16:53:16', '2023-06-29 16:53:16', NULL);
INSERT INTO `sys_menus` VALUES (106, '修改', '', 100, 2, '', '', '', '', '', 'F', '', '', '', 'device:group:edit', '0', 'panda', '', '', '2023-06-29 16:53:37', '2023-06-29 16:53:37', NULL);
INSERT INTO `sys_menus` VALUES (107, '删除', '', 100, 3, '', '', '', '', '', 'F', '', '', '', 'device:group:delete', '0', 'panda', '', '', '2023-06-29 16:53:56', '2023-06-29 16:53:56', NULL);
INSERT INTO `sys_menus` VALUES (114, '产品管理', '', 95, 2, 'elementCpu', '/device/product', '/device/product/index', '1', '', 'C', '0', '1', '1', 'device:product:list', '0', '', 'panda', '', '2023-06-30 14:13:39', '2023-07-21 16:03:31', NULL);
INSERT INTO `sys_menus` VALUES (115, '新增产品', '', 114, 1, '', '', '', '', '', 'F', '', '', '', 'device:product:add', '0', 'admin', '', '', '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL);
INSERT INTO `sys_menus` VALUES (116, '修改产品', '', 114, 2, '', '', '', '', '', 'F', '', '', '', 'device:product:edit', '0', 'admin', '', '', '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL);
INSERT INTO `sys_menus` VALUES (117, '删除产品', '', 114, 3, '', '', '', '', '', 'F', '', '', '', 'device:product:delete', '0', 'admin', '', '', '2023-06-30 14:13:39', '2023-06-30 14:13:39', NULL);
INSERT INTO `sys_menus` VALUES (118, '设备管理', '', 95, 4, 'elementSetting', '/device/device', '/device/device/index', '1', '', 'C', '0', '1', '1', 'device:device:list', '0', '', 'panda', '', '2023-06-30 14:20:03', '2023-07-21 16:03:41', NULL);
INSERT INTO `sys_menus` VALUES (119, '修改设备', '', 118, 2, '', '', '', '', '', 'F', '', '', '', 'device:device:edit', '0', 'admin', '', '', '2023-06-30 14:20:03', '2023-06-30 14:20:03', NULL);
INSERT INTO `sys_menus` VALUES (120, '新增设备', '', 118, 1, '', '', '', '', '', 'F', '', '', '', 'device:device:add', '0', 'admin', '', '', '2023-06-30 14:20:03', '2023-06-30 14:20:03', NULL);
INSERT INTO `sys_menus` VALUES (121, '删除设备', '', 118, 3, '', '', '', '', '', 'F', '', '', '', 'device:device:delete', '0', 'admin', '', '', '2023-06-30 14:20:03', '2023-06-30 14:20:03', NULL);
INSERT INTO `sys_menus` VALUES (122, '查看', '', 114, 4, '', '', '', '', '', 'F', '', '', '', 'device:product:view', '0', 'panda', '', '', '2023-07-05 17:14:20', '2023-07-05 17:14:20', NULL);
INSERT INTO `sys_menus` VALUES (131, '查看设备', '', 118, 4, '', '', '', '', '', 'F', '', '', '', 'device:device:view', '0', 'panda', '', '', '2023-07-10 08:50:48', '2023-07-10 08:50:48', NULL);
INSERT INTO `sys_menus` VALUES (132, '规则设计', '', 96, 1, 'iconfont icon-shuxingtu', '/rule/chain', '/rule/chain/index', '1', '', 'C', '0', '1', '1', 'rule:chain:list', '0', '', 'panda', '', '2023-07-21 14:38:54', '2023-09-02 14:33:03', NULL);
INSERT INTO `sys_menus` VALUES (133, '克隆', '', 132, 1, '', '', '', '', '', 'F', '', '', '', 'rule:chain:clone', '0', '', 'panda', '', '2023-07-21 14:39:27', '2023-07-21 14:57:05', NULL);
INSERT INTO `sys_menus` VALUES (134, '设计', '', 132, 2, '', '', '', '', '', 'F', '', '', '', 'rule:chain:design', '0', '', 'panda', '', '2023-07-21 14:39:53', '2023-07-21 14:57:13', NULL);
INSERT INTO `sys_menus` VALUES (135, '预览', '', 132, 3, '', '', '', '', '', 'F', '', '', '', 'rule:chain:view', '0', '', 'panda', '', '2023-07-21 14:40:08', '2023-07-21 14:57:20', NULL);
INSERT INTO `sys_menus` VALUES (136, '修改', '', 132, 4, '', '', '', '', '', 'F', '', '', '', 'rule:chain:edit', '0', '', 'panda', '', '2023-07-21 14:40:31', '2023-07-21 14:57:26', NULL);
INSERT INTO `sys_menus` VALUES (137, '删除', '', 132, 5, '', '', '', '', '', 'F', '', '', '', 'rule:chain:delete', '0', '', 'panda', '', '2023-07-21 14:40:47', '2023-07-21 14:57:33', NULL);
INSERT INTO `sys_menus` VALUES (138, '添加', '', 132, 6, '', '', '', '', '', 'F', '', '', '', 'rule:chain:add', '0', '', 'panda', '', '2023-07-21 14:41:04', '2023-07-21 14:57:39', NULL);
INSERT INTO `sys_menus` VALUES (139, '大屏分组', '', 97, 1, 'iconfont icon-wenducanshu-05', '/visual/screen_group', '/visual/screen_group/index', '1', '', 'C', '0', '1', '1', 'screen:group:list', '0', 'panda', '', '', '2023-07-21 14:46:41', '2023-07-21 14:46:41', NULL);
INSERT INTO `sys_menus` VALUES (140, '组态大屏', '', 97, 2, 'iconfont icon-diannaobangong', '/visual/screen', '/visual/screen/index', '1', '', 'C', '0', '1', '1', 'visual:screen:list', '0', 'panda', '', '', '2023-07-21 14:47:46', '2023-07-21 14:47:46', NULL);
INSERT INTO `sys_menus` VALUES (141, '添加', '', 139, 1, '', '', '', '', '', 'F', '', '', '', 'screen:group:add', '0', 'panda', '', '', '2023-07-21 14:50:40', '2023-07-21 14:50:40', NULL);
INSERT INTO `sys_menus` VALUES (142, '编辑', '', 139, 2, '', '', '', '', '', 'F', '', '', '', 'screen:group:edit', '0', 'panda', '', '', '2023-07-21 14:50:56', '2023-07-21 14:50:56', NULL);
INSERT INTO `sys_menus` VALUES (143, '删除', '', 139, 3, '', '', '', '', '', 'F', '', '', '', '	 screen:group:delete', '0', 'panda', '', '', '2023-07-21 14:51:22', '2023-07-21 14:51:22', NULL);
INSERT INTO `sys_menus` VALUES (144, '新增组态', '', 140, 1, '', '', '', '', '', 'F', '', '', '', 'visual:screen:add', '0', 'panda', '', '', '2023-07-21 14:53:26', '2023-07-21 14:53:26', NULL);
INSERT INTO `sys_menus` VALUES (145, '修改大屏', '', 140, 2, '', '', '', '', '', 'F', '', '', '', 'visual:screen:edit', '0', 'panda', '', '', '2023-07-21 14:53:50', '2023-07-21 14:53:50', NULL);
INSERT INTO `sys_menus` VALUES (146, '删除大屏', '', 140, 3, '', '', '', '', '', 'F', '', '', '', 'visual:screen:delete', '0', 'panda', '', '', '2023-07-21 14:54:14', '2023-07-21 14:54:14', NULL);
INSERT INTO `sys_menus` VALUES (147, '克隆', '', 140, 4, '', '', '', '', '', 'F', '', '', '', 'visual:screen:clone', '0', 'panda', '', '', '2023-07-21 14:54:30', '2023-07-21 14:54:30', NULL);
INSERT INTO `sys_menus` VALUES (148, '设计', '', 140, 5, '', '', '', '', '', 'F', '', '', '', 'visual:screen:design', '0', 'panda', '', '', '2023-07-21 14:54:57', '2023-07-21 14:54:57', NULL);
INSERT INTO `sys_menus` VALUES (149, '预览', '', 140, 6, '', '', '', '', '', 'F', '', '', '', 'visual:screen:view', '0', 'panda', '', '', '2023-07-21 14:55:27', '2023-07-21 14:55:27', NULL);
INSERT INTO `sys_menus` VALUES (150, '报表管理', '', 0, 4, 'iconfont icon-putong', '/report', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'panda', 'panda', '', '2023-07-24 10:12:26', '2023-07-24 10:13:54', '2023-08-18 14:01:16');
INSERT INTO `sys_menus` VALUES (151, '报表设计', '', 150, 1, 'iconfont icon-dayin', '/report', '/layout/routerView/iframes', '0', 'http://101.35.247.125:9001/edit', 'C', '0', '1', '1', 'report:list', '0', 'panda', '', '', '2023-07-24 10:13:47', '2023-07-24 10:13:47', '2023-08-18 14:01:12');
INSERT INTO `sys_menus` VALUES (152, '任务中心', '', 0, 5, 'iconfont icon-dayin', '/job', 'Layout', '1', '', 'M', '0', '1', '1', '', '0', 'panda', 'panda', '', '2023-08-08 14:08:11', '2023-10-05 14:03:38', NULL);
INSERT INTO `sys_menus` VALUES (153, '任务中心', '', 152, 1, 'elementAlarmClock', '/job/job', '/job/job/index', '1', '', 'C', '0', '1', '1', 'job:list', '0', '', 'panda', '', '2023-08-08 14:10:37', '2023-08-08 14:12:49', NULL);
INSERT INTO `sys_menus` VALUES (154, '任务日志', '', 152, 2, 'elementDocument', '/job/log', '/job/log/index', '1', '', 'C', '0', '1', '1', 'job:log:list', '0', 'panda', '', '', '2023-08-08 14:12:37', '2023-08-08 14:12:37', NULL);
INSERT INTO `sys_menus` VALUES (155, '新增', '', 153, 1, '', '', '', '', '', 'F', '', '', '', 'job:add', '0', 'panda', '', '', '2023-08-08 14:20:17', '2023-08-08 14:20:17', NULL);
INSERT INTO `sys_menus` VALUES (156, '编辑', '', 153, 2, '', '', '', '', '', 'F', '', '', '', 'job:edit', '0', 'panda', '', '', '2023-08-08 14:20:44', '2023-08-08 14:20:44', NULL);
INSERT INTO `sys_menus` VALUES (157, '删除', '', 153, 3, '', '', '', '', '', 'F', '', '', '', 'job:delete', '0', 'panda', '', '', '2023-08-08 14:21:03', '2023-08-08 14:21:03', NULL);
INSERT INTO `sys_menus` VALUES (158, '运行启动', '', 153, 4, '', '', '', '', '', 'F', '', '', '', 'job:run', '0', 'panda', '', '', '2023-08-08 14:21:25', '2023-08-08 14:21:25', NULL);
INSERT INTO `sys_menus` VALUES (159, '删除', '', 154, 1, '', '', '', '', '', 'F', '', '', '', 'job:log:delete', '0', 'panda', '', '', '2023-08-08 14:22:05', '2023-08-08 14:22:05', NULL);
INSERT INTO `sys_menus` VALUES (160, '清空', '', 154, 2, '', '', '', '', '', 'F', '', '', '', 'job:log:clean', '0', 'panda', '', '', '2023-08-08 14:22:33', '2023-08-08 14:22:33', NULL);
INSERT INTO `sys_menus` VALUES (161, '视频监控', '', 0, 4, 'iconfont icon-step', '/video', 'Layout', '1', '', 'M', '1', '1', '1', '', '1', 'panda', 'panda', '', '2023-09-02 13:52:17', '2023-10-24 14:05:25', NULL);
INSERT INTO `sys_menus` VALUES (164, '视频广场', '', 161, 3, 'elementGrid', '/video/splitview', '/video/splitview/index', '1', '', 'C', '0', '1', '1', 'video:splitview:list', '0', 'panda', '', '', '2023-09-02 14:00:41', '2023-09-02 14:00:41', NULL);
INSERT INTO `sys_menus` VALUES (165, '规则审计', '', 96, 2, 'iconfont icon--chaifenhang', '/rule/log', '/rule/log/index', '1', '', 'C', '0', '1', '1', 'rule:log:list', '0', 'panda', '', '', '2023-09-02 14:05:46', '2023-09-02 14:05:46', NULL);
INSERT INTO `sys_menus` VALUES (168, '设备地图', '', 95, 5, 'iconfont icon-ditu', 'device:map', '/device/map/index', '1', '', 'C', '0', '1', '1', 'device:map:list', '0', 'panda', '', '', '2023-09-02 14:14:00', '2023-09-02 14:14:00', NULL);
INSERT INTO `sys_menus` VALUES (169, '边缘管理', '', 0, 7, 'iconfont icon-wendu', '/edge', 'Layout', '1', '', 'M', '1', '1', '1', '', '1', 'panda', 'panda', '', '2023-09-02 14:27:39', '2023-10-24 14:05:34', NULL);
INSERT INTO `sys_menus` VALUES (170, '网关管理', '', 169, 1, 'iconfont icon-gouxuan-weixuanzhong-xianxingfangkuang', '/edge/gateway', '/edge/gateway/index', '1', '', 'C', '0', '1', '1', 'edge:gateway:list', '0', '', 'panda', '', '2023-09-02 14:44:13', '2023-09-19 10:20:34', NULL);
INSERT INTO `sys_menus` VALUES (171, '采集器', '', 169, 2, 'iconfont icon-wendu', '/edge/collector', '/edge/collector/index', '1', '', 'C', '0', '1', '1', 'edge:collector:list', '0', '', 'panda', '', '2023-09-02 14:45:31', '2023-09-19 10:20:57', NULL);
INSERT INTO `sys_menus` VALUES (172, '应用管理', '', 0, 8, 'iconfont icon-shoujidiannao', '/apply', 'Layout', '1', '', 'M', '1', '1', '1', '', '1', 'panda', 'panda', '', '2023-09-02 14:50:48', '2023-09-04 10:55:19', NULL);
INSERT INTO `sys_menus` VALUES (173, '应用商店', '', 172, 1, 'iconfont icon-shoujidiannao', '/apply/common', '/apply/common/index', '1', '', 'C', '0', '1', '1', 'apply:common:list', '0', 'panda', '', '', '2023-09-02 14:51:56', '2023-09-02 14:51:56', NULL);
INSERT INTO `sys_menus` VALUES (174, '我的应用', '', 172, 2, 'iconfont icon-LoggedinPC', '/apply/meapp', '/apply/meapp/index', '1', '', 'C', '0', '1', '1', 'apply:meapp:list', '0', 'panda', '', '', '2023-09-02 14:52:45', '2023-09-02 14:52:45', NULL);
INSERT INTO `sys_menus` VALUES (175, '萤石设备', '', 161, 2, 'iconfont icon-gerenzhongxin', '/video/ezviz', '/video/ezviz/index', '1', '', 'C', '0', '1', '1', 'video:ezviz:list', '0', 'panda', '', '', '2023-09-05 10:05:27', '2023-09-05 10:05:27', NULL);
INSERT INTO `sys_menus` VALUES (176, '国标设备', '', 161, 1, 'iconfont icon-wendu', '/video/gb28181', '/video/gb28181/index', '1', '', 'C', '0', '1', '1', 'video:gb28181:list', '0', 'panda', '', '', '2023-09-05 10:07:07', '2023-09-05 10:07:07', NULL);
INSERT INTO `sys_menus` VALUES (177, '分配设备', '', 118, 5, '', '', '', '', '', 'F', '', '', '', 'device:device:allot', '0', 'panda', '', '', '2023-09-15 17:32:08', '2023-09-15 17:32:08', NULL);
INSERT INTO `sys_menus` VALUES (178, '告警中心', '', 44, 1, 'iconfont icon-radio-off-full', '/tool/alarm', '/tool/alarm/index', '1', '', 'C', '0', '1', '1', 'tool:alarm:list', '0', 'panda', '', '', '2023-10-18 10:20:51', '2023-10-18 10:20:51', NULL);

-- ----------------------------
-- Table structure for sys_notices
-- ----------------------------
DROP TABLE IF EXISTS `sys_notices`;
CREATE TABLE `sys_notices`  (
  `notice_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '标题',
  `notice_type` varchar(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '通知类型',
  `organization_id` int(0) NULL DEFAULT NULL COMMENT '部门Id,部门及子部门',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  `user_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_notices
-- ----------------------------
INSERT INTO `sys_notices` VALUES (1, '关于学习交流的通知', '<p>发布<b>入群</b>通知&nbsp;<span style=\"color: var(--el-text-color-regular);\">467890197, 交流学习</span><span style=\"font-size: 14px; color: var(--el-text-color-regular);\"></span></p>', '1', 0, '2021-12-26 15:29:25', '2021-12-26 16:19:48', NULL, 'panda');

-- ----------------------------
-- Table structure for sys_organizations
-- ----------------------------
DROP TABLE IF EXISTS `sys_organizations`;
CREATE TABLE `sys_organizations`  (
  `organization_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `parent_id` int(0) NULL DEFAULT NULL COMMENT '上级组织',
  `organization_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织路径',
  `organization_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织名称',
  `sort` int(0) NULL DEFAULT NULL COMMENT '排序',
  `leader` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '修改人',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`organization_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_organizations
-- ----------------------------
INSERT INTO `sys_organizations` VALUES (2, 0, '/0/2', '熊猫科技', 0, 'xm', '18353366836', '342@qq.com', '0', 'admin', 'admin', '2021-12-01 17:31:53', '2021-12-02 08:56:19', NULL);
INSERT INTO `sys_organizations` VALUES (3, 2, '/0/2/3', '研发部', 1, 'panda', '18353366543', 'ewr@qq.com', '0', 'admin', 'admin', '2021-12-01 17:37:43', '2021-12-02 08:55:56', NULL);
INSERT INTO `sys_organizations` VALUES (7, 2, '/0/2/7', '营销部', 2, 'panda', '18353333333', '342@qq.com', '0', 'panda', 'panda', '2021-12-24 10:46:24', '2021-12-24 10:47:15', NULL);

-- ----------------------------
-- Table structure for sys_posts
-- ----------------------------
DROP TABLE IF EXISTS `sys_posts`;
CREATE TABLE `sys_posts`  (
  `post_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位名称',
  `post_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位代码',
  `sort` int(0) NULL DEFAULT NULL COMMENT '岗位排序',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_posts
-- ----------------------------
INSERT INTO `sys_posts` VALUES (1, '首席执行官', 'CEO', 1, '0', '首席执行官', 'panda', '', '2021-12-02 09:21:44', '2022-07-16 17:36:32', NULL);
INSERT INTO `sys_posts` VALUES (4, '首席技术执行官', 'CTO', 2, '0', '', 'panda', '', '2021-12-02 09:21:44', '2022-07-16 17:37:42', NULL);

-- ----------------------------
-- Table structure for sys_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menus`;
CREATE TABLE `sys_role_menus`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `role_id` int(0) NULL DEFAULT NULL,
  `menu_id` int(0) NULL DEFAULT NULL,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8253 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menus
-- ----------------------------
INSERT INTO `sys_role_menus` VALUES (6590, 5, 1, 'test');
INSERT INTO `sys_role_menus` VALUES (6591, 5, 3, 'test');
INSERT INTO `sys_role_menus` VALUES (6592, 5, 4, 'test');
INSERT INTO `sys_role_menus` VALUES (6593, 5, 5, 'test');
INSERT INTO `sys_role_menus` VALUES (6594, 5, 6, 'test');
INSERT INTO `sys_role_menus` VALUES (6595, 5, 7, 'test');
INSERT INTO `sys_role_menus` VALUES (6596, 5, 8, 'test');
INSERT INTO `sys_role_menus` VALUES (6597, 5, 9, 'test');
INSERT INTO `sys_role_menus` VALUES (6598, 5, 10, 'test');
INSERT INTO `sys_role_menus` VALUES (6599, 5, 11, 'test');
INSERT INTO `sys_role_menus` VALUES (6600, 5, 13, 'test');
INSERT INTO `sys_role_menus` VALUES (6601, 5, 14, 'test');
INSERT INTO `sys_role_menus` VALUES (6602, 5, 15, 'test');
INSERT INTO `sys_role_menus` VALUES (6603, 5, 16, 'test');
INSERT INTO `sys_role_menus` VALUES (6604, 5, 17, 'test');
INSERT INTO `sys_role_menus` VALUES (6605, 5, 18, 'test');
INSERT INTO `sys_role_menus` VALUES (6606, 5, 19, 'test');
INSERT INTO `sys_role_menus` VALUES (6607, 5, 20, 'test');
INSERT INTO `sys_role_menus` VALUES (6608, 5, 21, 'test');
INSERT INTO `sys_role_menus` VALUES (6609, 5, 22, 'test');
INSERT INTO `sys_role_menus` VALUES (6610, 5, 23, 'test');
INSERT INTO `sys_role_menus` VALUES (6611, 5, 24, 'test');
INSERT INTO `sys_role_menus` VALUES (6612, 5, 25, 'test');
INSERT INTO `sys_role_menus` VALUES (6613, 5, 26, 'test');
INSERT INTO `sys_role_menus` VALUES (6614, 5, 28, 'test');
INSERT INTO `sys_role_menus` VALUES (6615, 5, 29, 'test');
INSERT INTO `sys_role_menus` VALUES (6616, 5, 30, 'test');
INSERT INTO `sys_role_menus` VALUES (6617, 5, 31, 'test');
INSERT INTO `sys_role_menus` VALUES (6618, 5, 32, 'test');
INSERT INTO `sys_role_menus` VALUES (6619, 5, 33, 'test');
INSERT INTO `sys_role_menus` VALUES (6620, 5, 34, 'test');
INSERT INTO `sys_role_menus` VALUES (6621, 5, 35, 'test');
INSERT INTO `sys_role_menus` VALUES (6622, 5, 36, 'test');
INSERT INTO `sys_role_menus` VALUES (6623, 5, 37, 'test');
INSERT INTO `sys_role_menus` VALUES (6624, 5, 38, 'test');
INSERT INTO `sys_role_menus` VALUES (6625, 5, 39, 'test');
INSERT INTO `sys_role_menus` VALUES (6626, 5, 40, 'test');
INSERT INTO `sys_role_menus` VALUES (6627, 5, 41, 'test');
INSERT INTO `sys_role_menus` VALUES (6628, 5, 42, 'test');
INSERT INTO `sys_role_menus` VALUES (7878, 2, 1, 'manage');
INSERT INTO `sys_role_menus` VALUES (7879, 2, 3, 'manage');
INSERT INTO `sys_role_menus` VALUES (7880, 2, 4, 'manage');
INSERT INTO `sys_role_menus` VALUES (7881, 2, 5, 'manage');
INSERT INTO `sys_role_menus` VALUES (7882, 2, 6, 'manage');
INSERT INTO `sys_role_menus` VALUES (7883, 2, 7, 'manage');
INSERT INTO `sys_role_menus` VALUES (7884, 2, 8, 'manage');
INSERT INTO `sys_role_menus` VALUES (7885, 2, 9, 'manage');
INSERT INTO `sys_role_menus` VALUES (7886, 2, 10, 'manage');
INSERT INTO `sys_role_menus` VALUES (7887, 2, 11, 'manage');
INSERT INTO `sys_role_menus` VALUES (7888, 2, 12, 'manage');
INSERT INTO `sys_role_menus` VALUES (7889, 2, 13, 'manage');
INSERT INTO `sys_role_menus` VALUES (7890, 2, 14, 'manage');
INSERT INTO `sys_role_menus` VALUES (7891, 2, 15, 'manage');
INSERT INTO `sys_role_menus` VALUES (7892, 2, 16, 'manage');
INSERT INTO `sys_role_menus` VALUES (7893, 2, 17, 'manage');
INSERT INTO `sys_role_menus` VALUES (7894, 2, 18, 'manage');
INSERT INTO `sys_role_menus` VALUES (7895, 2, 19, 'manage');
INSERT INTO `sys_role_menus` VALUES (7896, 2, 20, 'manage');
INSERT INTO `sys_role_menus` VALUES (7897, 2, 21, 'manage');
INSERT INTO `sys_role_menus` VALUES (7898, 2, 22, 'manage');
INSERT INTO `sys_role_menus` VALUES (7899, 2, 23, 'manage');
INSERT INTO `sys_role_menus` VALUES (7900, 2, 24, 'manage');
INSERT INTO `sys_role_menus` VALUES (7901, 2, 25, 'manage');
INSERT INTO `sys_role_menus` VALUES (7902, 2, 26, 'manage');
INSERT INTO `sys_role_menus` VALUES (7903, 2, 28, 'manage');
INSERT INTO `sys_role_menus` VALUES (7904, 2, 29, 'manage');
INSERT INTO `sys_role_menus` VALUES (7905, 2, 30, 'manage');
INSERT INTO `sys_role_menus` VALUES (7906, 2, 31, 'manage');
INSERT INTO `sys_role_menus` VALUES (7907, 2, 32, 'manage');
INSERT INTO `sys_role_menus` VALUES (7908, 2, 33, 'manage');
INSERT INTO `sys_role_menus` VALUES (7909, 2, 34, 'manage');
INSERT INTO `sys_role_menus` VALUES (7910, 2, 35, 'manage');
INSERT INTO `sys_role_menus` VALUES (7911, 2, 36, 'manage');
INSERT INTO `sys_role_menus` VALUES (7912, 2, 37, 'manage');
INSERT INTO `sys_role_menus` VALUES (7913, 2, 38, 'manage');
INSERT INTO `sys_role_menus` VALUES (7914, 2, 39, 'manage');
INSERT INTO `sys_role_menus` VALUES (7915, 2, 40, 'manage');
INSERT INTO `sys_role_menus` VALUES (7916, 2, 41, 'manage');
INSERT INTO `sys_role_menus` VALUES (7917, 2, 42, 'manage');
INSERT INTO `sys_role_menus` VALUES (7918, 2, 43, 'manage');
INSERT INTO `sys_role_menus` VALUES (7919, 2, 44, 'manage');
INSERT INTO `sys_role_menus` VALUES (7920, 2, 45, 'manage');
INSERT INTO `sys_role_menus` VALUES (7921, 2, 46, 'manage');
INSERT INTO `sys_role_menus` VALUES (7922, 2, 47, 'manage');
INSERT INTO `sys_role_menus` VALUES (7923, 2, 49, 'manage');
INSERT INTO `sys_role_menus` VALUES (7924, 2, 50, 'manage');
INSERT INTO `sys_role_menus` VALUES (7925, 2, 51, 'manage');
INSERT INTO `sys_role_menus` VALUES (7926, 2, 52, 'manage');
INSERT INTO `sys_role_menus` VALUES (7927, 2, 55, 'manage');
INSERT INTO `sys_role_menus` VALUES (7928, 2, 59, 'manage');
INSERT INTO `sys_role_menus` VALUES (7929, 2, 60, 'manage');
INSERT INTO `sys_role_menus` VALUES (7930, 2, 63, 'manage');
INSERT INTO `sys_role_menus` VALUES (7931, 2, 64, 'manage');
INSERT INTO `sys_role_menus` VALUES (7932, 2, 69, 'manage');
INSERT INTO `sys_role_menus` VALUES (7933, 2, 70, 'manage');
INSERT INTO `sys_role_menus` VALUES (7934, 2, 71, 'manage');
INSERT INTO `sys_role_menus` VALUES (7935, 2, 72, 'manage');
INSERT INTO `sys_role_menus` VALUES (7936, 2, 73, 'manage');
INSERT INTO `sys_role_menus` VALUES (7937, 2, 74, 'manage');
INSERT INTO `sys_role_menus` VALUES (7938, 2, 75, 'manage');
INSERT INTO `sys_role_menus` VALUES (7939, 2, 76, 'manage');
INSERT INTO `sys_role_menus` VALUES (7940, 2, 77, 'manage');
INSERT INTO `sys_role_menus` VALUES (7941, 2, 95, 'manage');
INSERT INTO `sys_role_menus` VALUES (7942, 2, 96, 'manage');
INSERT INTO `sys_role_menus` VALUES (7943, 2, 97, 'manage');
INSERT INTO `sys_role_menus` VALUES (7944, 2, 98, 'manage');
INSERT INTO `sys_role_menus` VALUES (7945, 2, 100, 'manage');
INSERT INTO `sys_role_menus` VALUES (7946, 2, 102, 'manage');
INSERT INTO `sys_role_menus` VALUES (7947, 2, 103, 'manage');
INSERT INTO `sys_role_menus` VALUES (7948, 2, 104, 'manage');
INSERT INTO `sys_role_menus` VALUES (7949, 2, 105, 'manage');
INSERT INTO `sys_role_menus` VALUES (7950, 2, 106, 'manage');
INSERT INTO `sys_role_menus` VALUES (7951, 2, 107, 'manage');
INSERT INTO `sys_role_menus` VALUES (7952, 2, 114, 'manage');
INSERT INTO `sys_role_menus` VALUES (7953, 2, 115, 'manage');
INSERT INTO `sys_role_menus` VALUES (7954, 2, 116, 'manage');
INSERT INTO `sys_role_menus` VALUES (7955, 2, 117, 'manage');
INSERT INTO `sys_role_menus` VALUES (7956, 2, 118, 'manage');
INSERT INTO `sys_role_menus` VALUES (7957, 2, 119, 'manage');
INSERT INTO `sys_role_menus` VALUES (7958, 2, 120, 'manage');
INSERT INTO `sys_role_menus` VALUES (7959, 2, 121, 'manage');
INSERT INTO `sys_role_menus` VALUES (7960, 2, 122, 'manage');
INSERT INTO `sys_role_menus` VALUES (7961, 2, 131, 'manage');
INSERT INTO `sys_role_menus` VALUES (7962, 2, 132, 'manage');
INSERT INTO `sys_role_menus` VALUES (7963, 2, 133, 'manage');
INSERT INTO `sys_role_menus` VALUES (7964, 2, 134, 'manage');
INSERT INTO `sys_role_menus` VALUES (7965, 2, 135, 'manage');
INSERT INTO `sys_role_menus` VALUES (7966, 2, 136, 'manage');
INSERT INTO `sys_role_menus` VALUES (7967, 2, 137, 'manage');
INSERT INTO `sys_role_menus` VALUES (7968, 2, 138, 'manage');
INSERT INTO `sys_role_menus` VALUES (7969, 2, 139, 'manage');
INSERT INTO `sys_role_menus` VALUES (7970, 2, 140, 'manage');
INSERT INTO `sys_role_menus` VALUES (7971, 2, 141, 'manage');
INSERT INTO `sys_role_menus` VALUES (7972, 2, 142, 'manage');
INSERT INTO `sys_role_menus` VALUES (7973, 2, 143, 'manage');
INSERT INTO `sys_role_menus` VALUES (7974, 2, 144, 'manage');
INSERT INTO `sys_role_menus` VALUES (7975, 2, 145, 'manage');
INSERT INTO `sys_role_menus` VALUES (7976, 2, 146, 'manage');
INSERT INTO `sys_role_menus` VALUES (7977, 2, 147, 'manage');
INSERT INTO `sys_role_menus` VALUES (7978, 2, 148, 'manage');
INSERT INTO `sys_role_menus` VALUES (7979, 2, 149, 'manage');
INSERT INTO `sys_role_menus` VALUES (7980, 2, 152, 'manage');
INSERT INTO `sys_role_menus` VALUES (7981, 2, 153, 'manage');
INSERT INTO `sys_role_menus` VALUES (7982, 2, 154, 'manage');
INSERT INTO `sys_role_menus` VALUES (7983, 2, 155, 'manage');
INSERT INTO `sys_role_menus` VALUES (7984, 2, 156, 'manage');
INSERT INTO `sys_role_menus` VALUES (7985, 2, 157, 'manage');
INSERT INTO `sys_role_menus` VALUES (7986, 2, 158, 'manage');
INSERT INTO `sys_role_menus` VALUES (7987, 2, 159, 'manage');
INSERT INTO `sys_role_menus` VALUES (7988, 2, 160, 'manage');
INSERT INTO `sys_role_menus` VALUES (7989, 2, 161, 'manage');
INSERT INTO `sys_role_menus` VALUES (7990, 2, 164, 'manage');
INSERT INTO `sys_role_menus` VALUES (7991, 2, 165, 'manage');
INSERT INTO `sys_role_menus` VALUES (7992, 2, 166, 'manage');
INSERT INTO `sys_role_menus` VALUES (7993, 2, 167, 'manage');
INSERT INTO `sys_role_menus` VALUES (7994, 2, 168, 'manage');
INSERT INTO `sys_role_menus` VALUES (7995, 2, 169, 'manage');
INSERT INTO `sys_role_menus` VALUES (7996, 2, 170, 'manage');
INSERT INTO `sys_role_menus` VALUES (7997, 2, 171, 'manage');
INSERT INTO `sys_role_menus` VALUES (7998, 2, 175, 'manage');
INSERT INTO `sys_role_menus` VALUES (7999, 2, 176, 'manage');
INSERT INTO `sys_role_menus` VALUES (8000, 2, 177, 'manage');
INSERT INTO `sys_role_menus` VALUES (8128, 1, 1, 'admin');
INSERT INTO `sys_role_menus` VALUES (8129, 1, 3, 'admin');
INSERT INTO `sys_role_menus` VALUES (8130, 1, 4, 'admin');
INSERT INTO `sys_role_menus` VALUES (8131, 1, 5, 'admin');
INSERT INTO `sys_role_menus` VALUES (8132, 1, 6, 'admin');
INSERT INTO `sys_role_menus` VALUES (8133, 1, 7, 'admin');
INSERT INTO `sys_role_menus` VALUES (8134, 1, 8, 'admin');
INSERT INTO `sys_role_menus` VALUES (8135, 1, 9, 'admin');
INSERT INTO `sys_role_menus` VALUES (8136, 1, 10, 'admin');
INSERT INTO `sys_role_menus` VALUES (8137, 1, 11, 'admin');
INSERT INTO `sys_role_menus` VALUES (8138, 1, 12, 'admin');
INSERT INTO `sys_role_menus` VALUES (8139, 1, 13, 'admin');
INSERT INTO `sys_role_menus` VALUES (8140, 1, 14, 'admin');
INSERT INTO `sys_role_menus` VALUES (8141, 1, 15, 'admin');
INSERT INTO `sys_role_menus` VALUES (8142, 1, 16, 'admin');
INSERT INTO `sys_role_menus` VALUES (8143, 1, 17, 'admin');
INSERT INTO `sys_role_menus` VALUES (8144, 1, 18, 'admin');
INSERT INTO `sys_role_menus` VALUES (8145, 1, 19, 'admin');
INSERT INTO `sys_role_menus` VALUES (8146, 1, 20, 'admin');
INSERT INTO `sys_role_menus` VALUES (8147, 1, 21, 'admin');
INSERT INTO `sys_role_menus` VALUES (8148, 1, 22, 'admin');
INSERT INTO `sys_role_menus` VALUES (8149, 1, 23, 'admin');
INSERT INTO `sys_role_menus` VALUES (8150, 1, 24, 'admin');
INSERT INTO `sys_role_menus` VALUES (8151, 1, 25, 'admin');
INSERT INTO `sys_role_menus` VALUES (8152, 1, 26, 'admin');
INSERT INTO `sys_role_menus` VALUES (8153, 1, 28, 'admin');
INSERT INTO `sys_role_menus` VALUES (8154, 1, 29, 'admin');
INSERT INTO `sys_role_menus` VALUES (8155, 1, 30, 'admin');
INSERT INTO `sys_role_menus` VALUES (8156, 1, 31, 'admin');
INSERT INTO `sys_role_menus` VALUES (8157, 1, 32, 'admin');
INSERT INTO `sys_role_menus` VALUES (8158, 1, 33, 'admin');
INSERT INTO `sys_role_menus` VALUES (8159, 1, 34, 'admin');
INSERT INTO `sys_role_menus` VALUES (8160, 1, 35, 'admin');
INSERT INTO `sys_role_menus` VALUES (8161, 1, 36, 'admin');
INSERT INTO `sys_role_menus` VALUES (8162, 1, 37, 'admin');
INSERT INTO `sys_role_menus` VALUES (8163, 1, 38, 'admin');
INSERT INTO `sys_role_menus` VALUES (8164, 1, 39, 'admin');
INSERT INTO `sys_role_menus` VALUES (8165, 1, 40, 'admin');
INSERT INTO `sys_role_menus` VALUES (8166, 1, 41, 'admin');
INSERT INTO `sys_role_menus` VALUES (8167, 1, 42, 'admin');
INSERT INTO `sys_role_menus` VALUES (8168, 1, 43, 'admin');
INSERT INTO `sys_role_menus` VALUES (8169, 1, 44, 'admin');
INSERT INTO `sys_role_menus` VALUES (8170, 1, 45, 'admin');
INSERT INTO `sys_role_menus` VALUES (8171, 1, 46, 'admin');
INSERT INTO `sys_role_menus` VALUES (8172, 1, 47, 'admin');
INSERT INTO `sys_role_menus` VALUES (8173, 1, 49, 'admin');
INSERT INTO `sys_role_menus` VALUES (8174, 1, 50, 'admin');
INSERT INTO `sys_role_menus` VALUES (8175, 1, 51, 'admin');
INSERT INTO `sys_role_menus` VALUES (8176, 1, 52, 'admin');
INSERT INTO `sys_role_menus` VALUES (8177, 1, 55, 'admin');
INSERT INTO `sys_role_menus` VALUES (8178, 1, 59, 'admin');
INSERT INTO `sys_role_menus` VALUES (8179, 1, 60, 'admin');
INSERT INTO `sys_role_menus` VALUES (8180, 1, 63, 'admin');
INSERT INTO `sys_role_menus` VALUES (8181, 1, 64, 'admin');
INSERT INTO `sys_role_menus` VALUES (8182, 1, 69, 'admin');
INSERT INTO `sys_role_menus` VALUES (8183, 1, 70, 'admin');
INSERT INTO `sys_role_menus` VALUES (8184, 1, 71, 'admin');
INSERT INTO `sys_role_menus` VALUES (8185, 1, 72, 'admin');
INSERT INTO `sys_role_menus` VALUES (8186, 1, 73, 'admin');
INSERT INTO `sys_role_menus` VALUES (8187, 1, 74, 'admin');
INSERT INTO `sys_role_menus` VALUES (8188, 1, 75, 'admin');
INSERT INTO `sys_role_menus` VALUES (8189, 1, 76, 'admin');
INSERT INTO `sys_role_menus` VALUES (8190, 1, 77, 'admin');
INSERT INTO `sys_role_menus` VALUES (8191, 1, 95, 'admin');
INSERT INTO `sys_role_menus` VALUES (8192, 1, 96, 'admin');
INSERT INTO `sys_role_menus` VALUES (8193, 1, 97, 'admin');
INSERT INTO `sys_role_menus` VALUES (8194, 1, 98, 'admin');
INSERT INTO `sys_role_menus` VALUES (8195, 1, 100, 'admin');
INSERT INTO `sys_role_menus` VALUES (8196, 1, 102, 'admin');
INSERT INTO `sys_role_menus` VALUES (8197, 1, 103, 'admin');
INSERT INTO `sys_role_menus` VALUES (8198, 1, 104, 'admin');
INSERT INTO `sys_role_menus` VALUES (8199, 1, 105, 'admin');
INSERT INTO `sys_role_menus` VALUES (8200, 1, 106, 'admin');
INSERT INTO `sys_role_menus` VALUES (8201, 1, 107, 'admin');
INSERT INTO `sys_role_menus` VALUES (8202, 1, 114, 'admin');
INSERT INTO `sys_role_menus` VALUES (8203, 1, 115, 'admin');
INSERT INTO `sys_role_menus` VALUES (8204, 1, 116, 'admin');
INSERT INTO `sys_role_menus` VALUES (8205, 1, 117, 'admin');
INSERT INTO `sys_role_menus` VALUES (8206, 1, 118, 'admin');
INSERT INTO `sys_role_menus` VALUES (8207, 1, 119, 'admin');
INSERT INTO `sys_role_menus` VALUES (8208, 1, 120, 'admin');
INSERT INTO `sys_role_menus` VALUES (8209, 1, 121, 'admin');
INSERT INTO `sys_role_menus` VALUES (8210, 1, 122, 'admin');
INSERT INTO `sys_role_menus` VALUES (8211, 1, 131, 'admin');
INSERT INTO `sys_role_menus` VALUES (8212, 1, 132, 'admin');
INSERT INTO `sys_role_menus` VALUES (8213, 1, 133, 'admin');
INSERT INTO `sys_role_menus` VALUES (8214, 1, 134, 'admin');
INSERT INTO `sys_role_menus` VALUES (8215, 1, 135, 'admin');
INSERT INTO `sys_role_menus` VALUES (8216, 1, 136, 'admin');
INSERT INTO `sys_role_menus` VALUES (8217, 1, 137, 'admin');
INSERT INTO `sys_role_menus` VALUES (8218, 1, 138, 'admin');
INSERT INTO `sys_role_menus` VALUES (8219, 1, 139, 'admin');
INSERT INTO `sys_role_menus` VALUES (8220, 1, 140, 'admin');
INSERT INTO `sys_role_menus` VALUES (8221, 1, 141, 'admin');
INSERT INTO `sys_role_menus` VALUES (8222, 1, 142, 'admin');
INSERT INTO `sys_role_menus` VALUES (8223, 1, 143, 'admin');
INSERT INTO `sys_role_menus` VALUES (8224, 1, 144, 'admin');
INSERT INTO `sys_role_menus` VALUES (8225, 1, 145, 'admin');
INSERT INTO `sys_role_menus` VALUES (8226, 1, 146, 'admin');
INSERT INTO `sys_role_menus` VALUES (8227, 1, 147, 'admin');
INSERT INTO `sys_role_menus` VALUES (8228, 1, 148, 'admin');
INSERT INTO `sys_role_menus` VALUES (8229, 1, 149, 'admin');
INSERT INTO `sys_role_menus` VALUES (8230, 1, 152, 'admin');
INSERT INTO `sys_role_menus` VALUES (8231, 1, 153, 'admin');
INSERT INTO `sys_role_menus` VALUES (8232, 1, 154, 'admin');
INSERT INTO `sys_role_menus` VALUES (8233, 1, 155, 'admin');
INSERT INTO `sys_role_menus` VALUES (8234, 1, 156, 'admin');
INSERT INTO `sys_role_menus` VALUES (8235, 1, 157, 'admin');
INSERT INTO `sys_role_menus` VALUES (8236, 1, 158, 'admin');
INSERT INTO `sys_role_menus` VALUES (8237, 1, 159, 'admin');
INSERT INTO `sys_role_menus` VALUES (8238, 1, 160, 'admin');
INSERT INTO `sys_role_menus` VALUES (8239, 1, 161, 'admin');
INSERT INTO `sys_role_menus` VALUES (8240, 1, 164, 'admin');
INSERT INTO `sys_role_menus` VALUES (8241, 1, 165, 'admin');
INSERT INTO `sys_role_menus` VALUES (8242, 1, 168, 'admin');
INSERT INTO `sys_role_menus` VALUES (8243, 1, 169, 'admin');
INSERT INTO `sys_role_menus` VALUES (8244, 1, 170, 'admin');
INSERT INTO `sys_role_menus` VALUES (8245, 1, 171, 'admin');
INSERT INTO `sys_role_menus` VALUES (8246, 1, 172, 'admin');
INSERT INTO `sys_role_menus` VALUES (8247, 1, 173, 'admin');
INSERT INTO `sys_role_menus` VALUES (8248, 1, 174, 'admin');
INSERT INTO `sys_role_menus` VALUES (8249, 1, 175, 'admin');
INSERT INTO `sys_role_menus` VALUES (8250, 1, 176, 'admin');
INSERT INTO `sys_role_menus` VALUES (8251, 1, 177, 'admin');
INSERT INTO `sys_role_menus` VALUES (8252, 1, 178, 'admin');

-- ----------------------------
-- Table structure for sys_role_organizations
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_organizations`;
CREATE TABLE `sys_role_organizations`  (
  `role_id` int(0) NULL DEFAULT NULL,
  `organization_id` int(0) NULL DEFAULT NULL,
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_organizations
-- ----------------------------
INSERT INTO `sys_role_organizations` VALUES (1, 2, 9);
INSERT INTO `sys_role_organizations` VALUES (1, 3, 10);
INSERT INTO `sys_role_organizations` VALUES (1, 7, 11);
INSERT INTO `sys_role_organizations` VALUES (2, 2, 12);
INSERT INTO `sys_role_organizations` VALUES (2, 3, 13);
INSERT INTO `sys_role_organizations` VALUES (2, 7, 14);

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `role_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '状态',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色代码',
  `data_scope` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `role_sort` int(0) NULL DEFAULT NULL COMMENT '角色排序',
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO `sys_roles` VALUES (1, '超管理员', '0', 'admin', '4', 1, 'admin', 'panda', '超级管理', '2021-12-02 16:03:26', '2023-10-18 10:23:08', NULL);
INSERT INTO `sys_roles` VALUES (2, '管理员', '0', 'manage', '1', 2, 'panda', 'panda', '', '2021-12-19 16:06:20', '2023-10-07 11:37:58', NULL);
INSERT INTO `sys_roles` VALUES (5, '测试', '0', 'test', '0', 3, 'panda', '', '', '2023-09-14 11:34:44', '2023-09-14 11:34:44', NULL);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `user_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role_id` int(0) NULL DEFAULT NULL,
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `organization_id` int(0) NULL DEFAULT NULL,
  `post_id` int(0) NULL DEFAULT NULL,
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `delete_time` datetime(0) NULL DEFAULT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '多角色',
  `post_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '多岗位',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, 'pandax', '13818888888', 1, NULL, '', '0', '1@qq.com', 2, 4, 'panda', '1', NULL, '0', '2021-12-03 09:46:55', '2022-02-09 13:28:49', NULL, 'panda', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '1', '1,4');
INSERT INTO `sys_users` VALUES (3, '测试用户', '18435234356', 2, '', '', '0', '342@163.com', 3, 1, 'test', '', '', '0', '2021-12-06 15:16:53', '2022-05-10 19:19:25', NULL, 'test', '$2a$10$4cHTracxWJLdhMmazvbm1urKyD3v5N2AYxAFtNYh6juU39kgae73e', '2', '1,4');
INSERT INTO `sys_users` VALUES (4, 'panda', '18353366912', 2, '', '', '0', '2417920382@qq.com', 2, 4, 'panda', '', '', '0', '2021-12-19 15:58:09', '2021-12-19 16:06:54', NULL, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2', '2', '4,1');
INSERT INTO `sys_users` VALUES (5, 'tenant', '', 1, '', '', '0', '', 3, 1, 'panda', '1', '', '0', '2021-12-03 09:46:55', '2022-02-09 13:28:49', NULL, 'tenant', '$2a$10$ycRsRdsrNQInLB2Ib0maOetsWZ0kFctmF6ytAErWTjOx5cWdeJMcK', '1', '1,4');

SET FOREIGN_KEY_CHECKS = 1;
