DROP TABLE IF EXISTS `friends`;

CREATE TABLE `friends` (
    `user_id`  INT NOT NULL,
    `friend_user_id`  INT NOT NULL,
    FOREIGN KEY (`friend_user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);
