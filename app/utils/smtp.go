package utils

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

type EmailBuilder struct {
	message *gomail.Message
}

type MailContent struct {
	Subject string
	Html    string
}

func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{message: gomail.NewMessage()}
}

func (b *EmailBuilder) To(to []string) *EmailBuilder {
	b.message.SetHeader("To", to...)
	return b
}

func (b *EmailBuilder) Content(content MailContent) *EmailBuilder {
	senderName := os.Getenv("MAIL_FROM_NAME")
	sender := os.Getenv("MAIL_FROM_ADDRESS")

	b.message.SetHeader("From", fmt.Sprintf("%s <%s>", senderName, sender))
	b.message.SetHeader("Subject", content.Subject)
	b.message.SetBody("text/html", content.Html)
	b.message.SetHeader("Message-ID", fmt.Sprintf("%v", time.Now().Unix()))

	return b
}

func (b *EmailBuilder) Attach(filename string, settings ...gomail.FileSetting) *EmailBuilder {
	b.message.Attach(filename, settings...)
	return b
}

func (b *EmailBuilder) Send() error {
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")

	// conv
	convSmtpPort, _ := strconv.Atoi(smtpPort)

	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	dialer := gomail.NewDialer(smtpHost, convSmtpPort, username, password)

	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	if err := dialer.DialAndSend(b.message); err != nil {
		return err
	}

	return nil
}
