package routers

import (
	"github.com/DowneyL/the-way-to-gin/middleware/jwt"
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"github.com/DowneyL/the-way-to-gin/pkg/upload"
	"github.com/DowneyL/the-way-to-gin/routers/api"
	v1 "github.com/DowneyL/the-way-to-gin/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.POST("/api/auth", api.GetAuth)
	r.POST("/api/upload", api.UploadImage)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		// tags
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		// articles
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.EditArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
