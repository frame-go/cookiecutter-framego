CREATE DATABASE IF NOT EXISTS `{{ cookiecutter.service_name }}_db`
    DEFAULT CHARACTER SET = ascii
    DEFAULT COLLATE = ascii_general_ci;

USE `{{ cookiecutter.service_name }}_db`;

CREATE TABLE IF NOT EXISTS `{{ cookiecutter.service_name }}_db`.`object_tab`
(
    `id`          BIGINT UNSIGNED                                                NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(50)                                                    NOT NULL,
    `data`        VARCHAR(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `create_time` INT UNSIGNED                                                   NOT NULL,
    `update_time` INT UNSIGNED                                                   NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY idx_name (`name`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = ascii
    DEFAULT COLLATE = ascii_general_ci;