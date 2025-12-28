package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	EmailQueueName = "email-jobs"
)

type EmailJobType string

const (
	EmailVerification EmailJobType = "verification"
	EmailWelcome      EmailJobType = "welcome"
)

// VerificationEmailData contains data for verification emails
type VerificationEmailData struct {
	Username string `json:"username"`
	Url      string `json:"url"`
}

// WelcomeEmailData contains data for welcome emails
type WelcomeEmailData struct {
	Username string `json:"username"`
}

// EmailJob represents a generic email job
type EmailJob struct {
	Type      EmailJobType `json:"type"`
	To        string       `json:"to"`
	FromEmail string       `json:"fromEmail"`
	FromName  string       `json:"fromName"`
	Data      interface{}  `json:"data"`
}

// VerificationEmailJob is a strongly-typed verification email job
type VerificationEmailJob struct {
	Type      EmailJobType          `json:"type"`
	To        string                `json:"to"`
	FromEmail string                `json:"fromEmail"`
	FromName  string                `json:"fromName"`
	Data      VerificationEmailData `json:"data"`
}

// WelcomeEmailJob is a strongly-typed welcome email job
type WelcomeEmailJob struct {
	Type      EmailJobType     `json:"type"`
	To        string           `json:"to"`
	FromEmail string           `json:"fromEmail"`
	FromName  string           `json:"fromName"`
	Data      WelcomeEmailData `json:"data"`
}

type EmailQueue struct {
	client *redis.Client
}

// NewEmailQueue creates a new email queue publisher
func NewEmailQueue(client *redis.Client) *EmailQueue {
	return &EmailQueue{client: client}
}

// PublishVerificationEmail queues a verification email job
func (q *EmailQueue) PublishVerificationEmail(ctx context.Context, to string, data VerificationEmailData, fromEmail, fromName string) error {
	job := VerificationEmailJob{
		Type:      EmailVerification,
		To:        to,
		FromEmail: fromEmail,
		FromName:  fromName,
		Data:      data,
	}

	return q.publishTyped(ctx, job)
}

// PublishWelcomeEmail queues a welcome email job
func (q *EmailQueue) PublishWelcomeEmail(ctx context.Context, to string, data WelcomeEmailData, fromEmail, fromName string) error {
	job := WelcomeEmailJob{
		Type:      EmailWelcome,
		To:        to,
		FromEmail: fromEmail,
		FromName:  fromName,
		Data:      data,
	}

	return q.publishTyped(ctx, job)
}

// publishTyped pushes a typed job to the Redis queue using BullMQ format
func (q *EmailQueue) publishTyped(ctx context.Context, job interface{}) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	// Generate a unique job ID
	jobID := generateJobID()

	// BullMQ stores jobs in a hash with separate fields
	pipe := q.client.Pipeline()

	// Store job data in hash with BullMQ structure
	jobKey := EmailQueueName + ":" + jobID
	pipe.HSet(ctx, jobKey, map[string]any{
		"data":      string(data),
		"opts":      "{}",
		"name":      "__default__",
		"timestamp": time.Now().UnixMilli(),
	})

	// Add job ID to wait list
	pipe.RPush(ctx, EmailQueueName+":wait", jobID)

	// Publish event to notify workers (BullMQ v5+ uses streams for events)
	pipe.XAdd(ctx, &redis.XAddArgs{
		Stream: EmailQueueName + ":events",
		Values: map[string]interface{}{
			"event": "added",
			"jobId": jobID,
		},
	})

	// Execute pipeline
	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Printf("Failed to publish email job: %v", err)
		return err
	}

	log.Printf("Published email job %s to queue", jobID)
	return nil
}

// generateJobID generates a unique job ID
func generateJobID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
