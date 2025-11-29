package logic

import (
	"context"
	"database/sql"
	"wayToGoZero/app/article/model"

	"wayToGoZero/app/article/cmd/rpc/internal/svc"
	"wayToGoZero/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------article-----------------------
func (l *AddArticleLogic) AddArticle(in *pb.AddArticleReq) (*pb.AddArticleResp, error) {
	// todo: add your logic here and delete this line
	article := new(model.Article)
	article.Title = in.Title
	article.Content = sql.NullString{
		String: in.Content,
		Valid:  len(in.Content) > 0, // 如果内容不为空，则设置为有效
	}
	// 调用model层的方法
	_, err := l.svcCtx.ArticleModel.Insert(l.ctx, article)
	if err != nil {
		return nil, err
	}
	return &pb.AddArticleResp{}, nil
}
