CREATE TABLE IF NOT EXISTS `btc_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `eth_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `bch_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `bsv_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `dash_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `doge_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `ltc_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `qtum_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `trx_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `sol_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `fil_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `luna_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `ada_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `xrp_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `xmr_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `near_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `dot_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `avax_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `ftm_address` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `address` varchar(300) NOT NULL COMMENT '地址',
    `priv_key` text NOT NULL,
    `is_used` tinyint(4) DEFAULT 0,
    `multi` tinyint(4) DEFAULT 1,
    `create_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
    `update_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6) ON UPDATE current_timestamp(6),
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;