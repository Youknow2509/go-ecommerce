package product

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/repo/create"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Product = new(cProduct)

type cProduct struct {
}

// @Summary      Create clothing product
// @Description  Create clothing product
// @Tags         product management
// @Accept       json
// @Produce      json
// @Param        payload body model.ClothingProductRequest true "clothing payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /v1/clothing/create [post]
func (cU *cProduct) CreateClothing(c *gin.Context) {
    var request model.ClothingProductRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        response.ErrorResponse(c, response.ErrCodeBindVerifyInput, err.Error())
        return
    }
    // validate request
    if ok, message := request.Validate(); !ok {
		response.ErrorResponse(c, response.ErrSchemaValidate, message)
	}
    
    err := create.CreateProduct(model.ClothingCollection).InsertProduct(request)
    if err != nil {
        global.Logger.Error("create product error", zap.Error(err))
        response.ErrorResponse(c, response.ErrCodeCreateProductError, err.Error())
        return
    }
    
    response.SuccessResponse(c, response.ErrCodeSuccess, nil)
}
