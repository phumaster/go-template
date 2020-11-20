package connectors

import (
	"fmt"
	"go-template/configs"
	"go-template/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB var
var DB *gorm.DB

// InitDB init database
func InitDB() {
	dbconfig := (*configs.DB)(nil).GetConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbconfig.User,
		dbconfig.Password,
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	db.AutoMigrate(&models.User{})
}
