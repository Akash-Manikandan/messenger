package user

import (
	"context"

	"github.com/Akash-Manikandan/app-backend/pkg/queue"
	"github.com/gofiber/fiber/v2/log"
)

// QueueVerificationEmail sends a verification email asynchronously
func QueueVerificationEmail(emailQueue *queue.EmailQueue, email, username, url string) {
	go func() {
		ctx := context.Background()
		data := queue.VerificationEmailData{
			Username: username,
			Url:      url,
		}
		if err := emailQueue.PublishVerificationEmail(ctx, email, data, OnboardingEmail, SenderName); err != nil {
			log.Error("Failed to queue verification email", "error", err, "email", email)
		}
	}()
}

// QueueWelcomeEmail sends a welcome email asynchronously
func QueueWelcomeEmail(emailQueue *queue.EmailQueue, email, username string) {
	go func() {
		ctx := context.Background()
		data := queue.WelcomeEmailData{
			Username: username,
		}
		if err := emailQueue.PublishWelcomeEmail(ctx, email, data, WelcomeEmail, SenderName); err != nil {
			log.Error("Failed to queue welcome email", "error", err, "email", email)
		}
	}()
}
