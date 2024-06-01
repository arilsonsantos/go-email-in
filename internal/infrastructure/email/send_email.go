package email

import (
	"emailn/internal/domain/campaign"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func SendEmail(campaign *campaign.Campaign) error {
	smtpHost := os.Getenv("EMAIL_SMTP_HOST")
	smtpPort := os.Getenv("EMAIL_SMTP_PORT")
	email := os.Getenv("EMAIL")
	smtpUser := os.Getenv("SMTP_USER")
	password := os.Getenv("EMAIL_SMTP_PASSWORD")

	var emailsTo []string
	for _, contact := range campaign.Contacts {
		emailsTo = append(emailsTo, contact.Email)
	}

	err := sendEmail(campaign, email, smtpUser, emailsTo, password, smtpHost, smtpPort)
	if err != nil {
		return err
	}

	fmt.Println("E-mail enviado com sucesso!")

	return nil
}

func sendEmail(campaign *campaign.Campaign, email string, smtpUser string,
	emailsTo []string, password string, smtpHost string, smtpPort string) error {
	message := "From: " + email + "\n" +
		"To: " + strings.Join(emailsTo, ",") + "\n" +
		"Subject: " + campaign.Name + "\n\n" +
		campaign.Content

	// Autenticação com o servidor SMTP
	auth := smtp.PlainAuth("", smtpUser, password, smtpHost)

	// Envio do e-mail
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, emailsTo, []byte(message))
	if err != nil {
		fmt.Println("Erro ao enviar e-mail:", err)
		return err
	}
	return nil
}
