package pageSize

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPageSize(page,size int) (int, int) {
	offset := (page - 1) * size
	return offset, size
}

func GinPageSize(c *gin.Context) (int, int) {
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	size := com.StrTo(c.DefaultQuery("size", "10")).MustInt()
	offset := (page - 1) * size
	return offset, size
}
