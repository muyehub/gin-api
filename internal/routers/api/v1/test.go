package v1

import (
	"github.com/gin-gonic/gin"
)

// Test 文章结构体
type Test struct{}

// NewTest 新建 Test 类
func NewTest() Test {
	return Test{}
}

// Get 获取单个文章
func (t Test) Get(c *gin.Context) {}

// List 获取多个文章
func (t Test) List(c *gin.Context) {}

// Create 创建文章
func (t Test) Create(c *gin.Context) {}

// Update 更新文章
func (t Test) Update(c *gin.Context) {}

// Delete 删除文章
func (t Test) Delete(c *gin.Context) {}
