package smtp

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/utils/sendto"
	"go.uber.org/zap"
)

const (
	SMTP_HOST     = "smtp.gmail.com"
	SMTP_PORT     = "587"
	SMTP_USERNAME = "api"
	SMTP_PASSWORD = ""
)

// define struct sendToWithSMTP
type sendToWithSMTP struct {
}

// SendKafkaEmailOTP implements sendto.ISendTo.
func (s *sendToWithSMTP) SendKafkaEmailOTP(to string, from string, otp string) error {
	panic("unimplemented")
}

// implement interface ISendTo to new struct sendToWithSMTP
func NewSendToWithSMTP() sendto.ISendTo {
	return &sendToWithSMTP{}
}

// SendAPIEmailOTP implements sendto.ISendTo.
func (s *sendToWithSMTP) SendAPIEmailOTP(to string, from string, otp string) error {
	panic("unimplemented")
}

// Send a simple text OTP email
func (s *sendToWithSMTP) SendTextEmailOTP(
	to []string,
	from string,
	otp string,
) error {
	content := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Admin",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is: %s, Please enter it to verify your account.", otp),
	}

	messageMail := BuildMessage(content)

	// send smtp message
	auth := smtp.PlainAuth("", SMTP_USERNAME, SMTP_PASSWORD, SMTP_HOST)

	err := smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send failed:: ", zap.Error(err))
		return err
	}

	return nil
}

// Send a template html OTP email
func (s *sendToWithSMTP) SendTemplateEmailOTP(
	to []string,
	from string,
	nameTemplate string,
	dataTemplate map[string]interface{},
) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)

	if err != nil {
		return err
	}

	return send(to, from, htmlBody)
}

// help get template html
func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t := template.Must(
		template.New(
			nameTemplate).ParseFiles(
			"templates-email/" + nameTemplate + ".html"))

	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

// help sends
func send(
	to []string,
	from string,
	htmlTemplateTemplate string,
) error {
	content := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Admin",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplateTemplate,
	}

	messageMail := BuildMessage(content)

	// send smtp message
	auth := smtp.PlainAuth("", SMTP_USERNAME, SMTP_PASSWORD, SMTP_HOST)

	err := smtp.SendMail(
		SMTP_HOST+":"+SMTP_PORT,
		auth,
		from,
		to,
		[]byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send failed:: ", zap.Error(err))
		return err
	}

	return nil
}

// define struct EmailAddress
type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

// define struct Mail
type Mail struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

// helpper build message
func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\"; \r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}
