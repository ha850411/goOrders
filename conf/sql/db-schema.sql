CREATE TABLE `menu` (
  `id` int NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `active` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sort` int NOT NULL,
  `updateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `menu` (`id`, `name`, `icon`, `path`, `active`, `sort`, `updateTime`) VALUES
(1, '存貨管理', 'fa-solid fa-cart-shopping', '/product', 'product', 1, '2023-02-10 11:07:15'),
(2, '存貨分析', 'fa-solid fa-chart-simple', '/analysis', 'analysis', 2, '2023-02-10 11:07:22'),
(3, '訂單管理', 'fa-solid fa-list', '/order', 'order', 3, '2023-02-10 11:17:56'),
(4, '產品類別', 'fa-solid fa-tag', '/product/type', 'product.type', 4, '2023-02-15 10:10:48'),
(5, '輪播管理', 'fa-regular fa-image', '/banner', 'banner', 5, '2023-03-08 18:05:41'),
(6, '關於我們', 'fa-solid fa-circle-info', '/about', 'about', 6, '2023-03-09 10:08:02'),
(7, '聯絡我們', 'fa-solid fa-phone', '/contact', 'contact', 7, '2023-03-09 10:08:08'),
(8, '設定', 'fa-solid fa-gear', '/setting', 'setting', 8, '2023-03-09 10:08:14');

CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `contact` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remark` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `amount` int NOT NULL,
  `amountNotice` int NOT NULL DEFAULT '0',
  `price` int NOT NULL DEFAULT '0',
  `discount_price` INT NOT NULL DEFAULT '0',
  `status` int NOT NULL DEFAULT '1' COMMENT '0:刪除 1:存在',
  `type` int NOT NULL DEFAULT '0' COMMENT '商品分類',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '產品介紹\r\n',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `products_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `updateDate` date NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `pid_2` (`pid`,`updateDate`),
  KEY `pid` (`pid`),
  KEY `pid_3` (`pid`,`updateDate`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `linebot_token` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'linebot access_token',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `users` VALUES 
(1,'admin','admin',NULL),
(2,'eason','eason',NULL),
(3,'niigataiwan','niigataiwan',NULL);

CREATE TABLE `orders_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL,
  `pid` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `price` int NOT NULL DEFAULT '0',
  `total` int NOT NULL DEFAULT '0',
  `updateTime` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  KEY `pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `products_picture` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pid` int NOT NULL,
  `picture` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `products_type`;
CREATE TABLE `products_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '0:刪除 1:存在',
  `createTime` datetime NOT NULL,
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `products_type` (`id`, `name`, `status`, `createTime`, `updateTime`) VALUES
(1, '酒類', 1, '2023-02-14 09:49:43', '2023-02-15 10:45:52');

CREATE TABLE `banner` (
  `id` int NOT NULL AUTO_INCREMENT,
  `picture` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sort` int NOT NULL,
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `about` (
  `id` int NOT NULL DEFAULT '1',
  `title1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content1` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `picture1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `title2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content2` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `picture2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `contact` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `contact` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;