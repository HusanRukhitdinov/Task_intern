package email

import (
	"fmt"
	"net/smtp"
)

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
	FromName     string
}

type EmailSender struct {
	config EmailConfig
}

func NewEmailSender(config EmailConfig) *EmailSender {
	return &EmailSender{
		config: config,
	}
}

func (e *EmailSender) SendOTP(toEmail, code string) error {
	subject := "Your OTP Code"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Your OTP Code</h2>
			<p>Your OTP code is: <strong>%s</strong></p>
			<p>This code will expire in 3 minutes.</p>
			<p>If you didn't request this code, please ignore this email.</p>
		</body>
		</html>
	`, code)

	return e.SendEmail(toEmail, subject, body)
}

func (e *EmailSender) SendEmail(to, subject, body string) error {
	from := e.config.FromEmail

	auth := smtp.PlainAuth("", e.config.SMTPUsername, e.config.SMTPPassword, e.config.SMTPHost)

	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", e.config.FromName, from)
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	addr := fmt.Sprintf("%s:%s", e.config.SMTPHost, e.config.SMTPPort)
	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
