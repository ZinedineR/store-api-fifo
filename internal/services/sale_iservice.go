package service

import (
	"boiler-plate-clean/internal/model"
	"boiler-plate-clean/pkg/exception"
	"context"
)

type SaleService interface {
	// CRUD operations for Sale
	Create(ctx context.Context, req *model.CreateSaleReq) (*model.CreateSaleRes, *exception.Exception)
	GetProfitReport(
		ctx context.Context, req *model.ProfitReportReq,
	) (*model.ProfitReportRes, *exception.Exception)
	//GetById(ctx context.Context, req *model.GetSaleByIdReq) (*model.GetSaleByIdRes, *exception.Exception)
	//GetList(ctx context.Context, req *model.GetListSaleReq) (*model.GetListSaleRes, *exception.Exception)
	//Update(ctx context.Context, req *model.UpdateSaleReq) (*model.UpdateSaleRes, *exception.Exception)
	//Delete(ctx context.Context, req *model.DeleteSaleReq) (*model.DeleteSaleRes, *exception.Exception)
}
