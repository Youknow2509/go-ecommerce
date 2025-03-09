package initialize

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/internal/service/impl"
)

// initialize services
func InitServiceInterface() {
	q := database.New(global.Mdbc)
    service.InitUserLogin(impl.NewSUserLogin(q))
	service.InitTicketItem(impl.NewTicketItemImpl(q))
	// ...
}