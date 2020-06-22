package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"questionnaire/models"
)


func FindUsers(c *gin.Context) {
	users := models.UserList()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})

}
// @description 通过id获取用户
// @version 1.0
// @param id path int true "id"
// @router /user/{id} [get]
func FindUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"operation": "Find one"})
}

func CreateUser(c *gin.Context) {
	name := c.Request.FormValue("name")
	gender := c.Request.FormValue("gender")
	telephoneNumber := c.Request.FormValue("telephoneNumber")

	user := models.Consumer{Name: name, Gender: gender, TelephoneNumber: telephoneNumber}
	ra := models.UserAdd(user)
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"operation": "Update"})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"operation": "Delete"})
}

func CreateQuestionnaire(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"operation": "提交问卷成功"})
}
