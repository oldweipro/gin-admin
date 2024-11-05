/*
 Navicat Premium Dump SQL

 Source Server         : 192.168.1.65_3306
 Source Server Type    : MySQL
 Source Server Version : 90100 (9.1.0)
 Source Host           : 192.168.1.65:3306
 Source Schema         : gin_admin

 Target Server Type    : MySQL
 Target Server Version : 90100 (9.1.0)
 File Encoding         : 65001

 Date: 05/11/2024 18:03:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for announcements_info
-- ----------------------------
DROP TABLE IF EXISTS `announcements_info`;
CREATE TABLE `announcements_info`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公告标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '公告内容',
  `user_id` bigint NULL DEFAULT NULL COMMENT '发布者',
  `attachments` json NULL COMMENT '相关附件',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_gva_announcements_info_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_announcements_info_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of announcements_info
-- ----------------------------
INSERT INTO `announcements_info` VALUES (1, '2024-10-09 16:10:56.782', '2024-10-09 16:10:56.782', NULL, '1', '<p>11</p>', 1, '[]');

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2255 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (2242, 'p', '1', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2241, 'p', '1', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2236, 'p', '1', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2233, 'p', '1', '/api/enterSyncApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2238, 'p', '1', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2237, 'p', '1', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2234, 'p', '1', '/api/getApiGroups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2239, 'p', '1', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2232, 'p', '1', '/api/ignoreApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2235, 'p', '1', '/api/syncApi', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2240, 'p', '1', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2231, 'p', '1', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2230, 'p', '1', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2229, 'p', '1', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2227, 'p', '1', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2226, 'p', '1', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2228, 'p', '1', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2162, 'p', '1', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2163, 'p', '1', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2164, 'p', '1', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2182, 'p', '1', '/autoCode/addFunc', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2190, 'p', '1', '/autoCode/createPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2195, 'p', '1', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2187, 'p', '1', '/autoCode/delPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2183, 'p', '1', '/autoCode/delSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2193, 'p', '1', '/autoCode/getColumn', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2197, 'p', '1', '/autoCode/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2186, 'p', '1', '/autoCode/getMeta', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2188, 'p', '1', '/autoCode/getPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2184, 'p', '1', '/autoCode/getSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2196, 'p', '1', '/autoCode/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2189, 'p', '1', '/autoCode/getTemplates', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2192, 'p', '1', '/autoCode/installPlugin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2194, 'p', '1', '/autoCode/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2191, 'p', '1', '/autoCode/pubPlug', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2185, 'p', '1', '/autoCode/rollback', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2224, 'p', '1', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2225, 'p', '1', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2200, 'p', '1', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2199, 'p', '1', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2201, 'p', '1', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2202, 'p', '1', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2198, 'p', '1', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2213, 'p', '1', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2212, 'p', '1', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2209, 'p', '1', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2208, 'p', '1', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2214, 'p', '1', '/fileUploadAndDownload/findFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2207, 'p', '1', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2206, 'p', '1', '/fileUploadAndDownload/importURL', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2211, 'p', '1', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2210, 'p', '1', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2254, 'p', '1', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2223, 'p', '1', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2215, 'p', '1', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2221, 'p', '1', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2219, 'p', '1', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2217, 'p', '1', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2222, 'p', '1', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2216, 'p', '1', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2218, 'p', '1', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2220, 'p', '1', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2176, 'p', '1', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2175, 'p', '1', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2173, 'p', '1', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2172, 'p', '1', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2174, 'p', '1', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2180, 'p', '1', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2179, 'p', '1', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2178, 'p', '1', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2177, 'p', '1', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2181, 'p', '1', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2161, 'p', '1', '/sysExportTemplate/createSysExportTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2160, 'p', '1', '/sysExportTemplate/deleteSysExportTemplate', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2159, 'p', '1', '/sysExportTemplate/deleteSysExportTemplateByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2155, 'p', '1', '/sysExportTemplate/exportExcel', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2154, 'p', '1', '/sysExportTemplate/exportTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2157, 'p', '1', '/sysExportTemplate/findSysExportTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2156, 'p', '1', '/sysExportTemplate/getSysExportTemplateList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2153, 'p', '1', '/sysExportTemplate/importExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2158, 'p', '1', '/sysExportTemplate/updateSysExportTemplate', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2171, 'p', '1', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2168, 'p', '1', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2167, 'p', '1', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2170, 'p', '1', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2169, 'p', '1', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2205, 'p', '1', '/system/getServerInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2204, 'p', '1', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2203, 'p', '1', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2252, 'p', '1', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2246, 'p', '1', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2253, 'p', '1', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (2248, 'p', '1', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (2251, 'p', '1', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2244, 'p', '1', '/user/resetPassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2249, 'p', '1', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2243, 'p', '1', '/user/setSelfSetting', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (2247, 'p', '1', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2245, 'p', '1', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (2250, 'p', '1', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (117, 'p', '2', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (120, 'p', '2', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (122, 'p', '2', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (119, 'p', '2', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (116, 'p', '2', '/api/getApiGroups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (118, 'p', '2', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (121, 'p', '2', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (123, 'p', '2', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (124, 'p', '2', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (125, 'p', '2', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (126, 'p', '2', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (145, 'p', '2', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (144, 'p', '2', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (151, 'p', '2', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (152, 'p', '2', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (149, 'p', '2', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (150, 'p', '2', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (153, 'p', '2', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (141, 'p', '2', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (142, 'p', '2', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (140, 'p', '2', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (143, 'p', '2', '/fileUploadAndDownload/importURL', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (139, 'p', '2', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (146, 'p', '2', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (129, 'p', '2', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (131, 'p', '2', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (133, 'p', '2', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (135, 'p', '2', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (130, 'p', '2', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (127, 'p', '2', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (132, 'p', '2', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (128, 'p', '2', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (134, 'p', '2', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (147, 'p', '2', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (148, 'p', '2', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (115, 'p', '2', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (136, 'p', '2', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (154, 'p', '2', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (137, 'p', '2', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (138, 'p', '2', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1431, 'p', '3', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1438, 'p', '3', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1432, 'p', '3', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1427, 'p', '3', '/paymentUserTool/findCurrentPaymentUserTool', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1430, 'p', '3', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1437, 'p', '3', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1434, 'p', '3', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (1435, 'p', '3', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (1436, 'p', '3', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (1433, 'p', '3', '/user/setUserAuthority', 'POST', '', '', '');

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------
INSERT INTO `exa_customers` VALUES (1, '2024-10-09 16:11:56.670', '2024-10-09 16:11:56.670', NULL, '11', '111', 1, 1);

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `exa_file_id` bigint UNSIGNED NULL DEFAULT NULL,
  `file_chunk_number` bigint NULL DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '2024-10-09 00:20:09.326', '2024-10-09 00:20:09.326', NULL, 'avatar.jpg', 'https://oldwei.oss-cn-hangzhou.aliyuncs.com/pics/avatar.jpg', 'jpg', 'avatar.jpg');

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `chunk_total` bigint NULL DEFAULT NULL,
  `is_finish` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_files_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of exa_files
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (47, '2024-10-30 14:51:47.163', '2024-10-30 14:51:47.163', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZTFhMDYzMTMtNTEyOS00NzkyLTkxYzgtMmRiNzNiYTk2ZWU3IiwiSUQiOjQwLCJVc2VybmFtZSI6IjE4NjI5NTcwNTg2IiwiTmlja05hbWUiOiIxODYyOTU3MDU4NiIsIkF1dGhvcml0eUlkIjozLCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoib2xkd2VpcHJvIiwiYXVkIjpbIkdWQSJdLCJleHAiOjE3MzA4OTk4MzYsIm5iZiI6MTczMDI5NTAzNn0.c72zSPBmMNwGh7VTkTLgFMWG7q5g0GFJtx2JIMu1PX8');

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 214 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` VALUES (3, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/admin_register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (4, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (6, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (7, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` VALUES (8, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (9, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/user/setSelfSetting', '用户界面配置', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (13, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES (20, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/syncApi', '获取待同步API', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (21, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/getApiGroups', '获取路由组', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (22, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/enterSyncApi', '确认同步API', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/api/ignoreApi', '忽略API', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (24, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (26, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (27, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` VALUES (28, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (37, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (38, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (40, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (41, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET');
INSERT INTO `sys_apis` VALUES (42, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (44, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (45, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/upload', '文件上传（建议选择）', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (46, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (47, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (48, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (49, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/fileUploadAndDownload/importURL', '导入URL', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (50, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (51, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (52, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (53, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/customer/customer', '更新客户', '客户', 'PUT');
INSERT INTO `sys_apis` VALUES (54, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/customer/customer', '创建客户', '客户', 'POST');
INSERT INTO `sys_apis` VALUES (55, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/customer/customer', '删除客户', '客户', 'DELETE');
INSERT INTO `sys_apis` VALUES (56, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/customer/customer', '获取单一客户', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (57, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (58, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (59, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (60, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (61, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (62, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (63, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (64, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/pubPlug', '打包插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (65, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/createPackage', '配置模板', '模板配置', 'POST');
INSERT INTO `sys_apis` VALUES (66, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/getTemplates', '获取模板文件', '模板配置', 'GET');
INSERT INTO `sys_apis` VALUES (67, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/getPackage', '获取所有模板', '模板配置', 'POST');
INSERT INTO `sys_apis` VALUES (68, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/delPackage', '删除模板', '模板配置', 'POST');
INSERT INTO `sys_apis` VALUES (69, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (70, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (71, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (72, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (73, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/autoCode/addFunc', '增加模板方法', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (74, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT');
INSERT INTO `sys_apis` VALUES (75, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST');
INSERT INTO `sys_apis` VALUES (76, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE');
INSERT INTO `sys_apis` VALUES (77, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (78, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (79, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST');
INSERT INTO `sys_apis` VALUES (80, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE');
INSERT INTO `sys_apis` VALUES (81, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT');
INSERT INTO `sys_apis` VALUES (82, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典（建议选择）', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (83, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (84, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST');
INSERT INTO `sys_apis` VALUES (85, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (86, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (87, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (88, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (92, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', '2024-11-05 17:13:48.015', '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (93, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', '2024-11-05 17:13:48.015', '/email/sendEmail', '发送邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (94, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (95, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (96, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (97, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/createSysExportTemplate', '新增导出模板', '表格模板', 'POST');
INSERT INTO `sys_apis` VALUES (98, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/deleteSysExportTemplate', '删除导出模板', '表格模板', 'DELETE');
INSERT INTO `sys_apis` VALUES (99, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/deleteSysExportTemplateByIds', '批量删除导出模板', '表格模板', 'DELETE');
INSERT INTO `sys_apis` VALUES (100, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/updateSysExportTemplate', '更新导出模板', '表格模板', 'PUT');
INSERT INTO `sys_apis` VALUES (101, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/findSysExportTemplate', '根据ID获取导出模板', '表格模板', 'GET');
INSERT INTO `sys_apis` VALUES (102, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/getSysExportTemplateList', '获取导出模板列表', '表格模板', 'GET');
INSERT INTO `sys_apis` VALUES (103, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/exportExcel', '导出Excel', '表格模板', 'GET');
INSERT INTO `sys_apis` VALUES (104, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/exportTemplate', '下载模板', '表格模板', 'GET');
INSERT INTO `sys_apis` VALUES (105, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', NULL, '/sysExportTemplate/importExcel', '导入Excel', '表格模板', 'POST');
INSERT INTO `sys_apis` VALUES (106, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', '2024-11-05 17:13:56.563', '/info/createInfo', '新建公告', '公告', 'POST');
INSERT INTO `sys_apis` VALUES (107, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', '2024-11-05 17:13:56.563', '/info/deleteInfo', '删除公告', '公告', 'DELETE');
INSERT INTO `sys_apis` VALUES (108, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', '2024-11-05 17:13:56.563', '/info/deleteInfoByIds', '批量删除公告', '公告', 'DELETE');
INSERT INTO `sys_apis` VALUES (109, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', '2024-11-05 17:13:56.563', '/info/updateInfo', '更新公告', '公告', 'PUT');
INSERT INTO `sys_apis` VALUES (110, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', '2024-11-05 17:13:56.563', '/info/findInfo', '根据ID获取公告', '公告', 'GET');
INSERT INTO `sys_apis` VALUES (111, '2024-10-09 00:20:01.996', '2024-10-09 00:20:01.996', '2024-11-05 17:13:56.563', '/info/getInfoList', '获取公告列表', '公告', 'GET');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `authority_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `uni_sys_authorities_authority_id`(`authority_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2024-10-09 00:20:02.677', '2024-11-05 12:14:21.267', NULL, 1, '管理员', 0, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2024-10-09 00:20:02.677', '2024-10-09 00:20:08.183', NULL, 2, '子角色', 1, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2024-10-09 00:20:02.677', '2024-10-11 05:12:39.657', NULL, 3, '普通用户', 0, 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns`  (
  `authority_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE = InnoDB AUTO_INCREMENT = 56 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES (1, 1);
INSERT INTO `sys_authority_menus` VALUES (1, 2);
INSERT INTO `sys_authority_menus` VALUES (1, 3);
INSERT INTO `sys_authority_menus` VALUES (2, 1);
INSERT INTO `sys_authority_menus` VALUES (2, 2);
INSERT INTO `sys_authority_menus` VALUES (3, 1);
INSERT INTO `sys_authority_menus` VALUES (4, 1);
INSERT INTO `sys_authority_menus` VALUES (4, 2);
INSERT INTO `sys_authority_menus` VALUES (5, 1);
INSERT INTO `sys_authority_menus` VALUES (5, 2);
INSERT INTO `sys_authority_menus` VALUES (6, 1);
INSERT INTO `sys_authority_menus` VALUES (6, 2);
INSERT INTO `sys_authority_menus` VALUES (7, 1);
INSERT INTO `sys_authority_menus` VALUES (8, 1);
INSERT INTO `sys_authority_menus` VALUES (9, 1);
INSERT INTO `sys_authority_menus` VALUES (9, 3);
INSERT INTO `sys_authority_menus` VALUES (10, 1);
INSERT INTO `sys_authority_menus` VALUES (11, 1);
INSERT INTO `sys_authority_menus` VALUES (12, 1);
INSERT INTO `sys_authority_menus` VALUES (13, 1);
INSERT INTO `sys_authority_menus` VALUES (14, 1);
INSERT INTO `sys_authority_menus` VALUES (15, 1);
INSERT INTO `sys_authority_menus` VALUES (16, 1);
INSERT INTO `sys_authority_menus` VALUES (17, 1);
INSERT INTO `sys_authority_menus` VALUES (18, 1);
INSERT INTO `sys_authority_menus` VALUES (19, 1);
INSERT INTO `sys_authority_menus` VALUES (20, 1);
INSERT INTO `sys_authority_menus` VALUES (21, 1);
INSERT INTO `sys_authority_menus` VALUES (27, 1);
INSERT INTO `sys_authority_menus` VALUES (29, 1);

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表名',
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模块名/插件名',
  `request` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '前端传入的结构化信息',
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '结构体名称',
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务库',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'Struct中文名称',
  `templates` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '模板信息',
  `Injections` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '注入路径',
  `flag` bigint NULL DEFAULT NULL COMMENT '[0:创建,1:回滚]',
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api表注册内容',
  `menu_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `export_template_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '导出模板ID',
  `package_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '包ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_histories_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------

-- ----------------------------
-- Table structure for sys_auto_code_packages
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_packages`;
CREATE TABLE `sys_auto_code_packages`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示名',
  `template` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模版',
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '包名',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_packages_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_auto_code_packages
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sys_base_menu_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_btns_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 56 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------
INSERT INTO `sys_base_menu_btns` VALUES (1, '2024-10-09 16:28:22.972', '2024-10-09 16:28:22.972', '2024-11-05 11:55:50.588', 'add', '新增', 30);
INSERT INTO `sys_base_menu_btns` VALUES (2, '2024-10-09 16:28:22.972', '2024-10-09 16:28:22.972', '2024-11-05 11:55:50.588', 'batchDelete', '批量删除', 30);
INSERT INTO `sys_base_menu_btns` VALUES (3, '2024-10-09 16:28:22.972', '2024-10-09 16:28:22.972', '2024-11-05 11:55:50.588', 'delete', '删除', 30);
INSERT INTO `sys_base_menu_btns` VALUES (4, '2024-10-09 16:28:22.972', '2024-10-09 16:28:22.972', '2024-11-05 11:55:50.588', 'edit', '编辑', 30);
INSERT INTO `sys_base_menu_btns` VALUES (5, '2024-10-09 16:28:22.972', '2024-10-09 16:28:22.972', '2024-11-05 11:55:50.588', 'info', '详情', 30);
INSERT INTO `sys_base_menu_btns` VALUES (6, '2024-10-09 16:37:29.341', '2024-10-09 16:37:29.341', '2024-11-05 11:55:53.470', 'add', '新增', 31);
INSERT INTO `sys_base_menu_btns` VALUES (7, '2024-10-09 16:37:29.341', '2024-10-09 16:37:29.341', '2024-11-05 11:55:53.470', 'batchDelete', '批量删除', 31);
INSERT INTO `sys_base_menu_btns` VALUES (8, '2024-10-09 16:37:29.341', '2024-10-09 16:37:29.341', '2024-11-05 11:55:53.470', 'delete', '删除', 31);
INSERT INTO `sys_base_menu_btns` VALUES (9, '2024-10-09 16:37:29.341', '2024-10-09 16:37:29.341', '2024-11-05 11:55:53.470', 'edit', '编辑', 31);
INSERT INTO `sys_base_menu_btns` VALUES (10, '2024-10-09 16:37:29.341', '2024-10-09 16:37:29.341', '2024-11-05 11:55:53.470', 'info', '详情', 31);
INSERT INTO `sys_base_menu_btns` VALUES (11, '2024-10-09 16:42:20.080', '2024-10-09 16:42:20.080', '2024-11-05 11:55:55.779', 'add', '新增', 32);
INSERT INTO `sys_base_menu_btns` VALUES (12, '2024-10-09 16:42:20.080', '2024-10-09 16:42:20.080', '2024-11-05 11:55:55.779', 'batchDelete', '批量删除', 32);
INSERT INTO `sys_base_menu_btns` VALUES (13, '2024-10-09 16:42:20.080', '2024-10-09 16:42:20.080', '2024-11-05 11:55:55.779', 'delete', '删除', 32);
INSERT INTO `sys_base_menu_btns` VALUES (14, '2024-10-09 16:42:20.080', '2024-10-09 16:42:20.080', '2024-11-05 11:55:55.779', 'edit', '编辑', 32);
INSERT INTO `sys_base_menu_btns` VALUES (15, '2024-10-09 16:42:20.080', '2024-10-09 16:42:20.080', '2024-11-05 11:55:55.779', 'info', '详情', 32);
INSERT INTO `sys_base_menu_btns` VALUES (16, '2024-10-09 16:57:42.782', '2024-10-09 16:57:42.782', '2024-11-05 11:56:42.373', 'add', '新增', 36);
INSERT INTO `sys_base_menu_btns` VALUES (17, '2024-10-09 16:57:42.782', '2024-10-09 16:57:42.782', '2024-11-05 11:56:42.373', 'batchDelete', '批量删除', 36);
INSERT INTO `sys_base_menu_btns` VALUES (18, '2024-10-09 16:57:42.782', '2024-10-09 16:57:42.782', '2024-11-05 11:56:42.373', 'delete', '删除', 36);
INSERT INTO `sys_base_menu_btns` VALUES (19, '2024-10-09 16:57:42.782', '2024-10-09 16:57:42.782', '2024-11-05 11:56:42.373', 'edit', '编辑', 36);
INSERT INTO `sys_base_menu_btns` VALUES (20, '2024-10-09 16:57:42.782', '2024-10-09 16:57:42.782', '2024-11-05 11:56:42.373', 'info', '详情', 36);
INSERT INTO `sys_base_menu_btns` VALUES (21, '2024-10-09 17:00:04.760', '2024-10-09 17:00:04.760', '2024-11-05 11:56:38.378', 'add', '新增', 37);
INSERT INTO `sys_base_menu_btns` VALUES (22, '2024-10-09 17:00:04.760', '2024-10-09 17:00:04.760', '2024-11-05 11:56:38.378', 'batchDelete', '批量删除', 37);
INSERT INTO `sys_base_menu_btns` VALUES (23, '2024-10-09 17:00:04.760', '2024-10-09 17:00:04.760', '2024-11-05 11:56:38.378', 'delete', '删除', 37);
INSERT INTO `sys_base_menu_btns` VALUES (24, '2024-10-09 17:00:04.760', '2024-10-09 17:00:04.760', '2024-11-05 11:56:38.378', 'edit', '编辑', 37);
INSERT INTO `sys_base_menu_btns` VALUES (25, '2024-10-09 17:00:04.760', '2024-10-09 17:00:04.760', '2024-11-05 11:56:38.378', 'info', '详情', 37);
INSERT INTO `sys_base_menu_btns` VALUES (26, '2024-10-09 17:00:48.762', '2024-10-09 17:00:48.762', '2024-11-05 11:56:40.575', 'add', '新增', 38);
INSERT INTO `sys_base_menu_btns` VALUES (27, '2024-10-09 17:00:48.762', '2024-10-09 17:00:48.762', '2024-11-05 11:56:40.575', 'batchDelete', '批量删除', 38);
INSERT INTO `sys_base_menu_btns` VALUES (28, '2024-10-09 17:00:48.762', '2024-10-09 17:00:48.762', '2024-11-05 11:56:40.575', 'delete', '删除', 38);
INSERT INTO `sys_base_menu_btns` VALUES (29, '2024-10-09 17:00:48.762', '2024-10-09 17:00:48.762', '2024-11-05 11:56:40.575', 'edit', '编辑', 38);
INSERT INTO `sys_base_menu_btns` VALUES (30, '2024-10-09 17:00:48.762', '2024-10-09 17:00:48.762', '2024-11-05 11:56:40.575', 'info', '详情', 38);
INSERT INTO `sys_base_menu_btns` VALUES (31, '2024-10-09 17:11:09.445', '2024-10-09 17:11:09.445', '2024-11-05 11:56:05.639', 'add', '新增', 39);
INSERT INTO `sys_base_menu_btns` VALUES (32, '2024-10-09 17:11:09.445', '2024-10-09 17:11:09.445', '2024-11-05 11:56:05.639', 'batchDelete', '批量删除', 39);
INSERT INTO `sys_base_menu_btns` VALUES (33, '2024-10-09 17:11:09.445', '2024-10-09 17:11:09.445', '2024-11-05 11:56:05.639', 'delete', '删除', 39);
INSERT INTO `sys_base_menu_btns` VALUES (34, '2024-10-09 17:11:09.445', '2024-10-09 17:11:09.445', '2024-11-05 11:56:05.639', 'edit', '编辑', 39);
INSERT INTO `sys_base_menu_btns` VALUES (35, '2024-10-09 17:11:09.445', '2024-10-09 17:11:09.445', '2024-11-05 11:56:05.639', 'info', '详情', 39);
INSERT INTO `sys_base_menu_btns` VALUES (36, '2024-10-09 17:13:36.152', '2024-10-09 17:13:36.152', '2024-11-05 11:56:27.224', 'add', '新增', 40);
INSERT INTO `sys_base_menu_btns` VALUES (37, '2024-10-09 17:13:36.152', '2024-10-09 17:13:36.152', '2024-11-05 11:56:27.224', 'batchDelete', '批量删除', 40);
INSERT INTO `sys_base_menu_btns` VALUES (38, '2024-10-09 17:13:36.152', '2024-10-09 17:13:36.152', '2024-11-05 11:56:27.224', 'delete', '删除', 40);
INSERT INTO `sys_base_menu_btns` VALUES (39, '2024-10-09 17:13:36.152', '2024-10-09 17:13:36.152', '2024-11-05 11:56:27.224', 'edit', '编辑', 40);
INSERT INTO `sys_base_menu_btns` VALUES (40, '2024-10-09 17:13:36.152', '2024-10-09 17:13:36.152', '2024-11-05 11:56:27.224', 'info', '详情', 40);
INSERT INTO `sys_base_menu_btns` VALUES (41, '2024-10-09 17:14:51.099', '2024-10-09 17:14:51.099', '2024-11-05 11:56:24.074', 'add', '新增', 41);
INSERT INTO `sys_base_menu_btns` VALUES (42, '2024-10-09 17:14:51.099', '2024-10-09 17:14:51.099', '2024-11-05 11:56:24.074', 'batchDelete', '批量删除', 41);
INSERT INTO `sys_base_menu_btns` VALUES (43, '2024-10-09 17:14:51.099', '2024-10-09 17:14:51.099', '2024-11-05 11:56:24.074', 'delete', '删除', 41);
INSERT INTO `sys_base_menu_btns` VALUES (44, '2024-10-09 17:14:51.099', '2024-10-09 17:14:51.099', '2024-11-05 11:56:24.074', 'edit', '编辑', 41);
INSERT INTO `sys_base_menu_btns` VALUES (45, '2024-10-09 17:14:51.099', '2024-10-09 17:14:51.099', '2024-11-05 11:56:24.074', 'info', '详情', 41);
INSERT INTO `sys_base_menu_btns` VALUES (46, '2024-10-09 17:15:13.958', '2024-10-09 17:15:13.958', '2024-11-05 11:56:29.690', 'add', '新增', 42);
INSERT INTO `sys_base_menu_btns` VALUES (47, '2024-10-09 17:15:13.958', '2024-10-09 17:15:13.958', '2024-11-05 11:56:29.690', 'batchDelete', '批量删除', 42);
INSERT INTO `sys_base_menu_btns` VALUES (48, '2024-10-09 17:15:13.958', '2024-10-09 17:15:13.958', '2024-11-05 11:56:29.690', 'delete', '删除', 42);
INSERT INTO `sys_base_menu_btns` VALUES (49, '2024-10-09 17:15:13.958', '2024-10-09 17:15:13.958', '2024-11-05 11:56:29.690', 'edit', '编辑', 42);
INSERT INTO `sys_base_menu_btns` VALUES (50, '2024-10-09 17:15:13.958', '2024-10-09 17:15:13.958', '2024-11-05 11:56:29.690', 'info', '详情', 42);
INSERT INTO `sys_base_menu_btns` VALUES (51, '2024-10-09 17:15:36.699', '2024-10-09 17:15:36.699', '2024-11-05 11:56:32.189', 'add', '新增', 43);
INSERT INTO `sys_base_menu_btns` VALUES (52, '2024-10-09 17:15:36.699', '2024-10-09 17:15:36.699', '2024-11-05 11:56:32.189', 'batchDelete', '批量删除', 43);
INSERT INTO `sys_base_menu_btns` VALUES (53, '2024-10-09 17:15:36.699', '2024-10-09 17:15:36.699', '2024-11-05 11:56:32.189', 'delete', '删除', 43);
INSERT INTO `sys_base_menu_btns` VALUES (54, '2024-10-09 17:15:36.699', '2024-10-09 17:15:36.699', '2024-11-05 11:56:32.189', 'edit', '编辑', 43);
INSERT INTO `sys_base_menu_btns` VALUES (55, '2024-10-09 17:15:36.699', '2024-10-09 17:15:36.699', '2024-11-05 11:56:32.189', 'info', '详情', 43);

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `sys_base_menu_id` bigint UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `menu_level` bigint UNSIGNED NULL DEFAULT NULL,
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint NULL DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '', 0, 0, '仪表板', 'odometer', 0);
INSERT INTO `sys_base_menus` VALUES (2, '2024-10-09 00:20:06.108', '2024-10-12 01:07:20.300', NULL, 0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '', 0, 0, '系统管理', 'setting', 0);
INSERT INTO `sys_base_menus` VALUES (3, '2024-10-09 00:20:06.108', '2024-10-12 01:07:48.363', NULL, 0, 2, 'user', 'user', 0, 'view/superAdmin/user/user.vue', 1, '', 0, 0, '用户管理', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (4, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 2, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 2, '', 0, 0, '角色管理', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (5, '2024-10-09 00:20:06.108', '2024-10-12 01:10:42.855', NULL, 0, 2, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 3, '', 1, 0, '菜单管理', 'grid', 0);
INSERT INTO `sys_base_menus` VALUES (6, '2024-10-09 00:20:06.108', '2024-10-12 01:12:48.886', NULL, 0, 2, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 4, '', 1, 0, '接口管理', 'turn-off', 0);
INSERT INTO `sys_base_menus` VALUES (7, '2024-10-09 00:20:06.108', '2024-10-12 01:13:05.233', NULL, 0, 2, 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, '字典管理', 'key', 0);
INSERT INTO `sys_base_menus` VALUES (8, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 2, 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '', 0, 0, '操作历史', 'pie-chart', 0);
INSERT INTO `sys_base_menus` VALUES (9, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 0, 'person', 'person', 1, 'view/person/person.vue', 4, '', 0, 0, '个人信息', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (10, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 0, 'example', 'example', 0, 'view/example/index.vue', 7, '', 0, 0, '示例文件', 'management', 0);
INSERT INTO `sys_base_menus` VALUES (11, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 10, 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, '', 0, 0, '媒体库（上传下载）', 'upload', 0);
INSERT INTO `sys_base_menus` VALUES (12, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 10, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '', 0, 0, '断点续传', 'upload-filled', 0);
INSERT INTO `sys_base_menus` VALUES (13, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 10, 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, '', 0, 0, '客户列表（资源示例）', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (14, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 0, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '', 0, 0, '系统工具', 'tools', 0);
INSERT INTO `sys_base_menus` VALUES (15, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 14, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, '代码生成器', 'cpu', 0);
INSERT INTO `sys_base_menus` VALUES (16, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 14, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 3, '', 1, 0, '表单生成器', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (17, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 14, 'system', 'system', 0, 'view/systemTools/system/system.vue', 4, '', 0, 0, '系统配置', 'operation', 0);
INSERT INTO `sys_base_menus` VALUES (18, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 14, 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 2, '', 0, 0, '自动化代码管理', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (19, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 14, 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, '自动化代码-${id}', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (20, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 14, 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, '模板配置', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (21, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 0, 'state', 'state', 0, 'view/system/state.vue', 8, '', 0, 0, '服务器状态', 'cloudy', 0);
INSERT INTO `sys_base_menus` VALUES (22, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', '2024-11-05 17:12:56.026', 0, 0, 'plugin', 'plugin', 0, 'view/routerHolder.vue', 6, '', 0, 0, '插件系统', 'cherry', 0);
INSERT INTO `sys_base_menus` VALUES (23, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', '2024-11-05 17:12:52.659', 0, 22, 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', 0, 'https://plugin.gin-vue-admin.com/', 0, '', 0, 0, '插件市场', 'shop', 0);
INSERT INTO `sys_base_menus` VALUES (24, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', '2024-11-05 17:12:48.830', 0, 22, 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '', 0, 0, '插件安装', 'box', 0);
INSERT INTO `sys_base_menus` VALUES (25, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', '2024-11-05 17:12:45.929', 0, 22, 'pubPlug', 'pubPlug', 0, 'view/systemTools/pubPlug/pubPlug.vue', 3, '', 0, 0, '打包插件', 'files', 0);
INSERT INTO `sys_base_menus` VALUES (26, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', '2024-11-05 17:12:42.870', 0, 22, 'plugin-email', 'plugin-email', 0, 'plugin/email/view/index.vue', 4, '', 0, 0, '邮件插件', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (27, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 14, 'exportTemplate', 'exportTemplate', 0, 'view/systemTools/exportTemplate/exportTemplate.vue', 5, '', 0, 0, '表格模板', 'reading', 0);
INSERT INTO `sys_base_menus` VALUES (28, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', '2024-11-05 17:12:39.870', 0, 22, 'anInfo', 'anInfo', 0, 'plugin/announcement/view/info.vue', 5, '', 0, 0, '公告管理[示例]', 'scaleToOriginal', 0);
INSERT INTO `sys_base_menus` VALUES (29, '2024-10-09 00:20:06.108', '2024-10-09 00:20:06.108', NULL, 0, 0, 'about', 'about', 0, 'view/about/index.vue', 9, '', 0, 0, '关于我们', 'info-filled', 0);

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES (1, 1);
INSERT INTO `sys_data_authority_id` VALUES (1, 2);
INSERT INTO `sys_data_authority_id` VALUES (1, 3);
INSERT INTO `sys_data_authority_id` VALUES (2, 2);
INSERT INTO `sys_data_authority_id` VALUES (2, 3);
INSERT INTO `sys_data_authority_id` VALUES (3, 3);

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2024-10-09 00:20:03.685', '2024-10-09 00:20:03.957', NULL, '性别', 'gender', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2024-10-09 00:20:03.685', '2024-10-09 00:20:04.294', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2024-10-09 00:20:03.685', '2024-10-09 00:20:04.630', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2024-10-09 00:20:03.685', '2024-10-09 00:20:04.975', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2024-10-09 00:20:03.685', '2024-10-09 00:20:05.369', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2024-10-09 00:20:03.685', '2024-10-09 00:20:05.705', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '启用状态',
  `sort` bigint NULL DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2024-10-09 00:20:04.024', '2024-10-09 00:20:04.024', NULL, '男', '1', '', 1, 1, 1);
INSERT INTO `sys_dictionary_details` VALUES (2, '2024-10-09 00:20:04.024', '2024-10-09 00:20:04.024', NULL, '女', '2', '', 1, 2, 1);
INSERT INTO `sys_dictionary_details` VALUES (3, '2024-10-09 00:20:04.361', '2024-10-09 00:20:04.361', NULL, 'smallint', '1', 'mysql', 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2024-10-09 00:20:04.361', '2024-10-09 00:20:04.361', NULL, 'mediumint', '2', 'mysql', 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2024-10-09 00:20:04.361', '2024-10-09 00:20:04.361', NULL, 'int', '3', 'mysql', 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (6, '2024-10-09 00:20:04.361', '2024-10-09 00:20:04.361', NULL, 'bigint', '4', 'mysql', 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (7, '2024-10-09 00:20:04.361', '2024-10-09 00:20:04.361', NULL, 'int2', '5', 'pgsql', 1, 5, 2);
INSERT INTO `sys_dictionary_details` VALUES (8, '2024-10-09 00:20:04.361', '2024-10-09 00:20:04.361', NULL, 'int4', '6', 'pgsql', 1, 6, 2);
INSERT INTO `sys_dictionary_details` VALUES (9, '2024-10-09 00:20:04.361', '2024-10-09 00:20:04.361', NULL, 'int6', '7', 'pgsql', 1, 7, 2);
INSERT INTO `sys_dictionary_details` VALUES (10, '2024-10-09 00:20:04.361', '2024-10-09 00:20:04.361', NULL, 'int8', '8', 'pgsql', 1, 8, 2);
INSERT INTO `sys_dictionary_details` VALUES (11, '2024-10-09 00:20:04.697', '2024-10-09 00:20:04.697', NULL, 'date', '', '', 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (12, '2024-10-09 00:20:04.697', '2024-10-09 00:20:04.697', NULL, 'time', '1', 'mysql', 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (13, '2024-10-09 00:20:04.697', '2024-10-09 00:20:04.697', NULL, 'year', '2', 'mysql', 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (14, '2024-10-09 00:20:04.697', '2024-10-09 00:20:04.697', NULL, 'datetime', '3', 'mysql', 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (15, '2024-10-09 00:20:04.697', '2024-10-09 00:20:04.697', NULL, 'timestamp', '5', 'mysql', 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (16, '2024-10-09 00:20:04.697', '2024-10-09 00:20:04.697', NULL, 'timestamptz', '6', 'pgsql', 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (17, '2024-10-09 00:20:05.042', '2024-10-09 00:20:05.042', NULL, 'float', '', '', 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (18, '2024-10-09 00:20:05.042', '2024-10-09 00:20:05.042', NULL, 'double', '1', 'mysql', 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (19, '2024-10-09 00:20:05.042', '2024-10-09 00:20:05.042', NULL, 'decimal', '2', 'mysql', 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (20, '2024-10-09 00:20:05.042', '2024-10-09 00:20:05.042', NULL, 'numeric', '3', 'pgsql', 1, 3, 4);
INSERT INTO `sys_dictionary_details` VALUES (21, '2024-10-09 00:20:05.042', '2024-10-09 00:20:05.042', NULL, 'smallserial', '4', 'pgsql', 1, 4, 4);
INSERT INTO `sys_dictionary_details` VALUES (22, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'char', '', '', 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'varchar', '1', 'mysql', 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (24, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'tinyblob', '2', 'mysql', 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (25, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'tinytext', '3', 'mysql', 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (26, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'text', '4', 'mysql', 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (27, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'blob', '5', 'mysql', 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (28, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'mediumblob', '6', 'mysql', 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (29, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'mediumtext', '7', 'mysql', 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (30, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'longblob', '8', 'mysql', 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (31, '2024-10-09 00:20:05.436', '2024-10-09 00:20:05.436', NULL, 'longtext', '9', 'mysql', 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (32, '2024-10-09 00:20:05.772', '2024-10-09 00:20:05.772', NULL, 'tinyint', '1', 'mysql', 1, 0, 6);
INSERT INTO `sys_dictionary_details` VALUES (33, '2024-10-09 00:20:05.772', '2024-10-09 00:20:05.772', NULL, 'bool', '2', 'pgsql', 1, 0, 6);

-- ----------------------------
-- Table structure for sys_export_template_condition
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_template_condition`;
CREATE TABLE `sys_export_template_condition`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_export_template_condition_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_export_template_condition
-- ----------------------------

-- ----------------------------
-- Table structure for sys_export_template_join
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_template_join`;
CREATE TABLE `sys_export_template_join`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板标识',
  `joins` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联',
  `table` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联表',
  `on` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关联条件',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_export_template_join_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_export_template_join
-- ----------------------------

-- ----------------------------
-- Table structure for sys_export_templates
-- ----------------------------
DROP TABLE IF EXISTS `sys_export_templates`;
CREATE TABLE `sys_export_templates`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `db_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板标识',
  `template_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `limit` bigint NULL DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_export_templates_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_export_templates
-- ----------------------------
INSERT INTO `sys_export_templates` VALUES (1, '2024-10-09 00:20:07.507', '2024-10-09 00:20:07.507', NULL, '', 'api', 'sys_apis', 'api', '{\n\"path\":\"路径\",\n\"method\":\"方法（大写）\",\n\"description\":\"方法介绍\",\n\"api_group\":\"方法分组\"\n}', NULL, '');

-- ----------------------------
-- Table structure for sys_ignore_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_ignore_apis`;
CREATE TABLE `sys_ignore_apis`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_ignore_apis_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 105 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_ignore_apis
-- ----------------------------

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路径',
  `status` bigint NULL DEFAULT NULL COMMENT '请求状态',
  `latency` bigint NULL DEFAULT NULL COMMENT '延迟',
  `agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  `user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 439 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, 1);
INSERT INTO `sys_user_authority` VALUES (2, 3);
INSERT INTO `sys_user_authority` VALUES (5, 3);
INSERT INTO `sys_user_authority` VALUES (31, 3);
INSERT INTO `sys_user_authority` VALUES (32, 3);
INSERT INTO `sys_user_authority` VALUES (33, 3);
INSERT INTO `sys_user_authority` VALUES (34, 3);
INSERT INTO `sys_user_authority` VALUES (35, 3);
INSERT INTO `sys_user_authority` VALUES (36, 3);
INSERT INTO `sys_user_authority` VALUES (37, 3);
INSERT INTO `sys_user_authority` VALUES (39, 3);
INSERT INTO `sys_user_authority` VALUES (40, 3);
INSERT INTO `sys_user_authority` VALUES (41, 3);
INSERT INTO `sys_user_authority` VALUES (42, 3);
INSERT INTO `sys_user_authority` VALUES (43, 3);
INSERT INTO `sys_user_authority` VALUES (44, 3);
INSERT INTO `sys_user_authority` VALUES (45, 3);
INSERT INTO `sys_user_authority` VALUES (46, 3);
INSERT INTO `sys_user_authority` VALUES (47, 3);
INSERT INTO `sys_user_authority` VALUES (48, 3);
INSERT INTO `sys_user_authority` VALUES (49, 3);
INSERT INTO `sys_user_authority` VALUES (50, 3);
INSERT INTO `sys_user_authority` VALUES (51, 1);
INSERT INTO `sys_user_authority` VALUES (52, 3);
INSERT INTO `sys_user_authority` VALUES (53, 3);
INSERT INTO `sys_user_authority` VALUES (54, 3);
INSERT INTO `sys_user_authority` VALUES (55, 3);
INSERT INTO `sys_user_authority` VALUES (56, 3);
INSERT INTO `sys_user_authority` VALUES (57, 3);
INSERT INTO `sys_user_authority` VALUES (58, 3);
INSERT INTO `sys_user_authority` VALUES (59, 3);
INSERT INTO `sys_user_authority` VALUES (60, 3);
INSERT INTO `sys_user_authority` VALUES (61, 3);
INSERT INTO `sys_user_authority` VALUES (62, 3);
INSERT INTO `sys_user_authority` VALUES (63, 3);
INSERT INTO `sys_user_authority` VALUES (64, 3);
INSERT INTO `sys_user_authority` VALUES (65, 3);
INSERT INTO `sys_user_authority` VALUES (66, 3);
INSERT INTO `sys_user_authority` VALUES (67, 3);
INSERT INTO `sys_user_authority` VALUES (68, 3);
INSERT INTO `sys_user_authority` VALUES (69, 3);
INSERT INTO `sys_user_authority` VALUES (70, 3);
INSERT INTO `sys_user_authority` VALUES (71, 3);
INSERT INTO `sys_user_authority` VALUES (72, 3);
INSERT INTO `sys_user_authority` VALUES (73, 3);
INSERT INTO `sys_user_authority` VALUES (74, 3);
INSERT INTO `sys_user_authority` VALUES (75, 3);
INSERT INTO `sys_user_authority` VALUES (76, 3);
INSERT INTO `sys_user_authority` VALUES (77, 3);
INSERT INTO `sys_user_authority` VALUES (78, 3);
INSERT INTO `sys_user_authority` VALUES (79, 3);
INSERT INTO `sys_user_authority` VALUES (80, 3);
INSERT INTO `sys_user_authority` VALUES (81, 3);
INSERT INTO `sys_user_authority` VALUES (82, 3);
INSERT INTO `sys_user_authority` VALUES (83, 3);
INSERT INTO `sys_user_authority` VALUES (84, 3);
INSERT INTO `sys_user_authority` VALUES (85, 3);
INSERT INTO `sys_user_authority` VALUES (86, 3);
INSERT INTO `sys_user_authority` VALUES (87, 3);
INSERT INTO `sys_user_authority` VALUES (88, 3);
INSERT INTO `sys_user_authority` VALUES (89, 3);
INSERT INTO `sys_user_authority` VALUES (90, 3);
INSERT INTO `sys_user_authority` VALUES (91, 3);
INSERT INTO `sys_user_authority` VALUES (92, 3);
INSERT INTO `sys_user_authority` VALUES (93, 3);
INSERT INTO `sys_user_authority` VALUES (94, 3);
INSERT INTO `sys_user_authority` VALUES (95, 3);
INSERT INTO `sys_user_authority` VALUES (96, 3);
INSERT INTO `sys_user_authority` VALUES (97, 3);
INSERT INTO `sys_user_authority` VALUES (98, 3);
INSERT INTO `sys_user_authority` VALUES (99, 3);
INSERT INTO `sys_user_authority` VALUES (100, 3);
INSERT INTO `sys_user_authority` VALUES (101, 3);
INSERT INTO `sys_user_authority` VALUES (102, 3);
INSERT INTO `sys_user_authority` VALUES (103, 3);
INSERT INTO `sys_user_authority` VALUES (104, 3);
INSERT INTO `sys_user_authority` VALUES (105, 3);
INSERT INTO `sys_user_authority` VALUES (106, 3);
INSERT INTO `sys_user_authority` VALUES (107, 3);
INSERT INTO `sys_user_authority` VALUES (108, 3);
INSERT INTO `sys_user_authority` VALUES (109, 3);
INSERT INTO `sys_user_authority` VALUES (110, 3);
INSERT INTO `sys_user_authority` VALUES (111, 3);
INSERT INTO `sys_user_authority` VALUES (112, 3);
INSERT INTO `sys_user_authority` VALUES (113, 3);
INSERT INTO `sys_user_authority` VALUES (114, 3);
INSERT INTO `sys_user_authority` VALUES (115, 3);
INSERT INTO `sys_user_authority` VALUES (116, 3);
INSERT INTO `sys_user_authority` VALUES (117, 3);
INSERT INTO `sys_user_authority` VALUES (118, 3);
INSERT INTO `sys_user_authority` VALUES (119, 3);
INSERT INTO `sys_user_authority` VALUES (120, 3);
INSERT INTO `sys_user_authority` VALUES (121, 3);
INSERT INTO `sys_user_authority` VALUES (122, 3);
INSERT INTO `sys_user_authority` VALUES (123, 3);
INSERT INTO `sys_user_authority` VALUES (124, 3);
INSERT INTO `sys_user_authority` VALUES (125, 3);
INSERT INTO `sys_user_authority` VALUES (126, 3);
INSERT INTO `sys_user_authority` VALUES (127, 3);
INSERT INTO `sys_user_authority` VALUES (128, 3);
INSERT INTO `sys_user_authority` VALUES (129, 3);
INSERT INTO `sys_user_authority` VALUES (130, 3);
INSERT INTO `sys_user_authority` VALUES (131, 3);
INSERT INTO `sys_user_authority` VALUES (132, 3);
INSERT INTO `sys_user_authority` VALUES (133, 3);
INSERT INTO `sys_user_authority` VALUES (134, 3);
INSERT INTO `sys_user_authority` VALUES (135, 3);
INSERT INTO `sys_user_authority` VALUES (136, 3);
INSERT INTO `sys_user_authority` VALUES (137, 3);
INSERT INTO `sys_user_authority` VALUES (138, 3);
INSERT INTO `sys_user_authority` VALUES (139, 3);
INSERT INTO `sys_user_authority` VALUES (140, 3);
INSERT INTO `sys_user_authority` VALUES (141, 3);
INSERT INTO `sys_user_authority` VALUES (142, 3);
INSERT INTO `sys_user_authority` VALUES (143, 3);
INSERT INTO `sys_user_authority` VALUES (144, 3);
INSERT INTO `sys_user_authority` VALUES (145, 3);
INSERT INTO `sys_user_authority` VALUES (146, 3);
INSERT INTO `sys_user_authority` VALUES (147, 3);
INSERT INTO `sys_user_authority` VALUES (148, 3);
INSERT INTO `sys_user_authority` VALUES (149, 3);
INSERT INTO `sys_user_authority` VALUES (150, 3);
INSERT INTO `sys_user_authority` VALUES (151, 3);
INSERT INTO `sys_user_authority` VALUES (152, 3);
INSERT INTO `sys_user_authority` VALUES (153, 3);
INSERT INTO `sys_user_authority` VALUES (154, 3);
INSERT INTO `sys_user_authority` VALUES (155, 3);
INSERT INTO `sys_user_authority` VALUES (156, 3);
INSERT INTO `sys_user_authority` VALUES (157, 3);
INSERT INTO `sys_user_authority` VALUES (158, 3);
INSERT INTO `sys_user_authority` VALUES (159, 3);
INSERT INTO `sys_user_authority` VALUES (160, 3);
INSERT INTO `sys_user_authority` VALUES (161, 3);
INSERT INTO `sys_user_authority` VALUES (162, 3);
INSERT INTO `sys_user_authority` VALUES (163, 3);
INSERT INTO `sys_user_authority` VALUES (164, 3);
INSERT INTO `sys_user_authority` VALUES (165, 3);
INSERT INTO `sys_user_authority` VALUES (166, 3);
INSERT INTO `sys_user_authority` VALUES (167, 3);
INSERT INTO `sys_user_authority` VALUES (168, 3);
INSERT INTO `sys_user_authority` VALUES (169, 3);
INSERT INTO `sys_user_authority` VALUES (170, 3);
INSERT INTO `sys_user_authority` VALUES (171, 3);
INSERT INTO `sys_user_authority` VALUES (172, 3);
INSERT INTO `sys_user_authority` VALUES (173, 3);
INSERT INTO `sys_user_authority` VALUES (174, 3);
INSERT INTO `sys_user_authority` VALUES (175, 3);
INSERT INTO `sys_user_authority` VALUES (176, 3);
INSERT INTO `sys_user_authority` VALUES (177, 3);
INSERT INTO `sys_user_authority` VALUES (178, 3);
INSERT INTO `sys_user_authority` VALUES (179, 3);
INSERT INTO `sys_user_authority` VALUES (180, 3);
INSERT INTO `sys_user_authority` VALUES (181, 3);
INSERT INTO `sys_user_authority` VALUES (182, 3);
INSERT INTO `sys_user_authority` VALUES (183, 3);
INSERT INTO `sys_user_authority` VALUES (184, 3);
INSERT INTO `sys_user_authority` VALUES (185, 3);
INSERT INTO `sys_user_authority` VALUES (186, 3);
INSERT INTO `sys_user_authority` VALUES (187, 3);
INSERT INTO `sys_user_authority` VALUES (188, 3);
INSERT INTO `sys_user_authority` VALUES (189, 3);
INSERT INTO `sys_user_authority` VALUES (190, 3);
INSERT INTO `sys_user_authority` VALUES (191, 3);
INSERT INTO `sys_user_authority` VALUES (192, 3);
INSERT INTO `sys_user_authority` VALUES (193, 3);
INSERT INTO `sys_user_authority` VALUES (194, 3);
INSERT INTO `sys_user_authority` VALUES (195, 3);
INSERT INTO `sys_user_authority` VALUES (196, 3);
INSERT INTO `sys_user_authority` VALUES (197, 3);
INSERT INTO `sys_user_authority` VALUES (198, 3);
INSERT INTO `sys_user_authority` VALUES (199, 3);
INSERT INTO `sys_user_authority` VALUES (200, 3);
INSERT INTO `sys_user_authority` VALUES (201, 3);
INSERT INTO `sys_user_authority` VALUES (202, 3);
INSERT INTO `sys_user_authority` VALUES (203, 3);
INSERT INTO `sys_user_authority` VALUES (204, 3);
INSERT INTO `sys_user_authority` VALUES (205, 3);
INSERT INTO `sys_user_authority` VALUES (206, 3);
INSERT INTO `sys_user_authority` VALUES (207, 3);
INSERT INTO `sys_user_authority` VALUES (208, 3);
INSERT INTO `sys_user_authority` VALUES (209, 3);
INSERT INTO `sys_user_authority` VALUES (210, 3);
INSERT INTO `sys_user_authority` VALUES (211, 3);
INSERT INTO `sys_user_authority` VALUES (212, 3);
INSERT INTO `sys_user_authority` VALUES (213, 3);
INSERT INTO `sys_user_authority` VALUES (215, 3);
INSERT INTO `sys_user_authority` VALUES (217, 3);
INSERT INTO `sys_user_authority` VALUES (218, 3);
INSERT INTO `sys_user_authority` VALUES (219, 3);
INSERT INTO `sys_user_authority` VALUES (220, 3);
INSERT INTO `sys_user_authority` VALUES (221, 3);
INSERT INTO `sys_user_authority` VALUES (222, 3);
INSERT INTO `sys_user_authority` VALUES (223, 3);
INSERT INTO `sys_user_authority` VALUES (224, 3);
INSERT INTO `sys_user_authority` VALUES (225, 3);
INSERT INTO `sys_user_authority` VALUES (226, 3);
INSERT INTO `sys_user_authority` VALUES (227, 3);
INSERT INTO `sys_user_authority` VALUES (228, 3);
INSERT INTO `sys_user_authority` VALUES (229, 3);
INSERT INTO `sys_user_authority` VALUES (230, 3);
INSERT INTO `sys_user_authority` VALUES (231, 3);
INSERT INTO `sys_user_authority` VALUES (232, 3);
INSERT INTO `sys_user_authority` VALUES (233, 3);
INSERT INTO `sys_user_authority` VALUES (234, 3);
INSERT INTO `sys_user_authority` VALUES (235, 3);
INSERT INTO `sys_user_authority` VALUES (236, 3);
INSERT INTO `sys_user_authority` VALUES (237, 3);
INSERT INTO `sys_user_authority` VALUES (238, 3);
INSERT INTO `sys_user_authority` VALUES (239, 3);
INSERT INTO `sys_user_authority` VALUES (240, 3);
INSERT INTO `sys_user_authority` VALUES (241, 3);
INSERT INTO `sys_user_authority` VALUES (242, 3);
INSERT INTO `sys_user_authority` VALUES (243, 3);
INSERT INTO `sys_user_authority` VALUES (244, 3);
INSERT INTO `sys_user_authority` VALUES (245, 3);
INSERT INTO `sys_user_authority` VALUES (246, 3);
INSERT INTO `sys_user_authority` VALUES (247, 3);
INSERT INTO `sys_user_authority` VALUES (248, 3);
INSERT INTO `sys_user_authority` VALUES (249, 3);
INSERT INTO `sys_user_authority` VALUES (250, 3);
INSERT INTO `sys_user_authority` VALUES (251, 3);
INSERT INTO `sys_user_authority` VALUES (252, 3);
INSERT INTO `sys_user_authority` VALUES (253, 3);
INSERT INTO `sys_user_authority` VALUES (254, 3);
INSERT INTO `sys_user_authority` VALUES (255, 3);
INSERT INTO `sys_user_authority` VALUES (256, 3);
INSERT INTO `sys_user_authority` VALUES (257, 3);
INSERT INTO `sys_user_authority` VALUES (258, 3);
INSERT INTO `sys_user_authority` VALUES (259, 3);
INSERT INTO `sys_user_authority` VALUES (260, 3);
INSERT INTO `sys_user_authority` VALUES (261, 3);
INSERT INTO `sys_user_authority` VALUES (262, 3);
INSERT INTO `sys_user_authority` VALUES (263, 3);
INSERT INTO `sys_user_authority` VALUES (264, 3);
INSERT INTO `sys_user_authority` VALUES (265, 3);
INSERT INTO `sys_user_authority` VALUES (266, 3);
INSERT INTO `sys_user_authority` VALUES (267, 3);
INSERT INTO `sys_user_authority` VALUES (268, 3);
INSERT INTO `sys_user_authority` VALUES (269, 3);
INSERT INTO `sys_user_authority` VALUES (270, 3);
INSERT INTO `sys_user_authority` VALUES (271, 3);
INSERT INTO `sys_user_authority` VALUES (272, 3);
INSERT INTO `sys_user_authority` VALUES (273, 3);
INSERT INTO `sys_user_authority` VALUES (274, 3);
INSERT INTO `sys_user_authority` VALUES (275, 3);
INSERT INTO `sys_user_authority` VALUES (276, 3);
INSERT INTO `sys_user_authority` VALUES (277, 3);
INSERT INTO `sys_user_authority` VALUES (278, 3);
INSERT INTO `sys_user_authority` VALUES (279, 3);
INSERT INTO `sys_user_authority` VALUES (280, 3);
INSERT INTO `sys_user_authority` VALUES (281, 3);
INSERT INTO `sys_user_authority` VALUES (282, 3);
INSERT INTO `sys_user_authority` VALUES (283, 3);
INSERT INTO `sys_user_authority` VALUES (284, 3);
INSERT INTO `sys_user_authority` VALUES (285, 3);
INSERT INTO `sys_user_authority` VALUES (286, 3);
INSERT INTO `sys_user_authority` VALUES (287, 3);
INSERT INTO `sys_user_authority` VALUES (288, 3);
INSERT INTO `sys_user_authority` VALUES (289, 3);
INSERT INTO `sys_user_authority` VALUES (290, 3);
INSERT INTO `sys_user_authority` VALUES (291, 3);
INSERT INTO `sys_user_authority` VALUES (292, 3);
INSERT INTO `sys_user_authority` VALUES (293, 3);
INSERT INTO `sys_user_authority` VALUES (294, 3);
INSERT INTO `sys_user_authority` VALUES (295, 3);
INSERT INTO `sys_user_authority` VALUES (296, 3);
INSERT INTO `sys_user_authority` VALUES (297, 3);
INSERT INTO `sys_user_authority` VALUES (298, 3);
INSERT INTO `sys_user_authority` VALUES (299, 3);
INSERT INTO `sys_user_authority` VALUES (300, 3);
INSERT INTO `sys_user_authority` VALUES (301, 3);
INSERT INTO `sys_user_authority` VALUES (302, 3);
INSERT INTO `sys_user_authority` VALUES (303, 3);
INSERT INTO `sys_user_authority` VALUES (304, 3);
INSERT INTO `sys_user_authority` VALUES (305, 3);
INSERT INTO `sys_user_authority` VALUES (306, 3);
INSERT INTO `sys_user_authority` VALUES (307, 3);
INSERT INTO `sys_user_authority` VALUES (308, 3);
INSERT INTO `sys_user_authority` VALUES (309, 3);
INSERT INTO `sys_user_authority` VALUES (310, 3);
INSERT INTO `sys_user_authority` VALUES (311, 3);
INSERT INTO `sys_user_authority` VALUES (312, 3);
INSERT INTO `sys_user_authority` VALUES (313, 3);
INSERT INTO `sys_user_authority` VALUES (314, 3);
INSERT INTO `sys_user_authority` VALUES (315, 3);
INSERT INTO `sys_user_authority` VALUES (316, 3);
INSERT INTO `sys_user_authority` VALUES (317, 3);
INSERT INTO `sys_user_authority` VALUES (318, 3);
INSERT INTO `sys_user_authority` VALUES (319, 3);
INSERT INTO `sys_user_authority` VALUES (320, 3);
INSERT INTO `sys_user_authority` VALUES (321, 3);
INSERT INTO `sys_user_authority` VALUES (322, 3);
INSERT INTO `sys_user_authority` VALUES (323, 3);
INSERT INTO `sys_user_authority` VALUES (324, 3);
INSERT INTO `sys_user_authority` VALUES (325, 3);
INSERT INTO `sys_user_authority` VALUES (326, 3);
INSERT INTO `sys_user_authority` VALUES (327, 3);
INSERT INTO `sys_user_authority` VALUES (328, 3);
INSERT INTO `sys_user_authority` VALUES (329, 3);
INSERT INTO `sys_user_authority` VALUES (330, 3);
INSERT INTO `sys_user_authority` VALUES (331, 3);
INSERT INTO `sys_user_authority` VALUES (332, 3);
INSERT INTO `sys_user_authority` VALUES (333, 3);
INSERT INTO `sys_user_authority` VALUES (334, 3);
INSERT INTO `sys_user_authority` VALUES (335, 3);
INSERT INTO `sys_user_authority` VALUES (336, 3);
INSERT INTO `sys_user_authority` VALUES (337, 3);
INSERT INTO `sys_user_authority` VALUES (338, 3);
INSERT INTO `sys_user_authority` VALUES (339, 3);
INSERT INTO `sys_user_authority` VALUES (340, 3);
INSERT INTO `sys_user_authority` VALUES (341, 3);
INSERT INTO `sys_user_authority` VALUES (342, 3);
INSERT INTO `sys_user_authority` VALUES (343, 3);
INSERT INTO `sys_user_authority` VALUES (344, 3);
INSERT INTO `sys_user_authority` VALUES (345, 3);
INSERT INTO `sys_user_authority` VALUES (346, 3);
INSERT INTO `sys_user_authority` VALUES (347, 3);
INSERT INTO `sys_user_authority` VALUES (348, 3);
INSERT INTO `sys_user_authority` VALUES (349, 3);
INSERT INTO `sys_user_authority` VALUES (350, 3);
INSERT INTO `sys_user_authority` VALUES (351, 3);
INSERT INTO `sys_user_authority` VALUES (352, 3);
INSERT INTO `sys_user_authority` VALUES (353, 3);
INSERT INTO `sys_user_authority` VALUES (354, 3);
INSERT INTO `sys_user_authority` VALUES (355, 3);
INSERT INTO `sys_user_authority` VALUES (356, 3);
INSERT INTO `sys_user_authority` VALUES (357, 3);
INSERT INTO `sys_user_authority` VALUES (358, 3);
INSERT INTO `sys_user_authority` VALUES (359, 3);
INSERT INTO `sys_user_authority` VALUES (360, 3);
INSERT INTO `sys_user_authority` VALUES (361, 3);
INSERT INTO `sys_user_authority` VALUES (362, 3);
INSERT INTO `sys_user_authority` VALUES (363, 3);
INSERT INTO `sys_user_authority` VALUES (364, 3);
INSERT INTO `sys_user_authority` VALUES (365, 3);
INSERT INTO `sys_user_authority` VALUES (366, 3);
INSERT INTO `sys_user_authority` VALUES (367, 3);
INSERT INTO `sys_user_authority` VALUES (368, 3);
INSERT INTO `sys_user_authority` VALUES (369, 3);
INSERT INTO `sys_user_authority` VALUES (370, 3);
INSERT INTO `sys_user_authority` VALUES (371, 3);
INSERT INTO `sys_user_authority` VALUES (372, 3);
INSERT INTO `sys_user_authority` VALUES (373, 3);
INSERT INTO `sys_user_authority` VALUES (374, 3);
INSERT INTO `sys_user_authority` VALUES (375, 3);
INSERT INTO `sys_user_authority` VALUES (376, 3);
INSERT INTO `sys_user_authority` VALUES (377, 3);
INSERT INTO `sys_user_authority` VALUES (378, 3);
INSERT INTO `sys_user_authority` VALUES (379, 3);
INSERT INTO `sys_user_authority` VALUES (380, 3);
INSERT INTO `sys_user_authority` VALUES (381, 3);
INSERT INTO `sys_user_authority` VALUES (382, 3);
INSERT INTO `sys_user_authority` VALUES (383, 3);
INSERT INTO `sys_user_authority` VALUES (384, 3);
INSERT INTO `sys_user_authority` VALUES (385, 3);
INSERT INTO `sys_user_authority` VALUES (386, 3);
INSERT INTO `sys_user_authority` VALUES (387, 3);
INSERT INTO `sys_user_authority` VALUES (388, 3);
INSERT INTO `sys_user_authority` VALUES (389, 3);
INSERT INTO `sys_user_authority` VALUES (390, 3);
INSERT INTO `sys_user_authority` VALUES (391, 3);
INSERT INTO `sys_user_authority` VALUES (392, 3);
INSERT INTO `sys_user_authority` VALUES (393, 3);
INSERT INTO `sys_user_authority` VALUES (394, 3);
INSERT INTO `sys_user_authority` VALUES (395, 3);
INSERT INTO `sys_user_authority` VALUES (396, 3);
INSERT INTO `sys_user_authority` VALUES (397, 3);
INSERT INTO `sys_user_authority` VALUES (398, 3);
INSERT INTO `sys_user_authority` VALUES (399, 3);
INSERT INTO `sys_user_authority` VALUES (400, 3);
INSERT INTO `sys_user_authority` VALUES (401, 3);
INSERT INTO `sys_user_authority` VALUES (402, 3);
INSERT INTO `sys_user_authority` VALUES (403, 3);
INSERT INTO `sys_user_authority` VALUES (404, 3);
INSERT INTO `sys_user_authority` VALUES (405, 3);
INSERT INTO `sys_user_authority` VALUES (406, 3);
INSERT INTO `sys_user_authority` VALUES (407, 3);
INSERT INTO `sys_user_authority` VALUES (408, 3);
INSERT INTO `sys_user_authority` VALUES (409, 3);
INSERT INTO `sys_user_authority` VALUES (410, 3);
INSERT INTO `sys_user_authority` VALUES (411, 3);
INSERT INTO `sys_user_authority` VALUES (412, 3);
INSERT INTO `sys_user_authority` VALUES (413, 3);
INSERT INTO `sys_user_authority` VALUES (414, 3);
INSERT INTO `sys_user_authority` VALUES (415, 3);
INSERT INTO `sys_user_authority` VALUES (416, 3);
INSERT INTO `sys_user_authority` VALUES (417, 3);
INSERT INTO `sys_user_authority` VALUES (418, 3);
INSERT INTO `sys_user_authority` VALUES (419, 3);
INSERT INTO `sys_user_authority` VALUES (420, 3);
INSERT INTO `sys_user_authority` VALUES (421, 3);
INSERT INTO `sys_user_authority` VALUES (422, 3);
INSERT INTO `sys_user_authority` VALUES (423, 3);
INSERT INTO `sys_user_authority` VALUES (424, 3);
INSERT INTO `sys_user_authority` VALUES (425, 3);
INSERT INTO `sys_user_authority` VALUES (426, 3);
INSERT INTO `sys_user_authority` VALUES (427, 3);
INSERT INTO `sys_user_authority` VALUES (428, 3);
INSERT INTO `sys_user_authority` VALUES (429, 3);
INSERT INTO `sys_user_authority` VALUES (430, 3);
INSERT INTO `sys_user_authority` VALUES (431, 3);
INSERT INTO `sys_user_authority` VALUES (432, 3);
INSERT INTO `sys_user_authority` VALUES (433, 3);
INSERT INTO `sys_user_authority` VALUES (434, 3);
INSERT INTO `sys_user_authority` VALUES (435, 3);
INSERT INTO `sys_user_authority` VALUES (436, 3);
INSERT INTO `sys_user_authority` VALUES (437, 3);
INSERT INTO `sys_user_authority` VALUES (438, 3);
INSERT INTO `sys_user_authority` VALUES (439, 3);
INSERT INTO `sys_user_authority` VALUES (440, 3);
INSERT INTO `sys_user_authority` VALUES (441, 3);
INSERT INTO `sys_user_authority` VALUES (442, 3);
INSERT INTO `sys_user_authority` VALUES (443, 3);
INSERT INTO `sys_user_authority` VALUES (444, 3);
INSERT INTO `sys_user_authority` VALUES (445, 3);
INSERT INTO `sys_user_authority` VALUES (446, 3);
INSERT INTO `sys_user_authority` VALUES (447, 3);
INSERT INTO `sys_user_authority` VALUES (448, 3);
INSERT INTO `sys_user_authority` VALUES (449, 3);
INSERT INTO `sys_user_authority` VALUES (450, 3);
INSERT INTO `sys_user_authority` VALUES (451, 3);
INSERT INTO `sys_user_authority` VALUES (452, 3);
INSERT INTO `sys_user_authority` VALUES (453, 3);
INSERT INTO `sys_user_authority` VALUES (454, 3);
INSERT INTO `sys_user_authority` VALUES (455, 3);
INSERT INTO `sys_user_authority` VALUES (456, 3);
INSERT INTO `sys_user_authority` VALUES (457, 3);
INSERT INTO `sys_user_authority` VALUES (458, 3);
INSERT INTO `sys_user_authority` VALUES (459, 3);
INSERT INTO `sys_user_authority` VALUES (460, 3);
INSERT INTO `sys_user_authority` VALUES (461, 3);
INSERT INTO `sys_user_authority` VALUES (462, 3);
INSERT INTO `sys_user_authority` VALUES (463, 3);
INSERT INTO `sys_user_authority` VALUES (464, 3);
INSERT INTO `sys_user_authority` VALUES (465, 3);
INSERT INTO `sys_user_authority` VALUES (466, 3);
INSERT INTO `sys_user_authority` VALUES (467, 3);
INSERT INTO `sys_user_authority` VALUES (468, 3);
INSERT INTO `sys_user_authority` VALUES (469, 3);
INSERT INTO `sys_user_authority` VALUES (470, 3);
INSERT INTO `sys_user_authority` VALUES (471, 3);
INSERT INTO `sys_user_authority` VALUES (472, 3);
INSERT INTO `sys_user_authority` VALUES (473, 3);
INSERT INTO `sys_user_authority` VALUES (474, 3);
INSERT INTO `sys_user_authority` VALUES (475, 3);
INSERT INTO `sys_user_authority` VALUES (476, 3);
INSERT INTO `sys_user_authority` VALUES (477, 3);
INSERT INTO `sys_user_authority` VALUES (478, 3);
INSERT INTO `sys_user_authority` VALUES (479, 3);
INSERT INTO `sys_user_authority` VALUES (480, 3);
INSERT INTO `sys_user_authority` VALUES (481, 3);
INSERT INTO `sys_user_authority` VALUES (482, 3);
INSERT INTO `sys_user_authority` VALUES (483, 3);
INSERT INTO `sys_user_authority` VALUES (484, 3);
INSERT INTO `sys_user_authority` VALUES (485, 3);
INSERT INTO `sys_user_authority` VALUES (486, 3);
INSERT INTO `sys_user_authority` VALUES (487, 3);
INSERT INTO `sys_user_authority` VALUES (488, 3);
INSERT INTO `sys_user_authority` VALUES (489, 3);
INSERT INTO `sys_user_authority` VALUES (490, 3);
INSERT INTO `sys_user_authority` VALUES (491, 3);
INSERT INTO `sys_user_authority` VALUES (492, 3);
INSERT INTO `sys_user_authority` VALUES (493, 3);
INSERT INTO `sys_user_authority` VALUES (494, 3);
INSERT INTO `sys_user_authority` VALUES (495, 3);
INSERT INTO `sys_user_authority` VALUES (496, 3);
INSERT INTO `sys_user_authority` VALUES (497, 3);
INSERT INTO `sys_user_authority` VALUES (498, 3);
INSERT INTO `sys_user_authority` VALUES (499, 3);
INSERT INTO `sys_user_authority` VALUES (500, 3);
INSERT INTO `sys_user_authority` VALUES (501, 3);
INSERT INTO `sys_user_authority` VALUES (502, 3);
INSERT INTO `sys_user_authority` VALUES (503, 3);
INSERT INTO `sys_user_authority` VALUES (504, 3);
INSERT INTO `sys_user_authority` VALUES (505, 3);
INSERT INTO `sys_user_authority` VALUES (506, 3);
INSERT INTO `sys_user_authority` VALUES (507, 3);
INSERT INTO `sys_user_authority` VALUES (508, 3);
INSERT INTO `sys_user_authority` VALUES (509, 3);
INSERT INTO `sys_user_authority` VALUES (510, 3);
INSERT INTO `sys_user_authority` VALUES (511, 3);
INSERT INTO `sys_user_authority` VALUES (512, 3);
INSERT INTO `sys_user_authority` VALUES (513, 3);
INSERT INTO `sys_user_authority` VALUES (514, 3);
INSERT INTO `sys_user_authority` VALUES (515, 3);
INSERT INTO `sys_user_authority` VALUES (516, 3);
INSERT INTO `sys_user_authority` VALUES (517, 3);
INSERT INTO `sys_user_authority` VALUES (518, 3);
INSERT INTO `sys_user_authority` VALUES (519, 3);
INSERT INTO `sys_user_authority` VALUES (520, 3);
INSERT INTO `sys_user_authority` VALUES (521, 3);
INSERT INTO `sys_user_authority` VALUES (522, 3);
INSERT INTO `sys_user_authority` VALUES (523, 3);
INSERT INTO `sys_user_authority` VALUES (524, 3);
INSERT INTO `sys_user_authority` VALUES (525, 3);
INSERT INTO `sys_user_authority` VALUES (526, 3);
INSERT INTO `sys_user_authority` VALUES (527, 3);
INSERT INTO `sys_user_authority` VALUES (528, 3);
INSERT INTO `sys_user_authority` VALUES (529, 3);
INSERT INTO `sys_user_authority` VALUES (530, 3);
INSERT INTO `sys_user_authority` VALUES (531, 3);
INSERT INTO `sys_user_authority` VALUES (532, 3);
INSERT INTO `sys_user_authority` VALUES (533, 3);
INSERT INTO `sys_user_authority` VALUES (534, 3);
INSERT INTO `sys_user_authority` VALUES (535, 3);
INSERT INTO `sys_user_authority` VALUES (536, 3);
INSERT INTO `sys_user_authority` VALUES (537, 3);
INSERT INTO `sys_user_authority` VALUES (538, 3);
INSERT INTO `sys_user_authority` VALUES (539, 3);
INSERT INTO `sys_user_authority` VALUES (540, 3);
INSERT INTO `sys_user_authority` VALUES (541, 3);
INSERT INTO `sys_user_authority` VALUES (542, 3);
INSERT INTO `sys_user_authority` VALUES (543, 3);
INSERT INTO `sys_user_authority` VALUES (544, 3);
INSERT INTO `sys_user_authority` VALUES (545, 3);
INSERT INTO `sys_user_authority` VALUES (546, 3);
INSERT INTO `sys_user_authority` VALUES (547, 3);
INSERT INTO `sys_user_authority` VALUES (548, 3);
INSERT INTO `sys_user_authority` VALUES (549, 3);
INSERT INTO `sys_user_authority` VALUES (550, 3);
INSERT INTO `sys_user_authority` VALUES (551, 3);
INSERT INTO `sys_user_authority` VALUES (552, 3);
INSERT INTO `sys_user_authority` VALUES (553, 3);
INSERT INTO `sys_user_authority` VALUES (554, 3);
INSERT INTO `sys_user_authority` VALUES (555, 3);
INSERT INTO `sys_user_authority` VALUES (556, 3);
INSERT INTO `sys_user_authority` VALUES (557, 3);
INSERT INTO `sys_user_authority` VALUES (558, 3);
INSERT INTO `sys_user_authority` VALUES (559, 3);
INSERT INTO `sys_user_authority` VALUES (560, 3);
INSERT INTO `sys_user_authority` VALUES (561, 3);
INSERT INTO `sys_user_authority` VALUES (562, 3);
INSERT INTO `sys_user_authority` VALUES (563, 3);
INSERT INTO `sys_user_authority` VALUES (564, 3);
INSERT INTO `sys_user_authority` VALUES (565, 3);
INSERT INTO `sys_user_authority` VALUES (566, 3);
INSERT INTO `sys_user_authority` VALUES (567, 3);
INSERT INTO `sys_user_authority` VALUES (568, 3);
INSERT INTO `sys_user_authority` VALUES (569, 3);
INSERT INTO `sys_user_authority` VALUES (570, 3);
INSERT INTO `sys_user_authority` VALUES (571, 3);
INSERT INTO `sys_user_authority` VALUES (572, 3);
INSERT INTO `sys_user_authority` VALUES (573, 3);
INSERT INTO `sys_user_authority` VALUES (574, 3);
INSERT INTO `sys_user_authority` VALUES (575, 3);
INSERT INTO `sys_user_authority` VALUES (576, 3);
INSERT INTO `sys_user_authority` VALUES (577, 3);
INSERT INTO `sys_user_authority` VALUES (578, 3);
INSERT INTO `sys_user_authority` VALUES (579, 3);
INSERT INTO `sys_user_authority` VALUES (580, 3);
INSERT INTO `sys_user_authority` VALUES (581, 3);
INSERT INTO `sys_user_authority` VALUES (582, 3);
INSERT INTO `sys_user_authority` VALUES (583, 3);
INSERT INTO `sys_user_authority` VALUES (584, 3);
INSERT INTO `sys_user_authority` VALUES (585, 3);
INSERT INTO `sys_user_authority` VALUES (586, 3);
INSERT INTO `sys_user_authority` VALUES (587, 3);
INSERT INTO `sys_user_authority` VALUES (588, 3);
INSERT INTO `sys_user_authority` VALUES (589, 3);
INSERT INTO `sys_user_authority` VALUES (590, 3);
INSERT INTO `sys_user_authority` VALUES (591, 3);
INSERT INTO `sys_user_authority` VALUES (592, 3);
INSERT INTO `sys_user_authority` VALUES (593, 3);
INSERT INTO `sys_user_authority` VALUES (594, 3);
INSERT INTO `sys_user_authority` VALUES (595, 3);
INSERT INTO `sys_user_authority` VALUES (596, 3);
INSERT INTO `sys_user_authority` VALUES (597, 3);
INSERT INTO `sys_user_authority` VALUES (598, 3);
INSERT INTO `sys_user_authority` VALUES (599, 3);
INSERT INTO `sys_user_authority` VALUES (600, 3);
INSERT INTO `sys_user_authority` VALUES (601, 3);
INSERT INTO `sys_user_authority` VALUES (602, 3);
INSERT INTO `sys_user_authority` VALUES (603, 3);
INSERT INTO `sys_user_authority` VALUES (604, 3);
INSERT INTO `sys_user_authority` VALUES (605, 3);
INSERT INTO `sys_user_authority` VALUES (606, 3);
INSERT INTO `sys_user_authority` VALUES (607, 3);
INSERT INTO `sys_user_authority` VALUES (608, 3);
INSERT INTO `sys_user_authority` VALUES (609, 3);
INSERT INTO `sys_user_authority` VALUES (610, 3);
INSERT INTO `sys_user_authority` VALUES (611, 3);
INSERT INTO `sys_user_authority` VALUES (612, 3);
INSERT INTO `sys_user_authority` VALUES (613, 3);
INSERT INTO `sys_user_authority` VALUES (614, 3);
INSERT INTO `sys_user_authority` VALUES (615, 3);
INSERT INTO `sys_user_authority` VALUES (616, 3);
INSERT INTO `sys_user_authority` VALUES (617, 3);
INSERT INTO `sys_user_authority` VALUES (618, 3);
INSERT INTO `sys_user_authority` VALUES (619, 3);
INSERT INTO `sys_user_authority` VALUES (620, 3);
INSERT INTO `sys_user_authority` VALUES (621, 3);
INSERT INTO `sys_user_authority` VALUES (622, 3);
INSERT INTO `sys_user_authority` VALUES (623, 3);
INSERT INTO `sys_user_authority` VALUES (624, 3);
INSERT INTO `sys_user_authority` VALUES (625, 3);
INSERT INTO `sys_user_authority` VALUES (626, 3);
INSERT INTO `sys_user_authority` VALUES (627, 3);
INSERT INTO `sys_user_authority` VALUES (628, 3);
INSERT INTO `sys_user_authority` VALUES (629, 3);
INSERT INTO `sys_user_authority` VALUES (630, 3);
INSERT INTO `sys_user_authority` VALUES (631, 3);
INSERT INTO `sys_user_authority` VALUES (632, 3);
INSERT INTO `sys_user_authority` VALUES (633, 3);
INSERT INTO `sys_user_authority` VALUES (634, 3);
INSERT INTO `sys_user_authority` VALUES (635, 3);
INSERT INTO `sys_user_authority` VALUES (636, 3);
INSERT INTO `sys_user_authority` VALUES (637, 3);
INSERT INTO `sys_user_authority` VALUES (638, 3);
INSERT INTO `sys_user_authority` VALUES (639, 3);
INSERT INTO `sys_user_authority` VALUES (640, 3);
INSERT INTO `sys_user_authority` VALUES (641, 3);
INSERT INTO `sys_user_authority` VALUES (642, 3);
INSERT INTO `sys_user_authority` VALUES (643, 3);
INSERT INTO `sys_user_authority` VALUES (644, 3);
INSERT INTO `sys_user_authority` VALUES (645, 3);
INSERT INTO `sys_user_authority` VALUES (646, 3);
INSERT INTO `sys_user_authority` VALUES (647, 3);
INSERT INTO `sys_user_authority` VALUES (648, 3);
INSERT INTO `sys_user_authority` VALUES (649, 3);
INSERT INTO `sys_user_authority` VALUES (650, 3);
INSERT INTO `sys_user_authority` VALUES (651, 3);
INSERT INTO `sys_user_authority` VALUES (652, 3);
INSERT INTO `sys_user_authority` VALUES (653, 3);
INSERT INTO `sys_user_authority` VALUES (654, 3);
INSERT INTO `sys_user_authority` VALUES (655, 3);
INSERT INTO `sys_user_authority` VALUES (656, 3);
INSERT INTO `sys_user_authority` VALUES (657, 3);
INSERT INTO `sys_user_authority` VALUES (658, 3);
INSERT INTO `sys_user_authority` VALUES (659, 3);
INSERT INTO `sys_user_authority` VALUES (660, 3);
INSERT INTO `sys_user_authority` VALUES (661, 3);
INSERT INTO `sys_user_authority` VALUES (662, 3);
INSERT INTO `sys_user_authority` VALUES (663, 3);
INSERT INTO `sys_user_authority` VALUES (664, 3);
INSERT INTO `sys_user_authority` VALUES (665, 3);
INSERT INTO `sys_user_authority` VALUES (666, 3);
INSERT INTO `sys_user_authority` VALUES (667, 3);
INSERT INTO `sys_user_authority` VALUES (668, 3);
INSERT INTO `sys_user_authority` VALUES (669, 3);
INSERT INTO `sys_user_authority` VALUES (670, 3);
INSERT INTO `sys_user_authority` VALUES (671, 3);
INSERT INTO `sys_user_authority` VALUES (672, 3);
INSERT INTO `sys_user_authority` VALUES (673, 3);
INSERT INTO `sys_user_authority` VALUES (674, 3);
INSERT INTO `sys_user_authority` VALUES (675, 3);
INSERT INTO `sys_user_authority` VALUES (676, 3);
INSERT INTO `sys_user_authority` VALUES (677, 3);
INSERT INTO `sys_user_authority` VALUES (678, 3);
INSERT INTO `sys_user_authority` VALUES (679, 3);
INSERT INTO `sys_user_authority` VALUES (680, 3);
INSERT INTO `sys_user_authority` VALUES (681, 3);
INSERT INTO `sys_user_authority` VALUES (682, 3);
INSERT INTO `sys_user_authority` VALUES (683, 3);
INSERT INTO `sys_user_authority` VALUES (684, 3);
INSERT INTO `sys_user_authority` VALUES (685, 3);
INSERT INTO `sys_user_authority` VALUES (686, 3);
INSERT INTO `sys_user_authority` VALUES (687, 3);
INSERT INTO `sys_user_authority` VALUES (688, 3);
INSERT INTO `sys_user_authority` VALUES (689, 3);
INSERT INTO `sys_user_authority` VALUES (690, 3);
INSERT INTO `sys_user_authority` VALUES (691, 3);
INSERT INTO `sys_user_authority` VALUES (692, 3);
INSERT INTO `sys_user_authority` VALUES (693, 3);
INSERT INTO `sys_user_authority` VALUES (694, 3);
INSERT INTO `sys_user_authority` VALUES (695, 3);
INSERT INTO `sys_user_authority` VALUES (696, 3);
INSERT INTO `sys_user_authority` VALUES (697, 3);
INSERT INTO `sys_user_authority` VALUES (698, 3);
INSERT INTO `sys_user_authority` VALUES (699, 3);
INSERT INTO `sys_user_authority` VALUES (700, 3);
INSERT INTO `sys_user_authority` VALUES (701, 3);
INSERT INTO `sys_user_authority` VALUES (702, 3);
INSERT INTO `sys_user_authority` VALUES (703, 3);
INSERT INTO `sys_user_authority` VALUES (704, 3);
INSERT INTO `sys_user_authority` VALUES (705, 3);
INSERT INTO `sys_user_authority` VALUES (706, 3);
INSERT INTO `sys_user_authority` VALUES (707, 3);
INSERT INTO `sys_user_authority` VALUES (708, 3);
INSERT INTO `sys_user_authority` VALUES (709, 3);
INSERT INTO `sys_user_authority` VALUES (710, 3);
INSERT INTO `sys_user_authority` VALUES (711, 3);
INSERT INTO `sys_user_authority` VALUES (712, 3);
INSERT INTO `sys_user_authority` VALUES (713, 3);
INSERT INTO `sys_user_authority` VALUES (714, 3);
INSERT INTO `sys_user_authority` VALUES (715, 3);
INSERT INTO `sys_user_authority` VALUES (716, 3);
INSERT INTO `sys_user_authority` VALUES (717, 3);
INSERT INTO `sys_user_authority` VALUES (718, 3);
INSERT INTO `sys_user_authority` VALUES (719, 3);
INSERT INTO `sys_user_authority` VALUES (720, 3);
INSERT INTO `sys_user_authority` VALUES (721, 3);
INSERT INTO `sys_user_authority` VALUES (722, 3);
INSERT INTO `sys_user_authority` VALUES (723, 3);
INSERT INTO `sys_user_authority` VALUES (724, 3);
INSERT INTO `sys_user_authority` VALUES (725, 3);
INSERT INTO `sys_user_authority` VALUES (726, 3);
INSERT INTO `sys_user_authority` VALUES (727, 3);
INSERT INTO `sys_user_authority` VALUES (728, 3);
INSERT INTO `sys_user_authority` VALUES (729, 3);
INSERT INTO `sys_user_authority` VALUES (730, 3);
INSERT INTO `sys_user_authority` VALUES (731, 3);
INSERT INTO `sys_user_authority` VALUES (732, 3);
INSERT INTO `sys_user_authority` VALUES (733, 3);
INSERT INTO `sys_user_authority` VALUES (734, 3);
INSERT INTO `sys_user_authority` VALUES (735, 3);
INSERT INTO `sys_user_authority` VALUES (736, 3);
INSERT INTO `sys_user_authority` VALUES (737, 3);
INSERT INTO `sys_user_authority` VALUES (738, 3);
INSERT INTO `sys_user_authority` VALUES (739, 3);
INSERT INTO `sys_user_authority` VALUES (740, 3);
INSERT INTO `sys_user_authority` VALUES (741, 3);
INSERT INTO `sys_user_authority` VALUES (742, 3);
INSERT INTO `sys_user_authority` VALUES (743, 3);
INSERT INTO `sys_user_authority` VALUES (744, 3);
INSERT INTO `sys_user_authority` VALUES (745, 3);
INSERT INTO `sys_user_authority` VALUES (746, 3);
INSERT INTO `sys_user_authority` VALUES (747, 3);
INSERT INTO `sys_user_authority` VALUES (748, 3);
INSERT INTO `sys_user_authority` VALUES (749, 3);
INSERT INTO `sys_user_authority` VALUES (750, 3);
INSERT INTO `sys_user_authority` VALUES (751, 3);
INSERT INTO `sys_user_authority` VALUES (752, 3);
INSERT INTO `sys_user_authority` VALUES (753, 3);
INSERT INTO `sys_user_authority` VALUES (754, 3);
INSERT INTO `sys_user_authority` VALUES (755, 3);
INSERT INTO `sys_user_authority` VALUES (756, 3);
INSERT INTO `sys_user_authority` VALUES (757, 3);
INSERT INTO `sys_user_authority` VALUES (758, 3);
INSERT INTO `sys_user_authority` VALUES (759, 3);
INSERT INTO `sys_user_authority` VALUES (760, 3);
INSERT INTO `sys_user_authority` VALUES (761, 3);
INSERT INTO `sys_user_authority` VALUES (762, 3);
INSERT INTO `sys_user_authority` VALUES (763, 3);
INSERT INTO `sys_user_authority` VALUES (764, 3);
INSERT INTO `sys_user_authority` VALUES (765, 3);
INSERT INTO `sys_user_authority` VALUES (766, 3);
INSERT INTO `sys_user_authority` VALUES (767, 3);
INSERT INTO `sys_user_authority` VALUES (768, 3);
INSERT INTO `sys_user_authority` VALUES (769, 3);
INSERT INTO `sys_user_authority` VALUES (770, 3);
INSERT INTO `sys_user_authority` VALUES (771, 3);
INSERT INTO `sys_user_authority` VALUES (772, 3);
INSERT INTO `sys_user_authority` VALUES (773, 3);
INSERT INTO `sys_user_authority` VALUES (774, 3);
INSERT INTO `sys_user_authority` VALUES (775, 3);
INSERT INTO `sys_user_authority` VALUES (776, 3);
INSERT INTO `sys_user_authority` VALUES (777, 3);
INSERT INTO `sys_user_authority` VALUES (778, 3);
INSERT INTO `sys_user_authority` VALUES (779, 3);
INSERT INTO `sys_user_authority` VALUES (780, 3);
INSERT INTO `sys_user_authority` VALUES (781, 3);
INSERT INTO `sys_user_authority` VALUES (782, 3);
INSERT INTO `sys_user_authority` VALUES (783, 3);
INSERT INTO `sys_user_authority` VALUES (784, 3);
INSERT INTO `sys_user_authority` VALUES (785, 3);
INSERT INTO `sys_user_authority` VALUES (786, 3);
INSERT INTO `sys_user_authority` VALUES (787, 3);
INSERT INTO `sys_user_authority` VALUES (788, 3);
INSERT INTO `sys_user_authority` VALUES (789, 3);
INSERT INTO `sys_user_authority` VALUES (790, 3);
INSERT INTO `sys_user_authority` VALUES (791, 3);
INSERT INTO `sys_user_authority` VALUES (792, 3);
INSERT INTO `sys_user_authority` VALUES (793, 3);
INSERT INTO `sys_user_authority` VALUES (794, 3);
INSERT INTO `sys_user_authority` VALUES (795, 3);
INSERT INTO `sys_user_authority` VALUES (796, 3);
INSERT INTO `sys_user_authority` VALUES (797, 3);
INSERT INTO `sys_user_authority` VALUES (798, 3);
INSERT INTO `sys_user_authority` VALUES (799, 3);
INSERT INTO `sys_user_authority` VALUES (800, 3);
INSERT INTO `sys_user_authority` VALUES (801, 3);
INSERT INTO `sys_user_authority` VALUES (802, 3);
INSERT INTO `sys_user_authority` VALUES (803, 3);
INSERT INTO `sys_user_authority` VALUES (804, 3);
INSERT INTO `sys_user_authority` VALUES (805, 3);
INSERT INTO `sys_user_authority` VALUES (806, 3);
INSERT INTO `sys_user_authority` VALUES (807, 3);
INSERT INTO `sys_user_authority` VALUES (808, 3);
INSERT INTO `sys_user_authority` VALUES (809, 3);
INSERT INTO `sys_user_authority` VALUES (810, 3);
INSERT INTO `sys_user_authority` VALUES (811, 3);
INSERT INTO `sys_user_authority` VALUES (812, 3);
INSERT INTO `sys_user_authority` VALUES (813, 3);
INSERT INTO `sys_user_authority` VALUES (814, 3);
INSERT INTO `sys_user_authority` VALUES (815, 3);
INSERT INTO `sys_user_authority` VALUES (816, 3);
INSERT INTO `sys_user_authority` VALUES (817, 3);
INSERT INTO `sys_user_authority` VALUES (818, 3);
INSERT INTO `sys_user_authority` VALUES (819, 3);
INSERT INTO `sys_user_authority` VALUES (820, 3);
INSERT INTO `sys_user_authority` VALUES (821, 3);
INSERT INTO `sys_user_authority` VALUES (822, 3);
INSERT INTO `sys_user_authority` VALUES (823, 3);
INSERT INTO `sys_user_authority` VALUES (824, 3);
INSERT INTO `sys_user_authority` VALUES (825, 3);
INSERT INTO `sys_user_authority` VALUES (826, 3);
INSERT INTO `sys_user_authority` VALUES (827, 3);
INSERT INTO `sys_user_authority` VALUES (828, 3);
INSERT INTO `sys_user_authority` VALUES (829, 3);
INSERT INTO `sys_user_authority` VALUES (830, 3);
INSERT INTO `sys_user_authority` VALUES (831, 3);
INSERT INTO `sys_user_authority` VALUES (832, 3);
INSERT INTO `sys_user_authority` VALUES (833, 3);
INSERT INTO `sys_user_authority` VALUES (834, 3);
INSERT INTO `sys_user_authority` VALUES (835, 3);
INSERT INTO `sys_user_authority` VALUES (836, 3);
INSERT INTO `sys_user_authority` VALUES (837, 3);
INSERT INTO `sys_user_authority` VALUES (838, 3);
INSERT INTO `sys_user_authority` VALUES (839, 3);
INSERT INTO `sys_user_authority` VALUES (840, 3);
INSERT INTO `sys_user_authority` VALUES (841, 3);
INSERT INTO `sys_user_authority` VALUES (842, 3);
INSERT INTO `sys_user_authority` VALUES (843, 3);
INSERT INTO `sys_user_authority` VALUES (844, 3);
INSERT INTO `sys_user_authority` VALUES (845, 3);
INSERT INTO `sys_user_authority` VALUES (846, 3);
INSERT INTO `sys_user_authority` VALUES (847, 3);
INSERT INTO `sys_user_authority` VALUES (848, 3);
INSERT INTO `sys_user_authority` VALUES (849, 3);
INSERT INTO `sys_user_authority` VALUES (850, 3);
INSERT INTO `sys_user_authority` VALUES (851, 3);
INSERT INTO `sys_user_authority` VALUES (852, 3);
INSERT INTO `sys_user_authority` VALUES (853, 3);
INSERT INTO `sys_user_authority` VALUES (854, 3);
INSERT INTO `sys_user_authority` VALUES (855, 3);
INSERT INTO `sys_user_authority` VALUES (856, 3);
INSERT INTO `sys_user_authority` VALUES (857, 3);
INSERT INTO `sys_user_authority` VALUES (858, 3);
INSERT INTO `sys_user_authority` VALUES (859, 3);
INSERT INTO `sys_user_authority` VALUES (860, 3);
INSERT INTO `sys_user_authority` VALUES (861, 3);
INSERT INTO `sys_user_authority` VALUES (862, 3);
INSERT INTO `sys_user_authority` VALUES (863, 3);
INSERT INTO `sys_user_authority` VALUES (864, 3);
INSERT INTO `sys_user_authority` VALUES (865, 3);
INSERT INTO `sys_user_authority` VALUES (866, 3);
INSERT INTO `sys_user_authority` VALUES (867, 3);
INSERT INTO `sys_user_authority` VALUES (868, 3);
INSERT INTO `sys_user_authority` VALUES (869, 3);
INSERT INTO `sys_user_authority` VALUES (870, 3);
INSERT INTO `sys_user_authority` VALUES (871, 3);
INSERT INTO `sys_user_authority` VALUES (872, 3);
INSERT INTO `sys_user_authority` VALUES (873, 3);
INSERT INTO `sys_user_authority` VALUES (874, 3);
INSERT INTO `sys_user_authority` VALUES (875, 3);
INSERT INTO `sys_user_authority` VALUES (876, 3);
INSERT INTO `sys_user_authority` VALUES (877, 3);
INSERT INTO `sys_user_authority` VALUES (878, 3);
INSERT INTO `sys_user_authority` VALUES (879, 3);
INSERT INTO `sys_user_authority` VALUES (880, 3);
INSERT INTO `sys_user_authority` VALUES (881, 3);
INSERT INTO `sys_user_authority` VALUES (882, 3);
INSERT INTO `sys_user_authority` VALUES (883, 3);
INSERT INTO `sys_user_authority` VALUES (884, 3);
INSERT INTO `sys_user_authority` VALUES (885, 3);
INSERT INTO `sys_user_authority` VALUES (886, 3);
INSERT INTO `sys_user_authority` VALUES (887, 3);
INSERT INTO `sys_user_authority` VALUES (888, 3);
INSERT INTO `sys_user_authority` VALUES (889, 3);
INSERT INTO `sys_user_authority` VALUES (890, 3);
INSERT INTO `sys_user_authority` VALUES (891, 3);
INSERT INTO `sys_user_authority` VALUES (892, 3);
INSERT INTO `sys_user_authority` VALUES (893, 3);
INSERT INTO `sys_user_authority` VALUES (894, 3);
INSERT INTO `sys_user_authority` VALUES (895, 3);
INSERT INTO `sys_user_authority` VALUES (896, 3);
INSERT INTO `sys_user_authority` VALUES (897, 3);
INSERT INTO `sys_user_authority` VALUES (898, 3);
INSERT INTO `sys_user_authority` VALUES (899, 3);
INSERT INTO `sys_user_authority` VALUES (900, 3);
INSERT INTO `sys_user_authority` VALUES (901, 3);
INSERT INTO `sys_user_authority` VALUES (902, 3);
INSERT INTO `sys_user_authority` VALUES (903, 3);
INSERT INTO `sys_user_authority` VALUES (904, 3);
INSERT INTO `sys_user_authority` VALUES (905, 3);
INSERT INTO `sys_user_authority` VALUES (906, 3);
INSERT INTO `sys_user_authority` VALUES (907, 3);
INSERT INTO `sys_user_authority` VALUES (908, 3);
INSERT INTO `sys_user_authority` VALUES (909, 3);
INSERT INTO `sys_user_authority` VALUES (910, 3);
INSERT INTO `sys_user_authority` VALUES (911, 3);
INSERT INTO `sys_user_authority` VALUES (912, 3);
INSERT INTO `sys_user_authority` VALUES (913, 3);
INSERT INTO `sys_user_authority` VALUES (914, 3);
INSERT INTO `sys_user_authority` VALUES (915, 3);
INSERT INTO `sys_user_authority` VALUES (916, 3);
INSERT INTO `sys_user_authority` VALUES (917, 3);
INSERT INTO `sys_user_authority` VALUES (918, 3);
INSERT INTO `sys_user_authority` VALUES (919, 3);
INSERT INTO `sys_user_authority` VALUES (920, 3);
INSERT INTO `sys_user_authority` VALUES (921, 3);
INSERT INTO `sys_user_authority` VALUES (922, 3);
INSERT INTO `sys_user_authority` VALUES (923, 3);
INSERT INTO `sys_user_authority` VALUES (924, 3);
INSERT INTO `sys_user_authority` VALUES (925, 3);
INSERT INTO `sys_user_authority` VALUES (926, 3);
INSERT INTO `sys_user_authority` VALUES (927, 3);
INSERT INTO `sys_user_authority` VALUES (928, 3);
INSERT INTO `sys_user_authority` VALUES (929, 3);
INSERT INTO `sys_user_authority` VALUES (930, 3);
INSERT INTO `sys_user_authority` VALUES (931, 3);
INSERT INTO `sys_user_authority` VALUES (932, 3);
INSERT INTO `sys_user_authority` VALUES (933, 3);
INSERT INTO `sys_user_authority` VALUES (934, 3);
INSERT INTO `sys_user_authority` VALUES (935, 3);
INSERT INTO `sys_user_authority` VALUES (936, 3);
INSERT INTO `sys_user_authority` VALUES (937, 3);
INSERT INTO `sys_user_authority` VALUES (938, 3);
INSERT INTO `sys_user_authority` VALUES (939, 3);
INSERT INTO `sys_user_authority` VALUES (940, 3);
INSERT INTO `sys_user_authority` VALUES (941, 3);
INSERT INTO `sys_user_authority` VALUES (942, 3);
INSERT INTO `sys_user_authority` VALUES (943, 3);
INSERT INTO `sys_user_authority` VALUES (944, 3);
INSERT INTO `sys_user_authority` VALUES (945, 3);
INSERT INTO `sys_user_authority` VALUES (946, 3);
INSERT INTO `sys_user_authority` VALUES (947, 3);
INSERT INTO `sys_user_authority` VALUES (948, 3);
INSERT INTO `sys_user_authority` VALUES (949, 3);
INSERT INTO `sys_user_authority` VALUES (950, 3);
INSERT INTO `sys_user_authority` VALUES (951, 3);
INSERT INTO `sys_user_authority` VALUES (952, 3);
INSERT INTO `sys_user_authority` VALUES (953, 3);
INSERT INTO `sys_user_authority` VALUES (954, 3);
INSERT INTO `sys_user_authority` VALUES (955, 3);
INSERT INTO `sys_user_authority` VALUES (956, 3);
INSERT INTO `sys_user_authority` VALUES (957, 3);
INSERT INTO `sys_user_authority` VALUES (958, 3);
INSERT INTO `sys_user_authority` VALUES (959, 3);
INSERT INTO `sys_user_authority` VALUES (960, 3);
INSERT INTO `sys_user_authority` VALUES (961, 3);
INSERT INTO `sys_user_authority` VALUES (962, 3);
INSERT INTO `sys_user_authority` VALUES (963, 3);
INSERT INTO `sys_user_authority` VALUES (964, 3);
INSERT INTO `sys_user_authority` VALUES (965, 3);
INSERT INTO `sys_user_authority` VALUES (966, 3);
INSERT INTO `sys_user_authority` VALUES (967, 3);
INSERT INTO `sys_user_authority` VALUES (968, 3);
INSERT INTO `sys_user_authority` VALUES (969, 3);
INSERT INTO `sys_user_authority` VALUES (970, 3);
INSERT INTO `sys_user_authority` VALUES (971, 3);
INSERT INTO `sys_user_authority` VALUES (972, 3);
INSERT INTO `sys_user_authority` VALUES (973, 3);
INSERT INTO `sys_user_authority` VALUES (974, 3);
INSERT INTO `sys_user_authority` VALUES (975, 3);
INSERT INTO `sys_user_authority` VALUES (976, 3);
INSERT INTO `sys_user_authority` VALUES (977, 3);
INSERT INTO `sys_user_authority` VALUES (978, 3);
INSERT INTO `sys_user_authority` VALUES (979, 3);
INSERT INTO `sys_user_authority` VALUES (980, 3);
INSERT INTO `sys_user_authority` VALUES (981, 3);
INSERT INTO `sys_user_authority` VALUES (982, 3);
INSERT INTO `sys_user_authority` VALUES (983, 3);
INSERT INTO `sys_user_authority` VALUES (984, 3);
INSERT INTO `sys_user_authority` VALUES (985, 3);
INSERT INTO `sys_user_authority` VALUES (986, 3);
INSERT INTO `sys_user_authority` VALUES (987, 3);
INSERT INTO `sys_user_authority` VALUES (988, 3);
INSERT INTO `sys_user_authority` VALUES (989, 3);
INSERT INTO `sys_user_authority` VALUES (990, 3);
INSERT INTO `sys_user_authority` VALUES (991, 3);
INSERT INTO `sys_user_authority` VALUES (992, 3);
INSERT INTO `sys_user_authority` VALUES (993, 3);
INSERT INTO `sys_user_authority` VALUES (994, 3);
INSERT INTO `sys_user_authority` VALUES (995, 3);
INSERT INTO `sys_user_authority` VALUES (996, 3);
INSERT INTO `sys_user_authority` VALUES (997, 3);
INSERT INTO `sys_user_authority` VALUES (998, 3);
INSERT INTO `sys_user_authority` VALUES (999, 3);
INSERT INTO `sys_user_authority` VALUES (1000, 3);
INSERT INTO `sys_user_authority` VALUES (1001, 3);
INSERT INTO `sys_user_authority` VALUES (1002, 3);
INSERT INTO `sys_user_authority` VALUES (1003, 3);
INSERT INTO `sys_user_authority` VALUES (1004, 3);
INSERT INTO `sys_user_authority` VALUES (1005, 3);
INSERT INTO `sys_user_authority` VALUES (1006, 3);
INSERT INTO `sys_user_authority` VALUES (1007, 3);
INSERT INTO `sys_user_authority` VALUES (1008, 3);
INSERT INTO `sys_user_authority` VALUES (1009, 3);
INSERT INTO `sys_user_authority` VALUES (1010, 3);
INSERT INTO `sys_user_authority` VALUES (1011, 3);
INSERT INTO `sys_user_authority` VALUES (1012, 3);
INSERT INTO `sys_user_authority` VALUES (1013, 3);
INSERT INTO `sys_user_authority` VALUES (1014, 3);
INSERT INTO `sys_user_authority` VALUES (1015, 3);
INSERT INTO `sys_user_authority` VALUES (1016, 3);
INSERT INTO `sys_user_authority` VALUES (1017, 3);
INSERT INTO `sys_user_authority` VALUES (1018, 3);
INSERT INTO `sys_user_authority` VALUES (1019, 3);
INSERT INTO `sys_user_authority` VALUES (1020, 3);
INSERT INTO `sys_user_authority` VALUES (1021, 3);
INSERT INTO `sys_user_authority` VALUES (1022, 3);
INSERT INTO `sys_user_authority` VALUES (1023, 3);
INSERT INTO `sys_user_authority` VALUES (1024, 3);
INSERT INTO `sys_user_authority` VALUES (1025, 3);
INSERT INTO `sys_user_authority` VALUES (1026, 3);
INSERT INTO `sys_user_authority` VALUES (1027, 3);
INSERT INTO `sys_user_authority` VALUES (1028, 3);
INSERT INTO `sys_user_authority` VALUES (1029, 3);
INSERT INTO `sys_user_authority` VALUES (1030, 3);
INSERT INTO `sys_user_authority` VALUES (1031, 3);
INSERT INTO `sys_user_authority` VALUES (1032, 3);
INSERT INTO `sys_user_authority` VALUES (1033, 3);
INSERT INTO `sys_user_authority` VALUES (1034, 3);
INSERT INTO `sys_user_authority` VALUES (1035, 3);
INSERT INTO `sys_user_authority` VALUES (1036, 3);
INSERT INTO `sys_user_authority` VALUES (1037, 3);
INSERT INTO `sys_user_authority` VALUES (1038, 3);
INSERT INTO `sys_user_authority` VALUES (1039, 3);
INSERT INTO `sys_user_authority` VALUES (1040, 3);
INSERT INTO `sys_user_authority` VALUES (1041, 3);
INSERT INTO `sys_user_authority` VALUES (1042, 3);
INSERT INTO `sys_user_authority` VALUES (1043, 3);
INSERT INTO `sys_user_authority` VALUES (1044, 3);
INSERT INTO `sys_user_authority` VALUES (1045, 3);
INSERT INTO `sys_user_authority` VALUES (1046, 3);
INSERT INTO `sys_user_authority` VALUES (1047, 3);
INSERT INTO `sys_user_authority` VALUES (1048, 3);
INSERT INTO `sys_user_authority` VALUES (1049, 3);
INSERT INTO `sys_user_authority` VALUES (1050, 3);
INSERT INTO `sys_user_authority` VALUES (1051, 3);
INSERT INTO `sys_user_authority` VALUES (1052, 3);
INSERT INTO `sys_user_authority` VALUES (1053, 3);
INSERT INTO `sys_user_authority` VALUES (1054, 3);
INSERT INTO `sys_user_authority` VALUES (1055, 3);
INSERT INTO `sys_user_authority` VALUES (1056, 3);
INSERT INTO `sys_user_authority` VALUES (1057, 3);
INSERT INTO `sys_user_authority` VALUES (1058, 3);
INSERT INTO `sys_user_authority` VALUES (1059, 3);
INSERT INTO `sys_user_authority` VALUES (1060, 3);
INSERT INTO `sys_user_authority` VALUES (1061, 3);
INSERT INTO `sys_user_authority` VALUES (1062, 3);
INSERT INTO `sys_user_authority` VALUES (1063, 3);
INSERT INTO `sys_user_authority` VALUES (1064, 3);
INSERT INTO `sys_user_authority` VALUES (1065, 3);
INSERT INTO `sys_user_authority` VALUES (1066, 3);
INSERT INTO `sys_user_authority` VALUES (1067, 3);
INSERT INTO `sys_user_authority` VALUES (1068, 3);
INSERT INTO `sys_user_authority` VALUES (1069, 3);
INSERT INTO `sys_user_authority` VALUES (1070, 3);
INSERT INTO `sys_user_authority` VALUES (1071, 3);
INSERT INTO `sys_user_authority` VALUES (1072, 3);
INSERT INTO `sys_user_authority` VALUES (1073, 3);
INSERT INTO `sys_user_authority` VALUES (1074, 3);
INSERT INTO `sys_user_authority` VALUES (1075, 3);
INSERT INTO `sys_user_authority` VALUES (1076, 3);
INSERT INTO `sys_user_authority` VALUES (1077, 3);
INSERT INTO `sys_user_authority` VALUES (1078, 3);
INSERT INTO `sys_user_authority` VALUES (1079, 3);
INSERT INTO `sys_user_authority` VALUES (1080, 3);
INSERT INTO `sys_user_authority` VALUES (1081, 3);
INSERT INTO `sys_user_authority` VALUES (1082, 3);
INSERT INTO `sys_user_authority` VALUES (1083, 3);
INSERT INTO `sys_user_authority` VALUES (1084, 3);
INSERT INTO `sys_user_authority` VALUES (1085, 3);
INSERT INTO `sys_user_authority` VALUES (1086, 3);
INSERT INTO `sys_user_authority` VALUES (1087, 3);
INSERT INTO `sys_user_authority` VALUES (1088, 3);
INSERT INTO `sys_user_authority` VALUES (1089, 3);
INSERT INTO `sys_user_authority` VALUES (1090, 3);
INSERT INTO `sys_user_authority` VALUES (1091, 3);
INSERT INTO `sys_user_authority` VALUES (1092, 3);
INSERT INTO `sys_user_authority` VALUES (1093, 3);
INSERT INTO `sys_user_authority` VALUES (1094, 3);
INSERT INTO `sys_user_authority` VALUES (1095, 3);
INSERT INTO `sys_user_authority` VALUES (1096, 3);
INSERT INTO `sys_user_authority` VALUES (1097, 3);
INSERT INTO `sys_user_authority` VALUES (1098, 3);
INSERT INTO `sys_user_authority` VALUES (1099, 3);
INSERT INTO `sys_user_authority` VALUES (1100, 3);
INSERT INTO `sys_user_authority` VALUES (1101, 3);
INSERT INTO `sys_user_authority` VALUES (1102, 3);
INSERT INTO `sys_user_authority` VALUES (1103, 3);
INSERT INTO `sys_user_authority` VALUES (1104, 3);
INSERT INTO `sys_user_authority` VALUES (1105, 3);
INSERT INTO `sys_user_authority` VALUES (1106, 3);
INSERT INTO `sys_user_authority` VALUES (1107, 3);
INSERT INTO `sys_user_authority` VALUES (1108, 3);
INSERT INTO `sys_user_authority` VALUES (1109, 3);
INSERT INTO `sys_user_authority` VALUES (1110, 3);
INSERT INTO `sys_user_authority` VALUES (1111, 3);
INSERT INTO `sys_user_authority` VALUES (1112, 3);
INSERT INTO `sys_user_authority` VALUES (1113, 3);
INSERT INTO `sys_user_authority` VALUES (1114, 3);
INSERT INTO `sys_user_authority` VALUES (1115, 3);
INSERT INTO `sys_user_authority` VALUES (1116, 3);
INSERT INTO `sys_user_authority` VALUES (1117, 3);
INSERT INTO `sys_user_authority` VALUES (1118, 3);
INSERT INTO `sys_user_authority` VALUES (1119, 3);
INSERT INTO `sys_user_authority` VALUES (1120, 3);
INSERT INTO `sys_user_authority` VALUES (1121, 3);
INSERT INTO `sys_user_authority` VALUES (1122, 3);
INSERT INTO `sys_user_authority` VALUES (1123, 3);
INSERT INTO `sys_user_authority` VALUES (1124, 3);
INSERT INTO `sys_user_authority` VALUES (1125, 3);
INSERT INTO `sys_user_authority` VALUES (1126, 3);
INSERT INTO `sys_user_authority` VALUES (1127, 3);
INSERT INTO `sys_user_authority` VALUES (1128, 3);
INSERT INTO `sys_user_authority` VALUES (1129, 3);
INSERT INTO `sys_user_authority` VALUES (1130, 3);
INSERT INTO `sys_user_authority` VALUES (1131, 3);
INSERT INTO `sys_user_authority` VALUES (1132, 3);
INSERT INTO `sys_user_authority` VALUES (1133, 3);
INSERT INTO `sys_user_authority` VALUES (1134, 3);
INSERT INTO `sys_user_authority` VALUES (1135, 3);
INSERT INTO `sys_user_authority` VALUES (1136, 3);
INSERT INTO `sys_user_authority` VALUES (1137, 3);
INSERT INTO `sys_user_authority` VALUES (1138, 3);
INSERT INTO `sys_user_authority` VALUES (1139, 3);
INSERT INTO `sys_user_authority` VALUES (1140, 3);
INSERT INTO `sys_user_authority` VALUES (1141, 3);
INSERT INTO `sys_user_authority` VALUES (1142, 3);
INSERT INTO `sys_user_authority` VALUES (1143, 3);
INSERT INTO `sys_user_authority` VALUES (1144, 3);
INSERT INTO `sys_user_authority` VALUES (1145, 3);
INSERT INTO `sys_user_authority` VALUES (1146, 3);
INSERT INTO `sys_user_authority` VALUES (1147, 3);
INSERT INTO `sys_user_authority` VALUES (1148, 3);
INSERT INTO `sys_user_authority` VALUES (1149, 3);
INSERT INTO `sys_user_authority` VALUES (1150, 3);
INSERT INTO `sys_user_authority` VALUES (1151, 3);
INSERT INTO `sys_user_authority` VALUES (1152, 3);
INSERT INTO `sys_user_authority` VALUES (1153, 3);
INSERT INTO `sys_user_authority` VALUES (1154, 3);
INSERT INTO `sys_user_authority` VALUES (1155, 3);
INSERT INTO `sys_user_authority` VALUES (1156, 3);
INSERT INTO `sys_user_authority` VALUES (1157, 3);
INSERT INTO `sys_user_authority` VALUES (1158, 3);
INSERT INTO `sys_user_authority` VALUES (1159, 3);
INSERT INTO `sys_user_authority` VALUES (1160, 3);
INSERT INTO `sys_user_authority` VALUES (1161, 3);
INSERT INTO `sys_user_authority` VALUES (1162, 3);
INSERT INTO `sys_user_authority` VALUES (1163, 3);
INSERT INTO `sys_user_authority` VALUES (1164, 3);
INSERT INTO `sys_user_authority` VALUES (1165, 3);
INSERT INTO `sys_user_authority` VALUES (1166, 3);
INSERT INTO `sys_user_authority` VALUES (1167, 3);
INSERT INTO `sys_user_authority` VALUES (1168, 3);
INSERT INTO `sys_user_authority` VALUES (1169, 3);
INSERT INTO `sys_user_authority` VALUES (1170, 3);
INSERT INTO `sys_user_authority` VALUES (1171, 3);
INSERT INTO `sys_user_authority` VALUES (1172, 3);
INSERT INTO `sys_user_authority` VALUES (1173, 3);
INSERT INTO `sys_user_authority` VALUES (1174, 3);
INSERT INTO `sys_user_authority` VALUES (1175, 3);
INSERT INTO `sys_user_authority` VALUES (1176, 3);
INSERT INTO `sys_user_authority` VALUES (1177, 3);
INSERT INTO `sys_user_authority` VALUES (1178, 3);
INSERT INTO `sys_user_authority` VALUES (1179, 3);
INSERT INTO `sys_user_authority` VALUES (1180, 3);
INSERT INTO `sys_user_authority` VALUES (1181, 3);
INSERT INTO `sys_user_authority` VALUES (1182, 3);
INSERT INTO `sys_user_authority` VALUES (1183, 3);
INSERT INTO `sys_user_authority` VALUES (1184, 3);
INSERT INTO `sys_user_authority` VALUES (1185, 3);
INSERT INTO `sys_user_authority` VALUES (1186, 3);
INSERT INTO `sys_user_authority` VALUES (1187, 3);
INSERT INTO `sys_user_authority` VALUES (1188, 3);
INSERT INTO `sys_user_authority` VALUES (1189, 3);
INSERT INTO `sys_user_authority` VALUES (1190, 3);
INSERT INTO `sys_user_authority` VALUES (1191, 3);
INSERT INTO `sys_user_authority` VALUES (1192, 3);
INSERT INTO `sys_user_authority` VALUES (1193, 3);
INSERT INTO `sys_user_authority` VALUES (1194, 3);
INSERT INTO `sys_user_authority` VALUES (1195, 3);
INSERT INTO `sys_user_authority` VALUES (1196, 3);
INSERT INTO `sys_user_authority` VALUES (1197, 3);
INSERT INTO `sys_user_authority` VALUES (1198, 3);
INSERT INTO `sys_user_authority` VALUES (1199, 3);
INSERT INTO `sys_user_authority` VALUES (1200, 3);
INSERT INTO `sys_user_authority` VALUES (1201, 3);
INSERT INTO `sys_user_authority` VALUES (1202, 3);
INSERT INTO `sys_user_authority` VALUES (1203, 3);
INSERT INTO `sys_user_authority` VALUES (1204, 3);
INSERT INTO `sys_user_authority` VALUES (1205, 3);
INSERT INTO `sys_user_authority` VALUES (1206, 3);
INSERT INTO `sys_user_authority` VALUES (1207, 3);
INSERT INTO `sys_user_authority` VALUES (1208, 3);
INSERT INTO `sys_user_authority` VALUES (1209, 3);
INSERT INTO `sys_user_authority` VALUES (1210, 3);
INSERT INTO `sys_user_authority` VALUES (1211, 3);
INSERT INTO `sys_user_authority` VALUES (1212, 3);
INSERT INTO `sys_user_authority` VALUES (1213, 3);
INSERT INTO `sys_user_authority` VALUES (1214, 3);
INSERT INTO `sys_user_authority` VALUES (1215, 3);
INSERT INTO `sys_user_authority` VALUES (1216, 3);
INSERT INTO `sys_user_authority` VALUES (1217, 3);
INSERT INTO `sys_user_authority` VALUES (1218, 3);
INSERT INTO `sys_user_authority` VALUES (1219, 3);
INSERT INTO `sys_user_authority` VALUES (1220, 3);
INSERT INTO `sys_user_authority` VALUES (1221, 3);
INSERT INTO `sys_user_authority` VALUES (1222, 3);
INSERT INTO `sys_user_authority` VALUES (1223, 3);
INSERT INTO `sys_user_authority` VALUES (1224, 3);
INSERT INTO `sys_user_authority` VALUES (1225, 3);
INSERT INTO `sys_user_authority` VALUES (1226, 3);
INSERT INTO `sys_user_authority` VALUES (1227, 3);
INSERT INTO `sys_user_authority` VALUES (1228, 3);
INSERT INTO `sys_user_authority` VALUES (1229, 3);
INSERT INTO `sys_user_authority` VALUES (1230, 3);
INSERT INTO `sys_user_authority` VALUES (1231, 3);
INSERT INTO `sys_user_authority` VALUES (1232, 3);
INSERT INTO `sys_user_authority` VALUES (1233, 3);
INSERT INTO `sys_user_authority` VALUES (1234, 3);
INSERT INTO `sys_user_authority` VALUES (1235, 3);
INSERT INTO `sys_user_authority` VALUES (1236, 3);
INSERT INTO `sys_user_authority` VALUES (1237, 3);
INSERT INTO `sys_user_authority` VALUES (1238, 3);
INSERT INTO `sys_user_authority` VALUES (1239, 3);
INSERT INTO `sys_user_authority` VALUES (1240, 3);
INSERT INTO `sys_user_authority` VALUES (1241, 3);
INSERT INTO `sys_user_authority` VALUES (1242, 3);
INSERT INTO `sys_user_authority` VALUES (1243, 3);
INSERT INTO `sys_user_authority` VALUES (1244, 3);
INSERT INTO `sys_user_authority` VALUES (1245, 3);
INSERT INTO `sys_user_authority` VALUES (1246, 3);
INSERT INTO `sys_user_authority` VALUES (1247, 3);
INSERT INTO `sys_user_authority` VALUES (1248, 3);
INSERT INTO `sys_user_authority` VALUES (1249, 3);
INSERT INTO `sys_user_authority` VALUES (1250, 3);
INSERT INTO `sys_user_authority` VALUES (1251, 3);
INSERT INTO `sys_user_authority` VALUES (1252, 3);
INSERT INTO `sys_user_authority` VALUES (1253, 3);
INSERT INTO `sys_user_authority` VALUES (1254, 3);
INSERT INTO `sys_user_authority` VALUES (1255, 3);
INSERT INTO `sys_user_authority` VALUES (1256, 3);
INSERT INTO `sys_user_authority` VALUES (1257, 3);
INSERT INTO `sys_user_authority` VALUES (1258, 3);
INSERT INTO `sys_user_authority` VALUES (1259, 3);
INSERT INTO `sys_user_authority` VALUES (1260, 3);
INSERT INTO `sys_user_authority` VALUES (1261, 3);
INSERT INTO `sys_user_authority` VALUES (1262, 3);
INSERT INTO `sys_user_authority` VALUES (1263, 3);
INSERT INTO `sys_user_authority` VALUES (1264, 3);
INSERT INTO `sys_user_authority` VALUES (1265, 3);
INSERT INTO `sys_user_authority` VALUES (1266, 3);
INSERT INTO `sys_user_authority` VALUES (1267, 3);
INSERT INTO `sys_user_authority` VALUES (1268, 3);
INSERT INTO `sys_user_authority` VALUES (1269, 3);
INSERT INTO `sys_user_authority` VALUES (1270, 3);
INSERT INTO `sys_user_authority` VALUES (1271, 3);
INSERT INTO `sys_user_authority` VALUES (1272, 3);
INSERT INTO `sys_user_authority` VALUES (1273, 3);
INSERT INTO `sys_user_authority` VALUES (1274, 3);
INSERT INTO `sys_user_authority` VALUES (1275, 3);
INSERT INTO `sys_user_authority` VALUES (1276, 3);
INSERT INTO `sys_user_authority` VALUES (1277, 3);
INSERT INTO `sys_user_authority` VALUES (1278, 3);
INSERT INTO `sys_user_authority` VALUES (1279, 3);
INSERT INTO `sys_user_authority` VALUES (1280, 3);
INSERT INTO `sys_user_authority` VALUES (1281, 3);
INSERT INTO `sys_user_authority` VALUES (1282, 3);
INSERT INTO `sys_user_authority` VALUES (1283, 3);
INSERT INTO `sys_user_authority` VALUES (1284, 3);
INSERT INTO `sys_user_authority` VALUES (1285, 3);
INSERT INTO `sys_user_authority` VALUES (1286, 3);
INSERT INTO `sys_user_authority` VALUES (1287, 3);
INSERT INTO `sys_user_authority` VALUES (1288, 3);
INSERT INTO `sys_user_authority` VALUES (1289, 3);
INSERT INTO `sys_user_authority` VALUES (1290, 3);
INSERT INTO `sys_user_authority` VALUES (1291, 3);
INSERT INTO `sys_user_authority` VALUES (1292, 3);
INSERT INTO `sys_user_authority` VALUES (1293, 3);
INSERT INTO `sys_user_authority` VALUES (1294, 3);
INSERT INTO `sys_user_authority` VALUES (1295, 3);
INSERT INTO `sys_user_authority` VALUES (1296, 3);
INSERT INTO `sys_user_authority` VALUES (1297, 3);
INSERT INTO `sys_user_authority` VALUES (1298, 3);
INSERT INTO `sys_user_authority` VALUES (1299, 3);
INSERT INTO `sys_user_authority` VALUES (1300, 3);
INSERT INTO `sys_user_authority` VALUES (1301, 3);
INSERT INTO `sys_user_authority` VALUES (1302, 3);
INSERT INTO `sys_user_authority` VALUES (1303, 3);
INSERT INTO `sys_user_authority` VALUES (1304, 3);
INSERT INTO `sys_user_authority` VALUES (1305, 3);
INSERT INTO `sys_user_authority` VALUES (1306, 3);
INSERT INTO `sys_user_authority` VALUES (1307, 3);
INSERT INTO `sys_user_authority` VALUES (1308, 3);
INSERT INTO `sys_user_authority` VALUES (1309, 3);
INSERT INTO `sys_user_authority` VALUES (1310, 3);
INSERT INTO `sys_user_authority` VALUES (1311, 3);
INSERT INTO `sys_user_authority` VALUES (1312, 3);
INSERT INTO `sys_user_authority` VALUES (1313, 3);
INSERT INTO `sys_user_authority` VALUES (1314, 3);
INSERT INTO `sys_user_authority` VALUES (1315, 3);
INSERT INTO `sys_user_authority` VALUES (1316, 3);
INSERT INTO `sys_user_authority` VALUES (1317, 3);
INSERT INTO `sys_user_authority` VALUES (1318, 3);
INSERT INTO `sys_user_authority` VALUES (1319, 3);
INSERT INTO `sys_user_authority` VALUES (1320, 3);
INSERT INTO `sys_user_authority` VALUES (1321, 3);
INSERT INTO `sys_user_authority` VALUES (1322, 3);
INSERT INTO `sys_user_authority` VALUES (1323, 3);
INSERT INTO `sys_user_authority` VALUES (1324, 3);
INSERT INTO `sys_user_authority` VALUES (1325, 3);
INSERT INTO `sys_user_authority` VALUES (1326, 3);
INSERT INTO `sys_user_authority` VALUES (1327, 3);
INSERT INTO `sys_user_authority` VALUES (1328, 3);
INSERT INTO `sys_user_authority` VALUES (1329, 3);
INSERT INTO `sys_user_authority` VALUES (1330, 3);
INSERT INTO `sys_user_authority` VALUES (1331, 3);
INSERT INTO `sys_user_authority` VALUES (1332, 3);
INSERT INTO `sys_user_authority` VALUES (1333, 3);
INSERT INTO `sys_user_authority` VALUES (1334, 3);
INSERT INTO `sys_user_authority` VALUES (1335, 3);
INSERT INTO `sys_user_authority` VALUES (1336, 3);
INSERT INTO `sys_user_authority` VALUES (1337, 3);
INSERT INTO `sys_user_authority` VALUES (1338, 3);
INSERT INTO `sys_user_authority` VALUES (1339, 3);
INSERT INTO `sys_user_authority` VALUES (1340, 3);
INSERT INTO `sys_user_authority` VALUES (1341, 3);
INSERT INTO `sys_user_authority` VALUES (1342, 3);
INSERT INTO `sys_user_authority` VALUES (1343, 3);
INSERT INTO `sys_user_authority` VALUES (1344, 3);
INSERT INTO `sys_user_authority` VALUES (1345, 3);
INSERT INTO `sys_user_authority` VALUES (1346, 3);
INSERT INTO `sys_user_authority` VALUES (1347, 3);
INSERT INTO `sys_user_authority` VALUES (1348, 3);
INSERT INTO `sys_user_authority` VALUES (1349, 3);
INSERT INTO `sys_user_authority` VALUES (1350, 3);
INSERT INTO `sys_user_authority` VALUES (1351, 3);
INSERT INTO `sys_user_authority` VALUES (1352, 3);
INSERT INTO `sys_user_authority` VALUES (1353, 3);
INSERT INTO `sys_user_authority` VALUES (1354, 3);
INSERT INTO `sys_user_authority` VALUES (1355, 3);
INSERT INTO `sys_user_authority` VALUES (1356, 3);
INSERT INTO `sys_user_authority` VALUES (1357, 3);
INSERT INTO `sys_user_authority` VALUES (1358, 3);
INSERT INTO `sys_user_authority` VALUES (1359, 3);
INSERT INTO `sys_user_authority` VALUES (1360, 3);
INSERT INTO `sys_user_authority` VALUES (1361, 3);
INSERT INTO `sys_user_authority` VALUES (1362, 3);
INSERT INTO `sys_user_authority` VALUES (1363, 3);
INSERT INTO `sys_user_authority` VALUES (1364, 3);
INSERT INTO `sys_user_authority` VALUES (1365, 3);
INSERT INTO `sys_user_authority` VALUES (1366, 3);
INSERT INTO `sys_user_authority` VALUES (1367, 3);
INSERT INTO `sys_user_authority` VALUES (1368, 3);
INSERT INTO `sys_user_authority` VALUES (1369, 3);
INSERT INTO `sys_user_authority` VALUES (1370, 3);
INSERT INTO `sys_user_authority` VALUES (1371, 3);
INSERT INTO `sys_user_authority` VALUES (1372, 3);
INSERT INTO `sys_user_authority` VALUES (1373, 3);
INSERT INTO `sys_user_authority` VALUES (1374, 3);
INSERT INTO `sys_user_authority` VALUES (1375, 3);
INSERT INTO `sys_user_authority` VALUES (1376, 3);
INSERT INTO `sys_user_authority` VALUES (1377, 3);
INSERT INTO `sys_user_authority` VALUES (1378, 3);
INSERT INTO `sys_user_authority` VALUES (1379, 3);
INSERT INTO `sys_user_authority` VALUES (1380, 3);
INSERT INTO `sys_user_authority` VALUES (1381, 3);
INSERT INTO `sys_user_authority` VALUES (1382, 3);
INSERT INTO `sys_user_authority` VALUES (1383, 3);
INSERT INTO `sys_user_authority` VALUES (1384, 3);
INSERT INTO `sys_user_authority` VALUES (1385, 3);
INSERT INTO `sys_user_authority` VALUES (1386, 3);
INSERT INTO `sys_user_authority` VALUES (1387, 3);
INSERT INTO `sys_user_authority` VALUES (1388, 3);
INSERT INTO `sys_user_authority` VALUES (1389, 3);
INSERT INTO `sys_user_authority` VALUES (1390, 3);
INSERT INTO `sys_user_authority` VALUES (1391, 3);
INSERT INTO `sys_user_authority` VALUES (1392, 3);
INSERT INTO `sys_user_authority` VALUES (1393, 3);
INSERT INTO `sys_user_authority` VALUES (1394, 3);
INSERT INTO `sys_user_authority` VALUES (1395, 3);
INSERT INTO `sys_user_authority` VALUES (1396, 3);
INSERT INTO `sys_user_authority` VALUES (1397, 3);
INSERT INTO `sys_user_authority` VALUES (1398, 3);
INSERT INTO `sys_user_authority` VALUES (1399, 3);
INSERT INTO `sys_user_authority` VALUES (1400, 3);
INSERT INTO `sys_user_authority` VALUES (1401, 3);
INSERT INTO `sys_user_authority` VALUES (1402, 3);
INSERT INTO `sys_user_authority` VALUES (1403, 3);
INSERT INTO `sys_user_authority` VALUES (1404, 3);
INSERT INTO `sys_user_authority` VALUES (1405, 3);
INSERT INTO `sys_user_authority` VALUES (1406, 3);
INSERT INTO `sys_user_authority` VALUES (1407, 3);
INSERT INTO `sys_user_authority` VALUES (1408, 3);
INSERT INTO `sys_user_authority` VALUES (1409, 3);
INSERT INTO `sys_user_authority` VALUES (1410, 3);
INSERT INTO `sys_user_authority` VALUES (1411, 3);
INSERT INTO `sys_user_authority` VALUES (1412, 3);
INSERT INTO `sys_user_authority` VALUES (1413, 3);
INSERT INTO `sys_user_authority` VALUES (1414, 3);
INSERT INTO `sys_user_authority` VALUES (1415, 3);
INSERT INTO `sys_user_authority` VALUES (1416, 3);
INSERT INTO `sys_user_authority` VALUES (1417, 3);
INSERT INTO `sys_user_authority` VALUES (1418, 3);
INSERT INTO `sys_user_authority` VALUES (1419, 3);
INSERT INTO `sys_user_authority` VALUES (1420, 3);
INSERT INTO `sys_user_authority` VALUES (1421, 3);
INSERT INTO `sys_user_authority` VALUES (1422, 3);
INSERT INTO `sys_user_authority` VALUES (1423, 3);
INSERT INTO `sys_user_authority` VALUES (1424, 3);
INSERT INTO `sys_user_authority` VALUES (1425, 3);
INSERT INTO `sys_user_authority` VALUES (1426, 3);
INSERT INTO `sys_user_authority` VALUES (1427, 3);
INSERT INTO `sys_user_authority` VALUES (1428, 3);
INSERT INTO `sys_user_authority` VALUES (1429, 3);
INSERT INTO `sys_user_authority` VALUES (1430, 3);
INSERT INTO `sys_user_authority` VALUES (1431, 3);
INSERT INTO `sys_user_authority` VALUES (1432, 3);
INSERT INTO `sys_user_authority` VALUES (1433, 3);
INSERT INTO `sys_user_authority` VALUES (1434, 3);
INSERT INTO `sys_user_authority` VALUES (1435, 3);
INSERT INTO `sys_user_authority` VALUES (1436, 3);
INSERT INTO `sys_user_authority` VALUES (1437, 3);
INSERT INTO `sys_user_authority` VALUES (1438, 3);
INSERT INTO `sys_user_authority` VALUES (1439, 3);
INSERT INTO `sys_user_authority` VALUES (1440, 3);
INSERT INTO `sys_user_authority` VALUES (1441, 3);
INSERT INTO `sys_user_authority` VALUES (1442, 3);
INSERT INTO `sys_user_authority` VALUES (1443, 3);
INSERT INTO `sys_user_authority` VALUES (1444, 3);
INSERT INTO `sys_user_authority` VALUES (1445, 3);
INSERT INTO `sys_user_authority` VALUES (1446, 3);
INSERT INTO `sys_user_authority` VALUES (1447, 3);
INSERT INTO `sys_user_authority` VALUES (1448, 3);
INSERT INTO `sys_user_authority` VALUES (1449, 3);
INSERT INTO `sys_user_authority` VALUES (1450, 3);
INSERT INTO `sys_user_authority` VALUES (1451, 3);
INSERT INTO `sys_user_authority` VALUES (1452, 3);
INSERT INTO `sys_user_authority` VALUES (1453, 3);
INSERT INTO `sys_user_authority` VALUES (1454, 3);
INSERT INTO `sys_user_authority` VALUES (1455, 3);
INSERT INTO `sys_user_authority` VALUES (1456, 3);
INSERT INTO `sys_user_authority` VALUES (1457, 3);
INSERT INTO `sys_user_authority` VALUES (1458, 3);
INSERT INTO `sys_user_authority` VALUES (1459, 3);
INSERT INTO `sys_user_authority` VALUES (1460, 3);
INSERT INTO `sys_user_authority` VALUES (1461, 3);
INSERT INTO `sys_user_authority` VALUES (1462, 3);
INSERT INTO `sys_user_authority` VALUES (1463, 3);
INSERT INTO `sys_user_authority` VALUES (1464, 3);
INSERT INTO `sys_user_authority` VALUES (1465, 3);
INSERT INTO `sys_user_authority` VALUES (1466, 3);
INSERT INTO `sys_user_authority` VALUES (1467, 3);
INSERT INTO `sys_user_authority` VALUES (1468, 3);
INSERT INTO `sys_user_authority` VALUES (1469, 3);
INSERT INTO `sys_user_authority` VALUES (1470, 3);
INSERT INTO `sys_user_authority` VALUES (1471, 3);
INSERT INTO `sys_user_authority` VALUES (1472, 3);
INSERT INTO `sys_user_authority` VALUES (1473, 3);
INSERT INTO `sys_user_authority` VALUES (1474, 3);
INSERT INTO `sys_user_authority` VALUES (1475, 3);
INSERT INTO `sys_user_authority` VALUES (1476, 3);
INSERT INTO `sys_user_authority` VALUES (1477, 3);
INSERT INTO `sys_user_authority` VALUES (1478, 3);
INSERT INTO `sys_user_authority` VALUES (1479, 3);
INSERT INTO `sys_user_authority` VALUES (1480, 3);
INSERT INTO `sys_user_authority` VALUES (1481, 3);
INSERT INTO `sys_user_authority` VALUES (1482, 3);
INSERT INTO `sys_user_authority` VALUES (1483, 3);
INSERT INTO `sys_user_authority` VALUES (1484, 3);
INSERT INTO `sys_user_authority` VALUES (1485, 3);
INSERT INTO `sys_user_authority` VALUES (1486, 3);
INSERT INTO `sys_user_authority` VALUES (1487, 3);
INSERT INTO `sys_user_authority` VALUES (1488, 3);
INSERT INTO `sys_user_authority` VALUES (1489, 3);
INSERT INTO `sys_user_authority` VALUES (1490, 3);
INSERT INTO `sys_user_authority` VALUES (1491, 3);
INSERT INTO `sys_user_authority` VALUES (1492, 3);
INSERT INTO `sys_user_authority` VALUES (1493, 3);
INSERT INTO `sys_user_authority` VALUES (1494, 3);
INSERT INTO `sys_user_authority` VALUES (1495, 3);
INSERT INTO `sys_user_authority` VALUES (1496, 3);
INSERT INTO `sys_user_authority` VALUES (1497, 3);
INSERT INTO `sys_user_authority` VALUES (1498, 3);
INSERT INTO `sys_user_authority` VALUES (1499, 3);
INSERT INTO `sys_user_authority` VALUES (1500, 3);
INSERT INTO `sys_user_authority` VALUES (1501, 3);
INSERT INTO `sys_user_authority` VALUES (1502, 3);
INSERT INTO `sys_user_authority` VALUES (1503, 3);
INSERT INTO `sys_user_authority` VALUES (1504, 3);
INSERT INTO `sys_user_authority` VALUES (1505, 3);
INSERT INTO `sys_user_authority` VALUES (1506, 3);
INSERT INTO `sys_user_authority` VALUES (1507, 3);
INSERT INTO `sys_user_authority` VALUES (1508, 3);
INSERT INTO `sys_user_authority` VALUES (1509, 3);
INSERT INTO `sys_user_authority` VALUES (1510, 3);
INSERT INTO `sys_user_authority` VALUES (1511, 3);
INSERT INTO `sys_user_authority` VALUES (1512, 3);
INSERT INTO `sys_user_authority` VALUES (1513, 3);
INSERT INTO `sys_user_authority` VALUES (1514, 3);
INSERT INTO `sys_user_authority` VALUES (1515, 3);
INSERT INTO `sys_user_authority` VALUES (1516, 3);
INSERT INTO `sys_user_authority` VALUES (1517, 3);
INSERT INTO `sys_user_authority` VALUES (1518, 3);
INSERT INTO `sys_user_authority` VALUES (1519, 3);
INSERT INTO `sys_user_authority` VALUES (1520, 3);
INSERT INTO `sys_user_authority` VALUES (1521, 3);
INSERT INTO `sys_user_authority` VALUES (1522, 3);
INSERT INTO `sys_user_authority` VALUES (1523, 3);
INSERT INTO `sys_user_authority` VALUES (1524, 3);
INSERT INTO `sys_user_authority` VALUES (1525, 3);
INSERT INTO `sys_user_authority` VALUES (1526, 3);
INSERT INTO `sys_user_authority` VALUES (1527, 3);
INSERT INTO `sys_user_authority` VALUES (1528, 3);
INSERT INTO `sys_user_authority` VALUES (1529, 3);
INSERT INTO `sys_user_authority` VALUES (1530, 3);
INSERT INTO `sys_user_authority` VALUES (1531, 3);
INSERT INTO `sys_user_authority` VALUES (1532, 3);
INSERT INTO `sys_user_authority` VALUES (1533, 3);
INSERT INTO `sys_user_authority` VALUES (1534, 3);
INSERT INTO `sys_user_authority` VALUES (1535, 3);
INSERT INTO `sys_user_authority` VALUES (1536, 3);
INSERT INTO `sys_user_authority` VALUES (1537, 3);
INSERT INTO `sys_user_authority` VALUES (1538, 3);
INSERT INTO `sys_user_authority` VALUES (1539, 3);
INSERT INTO `sys_user_authority` VALUES (1540, 3);
INSERT INTO `sys_user_authority` VALUES (1541, 3);
INSERT INTO `sys_user_authority` VALUES (1542, 3);
INSERT INTO `sys_user_authority` VALUES (1543, 3);
INSERT INTO `sys_user_authority` VALUES (1544, 3);
INSERT INTO `sys_user_authority` VALUES (1545, 3);
INSERT INTO `sys_user_authority` VALUES (1546, 3);
INSERT INTO `sys_user_authority` VALUES (1547, 3);
INSERT INTO `sys_user_authority` VALUES (1548, 3);
INSERT INTO `sys_user_authority` VALUES (1549, 3);
INSERT INTO `sys_user_authority` VALUES (1550, 3);
INSERT INTO `sys_user_authority` VALUES (1551, 3);
INSERT INTO `sys_user_authority` VALUES (1552, 3);
INSERT INTO `sys_user_authority` VALUES (1553, 3);
INSERT INTO `sys_user_authority` VALUES (1554, 3);
INSERT INTO `sys_user_authority` VALUES (1555, 3);
INSERT INTO `sys_user_authority` VALUES (1556, 3);
INSERT INTO `sys_user_authority` VALUES (1557, 3);
INSERT INTO `sys_user_authority` VALUES (1558, 3);
INSERT INTO `sys_user_authority` VALUES (1559, 3);
INSERT INTO `sys_user_authority` VALUES (1560, 3);
INSERT INTO `sys_user_authority` VALUES (1561, 3);
INSERT INTO `sys_user_authority` VALUES (1562, 3);
INSERT INTO `sys_user_authority` VALUES (1563, 3);
INSERT INTO `sys_user_authority` VALUES (1564, 3);
INSERT INTO `sys_user_authority` VALUES (1565, 3);
INSERT INTO `sys_user_authority` VALUES (1566, 3);
INSERT INTO `sys_user_authority` VALUES (1567, 3);
INSERT INTO `sys_user_authority` VALUES (1568, 3);
INSERT INTO `sys_user_authority` VALUES (1569, 3);
INSERT INTO `sys_user_authority` VALUES (1570, 3);
INSERT INTO `sys_user_authority` VALUES (1571, 3);
INSERT INTO `sys_user_authority` VALUES (1572, 3);
INSERT INTO `sys_user_authority` VALUES (1573, 3);
INSERT INTO `sys_user_authority` VALUES (1574, 3);
INSERT INTO `sys_user_authority` VALUES (1575, 3);
INSERT INTO `sys_user_authority` VALUES (1576, 3);
INSERT INTO `sys_user_authority` VALUES (1577, 3);
INSERT INTO `sys_user_authority` VALUES (1578, 3);
INSERT INTO `sys_user_authority` VALUES (1579, 3);
INSERT INTO `sys_user_authority` VALUES (1580, 3);
INSERT INTO `sys_user_authority` VALUES (1581, 3);
INSERT INTO `sys_user_authority` VALUES (1582, 3);
INSERT INTO `sys_user_authority` VALUES (1583, 3);
INSERT INTO `sys_user_authority` VALUES (1584, 3);
INSERT INTO `sys_user_authority` VALUES (1585, 3);
INSERT INTO `sys_user_authority` VALUES (1586, 3);
INSERT INTO `sys_user_authority` VALUES (1587, 3);
INSERT INTO `sys_user_authority` VALUES (1588, 3);
INSERT INTO `sys_user_authority` VALUES (1589, 3);
INSERT INTO `sys_user_authority` VALUES (1590, 3);
INSERT INTO `sys_user_authority` VALUES (1591, 3);
INSERT INTO `sys_user_authority` VALUES (1592, 3);
INSERT INTO `sys_user_authority` VALUES (1593, 3);
INSERT INTO `sys_user_authority` VALUES (1594, 3);
INSERT INTO `sys_user_authority` VALUES (1595, 3);
INSERT INTO `sys_user_authority` VALUES (1596, 3);
INSERT INTO `sys_user_authority` VALUES (1597, 3);
INSERT INTO `sys_user_authority` VALUES (1598, 3);
INSERT INTO `sys_user_authority` VALUES (1599, 3);
INSERT INTO `sys_user_authority` VALUES (1600, 3);
INSERT INTO `sys_user_authority` VALUES (1601, 3);
INSERT INTO `sys_user_authority` VALUES (1602, 3);
INSERT INTO `sys_user_authority` VALUES (1603, 3);
INSERT INTO `sys_user_authority` VALUES (1604, 3);
INSERT INTO `sys_user_authority` VALUES (1605, 3);
INSERT INTO `sys_user_authority` VALUES (1606, 3);
INSERT INTO `sys_user_authority` VALUES (1607, 3);
INSERT INTO `sys_user_authority` VALUES (1608, 3);
INSERT INTO `sys_user_authority` VALUES (1609, 3);
INSERT INTO `sys_user_authority` VALUES (1610, 3);
INSERT INTO `sys_user_authority` VALUES (1611, 3);
INSERT INTO `sys_user_authority` VALUES (1612, 3);
INSERT INTO `sys_user_authority` VALUES (1613, 3);
INSERT INTO `sys_user_authority` VALUES (1614, 3);
INSERT INTO `sys_user_authority` VALUES (1615, 3);
INSERT INTO `sys_user_authority` VALUES (1616, 3);
INSERT INTO `sys_user_authority` VALUES (1617, 3);
INSERT INTO `sys_user_authority` VALUES (1618, 3);
INSERT INTO `sys_user_authority` VALUES (1619, 3);
INSERT INTO `sys_user_authority` VALUES (1620, 3);
INSERT INTO `sys_user_authority` VALUES (1621, 3);
INSERT INTO `sys_user_authority` VALUES (1622, 3);
INSERT INTO `sys_user_authority` VALUES (1623, 3);
INSERT INTO `sys_user_authority` VALUES (1624, 3);
INSERT INTO `sys_user_authority` VALUES (1625, 3);
INSERT INTO `sys_user_authority` VALUES (1626, 3);
INSERT INTO `sys_user_authority` VALUES (1627, 3);
INSERT INTO `sys_user_authority` VALUES (1628, 3);
INSERT INTO `sys_user_authority` VALUES (1629, 3);
INSERT INTO `sys_user_authority` VALUES (1630, 3);
INSERT INTO `sys_user_authority` VALUES (1631, 3);
INSERT INTO `sys_user_authority` VALUES (1632, 3);
INSERT INTO `sys_user_authority` VALUES (1633, 3);
INSERT INTO `sys_user_authority` VALUES (1634, 3);
INSERT INTO `sys_user_authority` VALUES (1635, 3);
INSERT INTO `sys_user_authority` VALUES (1636, 3);
INSERT INTO `sys_user_authority` VALUES (1637, 3);
INSERT INTO `sys_user_authority` VALUES (1638, 3);
INSERT INTO `sys_user_authority` VALUES (1639, 3);
INSERT INTO `sys_user_authority` VALUES (1640, 3);
INSERT INTO `sys_user_authority` VALUES (1641, 3);
INSERT INTO `sys_user_authority` VALUES (1642, 3);
INSERT INTO `sys_user_authority` VALUES (1643, 3);
INSERT INTO `sys_user_authority` VALUES (1644, 3);
INSERT INTO `sys_user_authority` VALUES (1645, 3);
INSERT INTO `sys_user_authority` VALUES (1646, 3);
INSERT INTO `sys_user_authority` VALUES (1647, 3);
INSERT INTO `sys_user_authority` VALUES (1648, 3);
INSERT INTO `sys_user_authority` VALUES (1649, 3);
INSERT INTO `sys_user_authority` VALUES (1650, 3);
INSERT INTO `sys_user_authority` VALUES (1651, 3);
INSERT INTO `sys_user_authority` VALUES (1652, 3);
INSERT INTO `sys_user_authority` VALUES (1653, 3);
INSERT INTO `sys_user_authority` VALUES (1654, 3);
INSERT INTO `sys_user_authority` VALUES (1655, 3);
INSERT INTO `sys_user_authority` VALUES (1656, 3);
INSERT INTO `sys_user_authority` VALUES (1657, 3);
INSERT INTO `sys_user_authority` VALUES (1658, 3);
INSERT INTO `sys_user_authority` VALUES (1659, 3);
INSERT INTO `sys_user_authority` VALUES (1660, 3);
INSERT INTO `sys_user_authority` VALUES (1661, 3);
INSERT INTO `sys_user_authority` VALUES (1662, 3);
INSERT INTO `sys_user_authority` VALUES (1663, 3);
INSERT INTO `sys_user_authority` VALUES (1664, 3);
INSERT INTO `sys_user_authority` VALUES (1665, 3);
INSERT INTO `sys_user_authority` VALUES (1666, 3);
INSERT INTO `sys_user_authority` VALUES (1667, 3);
INSERT INTO `sys_user_authority` VALUES (1668, 3);
INSERT INTO `sys_user_authority` VALUES (1669, 3);
INSERT INTO `sys_user_authority` VALUES (1670, 3);
INSERT INTO `sys_user_authority` VALUES (1671, 3);
INSERT INTO `sys_user_authority` VALUES (1672, 3);
INSERT INTO `sys_user_authority` VALUES (1673, 3);
INSERT INTO `sys_user_authority` VALUES (1674, 3);
INSERT INTO `sys_user_authority` VALUES (1675, 3);
INSERT INTO `sys_user_authority` VALUES (1676, 3);
INSERT INTO `sys_user_authority` VALUES (1677, 3);
INSERT INTO `sys_user_authority` VALUES (1678, 3);
INSERT INTO `sys_user_authority` VALUES (1679, 3);
INSERT INTO `sys_user_authority` VALUES (1680, 3);
INSERT INTO `sys_user_authority` VALUES (1681, 3);
INSERT INTO `sys_user_authority` VALUES (1682, 3);
INSERT INTO `sys_user_authority` VALUES (1683, 3);
INSERT INTO `sys_user_authority` VALUES (1684, 3);
INSERT INTO `sys_user_authority` VALUES (1685, 3);
INSERT INTO `sys_user_authority` VALUES (1686, 3);
INSERT INTO `sys_user_authority` VALUES (1687, 3);
INSERT INTO `sys_user_authority` VALUES (1688, 3);
INSERT INTO `sys_user_authority` VALUES (1689, 3);
INSERT INTO `sys_user_authority` VALUES (1690, 3);
INSERT INTO `sys_user_authority` VALUES (1691, 3);
INSERT INTO `sys_user_authority` VALUES (1692, 3);
INSERT INTO `sys_user_authority` VALUES (1693, 3);
INSERT INTO `sys_user_authority` VALUES (1694, 3);
INSERT INTO `sys_user_authority` VALUES (1695, 3);
INSERT INTO `sys_user_authority` VALUES (1696, 3);
INSERT INTO `sys_user_authority` VALUES (1697, 3);
INSERT INTO `sys_user_authority` VALUES (1698, 3);
INSERT INTO `sys_user_authority` VALUES (1699, 3);
INSERT INTO `sys_user_authority` VALUES (1700, 3);
INSERT INTO `sys_user_authority` VALUES (1701, 3);
INSERT INTO `sys_user_authority` VALUES (1702, 3);
INSERT INTO `sys_user_authority` VALUES (1703, 3);
INSERT INTO `sys_user_authority` VALUES (1704, 3);
INSERT INTO `sys_user_authority` VALUES (1705, 3);
INSERT INTO `sys_user_authority` VALUES (1706, 3);
INSERT INTO `sys_user_authority` VALUES (1707, 3);
INSERT INTO `sys_user_authority` VALUES (1708, 3);
INSERT INTO `sys_user_authority` VALUES (1709, 3);
INSERT INTO `sys_user_authority` VALUES (1710, 3);
INSERT INTO `sys_user_authority` VALUES (1711, 3);
INSERT INTO `sys_user_authority` VALUES (1712, 3);
INSERT INTO `sys_user_authority` VALUES (1713, 3);
INSERT INTO `sys_user_authority` VALUES (1714, 3);
INSERT INTO `sys_user_authority` VALUES (1715, 3);
INSERT INTO `sys_user_authority` VALUES (1716, 3);
INSERT INTO `sys_user_authority` VALUES (1717, 3);
INSERT INTO `sys_user_authority` VALUES (1718, 3);
INSERT INTO `sys_user_authority` VALUES (1719, 3);
INSERT INTO `sys_user_authority` VALUES (1720, 3);
INSERT INTO `sys_user_authority` VALUES (1721, 3);
INSERT INTO `sys_user_authority` VALUES (1722, 3);
INSERT INTO `sys_user_authority` VALUES (1723, 3);
INSERT INTO `sys_user_authority` VALUES (1724, 3);
INSERT INTO `sys_user_authority` VALUES (1725, 3);
INSERT INTO `sys_user_authority` VALUES (1726, 3);
INSERT INTO `sys_user_authority` VALUES (1727, 3);
INSERT INTO `sys_user_authority` VALUES (1728, 3);
INSERT INTO `sys_user_authority` VALUES (1729, 3);
INSERT INTO `sys_user_authority` VALUES (1730, 3);
INSERT INTO `sys_user_authority` VALUES (1731, 3);
INSERT INTO `sys_user_authority` VALUES (1732, 3);
INSERT INTO `sys_user_authority` VALUES (1733, 3);
INSERT INTO `sys_user_authority` VALUES (1734, 3);
INSERT INTO `sys_user_authority` VALUES (1735, 3);
INSERT INTO `sys_user_authority` VALUES (1736, 3);
INSERT INTO `sys_user_authority` VALUES (1737, 3);
INSERT INTO `sys_user_authority` VALUES (1738, 3);
INSERT INTO `sys_user_authority` VALUES (1739, 3);
INSERT INTO `sys_user_authority` VALUES (1740, 3);
INSERT INTO `sys_user_authority` VALUES (1741, 3);
INSERT INTO `sys_user_authority` VALUES (1742, 3);
INSERT INTO `sys_user_authority` VALUES (1743, 3);
INSERT INTO `sys_user_authority` VALUES (1744, 3);
INSERT INTO `sys_user_authority` VALUES (1745, 3);
INSERT INTO `sys_user_authority` VALUES (1746, 3);
INSERT INTO `sys_user_authority` VALUES (1747, 3);
INSERT INTO `sys_user_authority` VALUES (1748, 3);
INSERT INTO `sys_user_authority` VALUES (1749, 3);
INSERT INTO `sys_user_authority` VALUES (1750, 3);
INSERT INTO `sys_user_authority` VALUES (1751, 3);
INSERT INTO `sys_user_authority` VALUES (1752, 3);
INSERT INTO `sys_user_authority` VALUES (1753, 3);
INSERT INTO `sys_user_authority` VALUES (1754, 3);
INSERT INTO `sys_user_authority` VALUES (1755, 3);
INSERT INTO `sys_user_authority` VALUES (1756, 3);
INSERT INTO `sys_user_authority` VALUES (1757, 3);
INSERT INTO `sys_user_authority` VALUES (1758, 3);
INSERT INTO `sys_user_authority` VALUES (1759, 3);
INSERT INTO `sys_user_authority` VALUES (1760, 3);
INSERT INTO `sys_user_authority` VALUES (1761, 3);
INSERT INTO `sys_user_authority` VALUES (1762, 3);
INSERT INTO `sys_user_authority` VALUES (1763, 3);
INSERT INTO `sys_user_authority` VALUES (1764, 3);
INSERT INTO `sys_user_authority` VALUES (1765, 3);
INSERT INTO `sys_user_authority` VALUES (1766, 3);
INSERT INTO `sys_user_authority` VALUES (1767, 3);
INSERT INTO `sys_user_authority` VALUES (1768, 3);
INSERT INTO `sys_user_authority` VALUES (1769, 3);
INSERT INTO `sys_user_authority` VALUES (1770, 3);
INSERT INTO `sys_user_authority` VALUES (1771, 3);
INSERT INTO `sys_user_authority` VALUES (1772, 3);
INSERT INTO `sys_user_authority` VALUES (1773, 3);
INSERT INTO `sys_user_authority` VALUES (1774, 3);
INSERT INTO `sys_user_authority` VALUES (1775, 3);
INSERT INTO `sys_user_authority` VALUES (1776, 3);
INSERT INTO `sys_user_authority` VALUES (1777, 3);
INSERT INTO `sys_user_authority` VALUES (1778, 3);
INSERT INTO `sys_user_authority` VALUES (1779, 3);
INSERT INTO `sys_user_authority` VALUES (1780, 3);
INSERT INTO `sys_user_authority` VALUES (1781, 3);
INSERT INTO `sys_user_authority` VALUES (1782, 3);
INSERT INTO `sys_user_authority` VALUES (1783, 3);
INSERT INTO `sys_user_authority` VALUES (1784, 3);
INSERT INTO `sys_user_authority` VALUES (1785, 3);
INSERT INTO `sys_user_authority` VALUES (1786, 3);
INSERT INTO `sys_user_authority` VALUES (1787, 3);
INSERT INTO `sys_user_authority` VALUES (1788, 3);
INSERT INTO `sys_user_authority` VALUES (1789, 3);
INSERT INTO `sys_user_authority` VALUES (1790, 3);
INSERT INTO `sys_user_authority` VALUES (1791, 3);
INSERT INTO `sys_user_authority` VALUES (1792, 3);
INSERT INTO `sys_user_authority` VALUES (1793, 3);
INSERT INTO `sys_user_authority` VALUES (1794, 3);
INSERT INTO `sys_user_authority` VALUES (1795, 3);
INSERT INTO `sys_user_authority` VALUES (1796, 3);
INSERT INTO `sys_user_authority` VALUES (1797, 3);
INSERT INTO `sys_user_authority` VALUES (1798, 3);
INSERT INTO `sys_user_authority` VALUES (1799, 3);
INSERT INTO `sys_user_authority` VALUES (1800, 3);
INSERT INTO `sys_user_authority` VALUES (1801, 3);
INSERT INTO `sys_user_authority` VALUES (1802, 3);
INSERT INTO `sys_user_authority` VALUES (1803, 3);
INSERT INTO `sys_user_authority` VALUES (1804, 3);
INSERT INTO `sys_user_authority` VALUES (1805, 3);
INSERT INTO `sys_user_authority` VALUES (1806, 3);
INSERT INTO `sys_user_authority` VALUES (1807, 3);
INSERT INTO `sys_user_authority` VALUES (1808, 3);
INSERT INTO `sys_user_authority` VALUES (1809, 3);
INSERT INTO `sys_user_authority` VALUES (1810, 3);
INSERT INTO `sys_user_authority` VALUES (1811, 3);
INSERT INTO `sys_user_authority` VALUES (1812, 3);
INSERT INTO `sys_user_authority` VALUES (1813, 3);
INSERT INTO `sys_user_authority` VALUES (1814, 3);
INSERT INTO `sys_user_authority` VALUES (1815, 3);
INSERT INTO `sys_user_authority` VALUES (1816, 3);
INSERT INTO `sys_user_authority` VALUES (1817, 3);
INSERT INTO `sys_user_authority` VALUES (1818, 3);
INSERT INTO `sys_user_authority` VALUES (1819, 3);
INSERT INTO `sys_user_authority` VALUES (1820, 3);
INSERT INTO `sys_user_authority` VALUES (1821, 3);
INSERT INTO `sys_user_authority` VALUES (1822, 3);
INSERT INTO `sys_user_authority` VALUES (1823, 3);
INSERT INTO `sys_user_authority` VALUES (1824, 3);
INSERT INTO `sys_user_authority` VALUES (1825, 3);
INSERT INTO `sys_user_authority` VALUES (1826, 3);
INSERT INTO `sys_user_authority` VALUES (1827, 3);
INSERT INTO `sys_user_authority` VALUES (1828, 3);
INSERT INTO `sys_user_authority` VALUES (1829, 3);
INSERT INTO `sys_user_authority` VALUES (1830, 3);
INSERT INTO `sys_user_authority` VALUES (1831, 3);
INSERT INTO `sys_user_authority` VALUES (1832, 3);
INSERT INTO `sys_user_authority` VALUES (1833, 3);
INSERT INTO `sys_user_authority` VALUES (1834, 3);
INSERT INTO `sys_user_authority` VALUES (1835, 3);
INSERT INTO `sys_user_authority` VALUES (1836, 3);
INSERT INTO `sys_user_authority` VALUES (1837, 3);
INSERT INTO `sys_user_authority` VALUES (1838, 3);
INSERT INTO `sys_user_authority` VALUES (1839, 3);
INSERT INTO `sys_user_authority` VALUES (1840, 3);
INSERT INTO `sys_user_authority` VALUES (1841, 3);
INSERT INTO `sys_user_authority` VALUES (1842, 3);
INSERT INTO `sys_user_authority` VALUES (1843, 3);
INSERT INTO `sys_user_authority` VALUES (1844, 3);
INSERT INTO `sys_user_authority` VALUES (1845, 3);
INSERT INTO `sys_user_authority` VALUES (1846, 3);
INSERT INTO `sys_user_authority` VALUES (1847, 3);
INSERT INTO `sys_user_authority` VALUES (1848, 3);
INSERT INTO `sys_user_authority` VALUES (1849, 3);
INSERT INTO `sys_user_authority` VALUES (1850, 3);
INSERT INTO `sys_user_authority` VALUES (1851, 3);
INSERT INTO `sys_user_authority` VALUES (1852, 3);
INSERT INTO `sys_user_authority` VALUES (1853, 3);
INSERT INTO `sys_user_authority` VALUES (1854, 3);
INSERT INTO `sys_user_authority` VALUES (1855, 3);
INSERT INTO `sys_user_authority` VALUES (1856, 3);
INSERT INTO `sys_user_authority` VALUES (1857, 3);
INSERT INTO `sys_user_authority` VALUES (1858, 3);
INSERT INTO `sys_user_authority` VALUES (1859, 3);
INSERT INTO `sys_user_authority` VALUES (1860, 3);
INSERT INTO `sys_user_authority` VALUES (1861, 3);
INSERT INTO `sys_user_authority` VALUES (1862, 3);
INSERT INTO `sys_user_authority` VALUES (1863, 3);
INSERT INTO `sys_user_authority` VALUES (1864, 3);
INSERT INTO `sys_user_authority` VALUES (1865, 3);
INSERT INTO `sys_user_authority` VALUES (1866, 3);
INSERT INTO `sys_user_authority` VALUES (1867, 3);
INSERT INTO `sys_user_authority` VALUES (1868, 3);
INSERT INTO `sys_user_authority` VALUES (1869, 3);
INSERT INTO `sys_user_authority` VALUES (1870, 3);
INSERT INTO `sys_user_authority` VALUES (1871, 3);
INSERT INTO `sys_user_authority` VALUES (1872, 3);
INSERT INTO `sys_user_authority` VALUES (1873, 3);
INSERT INTO `sys_user_authority` VALUES (1874, 3);
INSERT INTO `sys_user_authority` VALUES (1875, 3);
INSERT INTO `sys_user_authority` VALUES (1876, 3);
INSERT INTO `sys_user_authority` VALUES (1877, 3);
INSERT INTO `sys_user_authority` VALUES (1878, 3);
INSERT INTO `sys_user_authority` VALUES (1879, 3);
INSERT INTO `sys_user_authority` VALUES (1880, 3);
INSERT INTO `sys_user_authority` VALUES (1881, 3);
INSERT INTO `sys_user_authority` VALUES (1882, 3);
INSERT INTO `sys_user_authority` VALUES (1883, 3);
INSERT INTO `sys_user_authority` VALUES (1884, 3);
INSERT INTO `sys_user_authority` VALUES (1885, 3);
INSERT INTO `sys_user_authority` VALUES (1886, 3);
INSERT INTO `sys_user_authority` VALUES (1887, 3);
INSERT INTO `sys_user_authority` VALUES (1888, 3);
INSERT INTO `sys_user_authority` VALUES (1889, 3);
INSERT INTO `sys_user_authority` VALUES (1890, 3);
INSERT INTO `sys_user_authority` VALUES (1891, 3);
INSERT INTO `sys_user_authority` VALUES (1892, 3);
INSERT INTO `sys_user_authority` VALUES (1893, 3);
INSERT INTO `sys_user_authority` VALUES (1894, 3);
INSERT INTO `sys_user_authority` VALUES (1895, 3);
INSERT INTO `sys_user_authority` VALUES (1896, 3);
INSERT INTO `sys_user_authority` VALUES (1897, 3);
INSERT INTO `sys_user_authority` VALUES (1898, 3);
INSERT INTO `sys_user_authority` VALUES (1899, 3);
INSERT INTO `sys_user_authority` VALUES (1900, 3);
INSERT INTO `sys_user_authority` VALUES (1901, 3);
INSERT INTO `sys_user_authority` VALUES (1902, 3);
INSERT INTO `sys_user_authority` VALUES (1903, 3);
INSERT INTO `sys_user_authority` VALUES (1904, 3);
INSERT INTO `sys_user_authority` VALUES (1905, 3);
INSERT INTO `sys_user_authority` VALUES (1906, 3);
INSERT INTO `sys_user_authority` VALUES (1907, 3);
INSERT INTO `sys_user_authority` VALUES (1908, 3);
INSERT INTO `sys_user_authority` VALUES (1909, 3);
INSERT INTO `sys_user_authority` VALUES (1910, 3);
INSERT INTO `sys_user_authority` VALUES (1911, 3);
INSERT INTO `sys_user_authority` VALUES (1912, 3);
INSERT INTO `sys_user_authority` VALUES (1913, 3);
INSERT INTO `sys_user_authority` VALUES (1914, 3);
INSERT INTO `sys_user_authority` VALUES (1915, 3);
INSERT INTO `sys_user_authority` VALUES (1916, 3);
INSERT INTO `sys_user_authority` VALUES (1917, 3);
INSERT INTO `sys_user_authority` VALUES (1918, 3);
INSERT INTO `sys_user_authority` VALUES (1919, 3);
INSERT INTO `sys_user_authority` VALUES (1920, 3);
INSERT INTO `sys_user_authority` VALUES (1921, 3);
INSERT INTO `sys_user_authority` VALUES (1922, 3);
INSERT INTO `sys_user_authority` VALUES (1923, 3);
INSERT INTO `sys_user_authority` VALUES (1924, 3);
INSERT INTO `sys_user_authority` VALUES (1925, 3);
INSERT INTO `sys_user_authority` VALUES (1926, 3);
INSERT INTO `sys_user_authority` VALUES (1927, 3);
INSERT INTO `sys_user_authority` VALUES (1928, 3);
INSERT INTO `sys_user_authority` VALUES (1929, 3);
INSERT INTO `sys_user_authority` VALUES (1930, 3);
INSERT INTO `sys_user_authority` VALUES (1931, 3);
INSERT INTO `sys_user_authority` VALUES (1932, 3);
INSERT INTO `sys_user_authority` VALUES (1933, 3);
INSERT INTO `sys_user_authority` VALUES (1934, 3);
INSERT INTO `sys_user_authority` VALUES (1935, 3);
INSERT INTO `sys_user_authority` VALUES (1936, 3);
INSERT INTO `sys_user_authority` VALUES (1937, 3);
INSERT INTO `sys_user_authority` VALUES (1938, 3);
INSERT INTO `sys_user_authority` VALUES (1939, 3);
INSERT INTO `sys_user_authority` VALUES (1940, 3);
INSERT INTO `sys_user_authority` VALUES (1941, 3);
INSERT INTO `sys_user_authority` VALUES (1942, 3);
INSERT INTO `sys_user_authority` VALUES (1943, 3);
INSERT INTO `sys_user_authority` VALUES (1944, 3);
INSERT INTO `sys_user_authority` VALUES (1945, 3);
INSERT INTO `sys_user_authority` VALUES (1946, 3);
INSERT INTO `sys_user_authority` VALUES (1947, 3);
INSERT INTO `sys_user_authority` VALUES (1948, 3);
INSERT INTO `sys_user_authority` VALUES (1949, 3);
INSERT INTO `sys_user_authority` VALUES (1950, 3);
INSERT INTO `sys_user_authority` VALUES (1951, 3);
INSERT INTO `sys_user_authority` VALUES (1952, 3);
INSERT INTO `sys_user_authority` VALUES (1953, 3);
INSERT INTO `sys_user_authority` VALUES (1954, 3);
INSERT INTO `sys_user_authority` VALUES (1955, 3);
INSERT INTO `sys_user_authority` VALUES (1956, 3);
INSERT INTO `sys_user_authority` VALUES (1957, 3);
INSERT INTO `sys_user_authority` VALUES (1958, 3);
INSERT INTO `sys_user_authority` VALUES (1959, 3);
INSERT INTO `sys_user_authority` VALUES (1960, 3);
INSERT INTO `sys_user_authority` VALUES (1961, 3);
INSERT INTO `sys_user_authority` VALUES (1962, 3);
INSERT INTO `sys_user_authority` VALUES (1963, 3);
INSERT INTO `sys_user_authority` VALUES (1964, 3);
INSERT INTO `sys_user_authority` VALUES (1965, 3);
INSERT INTO `sys_user_authority` VALUES (1966, 3);
INSERT INTO `sys_user_authority` VALUES (1967, 3);
INSERT INTO `sys_user_authority` VALUES (1968, 3);
INSERT INTO `sys_user_authority` VALUES (1969, 3);
INSERT INTO `sys_user_authority` VALUES (1970, 3);
INSERT INTO `sys_user_authority` VALUES (1971, 3);
INSERT INTO `sys_user_authority` VALUES (1972, 3);
INSERT INTO `sys_user_authority` VALUES (1973, 3);
INSERT INTO `sys_user_authority` VALUES (1974, 3);
INSERT INTO `sys_user_authority` VALUES (1975, 3);
INSERT INTO `sys_user_authority` VALUES (1976, 3);
INSERT INTO `sys_user_authority` VALUES (1977, 3);
INSERT INTO `sys_user_authority` VALUES (1978, 3);
INSERT INTO `sys_user_authority` VALUES (1979, 3);
INSERT INTO `sys_user_authority` VALUES (1980, 3);
INSERT INTO `sys_user_authority` VALUES (1981, 3);
INSERT INTO `sys_user_authority` VALUES (1982, 3);
INSERT INTO `sys_user_authority` VALUES (1983, 3);
INSERT INTO `sys_user_authority` VALUES (1984, 3);
INSERT INTO `sys_user_authority` VALUES (1985, 3);
INSERT INTO `sys_user_authority` VALUES (1986, 3);
INSERT INTO `sys_user_authority` VALUES (1987, 3);
INSERT INTO `sys_user_authority` VALUES (1988, 3);
INSERT INTO `sys_user_authority` VALUES (1989, 3);
INSERT INTO `sys_user_authority` VALUES (1990, 3);
INSERT INTO `sys_user_authority` VALUES (1991, 3);
INSERT INTO `sys_user_authority` VALUES (1992, 3);
INSERT INTO `sys_user_authority` VALUES (1993, 3);
INSERT INTO `sys_user_authority` VALUES (1994, 3);
INSERT INTO `sys_user_authority` VALUES (1995, 3);
INSERT INTO `sys_user_authority` VALUES (1996, 3);
INSERT INTO `sys_user_authority` VALUES (1997, 3);
INSERT INTO `sys_user_authority` VALUES (1998, 3);
INSERT INTO `sys_user_authority` VALUES (1999, 3);
INSERT INTO `sys_user_authority` VALUES (2000, 3);
INSERT INTO `sys_user_authority` VALUES (2001, 3);
INSERT INTO `sys_user_authority` VALUES (2002, 3);
INSERT INTO `sys_user_authority` VALUES (2003, 3);
INSERT INTO `sys_user_authority` VALUES (2004, 3);
INSERT INTO `sys_user_authority` VALUES (2005, 3);
INSERT INTO `sys_user_authority` VALUES (2006, 3);
INSERT INTO `sys_user_authority` VALUES (2007, 3);
INSERT INTO `sys_user_authority` VALUES (2008, 3);
INSERT INTO `sys_user_authority` VALUES (2009, 3);
INSERT INTO `sys_user_authority` VALUES (2010, 3);
INSERT INTO `sys_user_authority` VALUES (2011, 3);
INSERT INTO `sys_user_authority` VALUES (2012, 3);
INSERT INTO `sys_user_authority` VALUES (2013, 3);
INSERT INTO `sys_user_authority` VALUES (2014, 3);
INSERT INTO `sys_user_authority` VALUES (2015, 3);
INSERT INTO `sys_user_authority` VALUES (2016, 3);
INSERT INTO `sys_user_authority` VALUES (2017, 3);
INSERT INTO `sys_user_authority` VALUES (2018, 3);
INSERT INTO `sys_user_authority` VALUES (2019, 3);
INSERT INTO `sys_user_authority` VALUES (2020, 3);
INSERT INTO `sys_user_authority` VALUES (2021, 3);
INSERT INTO `sys_user_authority` VALUES (2022, 3);
INSERT INTO `sys_user_authority` VALUES (2023, 3);
INSERT INTO `sys_user_authority` VALUES (2024, 3);
INSERT INTO `sys_user_authority` VALUES (2025, 3);
INSERT INTO `sys_user_authority` VALUES (2026, 3);
INSERT INTO `sys_user_authority` VALUES (2027, 3);
INSERT INTO `sys_user_authority` VALUES (2028, 3);
INSERT INTO `sys_user_authority` VALUES (2029, 3);
INSERT INTO `sys_user_authority` VALUES (2030, 3);
INSERT INTO `sys_user_authority` VALUES (2031, 3);
INSERT INTO `sys_user_authority` VALUES (2032, 3);
INSERT INTO `sys_user_authority` VALUES (2033, 3);
INSERT INTO `sys_user_authority` VALUES (2034, 3);
INSERT INTO `sys_user_authority` VALUES (2035, 3);
INSERT INTO `sys_user_authority` VALUES (2036, 3);
INSERT INTO `sys_user_authority` VALUES (2037, 3);
INSERT INTO `sys_user_authority` VALUES (2038, 3);
INSERT INTO `sys_user_authority` VALUES (2039, 3);
INSERT INTO `sys_user_authority` VALUES (2040, 3);
INSERT INTO `sys_user_authority` VALUES (2041, 3);
INSERT INTO `sys_user_authority` VALUES (2042, 3);
INSERT INTO `sys_user_authority` VALUES (2043, 3);
INSERT INTO `sys_user_authority` VALUES (2044, 3);
INSERT INTO `sys_user_authority` VALUES (2045, 3);
INSERT INTO `sys_user_authority` VALUES (2046, 3);
INSERT INTO `sys_user_authority` VALUES (2047, 3);
INSERT INTO `sys_user_authority` VALUES (2048, 3);
INSERT INTO `sys_user_authority` VALUES (2049, 3);
INSERT INTO `sys_user_authority` VALUES (2050, 3);
INSERT INTO `sys_user_authority` VALUES (2051, 3);
INSERT INTO `sys_user_authority` VALUES (2052, 3);
INSERT INTO `sys_user_authority` VALUES (2053, 3);
INSERT INTO `sys_user_authority` VALUES (2054, 3);
INSERT INTO `sys_user_authority` VALUES (2055, 3);
INSERT INTO `sys_user_authority` VALUES (2056, 3);
INSERT INTO `sys_user_authority` VALUES (2057, 3);
INSERT INTO `sys_user_authority` VALUES (2058, 3);
INSERT INTO `sys_user_authority` VALUES (2059, 3);
INSERT INTO `sys_user_authority` VALUES (2060, 3);
INSERT INTO `sys_user_authority` VALUES (2061, 3);
INSERT INTO `sys_user_authority` VALUES (2062, 3);
INSERT INTO `sys_user_authority` VALUES (2063, 3);
INSERT INTO `sys_user_authority` VALUES (2064, 3);
INSERT INTO `sys_user_authority` VALUES (2065, 3);
INSERT INTO `sys_user_authority` VALUES (2066, 3);
INSERT INTO `sys_user_authority` VALUES (2067, 3);
INSERT INTO `sys_user_authority` VALUES (2068, 3);
INSERT INTO `sys_user_authority` VALUES (2069, 3);
INSERT INTO `sys_user_authority` VALUES (2070, 3);
INSERT INTO `sys_user_authority` VALUES (2071, 3);
INSERT INTO `sys_user_authority` VALUES (2072, 3);
INSERT INTO `sys_user_authority` VALUES (2073, 3);
INSERT INTO `sys_user_authority` VALUES (2074, 3);
INSERT INTO `sys_user_authority` VALUES (2075, 3);
INSERT INTO `sys_user_authority` VALUES (2076, 3);
INSERT INTO `sys_user_authority` VALUES (2077, 3);
INSERT INTO `sys_user_authority` VALUES (2078, 3);
INSERT INTO `sys_user_authority` VALUES (2079, 3);
INSERT INTO `sys_user_authority` VALUES (2080, 3);
INSERT INTO `sys_user_authority` VALUES (2081, 3);
INSERT INTO `sys_user_authority` VALUES (2082, 3);
INSERT INTO `sys_user_authority` VALUES (2083, 3);
INSERT INTO `sys_user_authority` VALUES (2084, 3);
INSERT INTO `sys_user_authority` VALUES (2085, 3);
INSERT INTO `sys_user_authority` VALUES (2086, 3);
INSERT INTO `sys_user_authority` VALUES (2087, 3);
INSERT INTO `sys_user_authority` VALUES (2088, 3);
INSERT INTO `sys_user_authority` VALUES (2089, 3);
INSERT INTO `sys_user_authority` VALUES (2090, 3);
INSERT INTO `sys_user_authority` VALUES (2091, 3);
INSERT INTO `sys_user_authority` VALUES (2092, 3);
INSERT INTO `sys_user_authority` VALUES (2093, 3);
INSERT INTO `sys_user_authority` VALUES (2094, 3);
INSERT INTO `sys_user_authority` VALUES (2095, 3);
INSERT INTO `sys_user_authority` VALUES (2096, 3);
INSERT INTO `sys_user_authority` VALUES (2097, 3);
INSERT INTO `sys_user_authority` VALUES (2098, 3);
INSERT INTO `sys_user_authority` VALUES (2099, 3);
INSERT INTO `sys_user_authority` VALUES (2100, 3);
INSERT INTO `sys_user_authority` VALUES (2101, 3);
INSERT INTO `sys_user_authority` VALUES (2102, 3);
INSERT INTO `sys_user_authority` VALUES (2103, 3);
INSERT INTO `sys_user_authority` VALUES (2104, 3);
INSERT INTO `sys_user_authority` VALUES (2105, 3);
INSERT INTO `sys_user_authority` VALUES (2106, 3);
INSERT INTO `sys_user_authority` VALUES (2107, 3);
INSERT INTO `sys_user_authority` VALUES (2108, 3);
INSERT INTO `sys_user_authority` VALUES (2109, 3);
INSERT INTO `sys_user_authority` VALUES (2110, 3);
INSERT INTO `sys_user_authority` VALUES (2111, 3);
INSERT INTO `sys_user_authority` VALUES (2112, 3);
INSERT INTO `sys_user_authority` VALUES (2113, 3);
INSERT INTO `sys_user_authority` VALUES (2114, 3);
INSERT INTO `sys_user_authority` VALUES (2115, 3);
INSERT INTO `sys_user_authority` VALUES (2116, 3);
INSERT INTO `sys_user_authority` VALUES (2117, 3);
INSERT INTO `sys_user_authority` VALUES (2118, 3);
INSERT INTO `sys_user_authority` VALUES (2119, 3);
INSERT INTO `sys_user_authority` VALUES (2120, 3);
INSERT INTO `sys_user_authority` VALUES (2121, 3);
INSERT INTO `sys_user_authority` VALUES (2122, 3);
INSERT INTO `sys_user_authority` VALUES (2123, 3);
INSERT INTO `sys_user_authority` VALUES (2124, 3);
INSERT INTO `sys_user_authority` VALUES (2125, 3);
INSERT INTO `sys_user_authority` VALUES (2126, 3);
INSERT INTO `sys_user_authority` VALUES (2127, 3);
INSERT INTO `sys_user_authority` VALUES (2128, 3);
INSERT INTO `sys_user_authority` VALUES (2129, 3);
INSERT INTO `sys_user_authority` VALUES (2130, 3);
INSERT INTO `sys_user_authority` VALUES (2131, 3);
INSERT INTO `sys_user_authority` VALUES (2132, 3);
INSERT INTO `sys_user_authority` VALUES (2133, 3);
INSERT INTO `sys_user_authority` VALUES (2134, 3);
INSERT INTO `sys_user_authority` VALUES (2135, 3);
INSERT INTO `sys_user_authority` VALUES (2136, 3);
INSERT INTO `sys_user_authority` VALUES (2137, 3);
INSERT INTO `sys_user_authority` VALUES (2138, 3);
INSERT INTO `sys_user_authority` VALUES (2139, 3);
INSERT INTO `sys_user_authority` VALUES (2140, 3);
INSERT INTO `sys_user_authority` VALUES (2141, 3);
INSERT INTO `sys_user_authority` VALUES (2142, 3);
INSERT INTO `sys_user_authority` VALUES (2143, 3);
INSERT INTO `sys_user_authority` VALUES (2144, 3);
INSERT INTO `sys_user_authority` VALUES (2145, 3);
INSERT INTO `sys_user_authority` VALUES (2146, 3);
INSERT INTO `sys_user_authority` VALUES (2147, 3);
INSERT INTO `sys_user_authority` VALUES (2148, 3);
INSERT INTO `sys_user_authority` VALUES (2149, 3);
INSERT INTO `sys_user_authority` VALUES (2150, 3);
INSERT INTO `sys_user_authority` VALUES (2151, 3);
INSERT INTO `sys_user_authority` VALUES (2152, 3);
INSERT INTO `sys_user_authority` VALUES (2153, 3);
INSERT INTO `sys_user_authority` VALUES (2154, 3);
INSERT INTO `sys_user_authority` VALUES (2155, 3);
INSERT INTO `sys_user_authority` VALUES (2156, 3);
INSERT INTO `sys_user_authority` VALUES (2157, 3);
INSERT INTO `sys_user_authority` VALUES (2158, 3);
INSERT INTO `sys_user_authority` VALUES (2159, 3);
INSERT INTO `sys_user_authority` VALUES (2160, 3);
INSERT INTO `sys_user_authority` VALUES (2161, 3);
INSERT INTO `sys_user_authority` VALUES (2162, 3);
INSERT INTO `sys_user_authority` VALUES (2163, 3);
INSERT INTO `sys_user_authority` VALUES (2164, 3);
INSERT INTO `sys_user_authority` VALUES (2165, 3);
INSERT INTO `sys_user_authority` VALUES (2166, 3);
INSERT INTO `sys_user_authority` VALUES (2167, 3);
INSERT INTO `sys_user_authority` VALUES (2168, 3);
INSERT INTO `sys_user_authority` VALUES (2169, 3);
INSERT INTO `sys_user_authority` VALUES (2170, 3);
INSERT INTO `sys_user_authority` VALUES (2171, 3);
INSERT INTO `sys_user_authority` VALUES (2172, 3);
INSERT INTO `sys_user_authority` VALUES (2173, 3);
INSERT INTO `sys_user_authority` VALUES (2174, 3);
INSERT INTO `sys_user_authority` VALUES (2175, 3);
INSERT INTO `sys_user_authority` VALUES (2176, 3);
INSERT INTO `sys_user_authority` VALUES (2177, 3);
INSERT INTO `sys_user_authority` VALUES (2178, 3);
INSERT INTO `sys_user_authority` VALUES (2179, 3);
INSERT INTO `sys_user_authority` VALUES (2180, 3);
INSERT INTO `sys_user_authority` VALUES (2181, 3);
INSERT INTO `sys_user_authority` VALUES (2182, 3);
INSERT INTO `sys_user_authority` VALUES (2183, 3);
INSERT INTO `sys_user_authority` VALUES (2184, 3);
INSERT INTO `sys_user_authority` VALUES (2185, 3);
INSERT INTO `sys_user_authority` VALUES (2186, 3);
INSERT INTO `sys_user_authority` VALUES (2187, 3);
INSERT INTO `sys_user_authority` VALUES (2188, 3);
INSERT INTO `sys_user_authority` VALUES (2189, 3);
INSERT INTO `sys_user_authority` VALUES (2190, 3);
INSERT INTO `sys_user_authority` VALUES (2191, 3);
INSERT INTO `sys_user_authority` VALUES (2192, 3);
INSERT INTO `sys_user_authority` VALUES (2193, 3);
INSERT INTO `sys_user_authority` VALUES (2194, 3);
INSERT INTO `sys_user_authority` VALUES (2195, 3);
INSERT INTO `sys_user_authority` VALUES (2196, 3);
INSERT INTO `sys_user_authority` VALUES (2197, 3);
INSERT INTO `sys_user_authority` VALUES (2198, 3);
INSERT INTO `sys_user_authority` VALUES (2199, 3);
INSERT INTO `sys_user_authority` VALUES (2200, 3);
INSERT INTO `sys_user_authority` VALUES (2201, 3);
INSERT INTO `sys_user_authority` VALUES (2202, 3);
INSERT INTO `sys_user_authority` VALUES (2203, 3);
INSERT INTO `sys_user_authority` VALUES (2204, 3);
INSERT INTO `sys_user_authority` VALUES (2205, 3);
INSERT INTO `sys_user_authority` VALUES (2206, 3);
INSERT INTO `sys_user_authority` VALUES (2207, 3);
INSERT INTO `sys_user_authority` VALUES (2208, 3);
INSERT INTO `sys_user_authority` VALUES (2209, 3);
INSERT INTO `sys_user_authority` VALUES (2210, 3);
INSERT INTO `sys_user_authority` VALUES (2211, 3);
INSERT INTO `sys_user_authority` VALUES (2212, 3);
INSERT INTO `sys_user_authority` VALUES (2213, 3);
INSERT INTO `sys_user_authority` VALUES (2214, 3);
INSERT INTO `sys_user_authority` VALUES (2215, 3);
INSERT INTO `sys_user_authority` VALUES (2216, 3);
INSERT INTO `sys_user_authority` VALUES (2217, 3);
INSERT INTO `sys_user_authority` VALUES (2218, 3);
INSERT INTO `sys_user_authority` VALUES (2219, 3);
INSERT INTO `sys_user_authority` VALUES (2220, 3);
INSERT INTO `sys_user_authority` VALUES (2221, 3);
INSERT INTO `sys_user_authority` VALUES (2222, 3);
INSERT INTO `sys_user_authority` VALUES (2223, 3);
INSERT INTO `sys_user_authority` VALUES (2224, 3);
INSERT INTO `sys_user_authority` VALUES (2225, 3);
INSERT INTO `sys_user_authority` VALUES (2226, 3);
INSERT INTO `sys_user_authority` VALUES (2227, 3);
INSERT INTO `sys_user_authority` VALUES (2228, 3);
INSERT INTO `sys_user_authority` VALUES (2229, 3);
INSERT INTO `sys_user_authority` VALUES (2230, 3);
INSERT INTO `sys_user_authority` VALUES (2231, 3);
INSERT INTO `sys_user_authority` VALUES (2232, 3);
INSERT INTO `sys_user_authority` VALUES (2233, 3);
INSERT INTO `sys_user_authority` VALUES (2234, 3);
INSERT INTO `sys_user_authority` VALUES (2235, 3);
INSERT INTO `sys_user_authority` VALUES (2236, 3);
INSERT INTO `sys_user_authority` VALUES (2237, 3);
INSERT INTO `sys_user_authority` VALUES (2238, 3);
INSERT INTO `sys_user_authority` VALUES (2239, 3);
INSERT INTO `sys_user_authority` VALUES (2240, 3);
INSERT INTO `sys_user_authority` VALUES (2241, 3);
INSERT INTO `sys_user_authority` VALUES (2242, 3);
INSERT INTO `sys_user_authority` VALUES (2243, 3);
INSERT INTO `sys_user_authority` VALUES (2244, 3);
INSERT INTO `sys_user_authority` VALUES (2245, 3);
INSERT INTO `sys_user_authority` VALUES (2246, 3);
INSERT INTO `sys_user_authority` VALUES (2247, 3);
INSERT INTO `sys_user_authority` VALUES (2248, 3);
INSERT INTO `sys_user_authority` VALUES (2249, 3);
INSERT INTO `sys_user_authority` VALUES (2250, 3);
INSERT INTO `sys_user_authority` VALUES (2251, 3);
INSERT INTO `sys_user_authority` VALUES (2252, 3);
INSERT INTO `sys_user_authority` VALUES (2253, 3);
INSERT INTO `sys_user_authority` VALUES (2254, 3);
INSERT INTO `sys_user_authority` VALUES (2255, 3);
INSERT INTO `sys_user_authority` VALUES (2256, 3);
INSERT INTO `sys_user_authority` VALUES (2257, 3);
INSERT INTO `sys_user_authority` VALUES (2258, 3);
INSERT INTO `sys_user_authority` VALUES (2259, 3);
INSERT INTO `sys_user_authority` VALUES (2260, 3);
INSERT INTO `sys_user_authority` VALUES (2261, 3);
INSERT INTO `sys_user_authority` VALUES (2262, 3);
INSERT INTO `sys_user_authority` VALUES (2263, 3);
INSERT INTO `sys_user_authority` VALUES (2264, 3);
INSERT INTO `sys_user_authority` VALUES (2265, 3);
INSERT INTO `sys_user_authority` VALUES (2266, 3);
INSERT INTO `sys_user_authority` VALUES (2267, 3);
INSERT INTO `sys_user_authority` VALUES (2268, 3);
INSERT INTO `sys_user_authority` VALUES (2269, 3);
INSERT INTO `sys_user_authority` VALUES (2270, 3);
INSERT INTO `sys_user_authority` VALUES (2271, 3);
INSERT INTO `sys_user_authority` VALUES (2272, 3);
INSERT INTO `sys_user_authority` VALUES (2273, 3);
INSERT INTO `sys_user_authority` VALUES (2274, 3);
INSERT INTO `sys_user_authority` VALUES (2275, 3);
INSERT INTO `sys_user_authority` VALUES (2276, 3);
INSERT INTO `sys_user_authority` VALUES (2277, 3);
INSERT INTO `sys_user_authority` VALUES (2278, 3);
INSERT INTO `sys_user_authority` VALUES (2279, 3);
INSERT INTO `sys_user_authority` VALUES (2280, 3);
INSERT INTO `sys_user_authority` VALUES (2281, 3);
INSERT INTO `sys_user_authority` VALUES (2282, 3);
INSERT INTO `sys_user_authority` VALUES (2283, 3);
INSERT INTO `sys_user_authority` VALUES (2284, 3);
INSERT INTO `sys_user_authority` VALUES (2285, 3);
INSERT INTO `sys_user_authority` VALUES (2286, 3);
INSERT INTO `sys_user_authority` VALUES (2287, 3);
INSERT INTO `sys_user_authority` VALUES (2288, 3);
INSERT INTO `sys_user_authority` VALUES (2289, 3);
INSERT INTO `sys_user_authority` VALUES (2290, 3);
INSERT INTO `sys_user_authority` VALUES (2291, 3);
INSERT INTO `sys_user_authority` VALUES (2292, 3);
INSERT INTO `sys_user_authority` VALUES (2293, 3);
INSERT INTO `sys_user_authority` VALUES (2294, 3);
INSERT INTO `sys_user_authority` VALUES (2295, 3);
INSERT INTO `sys_user_authority` VALUES (2296, 3);
INSERT INTO `sys_user_authority` VALUES (2297, 3);
INSERT INTO `sys_user_authority` VALUES (2298, 3);
INSERT INTO `sys_user_authority` VALUES (2299, 3);
INSERT INTO `sys_user_authority` VALUES (2300, 3);
INSERT INTO `sys_user_authority` VALUES (2301, 3);
INSERT INTO `sys_user_authority` VALUES (2302, 3);
INSERT INTO `sys_user_authority` VALUES (2303, 3);
INSERT INTO `sys_user_authority` VALUES (2304, 3);
INSERT INTO `sys_user_authority` VALUES (2305, 3);
INSERT INTO `sys_user_authority` VALUES (2306, 3);
INSERT INTO `sys_user_authority` VALUES (2307, 3);
INSERT INTO `sys_user_authority` VALUES (2308, 3);
INSERT INTO `sys_user_authority` VALUES (2309, 3);
INSERT INTO `sys_user_authority` VALUES (2310, 3);
INSERT INTO `sys_user_authority` VALUES (2311, 3);
INSERT INTO `sys_user_authority` VALUES (2312, 3);
INSERT INTO `sys_user_authority` VALUES (2313, 3);
INSERT INTO `sys_user_authority` VALUES (2314, 3);
INSERT INTO `sys_user_authority` VALUES (2315, 3);
INSERT INTO `sys_user_authority` VALUES (2316, 3);
INSERT INTO `sys_user_authority` VALUES (2317, 3);
INSERT INTO `sys_user_authority` VALUES (2318, 3);
INSERT INTO `sys_user_authority` VALUES (2319, 3);
INSERT INTO `sys_user_authority` VALUES (2320, 3);
INSERT INTO `sys_user_authority` VALUES (2321, 3);
INSERT INTO `sys_user_authority` VALUES (2322, 3);
INSERT INTO `sys_user_authority` VALUES (2323, 3);
INSERT INTO `sys_user_authority` VALUES (2324, 3);
INSERT INTO `sys_user_authority` VALUES (2325, 3);
INSERT INTO `sys_user_authority` VALUES (2326, 3);
INSERT INTO `sys_user_authority` VALUES (2327, 3);
INSERT INTO `sys_user_authority` VALUES (2328, 3);
INSERT INTO `sys_user_authority` VALUES (2329, 3);
INSERT INTO `sys_user_authority` VALUES (2330, 3);
INSERT INTO `sys_user_authority` VALUES (2331, 3);
INSERT INTO `sys_user_authority` VALUES (2332, 3);
INSERT INTO `sys_user_authority` VALUES (2333, 3);
INSERT INTO `sys_user_authority` VALUES (2334, 3);
INSERT INTO `sys_user_authority` VALUES (2335, 3);
INSERT INTO `sys_user_authority` VALUES (2336, 3);
INSERT INTO `sys_user_authority` VALUES (2337, 3);
INSERT INTO `sys_user_authority` VALUES (2338, 3);
INSERT INTO `sys_user_authority` VALUES (2339, 3);
INSERT INTO `sys_user_authority` VALUES (2340, 3);
INSERT INTO `sys_user_authority` VALUES (2341, 3);
INSERT INTO `sys_user_authority` VALUES (2342, 3);
INSERT INTO `sys_user_authority` VALUES (2343, 3);
INSERT INTO `sys_user_authority` VALUES (2344, 3);
INSERT INTO `sys_user_authority` VALUES (2345, 3);
INSERT INTO `sys_user_authority` VALUES (2346, 3);
INSERT INTO `sys_user_authority` VALUES (2347, 3);
INSERT INTO `sys_user_authority` VALUES (2348, 3);
INSERT INTO `sys_user_authority` VALUES (2349, 3);
INSERT INTO `sys_user_authority` VALUES (2350, 3);
INSERT INTO `sys_user_authority` VALUES (2351, 3);
INSERT INTO `sys_user_authority` VALUES (2352, 3);
INSERT INTO `sys_user_authority` VALUES (2353, 3);
INSERT INTO `sys_user_authority` VALUES (2354, 3);
INSERT INTO `sys_user_authority` VALUES (2355, 3);
INSERT INTO `sys_user_authority` VALUES (2356, 3);
INSERT INTO `sys_user_authority` VALUES (2357, 3);
INSERT INTO `sys_user_authority` VALUES (2358, 3);
INSERT INTO `sys_user_authority` VALUES (2359, 3);
INSERT INTO `sys_user_authority` VALUES (2360, 3);
INSERT INTO `sys_user_authority` VALUES (2361, 3);
INSERT INTO `sys_user_authority` VALUES (2362, 3);
INSERT INTO `sys_user_authority` VALUES (2363, 3);
INSERT INTO `sys_user_authority` VALUES (2364, 3);
INSERT INTO `sys_user_authority` VALUES (2365, 3);
INSERT INTO `sys_user_authority` VALUES (2366, 3);
INSERT INTO `sys_user_authority` VALUES (2367, 3);
INSERT INTO `sys_user_authority` VALUES (2368, 3);
INSERT INTO `sys_user_authority` VALUES (2369, 3);
INSERT INTO `sys_user_authority` VALUES (2370, 3);
INSERT INTO `sys_user_authority` VALUES (2371, 3);
INSERT INTO `sys_user_authority` VALUES (2372, 3);
INSERT INTO `sys_user_authority` VALUES (2373, 3);
INSERT INTO `sys_user_authority` VALUES (2374, 3);
INSERT INTO `sys_user_authority` VALUES (2375, 3);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://oldwei.oss-cn-hangzhou.aliyuncs.com/pics/avatar.jpg' COMMENT '用户头像',
  `authority_id` bigint UNSIGNED NULL DEFAULT 3 COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  `origin_setting` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '配置',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_uuid`(`uuid` ASC) USING BTREE,
  INDEX `idx_sys_users_username`(`username` ASC) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2378 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2024-10-09 00:20:06.501', '2024-11-05 12:14:06.365', NULL, 'b3f36d1b-6a73-46a7-9b2a-191bcac41996', 'admin', '$2a$10$ARp9/rCo4Z1JdpUfYcHdo.oLAoxI1Na9vKEBf737zFz.OgvPg./ay', '系统管理员', 'https://oldwei.oss-cn-hangzhou.aliyuncs.com/pics/avatar.jpg', 1, '18788888888', '888888888@qq.com', 1, NULL);
INSERT INTO `sys_users` VALUES (2, '2024-10-09 00:20:06.501', '2024-11-05 12:14:09.707', NULL, '5e466dad-4cd5-4f3d-b6f5-4cb4f2f615e0', 'oldwei', '$2a$10$jWoxxvMfvgGOlAWTIa6pQ.Se.cxU5Kc.HX54sNSBvjNjHMHXcEvN6', '老卫同学', 'https://oldwei.oss-cn-hangzhou.aliyuncs.com/pics/avatar.jpg', 3, '18888888888', '666666666@qq.com', 1, NULL);
INSERT INTO `sys_users` VALUES (5, '2023-02-18 16:19:15.826', '2024-11-05 12:14:12.450', NULL, 'eef87926-b83e-4a77-b394-7f89791934c7', 'oooldwei', '$2a$10$da5BKWshwai6pfmn4syXnu0W3aKToSTvsW/JIAyc5etiu12dYGUSC', '老卫测试', 'https://oldwei.oss-cn-hangzhou.aliyuncs.com/pics/avatar.jpg', 3, '', 'oooldwei@oldwei.com', 1, NULL);

SET FOREIGN_KEY_CHECKS = 1;
