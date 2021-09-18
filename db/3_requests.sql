DROP TABLE IF EXISTS `requests`;

CREATE TABLE `requests` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `sender_id` INT NOT NULL, -- users.id
    `receiver_id` INT NOT NULL, -- users.id
    `message` text not null,
    `created_at` datetime not null default current_timestamp,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`sender_id`) REFERENCES `users`(`id`),
    FOREIGN key (`receiver_id`) REFERENCES `users`(`id`)
);
