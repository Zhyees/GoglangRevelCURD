/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 5.0.27-community-nt : Database - curd
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`curd` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `curd`;

/*Table structure for table `data` */

DROP TABLE IF EXISTS `data`;

CREATE TABLE `data` (
  `id` int(11) NOT NULL auto_increment,
  `name` varchar(50) default NULL,
  `email` varchar(50) default NULL,
  `phone` varchar(20) default NULL,
  `address` text,
  `photo` varchar(100) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `data` */

insert  into `data`(`id`,`name`,`email`,`phone`,`address`,`photo`) values 
(1,'Azis','khamdan.ngazis@gmail.com','085747975295','jln raya serpong km 8, pakulonan, serpong utara,Tangerang Selatan','001232.JPG');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
