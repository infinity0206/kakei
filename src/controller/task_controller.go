package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"main/service"
)

type TaskController struct{}

// GET /task
func (pc TaskController) Index(c *gin.Context) {
	// 検索処理
	var task service.TaskService
	result, err := task.FindAll()

	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}


// POST /task
func (pc TaskController) Create(c *gin.Context) {
	// 検索処理
	var task service.TaskService
	result, err := task.Create()

	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}