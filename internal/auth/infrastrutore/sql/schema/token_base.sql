-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `acc_token_0001` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `is_refresh` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'indicates if the token is a refresh token',
    `access_token` varchar(255) NOT NULL,
    `refresh_token` varchar(255) NOT NULL,
    `access_token_expires_at` datetime NOT NULL,
    `refresh_token_expires_at` datetime NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `acc_user_base_0001`
ADD INDEX `acc_user_base_0001_user_id_index`(`user_id`);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `acc_user_base_0001`
ADD INDEX `acc_user_base_0001_refresh_token_index`(`refresh_token`);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `acc_user_base_0001`
ADD INDEX `acc_user_base_0001_refresh_token_expires_at_index`(`refresh_token_expires_at`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `acc_token_0001`;
-- +goose StatementEnd
