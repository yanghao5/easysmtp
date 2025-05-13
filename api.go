package easysmtp

func EasySend(msg string) (err error) {
	_ = msg
	return nil
}
func SendMail(recipient string, cc []string, subject string, msg string) (err error) {
	_ = recipient
	_ = cc
	_ = subject
	_ = msg

	return nil
}
func Gmail(recipient string, cc []string, subject string, msg string) (err error) {
	_ = recipient
	_ = cc
	_ = subject
	_ = msg

	return nil
}
func Outlook(recipient string, cc []string, msg string, subject string) (err error) {
	_ = recipient
	_ = cc
	_ = subject
	_ = msg

	return nil
}
func QQFoxmail(recipient string, cc []string, msg string, subject string) (err error) {
	_ = recipient
	_ = cc
	_ = subject
	_ = msg

	return nil
}

func Send(smtp_server string, sender string, name string, passwd string, recipient string, subject string, msg string, cc []string, enable_html bool) (ret string, err error) {
	_ = smtp_server
	_ = sender
	_ = name
	_ = passwd
	_ = recipient
	_ = subject
	_ = msg
	_ = cc
	_ = enable_html
	return "", nil
}
