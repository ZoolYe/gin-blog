package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"gin-blog/service/article_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取单个个文章
func GetArticle(c *gin.Context) {

	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{ID: id}
	exist := articleService.ExistByID()
	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, article)
}

//获取多个文章
func GetArticles(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg)
		maps["state"] = state
	}
	articles := models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
	total := models.GetArticleTotal(maps)
	data["list"] = articles
	data["total"] = total
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code), "data": data})
}

//添加文章
func AddArticle(c *gin.Context) {

	tagId := com.StrTo(c.Query("tagId")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("createdBy")
	state := com.StrTo(c.DefaultQuery("state", com.StrTo(1).String())).MustInt()
	coverImageUrl := c.Query("coverImageUrl")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tagId").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("文章标题不能为空")
	valid.Required(desc, "desc").Message("文章描述不能为空")
	valid.Required(content, "content").Message("文章内容不能为空")
	valid.Required(createdBy, "createdBy").Message("文章作者不能为空")
	valid.Range(state, 0, 1, "state").Message("文章状态必须在0-1")
	valid.Required(coverImageUrl, "coverImageUrl").Message("文章图片不能为空")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["tagId"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["createdBy"] = createdBy
		data["state"] = state
		data["coverImageUrl"] = coverImageUrl
		models.AddArticle(data)
		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
}

//编辑文章
func EditArticle(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	modifiedBy := c.Query("modifiedBy")
	valid := validation.Validation{}

	data := make(map[string]interface{})

	if arg := c.Query("tagId"); arg != "" {
		tagId := com.StrTo(arg).MustInt()
		data["tagId"] = tagId
		data["modifiedBy"] = modifiedBy
	}

	if desc := c.Query("desc"); desc != "" {
		data["desc"] = desc
		data["modifiedBy"] = modifiedBy
	}

	if title := c.Query("title"); title != "" {
		data["title"] = title
		data["modifiedBy"] = modifiedBy
	}

	if content := c.Query("content"); content != "" {
		data["content"] = content
		data["modifiedBy"] = modifiedBy
	}

	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("文章状态必须在0-1")
		data["state"] = state
	}

	if coverImageUrl := c.Query("coverImageUrl"); coverImageUrl != "" {
		data["coverImageUrl"] = coverImageUrl
		data["modifiedBy"] = modifiedBy
	}

	valid.Min(id, 1, "id").Message("id不能小于0")
	valid.Required(modifiedBy, "modifiedBy").Message("修改人不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExisArticleById(id) {
			models.EditArticle(id, data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
}

//删除文章
func DeleteArtircle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id不能小于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExisArticleById(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
}
