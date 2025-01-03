-- +goose Up
-- +goose StatementBegin
CREATE TABLE `go_crm_user` (
  `usr_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
  `usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
  `usr_phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'Phone Number',
  `usr_username` varchar(30) NOT NULL DEFAULT '' COMMENT 'Username',
  `usr_password` varchar(32) NOT NULL DEFAULT '' COMMENT 'Password',
  `usr_created_at` int NOT NULL DEFAULT '0' COMMENT 'Created Time',
  `usr_updated_at` int NOT NULL DEFAULT '0' COMMENT 'Updated Time',
  `usr_create_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Created IP',
  `usr_last_login_at` int NOT NULL DEFAULT '0' COMMENT 'Last Login Time',
  `usr_last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'Last Login IP',
  `usr_login_times` int NOT NULL DEFAULT '0' COMMENT 'Login Times',
  `usr_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Status',
  PRIMARY KEY (`usr_id`),
  KEY `idx_email` (`usr_email`),
  KEY `idx_phone` (`usr_phone`),
  KEY `idx_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Account';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `go_crm_user`;
-- +goose StatementEnd
