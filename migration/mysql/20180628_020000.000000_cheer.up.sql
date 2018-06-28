SET foreign_key_checks = 1;
SET time_zone = '+00:00';

CREATE TABLE cheers (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    `from` TEXT NOT NULL,
    `to` TEXT NOT NULL,
    `plus` TEXT NOT NULL,
    `note` TEXT NOT NULL,
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    PRIMARY KEY (id)
);