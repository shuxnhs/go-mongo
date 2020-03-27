/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80017
 Source Host           : localhost:3306
 Source Schema         : go-mongo

 Target Server Type    : MySQL
 Target Server Version : 80017
 File Encoding         : 65001

 Date: 13/03/2020 16:02:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `mongo_key` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '项目的app_key',
  `host` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'mongodb的host',
  `port` int(10) NOT NULL DEFAULT '27017' COMMENT '端口',
  `user` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` blob NOT NULL COMMENT '密码',
  `dbname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '数据库名',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `udx_ak` (`mongo_key`) USING BTREE COMMENT 'ak唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='保存每个项目的mongodb配置';

-- ----------------------------
-- Records of config
-- ----------------------------
BEGIN;
INSERT INTO `config` VALUES (11, 'DA1CFA0CB817A9B2F3A3FE15604F3752', '127.0.0.1', 27017, '', '', 'testMongo');
COMMIT;

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `project_name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '项目名称',
  `mongo_key` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '项目使用配置的key',
  `is_deleted` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0未删除,1已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of project
-- ----------------------------
BEGIN;
INSERT INTO `project` VALUES (2, '本地测试', 'DA1CFA0CB817A9B2F3A3FE15604F3752', 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
