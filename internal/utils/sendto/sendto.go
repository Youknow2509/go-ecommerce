package sendto

import (
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

	// v.v
}

