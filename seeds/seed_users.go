package seeds

import (
	"go-todo-api/config"
	"go-todo-api/models"
	"go-todo-api/utilities"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	max := 20
	conf := config.New()
	pw := conf.Secret.UserPass
	hashed, _ := utilities.HashPassword(pw)
	for i := 0; i < max; i++ {
		user := models.User{
			UserName:       gofakeit.Username(),
			Email:          gofakeit.Email(),
			PassWordDigest: hashed,
			Random:         true,
		}
		result := db.Create(&user)
		if result.Error != nil {
			log.Default().Println("Error saving: " + user.UserName)
		}
	}
}
