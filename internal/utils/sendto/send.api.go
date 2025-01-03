package sendto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Youknow2509/go-ecommerce/global"
	"go.uber.org/zap"
)

type sendToWithApi struct {
}

// SendAPIEmailOTP implements ISendTo.
func (s *sendToWithApi) SendAPIEmailOTP(to string, from string, otp string) error {
	// url api endpoint
	postURL := "http://localhost:9080/api/v1/email/otp"

	// data json
	mailRequest := Mail{
		To: to,
		From: EmailAddress{
			Address: from,
			Name:    "Admin",
		},
		Subject:          "OTP Verification",
		PlainTextContent: fmt.Sprintf("Your OTP is: %s, Please enter it to verify your account.", otp),
		HtmlContent:      "/templates-email/otp.html",
	}

	// convert to json
	requestBody, err := json.Marshal(mailRequest)
	if err != nil {
		global.Logger.Error("Error convert mailRequest to json:: ", zap.Error(err))
		return err
	}

	// create request api
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(requestBody))
	if err != nil {
		global.Logger.Error("Error create request:: ", zap.Error(err))
		return err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

	// execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.Logger.Error("Error execute request:: ", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	fmt.Sprintln("response Status:", resp.Status)
	return nil
}

// SendTemplateEmailOTP implements ISendTo.
func (s *sendToWithApi) SendTemplateEmailOTP(to []string, from string, nameTemplate string, dataTemplate map[string]interface{}) error {
	panic("unimplemented")
}

// SendTextEmailOTP implements ISendTo.
func (s *sendToWithApi) SendTextEmailOTP(to []string, from string, otp string) error {
	panic("unimplemented")
}

func NewSendToWithApi() ISendTo {
	return &sendToWithApi{}
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