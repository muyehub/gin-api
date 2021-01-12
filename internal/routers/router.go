package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muyehub/gin-api/internal/routers/api/v1"
)

// NewRouter 新建路由
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	test := v1.NewTest()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/test", test.Get)
	}

	return r
}
