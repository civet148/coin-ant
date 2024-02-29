/*
SQLyog Trial v13.1.8 (64 bit)
MySQL - 8.0.23 : Database - coin-ant
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`coin-ant` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `coin-ant`;

/*Table structure for table `rich_list` */

CREATE TABLE `rich_list` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'auto-incr id',
  `chain_id` int NOT NULL DEFAULT '0' COMMENT 'block chain id [0=Bitcoin 1=Ethereum]',
  `chain_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'block chain name [Bitcoin/Ethereum]',
  `symbol` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'token symbol',
  `address` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'owner address',
  `balance` decimal(60,0) NOT NULL COMMENT 'owner balance',
  `contract_addr` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'contract address (empty means native token)',
  `is_ok` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'is cracked (0=no 1=yes)',
  `key_phrase` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'phrase words',
  `private_key` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'private key of hex string',
  `derivaton_path` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'derivation path',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  `extra_data` longtext COLLATE utf8mb4_unicode_ci COMMENT 'extra data',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQ_CHAIN_ADDRESS` (`chain_id`,`symbol`,`address`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*Table structure for table `token_list` */

CREATE TABLE `token_list` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'auto-incr id',
  `chain_id` bigint NOT NULL DEFAULT '0' COMMENT 'block chain id [0=Bitcoin 1=Ethereum]',
  `chain_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'block chain name [Bitcoin/Ethereum]',
  `symbol` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'token symbol',
  `decimals` int NOT NULL DEFAULT '0' COMMENT 'token decimals',
  `contract_addr` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'contract address (empty means native token)',
  `min_balance` decimal(50,5) NOT NULL DEFAULT '0.00000' COMMENT 'account''s minimal balance  to attack',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'is deleted (0=no 1=yes)',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updte time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
