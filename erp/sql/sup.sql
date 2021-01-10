CREATE TABLE `sup` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `py` varchar(50) NOT NULL DEFAULT '',
  `tel` varchar(50) NOT NULL DEFAULT '',
  `fax` varchar(50) NOT NULL DEFAULT '',
  `cell` varchar(50) NOT NULL DEFAULT '',
  `mail` varchar(50) NOT NULL DEFAULT '',
  `address` varchar(50) NOT NULL DEFAULT '',
  `note` varchar(255) NOT NULL DEFAULT '',
  `owe_amount_limit` int(11) NOT NULL DEFAULT '0',
  `owe_time_limit` int(11) NOT NULL DEFAULT '0',
  `emp_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
)
