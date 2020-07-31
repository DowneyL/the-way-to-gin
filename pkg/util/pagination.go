package util

import (
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		return (page - 1) * setting.AppSetting.PageSize
	}
	return 0
}
