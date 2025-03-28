package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/logger"
	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	"github.com/go-gomail/gomail"
	"go.uber.org/zap"
)

const (
	SMTPHost = "smtp.gmail.com"
	SMTPPort = 587
	SMTPUser = "d7234846@gmail.com"
	SMTPPass = "gksw mnjw xula fwqw"
)

func SendVerificationEmail(toEmail, verificationLink string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", SMTPUser)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Confirme seu e-mail")
	m.SetBody("text/html", fmt.Sprintf(`
		<h2>Confirmação de Cadastro</h2>
		<p>Clique no link abaixo para verificar seu e-mail:</p>
		<a href="%s">Verificar E-mail</a>
	`, verificationLink))

	d := gomail.NewDialer(SMTPHost, SMTPPort, SMTPUser, SMTPPass)

	// TODO fazer implementações assíncronas
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func VerifyEmail(token string) *rest_err.RestErr {
	mu.Lock()
	email, exists := verificationTokens[token]
	if !exists {
		mu.Unlock()
		return rest_err.NewBadRequestError("Invalid or expired verification token")
	}
	delete(verificationTokens, token)
	verifiedEmails[email] = true
	mu.Unlock()

	logger.Info("Email verified successfully", zap.String("email", email))
	return nil
}

func generateVerificationToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func storeVerificationToken(email, token string) {
	mu.Lock()
	defer mu.Unlock()
	verificationTokens[token] = email
}

