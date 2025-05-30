package ticket

import (
	"strconv"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// manager controller Ticket Item
var TicketItem = new(cTicketItem)

type cTicketItem struct{}

// @Summary      Get ticket item by id
// @Description  Get ticket item by id
// @Tags         ticket management
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Ticket ID"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /v1/ticket/item/{id} [get]
func (p *cTicketItem) GetTicketItemById(ctx *gin.Context) {
	// get id from url
	id := ctx.Param("id")
	if id == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "id is required")
		return
	}
	// convert id to int
	ticketItemId, err := strconv.Atoi(id)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "id must be a number")
		return
	}
	// call implementation
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, ticketItemId)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, ticketItem)
}

// @Summary      Decrease ticket item by id
// @Description  Decrease ticket item by id
// @Tags         ticket management
// @Accept       json
// @Produce      json
// @Param        input body model.TicketItemDecreaseRequest true "Ticket Item Decrease Request"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrResponseData
// @Router       /v1/ticket/item/decrease [post]
func (p *cTicketItem) DecreaseTicketItem(ctx *gin.Context) {
	// bind request body
	var req model.TicketItemDecreaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "invalid request body")
		return
	}
	// call implementation
	err := service.TicketItem().DecreaseTicketItem(ctx, req.TicketId, req.TicketInventory, req.Quantity)
	if err != nil {
		global.Logger.Warn("decrease ticket item error", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, "decrease ticket item successfully")
}
