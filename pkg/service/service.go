package service

import "time"

type Clock interface {
	Now() time.Time
}

type MailSender interface {
	SendMail(recipient, subject, body string) error
}
