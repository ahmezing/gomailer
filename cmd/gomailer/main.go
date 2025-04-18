package main

import (
	"gomailer/internal/config"
	"gomailer/internal/email"
	"gomailer/internal/templates"
	"log"
)

func main() {
	// load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// load email template
	templateLoader := templates.NewTemplateLoader()
	emailTemplate, err := templateLoader.LoadTemplate("templates/email.html")
	if err != nil {
		log.Fatalf("Failed to load email template: %v", err)
	}

	// create email client
	client, err := email.NewSMTPClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create SMTP client: %v", err)
	}

	// create email message
	message := email.NewMessage(
		cfg.FromEmail,
		[]string{cfg.ToEmail},
		"بلا عنوان",
		emailTemplate,
	)

	// send email
	if err := client.Send(message); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}