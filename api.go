package easysmtp

import (
	"easysmtp/common/conf"
	"easysmtp/common/validator"
	"easysmtp/internal"
	"fmt"
	"strings"

	netsmtp "net/smtp"
)

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
