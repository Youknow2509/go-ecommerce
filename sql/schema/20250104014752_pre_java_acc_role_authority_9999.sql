-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_java_acc_role_authority_9999`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `authority_id` BIGINT NOT NULL,
    `role_id` BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_java_acc_role_authority_9999`;
-- +goose StatementEnd
