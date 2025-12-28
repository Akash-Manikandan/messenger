package health

import "gorm.io/gorm"

type Service interface {
	Status() string
	PostgresStatus() string
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}

func (s *service) Status() string {
	return "ok"
}

func (s *service) PostgresStatus() string {
	if s.db == nil {
		return "disconnected"
	}

	sqlDB, err := s.db.DB()
	if err != nil {
		return "error"
	}

	if err := sqlDB.Ping(); err != nil {
		return "disconnected"
	}

	return "connected"
}
