package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"lirawx.cn/go-web/models"
	"lirawx.cn/go-web/services"
)

// CreateTodo 创建todo
func CreateTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// fmt.Println("create,",c.)
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2. 存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

// GetTodoList 获取todo
func GetTodoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// GetATodo 获取一个todo
func GetATodo(c *gin.Context) {
	id := c.Query("id") // c.Query("id") // 是 c.Request.URL.Query().Get("id") 的简写
	fmt.Printf("id string get, %v\n", id)
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := services.GetATodoService(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// GetTodo 获取一个todo
func GetTodo(c *gin.Context) {
	id := c.Param("id") // 路径参数
	fmt.Printf("id string get, %v\n", id)
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := services.GetATodoService(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}
