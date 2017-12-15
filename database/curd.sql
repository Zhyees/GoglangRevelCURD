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


CREATE TABLE `menu` (
  `id` int(11) NOT NULL auto_increment,
  `name` varchar(50) default NULL,
  `url` varchar(50) default NULL,
  `icon` varchar(20) default NULL,
  `description` varchar(50) default NULL,
  `modul_id` int(11) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `mini_app` (
  `id` int(11) NOT NULL auto_increment,
  `code` varchar(50) NOT NULL,
  `description` varchar(255) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL default '0',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `mini_app` (`id`, `code`, `description`, `is_deleted`) VALUES 
  (1,'HomeCtrl','Dashboard Module',0),
  (2,'EmployedCtrl','Employed Modul',0),
  (3,'DepartmentCtrl','Departement Modul',0);

CREATE TABLE `mini_app_map` (
  `id` int(11) NOT NULL auto_increment,
  `mini_app_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  PRIMARY KEY  (`id`),
  KEY `mini_app_id` (`mini_app_id`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `mini_app_map` (`id`, `mini_app_id`, `role_id`) VALUES 
  (1,1,1),
  (2,2,1),
  (3,1,2),
  (4,3,2);

CREATE TABLE `modul1` (
  `id` int(11) NOT NULL auto_increment,
  `name` varchar(20) default NULL,
  `url` varchar(50) default NULL,
  `icon` varchar(50) default NULL,
  `description` varchar(255) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL default '0',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `modul1` (`id`, `name`, `url`, `icon`, `description`, `is_deleted`) VALUES 
  (1,'Dashboard','/home/','fa fa-android','Menu Home',0),
  (2,'Employed','/employed/','fa fa-rocket','Menu Employed',0),
  (3,'Departement','/departtement/','fa fa-dashboard','Menu Departement',0),
  (4,'Sign Out','/adminctrl/logout/','glyphicon glyphicon-log-out','Sign Out',0);

CREATE TABLE `modul_role_mapping` (
  `id` int(11) NOT NULL auto_increment,
  `modul_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  PRIMARY KEY  (`id`),
  KEY `mini_app_id` (`modul_id`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `modul_role_mapping` (`id`, `modul_id`, `role_id`) VALUES 
  (1,1,1),
  (4,1,2),
  (5,2,1),
  (6,3,2),
  (7,4,1),
  (8,4,2);

CREATE TABLE `role` (
  `id` int(11) NOT NULL auto_increment,
  `name` varchar(100) NOT NULL,
  `description` varchar(255) NOT NULL,
  `parent_role_id` int(11) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL default '0',
  PRIMARY KEY  (`id`),
  KEY `parent_role_id` (`parent_role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `role` (`id`, `name`, `description`, `parent_role_id`, `is_deleted`) VALUES
  (1,'Head of Operation','Head of Operation',1,0),
  (2,'Operational Admin','Responsible admin for operational activity',1,0);

CREATE TABLE `superuser` (
  `id` int(11) NOT NULL auto_increment,
  `email` varchar(50) NOT NULL,
  `full_name` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `login_token` varchar(100) NOT NULL,
  `role_id` int(11) NOT NULL,
  `last_login` timestamp NOT NULL default CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL default '0',
  PRIMARY KEY  (`id`),
  KEY `role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `superuser` (`id`, `email`, `full_name`, `password`, `login_token`, `role_id`, `last_login`, `is_deleted`) VALUES 
  (1,'khamdan.ngazis@gmail.com','Azis','$2a$10$IVPBJEWDnNqT0DQiE7fs2OBcbZ2BeVNp7sBql2nc/WpsFaN91EFyi','FKmvfIsMcBLKKzEIFfCm',1,'0000-00-00 00:00:00',0),
  (2,'sardi@gmail.com','Sardi','$2a$10$IVPBJEWDnNqT0DQiE7fs2OBcbZ2BeVNp7sBql2nc/WpsFaN91EFyi','CaMPoqtYmvrfcNcHVnbr',2,'0000-00-00 00:00:00',0);






