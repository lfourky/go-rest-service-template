package integration

import (
	"fmt"
	"net"
	"time"

	"github.com/lfourky/go-rest-service-template/pkg/service/mail"
	"github.com/mhale/smtpd"
)

type SMTPServer struct {
	server  *smtpd.Server
	address string
}

func NewSMTPServer(cfg mail.Config) *SMTPServer {
	s := &SMTPServer{
		address: fmt.Sprintf(":%d", cfg.SMTPPort),
	}

	s.server = &smtpd.Server{
		Appname:  "smtp_server_integration_test",
		Hostname: cfg.SMTPHost,
		Timeout:  time.Second,
	}

	return s
}

func (s *SMTPServer) SetHandler(handler smtpd.Handler) {
	s.server.Handler = handler
}

func (s *SMTPServer) Run() (func(), error) {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := s.server.Serve(listener); err != nil {
			fmt.Println(err)
		}
	}()

	return func() {
		if err := listener.Close(); err != nil {
			fmt.Println(err)
		}
	}, nil
}
