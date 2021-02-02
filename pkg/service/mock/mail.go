package mock

import (
	"github.com/stretchr/testify/mock"
)

type MailSender struct {
	mock.Mock
}

func (m *MailSender) SendMail(recipient, subject, body string) error {
	args := m.Called(recipient, subject, body)

	return args.Error(0)
}
