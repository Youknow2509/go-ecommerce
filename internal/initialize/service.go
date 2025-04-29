package initialize

import (
	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/internal/service/impl"
)

// initialize services
func InitServiceInterface() {
	// initialize database mysql
	q := database.New(global.Mdbc)
	// initialize local cache
	restrettoCache := impl.NewRestrettoCache()
	service.InitLocalCache(restrettoCache)
	// initialize distributed cache
	redisCache := impl.NewRedisCache(global.Rdb)
	service.InitRedisCache(redisCache)
	//
    service.InitUserLogin(impl.NewSUserLogin(q))
	service.InitTicketItem(impl.NewTicketItemImpl(
		q,
		service.GetLocalCache(),
	))
	// additional service initializations can be added here
}