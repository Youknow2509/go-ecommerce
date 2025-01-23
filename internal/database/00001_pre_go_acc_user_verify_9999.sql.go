// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 00001_pre_go_acc_user_verify_9999.sql

package database

import (
	"context"
	"database/sql"
)

const getInfoOTP = `-- name: GetInfoOTP :one
SELECT verify_id, verify_otp, verify_key, verify_key_hash, verify_type, is_verified, is_deleted
FROM ` + "`" + `pre_go_acc_user_verify_9999` + "`" + `
WHERE verify_key_hash = ?
`

type GetInfoOTPRow struct {
	VerifyID      int32
	VerifyOtp     string
	VerifyKey     string
	VerifyKeyHash string
	VerifyType    sql.NullInt32
	IsVerified    sql.NullInt32
	IsDeleted     sql.NullInt32
}

func (q *Queries) GetInfoOTP(ctx context.Context, verifyKeyHash string) (GetInfoOTPRow, error) {
	row := q.db.QueryRowContext(ctx, getInfoOTP, verifyKeyHash)
	var i GetInfoOTPRow
	err := row.Scan(
		&i.VerifyID,
		&i.VerifyOtp,
		&i.VerifyKey,
		&i.VerifyKeyHash,
		&i.VerifyType,
		&i.IsVerified,
		&i.IsDeleted,
	)
	return i, err
}

const getValidOtp = `-- name: GetValidOtp :one
SELECT verify_otp, verify_key_hash, verify_key, verify_id
FROM ` + "`" + `pre_go_acc_user_verify_9999` + "`" + `
WHERE verify_key_hash = ? AND is_verified = 0
`

type GetValidOtpRow struct {
	VerifyOtp     string
	VerifyKeyHash string
	VerifyKey     string
	VerifyID      int32
}

func (q *Queries) GetValidOtp(ctx context.Context, verifyKeyHash string) (GetValidOtpRow, error) {
	row := q.db.QueryRowContext(ctx, getValidOtp, verifyKeyHash)
	var i GetValidOtpRow
	err := row.Scan(
		&i.VerifyOtp,
		&i.VerifyKeyHash,
		&i.VerifyKey,
		&i.VerifyID,
	)
	return i, err
}

const insertOTPVerify = `-- name: InsertOTPVerify :execresult
INSERT INTO ` + "`" + `pre_go_acc_user_verify_9999` + "`" + ` (
    verify_otp, 
    verify_key, 
    verify_key_hash, 
    verify_type,
    is_verified, 
    is_deleted,
    verify_created_at,
    verify_updated_at
)
VALUES (?, ?, ?, ?, 0, 0, now(), now())
`

type InsertOTPVerifyParams struct {
	VerifyOtp     string
	VerifyKey     string
	VerifyKeyHash string
	VerifyType    sql.NullInt32
}

func (q *Queries) InsertOTPVerify(ctx context.Context, arg InsertOTPVerifyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertOTPVerify,
		arg.VerifyOtp,
		arg.VerifyKey,
		arg.VerifyKeyHash,
		arg.VerifyType,
	)
}

const updateUserVerificationStatus = `-- name: UpdateUserVerificationStatus :exec
UPDATE ` + "`" + `pre_go_acc_user_verify_9999` + "`" + `
SET is_verified = 1, 
    verify_updated_at = now()
WHERE verify_key_hash = ?
`

// update lai
func (q *Queries) UpdateUserVerificationStatus(ctx context.Context, verifyKeyHash string) error {
	_, err := q.db.ExecContext(ctx, updateUserVerificationStatus, verifyKeyHash)
	return err
}
