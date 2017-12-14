/*
Navicat MySQL Data Transfer

Source Server         : Localhost MYSQL
Source Server Version : 50616
Source Host           : 127.0.0.1:3306
Source Database       : curd

Target Server Type    : MYSQL
Target Server Version : 50616
File Encoding         : 65001

Date: 2017-12-13 15:05:39
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for data
-- ----------------------------
DROP TABLE IF EXISTS `data`;
CREATE TABLE `data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `address` text,
  `photo` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of data
-- ----------------------------
INSERT INTO `data` VALUES ('1', 'Azis', 'khamdan.ngazis@gmail.com', '085747975295', 'jln raya serpong km 8, pakulonan, serpong utara,Tangerang Selatan', '001232.JPG');

-- ----------------------------
-- Table structure for department
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `departmentname` varchar(25) NOT NULL,
  `numberofemployees` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of department
-- ----------------------------
INSERT INTO `department` VALUES ('1', 'Production', '500');
