CREATE DATABASE IF NOT EXISTS vetautet
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_unicode_ci;

-- 1. ticket table
CREATE TABLE IF NOT EXISTS `vetautet`.`ticket` (
     `id` BIGINT(20)  NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `name` VARCHAR(50) NOT NULL COMMENT 'ticket name',
    `desc` TEXT COMMENT 'ticket description',
    `start_time` DATETIME NOT NULL COMMENT 'ticket sale start time',
    `end_time` DATETIME    NOT NULL COMMENT 'ticket sale end time',
    `status`   INT(11) NOT NULL DEFAULT 0 COMMENT 'ticket sale activity status', -- 0: deactive, 1: activity
    `updated_at` DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at`   DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`),
    KEY `idx_end_time` (`end_time`), -- Very high query runtime
    KEY `idx_start_time` (`start_time`), -- Very high query runtime
    KEY `idx_status` (`status`) -- Very high query runtime
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'ticket table';

-- 2. ticket detail (item) table
CREATE TABLE IF NOT EXISTS `vetautet`.`ticket_item` (
    `id` BIGINT(20)  NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `name` VARCHAR(50) NOT NULL COMMENT 'Ticket title',
    `description` TEXT COMMENT 'Ticket description',
    `stock_initial` INT(11) NOT NULL DEFAULT 0 COMMENT 'Initial stock quantity (e.g., 1000 tickets)',
    `stock_available` INT(11) NOT NULL DEFAULT 0 COMMENT 'Current available stock (e.g., 900 tickets)',
    `is_stock_prepared` BOOLEAN NOT NULL DEFAULT 0 COMMENT 'Indicates if stock is pre-warmed (0/1)', -- warm up cache
    `price_original` BIGINT(20) NOT NULL COMMENT 'Original ticket price', -- Giá gốc: ví dụ: 100K/ticket
    `price_flash` BIGINT(20) NOT NULL COMMENT 'Discounted price during flash sale', -- Giảm giá khung giờ vàng : ví dụ: 10K/ticket
    `sale_start_time` DATETIME NOT NULL COMMENT 'Flash sale start time',
    `sale_end_time` DATETIME NOT NULL COMMENT 'Flash sale end time',
    `status` INT(11) NOT NULL DEFAULT 0 COMMENT 'Ticket status (e.g., active/inactive)', -- Trạng thái của vé (ví dụ: hoạt động/không hoạt động)
    `activity_id` BIGINT(20) NOT NULL COMMENT 'ID of associated activity', -- ID của hoạt động liên quan đến vé
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp of the last update',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation timestamp',
    PRIMARY KEY (`id`),
    KEY `idx_end_time` (`sale_end_time`),
    KEY `idx_start_time` (`sale_start_time`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for ticket details';

-- INSERT MOCK DATA
-- Insert data into `ticket` table
INSERT INTO `vetautet`.`ticket` (`name`, `desc`, `start_time`, `end_time`, `status`, `updated_at`, `created_at`)
VALUES
    ('Đợt Mở Bán Vé Ngày 12/12', 'Sự kiện mở bán vé đặc biệt cho ngày 12/12', '2024-12-12 00:00:00', '2024-12-12 23:59:59', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Đợt Mở Bán Vé Ngày 01/01', 'Sự kiện mở bán vé cho ngày đầu năm mới 01/01', '2025-01-01 00:00:00', '2025-01-01 23:59:59', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert data into `ticket_item` table corresponding to each event in `ticket` table
INSERT INTO `vetautet`.`ticket_item` (`name`, `description`, `stock_initial`, `stock_available`, `is_stock_prepared`, `price_original`, `price_flash`, `sale_start_time`, `sale_end_time`, `status`, `activity_id`, `updated_at`, `created_at`)
VALUES
    -- Ticket items for the 12/12 event
    ('Vé Sự Kiện 12/12 - Hạng Phổ Thông', 'Vé phổ thông cho sự kiện ngày 12/12', 1000, 1000, 0, 100000, 10000, '2024-12-12 00:00:00', '2024-12-12 23:59:59', 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Vé Sự Kiện 12/12 - Hạng VIP', 'Vé VIP cho sự kiện ngày 12/12', 500, 500, 0, 200000, 15000, '2024-12-12 00:00:00', '2024-12-12 23:59:59', 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),

    -- Ticket items for the 01/01 event
    ('Vé Sự Kiện 01/01 - Hạng Phổ Thông', 'Vé phổ thông cho sự kiện ngày 01/01', 2000, 2000, 0, 100000, 10000, '2025-01-01 00:00:00', '2025-01-01 23:59:59', 1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Vé Sự Kiện 01/01 - Hạng VIP', 'Vé VIP cho sự kiện ngày 01/01', 1000, 1000, 0, 200000, 15000, '2025-01-01 00:00:00', '2025-01-01 23:59:59', 1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- 3. order table
CREATE TABLE IF NOT EXISTS `vetautet`.`ticket_order_202502` (
    id INT(8) NOT NULL AUTO_INCREMENT COMMENT 'Unique ticket sales ID',
    user_id INT(8) NOT NULL  COMMENT 'userId',
    order_number VARCHAR(50) NOT NULL COMMENT 'Unique order number',
    total_amount DECIMAL(10,3) NOT NULL COMMENT 'Total payment amount',
    terminal_id VARCHAR(20) NOT NULL COMMENT 'ID of the sales terminal',
    order_date TIMESTAMP NOT NULL COMMENT 'Date and time of the ticket purchase',
    order_notes VARCHAR(100) NULL DEFAULT 'None' COMMENT 'Additional notes for the order',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp of the last update',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation timestamp',
    PRIMARY KEY (id) USING BTREE,
    UNIQUE KEY order_number (order_number),
    KEY order_date (order_date),
    KEY index_usr_id (user_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'order table';
-- insert data: Tạo đơn hàng
INSERT INTO `vetautet`.`ticket_order_202502` (order_number, user_id, total_amount, terminal_id, order_date, order_notes)
VALUES ('ORD2025020001', 1001, 5.600, 'POS001', '2025-02-28 10:00:00', 'Family trip');


-- 4. order detail table
CREATE TABLE IF NOT EXISTS `vetautet`.`ticket_order_details_202502` (
 id INT(8) NOT NULL AUTO_INCREMENT COMMENT 'Unique ticket sales ID',
 ticket_item_id BIGINT(20) NOT NULL COMMENT 'ticket detail ID',
 order_number VARCHAR(50) NOT NULL COMMENT 'Reference to the order number',
 passenger_name VARCHAR(100) NOT NULL COMMENT 'Passenger full name',
 passenger_id VARCHAR(20) NOT NULL COMMENT 'National ID or passport number',
 departure_station VARCHAR(10) NOT NULL COMMENT 'Departure station code (e.g., SGN)',
 arrival_station VARCHAR(10) NOT NULL COMMENT 'Arrival station code (e.g., HAN)',
 departure_time DATETIME NOT NULL COMMENT 'Train departure time',
 seat_class ENUM('Economy', 'Business', 'First') NOT NULL COMMENT 'Seat class type',
 seat_number VARCHAR(10) NOT NULL COMMENT 'Seat number',
 ticket_price DECIMAL(10,3) NOT NULL COMMENT 'Price of the individual ticket',
 updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp of the last update',
 created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation timestamp',
 PRIMARY KEY (id) USING BTREE,
--  FOREIGN KEY (order_number) REFERENCES ticket_order_202502(order_number) ON DELETE CASCADE,
 KEY order_number (order_number),
 KEY ticket_item_id (ticket_item_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'order table';


-- Thêm 4 vé vào đơn hàng
INSERT INTO `vetautet`.`ticket_order_details_202502` (ticket_item_id, order_number, passenger_name, passenger_id, departure_station, arrival_station, departure_time, seat_class, seat_number, ticket_price)
VALUES
    (4,'ORD2025020001', 'Nguyen Van A', 'ID12345678', 'SGN', 'HAN', '2025-03-01 08:00:00', 'Economy', 'A1', 1.400),
    (4,'ORD2025020001', 'Nguyen Van B', 'ID12345679', 'SGN', 'HAN', '2025-03-01 08:00:00', 'Economy', 'A2', 1.400),
    (4,'ORD2025020001', 'Nguyen Van C', 'ID12345680', 'SGN', 'HAN', '2025-03-01 08:00:00', 'Economy', 'A3', 1.400),
    (4,'ORD2025020001', 'Nguyen Van D', 'ID12345681', 'SGN', 'HAN', '2025-03-01 08:00:00', 'Economy', 'A4', 1.400);