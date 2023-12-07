// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"zero-admin/rpc/ums/internal/logic/memberservice"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"
)

type MemberServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberServiceServer
}

func NewMemberServiceServer(svcCtx *svc.ServiceContext) *MemberServiceServer {
	return &MemberServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *MemberServiceServer) MemberAdd(ctx context.Context, in *umsclient.MemberAddReq) (*umsclient.MemberAddResp, error) {
	l := memberservicelogic.NewMemberAddLogic(ctx, s.svcCtx)
	return l.MemberAdd(in)
}

func (s *MemberServiceServer) MemberLogin(ctx context.Context, in *umsclient.MemberLoginReq) (*umsclient.MemberLoginResp, error) {
	l := memberservicelogic.NewMemberLoginLogic(ctx, s.svcCtx)
	return l.MemberLogin(in)
}

func (s *MemberServiceServer) MemberList(ctx context.Context, in *umsclient.MemberListReq) (*umsclient.MemberListResp, error) {
	l := memberservicelogic.NewMemberListLogic(ctx, s.svcCtx)
	return l.MemberList(in)
}

func (s *MemberServiceServer) MemberUpdate(ctx context.Context, in *umsclient.MemberUpdateReq) (*umsclient.MemberUpdateResp, error) {
	l := memberservicelogic.NewMemberUpdateLogic(ctx, s.svcCtx)
	return l.MemberUpdate(in)
}

func (s *MemberServiceServer) MemberDelete(ctx context.Context, in *umsclient.MemberDeleteReq) (*umsclient.MemberDeleteResp, error) {
	l := memberservicelogic.NewMemberDeleteLogic(ctx, s.svcCtx)
	return l.MemberDelete(in)
}

func (s *MemberServiceServer) QueryMemberById(ctx context.Context, in *umsclient.MemberByIdReq) (*umsclient.MemberListData, error) {
	l := memberservicelogic.NewQueryMemberByIdLogic(ctx, s.svcCtx)
	return l.QueryMemberById(in)
}

func (s *MemberServiceServer) MemberUpdatePassword(ctx context.Context, in *umsclient.MemberUpdatePasswordReq) (*umsclient.MemberUpdateResp, error) {
	l := memberservicelogic.NewMemberUpdatePasswordLogic(ctx, s.svcCtx)
	return l.MemberUpdatePassword(in)
}

func (s *MemberServiceServer) UpdateMemberIntegration(ctx context.Context, in *umsclient.UpdateMemberIntegrationReq) (*umsclient.UpdateMemberIntegrationResp, error) {
	l := memberservicelogic.NewUpdateMemberIntegrationLogic(ctx, s.svcCtx)
	return l.UpdateMemberIntegration(in)
}
