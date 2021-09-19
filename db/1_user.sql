DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `display_name` VARCHAR(255) NOT NULL,
    `twitter_user_id` VARCHAR(255) NOT NULL,
    `icon_url` text,
    PRIMARY KEY (id)
);