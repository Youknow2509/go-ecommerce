-- +goose Up
-- +goose StatementBegin
ALTER TABLE `pre_go_acc_user_base_9999`
ADD COLUMN `is_two_factor_enabled` INT(1) DEFAULT 0 COMMENT 'authentication two factor enabled for the user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `pre_go_acc_user_base_9999`
DROP COLUMN `is_two_factor_enabled`;
-- +goose StatementEnd
