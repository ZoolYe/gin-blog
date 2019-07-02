package api

import (
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/qrcode"
	"github.com/astaxie/beego/validation"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateArticlePoster(c *gin.Context) {
	appG := app.Gin{C: c}
	url := appG.C.PostForm("url")

	valid := validation.Validation{}
	valid.Required(url, "url").Message("url不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	qrCode := qrcode.NewQrCode(url, 300, 300, qr.M, qr.Auto)
	path := qrcode.GetQrCodeFullPath()
	_, _, err := qrCode.Encode(path)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}
