package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
)

type Sender struct{}

func (s *Sender) Send(content string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	from := os.Getenv("SMTP_USERNAME")
	to := os.Getenv("SMTP_USERNAME")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	message := []byte("To: " + to + "\r\n" +
		"Subject: Contact form\r\n" +
		"\r\n" +
		content + "\r\n")

	auth := smtp.PlainAuth("", username, password, host)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", host+":"+port, tlsconfig)

	if err != nil {
		return fmt.Errorf("dialing tls: %w", err)
	}

	client, err := smtp.NewClient(conn, host)

	if err != nil {
		return fmt.Errorf("creating new smtp client: %w", err)
	}

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("authenticating: %w", err)
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("setting from: %w", err)
	}

	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("setting to: %w", err)
	}

	w, err := client.Data()

	if err != nil {
		return fmt.Errorf("get client data: %w", err)
	}

	if _, err = w.Write(message); err != nil {
		return fmt.Errorf("write message: %w", err)
	}

	if err = w.Close(); err != nil {
		return fmt.Errorf("close write: %w", err)
	}

	return nil
}

func NewSender() *Sender {
	return &Sender{}
}
