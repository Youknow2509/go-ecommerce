-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_go_product_category_9999`(
    `category_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `parent_id` BIGINT NOT NULL,
    `category_name` VARCHAR(255) NOT NULL,
    `has_active_children` BOOLEAN NOT NULL,
    `category_spu_count` BIGINT NOT NULL,
    `category_status` BIGINT NOT NULL,
    `category_description` VARCHAR(255) NOT NULL,
    `category_icon` VARCHAR(255) NOT NULL,
    `category_sort` BIGINT NOT NULL,
    `category_deleted_at` TIMESTAMP NOT NULL,
    `category_created_at` TIMESTAMP NOT NULL,
    `category_updated_at` TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `pre_go_product_category_9999`
-- +goose StatementEnd
