package email

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/config"
)

type EmailService interface {
	SendVerificationEmail(toEmail string, token string) error
	SendPasswordResetEmail(toEmail string, token string) error
}

type emailService struct {
	host     string
	port     string
	user     string
	password string
	from     string
}

func NewEmailService() EmailService {
	return &emailService{
		host:     config.GetEnv("SMTP_HOST"),
		port:     config.GetEnv("SMTP_PORT"),
		user:     config.GetEnv("SMTP_USER"),
		password: config.GetEnv("SMTP_PASS"),
		from:     config.GetEnv("SMTP_USER"),
	}
}

func (s *emailService) sendMail(to []string, subject string, body string) error {
	auth := smtp.PlainAuth("", s.user, s.password, s.host)

	// We need to add necessary headers for HTML email
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte("From: " + s.from + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		mime + "\r\n" +
		body)

	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	return smtp.SendMail(addr, auth, s.from, to, msg)
}

func (s *emailService) SendVerificationEmail(toEmail string, token string) error {
	frontendURL := config.GetEnv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	verifyLink := fmt.Sprintf("%s/verify-email?token=%s", frontendURL, token)

	subject := "Nihongo Learning Journal - Email Verification"

	// Using a simple HTML template
	htmlBody := fmt.Sprintf(`
		<html>
		<body>
			<h2>Welcome to Nihongo Learning Journal!</h2>
			<p>Please click the link below to verify your email address:</p>
			<p><a href="%s" style="padding: 10px 20px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px;">Verify Email</a></p>
			<p>If the button doesn't work, you can copy and paste this link into your browser:</p>
			<p>%s</p>
			<br>
			<p>Arigatou gozaimasu!</p>
		</body>
		</html>
	`, verifyLink, verifyLink)

	return s.sendMail([]string{toEmail}, subject, htmlBody)
}

func (s *emailService) SendPasswordResetEmail(toEmail string, token string) error {
	frontendURL := config.GetEnv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", frontendURL, token)

	subject := "Nihongo Learning Journal - Password Reset Request"

	htmlBody := fmt.Sprintf(`
		<html>
		<body>
			<h2>Password Reset Request</h2>
			<p>We received a request to reset your password. Please click the link below to set a new password:</p>
			<p><a href="%s" style="padding: 10px 20px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px;">Reset Password</a></p>
			<p>If the button doesn't work, you can copy and paste this link into your browser:</p>
			<p>%s</p>
			<br>
			<p>If you didn't request this, you can safely ignore this email.</p>
		</body>
		</html>
	`, resetLink, resetLink)

	return s.sendMail([]string{toEmail}, subject, htmlBody)
}
