-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_verify_001` (
    `verify_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Verification ID',
    `verify_otp` VARCHAR(6) NOT NULL COMMENT 'One-time password',
    `verify_key` VARCHAR(255) NOT NULL COMMENT 'Verification key - email address, phone number, ....',
    `verify_key_hash` VARCHAR(255) NOT NULL COMMENT 'Hash of the verification key',
    `verify_type` INT NULL DEFAULT 1 COMMENT 'Verification type (e.g., 1 for email)',
    `is_verified` INT NULL COMMENT 'Verification status: 1 for verified, 0 for not verified',
    `is_deleted` INT NULL COMMENT 'Deletion flag: 1 for deleted, 0 for not deleted',
    `verify_created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time',
    `verify_updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update time'
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `user_verify_001`
ADD INDEX `user_verify_001_verify_otp_index`(`verify_otp`);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE `user_verify_001`
ADD UNIQUE `user_verify_001_verify_key_unique`(`verify_key`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_verify_001`;
-- +goose StatementEnd
