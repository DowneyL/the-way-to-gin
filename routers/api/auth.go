package api

import (
	"github.com/DowneyL/the-way-to-gin/models"
	"github.com/DowneyL/the-way-to-gin/pkg/e"
	"github.com/DowneyL/the-way-to-gin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{username, password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if !ok {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	} else {
		isExist := models.CheckAuth(username, password)
		if isExist {
			if token, err := util.GenerateToken(username, password); err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
