package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"main/service"
)

type AssigneeController struct{}

// GET /assignee
func (pc AssigneeController) Index(c *gin.Context) {
	// 検索処理
	var assignee service.AssigneeService
	result, err := assignee.FindAll()
	// fmt.Println(result)

	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}


// POST /assignee
func (pc AssigneeController) Create(c *gin.Context) {
	// 検索処理
	var assignee service.AssigneeService
	result, err := assignee.Create()

	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}