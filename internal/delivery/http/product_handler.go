package http

import (
	_ "boiler-plate-clean/internal/delivery/http/response"
	"boiler-plate-clean/internal/model"
	service "boiler-plate-clean/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHTTPHandler struct {
	Handler
	ProductService service.ProductService
}

func NewProductHTTPHandler(example service.ProductService) *ProductHTTPHandler {
	return &ProductHTTPHandler{
		ProductService: example,
	}
}

// Create godoc
// @Summary Create a new product
// @Description Creates a new product in the catalog
// @Tags Product
// @Accept json
// @Produce json
// @Param product body model.CreateProductReq true "Create Product Request"
// @Success 200 {object} response.DataResponse{data=model.CreateProductRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /product [post]
func (h ProductHTTPHandler) Create(ctx *gin.Context) {
	request := model.CreateProductReq{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	resp, errException := h.ProductService.Create(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, resp)
}

// Find godoc
// @Summary Get all product
// @Description Retrieves a list of all product with optional filters, pagination, and sorting
// @Tags Product
// @Accept json
// @Produce json
// @Param pageSize query string false "Number of items per page"
// @Param page query string false "Page number"
// @Param filter query string false "Filter rules<br><br>### Rules Filter<br>rule:<br>  * {Name of Field}:{value}:{Symbol}<br><br>Symbols:<br>  * eq (=)<br>  * lt (<)<br>  * gt (>)<br>  * lte (<=)<br>  * gte (>=)<br>  * in ( in)<br>  * like (like)<br><br>Field list:<br>  * id<br>  * name<br>  * created_at"
// @Param sort query string false "Sort rules:<br><br>### Rules Sort<br>rule:<br>  * {Name of Field}:{Symbol}<br><br>Symbols:<br>  * asc<br>  * desc<br><br>Field list:<br>  * id<br>  * name<br>  * created_at" default(id:desc)
// @Success 200 {object} response.DataResponse{data=model.GetListProductRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /product [get]
func (h ProductHTTPHandler) Find(ctx *gin.Context) {
	var req model.GetListProductReq
	var err error
	req.Page, req.Order, req.Filter, err = h.ParsePaginationParams(ctx)
	if err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	result, errException := h.ProductService.GetList(ctx, &req)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, result)
}

// FindOne godoc
// @Summary Get product details
// @Description Retrieves the details of a specific product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "product ID"
// @Success 200 {object} response.DataResponse{data=model.GetProductByIdRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /product/{id} [get]
func (h ProductHTTPHandler) FindOne(ctx *gin.Context) {
	var req model.GetProductByIdReq
	req.ID = h.GetIntId(ctx)
	result, errException := h.ProductService.GetById(ctx, &req)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, result)
}

// Update godoc
// @Summary Update an existing product
// @Description Updates product details
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "product ID"
// @Param product body model.UpdateProductReq true "Update Product Request"
// @Success 200 {object} response.DataResponse{data=model.UpdateProductRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /product/{id} [put]
func (h ProductHTTPHandler) Update(ctx *gin.Context) {
	// Get Info
	//idParam := ctx.Param("id")
	request := model.UpdateProductReq{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	request.ID = ctx.GetInt("id")

	resp, errException := h.ProductService.Update(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, resp)
}

// Delete godoc
// @Summary Delete a product
// @Description Deletes a product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "product ID"
// @Success 200 {object} response.DataResponse{data=model.DeleteProductRes} "success"
// @Failure 400 {object} response.DataResponse "error"
// @Router /product/{id} [delete]
func (h ProductHTTPHandler) Delete(ctx *gin.Context) {
	var req model.DeleteProductReq
	req.ID = ctx.GetInt("id")
	if _, errException := h.ProductService.Delete(ctx, &req); errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.SuccessMessageJSON(ctx, "product has been deleted")
}
