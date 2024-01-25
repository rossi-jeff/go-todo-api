package main

import (
	"fmt"
	"go-todo-api/controllers"
	"go-todo-api/initializers"

	_ "go-todo-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvironment()
	initializers.DatabaseConnect()
}

// @Title ToDo API in Go
// @Description CRUD API for ToDos with Authentication
// @version         0.1.0
// @contact.name   Jeff Rossi
// @contact.url    https://resume.jeff-rossi.com/
// @contact.email  inquiries@jeff-rossi.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	router := gin.Default()

	// docs route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// auth
	router.POST("/auth/register", controllers.AuthRegister)
	router.POST("/auth/login", controllers.AuthLogin)

	// user
	router.GET("/user", controllers.UserList)
	router.GET("/user/current", controllers.UserCurrent)
	router.PATCH("/user/change", controllers.UserChangePassWord)
	router.PATCH("/user/:id", controllers.UserUpdate)
	router.DELETE("/user/:id", controllers.UserDelete)

	// todo
	router.GET("/todo", controllers.ToDoList)

	router.Run()
	fmt.Println("main")
}
