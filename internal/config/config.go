package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	SMTPHost     string
	SMTPPort     string
	FromEmail    string
	EmailPassword string
	ToEmail      string
}

// loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	fromEmail := os.Getenv("EMAIL")
	if fromEmail == "" {
		return nil, errors.New("EMAIL environment variable is not set")
	}

	password := os.Getenv("EMAIL_PASSWORD")
	if password == "" {
		return nil, errors.New("EMAIL_PASSWORD environment variable is not set")
	}

	toEmail := os.Getenv("TO_EMAIL")
	if toEmail == "" {
		toEmail = "add@your.own" // Default recipient (add your own)
	}

	// For Gmail
	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		smtpHost = "smtp.gmail.com" // Default SMTP host
	}

	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		smtpPort = "465" // Default SMTP port for SSL
	}

	return &Config{
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		FromEmail:    fromEmail,
		EmailPassword: password,
		ToEmail:      toEmail,
	}, nil
}

// returns the full SMTP server address
func (c *Config) SMTPAddress() string {
	return fmt.Sprintf("%s:%s", c.SMTPHost, c.SMTPPort)
}