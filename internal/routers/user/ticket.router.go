package user

import (
	"github.com/Youknow2509/go-ecommerce/internal/controller/ticket"
	"github.com/gin-gonic/gin"
)

type TicketRouter struct{}

func (tr *TicketRouter) InitTicketRouter(Router *gin.RouterGroup) {
	// public router
	ticketRouterPublic := Router.Group("/ticket")
	{
		// ticketRouterPublic.GET("/search")
		ticketRouterPublic.GET("/item/:id", ticket.TicketItem.GetTicketItemById)
		ticketRouterPublic.POST("/item/decrease", ticket.TicketItem.DecreaseTicketItem)
		ticketRouterPublic.POST("/item/release/enable", ticket.TicketItem.ReleaseTicketItemEnable)
		ticketRouterPublic.POST("/item/release/disable", ticket.TicketItem.ReleaseTicketItemDisable)
		ticketRouterPublic.POST("/item/release/decrease", ticket.TicketItem.DecreaseTicketItemRelease)
	}
	// private router

}
