CREATE TABLE `go_chat`.`chat_logs` (
  `id` BIGINT NOT NULL auto_increment,
  `message` VARCHAR(1000) NULL,
  `addtime` DATETIME DEFAULT CURRENT_TIMESTAMP ,
  `address` VARCHAR(50) NULL,
  PRIMARY KEY (`id`));