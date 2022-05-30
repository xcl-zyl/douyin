/*
 Navicat Premium Data Transfer

 Source Server         : xcl
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 31/05/2022 04:55:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for favorite_1
-- ----------------------------
DROP TABLE IF EXISTS `favorite_1`;
CREATE TABLE `favorite_1`  (
  `userName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite_1
-- ----------------------------
INSERT INTO `favorite_1` VALUES ('abc');

-- ----------------------------
-- Table structure for favorite_2
-- ----------------------------
DROP TABLE IF EXISTS `favorite_2`;
CREATE TABLE `favorite_2`  (
  `userName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite_2
-- ----------------------------

-- ----------------------------
-- Table structure for favorite_3
-- ----------------------------
DROP TABLE IF EXISTS `favorite_3`;
CREATE TABLE `favorite_3`  (
  `userName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite_3
-- ----------------------------

-- ----------------------------
-- Table structure for follow_abc
-- ----------------------------
DROP TABLE IF EXISTS `follow_abc`;
CREATE TABLE `follow_abc`  (
  `followName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of follow_abc
-- ----------------------------
INSERT INTO `follow_abc` VALUES ('zhanglei');
INSERT INTO `follow_abc` VALUES ('a');

-- ----------------------------
-- Table structure for follow_test
-- ----------------------------
DROP TABLE IF EXISTS `follow_test`;
CREATE TABLE `follow_test`  (
  `followName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of follow_test
-- ----------------------------

-- ----------------------------
-- Table structure for follower_a
-- ----------------------------
DROP TABLE IF EXISTS `follower_a`;
CREATE TABLE `follower_a`  (
  `followerName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of follower_a
-- ----------------------------
INSERT INTO `follower_a` VALUES ('abc');

-- ----------------------------
-- Table structure for follower_zhanglei
-- ----------------------------
DROP TABLE IF EXISTS `follower_zhanglei`;
CREATE TABLE `follower_zhanglei`  (
  `followerName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of follower_zhanglei
-- ----------------------------
INSERT INTO `follower_zhanglei` VALUES ('abc');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `userId` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户Id',
  `userName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户密码',
  `followCount` int(0) NOT NULL DEFAULT 0 COMMENT '关注数',
  `followerCount` int(0) NOT NULL COMMENT '粉丝数',
  `work_count` int(0) NOT NULL COMMENT '作品数',
  `favorite_count` int(0) NOT NULL COMMENT '喜爱数',
  `total_favorited` int(0) NOT NULL COMMENT '总赞数',
  PRIMARY KEY (`userId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'zhanglei', 'douyin', 0, 1, 1, 0, 1);
INSERT INTO `user` VALUES (5, 'abc', '123456', 2, 0, 2, 1, 0);
INSERT INTO `user` VALUES (6, 'a', '123456', 0, 1, 1, 0, 0);

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `videoId` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `author` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '发布作者',
  `playUrl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '播放链接',
  `coverImg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频封面链接',
  `favoriteCount` int(0) NOT NULL COMMENT '视频被点赞数量',
  `commentCount` int(0) NOT NULL COMMENT '视频评论条数',
  PRIMARY KEY (`videoId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, 'zhanglei', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 1, 0);
INSERT INTO `video` VALUES (2, 'a', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0);
INSERT INTO `video` VALUES (3, 'abc', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0);
INSERT INTO `video` VALUES (13, 'abc', 'http://172.22.7.45:8080/static/5_國產av-糖心vlog-jk制服黑絲內射在屁股裏.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
