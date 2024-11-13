CREATE TABLE IF NOT EXISTS `users`(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `mobile` VARCHAR(255) NOT NULL,
    `role` ENUM('renter','owner') NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS `cars`(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `owner_id` INT NOT NULL,
    `model` VARCHAR(255) NOT NULL,
    `rent_per_hour` DECIMAL(10,2) NOT NULL,
    `location` POINT NOT NULL,
    `is_available` BOOLEAN NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    FOREIGN KEY (`owner_id`) REFERENCES users(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `availability` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `car_id` INT NOT NULL,
    `start_time` BIGINT NOT NULL,
    `end_time` BIGINT NOT NULL,
    `status` ENUM('available', 'unavailable') NOT NULL,
    FOREIGN KEY (`car_id`) REFERENCES cars(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `rental_transactions` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `renter_id` INT NOT NULL,
    `car_id` INT NOT NULL,
    `start_time` BIGINT NOT NULL,
    `end_time` BIGINT NOT NULL,
    `total_cost` DECIMAL(10,2) NOT NULL,
    `status` ENUM('pending', 'active', 'completed', 'cancelled') NOT NULL,
    FOREIGN KEY (`renter_id`) REFERENCES users(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`car_id`) REFERENCES cars(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `chat_initiations` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `renter_id` INT NOT NULL,
    `owner_id` INT NOT NULL,
    `car_id` INT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    FOREIGN KEY (`renter_id`) REFERENCES users(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`owner_id`) REFERENCES users(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`car_id`) REFERENCES cars(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `messages` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `chat_id` INT NOT NULL,
    `sender_id` INT NOT NULL,
    `message` TEXT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    FOREIGN KEY (`chat_id`) REFERENCES chat_initiations(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`sender_id`) REFERENCES users(`id`) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `message_notifications` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `chat_id` INT NOT NULL,
    `unread_count` INT DEFAULT 0,
    `last_message_id` INT,
    `created_at` BIGINT,
    `updated_at` BIGINT,
    FOREIGN KEY (`user_id`) REFERENCES users(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`chat_id`) REFERENCES chat_initiations(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`last_message_id`) REFERENCES messages(`id`) ON DELETE SET NULL
);
