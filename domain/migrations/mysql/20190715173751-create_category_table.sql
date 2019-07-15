-- +migrate Up

CREATE TABLE IF NOT EXISTS `privy-id`.`categories`
(
    `id`     INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`   VARCHAR(255) NOT NULL,
    `enable` TINYINT      NOT NULL,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `privy-id`.categories;