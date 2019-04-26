package service

import (
	"context"
	"go-api-scaffolding/internal/model"

	"github.com/bilibili/kratos/pkg/log"
)

func (s *Service) AddArticle(c context.Context, article *model.Article) (res *model.Article, err error) {
	res = &model.Article{}
	*res = *article
	if res.Id, err = s.dao.AddArticle(c, article); err != nil {
		log.Error("incr article base err(%v) in service", err)
	}

	return
}