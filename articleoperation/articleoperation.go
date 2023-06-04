package main

import (
	"articleoperation/db"
	"articleoperation/model"
	"fmt"
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
	req := new(model.AddArticle)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}
	item := model.Article{
		Title:    req.Title,
		Author:   req.Author,
		Context:  req.Context,
		CreateAt: addTime.Unix(),
	}
	err := db.Mysql().Model(&model.Article{}).Create(&item).Error
	if err != nil {
		fmt.Println("Create item err", err)
	}
	resp.Msg = "文章添加成功"
	c.JSON(200, resp.Msg)
}

//DelArticle 删除文章
func DelArticle(c *gin.Context) {

	resp := BaseResp{}
	req := new(model.DelArticle)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}

	err := db.Mysql().Model(&model.Article{}).Where("id", req.Id).Delete(&model.Article{}).Error
	if err != nil {
		resp.Msg = "文章删除失败"
		c.JSON(200, resp.Msg)
		return
	}
	resp.Msg = "文章删除成功"
	c.JSON(200, resp.Msg)
}

//UpdateArticle 修改文章
func UpdateArticle(c *gin.Context) {
	up_time := time.Now()
	resp := BaseResp{}
	req := new(model.UpArticle)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}

	article := model.Article{
		Title:    req.Title,
		Author:   req.Author,
		Context:  req.Context,
		UpdateAt: up_time.Unix(),
	}
	err := db.Mysql().Model(&model.Article{}).Where("id = ?", req.Id).Updates(&article).Error
	if err != nil {
		resp.Msg = "文章修改失败"
		c.JSON(200, resp.Msg)
		return
	}
	resp.Msg = "文章修改成功"
	c.JSON(200, resp.Msg)
}

//GetArticle 查询文章
func GetArticle(c *gin.Context) {
	resp := BaseResp{}
	req := new(model.SelArticle)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}
	article := []model.Article{}

	str := "%" + req.Value + "%"
	err := db.Mysql().Model(&model.Article{}).Debug().Where("title like ?", str).
		Or("context like ?", str). //fmt.Sprintf("context like %q", ("%" + req.Value + "%"))
		Or("author like ?", str).  //fmt.Sprintf("author like %q", ("%" + req.Value + "%"))
		Find(&article).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(article) == 0 {
		resp.Msg = "没有找到文章"
		c.JSON(200, resp)
		return
	}
	resp.Data = article
	c.JSON(200, resp)
}
