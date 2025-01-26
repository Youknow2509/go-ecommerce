-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_two_factor_9999` (
    `two_factor_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'Khoa chinh tu dong tang',
    `user_id` INT UNSIGNED NOT NULL COMMENT 'Khoa ngoai lien ket voi bang nguoi dung',
    `two_factor_auth_type` ENUM('EMAIL', 'SMS', 'APP') NOT NULL COMMENT 'Loai xac thuc 2fa',
    `two_factor_auth_secret` VARCHAR(255) NOT NULL COMMENT 'Ma bi mat xac thuc 2fa',
    `two_factor_phone` VARCHAR(20) NULL COMMENT 'So dien thoai xac thuc 2fa',
    `two_factor_email` VARCHAR(255) NULL COMMENT 'Email xac thuc 2fa',
    `two_factor_is_active` BOOLEAN NOT NULL DEFAULT TRUE COMMENT 'Trang thai kich hoat xac thuc 2fa',
    `two_factor_created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Thoi gian tao xac thuc 2fa',
    `two_factor_updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Thoi gian cap nhat xac thuc 2fa',

    -- Rang buoc khoa ngoai
    -- FOREIGN KEY (`user_id`) REFERENCES `pre_go_acc_user_base_9999`(`user_id`) ON DELETE CASCADE

    -- Chỉ mục để tối ưu hoá truy vấn theo `user_id` và `auth_type`
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_auth_type` (`two_factor_auth_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='pre_go_acc_user_two_factor_9999';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_two_factor_9999`;
-- +goose StatementEnd
