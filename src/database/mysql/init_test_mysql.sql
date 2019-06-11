-- MySQL Script generated by MySQL Workbench
-- Mon Jun 10 19:00:40 2019
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema orderfood_test_menu
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `orderfood_test_menu` ;

-- -----------------------------------------------------
-- Schema orderfood_test_menu
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `orderfood_test_menu` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema orderfood_test_member
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `orderfood_test_member` ;

-- -----------------------------------------------------
-- Schema orderfood_test_member
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `orderfood_test_member` ;
USE `orderfood_test_menu` ;

-- -----------------------------------------------------
-- Table `orderfood_test_menu`.`shops`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_test_menu`.`shops` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `orderfood_test_menu`.`item_groups`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_test_menu`.`item_groups` (
  `id` INT NOT NULL,
  `shop_id` INT NOT NULL,
  `least_select_num` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `shop_fk_idx` (`shop_id` ASC) VISIBLE,
  CONSTRAINT `shop_group_fk`
    FOREIGN KEY (`shop_id`)
    REFERENCES `orderfood_test_menu`.`shops` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `orderfood_test_menu`.`selections`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_test_menu`.`selections` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `group_id` INT NOT NULL,
  `price` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `group_selection_fk_idx` (`group_id` ASC) VISIBLE,
  CONSTRAINT `group_selection_fk`
    FOREIGN KEY (`group_id`)
    REFERENCES `orderfood_test_menu`.`item_groups` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `orderfood_test_menu`.`items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_test_menu`.`items` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `group_id` INT NOT NULL,
  `price` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  INDEX `group_fk_idx` (`group_id` ASC) VISIBLE,
  CONSTRAINT `group_item_fk`
    FOREIGN KEY (`group_id`)
    REFERENCES `orderfood_test_menu`.`item_groups` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

USE `orderfood_test_member` ;

-- -----------------------------------------------------
-- Table `orderfood_test_member`.`members`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `orderfood_test_member`.`members` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
