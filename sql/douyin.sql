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

 Date: 30/05/2022 17:57:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
  `favorite_count` int(0) NOT NULL COMMENT '喜爱数',
  `total_favorited` int(0) NOT NULL COMMENT '总赞数',
  PRIMARY KEY (`userId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'zhanglei', 'douyin', 0, 0, 0, 0);
INSERT INTO `user` VALUES (5, 'abc', '123456', 0, 0, 0, 0);
INSERT INTO `user` VALUES (6, 'a', '123456', 0, 0, 0, 0);

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
INSERT INTO `video` VALUES (1, 'zhanglei', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0);
INSERT INTO `video` VALUES (2, 'a', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0);
INSERT INTO `video` VALUES (3, 'abc', 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
