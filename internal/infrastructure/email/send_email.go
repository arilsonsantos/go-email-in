package email

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail() error {
	smtpHost := os.Getenv("EMAIL_SMTP_HOST")
	smtpPort := os.Getenv("EMAIL_SMTP_PORT")
	email := os.Getenv("EMAIL")
	smtpUser := os.Getenv("SMTP_USER")
	password := os.Getenv("EMAIL_SMTP_PASSWORD")

	err := sendEmail(email, smtpUser, password, smtpHost, smtpPort)
	if err != nil {
		return err
	}

	fmt.Println("E-mail enviado com sucesso!")

	return nil
}

func sendEmail(email string, smtpUser string, password string, smtpHost string, smtpPort string) error {
	// Crie a mensagem de e-mail
	to := []string{"arilsonsantos@gmail.com"}
	subject := "Teste de e-mail"
	body := "Olá,\n\nEste é um e-mail de teste enviado via Go."

	message := "From: " + email + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Autenticação com o servidor SMTP
	auth := smtp.PlainAuth("", smtpUser, password, smtpHost)

	// Envio do e-mail
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, to, []byte(message))
	if err != nil {
		fmt.Println("Erro ao enviar e-mail:", err)
		return err
	}
	return nil
}
