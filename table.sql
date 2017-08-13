CREATE TABLE `pigeye`.`service` (
  `id` BIGINT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(30) NOT NULL,
  `description` VARCHAR(50) NOT NULL,
  `creation_datetime` DATETIME NOT NULL,
  `updated_datetime` DATETIME NOT NULL,
  PRIMARY KEY (`id`));


CREATE TABLE `pigeye`.`service_api` (
  `id` BIGINT(11) NOT NULL AUTO_INCREMENT,
  `service_id` BIGINT(11) NOT NULL,
  `name` VARCHAR(30) NOT NULL,
  `description` VARCHAR(100) NOT NULL,
  `url` VARCHAR(1000) NOT NULL,
  `content_type` VARCHAR(500) NOT NULL,
  `method` VARCHAR(10) NOT NULL,
  `request_body` TEXT NOT NULL,
  `status` INT NOT NULL,
  `response_body` TEXT NOT NULL,
  `api` TINYINT(1) NOT NULL,
  `creation_datetime` DATETIME NOT NULL,
  `updated_datetime` DATETIME NOT NULL,
  PRIMARY KEY (id),
  INDEX service_api_idx1 (service_id) );