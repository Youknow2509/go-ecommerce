-- +goose Up
-- +goose StatementBegin
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP IF EXISTS TABLE `pre_go_acc_user_verify_9999`;
-- +goose StatementEnd
