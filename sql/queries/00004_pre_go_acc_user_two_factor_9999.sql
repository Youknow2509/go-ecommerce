
-- name: EnableTwoFactorTypeEmail :exec
INSERT INTO `pre_go_acc_user_two_factor_9999` (user_id, two_factor_auth_type, two_factor_email, two_factor_auth_secret, two_factor_is_active, two_factor_created_at, two_factor_updated_at)
VALUES (?, ?, ?, "OTP", FALSE, NOW(), NOW());

-- name: DisableTwoFactor :exec
UPDATE `pre_go_acc_user_two_factor_9999`
SET two_factor_is_active = FALSE, 
    two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ?;

-- name: UpdateTwoFactorStatus :exec
UPDATE `pre_go_acc_user_two_factor_9999`
SET two_factor_is_active = TRUE, two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ? AND two_factor_is_active = FALSE;

-- name: VerifyTwoFactor :one
SELECt count(*) 
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_auth_type = ? AND two_factor_is_active = TRUE;

-- name: GetTwoFactorStatus :one
SELECT two_factor_is_active
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_auth_type = ?;

-- name: IsTwoFactorEnabled :one
SELECT count(*)
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_is_active = TRUE;

-- name: AddOrUpdateEmail :exec
INSERT INTO `pre_go_acc_user_two_factor_9999` (
    user_id, two_factor_email, two_factor_is_active)
VALUES (?, ?, TRUE)
ON DUPLICATE KEY UPDATE 
    two_factor_email = ?, two_factor_updated_at = NOW();

-- name: GetUserTwoFactoryMethods :many
SELECT user_id, two_factor_auth_type, two_factor_auth_secret, 
    two_factor_email, two_factor_phone,
    two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ?;

-- name: ReactivateTwoFactor :exec
UPDATE `pre_go_acc_user_two_factor_9999`
SET two_factor_is_active = TRUE, two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ?;

-- name: RemoveTwoFactor :exec
DELETE FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_auth_type = ?;

-- CountActiveTwoFactorMethods :one
SELECT COUNT(*)
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_is_active = TRUE;

-- name: GetTwoFactorMethodByID :one
SELECT user_id, two_factor_auth_type, two_factor_auth_secret, 
    two_factor_email, two_factor_phone,
    two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ?;

-- name: GetTwoFactorMethodByIDAndByType :one
SELECT user_id, two_factor_auth_type, two_factor_auth_secret, 
    two_factor_email, two_factor_phone,
    two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_auth_type = ?;
