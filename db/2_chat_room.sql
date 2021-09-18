DROP TABLE IF EXISTS `rooms`;

CREATE TABLE `rooms` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `frend_id` INT NOT NULL,
    PRIMARY KEY (id)
);


DROP TABLE IF EXISTS `chats`;

CREATE TABLE `chats` (
    `id`   INT NOT NULL AUTO_INCREMENT,
    `room_id` INT NOT NULL, -- rooms.id
    `sender_id` INT NOT NULL, -- users.id
    `text` text NOT NULL,
    `created_at` datetime NOT NULL,
    PRIMARY KEY (id)
);

-- TODO rooms.id index