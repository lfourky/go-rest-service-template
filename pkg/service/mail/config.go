package mail

type Config struct {
	SMTPPort      uint
	SMTPHost      string
	SMTPUser      string
	SMTPPassword  string
	SenderName    string
	SenderAddress string
}
