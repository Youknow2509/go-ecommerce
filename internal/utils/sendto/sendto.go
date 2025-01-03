package sendto

import (
)

// Enum Type for sending
const (
    TYPE_SMTP = "smtp"
	TYPE_SENDGRID = "sendgrid"
	TYPE_API = "api"
)

// interface for sending email
type ISendTo interface {
	// Send a simple text OTP email
	SendTextEmailOTP(to []string, from string, otp string) error

	// Send a template html OTP email
	SendTemplateEmailOTP(
		to []string,
		from string,
		nameTemplate string,
		dataTemplate map[string]interface{},
	) error
	
	// Send api OTP email
	SendAPIEmailOTP(to string, from string, otp string) error

	// v.v
}

