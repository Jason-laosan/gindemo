package main

import (
	"gindemo/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/regester", controller.Register)
	return r
}
