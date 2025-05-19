# easysmtp

`easysmtp` is a lightweight and opinionated wrapper around Go's `net/smtp` package. 

It simplifies sending emails via STARTTLS (port 587) by abstracting configuration and offering easy-to-use helpers for Gmail, Outlook, and QQ Foxmail Mail.

# Install

```bash
go get github.com/yonomesh/easysmtp@v0.0.1
```

# Usage

## Send()

`Send()` is the basic function used to send emails.

```go
func Send(config conf.Config) (err error)
```

example
```go
import (
	"log"

	"github.com/yonomesh/easysmtp"
)

func main() {
	config := easysmtp.Config{
		SmtpServer: "smtp.qq.com",             // SMTP server
		Sender:     "xxxxxx@foxmail.com",      // Sender email address
		Name:       "Young Hall",              // Sender name
		Passwd:     "xxxxxxxxxxx",             // App password or SMTP auth password
		Recipient:  "someone@outlook.com",     // Recipient email address
		CC:         nil,                       // Optional CC addresses ([]string)
		Subject:    "Hello EasySmtp",          // Email subject
		Msg:        "This is a test message",  // Email body (plain text or HTML)
		EnableHTML: false,                     // Enable HTML grammar
	}

	if err := easysmtp.Send(config); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
}
```

## Using Environment Variables

You can simplify configuration by using environment variables:

```bash
export EASYSMTP_SERVER="xxxx.smtp.com" # (optional) SMTP server
export EASYSMTP_MAIL="xxxx@mail.com" # sender mail address
export ESSYSMTP_NAME= "xxxx" # your name
export EASYSMTP_PASSWD="xxxx" # App password or SMTP auth password
export EASYSMTP_RECIPIENT_MAIL="xxxx@mail.com" # (optional) recipient mail address
export EASYSMTP_ENABLE_HTML=enable # (optional) enable html grammar, default false
```

### EasySend()

- `EasySend()` need `EASYSMTP_SERVER`, `EASYSMTP_MAIL`, `ESSYSMTP_NAME`, `EASYSMTP_PASSWD`, `EASYSMTP_RECIPIENT_MAIL` 
- It will send mail without Subject and Cc

```go

if err:=easysmtp.EasySend(msg); err != nil {
    log.Fatalf("Failed to send email: %v", err)
}
```
### SendMail()

- `EasySend()` need `EASYSMTP_SERVER`, `EASYSMTP_MAIL`, `ESSYSMTP_NAME`, `EASYSMTP_PASSWD` 
- you need to set recipient, Cc, subject, and msg
```go
if err:=easysmtp.SendMail(recipient,cc,subject,msg); err != nil {
    log.Fatalf("Failed to send email: %v", err)
}
```

### Gmail()
-  `Gmail()` need `EASYSMTP_MAIL`, `ESSYSMTP_NAME`, `EASYSMTP_PASSWD` 
- you need to set recipient, Cc, subject, and msg
```go
if err:=easysmtp.Gmail(recipient,cc,subject,msg); err != nil {
    log.Fatalf("Failed to send email: %v", err)
}
```

### Outlook()
-  `Outlook()` need `EASYSMTP_MAIL`, `ESSYSMTP_NAME`, `EASYSMTP_PASSWD` 
- you need to set recipient, Cc, subject, and msg
```go
if err:=easysmtp.Outlook(recipient,cc,subject,msg); err != nil {
    log.Fatalf("Failed to send email: %v", err)
}
```
### QQFoxmail()
-  `QQFoxmail()` need `EASYSMTP_MAIL`, `ESSYSMTP_NAME`, `EASYSMTP_PASSWD` 
- you need to set recipient, Cc, subject, and msg
```go
if err:=easysmtp.QQFoxmail(recipient,cc,subject,msg); err != nil {
    log.Fatalf("Failed to send email: %v", err)
}
```
