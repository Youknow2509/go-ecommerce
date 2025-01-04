-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_java_acc_user_role_9999`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NULL,
    `role_id` BIGINT NULL
);
ALTER TABLE
    `pre_java_acc_user_role_9999` ADD INDEX `pre_java_acc_user_role_9999_user_id_index`(`user_id`);
ALTER TABLE
    `pre_java_acc_user_role_9999` ADD INDEX `pre_java_acc_user_role_9999_role_id_index`(`role_id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_java_acc_user_role_9999`;
-- +goose StatementEnd
