package sendto

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	// "net/smtp"
	"strings"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.uber.org/zap"
)

const (
	SMTP_HOST     = "smtp.gmail.com"
	SMTP_PORT     = "587"
	SMTP_USERNAME = "api"
	SMTP_PASSWORD = ""
)

const (
	SENDGRID_API_KEY = ""
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\"; \r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

func BuildMessageInSendGird(m Mail) *mail.SGMailV3{
	from := mail.NewEmail("t1", "lytranvinh.work@gmail.com")
    subject := m.Subject
    to := mail.NewEmail("", m.To[0])
    plainTextContent := ""
    htmlContent := m.Body
	return mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
}

func SendTextEmailOTP(to []string, from string, otp string) error {
	content := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Admin",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is: %s, Please enter it to verify your account.", otp),
	}

	// messageMail := BuildMessage(content)

	// // send smtp message
	// auth := smtp.PlainAuth("", SMTP_USERNAME, SMTP_PASSWORD, SMTP_HOST)

	// err := smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, from, to, []byte(messageMail))
	// if err != nil {
	// 	global.Logger.Error("Email send failed:: ", zap.Error(err))
	// 	return err
	// }
	mail_send := BuildMessageInSendGird(content)
	client := sendgrid.NewSendClient(SENDGRID_API_KEY)
	response, err := client.Send(mail_send)

	fmt.Println(0)
	if err != nil {
		global.Logger.Error("Email send failed:: ", zap.Error(err))
		fmt.Println("Email send failed:: ", err)

	fmt.Println(2)

		return err
	}
	if (response.StatusCode != 202) {
		fmt.Println("Email send failed:: ", response)

	fmt.Println(2)

		return errors.New("Email send failed")
	}
	global.Logger.Info("Email OTP send successful:: ", zap.Any("response", response))
	fmt.Println(3)


	return nil
}

func SendTemplateEmailOTP(
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

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates-email/" + nameTemplate + ".html"))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplateTemplate string) error {
	content := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Admin",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplateTemplate,
	}

	// messageMail := BuildMessage(content)

	// // send smtp message
	// auth := smtp.PlainAuth("", SMTP_USERNAME, SMTP_PASSWORD, SMTP_HOST)

	// err := smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, from, to, []byte(messageMail))
	// if err != nil {
	// 	global.Logger.Error("Email send failed:: ", zap.Error(err))
	// 	return err
	// }

	fmt.Println("Email sent ", content)

	return nil
}
