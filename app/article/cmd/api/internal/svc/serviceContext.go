// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wayToGoZero/app/article/cmd/api/internal/config"
	"wayToGoZero/app/article/cmd/rpc/article"
)

type ServiceContext struct {
	Config     config.Config
	ArticleRpc article.ArticleZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ArticleRpc: article.NewArticleZrpcClient(zrpc.MustNewClient(c.ArticleRpcConf)),
	}
}
