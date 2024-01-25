package initializers

import (
	"fmt"
	"go-todo-api/config"
	"log"

	"go-todo-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	conf := config.New()
	dsn := ConnectionString(conf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Migrate(db)
	DB = db
}

func ConnectionString(d config.DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.DbUser, d.DbPass, d.DbHost, d.DbPort, d.DbName)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.ToDo{})
}
