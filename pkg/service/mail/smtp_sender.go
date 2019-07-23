package mail

import (
	gomail "gopkg.in/gomail.v2"
)

type SMTPSender struct {
	dialer        *gomail.Dialer
	senderName    string
	senderAddress string
}

func NewSMTPSender(
	smtpPort uint,
	smtpHost,
	smtpUser,
	smtpPassword,
	senderName,
	senderAddress string,
) *SMTPSender {
	return &SMTPSender{
		dialer:        gomail.NewDialer(smtpHost, int(smtpPort), smtpUser, smtpPassword),
		senderName:    senderName,
		senderAddress: senderAddress,
	}
}

func (s *SMTPSender) SendMail(recipientAddress, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetBody("text/html", body)
	msg.SetHeaders(map[string][]string{
		"From":    {msg.FormatAddress(s.senderAddress, s.senderName)},
		"To":      {recipientAddress},
		"Subject": {subject},
	})

	return s.dialer.DialAndSend(msg)
}
