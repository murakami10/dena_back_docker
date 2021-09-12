CREATE DATABASE test_database;

USE test_database;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) NOT NULL,
    `age`  INT,
    PRIMARY KEY (id)
);

INSERT INTO `user` (`name`, `age`) VALUES ('murakami', 23);
INSERT INTO `user` (`name`, `age`) VALUES ('yamada', 40);
