package main

import (
	"github.com/yonomesh/easysmtp"
	"github.com/yonomesh/easysmtp/common/conf"
)

// It is a test
func main() {
	config := conf.Config{
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
