// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package productattributecategoryservice

import (
	"context"

	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AlbumAddReq                                = pmsclient.AlbumAddReq
	AlbumAddResp                               = pmsclient.AlbumAddResp
	AlbumDeleteReq                             = pmsclient.AlbumDeleteReq
	AlbumDeleteResp                            = pmsclient.AlbumDeleteResp
	AlbumListData                              = pmsclient.AlbumListData
	AlbumListReq                               = pmsclient.AlbumListReq
	AlbumListResp                              = pmsclient.AlbumListResp
	AlbumPicAddReq                             = pmsclient.AlbumPicAddReq
	AlbumPicAddResp                            = pmsclient.AlbumPicAddResp
	AlbumPicDeleteReq                          = pmsclient.AlbumPicDeleteReq
	AlbumPicDeleteResp                         = pmsclient.AlbumPicDeleteResp
	AlbumPicListData                           = pmsclient.AlbumPicListData
	AlbumPicListReq                            = pmsclient.AlbumPicListReq
	AlbumPicListResp                           = pmsclient.AlbumPicListResp
	AlbumPicUpdateReq                          = pmsclient.AlbumPicUpdateReq
	AlbumPicUpdateResp                         = pmsclient.AlbumPicUpdateResp
	AlbumUpdateReq                             = pmsclient.AlbumUpdateReq
	AlbumUpdateResp                            = pmsclient.AlbumUpdateResp
	BrandAddReq                                = pmsclient.BrandAddReq
	BrandAddResp                               = pmsclient.BrandAddResp
	BrandDeleteReq                             = pmsclient.BrandDeleteReq
	BrandDeleteResp                            = pmsclient.BrandDeleteResp
	BrandListByIdsReq                          = pmsclient.BrandListByIdsReq
	BrandListData                              = pmsclient.BrandListData
	BrandListReq                               = pmsclient.BrandListReq
	BrandListResp                              = pmsclient.BrandListResp
	BrandUpdateReq                             = pmsclient.BrandUpdateReq
	BrandUpdateResp                            = pmsclient.BrandUpdateResp
	CommentAddReq                              = pmsclient.CommentAddReq
	CommentAddResp                             = pmsclient.CommentAddResp
	CommentDeleteReq                           = pmsclient.CommentDeleteReq
	CommentDeleteResp                          = pmsclient.CommentDeleteResp
	CommentListData                            = pmsclient.CommentListData
	CommentListReq                             = pmsclient.CommentListReq
	CommentListResp                            = pmsclient.CommentListResp
	CommentReplayAddReq                        = pmsclient.CommentReplayAddReq
	CommentReplayAddResp                       = pmsclient.CommentReplayAddResp
	CommentReplayDeleteReq                     = pmsclient.CommentReplayDeleteReq
	CommentReplayDeleteResp                    = pmsclient.CommentReplayDeleteResp
	CommentReplayListData                      = pmsclient.CommentReplayListData
	CommentReplayListReq                       = pmsclient.CommentReplayListReq
	CommentReplayListResp                      = pmsclient.CommentReplayListResp
	CommentReplayUpdateReq                     = pmsclient.CommentReplayUpdateReq
	CommentReplayUpdateResp                    = pmsclient.CommentReplayUpdateResp
	CommentUpdateReq                           = pmsclient.CommentUpdateReq
	CommentUpdateResp                          = pmsclient.CommentUpdateResp
	FeightTemplateAddReq                       = pmsclient.FeightTemplateAddReq
	FeightTemplateAddResp                      = pmsclient.FeightTemplateAddResp
	FeightTemplateDeleteReq                    = pmsclient.FeightTemplateDeleteReq
	FeightTemplateDeleteResp                   = pmsclient.FeightTemplateDeleteResp
	FeightTemplateListData                     = pmsclient.FeightTemplateListData
	FeightTemplateListReq                      = pmsclient.FeightTemplateListReq
	FeightTemplateListResp                     = pmsclient.FeightTemplateListResp
	FeightTemplateUpdateReq                    = pmsclient.FeightTemplateUpdateReq
	FeightTemplateUpdateResp                   = pmsclient.FeightTemplateUpdateResp
	MemberPriceAddReq                          = pmsclient.MemberPriceAddReq
	MemberPriceAddResp                         = pmsclient.MemberPriceAddResp
	MemberPriceDeleteReq                       = pmsclient.MemberPriceDeleteReq
	MemberPriceDeleteResp                      = pmsclient.MemberPriceDeleteResp
	MemberPriceList                            = pmsclient.MemberPriceList
	MemberPriceListData                        = pmsclient.MemberPriceListData
	MemberPriceListReq                         = pmsclient.MemberPriceListReq
	MemberPriceListResp                        = pmsclient.MemberPriceListResp
	MemberPriceUpdateReq                       = pmsclient.MemberPriceUpdateReq
	MemberPriceUpdateResp                      = pmsclient.MemberPriceUpdateResp
	ProductAddReq                              = pmsclient.ProductAddReq
	ProductAddResp                             = pmsclient.ProductAddResp
	ProductAttributeAddReq                     = pmsclient.ProductAttributeAddReq
	ProductAttributeAddResp                    = pmsclient.ProductAttributeAddResp
	ProductAttributeCategoryAddReq             = pmsclient.ProductAttributeCategoryAddReq
	ProductAttributeCategoryAddResp            = pmsclient.ProductAttributeCategoryAddResp
	ProductAttributeCategoryDeleteReq          = pmsclient.ProductAttributeCategoryDeleteReq
	ProductAttributeCategoryDeleteResp         = pmsclient.ProductAttributeCategoryDeleteResp
	ProductAttributeCategoryListData           = pmsclient.ProductAttributeCategoryListData
	ProductAttributeCategoryListReq            = pmsclient.ProductAttributeCategoryListReq
	ProductAttributeCategoryListResp           = pmsclient.ProductAttributeCategoryListResp
	ProductAttributeCategoryUpdateReq          = pmsclient.ProductAttributeCategoryUpdateReq
	ProductAttributeCategoryUpdateResp         = pmsclient.ProductAttributeCategoryUpdateResp
	ProductAttributeDeleteReq                  = pmsclient.ProductAttributeDeleteReq
	ProductAttributeDeleteResp                 = pmsclient.ProductAttributeDeleteResp
	ProductAttributeListData                   = pmsclient.ProductAttributeListData
	ProductAttributeListReq                    = pmsclient.ProductAttributeListReq
	ProductAttributeListResp                   = pmsclient.ProductAttributeListResp
	ProductAttributeUpdateReq                  = pmsclient.ProductAttributeUpdateReq
	ProductAttributeUpdateResp                 = pmsclient.ProductAttributeUpdateResp
	ProductAttributeValueAddReq                = pmsclient.ProductAttributeValueAddReq
	ProductAttributeValueAddResp               = pmsclient.ProductAttributeValueAddResp
	ProductAttributeValueDeleteReq             = pmsclient.ProductAttributeValueDeleteReq
	ProductAttributeValueDeleteResp            = pmsclient.ProductAttributeValueDeleteResp
	ProductAttributeValueList                  = pmsclient.ProductAttributeValueList
	ProductAttributeValueListData              = pmsclient.ProductAttributeValueListData
	ProductAttributeValueListReq               = pmsclient.ProductAttributeValueListReq
	ProductAttributeValueListResp              = pmsclient.ProductAttributeValueListResp
	ProductAttributeValueUpdateReq             = pmsclient.ProductAttributeValueUpdateReq
	ProductAttributeValueUpdateResp            = pmsclient.ProductAttributeValueUpdateResp
	ProductByIdsReq                            = pmsclient.ProductByIdsReq
	ProductCategoryAddReq                      = pmsclient.ProductCategoryAddReq
	ProductCategoryAddResp                     = pmsclient.ProductCategoryAddResp
	ProductCategoryAttributeRelationAddReq     = pmsclient.ProductCategoryAttributeRelationAddReq
	ProductCategoryAttributeRelationAddResp    = pmsclient.ProductCategoryAttributeRelationAddResp
	ProductCategoryAttributeRelationDeleteReq  = pmsclient.ProductCategoryAttributeRelationDeleteReq
	ProductCategoryAttributeRelationDeleteResp = pmsclient.ProductCategoryAttributeRelationDeleteResp
	ProductCategoryAttributeRelationListData   = pmsclient.ProductCategoryAttributeRelationListData
	ProductCategoryAttributeRelationListReq    = pmsclient.ProductCategoryAttributeRelationListReq
	ProductCategoryAttributeRelationListResp   = pmsclient.ProductCategoryAttributeRelationListResp
	ProductCategoryAttributeRelationUpdateReq  = pmsclient.ProductCategoryAttributeRelationUpdateReq
	ProductCategoryAttributeRelationUpdateResp = pmsclient.ProductCategoryAttributeRelationUpdateResp
	ProductCategoryDeleteReq                   = pmsclient.ProductCategoryDeleteReq
	ProductCategoryDeleteResp                  = pmsclient.ProductCategoryDeleteResp
	ProductCategoryListData                    = pmsclient.ProductCategoryListData
	ProductCategoryListReq                     = pmsclient.ProductCategoryListReq
	ProductCategoryListResp                    = pmsclient.ProductCategoryListResp
	ProductCategoryUpdateReq                   = pmsclient.ProductCategoryUpdateReq
	ProductCategoryUpdateResp                  = pmsclient.ProductCategoryUpdateResp
	ProductDeleteReq                           = pmsclient.ProductDeleteReq
	ProductDeleteResp                          = pmsclient.ProductDeleteResp
	ProductDetailByIdReq                       = pmsclient.ProductDetailByIdReq
	ProductDetailByIdResp                      = pmsclient.ProductDetailByIdResp
	ProductFullReductionAddReq                 = pmsclient.ProductFullReductionAddReq
	ProductFullReductionAddResp                = pmsclient.ProductFullReductionAddResp
	ProductFullReductionDeleteReq              = pmsclient.ProductFullReductionDeleteReq
	ProductFullReductionDeleteResp             = pmsclient.ProductFullReductionDeleteResp
	ProductFullReductionList                   = pmsclient.ProductFullReductionList
	ProductFullReductionListData               = pmsclient.ProductFullReductionListData
	ProductFullReductionListReq                = pmsclient.ProductFullReductionListReq
	ProductFullReductionListResp               = pmsclient.ProductFullReductionListResp
	ProductFullReductionUpdateReq              = pmsclient.ProductFullReductionUpdateReq
	ProductFullReductionUpdateResp             = pmsclient.ProductFullReductionUpdateResp
	ProductLadderAddReq                        = pmsclient.ProductLadderAddReq
	ProductLadderAddResp                       = pmsclient.ProductLadderAddResp
	ProductLadderDeleteReq                     = pmsclient.ProductLadderDeleteReq
	ProductLadderDeleteResp                    = pmsclient.ProductLadderDeleteResp
	ProductLadderList                          = pmsclient.ProductLadderList
	ProductLadderListData                      = pmsclient.ProductLadderListData
	ProductLadderListReq                       = pmsclient.ProductLadderListReq
	ProductLadderListResp                      = pmsclient.ProductLadderListResp
	ProductLadderUpdateReq                     = pmsclient.ProductLadderUpdateReq
	ProductLadderUpdateResp                    = pmsclient.ProductLadderUpdateResp
	ProductListData                            = pmsclient.ProductListData
	ProductListReq                             = pmsclient.ProductListReq
	ProductListResp                            = pmsclient.ProductListResp
	ProductOperateLogAddReq                    = pmsclient.ProductOperateLogAddReq
	ProductOperateLogAddResp                   = pmsclient.ProductOperateLogAddResp
	ProductOperateLogDeleteReq                 = pmsclient.ProductOperateLogDeleteReq
	ProductOperateLogDeleteResp                = pmsclient.ProductOperateLogDeleteResp
	ProductOperateLogListData                  = pmsclient.ProductOperateLogListData
	ProductOperateLogListReq                   = pmsclient.ProductOperateLogListReq
	ProductOperateLogListResp                  = pmsclient.ProductOperateLogListResp
	ProductOperateLogUpdateReq                 = pmsclient.ProductOperateLogUpdateReq
	ProductOperateLogUpdateResp                = pmsclient.ProductOperateLogUpdateResp
	ProductUpdateReq                           = pmsclient.ProductUpdateReq
	ProductUpdateResp                          = pmsclient.ProductUpdateResp
	ProductVertifyRecordAddReq                 = pmsclient.ProductVertifyRecordAddReq
	ProductVertifyRecordAddResp                = pmsclient.ProductVertifyRecordAddResp
	ProductVertifyRecordDeleteReq              = pmsclient.ProductVertifyRecordDeleteReq
	ProductVertifyRecordDeleteResp             = pmsclient.ProductVertifyRecordDeleteResp
	ProductVertifyRecordListData               = pmsclient.ProductVertifyRecordListData
	ProductVertifyRecordListReq                = pmsclient.ProductVertifyRecordListReq
	ProductVertifyRecordListResp               = pmsclient.ProductVertifyRecordListResp
	ProductVertifyRecordUpdateReq              = pmsclient.ProductVertifyRecordUpdateReq
	ProductVertifyRecordUpdateResp             = pmsclient.ProductVertifyRecordUpdateResp
	ReleaseSkuStockLockData                    = pmsclient.ReleaseSkuStockLockData
	ReleaseSkuStockLockReq                     = pmsclient.ReleaseSkuStockLockReq
	ReleaseSkuStockLockResp                    = pmsclient.ReleaseSkuStockLockResp
	SkuStockAddReq                             = pmsclient.SkuStockAddReq
	SkuStockAddResp                            = pmsclient.SkuStockAddResp
	SkuStockDeleteReq                          = pmsclient.SkuStockDeleteReq
	SkuStockDeleteResp                         = pmsclient.SkuStockDeleteResp
	SkuStockList                               = pmsclient.SkuStockList
	SkuStockListData                           = pmsclient.SkuStockListData
	SkuStockListReq                            = pmsclient.SkuStockListReq
	SkuStockListResp                           = pmsclient.SkuStockListResp
	SkuStockUpdateReq                          = pmsclient.SkuStockUpdateReq
	SkuStockUpdateResp                         = pmsclient.SkuStockUpdateResp

	ProductAttributeCategoryService interface {
		ProductAttributeCategoryAdd(ctx context.Context, in *ProductAttributeCategoryAddReq, opts ...grpc.CallOption) (*ProductAttributeCategoryAddResp, error)
		ProductAttributeCategoryList(ctx context.Context, in *ProductAttributeCategoryListReq, opts ...grpc.CallOption) (*ProductAttributeCategoryListResp, error)
		ProductAttributeCategoryUpdate(ctx context.Context, in *ProductAttributeCategoryUpdateReq, opts ...grpc.CallOption) (*ProductAttributeCategoryUpdateResp, error)
		ProductAttributeCategoryDelete(ctx context.Context, in *ProductAttributeCategoryDeleteReq, opts ...grpc.CallOption) (*ProductAttributeCategoryDeleteResp, error)
	}

	defaultProductAttributeCategoryService struct {
		cli zrpc.Client
	}
)

func NewProductAttributeCategoryService(cli zrpc.Client) ProductAttributeCategoryService {
	return &defaultProductAttributeCategoryService{
		cli: cli,
	}
}

func (m *defaultProductAttributeCategoryService) ProductAttributeCategoryAdd(ctx context.Context, in *ProductAttributeCategoryAddReq, opts ...grpc.CallOption) (*ProductAttributeCategoryAddResp, error) {
	client := pmsclient.NewProductAttributeCategoryServiceClient(m.cli.Conn())
	return client.ProductAttributeCategoryAdd(ctx, in, opts...)
}

func (m *defaultProductAttributeCategoryService) ProductAttributeCategoryList(ctx context.Context, in *ProductAttributeCategoryListReq, opts ...grpc.CallOption) (*ProductAttributeCategoryListResp, error) {
	client := pmsclient.NewProductAttributeCategoryServiceClient(m.cli.Conn())
	return client.ProductAttributeCategoryList(ctx, in, opts...)
}

func (m *defaultProductAttributeCategoryService) ProductAttributeCategoryUpdate(ctx context.Context, in *ProductAttributeCategoryUpdateReq, opts ...grpc.CallOption) (*ProductAttributeCategoryUpdateResp, error) {
	client := pmsclient.NewProductAttributeCategoryServiceClient(m.cli.Conn())
	return client.ProductAttributeCategoryUpdate(ctx, in, opts...)
}

func (m *defaultProductAttributeCategoryService) ProductAttributeCategoryDelete(ctx context.Context, in *ProductAttributeCategoryDeleteReq, opts ...grpc.CallOption) (*ProductAttributeCategoryDeleteResp, error) {
	client := pmsclient.NewProductAttributeCategoryServiceClient(m.cli.Conn())
	return client.ProductAttributeCategoryDelete(ctx, in, opts...)
}
