package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=fazt password=mysecretpassword dbname"

func DBConnection() {
	gorm.Open(postgres.Open())
}
