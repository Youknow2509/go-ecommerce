-- name: GetTokenInfoWithId :one
SELECT 
    *
FROM `acc_token_0001` 
WHERE `id` = ?
LIMIT 1;

-- name: GetTokenInfoWithUserId :many
SELECT 
    *
FROM `acc_token_0001`
WHERE `user_id` = ?
ORDER BY `created_at` DESC
LIMIT 1;

-- name: GetTokenInfoWithAccessToken :one
SELECT 
    *
FROM `acc_token_0001`
WHERE `access_token` = ?
LIMIT 1;

-- name: GetTokenInfoWithRefreshToken :one
SELECT 
    *
FROM `acc_token_0001`
WHERE `refresh_token` = ?
LIMIT 1;

-- name: GetToken :one
SELECT 
    *
FROM `acc_token_0001`
WHERE `access_token` = ? AND `refresh_token` = ?
LIMIT 1;

-- name: CreateToken :exec
INSERT INTO `acc_token_0001` (
    `user_id`, 
    `is_refresh`, 
    `access_token`, 
    `refresh_token`, 
    `access_token_expires_at`, 
    `refresh_token_expires_at`,
    `created_at`,
    `updated_at`
) VALUES (?, 1, ?, ?, ?, ?, NOW(), NOW());

-- name: IsTokenRefresh :one
SELECT 
    `id`
FROM `acc_token_0001`
WHERE `refresh_token` = ? and `access_token` = ? and `is_refresh` = 1
LIMIT 1;

-- name: IsTokenRefreshWithId :one
SELECT 
    `id`
FROM `acc_token_0001`
WHERE `id` = ? and `is_refresh` = 1
LIMIT 1;

-- name: BlockRefreshToken :exec
UPDATE `acc_token_0001`
SET `is_refresh` = 0, `updated_at` = NOW()
WHERE `refresh_token` = ? and `access_token` = ?;

-- name: BlockRefreshTokenWithId :exec
UPDATE `acc_token_0001`
SET `is_refresh` = 0, `updated_at` = NOW()
WHERE `id` = ?;

-- name: DeleteTokenWithId :exec
DELETE FROM `acc_token_0001`
WHERE `id` = ?;

