CREATE TABLE `pre_go_acc_user_9999` (
    `user_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'User  ID',
    `user_account` VARCHAR(255) NOT NULL COMMENT 'User  account',
    `user_nickname` VARCHAR(255) NULL COMMENT 'User  nickname',
    `user_avatar` VARCHAR(255) NULL COMMENT 'User  avatar',
    `user_state` TINYINT UNSIGNED NOT NULL COMMENT 'User  state: 0-Locked, 1-Activated, 2-Not Activated',
    `user_mobile` VARCHAR(20) NULL COMMENT 'Mobile phone number',
    `user_gender` TINYINT UNSIGNED NULL COMMENT 'User  gender: 0-Secret, 1-Male, 2-Female',
    `user_birthday` DATE NULL COMMENT 'User  birthday',
    `user_email` VARCHAR(255) NULL COMMENT 'User  email address',
    `user_is_authentication` TINYINT UNSIGNED NOT NULL COMMENT 'Authentication status: 0-Not Authenticated, 1-Pending, 2-Authenticated, 3-Failed',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update time'
);
ALTER TABLE
    `pre_go_acc_user_9999` ADD UNIQUE `pre_go_acc_user_9999_user_account_unique`(`user_account`);
ALTER TABLE
    `pre_go_acc_user_9999` ADD INDEX `pre_go_acc_user_9999_user_state_index`(`user_state`);
ALTER TABLE
    `pre_go_acc_user_9999` ADD INDEX `pre_go_acc_user_9999_user_mobile_index`(`user_mobile`);
ALTER TABLE
    `pre_go_acc_user_9999` ADD INDEX `pre_go_acc_user_9999_user_email_index`(`user_email`);
ALTER TABLE
    `pre_go_acc_user_9999` ADD INDEX `pre_go_acc_user_9999_user_is_authentication_index`(`user_is_authentication`);

CREATE TABLE `pre_go_acc_user_base_9999` (
    `user_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_account` VARCHAR(255) NOT NULL,
    `user_password` VARCHAR(255) NOT NULL,
    `user_salt` VARCHAR(255) NOT NULL,
    `user_login_time` TIMESTAMP NULL DEFAULT NULL,
    `user_logout_time` TIMESTAMP NULL DEFAULT NULL,
    `user_login_ip` VARCHAR(45) NULL DEFAULT NULL,
    `user_created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `user_updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
ALTER TABLE
    `pre_go_acc_user_base_9999` ADD UNIQUE `pre_go_acc_user_base_9999_user_account_unique`(`user_account`);

CREATE TABLE `pre_go_acc_user_verify_9999`(
    `verify_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `verify_otp` VARCHAR(6) NOT NULL,
    `verify_key` VARCHAR(255) NOT NULL,
    `verify_key_hash` VARCHAR(255) NOT NULL,
    `verify_type` INT NULL DEFAULT '1',
    `is_verified` INT NULL,
    `is_deleted` INT NULL,
    `verify_created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP(), `verify_updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP());
ALTER TABLE
    `pre_go_acc_user_verify_9999` ADD INDEX `pre_go_acc_user_verify_9999_verify_otp_index`(`verify_otp`);
ALTER TABLE
    `pre_go_acc_user_verify_9999` ADD UNIQUE `pre_go_acc_user_verify_9999_verify_key_unique`(`verify_key`);

CREATE TABLE `pre_java_acc_role_9999`(
    `role_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name_role` VARCHAR(255) NULL,
    `create_user_id` BIGINT NULL
);

CREATE TABLE `pre_java_acc_user_role_9999`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NULL,
    `role_id` BIGINT NULL
);
ALTER TABLE
    `pre_java_acc_user_role_9999` ADD INDEX `pre_java_acc_user_role_9999_user_id_index`(`user_id`);
ALTER TABLE
    `pre_java_acc_user_role_9999` ADD INDEX `pre_java_acc_user_role_9999_role_id_index`(`role_id`);

CREATE TABLE `pre_java_acc_authority_9999`(
    `authority_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `authority_name` VARCHAR(255) NOT NULL,
    `menu_pid` VARCHAR(255) NULL DEFAULT '0',
    `authority_url` VARCHAR(255) NOT NULL,
    `authority_prefix` VARCHAR(255) NOT NULL,
    `create_user_id` BIGINT NOT NULL
);

CREATE TABLE `pre_java_acc_role_authority_9999`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `authority_id` BIGINT NOT NULL,
    `role_id` BIGINT NOT NULL
);

CREATE TABLE `pre_go_product_base_9999`(
    `spu_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `category_id` BIGINT NOT NULL,
    `shop_id` BIGINT NOT NULL,
    `brand_id` BIGINT NOT NULL,
    `spu_name` VARCHAR(255) NOT NULL,
    `spu_description` VARCHAR(255) NOT NULL,
    `spu_img_url` VARCHAR(255) NOT NULL,
    `spu_video_url` VARCHAR(255) NOT NULL,
    `spu_sort` BIGINT NOT NULL,
    `spu_price` DECIMAL(8, 2) NOT NULL,
    `spu_status` BIGINT NOT NULL,
    `spu_ created_at` TIMESTAMP NOT NULL,
    `spu_ updated_at` TIMESTAMP NOT NULL,
    `spu_deleted_at` TIMESTAMP NOT NULL
);
ALTER TABLE
    `pre_go_product_base_9999` ADD UNIQUE `pre_go_product_base_9999_spu_name_spu_deleted_at_unique`(`spu_name`, `spu_deleted_at`);

CREATE TABLE `pre_go_product_sku_9999`(
    `sku_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `spu_id` BIGINT NOT NULL,
    `sku_price` DECIMAL(8, 2) NOT NULL,
    `sku_stock` BIGINT NOT NULL,
    `sku_attribute_value` VARCHAR(255) NOT NULL,
    `sku_created_at` TIMESTAMP NOT NULL,
    `sku_updated_at` TIMESTAMP NOT NULL,
    `sku_deleted_at` TIMESTAMP NOT NULL
);

CREATE TABLE `pre_go_product_category_9999`(
    `category_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `parent_id` BIGINT NOT NULL,
    `category_name` VARCHAR(255) NOT NULL,
    `has_active_children` BOOLEAN NOT NULL,
    `category_spu_count` BIGINT NOT NULL,
    `category_status` BIGINT NOT NULL,
    `category_description` VARCHAR(255) NOT NULL,
    `category_icon` VARCHAR(255) NOT NULL,
    `category_sort` BIGINT NOT NULL,
    `category_deleted_at` TIMESTAMP NOT NULL,
    `category_created_at` TIMESTAMP NOT NULL,
    `category_updated_at` TIMESTAMP NOT NULL
);
ALTER TABLE
    `pre_go_product_base_9999` ADD CONSTRAINT `pre_go_product_base_9999_category_id_foreign` FOREIGN KEY(`category_id`) REFERENCES `pre_go_product_category_9999`(`category_id`);
ALTER TABLE
    `pre_go_product_sku_9999` ADD CONSTRAINT `pre_go_product_sku_9999_sku_id_foreign` FOREIGN KEY(`sku_id`) REFERENCES `pre_go_product_base_9999`(`spu_id`);
ALTER TABLE
    `pre_go_acc_user_verify_9999` ADD CONSTRAINT `pre_go_acc_user_verify_9999_verify_key_foreign` FOREIGN KEY(`verify_key`) REFERENCES `pre_go_acc_user_base_9999`(`user_account`);
ALTER TABLE
    `pre_go_acc_user_9999` ADD CONSTRAINT `pre_go_acc_user_9999_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `pre_java_acc_user_role_9999`(`id`);
ALTER TABLE
    `pre_java_acc_authority_9999` ADD CONSTRAINT `pre_java_acc_authority_9999_authority_id_foreign` FOREIGN KEY(`authority_id`) REFERENCES `pre_java_acc_role_authority_9999`(`id`);
ALTER TABLE
    `pre_java_acc_role_9999` ADD CONSTRAINT `pre_java_acc_role_9999_role_id_foreign` FOREIGN KEY(`role_id`) REFERENCES `pre_java_acc_role_authority_9999`(`id`);
ALTER TABLE
    `pre_java_acc_user_role_9999` ADD CONSTRAINT `pre_java_acc_user_role_9999_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `pre_java_acc_role_9999`(`role_id`);
ALTER TABLE
    `pre_go_acc_user_base_9999` ADD CONSTRAINT `pre_go_acc_user_base_9999_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `pre_go_acc_user_9999`(`user_id`);