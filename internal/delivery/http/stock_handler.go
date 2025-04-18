package http

import (
	_ "boiler-plate-clean/internal/delivery/http/response"
	"boiler-plate-clean/internal/model"
	service "boiler-plate-clean/internal/services"
	"github.com/gin-gonic/gin"
)

type StockHTTPHandler struct {
	Handler
	StockService service.StockService
}

func NewStockHTTPHandler(example service.StockService) *StockHTTPHandler {
	return &StockHTTPHandler{
		StockService: example,
	}
}

// Create godoc
// @Summary Create a new stock
// @Description Creates a new stock in the system
// @Tags Stock
// @Accept json
// @Produce json
// @Param stock body model.CreateStockReq true "Create Stock Request"
// @Success 200 {object} response.DataResponse{data=model.CreateStockRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /stock [post]
func (h StockHTTPHandler) Create(ctx *gin.Context) {
	request := model.CreateStockReq{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	resp, errException := h.StockService.Create(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, resp)
}

// Find godoc
// @Summary Get all stock
// @Description Retrieves a list of all stock with optional filters, pagination, and sorting
// @Tags Stock
// @Accept json
// @Produce json
// @Param pageSize query string false "Number of items per page"
// @Param page query string false "Page number"
// @Param filter query string false "Filter rules<br><br>### Rules Filter<br>rule:<br>  * {Name of Field}:{value}:{Symbol}<br><br>Symbols:<br>  * eq (=)<br>  * lt (<)<br>  * gt (>)<br>  * lte (<=)<br>  * gte (>=)<br>  * in ( in)<br>  * like (like)<br><br>Field list:<br>  * id<br>  * price" default(id:1:eq)
// @Param sort query string false "Sort rules:<br><br>### Rules Sort<br>rule:<br>  * {Name of Field}:{Symbol}<br><br>Symbols:<br>  * asc<br>  * desc<br><br>Field list:<br>  * id<br>  * price" default(id:desc)
// @Success 200 {object} response.DataResponse{data=model.GetListStockRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /stock [get]
func (h StockHTTPHandler) Find(ctx *gin.Context) {
	var req model.GetListStockReq
	var err error
	req.Page, req.Order, req.Filter, err = h.ParsePaginationParams(ctx)
	if err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	result, errException := h.StockService.GetList(ctx, &req)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, result)
}

// FindOne godoc
// @Summary Get stock details
// @Description Retrieves the details of a specific stock by ID
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path int true "stock ID"
// @Success 200 {object} response.DataResponse{data=model.GetStockByIdRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /stock/{id} [get]
func (h StockHTTPHandler) FindOne(ctx *gin.Context) {
	var req model.GetStockByIdReq
	req.ID = ctx.GetInt("id")
	result, errException := h.StockService.GetById(ctx, &req)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, result)
}

// Update godoc
// @Summary Update an existing stock
// @Description Updates stock details
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path int true "stock ID"
// @Param stock body model.UpdateStockReq true "Update Stock Request"
// @Success 200 {object} response.DataResponse{data=model.UpdateStockRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /stock/{id} [put]
func (h StockHTTPHandler) Update(ctx *gin.Context) {
	// Get Info
	//idParam := ctx.Param("id")
	request := model.UpdateStockReq{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	request.ID = ctx.GetInt("id")

	resp, errException := h.StockService.Update(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, resp)
}

// Delete godoc
// @Summary Delete a stock
// @Description Deletes a stock by ID
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path int true "stock ID"
// @Success 200 {object} response.DataResponse{data=model.DeleteStockRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /stock/{id} [delete]
func (h StockHTTPHandler) Delete(ctx *gin.Context) {
	var req model.DeleteStockReq
	req.ID = ctx.GetInt("id")
	if _, errException := h.StockService.Delete(ctx, &req); errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.SuccessMessageJSON(ctx, "entire stock has been deleted")
}
