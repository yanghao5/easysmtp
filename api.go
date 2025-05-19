package easysmtp

import (
	"easysmtp/common/conf"
	"easysmtp/common/validator"
	"easysmtp/internal"
	"fmt"
	"os"
	"strings"

	netsmtp "net/smtp"
)

func EasySend(msg string) (err error) {
	config := conf.Config{
		SmtpServer: os.Getenv("EASYSMTP_SERVER"),
		Sender:     os.Getenv("EASYSMTP_MAIL"),
		Name:       os.Getenv("EASYSMTP_NAME"),
		Passwd:     os.Getenv("EASYSMTP_PASSWD"),
		Recipient:  os.Getenv("EASYSMTP_RECIPIENT_MAIL"),
		CC:         []string{},
		Subject:    "",
		Msg:        msg,
		EnableHTML: os.Getenv("EASYSMTP_ENABLE_HTML") == "true",
	}

	if config.SmtpServer == "" || config.Sender == "" || config.Passwd == "" || config.Recipient == "" {
		return fmt.Errorf("missing required SMTP configuration environment variables")
	}
	if !validator.IsValidEmail(config.Sender) {
		return fmt.Errorf("sender mail address is error")
	}

	if !validator.IsValidEmail(config.Recipient) {
		return fmt.Errorf("recipient mail address is error")
	}

	Send(config)

	return nil
}
func SendMail(recipient string, cc []string, subject string, msg string) (err error) {

	config := conf.Config{
		SmtpServer: os.Getenv("EASYSMTP_SERVER"),
		Sender:     os.Getenv("EASYSMTP_MAIL"),
		Name:       os.Getenv("EASYSMTP_NAME"),
		Passwd:     os.Getenv("EASYSMTP_PASSWD"),
		Recipient:  recipient,
		CC:         cc,
		Subject:    subject,
		Msg:        msg,
		EnableHTML: os.Getenv("EASYSMTP_ENABLE_HTML") == "true",
	}

	if config.SmtpServer == "" || config.Name == "" || config.Sender == "" || config.Passwd == "" || config.Recipient == "" {
		return fmt.Errorf("missing required SMTP configuration environment variables")
	}

	if !validator.IsValidEmail(config.Sender) {
		return fmt.Errorf("sender mail address is error")
	}

	if !validator.IsValidEmail(config.Recipient) {
		return fmt.Errorf("recipient mail address is error")
	}

	for _, addr := range config.CC {
		if !validator.IsValidEmail(addr) {
			return fmt.Errorf("error Cc mail address")
		}
	}

	Send(config)

	return nil
}

func Gmail(recipient string, cc []string, subject string, msg string) (err error) {
	config := conf.Config{
		SmtpServer: "smtp.gmail.com",
		Sender:     os.Getenv("EASYSMTP_MAIL"),
		Name:       os.Getenv("EASYSMTP_NAME"),
		Passwd:     os.Getenv("EASYSMTP_PASSWD"),
		Recipient:  recipient,
		CC:         cc,
		Subject:    subject,
		Msg:        msg,
		EnableHTML: os.Getenv("EASYSMTP_ENABLE_HTML") == "true",
	}

	if config.Sender == "" || config.Name == "" || config.Passwd == "" || config.Recipient == "" {
		return fmt.Errorf("missing required SMTP configuration environment variables")
	}
	if !validator.IsValidEmail(config.Sender) {
		return fmt.Errorf("sender mail error")
	}

	if !validator.IsEmailFromProvider(config.Sender, "gmail") {
		return fmt.Errorf("sender mail is not gmail")
	}

	if !validator.IsValidEmail(config.Recipient) {
		return fmt.Errorf("recipient mail address is error")
	}

	for _, addr := range config.CC {
		if !validator.IsValidEmail(addr) {
			return fmt.Errorf("error Cc mail address")
		}
	}

	Send(config)

	return nil
}

func Send(config conf.Config) (ret string, err error) {
	// sender
	sender := ""
	if validator.IsValidEmail(config.Sender) {
		sender = config.Sender
	} else {
		return "", fmt.Errorf("invalid sender email: %s", config.Sender)
	}

	// recipient
	recipient := ""
	if validator.IsValidEmail(config.Recipient) {
		recipient = config.Recipient
	} else {
		return "", fmt.Errorf("invalid recipient email: %s", config.Recipient)
	}

	// Cc
	validCC := []string{}
	for _, addr := range config.CC {
		if validator.IsValidEmail(addr) {
			validCC = append(validCC, addr)
		}
	}

	// 构造邮件头
	header := map[string]string{
		"From":    fmt.Sprintf("%s <%s>", config.Name, sender),
		"To":      recipient,
		"Subject": config.Subject,
	}

	if len(validCC) > 0 {
		header["Cc"] = strings.Join(validCC, ", ")
	}
	if config.EnableHTML {
		header["Content-Type"] = "text/html; charset=UTF-8"
	} else {
		header["Content-Type"] = "text/plain; charset=UTF-8"
	}

	var message strings.Builder
	for k, v := range header {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	message.WriteString("\r\n" + config.Msg)

	allRecipients := append([]string{config.Recipient}, validCC...)

	auth := netsmtp.PlainAuth("", config.Sender, config.Passwd, config.SmtpServer)

	err = internal.SendMailUsingSTARTTLS(
		fmt.Sprintf("%s:%d", config.SmtpServer, 587),
		auth,
		config.Sender,
		allRecipients,
		[]byte(message.String()),
	)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Send mail success!")
	}
	return "", nil
}
