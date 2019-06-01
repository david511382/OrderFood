-- MySQL Script generated by MySQL Workbench
-- Sat Jun  1 17:07:49 2019
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema orderfood_menu
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `orderfood_menu` ;

-- -----------------------------------------------------
-- Schema orderfood_menu
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `orderfood_menu` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema orderfood_member
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `orderfood_member` ;

-- -----------------------------------------------------
-- Schema orderfood_member
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `orderfood_member` ;
USE `orderfood_menu` ;

-- -----------------------------------------------------
-- Table `orderfood_menu`.`shop`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_menu`.`shop` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `orderfood_menu`.`item`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_menu`.`item` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `shop_id` INT NOT NULL,
  `price` INT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  INDEX `shop_fk_idx` (`shop_id` ASC) VISIBLE,
  CONSTRAINT `shop_fk`
    FOREIGN KEY (`shop_id`)
    REFERENCES `orderfood_menu`.`shop` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `orderfood_menu`.`selection`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_menu`.`selection` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `orderfood_menu`.`option`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_menu`.`option` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `is_optional` TINYINT(1) NOT NULL DEFAULT 1,
  `select_num` INT NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `orderfood_menu`.`option_selection`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_menu`.`option_selection` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `option_id` INT NOT NULL,
  `selection_id` INT NOT NULL,
  `price` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `option_fk_idx` (`option_id` ASC) VISIBLE,
  INDEX `selection_fk_idx` (`selection_id` ASC) VISIBLE,
  CONSTRAINT `option_fk`
    FOREIGN KEY (`option_id`)
    REFERENCES `orderfood_menu`.`option` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `selection_fk`
    FOREIGN KEY (`selection_id`)
    REFERENCES `orderfood_menu`.`selection` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `orderfood_menu`.`item_option`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_menu`.`item_option` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `item_id` INT NOT NULL,
  `option_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `item_fk_idx` (`item_id` ASC) VISIBLE,
  INDEX `option_fk_idx` (`option_id` ASC) VISIBLE,
  CONSTRAINT `item_option_fk`
    FOREIGN KEY (`item_id`)
    REFERENCES `orderfood_menu`.`item` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `option_item_fk`
    FOREIGN KEY (`option_id`)
    REFERENCES `orderfood_menu`.`option` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

USE `orderfood_member` ;

-- -----------------------------------------------------
-- Table `orderfood_member`.`member`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_member`.`member` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
