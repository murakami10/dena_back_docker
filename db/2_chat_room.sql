DROP TABLE IF EXISTS `rooms`;

CREATE TABLE `rooms` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `room_members`;

CREATE TABLE `room_members` (
    `id`       INT NOT NULL AUTO_INCREMENT,
    `user_id`  INT NOT NULL,
    `room_id`  INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);


DROP TABLE IF EXISTS `chats`;

CREATE TABLE `chats` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `room_id` INT NOT NULL, -- rooms.id
    `sender_id` INT NOT NULL, -- users.id
    `text` text NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`),
    FOREIGN KEY (`sender_id`) REFERENCES `users`(`id`)
);
