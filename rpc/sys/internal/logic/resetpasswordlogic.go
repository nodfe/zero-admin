package logic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReSetPasswordLogic) ReSetPassword(in *sys.ReSetPasswordReq) (*sys.ReSetPasswordResp, error) {

	_ = l.svcCtx.UserModel.Update(l.ctx, &sysmodel.SysUser{
		Id:         in.Id,
		Password:   "123456",
		Salt:       "123456",
		UpdateBy:   sql.NullString{String: in.LastUpdateBy, Valid: true},
		UpdateTime: sql.NullTime{Time: time.Now()},
	})

	return &sys.ReSetPasswordResp{}, nil
}
