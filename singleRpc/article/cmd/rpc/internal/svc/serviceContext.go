package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wayToGoZero/singleRpc/article/cmd/rpc/internal/config"
	"wayToGoZero/singleRpc/article/model"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	ArticleModel model.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
