-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_java_acc_role_9999`(
    `role_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name_role` VARCHAR(255) NULL,
    `create_user_id` BIGINT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_java_acc_role_9999`;
-- +goose StatementEnd
