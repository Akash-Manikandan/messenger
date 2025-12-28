package user

import (
	"errors"

	"github.com/Akash-Manikandan/app-backend/internal/config"
	"github.com/Akash-Manikandan/app-backend/internal/models"
	"github.com/Akash-Manikandan/app-backend/pkg/crypto"
	"github.com/Akash-Manikandan/app-backend/pkg/queue"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Service interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	ListUsers(limit, offset int) ([]models.User, error)
	VerifyUser(userID string) error
}

type service struct {
	db         *gorm.DB
	emailQueue *queue.EmailQueue
}

func NewService(db *gorm.DB, redis *redis.Client) Service {
	return &service{
		db:         db,
		emailQueue: queue.NewEmailQueue(redis),
	}
}

func (s *service) CreateUser(user *models.User) error {
	// Validate password strength
	if err := crypto.ValidatePasswordStrength(user.Password); err != nil {
		return err
	}

	// Generate salt for this user
	salt, err := crypto.GenerateSalt()
	if err != nil {
		return err
	}
	user.Salt = salt

	// Hash the password with the salt
	hashedPassword, err := crypto.HashPassword(user.Password, salt)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Create user in database
	if err := s.db.Create(user).Error; err != nil {
		return err
	}

	cfg := config.Load()

	// Queue verification email asynchronously (non-blocking)
	verificationUrl := cfg.AppBackendURL + "/api/users/" + user.ID + "/verify"
	QueueVerificationEmail(s.emailQueue, user.Email, user.Username, verificationUrl)

	return nil
}

func (s *service) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *service) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *service) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := s.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *service) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

func (s *service) DeleteUser(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.User{}).Error
}

func (s *service) ListUsers(limit, offset int) ([]models.User, error) {
	var users []models.User
	err := s.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (s *service) VerifyUser(userID string) error {
	var user models.User

	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(ErrUserNotFound)
		}
		return err
	}

	if user.IsVerified {
		return errors.New(ErrUserAlreadyVerified)
	}

	user.IsVerified = true
	QueueWelcomeEmail(s.emailQueue, user.Email, user.Username)
	return s.db.Save(&user).Error
}
