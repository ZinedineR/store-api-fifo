package http

import (
	_ "boiler-plate-clean/internal/delivery/http/response"
	"boiler-plate-clean/internal/model"
	service "boiler-plate-clean/internal/services"
	"github.com/gin-gonic/gin"
)

type SaleHTTPHandler struct {
	Handler
	SaleService service.SaleService
}

func NewSaleHTTPHandler(example service.SaleService) *SaleHTTPHandler {
	return &SaleHTTPHandler{
		SaleService: example,
	}
}

// Create godoc
// @Summary Create a new sale
// @Description Creates a new sale and calculates HPP using FIFO
// @Tags Sale
// @Accept json
// @Produce json
// @Param sale body model.CreateSaleReq true "Create Sale Request"
// @Success 200 {object} response.DataResponse{data=model.CreateSaleRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /sale [post]
func (h SaleHTTPHandler) Create(ctx *gin.Context) {
	var request model.CreateSaleReq
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}

	resp, errException := h.SaleService.Create(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, resp)
}

// GetProfitReport godoc
// @Summary Get monthly profit report
// @Description Retrieves total sales, total HPP, and profit (laba) for a specific month and year
// @Tags Sale
// @Accept json
// @Produce json
// @Param month query int true "Month (1-12)"
// @Param year query int true "Year (e.g. 2023)"
// @Success 200 {object} response.DataResponse{data=model.ProfitReportRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /sale/report [get]
func (h SaleHTTPHandler) GetProfitReport(ctx *gin.Context) {
	var req model.ProfitReportReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}

	result, errException := h.SaleService.GetProfitReport(ctx, &req)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, result)
}

//
//// Find godoc
//// @Summary Get all stock
//// @Description Retrieves a list of all stock with optional filters, pagination, and sorting
//// @Tags Sale
//// @Accept json
//// @Produce json
//// @Param pageSize query string false "Number of items per page"
//// @Param page query string false "Page number"
//// @Param filter query string false "Filter rules<br><br>### Rules Filter<br>rule:<br>  * {Name of Field}:{value}:{Symbol}<br><br>Symbols:<br>  * eq (=)<br>  * lt (<)<br>  * gt (>)<br>  * lte (<=)<br>  * gte (>=)<br>  * in ( in)<br>  * like (like)<br><br>Field list:<br>  * id<br>  * price" default(id:1:eq)
//// @Param sort query string false "Sort rules:<br><br>### Rules Sort<br>rule:<br>  * {Name of Field}:{Symbol}<br><br>Symbols:<br>  * asc<br>  * desc<br><br>Field list:<br>  * id<br>  * price" default(id:desc)
//// @Success 200 {object} response.DataResponse{data=model.GetListSaleRes} "success"
//// @Failure 400 {object} response.DataResponse "error"
//// @Router /stock [get]
//func (h SaleHTTPHandler) Find(ctx *gin.Context) {
//	var req model.GetListSaleReq
//	var err error
//	req.Page, req.Order, req.Filter, err = h.ParsePaginationParams(ctx)
//	if err != nil {
//		h.BadRequestJSON(ctx, err.Error())
//		return
//	}
//	result, errException := h.SaleService.GetList(ctx, &req)
//	if errException != nil {
//		h.ExceptionJSON(ctx, errException)
//		return
//	}
//
//	h.DataJSON(ctx, result)
//}
//
//// FindOne godoc
//// @Summary Get stock details
//// @Description Retrieves the details of a specific stock by ID
//// @Tags Sale
//// @Accept json
//// @Produce json
//// @Param id path int true "stock ID"
//// @Success 200 {object} response.DataResponse{data=model.GetSaleByIdRes} "success"
//// @Failure 400 {object} response.DataResponse "error"
//// @Router /stock/{id} [get]
//func (h SaleHTTPHandler) FindOne(ctx *gin.Context) {
//	var req model.GetSaleByIdReq
//	req.ID = ctx.GetInt("id")
//	result, errException := h.SaleService.GetById(ctx, &req)
//	if errException != nil {
//		h.ExceptionJSON(ctx, errException)
//		return
//	}
//
//	h.DataJSON(ctx, result)
//}
//
//// Update godoc
//// @Summary Update an existing stock
//// @Description Updates stock details
//// @Tags Sale
//// @Accept json
//// @Produce json
//// @Param id path int true "stock ID"
//// @Param stock body model.UpdateSaleReq true "Update Sale Request"
//// @Success 200 {object} response.DataResponse{data=model.UpdateSaleRes} "success"
//// @Failure 400 {object} response.DataResponse "error"
//// @Router /stock/{id} [put]
//func (h SaleHTTPHandler) Update(ctx *gin.Context) {
//	// Get Info
//	//idParam := ctx.Param("id")
//	request := model.UpdateSaleReq{}
//	if err := ctx.ShouldBindJSON(&request); err != nil {
//		h.BadRequestJSON(ctx, err.Error())
//		return
//	}
//	request.ID = ctx.GetInt("id")
//
//	resp, errException := h.SaleService.Update(ctx, &request)
//	if errException != nil {
//		h.ExceptionJSON(ctx, errException)
//		return
//	}
//
//	h.DataJSON(ctx, resp)
//}
//
//// Delete godoc
//// @Summary Delete a stock
//// @Description Deletes a stock by ID
//// @Tags Sale
//// @Accept json
//// @Produce json
//// @Param id path int true "stock ID"
//// @Success 200 {object} response.DataResponse{data=model.DeleteSaleRes} "success"
//// @Failure 400 {object} response.DataResponse "error"
//// @Router /stock/{id} [delete]
//func (h SaleHTTPHandler) Delete(ctx *gin.Context) {
//	var req model.DeleteSaleReq
//	req.ID = ctx.GetInt("id")
//	if _, errException := h.SaleService.Delete(ctx, &req); errException != nil {
//		h.ExceptionJSON(ctx, errException)
//		return
//	}
//
//	h.SuccessMessageJSON(ctx, "entire stock has been deleted")
//}
