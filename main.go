package main

import (
	"gindemo/common"

	"github.com/gin-gonic/gin"
)

func main() {
	// db := common.GetDB()
	common.InitDB()

	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.GET("/student", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "查询学生成绩成功",
	// 	})
	// })
	// r.POST("/create_student", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "创建学生成功",
	// 	})
	// })
	// r.DELETE("/delete_student", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "删除学生成功",
	// 	})
	// })
	r = CollectRoute(r)

	panic(r.Run())
}
