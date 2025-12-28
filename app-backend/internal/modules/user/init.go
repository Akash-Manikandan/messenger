package user

import "gorm.io/gorm"

var dbInstance *gorm.DB

func InitModule(db *gorm.DB) {
	dbInstance = db
}
