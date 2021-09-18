DROP TABLE IF EXISTS `requests`;

CREATE TABLE `requests` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `sender_id` INT NOT NULL, -- users.id
    `reciever_id` INT NOT NULL, -- users.id
    `message` text not null,
    `created_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`sender_id`) REFERENCES `users`(`id`),
    FOREIGN key (`reciever_id`) REFERENCES `users`(`id`)
);
