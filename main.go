package main

import (
	"gindemo/common"
	"gindemo/route"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()

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
	port := viper.GetString("server.port")
	r = route.CollectRoute(r)
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}
func InitConfig() {
	workdDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workdDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
