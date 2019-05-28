package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//获取单个个文章
func GetArticle(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")
	valid.Min(id, 0, "id").Message("id不能为0")
	code := e.INVALID_PARAMS

	var data interface{}

	if !valid.HasErrors() {
		if models.ExisArticleById(id) {
			article := models.GetArticle(id)
			data = article
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
			data = models.Article{}
		}

	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code), "data": data})
}

//获取多个文章
func GetArticles(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg)
		maps["state"] = state
	}
	articles := models.GetArticles(util.GetPage(c), setting.PageSize, maps)
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

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tagId").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("文章标题不能为空")
	valid.Required(desc, "desc").Message("文章描述不能为空")
	valid.Required(content, "content").Message("文章内容不能为空")
	valid.Required(createdBy, "createdBy").Message("文章作者不能为空")
	valid.Range(state, 0, 1, "state").Message("文章状态必须在0-1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["tagId"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["createdBy"] = createdBy
		data["state"] = state
		models.AddArticle(data)
		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
}

//编辑文章
func EditArticle(c *gin.Context) {

}

//删除文章
func DeleteArtircle(c *gin.Context) {

}
