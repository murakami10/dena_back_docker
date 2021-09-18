DROP TABLE IF EXISTS `rooms`;

CREATE TABLE `rooms` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `friend_id` INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
    FOREIGN KEY (`friend_id`) REFERENCES `users`(`id`),
    INDEX (`user_id`, `friend_id`)
);

DROP TABLE IF EXISTS `chats`;

CREATE TABLE `chats` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `room_id` INT NOT NULL, 
    `sender_id` INT NOT NULL,
    `text` text NOT NULL,
    `created_at` datetime NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`),
    FOREIGN KEY (`sender_id`) REFERENCES `users`(`id`),
    INDEX (`room_id`)
);