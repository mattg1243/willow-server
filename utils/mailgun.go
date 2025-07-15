package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type SendEmailParams struct {
	Subject string
	Content string
	To      string
}

func SendEmail(params SendEmailParams) (string, error) {
	var apiKey = os.Getenv("MAILGUN_API_KEY")
	var domain = os.Getenv("MAILGUN_DOMAIN")
	// var address = os.Getenv("MAILGUN_SUPPORT_ADDRESS")

	var mg = mailgun.NewMailgun(domain, apiKey)

	// mg.SetAPIBase("https://api.eu.mailgun.net")
	m := mailgun.NewMessage(
		"Willow Support <support@willowapp.io>",
		params.Subject,
		"",
		params.To,
	)

	m.SetHTML(params.Content)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)

	return id, err
}

func CreateResetPasswordEmail(token string) string {
	baseURL := os.Getenv("CLIENT_HOST")

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", baseURL, token)

	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Reset Your Password</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px; text-align: center;">
			<div style="max-width: 600px; margin: auto; background-color: #ffffff; padding: 30px; border-radius: 8px; box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);">
				<h2 style="color: #333;">Reset Your Willow Password</h2>
				<p>Hello,</p>
				<p>We received a request to reset your password. Click the button below to set a new one:</p>
				<p style="text-align: center;">
					<a href="%s" style="display: inline-block; padding: 12px 24px; background-color: #03b126; color: white; text-decoration: none; border-radius: 4px;">Reset Password</a>
				</p>
				<p>If you didn't request a password reset, you can safely ignore this email.</p>
				<p style="color: #999; font-size: 12px;">This link will expire in 1 hour for your security.</p>
			</div>
		</body>
		</html>
	`, resetLink)

	return html
}
