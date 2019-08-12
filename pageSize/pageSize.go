package pageSize

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPageSize(page, size int) (int, int) {
	offset := (page - 1) * size
	return offset, size
}

func GinPageSize(c *gin.Context) (int, int) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	p, err1 := strconv.Atoi(page)
	s, err2 := strconv.Atoi(size)
	if err1 != nil || err2 != nil {
		p = 1
		s = 10
	}
	offset := (p - 1) * s
	return offset, s
}
