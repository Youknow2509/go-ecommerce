package create

import (
	"strings"

	"github.com/Youknow2509/go-ecommerce/internal/utils/sendto"
	"github.com/Youknow2509/go-ecommerce/internal/utils/sendto/sendgrid"
	"github.com/Youknow2509/go-ecommerce/internal/utils/sendto/smtp"
)

// Factory create struct for sending email
func FactoryCreateSendTo(
	typeSendTo string,
) sendto.ISendTo {
	switch strings.ToLower(typeSendTo) {
	case "smtp":
		return smtp.NewSendToWithSMTP()
	case "sendgrid":
		return sendgrid.NewSendToWithSendGrid()
	default:
		return nil
	}
}