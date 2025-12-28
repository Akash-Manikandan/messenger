package models

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID         string         `gorm:"primaryKey;type:char(26)" json:"id"`
	Username   string         `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email      string         `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password   string         `gorm:"not null;size:255" json:"-"`
	Salt       string         `gorm:"not null;size:64" json:"-"`
	FirstName  string         `gorm:"size:50" json:"first_name"`
	LastName   string         `gorm:"size:50" json:"last_name"`
	Avatar     string         `gorm:"size:255" json:"avatar"`
	Bio        string         `gorm:"type:text" json:"bio"`
	IsActive   bool           `gorm:"default:true" json:"is_active"`
	IsVerified bool           `gorm:"default:false" json:"is_verified"`
	LastLogin  *time.Time     `json:"last_login"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName overrides the table name
func (User) TableName() string {
	return "users"
}

// BeforeCreate hook to generate ULID before creating
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		id := ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader)
		u.ID = id.String()
	}
	return nil
}
