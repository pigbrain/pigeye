CREATE TABLE `api` (
  `api_id` bigint(11) NOT NULL AUTO_INCREMENT,
  `service_id` bigint(11) NOT NULL,
  `name` varchar(30) NOT NULL,
  `description` varchar(100) NOT NULL,
  `url` varchar(1000) NOT NULL,
  `user_agent` varchar(100) NOT NULL,
  `content_type` varchar(100) NOT NULL,
  `method` varchar(10) NOT NULL,
  `request_body` text NOT NULL,
  `status` int(11) NOT NULL,
  `response_body` text NOT NULL,
  `success` tinyint(1) DEFAULT 0,
  `notification_script` text,
  `creation_datetime` datetime NOT NULL,
  `updated_datetime` datetime NOT NULL,
  PRIMARY KEY (`api_id`),
  KEY `service_api_idx1` (`service_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;

CREATE TABLE `service` (
  `service_id` bigint(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `description` varchar(50) NOT NULL,
  `creation_datetime` datetime NOT NULL,
  `updated_datetime` datetime NOT NULL,
  PRIMARY KEY (`service_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
