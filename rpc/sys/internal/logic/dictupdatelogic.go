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

type DictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictUpdateLogic) DictUpdate(in *sys.DictUpdateReq) (*sys.DictUpdateResp, error) {
	err := l.svcCtx.DictModel.Update(l.ctx, &sysmodel.SysDict{
		Id:          in.Id,
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        float64(in.Sort),
		UpdateBy:    sql.NullString{String: in.LastUpdateBy, Valid: true},
		UpdateTime:  time.Now(),
		Remarks:     sql.NullString{String: in.Remarks},
		DelFlag:     0,
	})

	if err != nil {
		return nil, err
	}

	return &sys.DictUpdateResp{}, nil
}
