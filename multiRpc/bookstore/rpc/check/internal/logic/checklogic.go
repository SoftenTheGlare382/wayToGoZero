package logic

import (
	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.CheckReq) (*check.Response, error) {

	return &check.Response{
		Found: true,
		Price: 10,
	}, nil

	//resp, err := l.svcCtx.Model.FindOne(l.ctx, in.Book)
	//if err != nil {
	//	return nil, err
	//}
	//return &check.Response{
	//	Found: true,
	//	Price: resp.Price,
	//}, nil

	//ctx, cancle := context.WithTimeout(l.ctx, 8*time.Second)
	//defer cancle()

	//type result struct {
	//	price int64
	//	err   error
	//}
	//resultChan := make(chan model.Book, 1)
	//go func() {
	//	if r := recover(); r != nil {
	//		resp, err := l.svcCtx.Model.FindOne(l.ctx, in.Book)
	//		if err != nil {
	//			resultChan <- model.Book{}
	//		} else {
	//			resultChan <- *resp
	//		}
	//	}
	//}()
	//
	//select {
	//case res := <-resultChan:
	//	if res == (model.Book{}) {
	//		return &check.Response{
	//			Found: false,
	//		}, nil
	//	}
	//	return &check.Response{
	//		Found: true,
	//		Price: res.Price,
	//	}, nil
	//case <-ctx.Done():
	//	return nil, ctx.Err()
	//}

}
