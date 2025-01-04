-- +goose Up
-- +goose StatementBegin
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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_base_9999`;
-- +goose StatementEnd
