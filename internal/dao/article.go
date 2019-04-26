package dao

import (
	"context"
	xsql "database/sql"
	"github.com/bilibili/kratos/pkg/log"
	"go-api-scaffolding/internal/model"
)

const (
	_incrUser = "insert into tbl_article (`title`,`content`,`status`,`create_time`,`update_time`) values(?,?,?,?,?)"
)

//AddArticle .
func (d *Dao) AddArticle(c context.Context, article *model.Article) (num int64, err error) {
	var res xsql.Result
	if res, err = d.db.Exec(c, _incrUser, article.Title, article.Content, article.Status, article.CreateTime, article.UpdateTime); err != nil {
		log.Error("incr article base err(%v) in dao", err)
		return
	}
	return res.LastInsertId()
}