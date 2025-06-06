package routers

import (
	"github.com/Youknow2509/go-ecommerce/internal/routers/cron"
	"github.com/Youknow2509/go-ecommerce/internal/routers/manage"
	"github.com/Youknow2509/go-ecommerce/internal/routers/product"
	"github.com/Youknow2509/go-ecommerce/internal/routers/prometheus"
	"github.com/Youknow2509/go-ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User       user.UserRouterGroup
	Manage     manage.ManageRouterGroup
	Prometheus prometheus.PrometheusRouter
	Product    product.ProductGroupRouter
	Cron       cron.CronJobRouterGroup
}

var RouterGroupApp = new(RouterGroup)
