/*
SQLyog - Free MySQL GUI v5.11
Host - 5.0.41-community-nt : Database - goles
*********************************************************************
Server version : 5.0.41-community-nt
*/

SET NAMES utf8;

SET SQL_MODE='';

create database if not exists `goles`;

USE `goles`;

/*Table structure for table `content` */

DROP TABLE IF EXISTS `content`;

CREATE TABLE `content` (
  `Id` int(11) NOT NULL auto_increment,
  `subject` varchar(50) default NULL,
  `news` varchar(200) default NULL,
  `periode` varchar(6) default NULL,
  PRIMARY KEY  (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `content` */

insert into `content` (`Id`,`subject`,`news`,`periode`) values (1,'Belajar Golang','Go adalah sebuah bahasa pemrograman yang dibuat di google pada tahun 2009 oleh Robert Griesemer, Rob Pike, dan Ken Thompson','202003');
insert into `content` (`Id`,`subject`,`news`,`periode`) values (2,'Bluebird','Bluebird adalah perusahaan taxi terbesar di Indonesia','202003');
insert into `content` (`Id`,`subject`,`news`,`periode`) values (3,'Burung Biru','Berawal dari sebutir telur','202003');
insert into `content` (`Id`,`subject`,`news`,`periode`) values (4,'Bigbird Raih Penghargaan di Transportation Safety ','PT Big Bird Pusaka meraih penghargaan juara I kategori angkutan jalan dalam ajang Transportation Safety Management Award (TSMA) 2019 yang diselenggarakan Kementerian Perhubungan (Kemenhub).','202003');
insert into `content` (`Id`,`subject`,`news`,`periode`) values (5,'Test','Cobain add artikel','202003');
insert into `content` (`Id`,`subject`,`news`,`periode`) values (6,'Test','Cobain add artikel','202003');

/*Table structure for table `userinfo` */

DROP TABLE IF EXISTS `userinfo`;

CREATE TABLE `userinfo` (
  `username` varchar(64) NOT NULL,
  `password` varchar(10) NOT NULL,
  `department` varchar(64) NOT NULL,
  `isadmin` tinyint(1) NOT NULL,
  `created` date NOT NULL,
  PRIMARY KEY  (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `userinfo` */

insert into `userinfo` (`username`,`password`,`department`,`isadmin`,`created`) values ('admin','123','it',1,'2020-02-18');
insert into `userinfo` (`username`,`password`,`department`,`isadmin`,`created`) values ('budi','123','bd',0,'2019-01-01');
insert into `userinfo` (`username`,`password`,`department`,`isadmin`,`created`) values ('heru','1234','IT',0,'2020-01-01');
insert into `userinfo` (`username`,`password`,`department`,`isadmin`,`created`) values ('heruru','1234','User',0,'2020-02-20');
insert into `userinfo` (`username`,`password`,`department`,`isadmin`,`created`) values ('idhon','123','bd',0,'2019-01-01');
