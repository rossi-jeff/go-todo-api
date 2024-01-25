package seeds

import (
	"go-todo-api/models"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func SeedToDos(db *gorm.DB) {
	rows, err := db.Model(models.User{}).Where("random = ?", true).Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User

		db.ScanRows(rows, &user)

		max := gofakeit.Number(3, 10)
		for i := 0; i < max; i++ {
			wordcount := gofakeit.Number(5, 10)
			todo := models.ToDo{
				Task:      gofakeit.LoremIpsumSentence(wordcount),
				Completed: gofakeit.Bool(),
				UserId:    &user.ID,
			}
			result := db.Create(&todo)
			if result.Error != nil {
				log.Default().Println("Erros saving: " + todo.Task)
			}
		}
	}

}
