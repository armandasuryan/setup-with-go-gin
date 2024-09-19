package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SentEmail(headerFrom string, headerTo string, subject string, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", headerFrom)
	m.SetHeader("To", headerTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// smtp config
	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	user := os.Getenv("SMTP_USER")
	passwd := os.Getenv("SMTP_PASS")

	configEmail := gomail.NewDialer(host, port, user, passwd)

	if err := configEmail.DialAndSend(m); err != nil {
		fmt.Printf("Failed to sent email : %s", err)
		return err
	}

	fmt.Println("Email sent successfully!")

	return nil
}
