// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

// Account
type GoCrmUser struct {
	// Account ID
	UsrID uint32
	// Email
	UsrEmail string
	// Phone Number
	UsrPhone string
	// Username
	UsrUsername string
	// Password
	UsrPassword string
	// Created Time
	UsrCreatedAt int32
	// Updated Time
	UsrUpdatedAt int32
	// Created IP
	UsrCreateIpAt string
	// Last Login Time
	UsrLastLoginAt int32
	// Last Login IP
	UsrLastLoginIpAt string
	// Login Times
	UsrLoginTimes int32
	// Status
	UsrStatus bool
}