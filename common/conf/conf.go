package conf

type Config struct {
	SmtpServer string
	Sender     string
	Name       string
	Passwd     string
	Recipient  string
	CC         []string
	Subject    string
	Msg        string
	EnableHTML bool
}
