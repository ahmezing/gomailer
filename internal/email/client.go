package email

import (
	"crypto/tls"
	"gomailer/internal/config"
	"net/smtp"
)

type SMTPClient struct {
	config *config.Config
}

func NewSMTPClient(cfg *config.Config) (*SMTPClient, error) {
	return &SMTPClient{
		config: cfg,
	}, nil
}

func (c *SMTPClient) Send(msg *Message) error {
	content := msg.Build()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         c.config.SMTPHost,
	}

	conn, err := tls.Dial("tcp", c.config.SMTPAddress(), tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, c.config.SMTPHost)
	if err != nil {
		return err
	}
	defer client.Quit()

	auth := smtp.PlainAuth("", c.config.FromEmail, c.config.EmailPassword, c.config.SMTPHost)
	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(c.config.FromEmail); err != nil {
		return err
	}

	for _, recipient := range msg.To {
		if err = client.Rcpt(recipient); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(content)
	return err
}