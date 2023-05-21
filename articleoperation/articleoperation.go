package main

import (
	"articleoperation/db"
	"articleoperation/model"
	"time"

	"github.com/gin-gonic/gin"
)

//文章功能 {id,标题，作者，时间，内容} 需要实现文章的增删改查功能，支持模糊查询

func init() {
	//调用db.InitMysql方法，并把数据传给host
	db.InitMysql("root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
}

//文章相关操作
func main() {
	e := gin.Default()

	e.POST("/addarticle", AddArticle)
	e.POST("/delarticle", DelArticle)
	e.POST("/updatearticle", UpdateArticle)
	e.POST("/getarticle", GetArticle)

	e.Run()
}

type BaseResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//AddArticle 添加文章
func AddArticle(c *gin.Context) {
	addTime := time.Now()
	resp := BaseResp{}
	req := new(model.ArticleInfo)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}
	item := model.ArticleInfo{
		Title:        req.Title,
		Author:       req.Author,
		Context:      req.Context,
		CreationTime: addTime.Unix(),
	}
	db.Mysql().Debug().Table("article").Create(&item)
	resp.Msg = "文章添加成功"
	c.JSON(200, resp.Msg)
}

//DelArticle 删除文章
func DelArticle(c *gin.Context) {

	//还没做文章是否存在判断

	resp := BaseResp{}
	req := new(model.ArticleInfo)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}

	article := model.ArticleInfo{}
	db.Mysql().Debug().Table("article").Where("id = ?", req.Id).First(&article)
	if article.Title == "" {
		resp.Msg = "文章不存在"
		c.JSON(200, resp.Msg)
		return
	}

	db.Mysql().Debug().Table("article").Where("id = ?", req.Id).Delete(&article)
	resp.Msg = "文章删除成功"
	c.JSON(200, resp.Msg)
}

//UpdateArticle 修改文章
func UpdateArticle(c *gin.Context) {
	modifyTime := time.Now()
	resp := BaseResp{}
	req := new(model.ArticleInfo)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}
	article := model.ArticleInfo{}
	db.Mysql().Debug().Table("article").Where("id = ?", req.Id).First(&article)
	if article.Title != "" {
		article = model.ArticleInfo{
			Title:      req.Title,
			Author:     req.Author,
			Context:    req.Context,
			ModifyTime: modifyTime.Unix(),
		}
		db.Mysql().Debug().Table("article").Where("id = ?", req.Id).Updates(&article)
		resp.Msg = "文章修改成功"
		c.JSON(200, resp.Msg)
		return
	}
	resp.Msg = "文章不存在"
	c.JSON(200, resp.Msg)
}

//GetArticle 查询文章
func GetArticle(c *gin.Context) {
	resp := BaseResp{}
	req := new(model.ArticleInfo)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}
	article := []model.ArticleInfo{}
	db.Mysql().Debug().Table("article").Where("title like ? or context like ? or author like ?", req.Value, req.Value, req.Value).Find(&article)
	if len(article) == 0 {
		resp.Msg = "没有找到文章"
		c.JSON(200, resp)
		return
	}
	resp.Data = article
	c.JSON(200, resp)
}
