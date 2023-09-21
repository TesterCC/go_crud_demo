/*
 Navicat Premium Data Transfer

 Source Server         : mysql5.7-dev-192.168.80.129
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : 192.168.80.129:3306
 Source Schema         : crud_list

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 21/09/2023 22:24:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for list
-- ----------------------------
DROP TABLE IF EXISTS `list`;
CREATE TABLE `list`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `state` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `phone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `email` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_list_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of list
-- ----------------------------
INSERT INTO `list` VALUES (1, '2023-09-11 17:08:16.750', '2023-09-11 17:08:16.750', NULL, 'Alan', 'online', '13111111111', 'alan@test.com', 'China, Guangzhou Province, Shenzhen, Nanshan');
INSERT INTO `list` VALUES (2, '2023-09-11 17:11:47.043', '2023-09-11 17:11:47.043', NULL, 'Bob', 'offline', '15222222222', 'bob@test.com', 'China, Zhejiang Province, Ningbo');
INSERT INTO `list` VALUES (3, '2023-09-11 17:13:01.652', '2023-09-21 17:04:45.354', NULL, 'Chris223', 'offline', '17333333333', 'christina222@test.com', 'China, Sichuan Province, Chengdu');
INSERT INTO `list` VALUES (4, '2023-09-15 18:07:55.166', '2023-09-15 18:07:55.166', '2023-09-15 18:08:23.537', 'Chris2', 'online', '17333333333', 'christina@test.com', 'China, Sichuan Province, Chengdu');
INSERT INTO `list` VALUES (5, '2023-09-21 17:53:47.163', '2023-09-21 17:53:47.163', NULL, 'David', 'online', '17933334444', 'David@test.com', 'USA Washington DC');
INSERT INTO `list` VALUES (6, '2023-09-21 17:57:20.712', '2023-09-21 17:57:20.712', NULL, 'David', 'online', '18551850082', 'David82@test.com', 'Singapore, Ubi Tech Park');

SET FOREIGN_KEY_CHECKS = 1;
