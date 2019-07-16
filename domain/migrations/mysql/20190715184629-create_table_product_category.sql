-- +migrate Up

CREATE TABLE IF NOT EXISTS `privy-id`.`product_categories`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `product_id`  INT UNSIGNED NOT NULL,
    `category_id` INT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_product_category_product_idx` (`product_id` ASC),
    INDEX `fk_product_category_category1_idx` (`category_id` ASC),
    CONSTRAINT `fk_product_category_product`
        FOREIGN KEY (`product_id`)
            REFERENCES `privy-id`.`products` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_product_category_category1`
        FOREIGN KEY (`category_id`)
            REFERENCES `privy-id`.`categories` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down

DROP TABLE IF EXISTS `privy-id`.product_categories;