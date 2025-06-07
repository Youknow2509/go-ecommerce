-- name: GetOneUserInfo :one
SELECT user_id, user_account, user_password, user_salt
FROM `user_base_001`
WHERE user_account = ?;

-- name: GetOneUserInfoAdmin :one
SELECT user_id, user_account, user_password, user_salt, user_login_time, user_logout_time, user_login_ip
    , user_created_at, user_updated_at
FROM `user_base_001`
WHERE user_account = ?;

-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM `user_base_001`
WHERE user_account = ?;

-- name: AddUserBase :execresult
INSERT INTO `user_base_001` (
    user_account, user_password, user_salt, user_created_at, user_updated_at
) VALUES (
    ?, ?, ?, NOW(), NOW()
);

-- name: UpdatePassword :exec
UPDATE `user_base_001` 
SET user_password = ? WHERE user_id = ?;

-- name: LoginUserBase :exec
UPDATE `user_base_001`
SET user_login_time = NOW(), user_login_ip = ?
WHERE user_account = ? AND user_password = ?;

-- name: LogoutUserBase :exec
UPDATE `user_base_001`
SET user_logout_time = NOW()
WHERE user_account = ?;