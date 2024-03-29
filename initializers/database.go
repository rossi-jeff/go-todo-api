package initializers

import (
	"fmt"
	"go-todo-api/config"
	"go-todo-api/seeds"
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
	Seed(db)
	DB = db
}

func ConnectionString(d config.DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", d.DbUser, d.DbPass, d.DbHost, d.DbPort, d.DbName)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.ToDo{})
}

func Seed(db *gorm.DB) {
	var count int64
	db.Model(models.User{}).Where("random = ?", true).Count(&count)
	if count == 0 {
		seeds.SeedUsers(db)
		seeds.SeedToDos(db)
	}
}
