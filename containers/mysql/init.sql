CREATE USER 'gopher'@'localhost' IDENTIFIED BY 'gopher';
GRANT ALL PRIVILEGES ON *.* TO 'gopher'@'localhost' WITH GRANT OPTION;

CREATE USER 'gopher'@'%' IDENTIFIED BY 'gopher';
GRANT ALL PRIVILEGES ON *.* TO 'gopher'@'%' WITH GRANT OPTION;


CREATE DATABASE app;

CREATE TABLE `app`.`item_categories` (
        `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY
       ,`name` VARCHAR(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;

CREATE TABLE `app`.`items` (
        `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY
       ,`name_g` VARCHAR(200) NOT NULL
       ,`category_id` INT NOT NULL REFERENCES `app`.`item_category` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
       ,`modified_at` DATETIME NOT NULL
       ,`created_at` DATETIME NOT NULL
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;


INSERT INTO `app`.`item_categories` (`id`,`name`) VALUES (1,"a"),(2,"b"),(3,"c");

INSERT INTO `app`.`items` (`name_g`,`category_id`,`modified_at`,`created_at`) VALUES ("book", 1,NOW(),NOW()),("note",2,NOW(),NOW()),("pc",3,NOW(),NOW());
