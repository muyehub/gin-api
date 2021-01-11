package v1

import (
	"github.com/gin-gonic/gin"
)

// Article 文章结构体
type Article struct{}

// NewArticle 新建文章
func NewArticle() Article {
	return Article{}
}

// Get 获取单个文章
func (a Article) Get(c *gin.Context) {}

// List 获取多个文章
func (a Article) List(c *gin.Context) {}

// Create 创建文章
func (a Article) Create(c *gin.Context) {}

// Update 更新文章
func (a Article) Update(c *gin.Context) {}

// Delete 删除文章
func (a Article) Delete(c *gin.Context) {}
