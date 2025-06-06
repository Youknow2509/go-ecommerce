-- name: GetUser :one
SELECT
    user_id, 
    user_account, 
    user_nickname, 
    user_avatar, 
    user_state, 
    user_mobile, 
    user_gender, 
    user_birthday, 
    user_email, 
    user_is_authentication, 
    created_at, 
    updated_at
FROM `user_info_001`
WHERE user_id = ? LIMIT 1;

-- name: GetUsers :many
SELECT 
    user_id, 
    user_account, 
    user_nickname, 
    user_avatar, 
    user_state, 
    user_mobile, 
    user_gender, 
    user_birthday, 
    user_email,
    user_is_authentication,
    created_at,
    updated_at
FROM `user_info_001`
WHERE user_id IN (?);

-- name: FindUsers :many
SELECT * FROM user_info_001 
WHERE user_account LIKE ? OR user_nickname LIKE ?;

-- name: ListUsers :many
SELECT * FROM user_info_001 LIMIT ? OFFSET ?;

-- name: RemoveUser :exec
DELETE FROM user_info_001 WHERE user_id = ?;

-- name: AddUserAutoUserId :execresult
INSERT INTO `user_info_001` (
    user_account, user_nickname, user_avatar, 
    user_state, user_mobile, user_gender, 
    user_birthday, user_email, user_is_authentication)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: AddUserHaveUserId :execresult
INSERT INTO `user_info_001` (
    user_id, user_account, user_nickname, 
    user_avatar, user_state, user_mobile, 
    user_gender, user_birthday, user_email, 
    user_is_authentication)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: EditUserByUserId :execresult
UPDATE `user_info_001`
SET user_nickname = ?, user_avatar = ?, user_mobile = ?,
    user_gender = ?, user_birthday = ?, user_email = ?, 
    updated_at = NOW()
WHERE user_id = ? AND user_is_authentication = 1;