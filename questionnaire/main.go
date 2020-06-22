package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"questionnaire/controllers"
)

// @title gin 框架
// @version 1.0
// @description 给予gin web框架搭建的业务骨架
// @termsofservice http://swagger.io/terms/
// @contact.name jinchunguang
// @contact.email jin-chunguang@foxmail.com
// @host localhost:10010
func main() {
	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", controllers.FindUsers )
	r.GET("/user/:id", controllers.FindUser )
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.POST("/questionnaire", controllers.CreateQuestionnaire)

	_ = r.Run()
}
