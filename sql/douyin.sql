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

 Date: 29/05/2022 23:15:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `commentId` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `videoId` int(0) NULL DEFAULT NULL,
  `author` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `createDate` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`commentId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (2, 7, 'aaa', 'bbb', '05-01');
INSERT INTO `comments` VALUES (5, 9, 'abc', 'ccc', '05-29');
INSERT INTO `comments` VALUES (6, 8, 'abc', 'aaa', '05-29');
INSERT INTO `comments` VALUES (8, 8, 'a', '呵呵呵', '05-29');

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
INSERT INTO `favorite_2` VALUES ('a');
INSERT INTO `favorite_2` VALUES ('zhanglei');
INSERT INTO `favorite_2` VALUES ('abc');

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
-- Table structure for favorite_7
-- ----------------------------
DROP TABLE IF EXISTS `favorite_7`;
CREATE TABLE `favorite_7`  (
  `userName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite_7
-- ----------------------------
INSERT INTO `favorite_7` VALUES ('zhanglei');
INSERT INTO `favorite_7` VALUES ('abc');

-- ----------------------------
-- Table structure for favorite_8
-- ----------------------------
DROP TABLE IF EXISTS `favorite_8`;
CREATE TABLE `favorite_8`  (
  `userName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite_8
-- ----------------------------
INSERT INTO `favorite_8` VALUES ('a');
INSERT INTO `favorite_8` VALUES ('zhanglei');

-- ----------------------------
-- Table structure for favorite_9
-- ----------------------------
DROP TABLE IF EXISTS `favorite_9`;
CREATE TABLE `favorite_9`  (
  `userName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite_9
-- ----------------------------
INSERT INTO `favorite_9` VALUES ('abc');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `userId` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户Id',
  `userName` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户密码',
  PRIMARY KEY (`userId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'zhanglei', 'douyin');
INSERT INTO `user` VALUES (5, 'abc', '123456');
INSERT INTO `user` VALUES (6, 'a', '123456');

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
  `isFavorite` tinyint(1) NOT NULL COMMENT '视频是否被该用户点赞',
  PRIMARY KEY (`videoId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, 'zhanglei', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0, 0);
INSERT INTO `video` VALUES (2, 'a', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 3, 0, 0);
INSERT INTO `video` VALUES (3, 'abc', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0, 0);
INSERT INTO `video` VALUES (7, 'abc', 'http://172.22.18.151:8080/static/5_VIDEO_20220511_162019968.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 3, 0, 0);
INSERT INTO `video` VALUES (8, 'a', 'http://172.22.18.151:8080/static/6_VIDEO_20220511_162019968.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 2, 2, 0);
INSERT INTO `video` VALUES (9, 'abc', 'http://172.22.13.149:8080/static/5_VIDEO_20220511_162019968.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 2, 1, 0);

SET FOREIGN_KEY_CHECKS = 1;
