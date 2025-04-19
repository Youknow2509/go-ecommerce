package sendgrid

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/utils/sendto"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.uber.org/zap"
)

const (
	SENDER_AUTH_NAME  = "t1"
	SENDER_AUTH_EMAIL = "lytranvinh.work@gmail.com"
)

// define struct send grid
type sendToWithSendGrid struct {
}

// SendKafkaEmailOTP implements sendto.ISendTo.
func (s *sendToWithSendGrid) SendKafkaEmailOTP(to string, from string, otp string) error {
	panic("unimplemented")
}

// implement interface ISendTo to new struct sendToWithSendGrid
func NewSendToWithSendGrid() sendto.ISendTo {
	return &sendToWithSendGrid{}
}

// Send a simple text OTP email
func (s *sendToWithSendGrid) SendTextEmailOTP(
	to []string,
	from string,
	otp string,
) error {
	var c = 0
	for _, email := range to {
		mail := &Mail{
			From: EmailAddress{
				Address: from,
				Name:    "Admin",
			},
			To:               email,
			Subject:          "OTP Verification",
			PlainTextContent: fmt.Sprintf("Your OTP is: %s, Please enter it to verify your account.", otp),
			HtmlContent:      "",
		}
		err := sendMail(*mail)
		if err != nil {
			global.Logger.Error("Email send failed:: ", zap.Error(err))
			c++
		}
	}
	if c > 0 {
		return fmt.Errorf("Email send failed:: %v mail", c)
	}

	return nil
}

// SendAPIEmailOTP implements sendto.ISendTo.
func (s *sendToWithSendGrid) SendAPIEmailOTP(to string, from string, otp string) error {
	panic("unimplemented")
}

// Send a template html OTP email
func (s *sendToWithSendGrid) SendTemplateEmailOTP(
	to []string,
	from string,
	nameTemplate string,
	dataTemplate map[string]interface{},
) error {
	var c = 0

	mailTemplateHtml, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		global.Logger.Error("Error get mail htmp:: ", zap.Error(err))
		return err
	}

	for _, email := range to {
		mail := &Mail{
			From: EmailAddress{
				Address: from,
				Name:    "Admin",
			},
			To:          email,
			Subject:     "OTP Verification",
			HtmlContent: mailTemplateHtml,
		}
		err := sendMail(*mail)
		if err != nil {
			global.Logger.Error("Email send failed:: ", zap.Error(err))
			c++
		}
	}
	if c > 0 {
		return fmt.Errorf("Email send failed:: %v mail", c)
	}

	return nil
}

// help send
func sendMail(m Mail) error {
	message := BuildMessageInSendGird(m)
	client := sendgrid.NewSendClient(global.Config.SendGrid.APIKey)

	response, err := client.Send(message)
	if err != nil {
		global.Logger.Error("Email send failed:: ", zap.Error(err))
		return err
	}

	// check status code
	if response.StatusCode != 202 {
		global.Logger.Error("Email send failed:: ", zap.Any("response", response))
		return fmt.Errorf("Email send failed:: %v", response)
	}

	return nil
}

// Help get template email html
func getMailTemplate(
	nameTemplate string,
	dataTemplate map[string]interface{},
) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t := template.Must(
		template.New(
			nameTemplate).ParseFiles(
			"templates-email/" + nameTemplate))

	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

// Build message in SendGrid
func BuildMessageInSendGird(m Mail) *mail.SGMailV3 {
	from := mail.NewEmail(SENDER_AUTH_NAME, SENDER_AUTH_EMAIL)

	subject := m.Subject
	to := mail.NewEmail("", m.To)
	plainTextContent := m.PlainTextContent
	htmlContent := m.HtmlContent

	return mail.NewSingleEmail(
		from,
		subject,
		to,
		plainTextContent,
		htmlContent)
}

// Send a simple text email
type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

// Send a simple text email
type Mail struct {
	From             EmailAddress `json:"from"`
	To               string       `json:"to"`
	Subject          string       `json:"subject"`
	Body             string       `json:"body"`
	PlainTextContent string       `json:"plainTextContent"`
	HtmlContent      string       `json:"htmlContent"`
}
