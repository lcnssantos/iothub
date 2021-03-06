CREATE TABLE `users`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `password` TEXT NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `active` TINYINT(1) NOT NULL DEFAULT '0',
    `createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    UNIQUE `email` (`email`)
);

CREATE TABLE `accounts` 
( 
    `id` BIGINT NOT NULL AUTO_INCREMENT, 
    `login` TEXT NOT NULL,
    `password` TEXT NOT NULL,
    `userId` BIGINT NOT NULL,
    `createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);