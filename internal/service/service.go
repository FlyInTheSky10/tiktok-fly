// Package service 服务层，调用 dao 的方法及用 model 中的结构
package service

import (
	"context"
	"github.com/FlyInThesky10/TikTok-Fly/global"
	"github.com/FlyInThesky10/TikTok-Fly/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(global.DBEngine)
	return svc
}
