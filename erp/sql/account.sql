CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `no` varchar(50) NOT NULL DEFAULT '',
  `py` varchar(50) NOT NULL DEFAULT '',
  `note` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`) 
  UNIQUE KEY `no_UNIQUE` (`no`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
