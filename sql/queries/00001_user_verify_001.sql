-- name: GetValidOtp :one
SELECT verify_otp, verify_key_hash, verify_key, verify_id
FROM `user_verify_001`
WHERE verify_key_hash = ? AND is_verified = 0;

-- update lai
-- name: UpdateUserVerificationStatus :exec
UPDATE `user_verify_001`
SET is_verified = 1, 
    verify_updated_at = now()
WHERE verify_key_hash = ?;

-- name: InsertOTPVerify :execresult
INSERT INTO `user_verify_001` (
    verify_otp, 
    verify_key, 
    verify_key_hash, 
    verify_type,
    is_verified, 
    is_deleted,
    verify_created_at,
    verify_updated_at
)
VALUES (?, ?, ?, ?, 0, 0, now(), now());

-- name: GetInfoOTP :one
SELECT verify_id, verify_otp, verify_key, verify_key_hash, verify_type, is_verified, is_deleted
FROM `user_verify_001`
WHERE verify_key_hash = ?;