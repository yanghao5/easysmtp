package easysmtp

type Config struct {
	SmtpServer string   // SMTP server
	Sender     string   // Sender email address
	Name       string   // Sender name
	Passwd     string   // App password or SMTP auth password
	Recipient  string   // Recipient email address
	CC         []string // Cc addresses
	Subject    string   // Email subject
	Msg        string   // Email body (plain text or HTML)
	EnableHTML bool     // Enable HTML grammar
}
