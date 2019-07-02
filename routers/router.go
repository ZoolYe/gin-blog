package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/export"
	"gin-blog/pkg/qrcode"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/upload"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)
	r.POST("/qrcode/get", api.GenerateArticlePoster)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		apiv1.POST("/tags/export", v1.ExportTag)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArtircle)
	}

	return r
}
