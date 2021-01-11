package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muyehub/gin-api/internal/api/v1"
)

// NewRouter 新建路由
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/test")
	}

	return r
}
