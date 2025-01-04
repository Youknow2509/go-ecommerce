-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_java_acc_authority_9999`(
    `authority_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `authority_name` VARCHAR(255) NOT NULL,
    `menu_pid` VARCHAR(255) NULL DEFAULT '0',
    `authority_url` VARCHAR(255) NOT NULL,
    `authority_prefix` VARCHAR(255) NOT NULL,
    `create_user_id` BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_java_acc_authority_9999`;
-- +goose StatementEnd
