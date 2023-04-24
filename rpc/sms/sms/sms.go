// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package sms

import (
	"context"

	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CouponAddReq                            = smsclient.CouponAddReq
	CouponAddResp                           = smsclient.CouponAddResp
	CouponDeleteReq                         = smsclient.CouponDeleteReq
	CouponDeleteResp                        = smsclient.CouponDeleteResp
	CouponHistoryAddReq                     = smsclient.CouponHistoryAddReq
	CouponHistoryAddResp                    = smsclient.CouponHistoryAddResp
	CouponHistoryDeleteReq                  = smsclient.CouponHistoryDeleteReq
	CouponHistoryDeleteResp                 = smsclient.CouponHistoryDeleteResp
	CouponHistoryListData                   = smsclient.CouponHistoryListData
	CouponHistoryListReq                    = smsclient.CouponHistoryListReq
	CouponHistoryListResp                   = smsclient.CouponHistoryListResp
	CouponHistoryUpdateReq                  = smsclient.CouponHistoryUpdateReq
	CouponHistoryUpdateResp                 = smsclient.CouponHistoryUpdateResp
	CouponListData                          = smsclient.CouponListData
	CouponListReq                           = smsclient.CouponListReq
	CouponListResp                          = smsclient.CouponListResp
	CouponProductCategoryRelationAddReq     = smsclient.CouponProductCategoryRelationAddReq
	CouponProductCategoryRelationAddResp    = smsclient.CouponProductCategoryRelationAddResp
	CouponProductCategoryRelationDeleteReq  = smsclient.CouponProductCategoryRelationDeleteReq
	CouponProductCategoryRelationDeleteResp = smsclient.CouponProductCategoryRelationDeleteResp
	CouponProductCategoryRelationListData   = smsclient.CouponProductCategoryRelationListData
	CouponProductCategoryRelationListReq    = smsclient.CouponProductCategoryRelationListReq
	CouponProductCategoryRelationListResp   = smsclient.CouponProductCategoryRelationListResp
	CouponProductCategoryRelationUpdateReq  = smsclient.CouponProductCategoryRelationUpdateReq
	CouponProductCategoryRelationUpdateResp = smsclient.CouponProductCategoryRelationUpdateResp
	CouponProductRelationAddReq             = smsclient.CouponProductRelationAddReq
	CouponProductRelationAddResp            = smsclient.CouponProductRelationAddResp
	CouponProductRelationDeleteReq          = smsclient.CouponProductRelationDeleteReq
	CouponProductRelationDeleteResp         = smsclient.CouponProductRelationDeleteResp
	CouponProductRelationListData           = smsclient.CouponProductRelationListData
	CouponProductRelationListReq            = smsclient.CouponProductRelationListReq
	CouponProductRelationListResp           = smsclient.CouponProductRelationListResp
	CouponProductRelationUpdateReq          = smsclient.CouponProductRelationUpdateReq
	CouponProductRelationUpdateResp         = smsclient.CouponProductRelationUpdateResp
	CouponUpdateReq                         = smsclient.CouponUpdateReq
	CouponUpdateResp                        = smsclient.CouponUpdateResp
	FlashPromotionAddReq                    = smsclient.FlashPromotionAddReq
	FlashPromotionAddResp                   = smsclient.FlashPromotionAddResp
	FlashPromotionDeleteReq                 = smsclient.FlashPromotionDeleteReq
	FlashPromotionDeleteResp                = smsclient.FlashPromotionDeleteResp
	FlashPromotionListData                  = smsclient.FlashPromotionListData
	FlashPromotionListReq                   = smsclient.FlashPromotionListReq
	FlashPromotionListResp                  = smsclient.FlashPromotionListResp
	FlashPromotionLogAddReq                 = smsclient.FlashPromotionLogAddReq
	FlashPromotionLogAddResp                = smsclient.FlashPromotionLogAddResp
	FlashPromotionLogDeleteReq              = smsclient.FlashPromotionLogDeleteReq
	FlashPromotionLogDeleteResp             = smsclient.FlashPromotionLogDeleteResp
	FlashPromotionLogListData               = smsclient.FlashPromotionLogListData
	FlashPromotionLogListReq                = smsclient.FlashPromotionLogListReq
	FlashPromotionLogListResp               = smsclient.FlashPromotionLogListResp
	FlashPromotionLogUpdateReq              = smsclient.FlashPromotionLogUpdateReq
	FlashPromotionLogUpdateResp             = smsclient.FlashPromotionLogUpdateResp
	FlashPromotionProductRelationAddReq     = smsclient.FlashPromotionProductRelationAddReq
	FlashPromotionProductRelationAddResp    = smsclient.FlashPromotionProductRelationAddResp
	FlashPromotionProductRelationDeleteReq  = smsclient.FlashPromotionProductRelationDeleteReq
	FlashPromotionProductRelationDeleteResp = smsclient.FlashPromotionProductRelationDeleteResp
	FlashPromotionProductRelationListData   = smsclient.FlashPromotionProductRelationListData
	FlashPromotionProductRelationListReq    = smsclient.FlashPromotionProductRelationListReq
	FlashPromotionProductRelationListResp   = smsclient.FlashPromotionProductRelationListResp
	FlashPromotionProductRelationUpdateReq  = smsclient.FlashPromotionProductRelationUpdateReq
	FlashPromotionProductRelationUpdateResp = smsclient.FlashPromotionProductRelationUpdateResp
	FlashPromotionSessionAddReq             = smsclient.FlashPromotionSessionAddReq
	FlashPromotionSessionAddResp            = smsclient.FlashPromotionSessionAddResp
	FlashPromotionSessionDeleteReq          = smsclient.FlashPromotionSessionDeleteReq
	FlashPromotionSessionDeleteResp         = smsclient.FlashPromotionSessionDeleteResp
	FlashPromotionSessionListData           = smsclient.FlashPromotionSessionListData
	FlashPromotionSessionListReq            = smsclient.FlashPromotionSessionListReq
	FlashPromotionSessionListResp           = smsclient.FlashPromotionSessionListResp
	FlashPromotionSessionUpdateReq          = smsclient.FlashPromotionSessionUpdateReq
	FlashPromotionSessionUpdateResp         = smsclient.FlashPromotionSessionUpdateResp
	FlashPromotionUpdateReq                 = smsclient.FlashPromotionUpdateReq
	FlashPromotionUpdateResp                = smsclient.FlashPromotionUpdateResp
	HomeAdvertiseAddReq                     = smsclient.HomeAdvertiseAddReq
	HomeAdvertiseAddResp                    = smsclient.HomeAdvertiseAddResp
	HomeAdvertiseDeleteReq                  = smsclient.HomeAdvertiseDeleteReq
	HomeAdvertiseDeleteResp                 = smsclient.HomeAdvertiseDeleteResp
	HomeAdvertiseListData                   = smsclient.HomeAdvertiseListData
	HomeAdvertiseListReq                    = smsclient.HomeAdvertiseListReq
	HomeAdvertiseListResp                   = smsclient.HomeAdvertiseListResp
	HomeAdvertiseUpdateReq                  = smsclient.HomeAdvertiseUpdateReq
	HomeAdvertiseUpdateResp                 = smsclient.HomeAdvertiseUpdateResp
	HomeBrandAddData                        = smsclient.HomeBrandAddData
	HomeBrandAddReq                         = smsclient.HomeBrandAddReq
	HomeBrandAddResp                        = smsclient.HomeBrandAddResp
	HomeBrandDeleteReq                      = smsclient.HomeBrandDeleteReq
	HomeBrandDeleteResp                     = smsclient.HomeBrandDeleteResp
	HomeBrandListData                       = smsclient.HomeBrandListData
	HomeBrandListReq                        = smsclient.HomeBrandListReq
	HomeBrandListResp                       = smsclient.HomeBrandListResp
	HomeBrandUpdateReq                      = smsclient.HomeBrandUpdateReq
	HomeBrandUpdateResp                     = smsclient.HomeBrandUpdateResp
	HomeNewProductAddData                   = smsclient.HomeNewProductAddData
	HomeNewProductAddReq                    = smsclient.HomeNewProductAddReq
	HomeNewProductAddResp                   = smsclient.HomeNewProductAddResp
	HomeNewProductDeleteReq                 = smsclient.HomeNewProductDeleteReq
	HomeNewProductDeleteResp                = smsclient.HomeNewProductDeleteResp
	HomeNewProductListData                  = smsclient.HomeNewProductListData
	HomeNewProductListReq                   = smsclient.HomeNewProductListReq
	HomeNewProductListResp                  = smsclient.HomeNewProductListResp
	HomeNewProductUpdateReq                 = smsclient.HomeNewProductUpdateReq
	HomeNewProductUpdateResp                = smsclient.HomeNewProductUpdateResp
	HomeRecommendProductAddData             = smsclient.HomeRecommendProductAddData
	HomeRecommendProductAddReq              = smsclient.HomeRecommendProductAddReq
	HomeRecommendProductAddResp             = smsclient.HomeRecommendProductAddResp
	HomeRecommendProductDeleteReq           = smsclient.HomeRecommendProductDeleteReq
	HomeRecommendProductDeleteResp          = smsclient.HomeRecommendProductDeleteResp
	HomeRecommendProductListData            = smsclient.HomeRecommendProductListData
	HomeRecommendProductListReq             = smsclient.HomeRecommendProductListReq
	HomeRecommendProductListResp            = smsclient.HomeRecommendProductListResp
	HomeRecommendProductUpdateReq           = smsclient.HomeRecommendProductUpdateReq
	HomeRecommendProductUpdateResp          = smsclient.HomeRecommendProductUpdateResp
	HomeRecommendSubjectAddReq              = smsclient.HomeRecommendSubjectAddReq
	HomeRecommendSubjectAddResp             = smsclient.HomeRecommendSubjectAddResp
	HomeRecommendSubjectDeleteReq           = smsclient.HomeRecommendSubjectDeleteReq
	HomeRecommendSubjectDeleteResp          = smsclient.HomeRecommendSubjectDeleteResp
	HomeRecommendSubjectListData            = smsclient.HomeRecommendSubjectListData
	HomeRecommendSubjectListReq             = smsclient.HomeRecommendSubjectListReq
	HomeRecommendSubjectListResp            = smsclient.HomeRecommendSubjectListResp
	HomeRecommendSubjectUpdateReq           = smsclient.HomeRecommendSubjectUpdateReq
	HomeRecommendSubjectUpdateResp          = smsclient.HomeRecommendSubjectUpdateResp

	Sms interface {
		CouponAdd(ctx context.Context, in *CouponAddReq, opts ...grpc.CallOption) (*CouponAddResp, error)
		CouponList(ctx context.Context, in *CouponListReq, opts ...grpc.CallOption) (*CouponListResp, error)
		CouponUpdate(ctx context.Context, in *CouponUpdateReq, opts ...grpc.CallOption) (*CouponUpdateResp, error)
		CouponDelete(ctx context.Context, in *CouponDeleteReq, opts ...grpc.CallOption) (*CouponDeleteResp, error)
		CouponHistoryAdd(ctx context.Context, in *CouponHistoryAddReq, opts ...grpc.CallOption) (*CouponHistoryAddResp, error)
		CouponHistoryList(ctx context.Context, in *CouponHistoryListReq, opts ...grpc.CallOption) (*CouponHistoryListResp, error)
		CouponHistoryUpdate(ctx context.Context, in *CouponHistoryUpdateReq, opts ...grpc.CallOption) (*CouponHistoryUpdateResp, error)
		CouponHistoryDelete(ctx context.Context, in *CouponHistoryDeleteReq, opts ...grpc.CallOption) (*CouponHistoryDeleteResp, error)
		CouponProductCategoryRelationAdd(ctx context.Context, in *CouponProductCategoryRelationAddReq, opts ...grpc.CallOption) (*CouponProductCategoryRelationAddResp, error)
		CouponProductCategoryRelationList(ctx context.Context, in *CouponProductCategoryRelationListReq, opts ...grpc.CallOption) (*CouponProductCategoryRelationListResp, error)
		CouponProductCategoryRelationUpdate(ctx context.Context, in *CouponProductCategoryRelationUpdateReq, opts ...grpc.CallOption) (*CouponProductCategoryRelationUpdateResp, error)
		CouponProductCategoryRelationDelete(ctx context.Context, in *CouponProductCategoryRelationDeleteReq, opts ...grpc.CallOption) (*CouponProductCategoryRelationDeleteResp, error)
		CouponProductRelationAdd(ctx context.Context, in *CouponProductRelationAddReq, opts ...grpc.CallOption) (*CouponProductRelationAddResp, error)
		CouponProductRelationList(ctx context.Context, in *CouponProductRelationListReq, opts ...grpc.CallOption) (*CouponProductRelationListResp, error)
		CouponProductRelationUpdate(ctx context.Context, in *CouponProductRelationUpdateReq, opts ...grpc.CallOption) (*CouponProductRelationUpdateResp, error)
		CouponProductRelationDelete(ctx context.Context, in *CouponProductRelationDeleteReq, opts ...grpc.CallOption) (*CouponProductRelationDeleteResp, error)
		FlashPromotionLogAdd(ctx context.Context, in *FlashPromotionLogAddReq, opts ...grpc.CallOption) (*FlashPromotionLogAddResp, error)
		FlashPromotionLogList(ctx context.Context, in *FlashPromotionLogListReq, opts ...grpc.CallOption) (*FlashPromotionLogListResp, error)
		FlashPromotionLogUpdate(ctx context.Context, in *FlashPromotionLogUpdateReq, opts ...grpc.CallOption) (*FlashPromotionLogUpdateResp, error)
		FlashPromotionLogDelete(ctx context.Context, in *FlashPromotionLogDeleteReq, opts ...grpc.CallOption) (*FlashPromotionLogDeleteResp, error)
		FlashPromotionAdd(ctx context.Context, in *FlashPromotionAddReq, opts ...grpc.CallOption) (*FlashPromotionAddResp, error)
		FlashPromotionList(ctx context.Context, in *FlashPromotionListReq, opts ...grpc.CallOption) (*FlashPromotionListResp, error)
		FlashPromotionUpdate(ctx context.Context, in *FlashPromotionUpdateReq, opts ...grpc.CallOption) (*FlashPromotionUpdateResp, error)
		FlashPromotionDelete(ctx context.Context, in *FlashPromotionDeleteReq, opts ...grpc.CallOption) (*FlashPromotionDeleteResp, error)
		FlashPromotionProductRelationAdd(ctx context.Context, in *FlashPromotionProductRelationAddReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationAddResp, error)
		FlashPromotionProductRelationList(ctx context.Context, in *FlashPromotionProductRelationListReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationListResp, error)
		FlashPromotionProductRelationUpdate(ctx context.Context, in *FlashPromotionProductRelationUpdateReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationUpdateResp, error)
		FlashPromotionProductRelationDelete(ctx context.Context, in *FlashPromotionProductRelationDeleteReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationDeleteResp, error)
		FlashPromotionSessionAdd(ctx context.Context, in *FlashPromotionSessionAddReq, opts ...grpc.CallOption) (*FlashPromotionSessionAddResp, error)
		FlashPromotionSessionList(ctx context.Context, in *FlashPromotionSessionListReq, opts ...grpc.CallOption) (*FlashPromotionSessionListResp, error)
		FlashPromotionSessionUpdate(ctx context.Context, in *FlashPromotionSessionUpdateReq, opts ...grpc.CallOption) (*FlashPromotionSessionUpdateResp, error)
		FlashPromotionSessionDelete(ctx context.Context, in *FlashPromotionSessionDeleteReq, opts ...grpc.CallOption) (*FlashPromotionSessionDeleteResp, error)
		HomeAdvertiseAdd(ctx context.Context, in *HomeAdvertiseAddReq, opts ...grpc.CallOption) (*HomeAdvertiseAddResp, error)
		HomeAdvertiseList(ctx context.Context, in *HomeAdvertiseListReq, opts ...grpc.CallOption) (*HomeAdvertiseListResp, error)
		HomeAdvertiseUpdate(ctx context.Context, in *HomeAdvertiseUpdateReq, opts ...grpc.CallOption) (*HomeAdvertiseUpdateResp, error)
		HomeAdvertiseDelete(ctx context.Context, in *HomeAdvertiseDeleteReq, opts ...grpc.CallOption) (*HomeAdvertiseDeleteResp, error)
		HomeBrandAdd(ctx context.Context, in *HomeBrandAddReq, opts ...grpc.CallOption) (*HomeBrandAddResp, error)
		HomeBrandList(ctx context.Context, in *HomeBrandListReq, opts ...grpc.CallOption) (*HomeBrandListResp, error)
		HomeBrandUpdate(ctx context.Context, in *HomeBrandUpdateReq, opts ...grpc.CallOption) (*HomeBrandUpdateResp, error)
		HomeBrandDelete(ctx context.Context, in *HomeBrandDeleteReq, opts ...grpc.CallOption) (*HomeBrandDeleteResp, error)
		HomeNewProductAdd(ctx context.Context, in *HomeNewProductAddReq, opts ...grpc.CallOption) (*HomeNewProductAddResp, error)
		HomeNewProductList(ctx context.Context, in *HomeNewProductListReq, opts ...grpc.CallOption) (*HomeNewProductListResp, error)
		HomeNewProductUpdate(ctx context.Context, in *HomeNewProductUpdateReq, opts ...grpc.CallOption) (*HomeNewProductUpdateResp, error)
		HomeNewProductDelete(ctx context.Context, in *HomeNewProductDeleteReq, opts ...grpc.CallOption) (*HomeNewProductDeleteResp, error)
		HomeRecommendProductAdd(ctx context.Context, in *HomeRecommendProductAddReq, opts ...grpc.CallOption) (*HomeRecommendProductAddResp, error)
		HomeRecommendProductList(ctx context.Context, in *HomeRecommendProductListReq, opts ...grpc.CallOption) (*HomeRecommendProductListResp, error)
		HomeRecommendProductUpdate(ctx context.Context, in *HomeRecommendProductUpdateReq, opts ...grpc.CallOption) (*HomeRecommendProductUpdateResp, error)
		HomeRecommendProductDelete(ctx context.Context, in *HomeRecommendProductDeleteReq, opts ...grpc.CallOption) (*HomeRecommendProductDeleteResp, error)
		HomeRecommendSubjectAdd(ctx context.Context, in *HomeRecommendSubjectAddReq, opts ...grpc.CallOption) (*HomeRecommendSubjectAddResp, error)
		HomeRecommendSubjectList(ctx context.Context, in *HomeRecommendSubjectListReq, opts ...grpc.CallOption) (*HomeRecommendSubjectListResp, error)
		HomeRecommendSubjectUpdate(ctx context.Context, in *HomeRecommendSubjectUpdateReq, opts ...grpc.CallOption) (*HomeRecommendSubjectUpdateResp, error)
		HomeRecommendSubjectDelete(ctx context.Context, in *HomeRecommendSubjectDeleteReq, opts ...grpc.CallOption) (*HomeRecommendSubjectDeleteResp, error)
	}

	defaultSms struct {
		cli zrpc.Client
	}
)

func NewSms(cli zrpc.Client) Sms {
	return &defaultSms{
		cli: cli,
	}
}

func (m *defaultSms) CouponAdd(ctx context.Context, in *CouponAddReq, opts ...grpc.CallOption) (*CouponAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponAdd(ctx, in, opts...)
}

func (m *defaultSms) CouponList(ctx context.Context, in *CouponListReq, opts ...grpc.CallOption) (*CouponListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponList(ctx, in, opts...)
}

func (m *defaultSms) CouponUpdate(ctx context.Context, in *CouponUpdateReq, opts ...grpc.CallOption) (*CouponUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponUpdate(ctx, in, opts...)
}

func (m *defaultSms) CouponDelete(ctx context.Context, in *CouponDeleteReq, opts ...grpc.CallOption) (*CouponDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponDelete(ctx, in, opts...)
}

func (m *defaultSms) CouponHistoryAdd(ctx context.Context, in *CouponHistoryAddReq, opts ...grpc.CallOption) (*CouponHistoryAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponHistoryAdd(ctx, in, opts...)
}

func (m *defaultSms) CouponHistoryList(ctx context.Context, in *CouponHistoryListReq, opts ...grpc.CallOption) (*CouponHistoryListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponHistoryList(ctx, in, opts...)
}

func (m *defaultSms) CouponHistoryUpdate(ctx context.Context, in *CouponHistoryUpdateReq, opts ...grpc.CallOption) (*CouponHistoryUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponHistoryUpdate(ctx, in, opts...)
}

func (m *defaultSms) CouponHistoryDelete(ctx context.Context, in *CouponHistoryDeleteReq, opts ...grpc.CallOption) (*CouponHistoryDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponHistoryDelete(ctx, in, opts...)
}

func (m *defaultSms) CouponProductCategoryRelationAdd(ctx context.Context, in *CouponProductCategoryRelationAddReq, opts ...grpc.CallOption) (*CouponProductCategoryRelationAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponProductCategoryRelationAdd(ctx, in, opts...)
}

func (m *defaultSms) CouponProductCategoryRelationList(ctx context.Context, in *CouponProductCategoryRelationListReq, opts ...grpc.CallOption) (*CouponProductCategoryRelationListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponProductCategoryRelationList(ctx, in, opts...)
}

func (m *defaultSms) CouponProductCategoryRelationUpdate(ctx context.Context, in *CouponProductCategoryRelationUpdateReq, opts ...grpc.CallOption) (*CouponProductCategoryRelationUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponProductCategoryRelationUpdate(ctx, in, opts...)
}

func (m *defaultSms) CouponProductCategoryRelationDelete(ctx context.Context, in *CouponProductCategoryRelationDeleteReq, opts ...grpc.CallOption) (*CouponProductCategoryRelationDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponProductCategoryRelationDelete(ctx, in, opts...)
}

func (m *defaultSms) CouponProductRelationAdd(ctx context.Context, in *CouponProductRelationAddReq, opts ...grpc.CallOption) (*CouponProductRelationAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponProductRelationAdd(ctx, in, opts...)
}

func (m *defaultSms) CouponProductRelationList(ctx context.Context, in *CouponProductRelationListReq, opts ...grpc.CallOption) (*CouponProductRelationListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponProductRelationList(ctx, in, opts...)
}

func (m *defaultSms) CouponProductRelationUpdate(ctx context.Context, in *CouponProductRelationUpdateReq, opts ...grpc.CallOption) (*CouponProductRelationUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponProductRelationUpdate(ctx, in, opts...)
}

func (m *defaultSms) CouponProductRelationDelete(ctx context.Context, in *CouponProductRelationDeleteReq, opts ...grpc.CallOption) (*CouponProductRelationDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.CouponProductRelationDelete(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionLogAdd(ctx context.Context, in *FlashPromotionLogAddReq, opts ...grpc.CallOption) (*FlashPromotionLogAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionLogAdd(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionLogList(ctx context.Context, in *FlashPromotionLogListReq, opts ...grpc.CallOption) (*FlashPromotionLogListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionLogList(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionLogUpdate(ctx context.Context, in *FlashPromotionLogUpdateReq, opts ...grpc.CallOption) (*FlashPromotionLogUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionLogUpdate(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionLogDelete(ctx context.Context, in *FlashPromotionLogDeleteReq, opts ...grpc.CallOption) (*FlashPromotionLogDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionLogDelete(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionAdd(ctx context.Context, in *FlashPromotionAddReq, opts ...grpc.CallOption) (*FlashPromotionAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionAdd(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionList(ctx context.Context, in *FlashPromotionListReq, opts ...grpc.CallOption) (*FlashPromotionListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionList(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionUpdate(ctx context.Context, in *FlashPromotionUpdateReq, opts ...grpc.CallOption) (*FlashPromotionUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionUpdate(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionDelete(ctx context.Context, in *FlashPromotionDeleteReq, opts ...grpc.CallOption) (*FlashPromotionDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionDelete(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionProductRelationAdd(ctx context.Context, in *FlashPromotionProductRelationAddReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionProductRelationAdd(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionProductRelationList(ctx context.Context, in *FlashPromotionProductRelationListReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionProductRelationList(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionProductRelationUpdate(ctx context.Context, in *FlashPromotionProductRelationUpdateReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionProductRelationUpdate(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionProductRelationDelete(ctx context.Context, in *FlashPromotionProductRelationDeleteReq, opts ...grpc.CallOption) (*FlashPromotionProductRelationDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionProductRelationDelete(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionSessionAdd(ctx context.Context, in *FlashPromotionSessionAddReq, opts ...grpc.CallOption) (*FlashPromotionSessionAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionSessionAdd(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionSessionList(ctx context.Context, in *FlashPromotionSessionListReq, opts ...grpc.CallOption) (*FlashPromotionSessionListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionSessionList(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionSessionUpdate(ctx context.Context, in *FlashPromotionSessionUpdateReq, opts ...grpc.CallOption) (*FlashPromotionSessionUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionSessionUpdate(ctx, in, opts...)
}

func (m *defaultSms) FlashPromotionSessionDelete(ctx context.Context, in *FlashPromotionSessionDeleteReq, opts ...grpc.CallOption) (*FlashPromotionSessionDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.FlashPromotionSessionDelete(ctx, in, opts...)
}

func (m *defaultSms) HomeAdvertiseAdd(ctx context.Context, in *HomeAdvertiseAddReq, opts ...grpc.CallOption) (*HomeAdvertiseAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeAdvertiseAdd(ctx, in, opts...)
}

func (m *defaultSms) HomeAdvertiseList(ctx context.Context, in *HomeAdvertiseListReq, opts ...grpc.CallOption) (*HomeAdvertiseListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeAdvertiseList(ctx, in, opts...)
}

func (m *defaultSms) HomeAdvertiseUpdate(ctx context.Context, in *HomeAdvertiseUpdateReq, opts ...grpc.CallOption) (*HomeAdvertiseUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeAdvertiseUpdate(ctx, in, opts...)
}

func (m *defaultSms) HomeAdvertiseDelete(ctx context.Context, in *HomeAdvertiseDeleteReq, opts ...grpc.CallOption) (*HomeAdvertiseDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeAdvertiseDelete(ctx, in, opts...)
}

func (m *defaultSms) HomeBrandAdd(ctx context.Context, in *HomeBrandAddReq, opts ...grpc.CallOption) (*HomeBrandAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeBrandAdd(ctx, in, opts...)
}

func (m *defaultSms) HomeBrandList(ctx context.Context, in *HomeBrandListReq, opts ...grpc.CallOption) (*HomeBrandListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeBrandList(ctx, in, opts...)
}

func (m *defaultSms) HomeBrandUpdate(ctx context.Context, in *HomeBrandUpdateReq, opts ...grpc.CallOption) (*HomeBrandUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeBrandUpdate(ctx, in, opts...)
}

func (m *defaultSms) HomeBrandDelete(ctx context.Context, in *HomeBrandDeleteReq, opts ...grpc.CallOption) (*HomeBrandDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeBrandDelete(ctx, in, opts...)
}

func (m *defaultSms) HomeNewProductAdd(ctx context.Context, in *HomeNewProductAddReq, opts ...grpc.CallOption) (*HomeNewProductAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeNewProductAdd(ctx, in, opts...)
}

func (m *defaultSms) HomeNewProductList(ctx context.Context, in *HomeNewProductListReq, opts ...grpc.CallOption) (*HomeNewProductListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeNewProductList(ctx, in, opts...)
}

func (m *defaultSms) HomeNewProductUpdate(ctx context.Context, in *HomeNewProductUpdateReq, opts ...grpc.CallOption) (*HomeNewProductUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeNewProductUpdate(ctx, in, opts...)
}

func (m *defaultSms) HomeNewProductDelete(ctx context.Context, in *HomeNewProductDeleteReq, opts ...grpc.CallOption) (*HomeNewProductDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeNewProductDelete(ctx, in, opts...)
}

func (m *defaultSms) HomeRecommendProductAdd(ctx context.Context, in *HomeRecommendProductAddReq, opts ...grpc.CallOption) (*HomeRecommendProductAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeRecommendProductAdd(ctx, in, opts...)
}

func (m *defaultSms) HomeRecommendProductList(ctx context.Context, in *HomeRecommendProductListReq, opts ...grpc.CallOption) (*HomeRecommendProductListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeRecommendProductList(ctx, in, opts...)
}

func (m *defaultSms) HomeRecommendProductUpdate(ctx context.Context, in *HomeRecommendProductUpdateReq, opts ...grpc.CallOption) (*HomeRecommendProductUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeRecommendProductUpdate(ctx, in, opts...)
}

func (m *defaultSms) HomeRecommendProductDelete(ctx context.Context, in *HomeRecommendProductDeleteReq, opts ...grpc.CallOption) (*HomeRecommendProductDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeRecommendProductDelete(ctx, in, opts...)
}

func (m *defaultSms) HomeRecommendSubjectAdd(ctx context.Context, in *HomeRecommendSubjectAddReq, opts ...grpc.CallOption) (*HomeRecommendSubjectAddResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeRecommendSubjectAdd(ctx, in, opts...)
}

func (m *defaultSms) HomeRecommendSubjectList(ctx context.Context, in *HomeRecommendSubjectListReq, opts ...grpc.CallOption) (*HomeRecommendSubjectListResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeRecommendSubjectList(ctx, in, opts...)
}

func (m *defaultSms) HomeRecommendSubjectUpdate(ctx context.Context, in *HomeRecommendSubjectUpdateReq, opts ...grpc.CallOption) (*HomeRecommendSubjectUpdateResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeRecommendSubjectUpdate(ctx, in, opts...)
}

func (m *defaultSms) HomeRecommendSubjectDelete(ctx context.Context, in *HomeRecommendSubjectDeleteReq, opts ...grpc.CallOption) (*HomeRecommendSubjectDeleteResp, error) {
	client := smsclient.NewSmsClient(m.cli.Conn())
	return client.HomeRecommendSubjectDelete(ctx, in, opts...)
}
