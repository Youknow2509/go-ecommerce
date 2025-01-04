-- +goose Up
-- +goose StatementBegin
ALTER TABLE
    `pre_go_product_base_9999` ADD CONSTRAINT `pre_go_product_base_9999_category_id_foreign` FOREIGN KEY(`category_id`) REFERENCES `pre_go_product_category_9999`(`category_id`);
ALTER TABLE
    `pre_go_product_sku_9999` ADD CONSTRAINT `pre_go_product_sku_9999_sku_id_foreign` FOREIGN KEY(`sku_id`) REFERENCES `pre_go_product_base_9999`(`spu_id`);
ALTER TABLE
    `pre_go_acc_user_verify_9999` ADD CONSTRAINT `pre_go_acc_user_verify_9999_verify_key_foreign` FOREIGN KEY(`verify_key`) REFERENCES `pre_go_acc_user_base_9999`(`user_account`);
ALTER TABLE
    `pre_go_acc_user_9999` ADD CONSTRAINT `pre_go_acc_user_9999_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `pre_java_acc_user_role_9999`(`id`);
ALTER TABLE
    `pre_java_acc_authority_9999` ADD CONSTRAINT `pre_java_acc_authority_9999_authority_id_foreign` FOREIGN KEY(`authority_id`) REFERENCES `pre_java_acc_role_authority_9999`(`id`);
ALTER TABLE
    `pre_java_acc_role_9999` ADD CONSTRAINT `pre_java_acc_role_9999_role_id_foreign` FOREIGN KEY(`role_id`) REFERENCES `pre_java_acc_role_authority_9999`(`id`);
ALTER TABLE
    `pre_java_acc_user_role_9999` ADD CONSTRAINT `pre_java_acc_user_role_9999_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `pre_java_acc_role_9999`(`role_id`);
ALTER TABLE
    `pre_go_acc_user_base_9999` ADD CONSTRAINT `pre_go_acc_user_base_9999_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `pre_go_acc_user_9999`(`user_id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
