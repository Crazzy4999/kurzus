package email

import (
	"errors"
	config "hangryAPI/configs"
	"net/smtp"
	"os"
	"regexp"
)

type EmailService struct {
	cfg  *config.Config
	host string
	port string
}

func NewEmailService(cfg *config.Config, host, port string) *EmailService {
	return &EmailService{
		cfg:  cfg,
		host: host,
		port: port,
	}
}

func (s *EmailService) SendEmail(to []string, subject string, mime string, templatePath string, templateData string) error {
	from := s.cfg.EmailingAccountEmail
	password := s.cfg.EmailingAccountPassword

	bytes, err := os.ReadFile(templatePath)
	if err != nil {
		return errors.New("couldn't read in html template")
	}

	processed := regexp.MustCompile("{{.}}").ReplaceAll(bytes, []byte(templateData))

	mimedBytes := []byte(subject + mime + string(processed))

	auth := smtp.PlainAuth("", from, password, s.host)

	err = smtp.SendMail(s.host+":"+s.port, auth, from, to, mimedBytes)
	if err != nil {
		return errors.New("sending email failed")
	}

	return nil
}

func (s *EmailService) SendResetEmail(to string, subject string, mime string, templatePath string, resetKey string) error {
	toSlice := []string{to}
	err := s.SendEmail(toSlice, subject, mime, templatePath, resetKey)
	if err != nil {
		return errors.New("sending reset email failed")
	}

	return nil
}
