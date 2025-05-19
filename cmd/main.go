package main

import (
	"github.com/yonomesh/easysmtp"
)

// It is a test
func main() {
	config := easysmtp.Config{
		SmtpServer: "",
		Sender:     "",
		Name:       "",
		Passwd:     "",
		Recipient:  "",
		CC:         []string{"", ""},
		Subject:    "Hello EasySmpt",
		Msg:        "It is a test.",
	}
	easysmtp.Send(config)
}
