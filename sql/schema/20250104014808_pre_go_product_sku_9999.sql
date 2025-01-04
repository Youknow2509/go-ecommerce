-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pre_go_product_sku_9999`(
    `sku_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `spu_id` BIGINT NOT NULL,
    `sku_price` DECIMAL(8, 2) NOT NULL,
    `sku_stock` BIGINT NOT NULL,
    `sku_attribute_value` VARCHAR(255) NOT NULL,
    `sku_created_at` TIMESTAMP NOT NULL,
    `sku_updated_at` TIMESTAMP NOT NULL,
    `sku_deleted_at` TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop TABLE IF EXISTS `pre_go_product_sku_9999`;
-- +goose StatementEnd
