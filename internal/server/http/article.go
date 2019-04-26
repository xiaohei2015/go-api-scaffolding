package http

import (
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"go-api-scaffolding/internal/model"
	"time"
)

func ArticleAdd(c *bm.Context) {
	title := c.Request.PostForm.Get("title")
	if title == "" {
		c.JSON(nil, ecode.Error(4001, "请输入参数title"))
		return
	}
	content := c.Request.PostForm.Get("content")
	if content == "" {
		c.JSON(nil, ecode.Error(4001, "请输入参数content"))
		return
	}
	k := &model.Article{
		Title: title,
		Content: content,
		Status: 1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
		IsDeleted: 0,
	}
	res, err := svc.AddArticle(c, k)
	if err != nil{
		log.Error("incr article base err(%v) in http", err)
	}
	c.JSON(res, nil)
}
