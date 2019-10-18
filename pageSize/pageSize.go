package pageSize

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPageSize(page, size int) (int, int) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	if size > 500 {
		size = 500
	}
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
